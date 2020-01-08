package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dshulyak/seriesrw"
	"github.com/wcharczuk/go-chart"
)

func PlotMemory(toPlot []plot, step int, out string) error {
	series := make([]chart.Series, 0, len(toPlot))
	prefix := make([]byte, 0, len(toPlot))

	for _, plot := range toPlot {
		single := chart.ContinuousSeries{
			Name:    plot.name,
			XValues: []float64{},
			YValues: []float64{},
		}
		if len(plot.name) > 0 {
			prefix = append(prefix, plot.name[0])
		}

		r, err := seriesrw.NewFileReader(plot.path, 2048)
		if err != nil {
			return err
		}
		defer r.Close()
		total := 0
		val := uint64(0)
		for err := r.Read(&val); err == nil; err = r.Read(&val) {
			total += step
			single.YValues = append(single.YValues, float64(val)/1024/1024)
			single.XValues = append(single.XValues, float64(total)/1000)
		}
		series = append(series, single)
	}

	f, _ := os.Create(filepath.Join(out, fmt.Sprintf("memory-%s-%d.png", prefix, step)))
	defer f.Close()
	graph := chart.Chart{
		Width:  800,
		Height: 400,
		XAxis: chart.XAxis{
			Name: "Entries (thousands)",
		},
		YAxis: chart.YAxis{
			Name: "Memory (MBs)",
		},
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	return graph.Render(chart.PNG, f)
}

func PlotTimeSpent(toPlot []plot, step int, out string) error {
	series := make([]chart.Series, 0, len(toPlot))
	prefix := make([]byte, 0, len(toPlot))

	for _, plot := range toPlot {
		single := chart.ContinuousSeries{
			Name:    plot.name,
			XValues: []float64{},
			YValues: []float64{},
		}
		if len(plot.name) > 0 {
			prefix = append(prefix, plot.name[0])
		}

		r, err := seriesrw.NewFileReader(plot.path, 2048)
		if err != nil {
			return err
		}
		defer r.Close()
		total := 0
		val := uint64(0)
		for err := r.Read(&val); err == nil; err = r.Read(&val) {
			total += step
			single.YValues = append(single.YValues, float64(time.Duration(val).Milliseconds()))
			single.XValues = append(single.XValues, float64(total)/1000)
		}
		series = append(series, single)
	}

	f, _ := os.Create(filepath.Join(out, fmt.Sprintf("time-%s-%d.png", prefix, step)))
	defer f.Close()
	graph := chart.Chart{
		Width:  800,
		Height: 400,
		XAxis: chart.XAxis{
			Name: "Entries (thousands)",
		},
		YAxis: chart.YAxis{
			Name:           "Milliseconds",
			ValueFormatter: chart.FloatValueFormatter,
		},
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	return graph.Render(chart.PNG, f)
}
