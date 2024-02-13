package clusters

import (
	"fmt"
	"math"
)

type Coordinates []float64

type Observation interface {
	Coordinates() Coordinates
	Distance(point Coordinates) float64
}

type Observations []Observation

func (c Coordinates) Coordinates() Coordinates {
	return Coordinates(c)
}

func (c Coordinates) Distance(point Coordinates) float64 {
	var dist float64

	for i, v := range c {
		dist += math.Pow(point[i]-v, 2)
	}

	return dist
}

func (obs Observations) Center() (Coordinates, error) {
	size := len(obs)
	if size == 0 {
		return nil, fmt.Errorf("there is no mean for an empty set of points")
	}

	obsCenter := make([]float64, len(obs[0].Coordinates()))
	for _, point := range obs {
		for i, v := range point.Coordinates() {
			obsCenter[i] += v
		}
	}

	var center Coordinates
	for _, v := range obsCenter {
		center = append(center, v/float64(size))
	}
	return center, nil
}

const averageDistPrecision float64 = 1e-5

func AverageDistance(o Observation, observations Observations) float64 {
	var dist float64
	var size int
	for _, v := range observations {
		d := o.Distance(v.Coordinates())
		if d < averageDistPrecision {
			continue
		}

		size++
		dist += d
	}

	if size == 0 {
		return 0.0
	}

	return dist / float64(size)
}
