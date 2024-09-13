package erc721

import (
	"blockscan-go/internal/core/domain/erc721"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase erc721.UseCase
}

func Erc721Handler(s erc721.UseCase, e *echo.Echo, apiAuth echo.MiddlewareFunc) {
	//h := handler{s}
	//
	//erc721Group := e.Group("/api/v1/erc721")

	//erc721Group.GET("/nfts/:wallet_address", h.Erc721GetNFTsByWallet, apiAuth)
	//erc721Group.GET("/:hash/detail/:token_id", h.GetNftDetail, apiAuth)
}
