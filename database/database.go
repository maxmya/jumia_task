package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	CONNECTION *gorm.DB
)

func InitDatabase() {
	var err error
	CONNECTION, err = gorm.Open(sqlite.Open("sample.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")
}
