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
