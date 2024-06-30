package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	con := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBPort, config.DBUserName, config.DBName, config.DBUserPassword)

	DB, err = gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("? Connected to database")
}
