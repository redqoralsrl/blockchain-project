package config

import (
	"os"
)

type Config struct {
	Stage                     string
	ApiPort                   string
	DBUser                    string
	DBPassword                string
	DBName                    string
	DBHost                    string
	DBPort                    string
	RedisHost                 string
	RedisPort                 string
	JwtSecretKey              string
	XApiKey                   string
	XAdminApiKey              string
	CoinApiKey                string
	Network                   string
	GiantName                 string
	GiantEndPoint             string
	GiantChainId              string
	GiantChainSymbol          string
	GiantWrappedToken         string
	GiantWrappedSymbol        string
	TestnetGiantEndpoint      string
	TestnetGiantChainId       string
	TestnetGiantChainSymbol   string
	TestnetGiantWrappedSymbol string
	TestnetGiantWrappedToken  string
	CryptoniansPrivateKey     string
	CryptoniansMnemonic       string
}

func LoadConfig() *Config {

	return &Config{
		Stage:                     os.Getenv("STAGE"),
		ApiPort:                   os.Getenv("API_PORT"),
		DBUser:                    os.Getenv("DB_USER"),
		DBPassword:                os.Getenv("DB_PASSWORD"),
		DBName:                    os.Getenv("DB_NAME"),
		DBHost:                    os.Getenv("DB_HOST"),
		DBPort:                    os.Getenv("DB_PORT"),
		RedisHost:                 os.Getenv("REDIS_HOST"),
		RedisPort:                 os.Getenv("REDIS_PORT"),
		JwtSecretKey:              os.Getenv("JWT_SECRET_KEY"),
		XApiKey:                   os.Getenv("X_API_KEY"),
		XAdminApiKey:              os.Getenv("X_ADMIN_API_KEY"),
		CoinApiKey:                os.Getenv("COIN_API_KEY"),
		Network:                   os.Getenv("NETWORK"),
		GiantName:                 os.Getenv("GIANT_NAME"),
		GiantEndPoint:             os.Getenv("GIANT_ENDPOINT"),
		GiantChainId:              os.Getenv("GIANT_CHAIN_ID"),
		GiantChainSymbol:          os.Getenv("GIANT_CHAIN_SYMBOL"),
		GiantWrappedToken:         os.Getenv("GIANT_WRAPPED_TOKEN"),
		GiantWrappedSymbol:        os.Getenv("GIANT_WRAPPED_SYMBOL"),
		TestnetGiantEndpoint:      os.Getenv("TESTNET_GIANT_ENDPOINT"),
		TestnetGiantChainId:       os.Getenv("TESTNET_GIANT_CHAIN_ID"),
		TestnetGiantChainSymbol:   os.Getenv("TESTNET_GIANT_CHAIN_SYMBOL"),
		TestnetGiantWrappedSymbol: os.Getenv("TESTNET_GIANT_WRAPPED_SYMBOL"),
		TestnetGiantWrappedToken:  os.Getenv("TESTNET_GIANT_WRAPPED_TOKEN"),
		CryptoniansPrivateKey:     os.Getenv("CRYPTONIANS_PRIVATE_KEY"),
		CryptoniansMnemonic:       os.Getenv("CRYPTONIANS_MNEMONIC"),
	}
}
