package cart

import (
	"gorm.io/gorm"
)

type RepoCart struct {
	db *gorm.DB
}

func NewRepoCart(DB *gorm.DB) *RepoCart {
	return &RepoCart{
		db: DB,
	}
}
