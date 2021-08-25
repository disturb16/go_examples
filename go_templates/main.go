package main

import (
	"os"
	"text/template"
)

func main() {
	name := "John"
	t, err := template.New("greeting").Parse("Hello {{.}}, nice to see you")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}
