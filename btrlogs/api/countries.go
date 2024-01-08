package api

import (
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

func listCountries(c *gin.Context) {
	// Simulate an error 40% of the time
	if rand.Intn(100) > 60 {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, allCountries)
}

func getCountry(c *gin.Context) {
	shortName := c.Params.ByName("short_name")

	for _, country := range allCountries {
		if country.ShortName == shortName {
			c.JSON(http.StatusOK, country)
			return
		}
	}

	c.JSON(http.StatusNotFound, "country not found")
}
