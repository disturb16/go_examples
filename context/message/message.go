package message

import (
	"context"
	"fmt"
	"time"
)

// PrintMessage prints a message after the given delay.
func PrintMessage(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)

	case <-ctx.Done():
		fmt.Println("context is done")
	}
}
