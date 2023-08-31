package handler

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary GetUserSegmentsRequest
// @Tags userSegment
// @Description Get User Segments
// @Router /user-segments/users/{user_id} [get]
// @ID get-segment
// @Param user_id path int true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} []model.Segment
// @Failure 400,404 {string} string "error"
func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := getId(c, "userId")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UserSegments.GetUserSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segments": user.Segments,
	})
}

// @Summary UpdateUserSegmentsRequest
// @Tags userSegment
// @Description Update User Segments
// @Router /user-segments [put]
// @Param input body model.UserSegments true "user segments info"
// @ID update-user-segments
// @Accept json
// @Produce json
// @Success 200 {object} model.UserSegments
// @Failure 400,404 {string} string "error"
func (h *Handler) updateUserSegments(c *gin.Context) {
	var userSegments model.UserSegments

	if err := c.BindJSON(&userSegments); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.UserSegments.UpdateUserSegments(userSegments)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
