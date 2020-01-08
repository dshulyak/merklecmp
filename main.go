package main

import (
	"flag"
	"log"
	"strings"
)

var (
	plots    = flag.String("plots", "", "list of comma separated name to a file that will be plotted, e.g. iavl=iavl.out")
	plotType = flag.String("type", "memory", "memory|time")
	dir      = flag.String("dir", "_assets/", "output directory for a plot")
	step     = flag.Int("step", 0, "step that was used for plot collection, will be on x-axis")
)

type plot struct {
	name string
	path string
}

func parsePlots(plots string) []plot {
	parts := strings.Split(plots, ",")
	rst := make([]plot, len(parts))
	for i, part := range parts {
		np := strings.Split(part, "=")
		if len(np) != 2 {
			log.Fatalf("part has a wrong format: %v", part)
		}
		rst[i].name = np[0]
		rst[i].path = np[1]
	}
	return rst
}

func main() {
	flag.Parse()
	switch *plotType {
	case "memory":
		if err := PlotMemory(parsePlots(*plots), *step, *dir); err != nil {
			log.Fatalf("failed to plot memory: %v", err)
		}
	case "time":
		if err := PlotTimeSpent(parsePlots(*plots), *step, *dir); err != nil {
			log.Fatalf("failed to plot time: %v", err)
		}
	default:
		log.Fatalf("unknown plot type: %v", *plotType)
	}
}
