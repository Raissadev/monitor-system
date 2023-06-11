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

type Swap struct {
	Usage int
	graph *widgets.Plot
}

func (s *Swap) AddGraph() (*widgets.Plot, chan float64) {
	s.graph = widgets.NewPlot()
	s.graph.Title = "Swap Usage"
	s.graph.Data = s.pseudoData()
	s.graph.SetRect(0, 0, 50, 10)

	s.graph.LineColors[0] = ui.Color(13)

	_data := make(chan float64)

	go s.sender(_data)
	go s.receiver(_data)

	return s.graph, _data

}

func (s *Swap) update() (float64, error) {
	out, err := exec.Command("free").Output()
	if err != nil {
		return 0, fmt.Errorf("failed to exec command free: %v", err)
	}

	lines := strings.Split(string(out), "\n")

	if len(lines) < 3 {
		return 0, fmt.Errorf("unexpected format output")
	}

	sLine := strings.TrimSpace(lines[2])
	fields := strings.Fields(sLine)
	if len(fields) < 3 {
		return 0, fmt.Errorf("unexpected format output")
	}
	stotal, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse total swap: %v", err)
	}

	usage, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse used swap: %v", err)
	}

	swap := float64(float64(usage) / float64(stotal) * 100)

	return swap, nil
}

func (s *Swap) sender(_s chan<- float64) {
	for {
		usage, err := s.update()
		if err != nil {
			log.Fatalf("failed to get swap usage information: %v", err)
		}
		_s <- usage
		time.Sleep(time.Second)
	}
}

func (s *Swap) receiver(_s <-chan float64) {
	i := 0

	for {
		select {
		case data := <-_s:
			s.graph.Data[0] = append(s.graph.Data[0], float64(i))
			s.graph.Data[1] = append(s.graph.Data[1], data)
			i++

			if i > 220 {
				s.graph.Data[0] = s.graph.Data[0][1:]
				s.graph.Data[1] = s.graph.Data[1][1:]
			}
			ui.Render(s.graph)
		}
	}
}

func (s *Swap) pseudoData() [][]float64 {
	us, err := s.update()
	if err != nil {
		log.Fatalf("failed to get swap usage information: %v", err)
	}
	n := 220
	data := make([][]float64, 2)
	data[0] = make([]float64, n)
	data[1] = make([]float64, int(us))
	return data
}
