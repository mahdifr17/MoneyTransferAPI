package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *gorm.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := gorm.Open(
		postgres.Open(opts.Dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
	return &Repository{db}
}
