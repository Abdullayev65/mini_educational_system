package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) AttendanceAdd(c *gin.Context) {
	input := new(ioPut.AttendanceInput)
	e := c.Bind(input)
	_ = e
	if input.Type == nil {
		c.String(400, "type can not be null")
		return
	}
	userID := h.getUserID(c)
	attendance := &models.Attendance{Type: *input.Type, Time: time.Now(), UserID: userID}
	err := h.Service.AttendanceAdd(c, attendance)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.JSON(200, attendance)
}
func (h *Handler) AttendanceAll(c *gin.Context) {
	all := h.Service.AttendanceAll(c)
	c.JSON(200, all)
}
func (h *Handler) AttendanceByUserID(c *gin.Context) {
	userId := h.getUserID(c)
	all, err := h.Service.AttendanceAllByUserID(c, userId)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, all)
}
