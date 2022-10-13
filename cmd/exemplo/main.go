package main

import (
	"database/sql"
	"encoding/json"
	"github.com/vyctor/pfa-go/internal/order/infra/database"
	"github.com/vyctor/pfa-go/internal/order/usecase"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	
	repository := database.NewOrderRepository(db)

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewGetTotalUseCase(repository)
		output, err := uc.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(output)
	})

	go http.ListenAndServe(":8080", nil)
}
