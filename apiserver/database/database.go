package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"prismcloud.dev/apiserver/database/model"
)

func createOutClusterDatabaseClient() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=administrator dbname=app port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func CreateDatabaseClient(outOfCluster bool) *gorm.DB {
	var database *gorm.DB
	if outOfCluster {
		database = createOutClusterDatabaseClient()
	} else {
		panic("Not implemented")
	}

	database.Logger = logger.Discard

	err := database.AutoMigrate(&model.Namespace{})
	if err != nil {
		panic(err)
	}

	return database
}
