package stats

import (
	"fmt"
	"runtime"
)

type Consumption struct {
}

func (b *Consumption) memory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	usage := m.Alloc

	usageMB := float64(usage) / 1024 / 1024

	fmt.Printf("Consumo de memória: %.2f MB\n", usageMB)

}
