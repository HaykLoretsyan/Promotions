package main

import (
	"github.com/gin-gonic/gin"
)

// UseCORsMiddleware attaches middleware for allowing CORs(Cross origin requests)
func UseCORsMiddleware(g *gin.Engine) {

	// Allowing COR requests from all origins
	g.Use(func(c *gin.Context) {
		referer := c.Request.Referer()
		if len(referer) > 0 && referer[len(referer)-1] == '/' {
			referer = referer[:len(c.Request.Referer())-1]
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", referer)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}

// UseRecoveryMiddleware attaches middleware for auto recovering from panics
func UseRecoveryMiddleware(g *gin.Engine) {

	// Auto recover from panics (write internal server error as response)
	g.Use(gin.Recovery())
}
