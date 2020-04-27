package postgres

import (
	"fmt"

	"github.com/Jopoleon/geekbrains_0/config"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DB struct {
	Logger   *logrus.Logger
	DB       *sqlx.DB
	DBConfig *config.Config
}

func NewPostgres(cfg config.Config, logger *logrus.Logger) (*DB, error) {
	//postgresql://user:password@ip:port/database
	newSqlStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.DBUser, cfg.DB.DBPass, cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBName)
	//str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, cfg.DB.DBPass, cfg.DB.DBName)
	ommitesStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, "[ommited]", cfg.DB.DBName)

	pp.Println(newSqlStr)

	db, err := sqlx.Connect("postgres", newSqlStr)
	if err != nil {
		logger.Errorf("could not establish connection to ", ommitesStr, err)
		return nil, errors.WithStack(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return &DB{DB: db, Logger: logger, DBConfig: &cfg}, nil
}
