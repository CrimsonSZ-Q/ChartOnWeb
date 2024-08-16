package dataaccess

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var err error

type Config struct {
	DB struct {
		Server   string `json:"server"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"db"`
}

func InitializeDatabase() {
	config, err := readConfig("config/config.json")
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DB.User, config.DB.Password, config.DB.Server, config.DB.Port, config.DB.Database)

	log.Println("Attempting to open database connection...")
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	// defer DB.Close()
	log.Println("Database connection opened successfully.")
	fmt.Println("Connected to database!")
}

func readConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
