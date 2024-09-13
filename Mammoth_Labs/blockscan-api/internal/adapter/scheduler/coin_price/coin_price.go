package coin_price

import (
	"blockscan-go/internal/config"
	"blockscan-go/internal/core/common/utils"
	"blockscan-go/internal/core/domain/crypto_currency"
	"blockscan-go/internal/database/postgresql"
	"encoding/json"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CoinPrice struct {
	cryptoCurrencyService crypto_currency.CronUseCase
	txManager             postgresql.DBTransactionManager
	cron                  *cron.Cron
	logger                *zap.Logger
	conf                  *config.Config
}

func NewCoinPrice(r crypto_currency.CronUseCase, tx postgresql.DBTransactionManager, cron *cron.Cron, logger *zap.Logger, conf *config.Config) *CoinPrice {
	return &CoinPrice{
		cryptoCurrencyService: r,
		txManager:             tx,
		cron:                  cron,
		logger:                logger,
		conf:                  conf,
	}
}

func (s *CoinPrice) Start() {
	_, _ = s.cron.AddFunc("0 * * * *", s.Price) // 매시간마다
}

func (s *CoinPrice) Price() {
	timestamp := time.Now().Unix()
	loc, _ := time.LoadLocation("Asia/Seoul") // 한국 시간 설정
	date := time.Unix(timestamp, 0).In(loc)
	// 해당 날짜의 아침 시간 (00:00) 계산
	morningDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	// 해당 날짜의 자정 직전 시간 (23:59:59.999) 계산
	midnightDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, int(time.Millisecond)-1, loc)

	url := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?id=1027,20688,23326,3890,1839"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.logger.Error("http create request err", zap.Error(err))
		return
	}

	apiKey := s.conf.CoinApiKey
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error("http client.Do error", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	var result Response
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		s.logger.Error("Json decoder error", zap.Error(err))
		return
	}

	cryptoInput := &crypto_currency.CreateCryptoCurrencyInput{
		Timestamp:    int(timestamp),
		MorningDate:  int(morningDate.Unix()),
		MidnightDate: int(midnightDate.Unix()),
		EthPrice:     -1,
		MmtPrice:     -1,
		GmmtPrice:    -1,
		MaticPrice:   -1,
		BnbPrice:     -1,
	}
	if eth, ok := result.Data["1027"]; ok {
		if eth.Quote["USD"].Price != nil {
			cryptoInput.EthPrice = *eth.Quote["USD"].Price
		}
	} else {
		s.logger.Error("get eth data error", zap.Error(err))
		return
	}
	if mmt, ok := result.Data["20688"]; ok {
		if mmt.Quote["USD"].Price != nil {
			cryptoInput.MmtPrice = *mmt.Quote["USD"].Price
		}
	} else {
		s.logger.Error("get mmt data error", zap.Error(err))
		return
	}
	if gmmt, ok := result.Data["23326"]; ok {
		if gmmt.Quote["USD"].Price != nil {
			cryptoInput.GmmtPrice = *gmmt.Quote["USD"].Price
		}

		layout := "2006-01-02 15:04:05 -0700 MST"
		t, err := time.Parse(layout, gmmt.LastUpdated.String())
		if err != nil {
			s.logger.Error("lastUpdated time parse error", zap.Error(err))
			return
		}
		cryptoInput.LastUpdated = int(t.Unix())

	} else {
		s.logger.Error("get gmmt data error", zap.Error(err))
		return
	}
	if matic, ok := result.Data["3890"]; ok {
		if matic.Quote["USD"].Price != nil {
			cryptoInput.MaticPrice = *matic.Quote["USD"].Price
		}
	} else {
		s.logger.Error("get matic data error", zap.Error(err))
		return
	}
	if bnb, ok := result.Data["1839"]; ok {
		if bnb.Quote["USD"].Price != nil {
			cryptoInput.BnbPrice = *bnb.Quote["USD"].Price
		}
	} else {
		s.logger.Error("get bnb data error", zap.Error(err))
		return
	}

	validator := utils.NewCustomValidator()
	if err := validator.Validate(cryptoInput); err != nil {
		s.logger.Error("cryptoInput validation error", zap.Error(err))
		return
	}

	cryptoCurrencyErr := s.cryptoCurrencyService.Create(nil, cryptoInput)
	if cryptoCurrencyErr != nil {
		s.logger.Error("crypto_currency create error", zap.Error(err))
		return
	}
}
