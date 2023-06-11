package stats

import (
	"fmt"
	"runtime"
)

type Bench struct {
}

func (b *Bench) view() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	usage := m.Alloc

	usageMB := float64(usage) / 1024 / 1024

	fmt.Printf("Consumo de memória: %.2f MB\n", usageMB)

}
