package sys

import "runtime"

type Memory struct {
	Allocated      uint
	TotalAllocated uint
	Sys            uint
	NumGC          uint
	Usage          float64
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

func (m *Memory) GetUsage() float64 {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	return float64(stat.Alloc) / 1024 / 1024
}
