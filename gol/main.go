package main

import "flag"

// golParams provides the details of how to run the Game of Life and which image to load.
type golParams struct {
	turns       int
	imageWidth  int
	imageHeight int
}

// cell is used as the return type for the testing framework.
type cell struct {
	x, y int
}

// gameOfLife is the function called by the testing framework.
// It returns an array of alive cells.
func gameOfLife(p golParams, initialWorld [][]byte) [][]byte {

	world := initialWorld

	for turn := 0; turn < p.turns; turn++ {
		world = calculateNextState(p, world)
	}

	return world
}

// main is the function called when starting Game of Life
// Do not edit.
func main() {
	var params golParams
	var inputFilename string
	var outputFilename string

	flag.IntVar(
		&params.imageWidth,
		"w",
		256,
		"Specify the width of the image. Defaults to 256.")

	flag.IntVar(
		&params.imageHeight,
		"h",
		256,
		"Specify the height of the image. Defaults to 256.")

	flag.IntVar(
		&params.turns,
		"t",
		300,
		"Specify the number of turns to complete. Defaults to 300.")

	flag.StringVar(
		&inputFilename,
		"i",
		"images/secret.pgm",
		"Input file. Defaults to images/secret.pgm")

	flag.StringVar(
		&outputFilename,
		"o",
		"output.pgm",
		"Output file. Defaults to output.pgm")

	flag.Parse()

	initialWorld := readPgmImage(params, inputFilename)

	finalWorld := gameOfLife(params, initialWorld)

	writePgmImage(params, finalWorld, outputFilename)
}
