package sys

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Disk struct {
	Usage int
	graph *widgets.Gauge
}

func (d *Disk) update(path string) (int, error) {
	out, err := exec.Command("df", "-h", "--output=pcent", path).Output()

	if err != nil {
		return 0, fmt.Errorf("failed to execute coomand to disk: %v", err)
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("unexpected output format")
	}

	str := strings.TrimSpace(lines[1])
	str = strings.TrimSuffix(str, "%")
	usage, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to parse disk usage: %v", err)
	}

	return usage, nil
}

func (d *Disk) AddGauge() (*widgets.Gauge, chan int) {
	d.graph = widgets.NewGauge()
	d.graph.BarColor = ui.Color(50)
	d.graph.BorderStyle.Fg = ui.ColorWhite
	d.graph.TitleStyle.Fg = ui.ColorCyan
	d.graph.Title = "Disk Usage"
	d.graph.Percent = 0

	_data := make(chan int)

	go d.sender(_data)
	go d.receiver(_data)

	return d.graph, _data
}

func (d *Disk) sender(_d chan<- int) {
	for {
		usage, err := d.update("/")
		if err != nil {
			log.Fatalf("failed to get disk usage information: %v", err)
		}
		_d <- usage
		time.Sleep(time.Second)
	}
}

func (d *Disk) receiver(_d <-chan int) {
	for {
		select {
		case data := <-_d:
			d.graph.Percent = data
			ui.Render(d.graph)
		}
	}
}
