package reports

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
)

func GenerateGraph(stats map[string]int, title, filePath string) {
	/*
		This function generates a bar graph from the statistics.
	*/
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "Elements"
	p.Y.Label.Text = "Occurrences"

	barData := make(plotter.Values, len(stats))
	i := 0
	for _, count := range stats {
		barData[i] = float64(count)
		i++
	}

	bars, err := plotter.NewBarChart(barData, vg.Points(20))
	if err != nil {
		log.Fatalf("could not create bar chart: %v", err)
	}

	p.Add(bars)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, filePath); err != nil {
		log.Fatalf("could not save graph: %v", err)
	}
	fmt.Printf("Graph saved succefuly ! %s\n", filePath)
}
