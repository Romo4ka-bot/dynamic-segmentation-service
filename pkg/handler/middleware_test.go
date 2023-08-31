package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		c := &gin.Context{}
		c.Params = []gin.Param{{Key: "id", Value: fmt.Sprintf("%d", id)}}
		return c
	}

	testTable := []struct {
		name       string
		c          *gin.Context
		inputCtx   string
		id         int
		shouldFail bool
	}{
		{
			name:     "Ok",
			inputCtx: "id",
			c:        getContext(1),
			id:       1,
		},
		{
			c:          &gin.Context{},
			name:       "Empty",
			shouldFail: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getId(test.c, test.inputCtx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.id, id)
		})
	}
}
