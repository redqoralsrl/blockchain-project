package nft

import (
	"blockscan-go/internal/core/domain/attributes"
	"blockscan-go/internal/core/domain/erc1155"
	"blockscan-go/internal/core/domain/erc721"
	"blockscan-go/internal/database/postgresql"
	"encoding/json"
	"errors"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type NftImageParser struct {
	erc721Service     erc721.CronUseCase
	erc1155Service    erc1155.CronUseCase
	attributesService attributes.CronUseCase
	txManager         postgresql.DBTransactionManager
	cron              *cron.Cron
	logger            *zap.Logger
}

func NewNftImageParser(s erc721.CronUseCase, c erc1155.CronUseCase, a attributes.CronUseCase, tx postgresql.DBTransactionManager, cron *cron.Cron, logger *zap.Logger) *NftImageParser {
	return &NftImageParser{
		erc721Service:     s,
		erc1155Service:    c,
		attributesService: a,
		txManager:         tx,
		cron:              cron,
		logger:            logger,
	}
}

func (s *NftImageParser) Start() {
	_, _ = s.cron.AddFunc("*/1 * * * *", s.process721)  // 1분 마다
	_, _ = s.cron.AddFunc("*/1 * * * *", s.process1155) // 1분 마다
}

func (s *NftImageParser) process721() {
	nftList := s.erc721Service.GetEmptyUrlErc721List()

	if len(nftList) == 0 {
		//s.logger.Info("nft image parser done")
		return
	}

	for _, data := range nftList {

		txDb, txDbErr := s.txManager.Begin()
		if txDbErr != nil {
			s.logger.Error("Error starting transaction cron", zap.Error(txDbErr))
			return
		}
		defer func() {
			_ = txDb.Rollback()
		}()

		var url string
		ipfsUrl := "https://ipfs.io/ipfs/"

		var err error
		switch scheme := strings.Split(data.Url, "://")[0]; scheme {
		case "ipfs":
			url = ipfsUrl + strings.Split(data.Url, "://")[1]
		case "http", "https":
			url = data.Url
		default:
			err = errors.New("unsupported scheme")
			//s.logger.Error("Unsupported url", zap.Error(err))
			_ = s.erc721Service.UpdateIsUndefinedMetaData(txDb, data.Erc721ID)
			return
		}

		if len(url) > 0 {
			response, httpErr := http.Get(url)
			if httpErr != nil {
				//s.logger.Error("Error fetching URL", zap.String("url", url), zap.Error(httpErr))
				return
			}
			defer func() {
				if response.Body != nil {
					_ = response.Body.Close()
				}
			}()

			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				m := make(map[string]interface{})
				if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
					return
				}

				infoData := &attributes.CreateErc721AttributesInput{
					Erc721Id:      data.Erc721ID,
					AttributeList: []attributes.Erc721AttributesInput{},
				}

				if img, ok := m["image"]; ok {
					infoData.ImageUrl = img.(string)
				}
				if name, ok := m["name"]; ok {
					infoData.Name = name.(string)
				}
				if description, ok := m["description"]; ok {
					infoData.Description = description.(string)
				}
				if externalUrl, ok := m["external_url"]; ok {
					infoData.ExternalUrl = externalUrl.(string)
				}
				if attr, ok := m["attributes"]; ok {
					if attrArray, isArray := attr.([]interface{}); isArray {
						for _, attribute := range attrArray {
							if attributeMap, isMap := attribute.(map[string]interface{}); isMap {
								traitType, traitTypeOk := attributeMap["trait_type"].(string)
								value, valueOk := attributeMap["value"].(string)
								displayType, displayOk := attributeMap["display_type"].(string)
								var attributeData = &attributes.Erc721AttributesInput{
									ChainId:    data.ChainId,
									ContractId: data.ContractId,
									Erc721Id:   data.Erc721ID,
								}
								if traitTypeOk && valueOk {
									attributeData.TraitType = traitType
									attributeData.Value = value
									if displayOk {
										attributeData.DisplayType = displayType
									}
									infoData.AttributeList = append(infoData.AttributeList, *attributeData)
								}
							}
						}
					}
				}

				erc721Error := s.erc721Service.Update(txDb, infoData)
				if erc721Error != nil {
					return
				}
				attributesErr := s.attributesService.CreateErc721(txDb, infoData)
				if attributesErr != nil {
					return
				}
			} else {
				if response.StatusCode == 429 || response.StatusCode == 504 {
					return
				} else {
					_ = s.erc721Service.UpdateIsUndefinedMetaData(txDb, data.Erc721ID)
					return
				}
			}
			_ = txDb.Commit()
		}
	}

}

