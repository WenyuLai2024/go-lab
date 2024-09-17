package main

import (
	"fmt"
	"testing"
)

type args struct {
	p             golParams
	expectedAlive []cell
}

type test struct {
	name string
	args args
}

func TestGol(t *testing.T) {
	tests := []test{
		{"0 turns", args{
			p: golParams{
				turns:       0,
				imageWidth:  16,
				imageHeight: 16,
			},
			expectedAlive: []cell{
				{x: 4, y: 5},
				{x: 5, y: 6},
				{x: 3, y: 7},
				{x: 4, y: 7},
				{x: 5, y: 7},
			},
		}},

		{"1 turn", args{
			p: golParams{
				turns:       1,
				imageWidth:  16,
				imageHeight: 16,
			},
			expectedAlive: []cell{
				{x: 3, y: 6},
				{x: 5, y: 6},
				{x: 4, y: 7},
				{x: 5, y: 7},
				{x: 4, y: 8},
			},
		}},

		{"100 turns", args{
			p: golParams{
				turns:       100,
				imageWidth:  16,
				imageHeight: 16,
			},
			expectedAlive: []cell{
				{x: 12, y: 0},
				{x: 13, y: 0},
				{x: 14, y: 0},
				{x: 13, y: 14},
				{x: 14, y: 15},
			},
		}},
	}

	// Run normal tests
	for _, test := range tests {
		initial16x16World := readPgmImage(golParams{turns: 0, imageWidth: 16, imageHeight: 16}, "images/16x16.pgm")
		testName := fmt.Sprintf("%dx%dx%d", test.args.p.imageWidth, test.args.p.imageHeight, test.args.p.turns)
		t.Run(testName, func(t *testing.T) {
			world := gameOfLife(test.args.p, initial16x16World)
			aliveCells := calculateAliveCells(test.args.p, world)
			assertEqualBoard(t, aliveCells, test.args.expectedAlive, test.args.p)
		})
	}
}

func boardFail(t *testing.T, given, expected []cell, p golParams) bool {
	errorString := fmt.Sprintf("-----------------\n\n  FAILED TEST\n  16x16\n  %d Turns\n", p.turns)
	errorString = errorString + AliveCellsToString(given, expected, p.imageWidth, p.imageHeight)
	t.Error(errorString)
	return false
}

func assertEqualBoard(t *testing.T, given, expected []cell, p golParams) bool {
	givenLen := len(given)
	expectedLen := len(expected)

	if givenLen != expectedLen {
		return boardFail(t, given, expected, p)
	}

	visited := make([]bool, expectedLen)
	for i := 0; i < givenLen; i++ {
		element := given[i]
		found := false
		for j := 0; j < expectedLen; j++ {
			if visited[j] {
				continue
			}
			if expected[j] == element {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return boardFail(t, given, expected, p)
		}
	}

	return true
}
