package robotTourOptimization

import (
	"math"
)

type Point struct {
	x int
	y int
}

func distance(a, b Point) float64 {
	x := 0
	if a.x < b.x {
		x = b.x - a.x
	} else {
		x = a.x - b.x
	}
	y := 0
	if a.y < b.y {
		y = b.y - a.y
	} else {
		y = a.y - b.y
	}

	return math.Sqrt(math.Pow(float64(x), 2.0) + math.Pow(float64(y), 2.0))
}

func NearestNeighbor(points []Point) float64 {
	result := 0.0
	visited := map[int]bool{}
	p := points[0]
	visited[0] = true

	for {
		var indexOfNearest int
		dist := math.MaxFloat64

		for k, p1 := range points {
			if _, prs := visited[k]; prs {
				continue
			}
			newDist := distance(p, p1)
			if newDist < dist {
				indexOfNearest = k
				dist = newDist
			}
		}
		visited[indexOfNearest] = true
		p = points[indexOfNearest]
		result += dist

		// Return to start
		if len(visited) == len(points) {
			result += distance(p, points[0])
			break
		}
	}
	return result
}
