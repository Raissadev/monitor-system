package main

import (
	"log"
	term "system/src/stats"
	"system/src/sys"

	ui "github.com/gizak/termui/v3"
)

var c term.Channel
var r term.Render
var p sys.Proc

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p.ProcsListRenderer()

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent && e.ID == "q" {
				return
			}
		}
	}
}
