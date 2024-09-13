package postgresql

import (
	"blockscan-go/internal/core/domain/chain"
	"database/sql"
)

type ChainStorage struct {
	db *sql.DB
}

func NewChainStorage(db *sql.DB) *ChainStorage {
	return &ChainStorage{db}
}

func (s *ChainStorage) Get(chainId int) (*chain.Chain, error) {
	query := `
		select id, chain_id, contract_address, name, symbol, decimals, image_url, site, scan_site from chain
		where chain_id = $1 and contract_address = '0x';
	`

	var data = &chain.Chain{}

	err := s.db.QueryRow(
		query,
		chainId,
	).Scan(
		&data.ID,
		&data.ChainId,
		&data.ContractAddress,
		&data.Name,
		&data.Symbol,
		&data.Decimals,
		&data.ImageUrl,
		&data.Site,
		&data.ScanSite,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ChainStorage) GetToken(input *chain.GetTokenChainInput) (*chain.Chain, error) {
	query := `
		select id, chain_id, contract_address, name, symbol, decimals, image_url, site, scan_site from chain
		where chain_id = $1 and lower(contract_address) = lower($2) and contract_address != '0x';
	`

	var data = &chain.Chain{}

	err := s.db.QueryRow(
		query,
		input.ChainId,
		input.ContractAddress,
	).Scan(
		&data.ID,
		&data.ChainId,
		&data.ContractAddress,
		&data.Name,
		&data.Symbol,
		&data.Decimals,
		&data.ImageUrl,
		&data.Site,
		&data.ScanSite,
	)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ChainStorage) GetTokens(input *chain.GetTokensChainInput) ([]*chain.Chain, error) {
	query := `
		select id, chain_id, contract_address, name, symbol, decimals, image_url, site, scan_site from chain
		where chain_id = $1 and contract_address != '0x';
	`

	rows, err := s.db.Query(
		query,
		input.ChainId,
	)

	if err != nil {
		return []*chain.Chain{}, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var chainList []*chain.Chain
	for rows.Next() {
		var chainSingle chain.Chain

		err := rows.Scan(
			&chainSingle.ID,
			&chainSingle.ChainId,
			&chainSingle.ContractAddress,
			&chainSingle.Name,
			&chainSingle.Symbol,
			&chainSingle.Decimals,
			&chainSingle.ImageUrl,
			&chainSingle.Site,
			&chainSingle.ScanSite,
		)

		if err != nil {
			return nil, err
		}

		chainList = append(chainList, &chainSingle)
	}

	return chainList, nil
}
