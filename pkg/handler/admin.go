package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UsersByDepartment(c *gin.Context) {
	depID := c.GetInt("departmentID")
	_ = depID + 1
	depInfo, err := h.Service.UsersByDepID(c, depID)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, depInfo)
}
