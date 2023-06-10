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

	width, height := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, width, height)

	procs, _ := p.AddGraph()
	cpu, _ := c.AddGraph()
	mem, _ := m.AddGraph()

	grid.Set(
		ui.NewRow(0.1, ui.NewCol(0.5, cpu), ui.NewCol(0.5, mem)),
		ui.NewRow(1, procs),
	)

	ui.Render(grid)

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
