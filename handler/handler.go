package handler

import (
	"github.com/acentior/train-ticket/repository"
	"gorm.io/gorm"
)

type Handler struct {
	db   *gorm.DB
	Repo *repository.Repository
}
