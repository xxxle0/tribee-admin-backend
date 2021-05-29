package middlewares

import "github.com/gin-gonic/gin"

type RequestHeader struct {
	Authorization string `header:"Authorization"`
}

func Authentication(c *gin.Context) {
	header := RequestHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.AbortWithStatusJSON(403, "forbidden")
	}
	c.JSON(200, gin.H{
		"header": header,
	})
}
