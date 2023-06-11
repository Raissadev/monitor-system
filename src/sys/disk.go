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
	graph *widgets.BarChart
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

	fmt.Println(lines[1])
	str := strings.TrimSpace(lines[1])
	str = strings.TrimSuffix(str, "%")
	usage, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to parse disk usage: %v", err)
	}

	return usage, nil
}

func (d *Disk) AddGraph() (*widgets.BarChart, chan []float64) {
	d.graph = widgets.NewBarChart()
	d.graph.Title = "Disk Usage"
	d.graph.Labels = []string{"Usage"}
	d.graph.TitleStyle.Fg = ui.ColorWhite
	d.graph.SetRect(0, 0, 50, 10)

	_data := make(chan []float64)

	go d.sender(_data)
	go d.receiver(_data)

	return d.graph, _data
}

func (d *Disk) sender(_d chan<- []float64) {
	for {
		usage, err := d.update("/")
		if err != nil {
			log.Fatalf("failed to get disk usage information: %v", err)
		}
		_d <- []float64{float64(usage)}
		time.Sleep(time.Second)
	}
}

func (d *Disk) receiver(_d <-chan []float64) {
	for {
		select {
		case us := <-_d:
			d.graph.Data = us
			ui.Render(d.graph)
		}
	}
}
