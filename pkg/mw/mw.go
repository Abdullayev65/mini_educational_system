package mw

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/service"
	"github.com/Abdullayev65/mini_educational_system/pkg/utill"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MW struct {
	TokenJWT *utill.TokenJWT
	Serv     *service.Service
}

func New(tokenJWT *utill.TokenJWT, serv *service.Service) *MW {
	return &MW{TokenJWT: tokenJWT, Serv: serv}
}
func (mw *MW) SetIntFromParam(name string) func(*gin.Context) {
	return func(c *gin.Context) {
		param := c.Param(name)
		i, err := strconv.Atoi(param)
		if err != nil {
			c.String(400, "param "+name+" not found")
			c.Abort()
			return
		}
		c.Set(name, i)
	}
}
func (mw *MW) UserIDFromToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.String(http.StatusUnauthorized, "Authorization header is empty")
		c.Abort()
		return
	}
	userID, err := mw.TokenJWT.Parse(header)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	c.Set("userID", userID)
}
func (mw *MW) SetIntFromQuery(names ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, name := range names {
			query := c.Query(name)
			i, err := strconv.Atoi(query)
			if err != nil {
				c.String(400, "query "+name+" not found")
				c.Abort()
				return
			}
			c.Set(name, i)
		}
	}
}
func (mw *MW) CheckAdmin(c *gin.Context) {
	//todo

}
