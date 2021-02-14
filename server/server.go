package server

import (
	"github.com/acentior/train-ticket/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	h := handler.NewHandler(db)

	r := gin.Default()

	r.GET("/bookings", h.GetAllTickets)
	r.POST("/tickets", h.CreateTicket)
	r.PUT("/tickets/:id", h.UpdateTicket)
	r.GET("/tickets/:id", h.GetTicket)
	r.DELETE("/tickets/:id", h.DeleteTicket)

	return r
}
