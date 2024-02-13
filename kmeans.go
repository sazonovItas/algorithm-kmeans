package kmeans

import (
	"fmt"
	"math/rand"

	"github.com/sazonovItas/algorithm-kmeans/clusters"
	"github.com/sazonovItas/algorithm-kmeans/plotter"
)

type Kmeans struct {
	plotter    Plotter
	minPercent float64
	iterations int
}

type Plotter interface {
	Plot(clusters.Clusters, string) error
}

func NewWithOptions(iterations int, plotter Plotter, percentChanges float64) (Kmeans, error) {
	if percentChanges < 0.0 || percentChanges > 1.0 {
		return Kmeans{}, fmt.Errorf("percent of points changed cluster must be 0.0 < x < 1.0")
	}

	return Kmeans{
		plotter:    plotter,
		iterations: iterations,
		minPercent: percentChanges,
	}, nil
}

func New() Kmeans {
	return Kmeans{
		plotter: plotter.SimplePlotter{
			Width:  1024,
			Height: 1024,
		},
		iterations: 96,
		minPercent: 0.1,
	}
}

func (km Kmeans) Partition(
	dataset clusters.Observations,
	k int, toimg string,
) (clusters.Clusters, error) {
	if k > len(dataset) {
		return nil, fmt.Errorf("the size of data set must be greater or equal k")
	}

	cc, err := clusters.New(k, dataset)
	if err != nil {
		return cc, err
	}

	changes := 1
	pointsByCluster := make([]int, len(dataset))

	for i := 0; changes > 0; i++ {
		changes = 0

		cc.Reset()
		for p, point := range dataset {
			ci := cc.Nearest(point)
			cc[ci].Append(point)

			if pointsByCluster[p] != ci {
				pointsByCluster[p] = ci
				changes++
			}
		}

		for ci := 0; ci < len(cc); ci++ {
			if len(cc[ci].Observations) == 0 {

				var ri int
				for {
					ri = rand.Intn(len(dataset))
					if len(cc[pointsByCluster[ri]].Observations) > 1 {
						break
					}
				}

				cc[ci].Append(dataset[ri])
				pointsByCluster[ri] = ci
				changes = len(dataset)
			}
		}

		if changes > 0 {
			cc.Recenter()
		}

		err := km.plotter.Plot(cc, toimg)
		if err != nil {
			return nil, fmt.Errorf("error to save plot to the file")
		}

		if i == km.iterations || changes < int(float64(len(dataset)*int(km.minPercent))) {
			break
		}

	}

	return cc, nil
}
