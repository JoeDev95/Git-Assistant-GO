package main

import (
	"errors"
	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("nenhum arquivo .env encontrado"))
		log.Fatal("Nenhum arquivo .env encontrado")
	}

	porta, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT não definida no arquivo .env"))
		log.Fatal("PORT não definida no arquivo .env")
	}

	err := http.ListenAndServe(":"+porta, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
