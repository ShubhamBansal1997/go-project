package server

import (
	"go-assignment/config"
	"go-assignment/db"

	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		DB:     db.Init(cfg),
		Config: cfg,
	}
}
