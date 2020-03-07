package handler

import (
	"github.com/alitaso345-sandbox/newsfeeder/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}