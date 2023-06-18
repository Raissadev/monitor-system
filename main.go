package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"system/src/cli"
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
var flg cli.Flags
var gp cli.Graph

func main() {
	flags := flg.ManArgs()

	if flg.Contains(flags, "help") {
		return
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, width, height)

	grid = gp.Define(flags, grid)

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
