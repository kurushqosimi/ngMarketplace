package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// GET /api/categories/:id/attribute-schema
	return r
}

func WriteJSON(ctx *gin.Context, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		ctx.Header(key, strings.Join(value, ","))
	}

	ctx.Header("Content-Type", "application/json")

	ctx.Writer.WriteHeader(status)
	_, err = ctx.Writer.Write(js)
	if err != nil {
		return err
	}

	return nil
}
