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
	Total int
	graph *widgets.Paragraph
}

func (s *Swap) AddPlot() (*widgets.Paragraph, chan Swap) {
	s.graph = widgets.NewParagraph()
	s.graph.Title = "Swap Usage"

	_data := make(chan Swap)

	go s.sender(_data)
	go s.receiver(_data)

	return s.graph, _data

}

func (s *Swap) update() (*Swap, error) {
	out, err := exec.Command("free", "-m").Output()
	if err != nil {
		return s, fmt.Errorf("failed to exec command free: %v", err)
	}

	lines := strings.Split(string(out), "\n")

	if len(lines) < 3 {
		return s, fmt.Errorf("unexpected format output")
	}

	sLine := strings.TrimSpace(lines[2])
	fields := strings.Fields(sLine)
	if len(fields) < 3 {
		return s, fmt.Errorf("unexpected format output")
	}
	totalΔ, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return s, fmt.Errorf("failed to parse total swap: %v", err)
	}

	usageΔ, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return s, fmt.Errorf("failed to parse used swap: %v", err)
	}

	s.Usage = int(usageΔ)
	s.Total = int(totalΔ)

	return s, nil
}

func (s *Swap) sender(_s chan<- Swap) {
	for {
		usage, err := s.update()
		if err != nil {
			log.Fatalf("failed to get swap usage information: %v", err)
		}
		_s <- *usage
		time.Sleep(time.Second)
	}
}

func (s *Swap) receiver(_s <-chan Swap) {
	for {
		select {
		case data := <-_s:
			s.graph.Text = fmt.Sprintf("Used: %d MB (∂)\nTotal: %d MB (∑)", data.Usage, data.Total)
			ui.Render(s.graph)
		}
	}
}
