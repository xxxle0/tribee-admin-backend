package middelwares

import "github.com/gin-gonic/gin"

type Header struct {
	Authorization string `header:"Authorization"`
}

func Authentication(c *gin.Context) {
	header := Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, gin.H{
		"header": header,
	})
}
