package repository

import (
	"dynamic-segmentation-service/pkg/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/url"
)

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	Charset  string
	SSLMode  string
	Scheme   string
}

func NewPostgresDB(config DBConfig) *gorm.DB {
	dsn := url.URL{
		User:     url.UserPassword(config.Username, config.Password),
		Scheme:   config.Scheme,
		Host:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Path:     config.DBName,
		RawQuery: (&url.Values{"sslmode": []string{config.SSLMode}}).Encode(),
	}

	db, err := gorm.Open(config.Dialect, dsn.String())
	if err != nil {
		logrus.Fatalf("could not connect database: %s", err.Error())
	}

	return model.DBMigrate(db)
}
