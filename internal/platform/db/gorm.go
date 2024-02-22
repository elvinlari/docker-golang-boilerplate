package db

import (
	"fmt"
	"os"
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	taskDomain "github.com/elvinlari/docker-golang/internal/task/domain"
	companyDomain "github.com/elvinlari/docker-golang/internal/company/domain"
	userDomain "github.com/elvinlari/docker-golang/internal/user/domain"
	inviteDomain "github.com/elvinlari/docker-golang/internal/invite/domain"
)

func Connect() (db *gorm.DB, err error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
		os.Getenv("DB_TIMEZONE"))

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func RunMigration(db *gorm.DB) error {
    // Run migration for domain structs
    if err := db.AutoMigrate(&taskDomain.Task{}); err != nil {
        return err
    }
	if err := db.AutoMigrate(&companyDomain.Company{}); err != nil {
        return err
    }
	if err := db.AutoMigrate(&userDomain.User{}); err != nil {
        return err
    }
	if err := db.AutoMigrate(&inviteDomain.Invite{}); err != nil {
        return err
    }

    return nil
}
