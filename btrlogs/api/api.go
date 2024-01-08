package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", listCountries)
	r.GET("/:short_name", getCountry)

	return r
}
