package main

import (
	"btrlogs/api"
)

func main() {
  r := api.NewRouter()

  r.Run(":8080")
}
