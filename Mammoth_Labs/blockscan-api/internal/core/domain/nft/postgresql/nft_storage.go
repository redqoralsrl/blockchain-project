package postgresql

import (
	"blockscan-go/internal/core/domain/nft"
	"database/sql"
	"encoding/json"
	"fmt"
)

type NftStorage struct {
	db *sql.DB
}

func NewNftStorage(db *sql.DB) *NftStorage {
	return &NftStorage{db}
}

func (s *NftStorage) Get(input *nft.GetNftInput) ([]*nft.GetNftData, error) {
	limit := 10
	if input.Take > 100 {
		limit = 100
	} else if input.Take > 0 {
		limit = input.Take
	}
	offset := 0
	if input.Skip >= 0 {
		offset = input.Skip
	}

	query := `
		select
			nft.id,
			nft.chain_id,
			contract_id,
			nft.contract_hash,
			coalesce(nft.contract_name, '') as contract_name,
			coalesce(nft.contract_symbol, '') as contract_symbol,
			coalesce(nft.contract_description, '') as contract_description,
			coalesce(nft.aws_logo_image, '') as aws_logo_image,
			coalesce(nft.aws_banner_image, '') as aws_banner_image,
			nft.token_id::text as token_id,
			nft.amount::text as amount,
			nft.type,
			coalesce(nft.url, '') as url,
			coalesce(nft.image_url, '') as image_url,
			coalesce(nft.nft_name, '') as nft_name,
			coalesce(nft.nft_description, '') as nft_description,
			nft.attributes_array
		from (
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				e.amount,
				'erc721' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', a.trait_type,
						'value', a.value
					) 
				) filter (where a.id is not null) as attributes_array
			from erc721 e
			inner join contract c on e.contract_id = c.id
			inner join wallet w on e.wallet_id = w.id            
			left join attributes a on e.id = a.erc721_id
			where lower(w.address) = lower($1) and e.chain_id = $2
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, e.amount,
				e.url, e.image_url, e.name, e.description
			
			union all 
			
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				eo.amount,
				'erc1155' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', a.trait_type,
						'value', a.value
					) 
				) filter (where a.id is not null) as attributes_array
			from erc1155 e
			inner join contract c on e.contract_id = c.id
			inner join erc1155_owner eo on eo.erc1155_id = e.id and eo.wallet_id = (select id from wallet where lower(address) = lower($1))
			left join attributes a on e.id = a.erc1155_id
			where eo.wallet_id = (select id from wallet where lower(address) = lower($1)) and e.chain_id = $2
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, eo.amount,
				e.url, e.image_url, e.name, e.description
		) as nft
		order by contract_id desc, token_id::int desc
		limit $3 offset $4;
	`

	rows, err := s.db.Query(query,
		input.WalletAddress,
		input.ChainId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var nftList []*nft.GetNftData
	for rows.Next() {
		var nftData nft.GetNftData

		var getAttributesArray sql.NullString

		if errs := rows.Scan(
			&nftData.ID, &nftData.ChainId,
			&nftData.ContractId, &nftData.ContractHash, &nftData.ContractName, &nftData.ContractSymbol,
			&nftData.ContractDescription, &nftData.AwsLogoImage, &nftData.AwsBannerImage,
			&nftData.TokenId, &nftData.Amount, &nftData.Type, &nftData.Url, &nftData.ImageUrl, &nftData.NftName,
			&nftData.NftDescription, &getAttributesArray,
		); errs != nil {
			return nil, errs
		}

		if getAttributesArray.Valid {
			var attributesArray []nft.Attributes
			err := json.Unmarshal([]byte(getAttributesArray.String), &attributesArray)
			if err != nil {
				return nil, fmt.Errorf("attributesArray failed to convert json.Unmarshal")
			}
			nftData.AttributesArray = attributesArray
		} else {
			nftData.AttributesArray = []nft.Attributes{}
		}

		nftList = append(nftList, &nftData)
	}
	if errs := rows.Err(); errs != nil {
		return nil, errs
	}

	return nftList, nil
}

func (s *NftStorage) GetDetail(input *nft.GetNftDetailInput) (*nft.GetNFtDetailData, error) {
	query := `
		with contract_info as (
			select id from contract
			where lower(hash) = lower($1) and chain_id = $2
		),
		erc721_attributes as (
			select a.trait_type, a.value, e.token_id from attributes a
			inner join erc721 e on a.erc721_id = e.id
			where e.contract_id = (select id from contract_info)
		),
		erc1155_attributes as (
			select a.trait_type, a.value, e.token_id from attributes a
			inner join erc1155 e on a.erc1155_id = e.id
			where e.contract_id = (select id from contract_info)
		),
		combined_attributes as (
			select * from erc721_attributes
			union all
			select * from erc1155_attributes
		),
		total_counts_per_trait as (
			select
				trait_type, 
				count(value) as total_count
			from combined_attributes
			group by trait_type
		),
		specific_token_attributes as (
			select trait_type, value from combined_attributes
			where token_id = $3
		),
		specific_token_attributes_counts as (
			select
				sa.trait_type,
				sa.value,
				count(sa.value) as count
			from combined_attributes a
			inner join specific_token_attributes sa on a.trait_type = sa.trait_type and a.value = sa.value
			group by sa.trait_type, sa.value
		),
		attributes_details as (
			select
				s.trait_type,
				s.value,
				tc.total_count,
				sac.count as specific_count,
				(sac.count::DECIMAL / tc.total_count) * 100 as percentage
			from specific_token_attributes_counts sac
			join total_counts_per_trait tc on sac.trait_type = tc.trait_type
			join specific_token_attributes s on sac.trait_type = s.trait_type and sac.value = s.value
		)
		select
			nft.id,
			nft.chain_id,
			contract_id,
			nft.contract_hash,
			coalesce(nft.contract_name, '') as contract_name,
			coalesce(nft.contract_symbol, '') as contract_symbol,
			coalesce(nft.contract_description, '') as contract_description,
			coalesce(nft.aws_logo_image, '') as aws_logo_image,
			coalesce(nft.aws_banner_image, '') as aws_banner_image,
			nft.token_id::text as token_id,
			nft.amount::text as amount,
			nft.type,
			coalesce(nft.url, '') as url,
			coalesce(nft.image_url, '') as image_url,
			coalesce(nft.nft_name, '') as nft_name,
			coalesce(nft.nft_description, '') as nft_description,
			nft.attributes_details
		from (
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				e.amount,
				'erc721' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', attributes_details.trait_type,
						'value', attributes_details.value,
						'total_count', attributes_details.total_count,
						'specific_count', attributes_details.specific_count,
						'percentage', attributes_details.percentage
					) 
				) as attributes_details
			from erc721 e
			inner join contract c on e.contract_id = c.id
			inner join wallet w on e.wallet_id = w.id            
    		left join attributes_details on attributes_details.value is not null
			where e.chain_id = $2 and e.token_id = $3 and e.contract_id = (select id from contract_info)
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, e.amount,
				e.url, e.image_url, e.name, e.description
			
			union all 
			
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				e.amount,
				'erc1155' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', attributes_details.trait_type,
						'value', attributes_details.value,
						'total_count', attributes_details.total_count,
						'specific_count', attributes_details.specific_count,
						'percentage', attributes_details.percentage
					) 
				) as attributes_details
			from erc1155 e
			inner join contract c on e.contract_id = c.id
   			left join attributes_details on attributes_details.value is not null
			where e.chain_id = $2 and e.token_id = $3 and e.contract_id = (select id from contract_info)
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, e.amount,
				e.url, e.image_url, e.name, e.description
		) as nft;
	`

	var nftData = &nft.GetNFtDetailData{}

	var getAttributes sql.NullString

	err := s.db.QueryRow(
		query,
		input.Hash,
		input.ChainId,
		input.TokenId,
	).Scan(
		&nftData.ID,
		&nftData.ChainId,
		&nftData.ContractId,
		&nftData.ContractHash,
		&nftData.ContractName,
		&nftData.ContractSymbol,
		&nftData.ContractDescription,
		&nftData.AwsLogoImage,
		&nftData.AwsBannerImage,
		&nftData.TokenId,
		&nftData.Amount,
		&nftData.Type,
		&nftData.Url,
		&nftData.ImageUrl,
		&nftData.NftName,
		&nftData.NftDescription,
		&getAttributes,
	)

	if err != nil {
		return nil, err
	}

	if getAttributes.Valid {
		var attributesArray []nft.AttributesPercent
		err := json.Unmarshal([]byte(getAttributes.String), &attributesArray)
		if err != nil {
			return nil, fmt.Errorf("attributesArray failed to convert json.Unmarshal")
		}
		nftData.AttributesArray = attributesArray
	} else {
		nftData.AttributesArray = []nft.AttributesPercent{}
	}

	return nftData, nil
}

func (s *NftStorage) GetDetailOfWallet(input *nft.GetNftDetailOfWalletInput) (*nft.GetNFtDetailData, error) {
	query := `
		with contract_info as (
			select id from contract
			where lower(hash) = lower($1) and chain_id = $2
		),
		erc721_attributes as (
			select a.trait_type, a.value, e.token_id from attributes a
			inner join erc721 e on a.erc721_id = e.id
			where e.contract_id = (select id from contract_info)
		),
		erc1155_attributes as (
			select a.trait_type, a.value, e.token_id from attributes a
			inner join erc1155 e on a.erc1155_id = e.id
			where e.contract_id = (select id from contract_info)
		),
		combined_attributes as (
			select * from erc721_attributes
			union all
			select * from erc1155_attributes
		),
		total_counts_per_trait as (
			select
				trait_type, 
				count(value) as total_count
			from combined_attributes
			group by trait_type
		),
		specific_token_attributes as (
			select trait_type, value from combined_attributes
			where token_id = $3
		),
		specific_token_attributes_counts as (
			select
				sa.trait_type,
				sa.value,
				count(sa.value) as count
			from combined_attributes a
			inner join specific_token_attributes sa on a.trait_type = sa.trait_type and a.value = sa.value
			group by sa.trait_type, sa.value
		),
		attributes_details as (
			select
				s.trait_type,
				s.value,
				tc.total_count,
				sac.count as specific_count,
				(sac.count::DECIMAL / tc.total_count) * 100 as percentage
			from specific_token_attributes_counts sac
			join total_counts_per_trait tc on sac.trait_type = tc.trait_type
			join specific_token_attributes s on sac.trait_type = s.trait_type and sac.value = s.value
		)
		select
			nft.id,
			nft.chain_id,
			contract_id,
			nft.contract_hash,
			coalesce(nft.contract_name, '') as contract_name,
			coalesce(nft.contract_symbol, '') as contract_symbol,
			coalesce(nft.contract_description, '') as contract_description,
			coalesce(nft.aws_logo_image, '') as aws_logo_image,
			coalesce(nft.aws_banner_image, '') as aws_banner_image,
			nft.token_id::text as token_id,
			nft.amount::text as amount,
			nft.type,
			coalesce(nft.url, '') as url,
			coalesce(nft.image_url, '') as image_url,
			coalesce(nft.nft_name, '') as nft_name,
			coalesce(nft.nft_description, '') as nft_description,
			nft.attributes_details
		from (
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				e.amount,
				'erc721' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', attributes_details.trait_type,
						'value', attributes_details.value,
						'total_count', attributes_details.total_count,
						'specific_count', attributes_details.specific_count,
						'percentage', attributes_details.percentage
					) 
				) as attributes_details
			from erc721 e
			inner join contract c on e.contract_id = c.id
			inner join wallet w on e.wallet_id = w.id            
    		left join attributes_details on attributes_details.value is not null
			where e.chain_id = $2 and e.token_id = $3 and e.contract_id = (select id from contract_info) and lower(w.address) = lower($4)
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, e.amount,
				e.url, e.image_url, e.name, e.description
			
			union all 
			
			select
				e.id,
				e.chain_id,
				c.id as contract_id,
				c.hash as contract_hash,
				c.name as contract_name,
				c.symbol as contract_symbol,
				c.description as contract_description,
				c.aws_logo_image,
				c.aws_banner_image,
				e.token_id,
				eo.amount,
				'erc1155' as type,
				e.url,
				e.image_url,
				e.name as nft_name,
				e.description as nft_description,
				json_agg(
					json_build_object(
						'trait_type', attributes_details.trait_type,
						'value', attributes_details.value,
						'total_count', attributes_details.total_count,
						'specific_count', attributes_details.specific_count,
						'percentage', attributes_details.percentage
					) 
				) as attributes_details
			from erc1155 e
			inner join contract c on e.contract_id = c.id
			inner join erc1155_owner eo on eo.erc1155_id = e.id and eo.wallet_id = (select id from wallet where lower(address) = lower($4))
   			left join attributes_details on attributes_details.value is not null
			where e.chain_id = $2 and e.token_id = $3 and e.contract_id = (select id from contract_info) and eo.wallet_id = (select id from wallet where lower(address) = lower($4))
			group by
				e.id, e.chain_id, c.id, c.hash, c.name, c.symbol, c.description,
				c.aws_logo_image, c.aws_banner_image, e.token_id, eo.amount,
				e.url, e.image_url, e.name, e.description
		) as nft;
	`

	var nftData = &nft.GetNFtDetailData{}

	var getAttributes sql.NullString

	err := s.db.QueryRow(
		query,
		input.Hash,
		input.ChainId,
		input.TokenId,
		input.WalletAddress,
	).Scan(
		&nftData.ID,
		&nftData.ChainId,
		&nftData.ContractId,
		&nftData.ContractHash,
		&nftData.ContractName,
		&nftData.ContractSymbol,
		&nftData.ContractDescription,
		&nftData.AwsLogoImage,
		&nftData.AwsBannerImage,
		&nftData.TokenId,
		&nftData.Amount,
		&nftData.Type,
		&nftData.Url,
		&nftData.ImageUrl,
		&nftData.NftName,
		&nftData.NftDescription,
		&getAttributes,
	)

	if err != nil {
		return nil, err
	}

	if getAttributes.Valid {
		var attributesArray []nft.AttributesPercent
		err := json.Unmarshal([]byte(getAttributes.String), &attributesArray)
		if err != nil {
			return nil, fmt.Errorf("attributesArray failed to convert json.Unmarshal")
		}
		nftData.AttributesArray = attributesArray
	} else {
		nftData.AttributesArray = []nft.AttributesPercent{}
	}

	return nftData, nil
}
