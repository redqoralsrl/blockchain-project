package postgresql

import (
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
)

type AttributesStorage struct {
	db *sql.DB
}

func NewAttributesStorage(db *sql.DB) *AttributesStorage {
	return &AttributesStorage{db}
}

func (s *AttributesStorage) CreateErc721(queryRower postgresql.Query, input *attributes.CreateErc721AttributesInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	if len(input.AttributeList) > 0 {
		for _, attribute := range input.AttributeList {
			query := `
				insert into attributes (chain_id, contract_id, erc721_id, trait_type, value, display_type)
				values ($1, $2, $3, $4, $5, $6)
				returning id;
			`

			var displayType sql.NullString
			if attribute.DisplayType != "" {
				displayType.String = attribute.DisplayType
				displayType.Valid = true
			} else {
				displayType.Valid = false
			}

			_, err := queryRower.Exec(
				query,
				attribute.ChainId,
				attribute.ContractId,
				attribute.Erc721Id,
				attribute.TraitType,
				attribute.Value,
				displayType,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *AttributesStorage) CreateErc1155(queryRower postgresql.Query, input *attributes.CreateErc1155AttributesInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	if len(input.AttributeList) > 0 {
		for _, attribute := range input.AttributeList {
			query := `
				insert into attributes (chain_id, contract_id, erc1155_id, trait_type, value, display_type)
				values ($1, $2, $3, $4, $5, $6)
				returning id;
			`

			var displayType sql.NullString
			if attribute.DisplayType != "" {
				displayType.String = attribute.DisplayType
				displayType.Valid = true
			} else {
				displayType.Valid = false
			}

			_, err := queryRower.Exec(
				query,
				attribute.ChainId,
				attribute.ContractId,
				attribute.Erc1155Id,
				attribute.TraitType,
				attribute.Value,
				displayType,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
