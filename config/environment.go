package config

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/joho/godotenv"
)


var ServerPort string
var kafkaUrl string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env"))
	}


	ServerPort = os.Getenv("SERVER_PORT")
	kafkaUrl = os.Getenv("KAFKA_URL")
}
