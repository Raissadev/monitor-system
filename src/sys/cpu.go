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

func (c *CPU) update() int {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	total := stat.Sys
	idle := stat.Sys - stat.HeapInuse - stat.StackInuse

	totalDelta := float64(total - c.pTotal)
	idleDelta := float64(idle - c.pIdle)
	cpuPercent := (1.0 - idleDelta/totalDelta) * 100

	c.pTotal = total
	c.pIdle = idle

	return int(cpuPercent)
}

func (c *CPU) AddGraph() (*widgets.Gauge, chan int) {
	c.graph = widgets.NewGauge()
	c.graph.BarColor = ui.Color(300)
	c.graph.BorderStyle.Fg = ui.ColorWhite
	c.graph.TitleStyle.Fg = ui.ColorCyan
	c.graph.Title = "CPU Usage"
	c.graph.Percent = 0

	_data := make(chan int)

	go c.sender(_data)
	go c.receiver(_data)

	return c.graph, _data

}

func (c *CPU) sender(_c chan<- int) {
	for {
		_c <- c.update()
		time.Sleep(time.Second)
	}
}

func (c *CPU) receiver(_c <-chan int) {
	for {
		select {
		case data := <-_c:
			c.graph.Percent = data
			ui.Render(c.graph)
		}
	}
}
