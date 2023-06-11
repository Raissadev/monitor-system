package tests

import (
	"system/src/sys"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var procs sys.Proc

type Render struct {
}

func (r *Render) Test() {
	p := widgets.NewParagraph()
	p.Text = "42"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	_msgs := make(chan string)

	go func() {
		for {
			_msgs <- "¬¬ sasageyo"
			time.Sleep(time.Second)

		}
	}()

	go func() {
		for {
			select {
			case new := <-_msgs:
				ui.Clear()
				p.Text = new
				ui.Render(p)
			}
		}
	}()
}
