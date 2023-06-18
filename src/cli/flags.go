package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
)

type Flags struct {
}

func (f *Flags) ManArgs() []string {
	args := make([]string, 0)

	if len(os.Args[1:]) == 0 {
		return args
	}

	help := flag.Bool("help", false, "Display help message")
	specs := flag.String("show", "", "Specify a comma-separated list of specifications")

	// ? acronyms
	flag.BoolVar(help, "h", false, "Display help message")
	flag.StringVar(specs, "s", "", "Specify a comma-separated list of specifications (cpu,netrwork,swap,memory,procs,disk,entropy-kernel)")

	flag.Parse()

	if *help {
		f.printHelp()
		args = append(args, "help")
		return args
	}

	switch {
	case *specs != "":
		stats := strings.Split(*specs, ",")
		for _, spec := range stats {
			switch spec {
			case "cpu":
				args = append(args, "cpu")
				continue
			case "network":
				args = append(args, "network")
			case "swap":
				args = append(args, "swap")
			case "memory":
				args = append(args, "memory")
			case "procs", "processes":
				args = append(args, "procs")
			case "disk":
				args = append(args, "disk")
			case "entropy", "entropy-kernel":
				args = append(args, "entropy-kernel")
			default:
				return args
			}
		}

	default:
		return args
	}

	return args
}

func (f *Flags) printHelp() {
	fmt.Println("Usage: kenbunshoku-haki [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func (f *Flags) Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func (f *Flags) InterfaceSlice(slice []ui.GridItem) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
