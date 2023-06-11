package sys

import (
	"fmt"
	"runtime"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	used := stat.HeapInuse + stat.StackInuse
	total := stat.Sys

	usedMB := float64(used) / 1024 / 1024
	totalMB := float64(total) / 1024 / 1024

	text := fmt.Sprintf("Used: %.2f MB\nTotal: %.2f MB", usedMB, totalMB)

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
