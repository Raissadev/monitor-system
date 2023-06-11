package main

import (
	"log"
	"os"
	"system/src/sys"

	ui "github.com/gizak/termui/v3"
)

var c sys.CPU
var p sys.Proc
var m sys.Memory
var d sys.Disk
var s sys.Swap
var n sys.Network
var grid ui.Grid

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, width, height)

	procs, _ := p.AddList()
	cpu, _ := c.AddGauge()
	mem, _ := m.AddParagraph()
	ds, _ := d.AddGauge()
	sw, _ := s.AddPlot()
	nw, _ := n.AddPlot()

	grid.Set(
		ui.NewRow(.1, ui.NewCol(0.3, cpu), ui.NewCol(0.3, mem), ui.NewCol(0.3, ds)),
		ui.NewRow(.2, sw),
		ui.NewRow(.2, nw),
		ui.NewRow(.5, procs),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent && e.ID == "q" {
				ui.Close()
				os.Exit(0)
				return
			}
		}
	}
}
