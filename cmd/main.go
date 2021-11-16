package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/internal/app/handler"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/internal/app/repository"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

func main() {
	config := loadConfig()
	reposit, err := repository.NewRepos(config)
	if err != nil {
		log.Fatalf("Fail to open connection with db: %s", err.Error())
	}
	handler := handler.NewHandler(reposit)
	//создание строки для запуска на порте из config

	err = handler.ServerRun(":8080")
	if err != nil {
		log.Fatalf("Fail to server run: %s", err.Error())
	}
}

//LoadConfig инициализирует config для использования
//данных пользователя из .env
func loadConfig() *models.Config {
	config := models.NewConfig()
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatalf("No found .env")
	}
	config.Host = os.Getenv("host")
	config.Dbname = os.Getenv("dbname")
	config.Port = os.Getenv("port")
	config.Password = os.Getenv("password")
	config.User = os.Getenv("user")
	if (config.Dbname == "") || (config.Host == "") || (config.Port == "") || (config.Password == "") || (config.User == "") {
		log.Fatalf("No found params in .env")
	}
	return config
}
