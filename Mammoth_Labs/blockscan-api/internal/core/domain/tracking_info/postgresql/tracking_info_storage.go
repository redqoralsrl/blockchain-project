package postgresql

import (
	"blockscan-go/internal/core/domain/tracking_info"
	"blockscan-go/internal/database/postgresql"
	"database/sql"
	"math/big"
)

type TrackingInfoStorage struct {
	db *sql.DB
}

func NewTrackingInfoStorage(db *sql.DB) *TrackingInfoStorage {
	return &TrackingInfoStorage{db}
}

func (s *TrackingInfoStorage) Get(chainId int) (*tracking_info.TrackingInfo, error) {
	query := `select * from tracking_info where chain_id = $1;`

	row := s.db.QueryRow(query, chainId)

	var blockHeight int64
	trackingInfoData := &tracking_info.TrackingInfo{}
	if err := row.Scan(&trackingInfoData.ID, &trackingInfoData.CreatedAt, &trackingInfoData.ChainId, &blockHeight, &trackingInfoData.IsOperation); err != nil {
		return nil, err
	}
	trackingInfoData.BlockHeight = *big.NewInt(blockHeight)

	return trackingInfoData, nil
}

func (s *TrackingInfoStorage) Increase(queryRower postgresql.Query, chainId int) error {
	if queryRower == nil {
		queryRower = s.db
	}

	query := `
		update tracking_info
		set block_height = block_height + 1
		where chain_id = $1;
	`

	_, err := queryRower.Exec(query, chainId)
	return err
}
