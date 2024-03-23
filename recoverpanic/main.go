package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func logNum(n int) {
	log.Println(n)

	if n == 5 {
		panic("n is 5")
	}
}

var lastNum = 0

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("recovered from panic")
	// 		lastNum++
	// 		logNums()
	// 	}
	// }()

	// logNums()

	r := echo.New()
	r.Use(middleware.Recover())

	r.GET("/", func(c echo.Context) error {
		logNums()
		return c.String(200, "OK")
	})

}

func logNums() {
	for i := lastNum; i < 10; i++ {
		lastNum = i
		logNum(i)
	}
}
