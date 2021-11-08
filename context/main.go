package main

import (
	"context"

	"github.com/user/context/values"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "John")
	ctx = context.WithValue(ctx, "age", "30")

	values.PrintValues(ctx)
}
