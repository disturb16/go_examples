package api

import (
	"btrlogs/blog"
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

var allCountries = []Country{
	{
		Name:      "United States of America",
		ShortName: "USA",
	},
	{
		Name:      "United Kingdom",
		ShortName: "UK",
	},
	{
		Name:      "Canada",
		ShortName: "CA",
	},
}

func ListCountries(c *gin.Context) {
	ctx := c.Request.Context()
	logger := blog.New(ctx)

	// Simulate an error 40% of the time
	if rand.Intn(100) > 60 {

		logger.Error("error listing countries", slog.String("error", "simulated error"))
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, allCountries)
}

func GetCountry(c *gin.Context) {
	ctx := c.Request.Context()
	logger := blog.New(ctx)

	shortName := c.Params.ByName("short_name")

	for _, country := range allCountries {
		if country.ShortName == shortName {
			c.JSON(http.StatusOK, country)
			return
		}
	}

	logger.Error("could not found country", slog.String("short_name", shortName))
	c.JSON(http.StatusNotFound, "country not found")
}
