package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/Abdullayev65/mini_educational_system/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) LvlAdd(c *gin.Context) {
	var lvlInput ioPut.LvlInput
	c.Bind(&lvlInput)
	if lvlInput.Type == nil {
		response.Fail(c, "type not given")
		return
	}
	level := models.Level{Type: *lvlInput.Type}
	err := h.Service.LvlAdd(c, &level)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &level)
}
func (h *Handler) LvlAll(c *gin.Context) {
	lvls := h.Service.LvlAll(c)
	response.Success(c, lvls)
}
func (h *Handler) LvlUpdate(c *gin.Context) {
	var lvlInput ioPut.LvlInput
	c.Bind(&lvlInput)
	if lvlInput.Type == nil {
		response.Fail(c, "type not given")
		return
	}
	id := c.GetInt("id")
	level := models.Level{ID: id, Type: *lvlInput.Type}
	err := h.Service.LvlUpdate(c, &level)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &level)
}
func (h *Handler) LvlDelete(c *gin.Context) {
	id := c.GetInt("id")
	err := h.Service.LvlDelete(c, id)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, id)
}
