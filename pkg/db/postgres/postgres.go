package postgres

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectWithValues(SSLMode, Host, Port, User, DBName, Password string) *gorm.DB {
	connString := fmt.Sprintf("sslmode=%s host=%s port=%s user=%s dbname=%s password=%s",
		SSLMode, Host, Port, User, DBName, Password)

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panicf("%s\nfailed to connect database", err.Error())
	}
	return db
}
