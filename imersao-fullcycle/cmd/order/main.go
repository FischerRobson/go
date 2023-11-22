package main

import (
	"database/sql"
	"fmt"
	"imersao-fullcycle/internal/infra/database"
	"imersao-fullcycle/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	uc := usecase.CalculateFinalPrice{
		OrderRepository: &database.OrderRepository{
			Db: db,
		},
	}

	input := usecase.OrderInput{
		ID:    "5",
		Price: 10,
		Tax:   1,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *output)
}
