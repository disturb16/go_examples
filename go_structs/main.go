package main

import (
	"fmt"

	"github.com/disturb/go_structs/users"
)

func main() {
	// Id: 1, name: Martha, age: 20

	martha := users.User{Id: 1, Name: "Martha", Age: 20}
	pedro := users.User{Id: 2, Name: "Pedro", Age: 21}
	john := users.User{Id: 3, Name: "John", Age: 20}
	jane := users.User{Id: 4, Name: "Jane", Age: 21}

	martha.SayHello()
	martha.AddFriend(pedro)
	martha.AddFriend(john)
	martha.AddFriend(jane)

	martha.ListFriends()

	martha.RemoveFriend(john.Id)

	fmt.Println("========================")
	martha.ListFriends()
}
