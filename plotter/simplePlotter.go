package plotter

import (
	"bytes"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"

	"github.com/sazonovItas/algorithm-kmeans/clusters"
)

// 2d dimensional plotter
type SimplePlotter struct {
	Width  int
	Height int
}

// A monokai-ish color palette
var colors = []drawing.Color{
	drawing.ColorFromHex("f92672"),
	drawing.ColorFromHex("89bdff"),
	drawing.ColorFromHex("66d9ef"),
	drawing.ColorFromHex("67210c"),
	drawing.ColorFromHex("7acd10"),
	drawing.ColorFromHex("af619f"),
	drawing.ColorFromHex("fd971f"),
	drawing.ColorFromHex("dcc060"),
	drawing.ColorFromHex("545250"),
	drawing.ColorFromHex("4b7509"),
}

// Plot draw a 2-dimensional data set into a PNG file toimg
func (p SimplePlotter) Plot(cc clusters.Clusters, toimg string) error {
	var series []chart.Series

	// draw data points
	for i := 0; i < len(cc); i++ {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    4,
			},
			XValues: cc[i].PointsInDimension(0),
			YValues: cc[i].PointsInDimension(1),
		})
	}

	// draw cluster center points
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			StrokeWidth: chart.Disabled,
			DotColor:    drawing.ColorBlack,
			DotWidth:    8,
		},
		XValues: cc.CentersInDimension(0),
		YValues: cc.CentersInDimension(1),
	})

	graph := chart.Chart{
		Height: p.Height,
		Width:  p.Width,
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}

	img, err := os.Create(toimg)
	if err != nil {
		return err
	}
	defer img.Close()

	_, err = buffer.WriteTo(img)
	time.Sleep(time.Millisecond * 200)
	return err
}
