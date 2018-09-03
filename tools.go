package	main

import "os"
import "fmt"
import "os/exec"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	get_size(width int, height int) int {
	var size	int

	size = height - (height / 8)

	for size > width {
		size -= 1
	}

	return size
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	get_matrix(width int, height int) [][]uint8 {
	var size	int
	var result	[][]uint8

	size = get_size(width, height)
	result = make([][]uint8, size)
	
	for y := 0; y < size; y += 1 { 
		result[y] = make([]uint8, size)
		for x := 0; x < size; x += 1 {
			result[y][x] = 0
		}
	}

	return result
}

func 	tput(str string) {
	var cmd	*exec.Cmd
	var err	error

	cmd = exec.Command("tput", str)
	cmd.Stdout = os.Stdout
	
	if err = cmd.Run(); err != nil {
		_exit(err)
	}
}

func	_exit(err interface{}) {
	tput("cvvis")
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
