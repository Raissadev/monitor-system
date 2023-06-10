package sys

import (
	"log"
	"os/exec"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Proc struct {
}

func (p *Proc) ProcessLs() ([]string, error) {
	cmd := exec.Command("ps", "-e", "-o", "pid,ppid,user,%cpu,%mem,command")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(output), "\n")[1:]

	procs := make([]string, 0)
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row != "" {
			procs = append(procs, row)
		}
	}

	return procs, nil
}

func (p *Proc) ProcsListRenderer() {
	list := widgets.NewList()
	list.Title = "Procs list"
	list.Rows = []string{"loading..."}
	list.TextStyle = ui.NewStyle(ui.ColorCyan)
	list.WrapText = false

	list.SetRect(0, 0, 80, 24)

	ui.Render(list)

	_procs := make(chan []string)

	go func() {
		for {
			processes, err := p.ProcessLs()
			if err != nil {
				log.Printf("failed to get process list: %v", err)
			}

			_procs <- processes

			time.Sleep(time.Second)

		}
	}()

	go func() {
		for {
			select {
			case processes := <-_procs:
				ui.Clear()
				list.Rows = processes
				ui.Render(list)
			}
		}
	}()
}
