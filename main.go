package main

import (
	"log"
	term "system/src/stats"
	"system/src/sys"

	ui "github.com/gizak/termui/v3"
)

var c sys.CPU
var r term.Render
var p sys.Proc
var m sys.Memory
var grid ui.Grid

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	grid.SetRect(0, 0, 50, 10)

	// c.Graph()
	p.ProcsListRenderer()
	// m.Graph()

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
