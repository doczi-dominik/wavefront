package main

type Node struct {
	x     int
	y     int
	value int
}

// plan fills the waveMap by walking from the end
// to the start and storing a distance value based
// on a counter, the step type (orthogonal or diagonal)
// and the value of its parent
//
// Example:
// ```
// -----       -----
// |S  |       |S34|
// |   |  -->  |332|
// |  F|       |42F|
// -----       -----
// ```
//
// This allows easy traversal as `printSteps()` demonstrates.
func plan(m *Map) {
	queue := generateNeighbours(m, m.endX, m.endY, 0)
	distance := 1
	unreachable := true

	for len(queue) > 0 {
		newQueue := []Node{}

		for _, neighbour := range queue {
			mapVal := m.wavemap[neighbour.y][neighbour.x]

			if mapVal == WAVEMAP_END {
				unreachable = false
			}

			// Avoids infinite loops between adjacent
			// nodes who consider each others as neighbours
			if mapVal != WAVEMAP_EMPTY {
				continue
			}

			m.wavemap[neighbour.y][neighbour.x] = mapVal + neighbour.value

			newQueue = append(
				newQueue,
				generateNeighbours(m, neighbour.x, neighbour.y, distance)...,
			)

			distance++
		}

		queue = newQueue
	}

	if unreachable {
		panic("No route to end", nil)
	}
}
