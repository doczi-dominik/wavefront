package main

const (
	// ORT_VALUE represents the value increment
	// of an orthogonal step
	ORT_VALUE = 2
	// DIAG_VALUE represents the value increment
	// of a diagonal step
	DIAG_VALUE = 3
)

func generateNeighbours(m *Map, x int, y int) []Node {
	nodes := []Node{
		{x: x - 1, y: y, value: ORT_VALUE},
		{x: x, y: y - 1, value: ORT_VALUE},
		{x: x, y: y + 1, value: ORT_VALUE},
		{x: x + 1, y: y, value: ORT_VALUE},

		{x: x - 1, y: y - 1, value: DIAG_VALUE},
		{x: x - 1, y: y + 1, value: DIAG_VALUE},
		{x: x + 1, y: y - 1, value: DIAG_VALUE},
		{x: x + 1, y: y + 1, value: DIAG_VALUE},
	}

	// Avoids an extra array allocation
	// by creating a 0-length slice from the existing
	// array
	filtered := nodes[:0]

	for _, node := range nodes {
		xInbounds := node.x >= 0 && node.x <= m.width-1
		yInbounds := node.y >= 0 && node.y <= m.height-1

		if xInbounds && yInbounds {
			filtered = append(filtered, node)
		}
	}

	return filtered
}
