package main

import (
	"github.com/felipecveiga/cadastro-produto/handler"
	"github.com/felipecveiga/cadastro-produto/model"
	"github.com/felipecveiga/cadastro-produto/repository"
	"github.com/felipecveiga/cadastro-produto/service"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Connect to database
	dsn := "root:root@tcp(localhost:3306)/cadastro_produto?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha com a conexao com o database")
	}

	db.AutoMigrate(&model.Produtos{})

	Repo := repository.NewUserRepository(db)
	Service := service.NewUserService(Repo)
	Handler := handler.NewUserHandler(Service)

	e := echo.New()

	e.POST("create", Handler.CreateProductHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
