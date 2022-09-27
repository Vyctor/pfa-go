package main

import (
	"database/sql"
	"github.com/vyctor/pfa-go/internal/order/infra/database"
	"github.com/vyctor/pfa-go/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderOutputDTO{
		ID:    "123",
		Price: 10,
		Tax:   2,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	println(output.FinalPrice)

}
