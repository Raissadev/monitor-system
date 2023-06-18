package cli

import (
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
var flg Flags

type Graph struct {
}

func (g *Graph) Define(flags []string, grid *ui.Grid) *ui.Grid {
	if len(flags) == 0 {
		procs, _ := p.AddList()
		cpu, _ := c.AddGauge()
		mem, _ := m.AddParagraph()
		ds, _ := d.AddGauge()
		sw, _ := s.AddPlot()
		nw, _ := n.AddPlot()

		grid.Set(
			ui.NewRow(.1, ui.NewCol(0.33, cpu), ui.NewCol(0.17, mem), ui.NewCol(0.17, sw), ui.NewCol(0.33, ds)),
			ui.NewRow(.3, nw),
			ui.NewRow(.6, procs),
		)
	} else {
		rows := make([]ui.GridItem, 0)
		if flg.Contains(flags, "cpu") {
			cpu, _ := c.AddGauge()
			rows = append(rows, ui.NewRow(.1, cpu))
		}
		if flg.Contains(flags, "memory") {
			mem, _ := m.AddParagraph()
			rows = append(rows, ui.NewRow(.1, mem))
		}
		if flg.Contains(flags, "swap") {
			sw, _ := s.AddPlot()
			rows = append(rows, ui.NewRow(.1, sw))
		}
		if flg.Contains(flags, "disk") {
			ds, _ := d.AddGauge()
			rows = append(rows, ui.NewRow(.1, ds))
		}
		if flg.Contains(flags, "network") {
			nw, _ := n.AddPlot()
			rows = append(rows, ui.NewRow(.3, nw))
		}
		if flg.Contains(flags, "procs") {
			procs, _ := p.AddList()
			rows = append(rows, ui.NewRow(.6, procs))
		}
		grid.Set(flg.InterfaceSlice(rows)...)
	}

	return grid
}
