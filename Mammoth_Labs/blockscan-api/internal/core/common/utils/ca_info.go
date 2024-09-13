package utils

import "fmt"

type CaInfo struct {
	FactoryCA string
	RouterCA  string
}

func GetCA(chainId int) (*CaInfo, error) {
	switch chainId {
	case 8989:
		return &CaInfo{
			FactoryCA: "0x6aD630595ADC6717119aB5c8192e1CEd94E0C587",
			RouterCA:  "0x92d8bF464931aab6323dab18d56bBb37e119DE53",
		}, nil
	case 898989:
		return &CaInfo{
			FactoryCA: "0x0e9e740319A4A2f4ea89Af6fd3A8B8016D6fe7e9",
			RouterCA:  "0xd78eC64aFB1273d80ae947fF0797830828975c9D",
		}, nil
	default:
		return nil, fmt.Errorf("unsupported chainId: %d", chainId)
	}
}
