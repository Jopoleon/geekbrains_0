package controllers

import (
	"github.com/Jopoleon/geekbrains_0/config"
	"github.com/Jopoleon/geekbrains_0/storage"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"time"
)

type Controllers struct {
	StartTime  time.Time
	Logger     *logrus.Logger
	HttpPort   string
	Config     *config.Config
	Router     *chi.Mux
	Repository *storage.Storage
}

func NewControllers(rep *storage.Storage, log *logrus.Logger, cfg *config.Config) *Controllers {
	a := &Controllers{
		HttpPort:   cfg.HttpPort,
		StartTime:  time.Now(),
		Config:     cfg,
		Logger:     log,
		Repository: rep,
	}
	return a
}
