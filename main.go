package main

import (
	"hienviluong125/go-hex-app/logger"
	"hienviluong125/go-hex-app/middleware"
	"hienviluong125/go-hex-app/modules/productmodule"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	dbClient, err := sqlx.Connect("mysql", databaseUrl)

	if err != nil {
		panic(err.Error())
	}

	productRepo := productmodule.NewProductRepositoryMysql(dbClient)
	productService := productmodule.NewDefaultProductService(productRepo)
	productHandler := productmodule.NewProductHandler(productService)

	r := mux.NewRouter()

	r.Use(middleware.HandleErrorMiddleware)

	r.HandleFunc("/products", productHandler.Index).Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/products/{product_id}", productHandler.Show).Methods(http.MethodGet)
	r.HandleFunc("/products/{product_id}", productHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/products/{product_id}", productHandler.Destroy).Methods(http.MethodDelete)
	logger.Info("Listening server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
