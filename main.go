package main

import (
	"log"
	"system/src/term"

	ui "github.com/gizak/termui/v3"
)

var c term.Channel
var r term.Render

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	r.ProcsListRenderer()
}
