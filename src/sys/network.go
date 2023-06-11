package sys

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Network struct {
	graph *widgets.Plot
}

func (n *Network) AddPlot() (*widgets.Plot, chan []float64) {
	n.graph = widgets.NewPlot()
	n.graph.Title = "Network data"
	n.graph.DataLabels = []string{"Packets"}
	n.graph.Data = n.pseudoData()
	n.graph.SetRect(0, 0, 50, 10)

	n.graph.LineColors[0] = ui.Color(13)

	_data := make(chan []float64)

	go n.sender(_data)
	go n.receiver(_data)

	return n.graph, _data

}

func (n *Network) update() ([]float64, error) {
	cmd := exec.Command("ifconfig")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var data []float64
	rows := strings.Split(string(output), "\n")
	for _, row := range rows {
		if strings.Contains(row, "RX packets") {
			fields := strings.Fields(row)
			if len(fields) >= 5 {
				rxPkgs, _ := strconv.ParseFloat(strings.TrimSpace(fields[2]), 64)
				txPkgs, _ := strconv.ParseFloat(strings.TrimSpace(fields[5]), 64)
				totalPkgs := rxPkgs + txPkgs
				data = append(data, totalPkgs)
			}
		}
	}

	return data, nil
}
func (n *Network) sender(_n chan<- []float64) {
	for {
		data, err := n.update()
		if err != nil {
			log.Fatalf("failed: %v", err)
		}
		_n <- data
		time.Sleep(time.Second)
	}
}

func (n *Network) receiver(_n <-chan []float64) {
	i := 0
	for {
		select {
		case data := <-_n:
			n.graph.Data[0] = append(n.graph.Data[0], float64(i))
			n.graph.Data[1] = append(data)
			i++

			if i > 220 {
				n.graph.Data[0] = n.graph.Data[0][1:]
				n.graph.Data[1] = n.graph.Data[1][1:]
			}
			ui.Render(n.graph)
		}
	}
}

func (n *Network) pseudoData() [][]float64 {
	us, err := n.update()
	if err != nil {
		log.Fatalf("failed to get swap usage information: %v", err)
	}
	ŋ := 220
	data := make([][]float64, 2)
	data[0] = make([]float64, ŋ)
	data[1] = make([]float64, int(us[0]))
	return data
}
