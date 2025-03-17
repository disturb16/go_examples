package payments

import (
	"event-handler/payments/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	AccountID string `json:"account_id"`
	Total     int64  `json:"total"`
}

type Response struct {
	Data any `json:"data"`
}

func (h *handler) RegisterPayment(c echo.Context) error {
	req := Request{}
	ctx := c.Request().Context()
	log := logger.FromCtx(ctx)

	err := c.Bind(&req)
	if err != nil {
		log.WithError(err).Error("invalid request data")
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.WithAny("request", req).Info("")

	err = h.paymentManager.RegisterPayment(ctx, req.AccountID, req.Total)
	if err != nil {
		log.WithError(err).Error("failed to register payment")
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	resp := Response{
		Data: "success",
	}
	return c.JSON(http.StatusOK, resp)
}
