package users

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

var users []User = []User{
	{Id: 1, Username: "bob"},
	{Id: 2, Username: "alice"},
	{Id: 3, Username: "Mark"},
}

// ListUsers prints one line for each user every second.
func ListUsers(ctx context.Context, ch chan<- bool) {
	for _, user := range users {
		contextIsDone := false

		// wait for one second or until the context is done
		select {
		case <-ctx.Done():
			contextIsDone = true
			break
		case <-time.After(time.Second):
			fmt.Println(user)
		}

		// if the context is done, break out of the loop
		if contextIsDone {
			fmt.Println("context done")
			ch <- true
			break
		}
	}

	ch <- true
}
