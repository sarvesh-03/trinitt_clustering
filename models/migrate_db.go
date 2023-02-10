package models

import (
	"github.com/trinitt/config"
)

func MigrateDB() {
	db := config.GetDB()

	for _, model := range []interface{}{
		// User{},
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
