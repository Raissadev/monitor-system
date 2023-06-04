package term

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init: %v", err)
	}

	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hey"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
