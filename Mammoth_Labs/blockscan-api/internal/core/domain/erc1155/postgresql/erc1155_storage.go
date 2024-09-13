package postgresql

import (
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/core/domain/erc1155"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	cm "github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
)

type Erc1155Storage struct {
	db *sql.DB
}

func NewErc1155Storage(db *sql.DB) *Erc1155Storage {
	return &Erc1155Storage{db}
}

func (s *Erc1155Storage) MoveErc1155(queryRower postgresql.Query, input *erc1155.MoveErc1155Input) error {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		update contract
		set is_erc_20 = false, is_erc_721 = false, is_erc_1155 = true
		where id = $1;
	`
	_, err := queryRower.Exec(query, input.ContractId)
	if err != nil {
		return err
	}

	for i, tokenId := range input.Erc1155TokenId {
		value := input.Erc1155Value[i]
		url := input.Erc1155Url[i]

		var erc1155Id int
		tokenIdStr := tokenId.String()
		valueStr := value.String()

		getQuery := `
			select erc1155.id from erc1155
			where contract_id = $1 and chain_id = $2 and token_id = $3;
		`

		err := queryRower.QueryRow(
			getQuery,
			input.ContractId,
			input.ChainId,
			tokenIdStr,
		).Scan(
			&erc1155Id,
		)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("error checking erc1155 existence: %w", err)
		}

		if errors.Is(err, sql.ErrNoRows) {
			createQuery := `
				insert into erc1155 (chain_id, contract_id, token_id, amount, url)
				values ($1, $2, $3, $4, $5)
				returning id;
			`

			err = queryRower.QueryRow(
				createQuery,
				input.ChainId,
				input.ContractId,
				tokenIdStr,
				valueStr,
				url,
			).Scan(
				&erc1155Id,
			)
			if err != nil {
				return err
			}

			createOwnerQuery := `
				insert into erc1155_owner (chain_id, erc1155_id, wallet_id, amount)
				values($1, $2, (select id from wallet where lower(address) = lower($3)), $4)
				returning id;
			`

			_, err := queryRower.Exec(
				createOwnerQuery,
				input.ChainId,
				erc1155Id,
				input.To,
				valueStr,
			)
			if err != nil {
				return err
			}
		} else if erc1155Id != 0 {
			// from 삭제
			if cm.HexToAddress(input.From) == cm.HexToAddress("0x0000000000000000000000000000000000000000") {
				// mint 일 시 to 만 추가하고 erc1155 amount 값 증가
				updateQuery := `
					insert into erc1155_owner (chain_id, erc1155_id, wallet_id, amount)
					values($1, $2, (select id from wallet where lower(address) = lower($3)), $4)
					on conflict (erc1155_id, wallet_id)
            		do update set amount = erc1155_owner.amount + excluded.amount;
				`
				_, err := queryRower.Exec(
					updateQuery,
					input.ChainId,
					erc1155Id,
					input.To,
					valueStr,
				)
				if err != nil {
					return err
				}

				updateErc1155Query := `
					update erc1155
					set amount = amount + $1
					where id = $2;
				`
				_, err = queryRower.Exec(
					updateErc1155Query,
					valueStr,
					erc1155Id,
				)
				if err != nil {
					return err
				}

			} else {
				updateFromQuery := `
					update erc1155_owner
					set amount = amount - $1
					where erc1155_id = $2 and wallet_id = (select id from wallet where lower(address) = lower($3));
				`
				_, err := queryRower.Exec(
					updateFromQuery,
					valueStr,
					erc1155Id,
					input.From,
				)
				if err != nil {
					return err
				}

				deleteFromQuery := `
					delete from erc1155_owner
            		where erc1155_id = $1 and wallet_id = (select id from wallet where lower(address) = lower($2)) and amount <= 0;
				`
				_, err = queryRower.Exec(
					deleteFromQuery,
					erc1155Id,
					input.From,
				)
				if err != nil {
					return err
				}

				updateQuery := `
					insert into erc1155_owner (chain_id, erc1155_id, wallet_id, amount)
					values($1, $2, (select id from wallet where lower(address) = lower($3)), $4)
					on conflict (erc1155_id, wallet_id)
            		do update set amount = erc1155_owner.amount + excluded.amount
				`
				_, err = queryRower.Exec(
					updateQuery,
					input.ChainId,
					erc1155Id,
					input.To,
					valueStr,
				)

				if err != nil {
					return err
				}
			}
		}

	}

	return err
}

func (s *Erc1155Storage) GetEmptyUrlErc1155List() ([]*erc1155.GetEmptyUrlErc1155, error) {
	query := `
		select id, contract_id, url, chain_id from erc1155
		where is_undefined_metadata = false and erc1155.image_url is null
		order by random()
		limit 100 offset 0;
	`

	rows, err := s.db.Query(query)

	if errors.Is(err, sql.ErrNoRows) {
		return []*erc1155.GetEmptyUrlErc1155{}, nil
	} else if err != nil {
		return nil, err
	}

	var erc1155List []*erc1155.GetEmptyUrlErc1155
	for rows.Next() {
		var erc1155Single erc1155.GetEmptyUrlErc1155
		err := rows.Scan(&erc1155Single.Erc1155ID, &erc1155Single.ContractId, &erc1155Single.Url, &erc1155Single.ChainId)
		if err != nil {
			return nil, err
		}
		erc1155List = append(erc1155List, &erc1155Single)
	}

	if err := rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*erc1155.GetEmptyUrlErc1155{}, nil
		}
		return nil, err
	}

	return erc1155List, nil
}

func (s *Erc1155Storage) UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) error {
	if queryRower != nil {
		queryRower = s.db
	}

	query := `
		update erc1155
		set is_undefined_metadata = true
		where id = $1;
	`

	_, err := queryRower.Exec(
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Erc1155Storage) Update(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	var setText strings.Builder
	var values []interface{}

	if input.ImageUrl != "" {
		setText.WriteString("image_url = $")
		setText.WriteString(strconv.Itoa(len(values) + 1))
		values = append(values, input.ImageUrl)
	}

	if input.Name != "" {
		if setText.Len() > 0 {
			setText.WriteString(", ")
		}
		setText.WriteString("name = $")
		setText.WriteString(strconv.Itoa(len(values) + 1))
		values = append(values, input.Name)
	}

	if input.Description != "" {
		if setText.Len() > 0 {
			setText.WriteString(", ")
		}
		setText.WriteString("description = $")
		setText.WriteString(strconv.Itoa(len(values) + 1))
		values = append(values, input.Description)
	}

	if input.ExternalUrl != "" {
		if setText.Len() > 0 {
			setText.WriteString(", ")
		}
		setText.WriteString("external_url = $")
		setText.WriteString(strconv.Itoa(len(values) + 1))
		values = append(values, input.ExternalUrl)
	}

	var query string
	if setText.Len() > 0 {
		query = fmt.Sprintf(`
			update erc1155
			set is_undefined_metadata = true, %s
			where id = $%d;
		`, setText.String(), len(values)+1)
	}

	values = append(values, strconv.Itoa(input.Erc1155Id))

	_, err := queryRower.Exec(
		query,
		values...,
	)

	if err != nil {
		return err
	}

	return nil
}
