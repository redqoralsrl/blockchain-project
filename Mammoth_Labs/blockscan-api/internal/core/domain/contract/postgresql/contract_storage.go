package postgresql

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/contract"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"unicode/utf8"
)

type ContractStorage struct {
	db *sql.DB
}

func NewContractStorage(db *sql.DB) *ContractStorage {
	return &ContractStorage{db}
}

func (s *ContractStorage) Get(queryRower postgresql.Query, input *contract.GetContractInput) (*contract.Contract, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		select id, chain_id, transaction_create_id, timestamp, hash, name, symbol, decimals, is_erc_20, is_erc_721, is_erc_1155, total_supply, creator
		from contract
		where lower(hash)=lower($1) and chain_id = $2;
	`

	var contract *contract.Contract

	err := queryRower.QueryRow(
		query,
		input.Hash,
		input.ChainId,
	).Scan(
		&contract.ID,
		&contract.ChainId,
		&contract.TransactionCreateId,
		&contract.Timestamp,
		&contract.Hash,
		&contract.Name,
		&contract.Symbol,
		&contract.Decimals,
		&contract.IsErc20,
		&contract.IsErc721,
		&contract.IsErc1155,
		&contract.TotalSupply,
		&contract.Creator,
	)

	if err != nil {
		return nil, err
	}

	return contract, nil
}

func (s *ContractStorage) Create(queryRower postgresql.Query, input *contract.CreateContractInput) (*contract.Contract, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	if !utf8.ValidString(input.Name) {
		input.Name = ""
	}
	if !utf8.ValidString(input.Symbol) {
		input.Symbol = ""
	}

	query := `
		insert into contract (chain_id, transaction_create_id, timestamp, hash, name, symbol, decimals, is_erc_20, is_erc_721, is_erc_1155, total_supply, creator)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		returning id, chain_id, transaction_create_id, timestamp, hash, name, symbol, decimals, is_erc_20, is_erc_721, is_erc_1155, total_supply, creator;
	`

	totalSupplyStr := input.TotalSupply.String()
	var getTotalSupply sql.NullString

	contractData := &contract.Contract{}
	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.TransactionCreateId,
		input.Timestamp,
		input.Hash,
		input.Name,
		input.Symbol,
		input.Decimals,
		input.IsErc20,
		input.IsErc721,
		input.IsErc1155,
		totalSupplyStr,
		input.Creator,
	).Scan(
		&contractData.ID,
		&contractData.ChainId,
		&contractData.TransactionCreateId,
		&contractData.Timestamp,
		&contractData.Hash,
		&contractData.Name,
		&contractData.Symbol,
		&contractData.Decimals,
		&contractData.IsErc20,
		&contractData.IsErc721,
		&contractData.IsErc1155,
		&getTotalSupply,
		&contractData.Creator,
	)

	if getTotalSupply.Valid {
		totalSupplyInt := new(big.Int)
		totalSupplyInt, ok := totalSupplyInt.SetString(getTotalSupply.String, 10)
		if !ok {
			return nil, fmt.Errorf("totalSupplyInt failed to convert number string to big.Int")
		}
		contractData.TotalSupply = *totalSupplyInt
	}

	if err != nil {
		return nil, err
	}

	return contractData, nil
}

func (s *ContractStorage) GetId(queryRower postgresql.Query, input *contract.GetContractIdInput) (int, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
	   select id from contract where lower(hash) = lower($1) and chain_id = $2;
	`

	var contractId int
	err := queryRower.QueryRow(
		query,
		input.ContractAddress,
		input.ChainId,
	).Scan(
		&contractId,
	)

	if err != nil {
		return -1, err
	}

	return contractId, err
}

