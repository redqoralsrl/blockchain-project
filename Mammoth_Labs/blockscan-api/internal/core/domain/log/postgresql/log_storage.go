package postgresql

import (
	"blockscan-go/internal/core/domain/log"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"math/big"
)

type LogStorage struct {
	db *sql.DB
}

func NewLogStorage(db *sql.DB) *LogStorage {
	return &LogStorage{db}
}

func (s *LogStorage) Create(queryRower postgresql.Query, input *log.CreateLogInput) (*log.Log, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into log (chain_id, transaction_id, address, block_hash, block_number_hex, block_number, data, log_index, removed, transaction_hash, transaction_index, timestamp, function, type, dapp, from_address, to_address, value, token_id, url, name, symbol, decimals, erc1155_value, erc1155_token_id, erc1155_url, trade_nft_volume, trade_nft_volume_symbol, trade_nft_volume_contract, topics)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)
		returning id, chain_id, transaction_id, address, block_hash, block_number_hex, block_number, data, log_index, removed, transaction_hash, transaction_index, timestamp, function, type, dapp, from_address, to_address, value, token_id, url, name, symbol, decimals, erc1155_value, erc1155_token_id, erc1155_url, trade_nft_volume, trade_nft_volume_symbol, trade_nft_volume_contract, topics;
    `

	blockNumberStr := input.BlockNumber.String()
	var getBlockNumberStr string

	valueStr := input.Value.String()
	var getValueStr string

	tokenId := input.TokenId.String()
	var getTokenId string

	var erc1155ValueStrings []string
	for _, val := range input.Erc1155Value {
		erc1155ValueStrings = append(erc1155ValueStrings, val.String())
	}
	pqErc1155Value := pq.StringArray(erc1155ValueStrings)
	var getErc1155Value []string

	var erc1155TokenIdStrings []string
	for _, token := range input.Erc1155TokenId {
		erc1155TokenIdStrings = append(erc1155TokenIdStrings, token.String())
	}
	pqErc1155TokenId := pq.StringArray(erc1155TokenIdStrings)
	var getErc1155TokenId []string

	var erc1155UrlStrings []string
	for _, url := range input.Erc1155Url {
		erc1155UrlStrings = append(erc1155UrlStrings, url)
	}
	pqErc1155Url := pq.StringArray(erc1155UrlStrings)
	var getErc1155Url []string

	tradeNftVolume := input.TradeNftVolume.String()
	var getTradeNftVolume string

	var topicsStr []string
	for _, topic := range input.Topics {
		topicsStr = append(topicsStr, topic.Hex())
	}
	var getTopicArray pq.StringArray

	logData := &log.Log{}
	err := queryRower.QueryRow(
		query,
		input.ChainId,
		input.TransactionId,
		input.Address,
		input.BlockHash,
		input.BlockNumberHex,
		blockNumberStr,
		input.Data,
		input.LogIndex,
		input.Removed,
		input.TransactionHash,
		input.TransactionIndex,
		input.Timestamp,
		input.Function,
		input.Type,
		input.Dapp,
		input.FromAddress,
		input.ToAddress,
		valueStr,
		tokenId,
		input.Url,
		input.Name,
		input.Symbol,
		input.Decimals,
		pqErc1155Value,
		pqErc1155TokenId,
		pqErc1155Url,
		tradeNftVolume,
		input.TradeNftVolumeSymbol,
		input.TradeNftVolumeContract,
		pq.Array(topicsStr),
	).Scan(
		&logData.ID,
		&logData.ChainId,
		&logData.TransactionId,
		&logData.Address,
		&logData.BlockHash,
		&logData.BlockNumberHex,
		&getBlockNumberStr,
		&logData.Data,
		&logData.LogIndex,
		&logData.Removed,
		&logData.TransactionHash,
		&logData.TransactionIndex,
		&logData.Timestamp,
		&logData.Function,
		&logData.Type,
		&logData.Dapp,
		&logData.FromAddress,
		&logData.ToAddress,
		&getValueStr,
		&getTokenId,
		&logData.Url,
		&logData.Name,
		&logData.Symbol,
		&logData.Decimals,
		pq.Array(&getErc1155Value),
		pq.Array(&getErc1155TokenId),
		pq.Array(&getErc1155Url),
		&getTradeNftVolume,
		&logData.TradeNftVolumeSymbol,
		&logData.TradeNftVolumeContract,
		&getTopicArray,
	)

	if err != nil {
		return nil, err
	}

	if getBlockNumberStr != "" {
		blockNumberInt := new(big.Int)
		blockNumberInt, ok := blockNumberInt.SetString(getBlockNumberStr, 10)
		if !ok {
			return nil, fmt.Errorf("blockNumberInt failed to convert number string to big.Int")
		}
		logData.BlockNumber = *blockNumberInt
	}

	if getValueStr != "" {
		valueInt := new(big.Int)
		valueInt, ok := valueInt.SetString(getValueStr, 10)
		if !ok {
			return nil, fmt.Errorf("valueInt failed to convert number string to big.Int")
		}
		logData.Value = *valueInt
	}

	if getTokenId != "" {
		tokenIdInt := new(big.Int)
		tokenIdInt, ok := tokenIdInt.SetString(getTokenId, 10)
		if !ok {
			return nil, fmt.Errorf("tokenIdInt failed to convert number string to big.Int")
		}
		logData.TokenId = *tokenIdInt
	}

	if len(getErc1155Value) > 0 {
		var erc1155ValueInt []*big.Int
		for _, erc1155Val := range getErc1155Value {
			intValue := new(big.Int)
			intValue.SetString(erc1155Val, 10)
			erc1155ValueInt = append(erc1155ValueInt, intValue)
		}
		logData.Erc1155Value = erc1155ValueInt
	}

	if len(getErc1155TokenId) > 0 {
		var erc1155TokenIdInt []*big.Int
		for _, erc1155To := range getErc1155TokenId {
			erc1155TokenIdBigInt := new(big.Int)
			erc1155TokenIdBigInt.SetString(erc1155To, 10)
			erc1155TokenIdInt = append(erc1155TokenIdInt, erc1155TokenIdBigInt)
		}
		logData.Erc1155TokenId = erc1155TokenIdInt
	}

	if len(getErc1155Url) > 0 {
		logData.Erc1155Url = getErc1155Url
	}

	if getTradeNftVolume != "" {
		tradeNftVolumeInt := new(big.Int)
		tradeNftVolumeInt, ok := tradeNftVolumeInt.SetString(getTradeNftVolume, 10)
		if !ok {
			return nil, fmt.Errorf("tradeNftVolumeInt failed to convert number string to big.Int")
		}
		logData.TradeNftVolume = *tradeNftVolumeInt
	}

	if len(getTopicArray) > 0 {
		logData.Topics = []string(getTopicArray)
	}

	return logData, nil
}

func (s *LogStorage) GetNFTsByWallet(input *log.GetLogNFTsByWalletInput) ([]*log.GetLogNFTsByWalletData, error) {
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
		select
		    log.chain_id,
		    coalesce(log.address, '') as address,
		    log.block_hash,
		    log.block_number::text as block_number,
		    log.transaction_hash,
		    log.timestamp,
		    coalesce(log.function, '') as function,
		    coalesce(log.type, 0) as type, 
		    coalesce(log.dapp, '') as dapp,
		    coalesce(log.from_address, '') as from_address,
		    coalesce(log.to_address, '') as to_address,
		    coalesce(log.value, 0) as value,
		    coalesce(log.token_id, 0) as token_id,
		    coalesce(log.url, '') as url,
		    coalesce(log.name, '') as name,
		    coalesce(log.symbol, '') as symbol,
		    coalesce(log.trade_nft_volume, 0) as trade_nft_volume,
		    coalesce(log.trade_nft_volume_symbol, '') as trade_nft_volume_symbol,
		    coalesce(log.trade_nft_volume_contract, '') as trade_nft_volume_contract,
		    array_to_json(log.topics) as topics,
		    array_to_json(log.erc1155_value) as erc1155_value,
		    array_to_json(log.erc1155_token_id) as erc1155_token_id,
		    array_to_json(log.erc1155_url) as erc1155_url,
		    t.gas as gas,
		    t.gas_price as gas_price,
		    t.gas_used as gas_used,
		    json_agg(
				json_build_object(
					'chain_id', market_info.chain_id,
					'transaction_hash', market_info.transaction_hash,
					'collection',market_info.collection,
					'seller',market_info.seller,
					'buyer', market_info.buyer,
					'volume', market_info.volume::text,
					'volume_symbol', market_info.volume_symbol,
					'volume_contract_address', market_info.volume_contract_address
				)
			) filter (where market_info is not null) as market_volume_array,
		    json_agg(
				json_build_object(
					'chain_id', nft_volume.chain_id,
					'from_address', nft_volume.from_address,
					'to_address', nft_volume.to_address,
					'value', nft_volume.value::text,
					'contract', nft_volume.contract,
					'symbol', nft_volume.symbol,
					'timestamp', nft_volume.timestamp,
					'transaction_hash', nft_volume.transaction_hash,
					'event', nft_volume.event
				)
			) filter (where nft_volume is not null) as nft_volume_array
		from
		    log
		left join erc721 on erc721.token_id = log.token_id and erc721.contract_id = (select id from contract where contract.chain_id = log.chain_id and lower(contract.hash) = lower(log.address))
		left join erc1155 on erc1155.token_id = log.token_id and erc1155.contract_id = (select id from contract where contract.chain_id = log.chain_id and lower(contract.hash) = lower(log.address))
		inner join transaction t on t.id = log.transaction_id
		left join market_info on market_info.log_id = log.id
		left join nft_volume on log.id = nft_volume.log_id
		where (lower(log.from_address) = lower($1) or lower(log.to_address) = lower($1)) and (log.type = 721 or log.type = 1155) and log.chain_id = $4
		group by
		    log.id, log.chain_id, log.block_hash, 
		    coalesce(log.address, ''), log.chain_id, 
		    log.block_number::text, log.transaction_hash, 
		    log.timestamp, coalesce(log.function, ''), 
		    coalesce(log.type, 0), coalesce(log.dapp, ''), 
		    coalesce(log.from_address, ''), coalesce(log.to_address, ''), 
		    coalesce(log.value, 0), coalesce(log.token_id, 0), 
		    coalesce(log.url, ''), coalesce(log.name, ''), 
		    coalesce(log.symbol, ''), coalesce(log.trade_nft_volume, 0), 
		    coalesce(log.trade_nft_volume_symbol, ''), coalesce(log.trade_nft_volume_contract, ''), 
		    log.topics, log.erc1155_value, log.erc1155_token_id, log.erc1155_url, t.gas, t.gas_price, t.gas_used
		order by log.id desc
		limit $2 offset $3;
	`

	rows, err := s.db.Query(
		query,
		input.WalletAddress,
		limit,
		offset,
		input.ChainId,
	)
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return []*log.GetLogNFTsByWalletData{}, err
	}

	var logList []*log.GetLogNFTsByWalletData
	for rows.Next() {
		var logSingle log.GetLogNFTsByWalletData

		var getTopics sql.NullString
		var getErc1155Value sql.NullString
		var getErc1155TokenId sql.NullString
		var getErc1155Url sql.NullString
		var getMarketVolumeArray sql.NullString
		var getNftVolumeArray sql.NullString

		err := rows.Scan(
			&logSingle.ChainId,
			&logSingle.Address,
			&logSingle.BlockHash,
			&logSingle.BlockNumber,
			&logSingle.TransactionHash,
			&logSingle.Timestamp,
			&logSingle.Function,
			&logSingle.Type,
			&logSingle.DApp,
			&logSingle.FromAddress,
			&logSingle.ToAddress,
			&logSingle.Value,
			&logSingle.TokenId,
			&logSingle.Url,
			&logSingle.Name,
			&logSingle.Symbol,
			&logSingle.TradeNftVolume,
			&logSingle.TradeNftVolumeSymbol,
			&logSingle.TradeNftVolumeContract,
			&getTopics,
			&getErc1155Value,
			&getErc1155TokenId,
			&getErc1155Url,
			&logSingle.Gas,
			&logSingle.GasPrice,
			&logSingle.GasUsed,
			&getMarketVolumeArray,
			&getNftVolumeArray,
		)

		if err != nil {
			return []*log.GetLogNFTsByWalletData{}, err
		}

		if getTopics.Valid {
			err := json.Unmarshal([]byte(getTopics.String), &logSingle.Topics)
			if err != nil {
				return nil, fmt.Errorf("topics failed to convert json.Unmarshal")
			}
		} else {
			logSingle.Topics = []string{}
		}
		if getErc1155Value.Valid {
			var tempValues []big.Int
			err := json.Unmarshal([]byte(getErc1155Value.String), &tempValues)
			if err != nil {
				return nil, fmt.Errorf("erc1155_value failed to convert json.Unmarshal")
			}
			var bigIntValues []string
			for _, value := range tempValues {
				bigIntValues = append(bigIntValues, value.String())
			}
			logSingle.Erc1155Value = bigIntValues
		} else {
			logSingle.Erc1155Value = []string{}
		}

		if getErc1155TokenId.Valid {
			var tempValues []big.Int
			err := json.Unmarshal([]byte(getErc1155TokenId.String), &tempValues)
			if err != nil {
				return nil, fmt.Errorf("erc1155_value failed to convert json.Unmarshal")
			}
			var bigIntValues []string
			for _, value := range tempValues {
				bigIntValues = append(bigIntValues, value.String())
			}
			logSingle.Erc1155TokenId = bigIntValues
		} else {
			logSingle.Erc1155TokenId = []string{}
		}

		if getErc1155Url.Valid {
			err := json.Unmarshal([]byte(getErc1155Url.String), &logSingle.Erc1155Url)
			if err != nil {
				return nil, fmt.Errorf("erc1155_url failed to convert json.Unmarshal")
			}
		} else {
			logSingle.Erc1155Url = []string{}
		}

		if getMarketVolumeArray.Valid {
			var marketValues []log.MarketInfo
			err := json.Unmarshal([]byte(getMarketVolumeArray.String), &marketValues)
			if err != nil {
				return nil, fmt.Errorf("market_volume_array failed to convert json.Unmarshal")
			}
			logSingle.MarketInfoArray = marketValues
		} else {
			logSingle.MarketInfoArray = []log.MarketInfo{}
		}

		if getNftVolumeArray.Valid {
			var nftValues []log.NftVolumeData
			err := json.Unmarshal([]byte(getNftVolumeArray.String), &nftValues)
			if err != nil {
				return nil, fmt.Errorf("market_volume_array failed to convert json.Unmarshal")
			}
			logSingle.NftVolumeArray = nftValues
		} else {
			logSingle.NftVolumeArray = []log.NftVolumeData{}
		}

		logList = append(logList, &logSingle)
	}

	return logList, nil
}

