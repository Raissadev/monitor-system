package stats

import (
	"fmt"
	"system/src/sys"
	"time"
)

type Channel struct {
}

var cpu sys.CPU
var mem sys.Memory

func (c *Channel) Exec() {
	go func() {
		_cpu := make(chan float64)
		_mem := make(chan float64)

		go func() {
			for {
				cpu.Usage = cpu.GetUsage()
				_cpu <- cpu.Usage

				mem.Usage = mem.GetUsage()
				_mem <- mem.Usage

				time.Sleep(time.Second)
			}
		}()

		for {
			select {
			case cpu := <-_cpu:
				fmt.Println(cpu)

			case mem := <-_mem:
				fmt.Println(mem)
			}
		}
	}()
}
