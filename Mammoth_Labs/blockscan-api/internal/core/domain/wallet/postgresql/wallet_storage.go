package postgresql

import (
	"blockscan-go/internal/core/domain/wallet"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"fmt"
	"math/big"
)

type WalletStorage struct {
	db *sql.DB
}

func NewWalletStorage(db *sql.DB) *WalletStorage {
	return &WalletStorage{db}
}

func (s *WalletStorage) Get(queryRower postgresql.Query, address string) (*wallet.Wallet, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		select id, address, nick_name, profile, nonce, sign_hash, sign_timestamp from wallet
		where lower(address) = lower($1);
	`

	var getWalletNickName sql.NullString
	var getWalletProfile sql.NullString
	var getWalletNonce sql.NullString
	var getWalletSignHash sql.NullString
	var getWalletSignTimestamp sql.NullInt64

	var walletData = &wallet.Wallet{}
	err := queryRower.QueryRow(
		query,
		address,
	).Scan(
		&walletData.ID,
		&walletData.Address,
		&getWalletNickName,
		&getWalletProfile,
		&getWalletNonce,
		&getWalletSignHash,
		&getWalletSignTimestamp,
	)

	if err != nil {
		return nil, err
	}

	if getWalletNickName.Valid {
		walletData.NickName = getWalletNickName.String
	}
	if getWalletProfile.Valid {
		walletData.Profile = getWalletProfile.String
	}
	if getWalletNonce.Valid {
		walletNonceInt := new(big.Int)
		walletNonceInt, ok := walletNonceInt.SetString(getWalletNonce.String, 10)
		if !ok {
			return nil, fmt.Errorf("walletNonceInt failed to convert number string to big.Int")
		}

		walletData.Nonce = *walletNonceInt
	}
	if getWalletSignHash.Valid {
		walletData.SignHash = getWalletSignHash.String
	}
	if getWalletSignTimestamp.Valid {
		walletData.SignTimestamp = int(getWalletSignTimestamp.Int64)
	}

	return walletData, nil
}

func (s *WalletStorage) Create(queryRower postgresql.Query, address string) (int, error) {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		insert into wallet (address)
		values ($1)
		returning id;
	`

	var walletId int
	err := queryRower.QueryRow(
		query,
		address,
	).Scan(
		&walletId,
	)

	if err != nil {
		return -1, nil
	}

	return walletId, nil
}
