package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServer(port string) *http.Server {

	engine := gin.New()
	UseCORsMiddleware(engine)
	UseRecoveryMiddleware(engine)

	engine.GET("/promotions/:id", PromotionsHandler)

	return &http.Server{
		Addr:    port,
		Handler: engine,
	}
}

func PromotionsHandler(c *gin.Context) {

	id := c.Param("id")
	result, err := Get(id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = c.Writer.Write(data)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
