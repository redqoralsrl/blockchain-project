package main

import (
	"blockscan-go/internal/adapter/http"
	"blockscan-go/internal/adapter/rpc"
	"blockscan-go/internal/adapter/scheduler"
	"blockscan-go/internal/config"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/database/postgresql"
	"blockscan-go/internal/database/redis"
	"context"
	"log"
)

func main() {
	conf := config.LoadConfig()

	stage := conf.Stage
	if "" == stage {
		stage = "dev"
	}

	// database
	db := postgresql.ConnectDatabase(conf)
	defer db.Close()

	// cache database
	cache := redis.ConnectCache(conf)
	defer cache.Close()

	// transaction db
	txManager := postgresql.NewManager(db)

	// logger
	newLogger, err := utils.NewLogger(stage)
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer func() {
		if err := newLogger.Sync(); err != nil {
			log.Fatalf("Falied to sync logger: %v", err)
		}
	}()

	// scheduler cron
	cronService, nftCronService, coinPriceService := scheduler.Run(db, txManager, newLogger, conf)
	defer cronService.Stop()
	nftCronService.Start()
	coinPriceService.Start()

	// rpc adapter
	ctx, cancel := context.WithCancel(context.Background())
	go rpc.Run(ctx, conf, newLogger, db, txManager)
	defer cancel()

	// echo adapter
	http.Run(db, newLogger, txManager, conf)
}
