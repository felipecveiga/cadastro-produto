package model

type Produtos struct {
	ID        int     `gorm:"primary_key"`
	Nome      string  `gorm:"not null"`
	Descricao string  `gorm:"not null"`
	Preco     float64 `gorm:"not null"`
}
