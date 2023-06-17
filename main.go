package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"system/src/sys"
	"time"

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
		ui.NewRow(.1, ui.NewCol(0.33, cpu), ui.NewCol(0.17, mem), ui.NewCol(0.17, sw), ui.NewCol(0.33, ds)),
		ui.NewRow(.3, nw),
		ui.NewRow(.6, procs),
	)

	ui.Render(grid)

	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case sig := <-sig:
			log.Printf("Received signal: %v", sig)
			cancel()
		case <-ctx.Done():
		}
	}()

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent && e.ID == "q" {
				ui.Close()
				os.Exit(0)
				return
			}
		case <-ctx.Done():
			time.Sleep(100 * time.Millisecond)
			ui.Close()
			os.Exit(0)
			return
		}
	}
}
