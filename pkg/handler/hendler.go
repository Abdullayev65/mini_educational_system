package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/service"
	"github.com/Abdullayev65/mini_educational_system/pkg/utill"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service  *service.Service
	TokenJWT *utill.TokenJWT
}

func New(service *service.Service, TokenJWT *utill.TokenJWT) *Handler {
	return &Handler{Service: service, TokenJWT: TokenJWT}
}

func (h *Handler) getUserID(c *gin.Context) int {
	userID, _ := c.Get("userID")
	return userID.(int)
}
