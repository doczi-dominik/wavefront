package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INPUT_EMPTY = '0'
	INPUT_END   = 'F'
	INPUT_START = 'S'
	INPUT_WALL  = '1'
)

type Map struct {
	width  int
	height int

	startX int
	startY int

	endX int
	endY int

	wavemap Wavemap
}

func parseLine(line []byte) (out []int, startX int, endX int) {
	out = make([]int, len(line))
	startX = -1
	endX = -1

	for i, char := range line {
		switch char {
		case INPUT_EMPTY:
			out[i] = WAVEMAP_EMPTY
		case INPUT_END:
			out[i] = WAVEMAP_END
			endX = i
		case INPUT_START:
			out[i] = WAVEMAP_START
			startX = i
		case INPUT_WALL:
			out[i] = WAVEMAP_WALL
		default:
			panic("Unrecognized character", fmt.Errorf("%c (code %v)", char, char))
		}
	}

	return out, startX, endX
}

func parse(filename string) *Map {
	m := &Map{
		startX: -1,
		startY: -1,

		endX: -1,
		endY: -1,
	}

	waveMap := Wavemap{}

	f, err := os.Open(filename)

	if err != nil {
		panic("Could not read file", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		parsed, startX, endX := parseLine(line)

		if startX != -1 {
			if m.startX != -1 {
				panic("Multiple start positions found", nil)
			}

			m.startX = startX
			m.startY = lineIndex
		}

		if endX != -1 {
			if m.endX != -1 {
				panic("Multiple end positions found", nil)
			}

			m.endX = endX
			m.endY = lineIndex
		}

		waveMap = append(waveMap, parsed)

		lineIndex++
	}

	if m.startX == -1 {
		panic("No start position", nil)
	}

	if m.endX == -1 {
		panic("No end position", nil)
	}

	m.width = len(waveMap[1])
	m.height = lineIndex
	m.wavemap = waveMap

	return m
}
