package contracts

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo/v4"
)

type HttpHandler interface {
	RegisterRoutes(e *echo.Echo)
}

type EventHandler interface {
	Name() string
	Topic() string
	Handle(msg *message.Message) error
}
