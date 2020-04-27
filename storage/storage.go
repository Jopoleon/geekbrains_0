package storage

import (
	"github.com/Jopoleon/geekbrains_0/config"
	"github.com/Jopoleon/geekbrains_0/storage/postgres"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	Logger *logrus.Logger
	DB     *postgres.DB
}

func NewStorage(cfg config.Config, logger *logrus.Logger) (*Storage, error) {
	res := &Storage{}
	db, err := postgres.NewPostgres(cfg, logger)
	if err != nil {
		return res, errors.WithStack(err)
	}

	res.DB = db
	res.Logger = logger
	return res, nil
}
