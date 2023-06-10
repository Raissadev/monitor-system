package sys

import (
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Proc struct {
	graph *widgets.List
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

func (p *Proc) AddGraph() (*widgets.List, chan []string) {
	p.graph = widgets.NewList()
	p.graph.Title = "Procs list"
	p.graph.Rows = []string{"loading..."}
	p.graph.TextStyle = ui.NewStyle(ui.ColorCyan)
	p.graph.WrapText = false

	_procs := make(chan []string)

	go p.sender(_procs)
	go p.receiver(_procs)

	return p.graph, _procs
}

func (p *Proc) sender(_p chan<- []string) {
	for {
		procs, err := p.ProcessLs()
		if err != nil {
			log.Printf("failed to get process list: %v", err)
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(procs), func(i, j int) {
			procs[i], procs[j] = procs[j], procs[i]
		})

		_p <- procs
		time.Sleep(time.Second)
	}
}

func (p *Proc) receiver(_p <-chan []string) {
	for {
		select {
		case procs := <-_p:
			p.graph.Rows = procs
			ui.Render(p.graph)
		}
	}
}
