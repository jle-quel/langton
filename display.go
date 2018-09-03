package	main

import "github.com/buger/goterm"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	print_block(ant Ant, matrix [][]uint8, x int, y int) {
	if ant.X == x && ant.Y == y {
		goterm.Printf("%s", goterm.Color("*", goterm.RED))
	} else if matrix[y][x] == 0 {
		goterm.Printf("%s", goterm.Color(".", goterm.WHITE))
	} else {
		goterm.Printf("%s", goterm.Color(".", goterm.BLACK))
	}
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	print_flag(index int, environ Environ, width int, height int) {
	var title	string
	var x		int
	var y		int

	goterm.MoveCursor(0, 1)
	goterm.Printf("step: %d/%d", index, environ.Limit)

	title = "Langton's Ant"
	x = (width / 2) - (len(title) / 2)
	y = height / 16

	goterm.MoveCursor(x, y)
	goterm.Printf("%s", title)
}

func	print_matrix(ant Ant, matrix [][]uint8, width int, height int) {
	var length	int
	var pos_x	int
	var pos_y	int

	length = len(matrix)
	pos_y = height / 8

	for y, _ := range matrix {
		pos_x = (width / 2) - (length / 1)
		for x, _ := range matrix[y] {
			goterm.MoveCursor(pos_x, pos_y)
			print_block(ant, matrix, x, y)
			pos_x += 2
		}
		pos_y += 1
	}
}
