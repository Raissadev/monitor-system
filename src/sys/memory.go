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

func (m *Memory) Graph() {
	m.graph = widgets.NewParagraph()
	m.graph.Title = "Memory Usage"
	m.graph.Text = ""
	m.graph.SetRect(0, 0, 50, 5)

	ui.Render(m.graph)

	ticker := time.NewTicker(1 * time.Second)
	exit := make(chan struct{})

	go m.update(exit, ticker)

}

func (m *Memory) upInfo() string {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	used := stat.HeapInuse + stat.StackInuse
	total := stat.Sys

	usedMB := float64(used) / 1024 / 1024
	totalMB := float64(total) / 1024 / 1024

	text := fmt.Sprintf("Used: %.2f MB\nTotal: %.2f MB", usedMB, totalMB)

	return text
}

func (m *Memory) update(exit chan struct{}, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			m.graph.Text = m.upInfo()
			ui.Render(m.graph)
		case <-exit:
			ticker.Stop()
			return
		}
	}
}
