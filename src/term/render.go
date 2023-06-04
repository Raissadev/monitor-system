package term

import (
	"log"
	"system/src/sys"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var procs sys.Proc

type Render struct {
}

func (r *Render) Test() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init: %v", err)
	}

	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hey"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func (r *Render) ProcsListRenderer() {
	list := widgets.NewList()
	list.Title = "Procs list"
	list.Rows = []string{"loading..."}
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false

	list.SetRect(0, 0, 80, 24)

	ui.Render(list)

	go func() {
		for {
			processes, err := procs.ProcessLs()
			if err != nil {
				log.Printf("failed to get process list: %v", err)
			}

			list.Rows = processes
			ui.Render(list)

			time.Sleep(time.Second)
		}
	}()

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
