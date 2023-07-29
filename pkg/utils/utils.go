package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON marsha and unmarshal goes here

func ParseBody(c *gin.Context) {
	if body, err := ioutil.ReadAll(c.Request.Body); err == nil {
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}
}
