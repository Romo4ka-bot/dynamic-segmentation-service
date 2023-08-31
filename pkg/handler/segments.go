package handler

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary CreateSegmentRequest
// @Tags segment
// @Description Create Segment
// @Router /segments [post]
// @ID create-segment
// @Accept json
// @Produce json
// @Param input body model.Segment true "segment info"
// 400: BadRequest
// @Success 200 {object} model.Segment
// @Failure 400,404 {string} string "error"
func (h *Handler) createSegment(c *gin.Context) {
	var segment model.Segment

	if err := c.BindJSON(&segment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	segment, err := h.services.Segment.CreateSegment(segment)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segment": segment,
	})
}

// @Summary DeleteSegmentRequest
// @Tags segment
// @Description Delete Segment
// @Router /segments/{segment_id} [delete]
// @ID delete-segment
// @Param segment_id path int true "Segment ID"
// @Accept json
// @Produce json
// @Success 200 {string} nil
// @Failure 400,404 {string} string "error"
func (h *Handler) deleteSegment(c *gin.Context) {

	id, err := getId(c, "id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Segment.DeleteSegment(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
