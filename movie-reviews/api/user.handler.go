package api

import (
	"encoding/json"
	"net/http"

	"github.com/disturb16/go_examples/movie-reviews/api/dto"
	"github.com/labstack/echo/v4"
)

func (api *API) Register(c echo.Context) error {

	// ctx := c.Request().Context()
	params := dto.RegisterUser{}

	data := c.Request().Body
	decoder := json.NewDecoder(data)

	err := decoder.Decode(&params)
	if err != nil {
		// TODO: Handle dto errors
		return c.JSON(http.StatusBadRequest, err)
	}

	err = api.userService.Register(params.ToUserModel())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, err)
}
