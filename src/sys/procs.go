package sys

import (
	"os/exec"
	"strings"
)

type Proc struct {
}

func (p *Proc) ProcessLs() ([]string, error) {
	cmd := exec.Command("ps", "-e", "-o", "pid,ppid,user,%cpu,%mem,command")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(output), "\n")[1:]

	procs := make([]string, 0)
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row != "" {
			procs = append(procs, row)
		}
	}

	return procs, nil
}
