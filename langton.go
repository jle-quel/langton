package	main

import "time"
import "math/rand"
import "github.com/buger/goterm"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	(self *Ant) ChangePosition(matrix [][]uint8) {
	switch (*self).Direction {
		case 0:
			self.X += 1
		case 1:
			self.Y += 1
		case 2:
			self.X -= 1
		case 3:
			self.Y -= 1
	}
}

func	(self *Ant) CheckCollision(matrix [][]uint8) {
	if self.X >= len(matrix) || self.X < 0 {
			_exit("")
	}
	if self.Y >= len(matrix) || self.Y < 0 {
			_exit("")
	}
}

func	(self *Ant) ChangeDirection(matrix [][]uint8) {
	if matrix[self.Y][self.X] == 0 {
		self.Direction = (self.Direction + 1) % 4
	} else {
		self.Direction = (self.Direction - 1) % 4
	}
}

func	(self *Ant) Simulate(matrix [][]uint8) {
	matrix[(*self).Y][(*self).X] ^= 1

	(*self).ChangePosition(matrix)
	(*self).CheckCollision(matrix)
	(*self).ChangeDirection(matrix)
}

func	(self *Ant) Get(matrix [][]uint8) {
	var size	int

	size = len(matrix)

	(*self).X = size / 2
	(*self).Y = size  / 2
	(*self).Direction = uint8(rand.Int() % 3)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	langton(width int, height int, environ Environ) {
	var matrix	[][]uint8
	var ant		Ant

	matrix = get_matrix(width, height)
	ant.Get(matrix)
	goterm.Clear()
	
	for index := 0; index < environ.Limit; index += 1 {
		print_flag(index, environ, width, height)
		print_matrix(ant, matrix, width, height)
		ant.Simulate(matrix)
		goterm.Flush()
		time.Sleep(time.Duration(environ.Speed) * time.Millisecond)
	}
}
