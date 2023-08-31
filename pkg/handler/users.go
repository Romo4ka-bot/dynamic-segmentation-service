package handler

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary CreateUserRequest
// @Tags user
// @Description Create User
// @Router /users [post]
// @ID create-user
// @Accept json
// @Produce json
// @Param input body model.User true "user info"
// 400: BadRequest
// @Success 200 {object} model.User
// @Failure 400,404 {string} string "error"
func (h *Handler) createUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.User.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.HashPassword = ""

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
