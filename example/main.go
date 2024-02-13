package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"

	kmeans "github.com/sazonovItas/algorithm-kmeans"
	"github.com/sazonovItas/algorithm-kmeans/clusters"
	"github.com/sazonovItas/algorithm-kmeans/plotter"
)

const (
	outPng = "k-means.png"
)

var (
	index = `
  <html>
  <head>
    <link rel="stylesheet" href="style.css">
  </head>
  <body>
    <img src="./k-means.png" />
  </body>
  </html>
  `
	cntPoints   int = 1024
	cntClusters int = 10
)

type Color struct {
	colorful.Color
}

func main() {
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err == nil {
			cntPoints = v
		}
	}

	if len(os.Args) > 2 {
		v, err := strconv.Atoi(os.Args[2])
		if err == nil {
			cntClusters = v
		}
	}

	html, err := os.Create("index.html")
	if err != nil {
		log.Printf("error creating index.html: %s", err.Error())
		return
	}
	defer html.Close()

	_, err = html.WriteString(index)
	if err != nil {
		log.Printf("error write to file: %s", err.Error())
		return
	}

	var dataset clusters.Observations
	for x := 0; x < cntPoints; x++ {
		dataset = append(dataset, clusters.Coordinates{
			rand.Float64(),
			rand.Float64(),
		})
	}

	Kmeans, err := kmeans.NewWithOptions(100, plotter.SimplePlotter{Height: 800, Width: 800}, 0.05)
	if err != nil {
		log.Printf("error to create kmeans")
		return
	}

	clusters, _ := Kmeans.Partition(dataset, cntClusters, "./"+outPng)
	for i, c := range clusters {
		fmt.Printf("Cluster: %d, points: %d\n", i, len(c.Observations))
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
	}
}
