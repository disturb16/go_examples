package api

import (
	"btrlogs/blog"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func withRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(string(blog.RequestIDKey))
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, blog.RequestIDKey, requestID)

		req, _ := http.NewRequestWithContext(
			ctx,
			c.Request.Method,
			c.Request.URL.String(),
			c.Request.Body,
		)

		c.Request = req
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(withRequestID())

	r.GET("/", ListCountries)
	r.GET("/:short_name", GetCountry)

	return r
}
