package postgresql

import (
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/transaction"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
	"math/big"
	"strings"
)

type TransactionStorage struct {
	db *sql.DB
}

func NewTransactionStorage(db *sql.DB) *TransactionStorage {
	return &TransactionStorage{db}
}

func (s *TransactionStorage) Create(queryRower postgresql.Query, input *transaction.CreateTransactionInput) (*transaction.Transaction, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into transaction (chain_id, block_id, timestamp, block_hash, block_number_hex, block_number, from_address, gas, gas_price, hash, input, nonce, r, s, to_address, transaction_index, type, v, value, contract_address, cumulative_gas_used, gas_used, logs_bloom, status)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
		returning id, chain_id, block_id, timestamp, block_hash, block_number_hex, block_number, from_address, gas, gas_price, hash, input, nonce, r, s, to_address, transaction_index, type, v, value, contract_address, cumulative_gas_used, gas_used, logs_bloom, status;
	`

	blockNumberStr := input.BlockNumber.String()
	var getBlockNumberStr string

	valueStr := input.Value.String()
	var getValueStr string

	transactionData := &transaction.Transaction{}
	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.BlockId,
		input.Timestamp,
		input.BlockHash,
		input.BlockNumberHex,
		blockNumberStr,
		input.FromAddress,
		input.Gas,
		input.GasPrice,
		input.Hash,
		input.Input,
		input.Nonce,
		input.R,
		input.S,
		input.ToAddress,
		input.TransactionIndex,
		input.Type,
		input.V,
		valueStr,
		input.ContractAddress,
		input.CumulativeGasUsed,
		input.GasUsed,
		input.LogsBloom,
		input.Status,
	).Scan(
		&transactionData.ID,
		&transactionData.ChainId,
		&transactionData.BlockHash,
		&transactionData.Timestamp,
		&transactionData.BlockHash,
		&transactionData.BlockNumberHex,
		&getBlockNumberStr,
		&transactionData.FromAddress,
		&transactionData.Gas,
		&transactionData.GasPrice,
		&transactionData.Hash,
		&transactionData.Input,
		&transactionData.Nonce,
		&transactionData.R,
		&transactionData.S,
		&transactionData.ToAddress,
		&transactionData.TransactionIndex,
		&transactionData.Type,
		&transactionData.V,
		&getValueStr,
		&transactionData.ContractAddress,
		&transactionData.CumulativeGasUsed,
		&transactionData.GasUsed,
		&transactionData.LogsBloom,
		&transactionData.Status,
	)

	if getBlockNumberStr != "" {
		blockNumberInt := new(big.Int)
		blockNumberInt, ok := blockNumberInt.SetString(getBlockNumberStr, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert number string to big.Int")
		}
		transactionData.BlockNumber = *blockNumberInt
	}

	if getValueStr != "" {
		valueInt := new(big.Int)
		valueInt, ok := valueInt.SetString(getValueStr, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert number string to big.Int")
		}
		transactionData.Value = *valueInt
	}

	if err != nil {
		return nil, err
	}

	return transactionData, err
}

func (s *TransactionStorage) Update(queryRower postgresql.Query, input *transaction.UpdateTransactionInput) error {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		update transaction
		set contract_id = $2
		where id = $1;
	`

	_, err := s.db.Exec(
		query,
		input.TransactionId,
		input.ContractId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *TransactionStorage) GetCoin(input *transaction.GetTransactionInput) ([]*transaction.GetTransactionData, error) {
	typeOf := ""
	switch input.Type {
	case "send":
		typeOf = fmt.Sprintf("and lower(t.from_address) = lower('%s')", input.WalletAddress)
	case "received":
		typeOf = fmt.Sprintf("and lower(t.to_address) = lower('%s')", input.WalletAddress)
	default:
		typeOf = fmt.Sprintf("and (lower(t.from_address) = lower('%s') or lower(t.to_address) = lower('%s'))", input.WalletAddress, input.WalletAddress)
	}

	if input.Take > 100 {
		input.Take = 100
	}
	if input.Skip < 0 {
		input.Skip = 0
	}

	order := "asc"
	if strings.EqualFold(input.OrderBy, "desc") {
		order = "desc"
	}

	ca, err := utils.GetCA(input.ChainId)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		select 
		    t.hash, 
		    t.timestamp, 
		    t.from_address as from, 
		    t.to_address as to, 
		    t.value, 
		    t.gas_used, 
		    t.gas, 
		    t.gas_price,
		    case
		        when lower(t.from_address) = lower('%s') and t.to_address = '' then 'CONTRACT CREATE'
		    	when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
					then 'REMOVE LIQUIDITY'
		        when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
					then 'ADD LIQUIDITY'
				when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
					(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
					then 'SWAP'
		        when array_agg(distinct l.type) @> array[20]::integer[] and not (array_agg(distinct l.type) @> array[721,1155]::integer[]) then 'ERC20'
		        when array_agg(distinct l.type) @> array[721]::integer[] and not (array_agg(distinct l.type) @> array[20,1155]::integer[]) then 'ERC721'
		        when array_agg(distinct l.type) @> array[1155]::integer[] and not (array_agg(distinct l.type) @> array[20,721]::integer[]) then 'ERC1155'
		        when array_agg(distinct l.type) @> array[20, 721]::integer[] then 'ERC20, ERC721'
		        when array_agg(distinct l.type) @> array[20, 1155]::integer[] then 'ERC20, ERC1155'
		        when array_agg(distinct l.type) @> array[1155, 721]::integer[] then 'ERC721, ERC1155'
		        when left(t.input, 10) = '0x2e1a7d4d' then 'WITHDRAW'
				when left(t.input, 10) = '0x095ea7b3' then 'APPROVE'
		        when ct.chain_id is not null and lower(ct.hash) = lower(t.to_address) then 'CALL'
		        when count(distinct l.type) = 0 or (count(distinct l.type) = 1 and max(l.type) = 0) then 'COIN'
		        else 'UNKNOWN'
		    end as info,
		    case
		        when lower(t.from_address) = lower('%s') and t.to_address = '' then 'FEE'
		        when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
					then 'POOL REMOVED'
		    	when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
					then 'POOL CREATED'
				when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
					(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
					then 'SWAP'
		        when t.input = '0xd0e30db0' then 'DEPOSIT'
		        when left(t.input, 10) = '0x2e1a7d4d' then 'WITHDRAW'
				when left(t.input, 10) = '0x095ea7b3' then 'APPROVE'
			    when array_agg(distinct l.type) @> array[20]::integer[] and not (array_agg(distinct l.type) @> array[721,1155]::integer[]) then 'FEE'
		        when array_agg(distinct l.type) @> array[721]::integer[] and not (array_agg(distinct l.type) @> array[20,1155]::integer[]) then 'FEE'
		        when array_agg(distinct l.type) @> array[1155]::integer[] and not (array_agg(distinct l.type) @> array[20,721]::integer[]) then 'FEE'
		        when ct.chain_id is not null and lower(ct.hash) = lower(t.to_address) then 'CONTRACT CALL'
		        when lower(t.from_address) = lower('%s') then 'SEND'
		        when lower(t.to_address) = lower('%s') then 'RECEIVED'
		        else 'UNKNOWN'
		    end as type,
		    case
		        when t.status = '0x1' then 'SUCCESS'
				else 'FAILED'
			end as status,
			c.chain_id,
		    c.contract_address,
		    c.symbol,
		    c.decimals
		from transaction t
		inner join chain c on c.chain_id = t.chain_id
		left join contract ca on ca.transaction_create_id = t.id
		left join contract ct on ct.chain_id = c.chain_id and lower(ct.hash) = lower(to_address)
		left join log l on l.transaction_id = t.id
		where c.chain_id = $1 %s
		group by
		    t.id, t.hash, t.timestamp, t.from_address, t.to_address, t.value, t.gas_used, t.input,
		    t.gas, t.gas_price, c.chain_id, c.contract_address, c.symbol, c.decimals,
		    ca.transaction_create_id, t.status, ct.chain_id, ct.hash
		order by t.timestamp %s
		limit $2 offset $3;
	`, input.WalletAddress, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, input.WalletAddress, typeOf, order)

	rows, err := s.db.Query(query, input.ChainId, input.Take, input.Skip)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var transactions []*transaction.GetTransactionData
	for rows.Next() {
		var tx transaction.GetTransactionData
		if err := rows.Scan(&tx.Hash, &tx.Timestamp, &tx.From, &tx.To, &tx.Value, &tx.GasUsed, &tx.Gas, &tx.GasPrice, &tx.Info, &tx.Type, &tx.Status, &tx.ChainId, &tx.ContractAddress, &tx.Symbol, &tx.Decimals); err != nil {
			return nil, err
		}
		transactions = append(transactions, &tx)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *TransactionStorage) GetToken(input *transaction.GetTransactionInput) ([]*transaction.GetTransactionData, error) {
	contractFilter := fmt.Sprintf("and lower(l.address) = lower('%s')", input.ContractAddress)

	typeOf := ""
	switch input.Type {
	case "send":
		typeOf = fmt.Sprintf("and lower(l.from_address) = lower('%s')", input.WalletAddress)
	case "received":
		typeOf = fmt.Sprintf("and lower(l.to_address) = lower('%s')", input.WalletAddress)
	default:
		typeOf = fmt.Sprintf("and (lower(l.from_address) = lower('%s') or lower(l.to_address) = lower('%s'))", input.WalletAddress, input.WalletAddress)
	}

	if input.Take > 100 {
		input.Take = 100
	}
	if input.Skip < 0 {
		input.Skip = 0
	}

	order := "asc"
	if strings.EqualFold(input.OrderBy, "desc") {
		order = "desc"
	}

	ca, err := utils.GetCA(input.ChainId)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		select
			l.transaction_hash as hash,
			l.timestamp as timestamp,
			l.from_address as from,
		    l.to_address as to,
		    l.value as value,
		    t.gas_used,
		    t.gas,
		    t.gas_price,
		    case
		        when l.type = 20 then 'ERC20'
		        else 'UNKNOWN'
		    end as info,
		    case
		        when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
					then 'POOL REMOVED'
		    	when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
					(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
					then 'POOL CREATED'
				when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
					(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
					(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
					then 'SWAP'
		        when lower(l.from_address) = lower('0x0000000000000000000000000000000000000000') then 'MINT'
		        when lower(l.to_address) = lower('0x0000000000000000000000000000000000000000') or lower(l.to_address) = lower('0x000000000000000000000000000000000000dead') then 'BURN'
		        when lower(l.from_address) = lower('%s') then 'SEND'
		        when lower(l.to_address) = lower('%s') then 'RECEIVED'
		        when l.function != '' then l.function
		        else 'UNKNOWN'
		    end as type,
		    case
		    	when t.status = '0x1' then 'SUCCESS'
				else 'FAILED'
			end as status,
			l.chain_id,
			c.hash,
			c.symbol,
			c.decimals
		from log l
		inner join transaction t on l.transaction_id = t.id
		inner join contract c on c.chain_id = l.chain_id and lower(c.hash) = lower(l.address)
		where c.is_erc_20 = true and l.chain_id = $1 %s %s
		order by l.timestamp %s
		limit $2 offset $3;
	`, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress, input.WalletAddress, contractFilter, typeOf, order)

	rows, err := s.db.Query(query, input.ChainId, input.Take, input.Skip)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var transactions []*transaction.GetTransactionData
	for rows.Next() {
		var tx transaction.GetTransactionData
		if errs := rows.Scan(&tx.Hash, &tx.Timestamp, &tx.From, &tx.To, &tx.Value, &tx.GasUsed, &tx.Gas, &tx.GasPrice, &tx.Info, &tx.Type, &tx.Status, &tx.ChainId, &tx.ContractAddress, &tx.Symbol, &tx.Decimals); errs != nil {
			return nil, errs
		}
		transactions = append(transactions, &tx)
	}
	if errs := rows.Err(); errs != nil {
		return nil, errs
	}

	return transactions, nil
}

func (s *TransactionStorage) GetAll(input *transaction.GetAllTransactionInput) ([]*transaction.GetTransactionData, error) {
	if input.Take > 100 {
		input.Take = 100
	}
	if input.Skip < 0 {
		input.Skip = 0
	}

	ca, err := utils.GetCA(input.ChainId)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		select
			tx.hash, 
			tx.timestamp, 
			tx.from, 
			tx.to, 
			tx.value, 
			tx.gas_used, 
			tx.gas, 
			tx.gas_price,
			tx.info,
			tx.type,
			tx.status,
			tx.chain_id,
			tx.contract_address,
			tx.symbol,
			tx.decimals
		from (
			select
				t.hash, 
				t.timestamp, 
				t.from_address as from, 
				t.to_address as to, 
				t.value, 
				t.gas_used, 
				t.gas, 
				t.gas_price,
				case
					when lower(t.from_address) = lower('%s') and t.to_address = '' then 'CONTRACT CREATE'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
									then 'REMOVE LIQUIDITY'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
									then 'ADD LIQUIDITY'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
									(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
									then 'SWAP'
					when array_agg(distinct l.type) @> array[20]::integer[] and not (array_agg(distinct l.type) @> array[721,1155]::integer[]) then 'ERC20'
					when array_agg(distinct l.type) @> array[721]::integer[] and not (array_agg(distinct l.type) @> array[20,1155]::integer[]) then 'ERC721'
					when array_agg(distinct l.type) @> array[1155]::integer[] and not (array_agg(distinct l.type) @> array[20,721]::integer[]) then 'ERC1155'
					when array_agg(distinct l.type) @> array[20, 721]::integer[] then 'ERC20, ERC721'
					when array_agg(distinct l.type) @> array[20, 1155]::integer[] then 'ERC20, ERC1155'
					when array_agg(distinct l.type) @> array[1155, 721]::integer[] then 'ERC721, ERC1155'
					when left(t.input, 10) = '0x2e1a7d4d' then 'WITHDRAW'
					when left(t.input, 10) = '0x095ea7b3' then 'APPROVE'
					when ct.chain_id is not null and lower(ct.hash) = lower(t.to_address) then 'CALL'
					when count(distinct l.type) = 0 or (count(distinct l.type) = 1 and max(l.type) = 0) then 'COIN'
					else 'UNKNOWN'
				end as info,
				case
					when lower(t.from_address) = lower('%s') and t.to_address = '' then 'FEE'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
									then 'POOL REMOVED'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
									then 'POOL CREATED'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
									(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
									then 'SWAP'
					when t.input = '0xd0e30db0' then 'DEPOSIT'
					when left(t.input, 10) = '0x2e1a7d4d' then 'WITHDRAW'
					when left(t.input, 10) = '0x095ea7b3' then 'APPROVE'
					when array_agg(distinct l.type) @> array[20]::integer[] and not (array_agg(distinct l.type) @> array[721,1155]::integer[]) then 'FEE'
					when array_agg(distinct l.type) @> array[721]::integer[] and not (array_agg(distinct l.type) @> array[20,1155]::integer[]) then 'FEE'
					when array_agg(distinct l.type) @> array[1155]::integer[] and not (array_agg(distinct l.type) @> array[20,721]::integer[]) then 'FEE'
					when ct.chain_id is not null and lower(ct.hash) = lower(t.to_address) then 'CONTRACT CALL'
					when lower(t.from_address) = lower('%s') then 'SEND'
					when lower(t.to_address) = lower('%s') then 'RECEIVED'
					else 'UNKNOWN'
				end as type,
				case
					when t.status = '0x1' then 'SUCCESS'
					else 'FAILED'
				end as status,
				c.chain_id,
				c.contract_address,
				c.symbol,
				c.decimals
			from transaction t
			inner join chain c on c.chain_id = t.chain_id
			left join contract ca on ca.transaction_create_id = t.id
			left join contract ct on ct.chain_id = c.chain_id and lower(ct.hash) = lower(to_address)
			left join log l on l.transaction_id = t.id
			where c.chain_id = $3 and (lower(t.from_address) = lower('%s') or lower(t.to_address) = lower('%s'))
			group by
				t.id, t.hash, t.timestamp, t.from_address, t.to_address, t.value, t.gas_used, t.input,
				t.gas, t.gas_price, c.chain_id, c.contract_address, c.symbol, c.decimals,
				ca.transaction_create_id, t.status, ct.chain_id, ct.hash
		
			union all
		
			select
				l.transaction_hash as hash,
				l.timestamp as timestamp,
				l.from_address as from,
				l.to_address as to,
				l.value as value,
				t.gas_used,
				t.gas,
				t.gas_price,
				case
					when l.type = 20 then 'ERC20'
					else 'UNKNOWN'
				end as info,
				case
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'burn') = 1
									then 'POOL REMOVED'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l1 where l1.transaction_id = t.id and l1.type = 20) >= 3 and
									(select count(*) from log l2 where l2.transaction_id = t.id and l2.function = 'mint') = 2
									then 'POOL CREATED'
					when lower(t.from_address) = lower('%s') and lower(t.to_address) = lower('%s') and
									(select count(*) from log l3 where l3.transaction_id = t.id and l3.type = 20) >= 2 and
									(select count(*) from log l4 where l4.transaction_id = t.id and l4.function = 'transfer') = 2
									then 'SWAP'
					when lower(l.from_address) = lower('0x0000000000000000000000000000000000000000') then 'MINT'
					when lower(l.to_address) = lower('0x0000000000000000000000000000000000000000') or lower(l.to_address) = lower('0x000000000000000000000000000000000000dead') then 'BURN'
					when lower(l.from_address) = lower('%s') then 'SEND'
					when lower(l.to_address) = lower('%s') then 'RECEIVED'
					when l.function != '' then l.function
					else 'UNKNOWN'
				end as type,
				case
					when t.status = '0x1' then 'SUCCESS'
					else 'FAILED'
				end as status,
				l.chain_id,
				c.hash,
				c.symbol,
				c.decimals
			from log l
			inner join transaction t on l.transaction_id = t.id
			inner join contract c on c.chain_id = l.chain_id and lower(c.hash) = lower(l.address)
			where c.is_erc_20 = true and l.chain_id = $3 and (lower(l.from_address) = lower('%s') or lower(l.to_address) = lower('%s'))
		) as tx
		order by tx.timestamp desc
		limit $1 offset $2;
	`,
		input.WalletAddress, input.WalletAddress, input.WalletAddress, input.WalletAddress,
		input.WalletAddress, input.WalletAddress, ca.RouterCA, input.WalletAddress,
		input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA, input.WalletAddress,
		ca.RouterCA, input.WalletAddress, input.WalletAddress,
		input.WalletAddress, input.WalletAddress, input.WalletAddress, ca.RouterCA,
		input.WalletAddress, ca.RouterCA, input.WalletAddress, ca.RouterCA,
		input.WalletAddress, input.WalletAddress, input.WalletAddress, input.WalletAddress)

	rows, err := s.db.Query(query, input.Take, input.Skip, input.ChainId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var transactions []*transaction.GetTransactionData
	for rows.Next() {
		var tx transaction.GetTransactionData
		if errs := rows.Scan(&tx.Hash, &tx.Timestamp, &tx.From, &tx.To, &tx.Value, &tx.GasUsed, &tx.Gas, &tx.GasPrice, &tx.Info, &tx.Type, &tx.Status, &tx.ChainId, &tx.ContractAddress, &tx.Symbol, &tx.Decimals); errs != nil {
			return nil, errs
		}
		transactions = append(transactions, &tx)
	}
	if errs := rows.Err(); errs != nil {
		return nil, errs
	}

	return transactions, nil
}
