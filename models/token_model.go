package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID                   uint64    `json:"id" gorm:"column:id"`
	MintAddress          string    `json:"mintAddress" gorm:"column:token_mint_address"`
	Owner                string    `json:"owner" gorm:"column:owner"`
	Supply               uint64    `json:"supply" gorm:"column:supply"`
	Collection           string    `json:"collection" gorm:"column:collection"`
	CollectionName       string    `json:"collectionName" gorm:"column:collection_name"`
	Name                 string    `json:"name" gorm:"column:name"`
	UpdateAuthority      string    `json:"updateAuthority" gorm:"column:update_authority"`
	PrimarySaleHappened  bool      `json:"primarySaleHappened" gorm:"column:primary_sale_happened"`
	SellerFeeBasisPoints uint64    `json:"sellerFeeBasisPoints" gorm:"column:seller_fee_basis_points"`
	Image                string    `json:"image" gorm:"column:image"`
	ExternalUrl          string    `json:"externalUrl" gorm:"column:external_url"`
	Attributes           string    `json:"attributes" gorm:"column:attributes"`
	Properties           string    `json:"properties" gorm:"properties"`
	Price                float64   `json:"price" gorm:"price"`
	ListStatus           string    `json:"listStatus" gorm:"list_status"`
	TokenAddress         string    `json:"tokenAddress" gorm:"token_address"`
	CreatedAt            time.Time `json:"created_at" gorm:"column:created_at"`
}

type Repository struct {
	Database *gorm.DB
}

type MetaResponse struct {
	Page        int `json:"page"`
	Limit       int `json:"limit"`
	TotalRecord int `json:"total_records"`
	TotalPages  int `json:"total_pages"`
}

type ListTokensResponse struct {
	Tokens []Token `json:"tokens"`
}

func New(d *gorm.DB) Repository {
	return Repository{
		Database: d,
	}
}

func (r Repository) Save(w *Token) (*Token, error) {
	res := r.Database.Create(w)
	return w, res.Error
}

func (r Repository) GetByWallet(owner string) (ListTokensResponse, error) {

	tokens := []Token{}

	results := r.Database.Where("owner = ?", owner).Order("created_at desc").Find(&tokens)

	listTokensResponse := ListTokensResponse{
		Tokens: tokens,
	}

	return listTokensResponse, results.Error
}

func (r Repository) GetOne(token_address string) (*Token, error) {

	token := Token{}

	results := r.Database.Where("token_mint_address = ?", token_address).First(&token)

	return &token, results.Error
}

func (r Repository) Delete(id string) (int64, error) {
	token := Token{}

	results := r.Database.Where("token_mint_address = ?", id).Delete(&token)

	return results.RowsAffected, results.Error
}
