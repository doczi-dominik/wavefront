package main

// These are the special values in a Wavemap.
//
// Positive integers represent distance values
// from the finish to the start position.
//
// By using non-positive values and ordering them
// in descending order, one can easily restrict
// or filter the current types using
// numerical comparisons (<, <=, >, >=).
const (
	WAVEMAP_EMPTY = -iota
	WAVEMAP_END
	WAVEMAP_START
	WAVEMAP_WALL
)

type Wavemap [][]int
