package gin_router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ginEnginer struct {
}

func NewEnginerHandler() *ginEnginer {
	return &ginEnginer{}
}

func (h *ginEnginer) InitializeRoutes() http.Handler {
	r := gin.New()

	corsConf := cors.DefaultConfig()
	corsConf.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Api-Key", "Refreshtoken"}
	corsConf.AddExposeHeaders("Refreshtoken")
	corsConf.AllowOrigins = []string{"*"}

	r.Use(gin.Recovery(), cors.New(corsConf), gin.Logger())
	r.GET("/ping", h.ping())
	return r
}

func (h *ginEnginer) ping() gin.HandlerFunc {
	logrus.WithFields(logrus.Fields{"component": "gin - handler", "function": "ping"})
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}
