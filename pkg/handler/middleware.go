package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getId(c *gin.Context, ctx string) (int, error) {
	id := c.Param(ctx)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id is of invalid type")
	}

	return idInt, nil
}
