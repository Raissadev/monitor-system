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
var d sys.Disk
var s sys.Swap
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
	ds, _ := d.AddGraph()
	sw, _ := s.AddGraph()

	grid.Set(
		ui.NewRow(0.1, ui.NewCol(0.3, cpu), ui.NewCol(0.3, mem), ui.NewCol(0.3, ds)),
		ui.NewRow(.2, sw),
		ui.NewRow(.7, procs),
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
