package main

import (
	"fmt"
	"log"
	"pbi-task-5/initializers"
	"pbi-task-5/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {

	err := initializers.DB.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		fmt.Println("something wrong when trying to migrate user model.")
	}
	fmt.Println("? Migration complete")
}
