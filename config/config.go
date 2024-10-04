package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ConnectPort() string {
	// Nên setup 1 lần load file .env
	err := godotenv.Load("././.env")
	if err != nil {
		// Dùng thư viện log + log fatal nếu load file bị lỗi
		// Thì cần crash app để fix env
		fmt.Printf("Error loading .env file")
	}
	return ":" + os.Getenv("PORT")
}
