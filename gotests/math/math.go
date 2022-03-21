package math

import (
	"strconv"
	"time"
)

func sum(a, b int) int {
	return a + b
}

func sumStr(a, b string) int {

	time.Sleep(time.Millisecond * 100)

	c, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}

	d, err := strconv.Atoi(b)
	if err != nil {
		return 0
	}

	return c + d
}
