package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/models"
	"github.com/Abdullayev65/mini_educational_system/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SubjectUpdate(c *gin.Context) {
	var sInput ioPut.SubjectInput
	c.Bind(&sInput)
	if sInput.Name == nil && sInput.Type == nil {
		response.Fail(c, "name and type is null")
		return
	}
	id := c.GetInt("id")
	subject, err := h.Service.SubjectByID(c, id)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	if sInput.Name != nil {
		subject.Name = *sInput.Name
	}
	if sInput.Type != nil {
		subject.Type = *sInput.Type
	}
	err = h.Service.SubjectUpdate(c, subject)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &subject)
}
func (h *Handler) SubjectDelete(c *gin.Context) {
	id := c.GetInt("id")
	err := h.Service.SubjectDelete(c, id)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, id)
}
func (h *Handler) SubjectAdd(c *gin.Context) {
	var sInput ioPut.SubjectInput
	c.Bind(&sInput)
	if sInput.Name == nil || sInput.Type == nil {
		response.Fail(c, "name or type not given")
		return
	}
	subject := models.Subject{Name: *sInput.Name, Type: *sInput.Type}
	err := h.Service.SubjectAdd(c, &subject)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, &subject)
}
func (h *Handler) SubjectAll(c *gin.Context) {
	deps := h.Service.SubjectAll(c)
	response.Success(c, deps)
}
