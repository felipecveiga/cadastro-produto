package service

import (

	"github.com/felipecveiga/cadastro-produto/model"
	"github.com/felipecveiga/cadastro-produto/repository"
)

type Service struct {
	Repository *repository.Repository
}

func NewUserService(useRepo *repository.Repository) *Service {
	return &Service{
		Repository: useRepo,
	}
}

func (s *Service) CreateProduct(product *model.Produtos) error {

		return s.Repository.CreateUserFromBD(product)
	
}