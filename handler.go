package	main

import "os"
import "time"
import "syscall"
import "os/signal"
import "golang.org/x/crypto/ssh/terminal"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	handle_signal() {
	var channel	chan os.Signal

	channel = make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<- channel

	_exit("")
}

func	handle_window(default_width int, default_height int) {
	var width	int
	var height	int
	var err		error

	for {
		if width, height, err = terminal.GetSize(syscall.Stdin); err != nil {
			_exit("error: couldn't get size of terminal")
		}
		if width < default_width || height < default_height {
			_exit("")
		}
		time.Sleep(1 * time.Second)
	}
}
