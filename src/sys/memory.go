package sys

import (
	"fmt"
	"math"
	"runtime"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/mem"
)

type Memory struct {
	Allocated      uint
	TotalAllocated uint
	Sys            uint
	NumGC          uint
	Usage          float64
	graph          *widgets.Paragraph
}

func (m *Memory) Info() *Memory {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.Allocated = uint(memStats.Alloc)
	m.TotalAllocated = uint(memStats.TotalAlloc)
	m.Sys = uint(memStats.Sys)
	m.NumGC = uint(memStats.NumGC)

	return m
}

func (m *Memory) AddParagraph() (*widgets.Paragraph, chan string) {
	m.graph = widgets.NewParagraph()
	m.graph.Title = "Memory Usage"
	m.graph.Text = ""

	_msgs := make(chan string)

	go m.sender(_msgs)
	go m.receiver(_msgs)

	return m.graph, _msgs

}

func (m *Memory) update() string {
	memInfo, _ := mem.VirtualMemory()

	partialGB := float64(memInfo.Total) / math.Pow(1024, 3)
	sigmaGB := float64(memInfo.Used) / math.Pow(1024, 3)

	text := fmt.Sprintf("Used: %.2f GB (∂)\nTotal: %.2f GB (∑)", sigmaGB, partialGB)

	return text
}

func (m *Memory) sender(_m chan<- string) {
	for {
		_m <- m.update()
		time.Sleep(time.Second)
	}
}

func (m *Memory) receiver(_m <-chan string) {
	for {
		select {
		case msgs := <-_m:
			m.graph.Text = msgs
			ui.Render(m.graph)
		}
	}
}
