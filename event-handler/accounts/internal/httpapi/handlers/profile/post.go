package profile

import (
	"event-handler/accounts/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateProfile(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.FromCtx(ctx)

	accountID := uuid.New().String()
	subscriptionID := uuid.New().String()

	err := h.profileManager.Register(ctx, accountID, subscriptionID)
	if err != nil {
		log.WithError(err).Error("failed to register new account")
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	resp := Response{
		Data: accountID,
	}

	return c.JSON(http.StatusOK, resp)
}
