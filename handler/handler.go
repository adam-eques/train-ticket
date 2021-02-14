package handler

import (
	"github.com/midepeter/train-ticket/repository"
	"gorm.io/gorm"
)

type Handler struct {
	db   *gorm.DB
	Repo *repository.Repository
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}