func (s *NftImageParser) process1155() {
	nftList := s.erc1155Service.GetEmptyUrlErc1155List()

	if len(nftList) == 0 {
		return
	}

	for _, data := range nftList {
		txDb, txDbErr := s.txManager.Begin()
		if txDbErr != nil {
			s.logger.Error("Error starting transaction cron", zap.Error(txDbErr))
			return
		}
		defer func() {
			_ = txDb.Rollback()
		}()

		var url string
		ipfsUrl := "https://ipfs.io/ipfs/"

		var err error
		switch scheme := strings.Split(data.Url, "://")[0]; scheme {
		case "ipfs":
			url = ipfsUrl + strings.Split(data.Url, "://")[1]
		case "http", "https":
			url = data.Url
		default:
			err = errors.New("unsupported scheme")
			//s.logger.Error("Unsupported url", zap.Error(err))
			_ = s.erc1155Service.UpdateIsUndefinedMetaData(txDb, data.Erc1155ID)
			return
		}

		if len(url) > 0 {
			response, httpErr := http.Get(url)
			if httpErr != nil {
				//s.logger.Error("Error fetching URL", zap.String("url", url), zap.Error(httpErr))
				return
			}
			defer func() {
				if response.Body != nil {
					_ = response.Body.Close()
				}
			}()
			if httpErr != nil {
				return
			}

			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				m := make(map[string]interface{})
				if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
					return
				}

				infoData := &attributes.CreateErc1155AttributesInput{
					Erc1155Id:     data.Erc1155ID,
					AttributeList: []attributes.Erc1155AttributesInput{},
				}

				if img, ok := m["image"]; ok {
					infoData.ImageUrl = img.(string)
				}
				if name, ok := m["name"]; ok {
					infoData.Name = name.(string)
				}
				if description, ok := m["description"]; ok {
					infoData.Description = description.(string)
				}
				if externalUrl, ok := m["external_url"]; ok {
					infoData.ExternalUrl = externalUrl.(string)
				}
				if attr, ok := m["attributes"]; ok {
					if attrArray, isArray := attr.([]interface{}); isArray {
						for _, attribute := range attrArray {
							if attributeMap, isMap := attribute.(map[string]interface{}); isMap {
								traitType, traitTypeOk := attributeMap["trait_type"].(string)
								value, valueOk := attributeMap["value"].(string)
								displayType, displayOk := attributeMap["display_type"].(string)
								var attributeData = &attributes.Erc1155AttributesInput{
									ChainId:    data.ChainId,
									ContractId: data.ContractId,
									Erc1155Id:  data.Erc1155ID,
								}
								if traitTypeOk && valueOk {
									attributeData.TraitType = traitType
									attributeData.Value = value
									if displayOk {
										attributeData.DisplayType = displayType
									}
									infoData.AttributeList = append(infoData.AttributeList, *attributeData)
								}
							}
						}
					}
				}

				erc1155Error := s.erc1155Service.Update(txDb, infoData)
				if erc1155Error != nil {
					return
				}
				attributesErr := s.attributesService.CreateErc1155(txDb, infoData)
				if attributesErr != nil {
					return
				}
			} else {
				if response.StatusCode == 429 || response.StatusCode == 504 {
					return
				} else {
					_ = s.erc1155Service.UpdateIsUndefinedMetaData(txDb, data.Erc1155ID)
					return
				}
			}
			_ = txDb.Commit()
		}
	}
}
