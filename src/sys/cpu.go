package sys

import (
	"runtime"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type CPU struct {
	NumCPU       int
	NumCores     int
	MaxFrequency int
	Usage        float64
	pTotal       uint64
	pIdle        uint64
	graph        *widgets.Gauge
	stat         runtime.MemStats
}

func (c *CPU) Info() *CPU {
	c.NumCPU = runtime.NumCPU()
	c.NumCores = int(runtime.NumCgoCall())
	c.MaxFrequency = 0

	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	c.Usage = float64(stat.Sys) / float64(1<<20)

	return c
}

func (c *CPU) update() {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	total := stat.Sys
	idle := stat.Sys - stat.HeapInuse - stat.StackInuse

	totalDelta := float64(total - c.pTotal)
	idleDelta := float64(idle - c.pIdle)
	cpuPercent := (1.0 - idleDelta/totalDelta) * 100

	c.pTotal = total
	c.pIdle = idle

	c.graph.Percent = int(cpuPercent)
}

func (c *CPU) Graph() {
	c.graph = widgets.NewGauge()
	c.graph.Title = "CPU Usage"
	c.graph.Percent = 0
	c.graph.SetRect(0, 0, 100, 5)

	ui.Render(c.graph)

	ticker := time.NewTicker(1 * time.Second)

	exit := make(chan struct{})

	go c.converge(exit, ticker)

}

func (c *CPU) converge(exit chan struct{}, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			c.update()
			ui.Render(c.graph)
		case <-exit:
			ticker.Stop()
			return
		}
	}
}
