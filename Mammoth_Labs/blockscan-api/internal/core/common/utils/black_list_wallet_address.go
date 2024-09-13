package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strings"
)

func ValidateWalletAddress(walletAddress string) error {

	var re = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	res := re.MatchString(walletAddress)

	if !res {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid wallet address")
	}

	invalidAddress := []string{
		"0x0000000000000000000000000000000000000000",
		"0x000000000000000000000000000000000000dead",
	}

	for _, addr := range invalidAddress {
		if strings.EqualFold(walletAddress, addr) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid wallet address")
		}
	}

	return nil
}
