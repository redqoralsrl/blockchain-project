package postgresql

import (
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/core/domain/erc721"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Erc721Storage struct {
	db *sql.DB
}

func NewErc721Storage(db *sql.DB) *Erc721Storage {
	return &Erc721Storage{db}
}

func (s *Erc721Storage) MoveErc721(queryRower postgresql.Query, input *erc721.MoveErc721Input) error {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		update contract
		set is_erc_20 = false, is_erc_721 = true
		where id = $1;
	`
	_, err := queryRower.Exec(query, input.ContractId)
	if err != nil {
		return err
	}

	getQuery := `
		select erc721.id from erc721
		inner join wallet on erc721.wallet_id = wallet.id
		where contract_id = $1 and chain_id = $2 and token_id = $3;
	`

	tokenIdString := input.TokenId.String()
	var erc721Id int
	err = queryRower.QueryRow(
		getQuery,
		input.ContractId,
		input.ChainId,
		tokenIdString,
	).Scan(
		&erc721Id,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("Movenft function err : %w", err)
	}

	if errors.Is(err, sql.ErrNoRows) {
		createQuery := `
			insert into erc721 (chain_id, contract_id, wallet_id, token_id, amount, url)
			values ($1, $2, (select id from wallet where lower(address) = lower($3)), $4, $5, $6)
			returning id;
		`

		tokenIdStr := input.TokenId.String()
		amountStr := input.Amount.String()

		err = nil
		_, err = queryRower.Exec(createQuery, input.ChainId, input.ContractId, input.To, tokenIdStr, amountStr, input.Url)
		if err != nil {
			return err
		}
	} else if erc721Id != 0 {
		updateQuery := `
			update erc721
			set wallet_id = (select id from wallet where address = $1)
			where contract_id = $2 and chain_id = $3 and token_id = $4;
		`

		tokenIdStr := input.TokenId.String()

		err = nil
		_, err = queryRower.Exec(updateQuery, input.To, input.ContractId, input.ChainId, tokenIdStr)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *Erc721Storage) GetEmptyUrlErc721List() ([]*erc721.GetEmptyUrlErc721, error) {
	query := `
		select id, contract_id, url, chain_id from erc721
		where is_undefined_metadata = false and erc721.image_url is null
		order by random()
		limit 100 offset 0;
	`

	rows, err := s.db.Query(query)

	if errors.Is(err, sql.ErrNoRows) {
		return []*erc721.GetEmptyUrlErc721{}, nil
	} else if err != nil {
		return nil, err
	}

	var erc721List []*erc721.GetEmptyUrlErc721
	for rows.Next() {
		var erc721Single erc721.GetEmptyUrlErc721
		err := rows.Scan(&erc721Single.Erc721ID, &erc721Single.ContractId, &erc721Single.Url, &erc721Single.ChainId)
		if err != nil {
			return nil, err
		}
		erc721List = append(erc721List, &erc721Single)
	}

	if err := rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*erc721.GetEmptyUrlErc721{}, nil
		}
		return nil, err
	}

	return erc721List, nil
}

func (s *Erc721Storage) UpdateIsUndefinedMetaData(queryRower postgresql.Query, id int) error {
	if queryRower != nil {
		queryRower = s.db
	}

	query := `
		update erc721
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

func (s *Erc721Storage) Update(queryRower postgresql.Query, input *attributes.CreateErc721AttributesInput) error {
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
			update erc721
			set is_undefined_metadata = true, %s
			where id = $%d;
		`, setText.String(), len(values)+1)
	}

	values = append(values, strconv.Itoa(input.Erc721Id))

	_, err := queryRower.Exec(
		query,
		values...,
	)

	if err != nil {
		return err
	}

	return nil
}