func (s *ContractStorage) GetType(queryRower postgresql.Query, input *contract.GetContractTypeInput) (int, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		select is_erc_20, is_erc_721, is_erc_1155 from contract
		where lower(hash) = lower($1) and chain_id = $2;
	`

	var contractInfoData = &contract.GetContractTypeInfo{
		IsErc20:   false,
		IsErc721:  false,
		IsErc1155: false,
	}
	err := queryRower.QueryRow(
		query,
		input.Hash,
		input.ChainId,
	).Scan(
		&contractInfoData.IsErc20,
		&contractInfoData.IsErc721,
		&contractInfoData.IsErc1155,
	)

	if err != nil {
		// error
		return -1, err
	}

	if contractInfoData.IsErc20 {
		// if erc 20
		return 20, nil
	} else if contractInfoData.IsErc721 {
		// if erc721
		return 721, nil
	} else if contractInfoData.IsErc1155 {
		// if erc1155
		return 1155, nil
	}

	// or nothing
	return -1, nil
}

func (s *ContractStorage) UpdateType(queryRower postgresql.Query, input *contract.UpdateContractTypeInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	var setText string
	if input.Type == 20 {
		setText = "is_erc_20 = true"
	} else if input.Type == 721 {
		setText = "is_erc_721 = true"
	} else if input.Type == 1155 {
		setText = "is_erc_1155 = true"
	} else {
		return fmt.Errorf("not support input type")
	}

	query := fmt.Sprintf(`
		update contract
		set %s
		where lower(hash) = lower($1) and chain_id = $2;
	`, setText)

	_, err := queryRower.Exec(
		query,
		input.Hash,
		input.ChainId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *ContractStorage) UpdateAllVolume(queryRower postgresql.Query, input *contract.UpdateContractAllVolumeInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	getQuery := `
		select all_volume from contract
		where lower(hash) = lower($1) and chain_id = $2;
	`

	var getVolumeStr string
	err := queryRower.QueryRow(
		getQuery,
		input.Hash,
		input.ChainId,
	).Scan(
		&getVolumeStr,
	)

	if err != nil {
		return err
	}

	if getVolumeStr != "" {
		volumeInt := new(big.Int)
		volumeInt, ok := volumeInt.SetString(getVolumeStr, 10)
		if !ok {
			return fmt.Errorf("volumeInt failed to convert number string to big.Int")
		}

		volumeInt.Add(volumeInt, &input.Volume)

		query := `
			update contract
			set all_volume = $1
			where lower(hash) = lower($2) and chain_id = $3;
		`
		_, err = queryRower.Exec(
			query,
			volumeInt.String(),
			input.Hash,
			input.ChainId,
		)

		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("contract all_volume data is nil")
	}

	return nil
}

func (s *ContractStorage) GetWalletNFTsByCollection(input *contract.GetContractWalletNFTsByCollectionInput) ([]*contract.GetContractWalletNFTsByCollectionData, error) {
	limit := 10
	if input.Take > 100 {
		limit = 100
	} else if input.Take > 0 {
		limit = input.Take
	}

	offset := 0
	if input.Skip > 0 {
		offset = input.Skip
	}

	query := `
		with wallet_id as (
			select id from wallet
			where lower(address) = lower($1)
		),
		nft_counts_per_contract as (
		    select
		        e.contract_id,
		        count(distinct e.token_id) as unique_tokens_count,
		        sum(e.amount) as tokens_amount
			from
			    erc721 e
			inner join
				wallet_id w on e.wallet_id = w.id
			group by
			    e.contract_id
			union all
			select
				eo.contract_id,
				count(distinct eo.token_id) as unique_tokens_count,
				sum(eo.amount) as tokens_amount
			from
			    erc1155_owner e
			inner join
			    erc1155 eo on eo.id = e.erc1155_id
			inner join
				wallet_id w on e.wallet_id = w.id
			group by eo.contract_id
		)
		select
		    c.id,
			c.chain_id,
			c.hash,
			c.name,
			c.symbol,
			c.aws_logo_image,
			c.aws_banner_image,
			c.description,
			case
			    when c.is_erc_721 = true then '721'
			    when c.is_erc_1155 = true then '1155'
			    else 'UNKNOWN'
			end as erc_type,
			w.address as wallet_address,
			w.nick_name as wallet_nickname,
			w.profile as wallet_profile,
			json_agg(
				json_build_object(
					'token_id', e.token_id::text,
					'amount', e.amount,
					'url', e.url,
					'image_url', e.image_url,
					'name', e.name,
					'description', e.description,
					'aws_image_url', e.aws_image_url
				) order by e.token_id asc
			) as erc_list,
			n.unique_tokens_count,
			n.tokens_amount
		from contract c
		inner join wallet w on lower(w.address) = lower($1)
		left join lateral (
		    select
		        e.token_id as token_id,
    			e.amount::text as amount,
				coalesce(e.url, '') as url,
				coalesce(e.image_url, '') as image_url,
				coalesce(e.name, '') as name,
				coalesce(e.description, '') as description,
				coalesce(e.aws_image_url, '') as aws_image_url,
				e.wallet_id
		    from
		        erc721 e
		    where
		        e.contract_id = c.id and
		        e.wallet_id = (select id from wallet_id)
		    union all
		    select
		        e.token_id as token_id,
    			e.amount::text as amount,
				coalesce(e.url, '') as url,
				coalesce(e.image_url, '') as image_url,
				coalesce(e.name, '') as name,
				coalesce(e.description, '') as description,
				coalesce(e.aws_image_url, '') as aws_image_url,
				eo.wallet_id
		    from
		        erc1155_owner eo
		    inner join
		        erc1155 e on e.id = eo.erc1155_id
		    where
		        e.contract_id = c.id and
		        eo.wallet_id = (select id from wallet_id)  
		    order by token_id asc
		    limit 10
		) e on true
		left join nft_counts_per_contract n on n.contract_id = c.id
		where e.wallet_id = (select id from wallet_id) and c.chain_id = $2
		group by 
		    c.id, w.address, w.nick_name, 
		    w.profile, c.chain_id, c.hash, 
		    c.name, c.symbol, c.aws_logo_image, 
		    c.aws_banner_image, c.description, 
		    n.unique_tokens_count, n.tokens_amount
		limit $3 offset $4;
	`

	rows, err := s.db.Query(
		query,
		input.WalletAddress,
		input.ChainId,
		limit,
		offset,
	)
	if err != nil {
		return []*contract.GetContractWalletNFTsByCollectionData{}, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var contractList []*contract.GetContractWalletNFTsByCollectionData
	for rows.Next() {

		var contractSingle contract.GetContractWalletNFTsByCollectionData

		var getName sql.NullString
		var getSymbol sql.NullString
		var getAwsLogoImage sql.NullString
		var getAwsBannerImage sql.NullString
		var getWalletNickname sql.NullString
		var getWalletProfile sql.NullString
		var getErcList sql.NullString
		var getDescription sql.NullString
		var getUniqueTokensCount sql.NullString
		var getTokensAmount sql.NullString

		err := rows.Scan(
			&contractSingle.ID,
			&contractSingle.ChainId,
			&contractSingle.Hash,
			&getName,
			&getSymbol,
			&getAwsLogoImage,
			&getAwsBannerImage,
			&getDescription,
			&contractSingle.ErcType,
			&contractSingle.WalletAddress,
			&getWalletNickname,
			&getWalletProfile,
			&getErcList,
			&getUniqueTokensCount,
			&getTokensAmount,
		)
		if err != nil {
			return nil, err
		}

		contractSingle.Name = utils.GetString(getName)
		contractSingle.Symbol = utils.GetString(getSymbol)
		contractSingle.AwsLogoImage = utils.GetString(getAwsLogoImage)
		contractSingle.AwsBannerImage = utils.GetString(getAwsBannerImage)
		contractSingle.Description = utils.GetString(getDescription)
		contractSingle.WalletNickname = utils.GetString(getWalletNickname)
		contractSingle.WalletProfile = utils.GetString(getWalletProfile)

		if getErcList.Valid {
			var ercList []contract.ErcList
			err := json.Unmarshal([]byte(getErcList.String), &ercList)
			if err != nil {
				return nil, fmt.Errorf("ercList failed to convert json.Unmarshal")
			}
			contractSingle.ErcList = ercList
		} else {
			contractSingle.ErcList = []contract.ErcList{}
		}

		contractSingle.UniqueTokensCount = utils.GetString(getUniqueTokensCount)

		if getTokensAmount.Valid {
			contractSingle.TokensAmount = getTokensAmount.String
		} else {
			contractSingle.TokensAmount = "0"
		}

		contractList = append(contractList, &contractSingle)
	}

	return contractList, nil
}

func (s *ContractStorage) GetCollectionNFTsForWallet(input *contract.GetContractCollectionNFTsForWalletInput) (*contract.GetContractCollectionNFTsForWalletData, error) {
	limit := 10
	if input.Take > 100 {
		limit = 100
	} else if input.Take > 0 {
		limit = input.Take
	}

	offset := 0
	if input.Skip > 0 {
		offset = input.Skip
	}

	query := `
		with wallet_id as (
			select id from wallet
			where lower(address) = lower($1)
		)
		select
			c.id,
			c.chain_id,
			c.hash,
			c.name,
			c.symbol,
			c.aws_logo_image,
			c.aws_banner_image,
			c.description,
			case
			    when c.is_erc_721 = true then '721'
			    when c.is_erc_1155 = true then '1155'
			    else 'UNKNOWN'
			end as erc_type,
			w.address as wallet_address,
			w.nick_name as wallet_nickname,
			w.profile as wallet_profile,
			json_agg(
				json_build_object(
					'token_id', e.token_id::text,
					'amount', e.amount,
					'url', e.url,
					'image_url', e.image_url,
					'name',e.name,
					'description', e.description,
					'aws_image_url', e.aws_image_url
				) order by e.token_id asc
			) as erc_list
		from contract c
		inner join wallet w on lower(w.address) = lower($1)
		left join lateral (
		    select
				e.token_id as token_id,
    			e.amount::text as amount,
				coalesce(e.url, '') as url,
				coalesce(e.image_url, '') as image_url,
				coalesce(e.name, '') as name,
				coalesce(e.description, '') as description,
				coalesce(e.aws_image_url, '') as aws_image_url,
				e.wallet_id
			from
				erc721 e
			where
				e.contract_id = c.id and
				e.wallet_id = (select id from wallet_id)
			union all 
		    select
		        e.token_id as token_id,
    			e.amount::text as amount,
				coalesce(e.url, '') as url,
				coalesce(e.image_url, '') as image_url,
				coalesce(e.name, '') as name,
				coalesce(e.description, '') as description,
				coalesce(e.aws_image_url, '') as aws_image_url,
				eo.wallet_id
		    from
		        erc1155_owner eo
		    inner join
		        erc1155 e on e.id = eo.erc1155_id
		    where
		        e.contract_id = c.id and
		        eo.wallet_id = (select id from wallet_id)  
			order by token_id asc
			limit $3 offset $4
		) e on true
		where e.wallet_id = (select id from wallet_id) and c.chain_id = $2 and lower(c.hash) = lower($5)
		group by c.id, w.address, w.nick_name, w.profile, c.chain_id, c.hash, c.name, c.symbol, c.aws_logo_image, c.aws_banner_image, c.description;
	`

	var contractData contract.GetContractCollectionNFTsForWalletData

	var getName sql.NullString
	var getSymbol sql.NullString
	var getAwsLogoImage sql.NullString
	var getAwsBannerImage sql.NullString
	var getWalletNickname sql.NullString
	var getWalletProfile sql.NullString
	var getErcList sql.NullString
	var getDescription sql.NullString

	err := s.db.QueryRow(
		query,
		input.WalletAddress,
		input.ChainId,
		limit,
		offset,
		input.Hash,
	).Scan(
		&contractData.ID,
		&contractData.ChainId,
		&contractData.Hash,
		&getName,
		&getSymbol,
		&getAwsLogoImage,
		&getAwsBannerImage,
		&getDescription,
		&contractData.ErcType,
		&contractData.WalletAddress,
		&getWalletNickname,
		&getWalletProfile,
		&getErcList,
	)

	if err != nil {
		return nil, err
	}

	contractData.Name = utils.GetString(getName)
	contractData.Symbol = utils.GetString(getSymbol)
	contractData.AwsLogoImage = utils.GetString(getAwsLogoImage)
	contractData.AwsBannerImage = utils.GetString(getAwsBannerImage)
	contractData.Description = utils.GetString(getDescription)
	contractData.WalletNickname = utils.GetString(getWalletNickname)
	contractData.WalletProfile = utils.GetString(getWalletProfile)

	if getErcList.Valid {
		var ercList []contract.ErcList
		err := json.Unmarshal([]byte(getErcList.String), &ercList)
		if err != nil {
			return nil, fmt.Errorf("ercList failed to convert json.Unmarshal")
		}
		contractData.ErcList = ercList
	} else {
		contractData.ErcList = []contract.ErcList{}
	}

	return &contractData, nil
}
