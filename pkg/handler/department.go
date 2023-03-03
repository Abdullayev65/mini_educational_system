package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/Abdullayev65/mini_educational_system/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DepAdd(c *gin.Context) {
	var depInput ioPut.DepInput
	c.Bind(&depInput)
	if depInput.Name == nil {
		response.Fail(c, "name not given")
		return
	}
	department := models.Department{Name: *depInput.Name}
	err := h.Service.DepAdd(c, &department)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &department)
}
func (h *Handler) DepAll(c *gin.Context) {
	deps := h.Service.DepAll(c)
	response.Success(c, deps)
}
func (h *Handler) DepUpdate(c *gin.Context) {
	var depInput ioPut.DepInput
	c.Bind(&depInput)
	if depInput.Name == nil {
		response.Fail(c, "name not given")
		return
	}
	id := c.GetInt("id")
	department := models.Department{ID: id, Name: *depInput.Name}
	err := h.Service.DepUpdate(c, &department)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &department)
}
func (h *Handler) DepDelete(c *gin.Context) {
	id := c.GetInt("id")
	err := h.Service.DepDelete(c, id)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, id)
}
