package database

import (
	"deeploy-exam/app/config"
	"deeploy-exam/app/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sql *gorm.DB

func NewSQLClient(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.PQUser,
		cfg.PQPassword,
		cfg.PQDB,
		cfg.PQHost,
		cfg.PQPort,
	)
	var err error
	sql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("connect to database failed")
		log.Fatal(err)
	}
}

func UseSQL() *gorm.DB {
	return sql
}

func MigrateSQLSchema() {
	log.Print("Migrating Schema to PostgreSQL")
	err := sql.AutoMigrate(
		&models.Store{},
		&models.ProductCategory{},
		&models.Product{},
		&models.StoreProduct{},
	)
	if err != nil {
		log.Print("Migrating to database failed")
		log.Fatal(err)
	}
	log.Print("Migrating Succeed")
}
