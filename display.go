package	main

import "github.com/buger/goterm"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	(self *Ant) Draw(matrix [][]uint8, x int, y int) {
	if (*self).X == x && (*self).Y == y {
		goterm.Printf("🐜")
	} else if matrix[y][x] == 0 {
		goterm.Printf("◻️")
	} else {
		goterm.Printf("◼️")
	}
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	flag(index int, environ Environ, width int, height int) {
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

func	display(ant Ant, matrix [][]uint8, width int, height int) {
	var length	int
	var pos_x	int
	var pos_y	int

	length = len(matrix)
	pos_y = height / 8

	for y, _ := range matrix {
		pos_x = (width / 2) - (length / 1)
		for x, _ := range matrix[y] {
			goterm.MoveCursor(pos_x, pos_y)
			ant.Draw(matrix, x, y)
			pos_x += 2
		}
		pos_y += 1
	}
}
