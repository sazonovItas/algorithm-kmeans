package clusters

import (
	"fmt"
	"log"
	"math/rand"
)

type Cluster struct {
	Center       Coordinates
	Observations Observations
}

type Clusters []Cluster

func New(k int, dataset Observations) (Clusters, error) {
	var c Clusters
	if len(dataset) == 0 || len(dataset[0].Coordinates()) == 0 {
		return nil, fmt.Errorf("there is no mean for an empty dataset or 0 demensional coordinates")
	}
	if k <= 0 {
		return nil, fmt.Errorf("amount of the clusters in dataset should be positive")
	}

	for i := 0; i < k; i++ {
		var p Coordinates
		for j := 0; j < len(dataset[0].Coordinates()); j++ {
			p = append(p, rand.Float64())
		}

		c = append(c, Cluster{
			Center:       p,
			Observations: Observations{},
		})
	}

	return c, nil
}

func (c *Cluster) Append(point Observation) {
	c.Observations = append(c.Observations, point)
}

func (cs Clusters) Nearest(point Observation) int {
	ci, dist := 0, -1.0

	for i, v := range cs {
		d := v.Center.Distance(point.Coordinates())

		if dist < 0 || dist > d {
			ci = i
			dist = d
		}
	}

	return ci
}

func (cs *Cluster) FarthestPoint() Coordinates {
	dist, nd := -1.0, 0

	for i, v := range cs.Observations {
		d := cs.Center.Coordinates().Distance(v.Coordinates())

		if d > dist {
			nd = i
			dist = d
		}
	}

	point := make([]float64, len(cs.Observations[nd].Coordinates()))
	copy(point, cs.Observations[nd].Coordinates())
	return point
}

func (cs Clusters) AverageCentersDist() float64 {
	size := 0

	var dist float64
	for i := range cs {
		for j := i + 1; j < len(cs); j++ {
			size++
			dist += cs[i].Center.Distance(cs[j].Center)
		}
	}

	return dist / float64(size)
}

func (c *Cluster) Recenter() {
	center, err := c.Observations.Center()
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	c.Center = center
}

func (cs Clusters) Recenter() {
	for i := range cs {
		cs[i].Recenter()
	}
}

func (cs Clusters) Reset() {
	for i := range cs {
		cs[i].Observations = Observations{}
	}
}

func (c Cluster) PointsInDimension(n int) Coordinates {
	v := make([]float64, 0)
	for i := range c.Observations {
		v = append(v, c.Observations[i].Coordinates()[n])
	}

	return v
}

func (cs Clusters) CentersInDimension(n int) Coordinates {
	v := make([]float64, 0)
	for i := range cs {
		v = append(v, cs[i].Center.Coordinates()[n])
	}

	return v
}
