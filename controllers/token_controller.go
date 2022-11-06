package controllers

import (
	"10xers/configs"
	"10xers/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type iTokenRepositorySave interface {
	Save(*models.Token) (*models.Token, error)
}

type tokenServiceCreate struct {
	iTokenRepositorySave
}

type FilesObj struct {
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

type CreatorObj struct {
	Address string `json:"address"`
	Share   uint64 `json:"share"`
}

type PropertiesObj struct {
	Edition  uint64       `json:"edition"`
	Files    []FilesObj   `json:"files"`
	Category string       `json:"category"`
	Creators []CreatorObj `json:"creators"`
}

type NewToken struct {
	MintAddress          string        `json:"mintAddress"`
	Owner                string        `json:"owner"`
	Supply               uint64        `json:"supply"`
	Collection           string        `json:"collection"`
	CollectionName       string        `json:"collectionName"`
	Name                 string        `json:"name"`
	UpdateAuthority      string        `json:"updateAuthority"`
	PrimarySaleHappened  bool          `json:"primarySaleHappened"`
	SellerFeeBasisPoints uint64        `json:"sellerFeeBasisPoints"`
	Image                string        `json:"image"`
	ExternalUrl          string        `json:"externalUrl"`
	Attributes           []interface{} `json:"attributes"`
	Properties           PropertiesObj `json:"properties"`
	Price                float64       `json:"price"`
	ListStatus           string        `json:"listStatus"`
	TokenAddress         string        `json:"tokenAddress"`
}

type DeleteToken struct {
	MintAddress string `json:"token_mint_address"`
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var tokens []NewToken

		walletAddress := c.Query("wallet_address")
		url := fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/wallets/%s/tokens?offset=0&limit=10&listStatus=listed", walletAddress)

		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err)
		}
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&tokens); err != nil {
			fmt.Print(err)
		}

		for i := 0; i < len(tokens); i++ {

			toStringAttributes := fmt.Sprintf("%v", tokens[i].Attributes)
			toStringProperties := fmt.Sprintf("%v", tokens[i].Properties)

			newToken := models.Token{
				MintAddress:          tokens[i].MintAddress,
				Owner:                tokens[i].Owner,
				Supply:               tokens[i].Supply,
				Collection:           tokens[i].Collection,
				CollectionName:       tokens[i].CollectionName,
				Name:                 tokens[i].Name,
				UpdateAuthority:      tokens[i].UpdateAuthority,
				PrimarySaleHappened:  tokens[i].PrimarySaleHappened,
				SellerFeeBasisPoints: tokens[i].SellerFeeBasisPoints,
				Image:                tokens[i].Image,
				ExternalUrl:          tokens[i].ExternalUrl,
				Attributes:           toStringAttributes,
				Properties:           toStringProperties,
				Price:                tokens[i].Price,
				ListStatus:           tokens[i].ListStatus,
				TokenAddress:         tokens[i].TokenAddress,
			}

			database := configs.Connect()
			tokenRepository := models.New(database)
			tokenService := tokenServiceCreate{tokenRepository}
			res, err := tokenService.Save(&newToken)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})

				return
			}

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "Data": err})

			} else {
				c.JSON(http.StatusNoContent, gin.H{"message": "success", "Data": res})
			}

		}
	}
}

func Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		database := configs.Connect()
		tokenRepository := models.New(database)
		walletAddress := c.Param("wallet_address")

		results, err := tokenRepository.GetByWallet(walletAddress)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "Data": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success", "Data": results})

	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		var token DeleteToken

		if err := c.ShouldBindJSON(&token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		database := configs.Connect()
		tokenRepository := models.New(database)

		results, err := tokenRepository.Delete(token.MintAddress)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "Data": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success", "Data": results})

	}
}
