package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseQueryOffset(c *gin.Context) (int, error) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		return 0, fmt.Errorf("offset must be a non-negative integer")
	}

	return offset, nil
}

func ParseQueryLimit(c *gin.Context, max int) (int, error) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "25"))
	if err != nil || limit < 1 || limit > max {
		return 0, fmt.Errorf("limit must be between 1 and %d", max)
	}

	return limit, nil
}
