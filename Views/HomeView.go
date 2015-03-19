package Views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	obj := c.MustGet("data").(map[string]string)
	c.HTML(http.StatusOK, "home.html",obj)
}