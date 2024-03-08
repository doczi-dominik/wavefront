package main

import "fmt"

// DIRECTION_TO_STRING is used to convert an offset between
// positions to a short string describing the direction
var DIRECTION_TO_STRING = map[int]map[int]string{
	-1: {
		-1: "UL",
		0:  "U",
		1:  "UR",
	},
	0: {
		-1: "L",
		1:  "R",
	},
	1: {
		-1: "DL",
		0:  "D",
		1:  "DR",
	},
}

// printSteps starts walking from the starting
// position to the ending position by always
// choosing the tile with the lowest valid
// value
//
// Call `plan()` before to set up the wavemap
func printSteps(m *Map) {
	currentX := m.startX
	currentY := m.startY

	for !(currentX == m.endX && currentY == m.endY) {
		neighbours := generateNeighbours(m, currentX, currentY)

		var smallestNode *Node

		for _, node := range neighbours {
			nodeMapValue := m.wavemap[node.y][node.x]

			// Invalid moves are:
			// - stepping on the start tile
			// - stepping on a wall
			isValidMove := nodeMapValue == WAVEMAP_END || nodeMapValue > 0

			// take the first valid move as initial smallest
			if smallestNode == nil {
				if isValidMove {
					smallestNode = &node
				}
				continue
			}

			smallestMapValue := m.wavemap[smallestNode.y][smallestNode.x]

			smallerThan := nodeMapValue < smallestMapValue

			if isValidMove && smallerThan {
				smallestNode = &node
			}
		}

		if smallestNode == nil {
			panic("Cannot leave start position", nil)
		}

		nextX := smallestNode.x
		nextY := smallestNode.y

		diffX := nextX - currentX
		diffY := nextY - currentY

		fmt.Print(DIRECTION_TO_STRING[diffY][diffX] + " ")

		currentX = nextX
		currentY = nextY
	}

	fmt.Println()
}
