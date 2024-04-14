package repository

import (


	"github.com/felipecveiga/cadastro-produto/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: *db,
	}
}


func (r Repository) CreateUserFromBD(product *model.Produtos) error {

	return r.DB.Create(product).Error
}
