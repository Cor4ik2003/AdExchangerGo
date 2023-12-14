package handler

import (
	"AdExchangerGo/internal/model"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) SignUp(c *gin.Context) {
	request := model.User{}

	if err := c.BindJSON(&request); err != nil {
		log.Fatalf("error")
	}

}
