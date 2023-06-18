package sys

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/sys/unix"
)

/**
 * ?  Entropy is a measure of the unpredictability of a data set, popularly known as the degree of disorder of a system
 */

type EntropyKernel struct {
	graph *widgets.Plot
}

func (ek *EntropyKernel) AddPlot() (*widgets.Plot, chan float64) {
	ek.graph = widgets.NewPlot()
	ek.graph.Title = "Entropy kernel"
	ek.graph.Data = make([][]float64, 2)
	ek.graph.SetRect(0, 0, 50, 10)
	ek.graph.Marker = widgets.MarkerDot
	ek.graph.LineColors[0] = ui.Color(13)

	_data := make(chan float64)

	go ek.sender(_data)
	go ek.receiver(_data)

	return ek.graph, _data

}

func (ek *EntropyKernel) updateEnt() float64 {
	buf := make([]byte, 8)

	_, err := unix.Getrandom(buf, unix.GRND_NONBLOCK)
	if err != nil {
		log.Printf("Failed to get kernel entropy: %v", err)
		return 0.0
	}

	// ?   convert to 0 - 100
	return float64(buf[0]) / 255.0 * 100.0
}

func (ek *EntropyKernel) sender(_entropy chan<- float64) {
	for {
		_entropy <- ek.updateEnt()
		time.Sleep(time.Second)
	}
}

func (ek *EntropyKernel) receiver(_entropy <-chan float64) {
	i := 0
	for {
		select {
		case data := <-_entropy:
			ek.graph.Data[0] = append(ek.graph.Data[0], float64(i))
			ek.graph.Data[1] = append(ek.graph.Data[1], data)
			i++

			if i > 220 {
				ek.graph.Data[0] = ek.graph.Data[0][1:]
				ek.graph.Data[1] = ek.graph.Data[1][1:]
			}
			ui.Render(ek.graph)
		}
	}
}
