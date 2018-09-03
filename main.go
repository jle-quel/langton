package	main

import "syscall"
import "golang.org/x/crypto/ssh/terminal"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	main() {
	var width	int
	var height	int
	var environ	Environ
	var err		error

	if width, height, err = terminal.GetSize(syscall.Stdin); err != nil {
		_exit("error: couldn't get size of terminal")
	}
	if width < WIDTH_MIN || height < HEIGHT_MIN {
		_exit("error: terminal window is too small") 
	}
	if err = environ.Get(); err != nil {
		_exit(err)
	}

	tput("civis")
	go handle_signal()
	go handle_window(width, height)
	langton(width, height, environ)
	tput("cvvis")
}
