package main

import (
	"Pengenalan-Microservice/config"
	"Pengenalan-Microservice/database"
	"Pengenalan-Microservice/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Config{
		Database: config.Database{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "admin123",
			DbName:   "pengenalan_microservice",
			Config:   "charset=utf8&parseTime=True&loc=Local",
		},
		Auth: config.Auth{
			Host: "http://localhost:8001",
		},
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Panic(err)
		return
	}

	router := mux.NewRouter()

	menuHandler := handler.MenuHandler{
		Db: db,
	}
	authHandler := handler.AuthHandler{
		Config: cfg.Auth,
	}

	router.Handle("/add-item", authHandler.ValidateAdmin(menuHandler.AddMenu))
	router.Handle("/menu", http.HandlerFunc(menuHandler.GetMenu))

	fmt.Println("Menu service listen on port :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(database.Menu{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Database")

	return db, nil
}
