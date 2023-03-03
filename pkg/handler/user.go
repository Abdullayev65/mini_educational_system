package hendler

import (
	"github.com/Abdullayev65/mini_educational_system/pkg/ioPut"
	"github.com/Abdullayev65/mini_educational_system/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var userAdd ioPut.UserAdd
	c.Bind(&userAdd)
	if ok, fieldName := userAdd.ValidForAdding(); !ok {
		response.Fail(c, "invalid user field ["+fieldName+"]")
		return
	}
	user, position := userAdd.MapToUser()
	err := h.Service.AddUserWithPosition(c, user, position)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, user.ID)
}
func (h *Handler) SignIn(c *gin.Context) {
	var sign ioPut.Sign
	c.Bind(&sign)
	user, err := h.Service.UserByUsername(c, sign.Username)
	if err != nil || user.Password != sign.Password {
		response.FailErrOrMsg(c, err, "username or password wrong")
		return
	}
	token, err := h.TokenJWT.Generate(user.ID)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, token)
}
func (h *Handler) UserMe(c *gin.Context) {
	userID := h.getUserID(c)
	userInfo, err := h.Service.UserInfoByID(c, userID)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, userInfo)
}
func (h *Handler) PatchUser(c *gin.Context) {
	userID := h.getUserID(c)
	userAdd := new(ioPut.UserAdd)
	c.Bind(userAdd)
	err := h.Service.UserUpdate(c, userID, userAdd)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, userID)
}