func (s *LogStorage) GetSearchNFTsByWallet(input *log.GetSearchNFTsByWalletInput) ([]*log.GetLogNFTsByWalletData, error) {
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

	var dateSql string
	if input.Date == "7D" {
		dateSql = `AND log.timestamp >= extract(epoch from now() - interval '7 days')::integer`
	} else if input.Date == "30D" {
		dateSql = `AND log.timestamp >= extract(epoch from now() - interval '30 days')::integer`
	} else if input.Date == "1Y" {
		dateSql = `and log.timestamp >= extract(epoch from now() - interval '1 year')::integer`
	}

	var typeString string
	if input.Type == "mint" {
		typeString = "and lower(log.function) = 'mint'"
	} else if input.Type == "transfer" {
		typeString = "and log.trade_nft_volume = 0"
	} else if input.Type == "sale" {
		typeString = "and log.trade_nft_volume != 0"
	} else if input.Type == "burn" {
		typeString = "and lower(log.function) = 'burn'"
	}

	query := `
		select
		    log.chain_id,
		    coalesce(log.address, '') as address,
		    log.block_hash,
		    log.block_number::text as block_number,
		    log.transaction_hash,
		    log.timestamp,
		    coalesce(log.function, '') as function,
		    coalesce(log.type, 0) as type,
		    coalesce(log.dapp, '') as dapp,
		    coalesce(log.from_address, '') as from_address,
		    coalesce(log.to_address, '') as to_address,
		    coalesce(log.value, 0) as value,
		    coalesce(log.token_id, 0) as token_id,
		    coalesce(log.url, '') as url,
		    coalesce(log.name, '') as name,
		    coalesce(log.symbol, '') as symbol,
		    coalesce(log.trade_nft_volume, 0) as trade_nft_volume,
		    coalesce(log.trade_nft_volume_symbol, '') as trade_nft_volume_symbol,
		    coalesce(log.trade_nft_volume_contract, '') as trade_nft_volume_contract,
		    array_to_json(log.topics) as topics,
		    array_to_json(log.erc1155_value) as erc1155_value,
		    array_to_json(log.erc1155_token_id) as erc1155_token_id,
		    array_to_json(log.erc1155_url) as erc1155_url,
		    t.gas as gas,
		    t.gas_price as gas_price,
		    t.gas_used as gas_used,
		    json_agg(
				json_build_object(
					'chain_id', market_info.chain_id,
					'transaction_hash', market_info.transaction_hash,
					'collection',market_info.collection,
					'seller',market_info.seller,
					'buyer', market_info.buyer,
					'volume', market_info.volume::text,
					'volume_symbol', market_info.volume_symbol,
					'volume_contract_address', market_info.volume_contract_address
				)
			) filter (where market_info is not null) as market_volume_array,
		    json_agg(
				json_build_object(
					'chain_id', nft_volume.chain_id,
					'from_address', nft_volume.from_address,
					'to_address', nft_volume.to_address,
					'value', nft_volume.value::text,
					'contract', nft_volume.contract,
					'symbol', nft_volume.symbol,
					'timestamp', nft_volume.timestamp,
					'transaction_hash', nft_volume.transaction_hash,
					'event', nft_volume.event
				)
			) filter (where nft_volume is not null) as nft_volume_array
		from
		    log
		left join erc721 on erc721.token_id = log.token_id and erc721.contract_id = (select id from contract where contract.chain_id = log.chain_id and lower(contract.hash) = lower(log.address))
		left join erc1155 on erc1155.token_id = log.token_id and erc1155.contract_id = (select id from contract where contract.chain_id = log.chain_id and lower(contract.hash) = lower(log.address))
		inner join transaction t on t.id = log.transaction_id
		left join market_info on market_info.log_id = log.id
		left join nft_volume on log.id = nft_volume.log_id
		where (lower(log.from_address) = lower($1) or lower(log.to_address) = lower($1)) and (log.type = 721 or log.type = 1155) and log.chain_id = $4
			` + dateSql + typeString + `
		group by
		    log.id, log.chain_id, log.block_hash,
		    coalesce(log.address, ''), log.chain_id,
		    log.block_number::text, log.transaction_hash,
		    log.timestamp, coalesce(log.function, ''),
		    coalesce(log.type, 0), coalesce(log.dapp, ''),
		    coalesce(log.from_address, ''), coalesce(log.to_address, ''),
		    coalesce(log.value, 0), coalesce(log.token_id, 0),
		    coalesce(log.url, ''), coalesce(log.name, ''),
		    coalesce(log.symbol, ''), coalesce(log.trade_nft_volume, 0),
		    coalesce(log.trade_nft_volume_symbol, ''), coalesce(log.trade_nft_volume_contract, ''),
		    log.topics, log.erc1155_value, log.erc1155_token_id, log.erc1155_url, t.gas, t.gas_price, t.gas_used
		order by log.id desc
		limit $2 offset $3;
	`

	rows, err := s.db.Query(
		query,
		input.WalletAddress,
		limit,
		offset,
		input.ChainId,
	)
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return []*log.GetLogNFTsByWalletData{}, nil
	}

	var logList []*log.GetLogNFTsByWalletData
	for rows.Next() {
		var logSingle log.GetLogNFTsByWalletData

		var getTopics sql.NullString
		var getErc1155Value sql.NullString
		var getErc1155TokenId sql.NullString
		var getErc1155Url sql.NullString
		var getMarketVolumeArray sql.NullString
		var getNftVolumeArray sql.NullString

		err := rows.Scan(
			&logSingle.ChainId,
			&logSingle.Address,
			&logSingle.BlockHash,
			&logSingle.BlockNumber,
			&logSingle.TransactionHash,
			&logSingle.Timestamp,
			&logSingle.Function,
			&logSingle.Type,
			&logSingle.DApp,
			&logSingle.FromAddress,
			&logSingle.ToAddress,
			&logSingle.Value,
			&logSingle.TokenId,
			&logSingle.Url,
			&logSingle.Name,
			&logSingle.Symbol,
			&logSingle.TradeNftVolume,
			&logSingle.TradeNftVolumeSymbol,
			&logSingle.TradeNftVolumeContract,
			&getTopics,
			&getErc1155Value,
			&getErc1155TokenId,
			&getErc1155Url,
			&logSingle.Gas,
			&logSingle.GasPrice,
			&logSingle.GasUsed,
			&getMarketVolumeArray,
			&getNftVolumeArray,
		)

		if err != nil {
			return []*log.GetLogNFTsByWalletData{}, err
		}

		if getTopics.Valid {
			err := json.Unmarshal([]byte(getTopics.String), &logSingle.Topics)
			if err != nil {
				return nil, fmt.Errorf("topics failed to convert json.Unmarshal")
			}
		} else {
			logSingle.Topics = []string{}
		}
		if getErc1155Value.Valid {
			var tempValues []big.Int
			err := json.Unmarshal([]byte(getErc1155Value.String), &tempValues)
			if err != nil {
				return nil, fmt.Errorf("erc1155_value failed to convert json.Unmarshal")
			}
			var bigIntValues []string
			for _, value := range tempValues {
				bigIntValues = append(bigIntValues, value.String())
			}
			logSingle.Erc1155Value = bigIntValues
		} else {
			logSingle.Erc1155Value = []string{}
		}

		if getErc1155TokenId.Valid {
			var tempValues []big.Int
			err := json.Unmarshal([]byte(getErc1155TokenId.String), &tempValues)
			if err != nil {
				return nil, fmt.Errorf("erc1155_value failed to convert json.Unmarshal")
			}
			var bigIntValues []string
			for _, value := range tempValues {
				bigIntValues = append(bigIntValues, value.String())
			}
			logSingle.Erc1155TokenId = bigIntValues
		} else {
			logSingle.Erc1155TokenId = []string{}
		}

		if getErc1155Url.Valid {
			err := json.Unmarshal([]byte(getErc1155Url.String), &logSingle.Erc1155Url)
			if err != nil {
				return nil, fmt.Errorf("erc1155_url failed to convert json.Unmarshal")
			}
		} else {
			logSingle.Erc1155Url = []string{}
		}

		if getMarketVolumeArray.Valid {
			var marketValues []log.MarketInfo
			err := json.Unmarshal([]byte(getMarketVolumeArray.String), &marketValues)
			if err != nil {
				return nil, fmt.Errorf("market_volume_array failed to convert json.Unmarshal")
			}
			logSingle.MarketInfoArray = marketValues
		} else {
			logSingle.MarketInfoArray = []log.MarketInfo{}
		}

		if getNftVolumeArray.Valid {
			var nftValues []log.NftVolumeData
			err := json.Unmarshal([]byte(getNftVolumeArray.String), &nftValues)
			if err != nil {
				return nil, fmt.Errorf("market_volume_array failed to convert json.Unmarshal")
			}
			logSingle.NftVolumeArray = nftValues
		} else {
			logSingle.NftVolumeArray = []log.NftVolumeData{}
		}

		logList = append(logList, &logSingle)
	}

	return logList, nil
}
