package main

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {

	rend := render()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = rend.Render(w)
	})
	http.ListenAndServe(":9000", nil)

}

type Renderer interface {
	Render(w io.Writer) error
}

func render() Renderer {
	// create a new sankey instance
	sankey := charts.NewSankey()
	// set some global options like Title/Legend/ToolTip or anything else
	sankey.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))
	nodes := []opts.SankeyNode{
		{
			Name: "transfer",
		},
		{
			Name: "income",
		},
		{
			Name: "coffee",
		},
		{Name: "cash"},
	}
	links := []opts.SankeyLink{
		{
			Source: "transfer",
			Target: "cash",
			Value:  100.0,
		},
		{
			Source: "income",
			Target: "cash",
			Value:  100.0,
		},
		{
			Source: "cash",
			Target: "coffee",
			Value:  5.0,
		},
	}
	sankey.AddSeries("history", nodes, links).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Color:     "source",
				Curveness: 0.5,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return sankey
}
