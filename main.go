package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vishals11/gouser/config"
)

func main() {
	cfg := config.Get()

	r := InitRoutes()
	fmt.Printf("Server listening on port: %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
