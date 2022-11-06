package main

import (
	"10xers/configs"
	"10xers/models"
	"fmt"
)

func init() {

	configs.Connect()

}

func main() {
	configs.DB.AutoMigrate(&models.Token{})
	fmt.Println("? Migration complete")
}
