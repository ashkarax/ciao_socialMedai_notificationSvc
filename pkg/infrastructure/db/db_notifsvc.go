package db_notifSvc

import (
	"database/sql"
	"fmt"
	"time"

	domain_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/domain"
	config_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/config"
	interface_hash_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/utils/hash_password/interface"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config_notifSvc.DataBase, hashUtil interface_hash_notifSvc.IhashPassword) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBPort)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set the timezone to UTC
		},
	})
	if dberr != nil {
		return DB, nil
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain_notifSvc.Notification{}); err != nil {
		return DB, err
	}

	return DB, nil
}
