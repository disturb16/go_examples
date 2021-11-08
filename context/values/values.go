package values

import (
	"context"
	"fmt"
)

func PrintValues(ctx context.Context) {

	name, ok := ctx.Value("name").(string)
	if !ok {
		fmt.Println("name not found")
		return
	}

	age, ok := ctx.Value("age").(int)
	if !ok {
		fmt.Println("age not found")
		return
	}

	fmt.Println("name:", name)
	fmt.Println("age:", age)
}
