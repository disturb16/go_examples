package profile

import (
	"event-handler/accounts/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) ProfileByAccountID(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.FromCtx(ctx)
	accountID := c.Param("accountID")

	log.Info("trying to get account")
	p, err := h.profileManager.ByID(ctx, accountID)
	if err != nil {
		log.WithError(err).Error("failed to get profile by id")
	}

	resp := Response{
		Data: p,
	}

	return c.JSON(http.StatusOK, resp)
}
