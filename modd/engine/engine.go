package engine

import (
	_ "embed"
	"log"
)

//go:embed templates/index.html
var index string

// Engine is the engine that renders the index.html file.
type Engine struct {
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) RenderIndex() string {
	log.Println("Rendering index")
	return index
}
