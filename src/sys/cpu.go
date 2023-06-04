package sys

import "runtime"

type CPU struct {
	NumCPU       int
	NumCores     int
	MaxFrequency int
	Usage        float64
}

func (c *CPU) Info() *CPU {
	c.NumCPU = runtime.NumCPU()
	c.NumCores = int(runtime.NumCgoCall())
	c.MaxFrequency = 0
	c.Usage = c.GetUsage()

	return c
}

func (c *CPU) GetUsage() float64 {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	return float64(stat.Sys) / float64(1<<20)
}
