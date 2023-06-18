package sys

import (
	"fmt"
	"runtime"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/cpu"
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

// ? Consumption of all cores
func (c *CPU) update() (int, error) {
	resps, err := cpu.Percent(time.Second, false)

	if err != nil {
		return 0, fmt.Errorf("failed get data of cpu: %v", err)
	}

	var pΔ float64
	for _, percentage := range resps {
		pΔ += percentage
	}

	return int(pΔ), nil
}

func (c *CPU) AddGauge() (*widgets.Gauge, chan int) {
	c.graph = widgets.NewGauge()
	c.graph.BarColor = ui.Color(50)
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
		data, _ := c.update()
		_c <- data
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
