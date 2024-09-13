package nft

import "blockscan-go/internal/core/common/utils"

type Reader interface {
	Get(input *GetNftInput) ([]*GetNftData, error)
	GetDetail(input *GetNftDetailInput) (*GetNFtDetailData, error)
	GetDetailOfWallet(input *GetNftDetailOfWalletInput) (*GetNFtDetailData, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(input *GetNftInput) ([]*GetNftData, *utils.ServiceError)
	GetDetail(input *GetNftDetailInput) (*GetNFtDetailData, *utils.ServiceError)
	GetDetailOfWallet(input *GetNftDetailOfWalletInput) (*GetNFtDetailData, *utils.ServiceError)
}

type RpcUseCase interface {
}
