package	main

import "os"
import "fmt"
import "strconv"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func	help(argv []string) error {
	fmt.Println("Langton's Ant simulation")
	fmt.Println()

	fmt.Println("Usage:\n\n\tlangton [command or option] [argument]")
	fmt.Println()

	fmt.Println("The commands are:\n")
	fmt.Println("\thelp\tinformation about langton")
	fmt.Println()

	fmt.Println("The options are:\n")
	fmt.Println("\t-s\tnumber of step by millisecond done by the ant(s)")
	fmt.Println("\t-l\tnumber of step until the simulation finish")
	fmt.Println()

	fmt.Println("If no argv is given, the default value are [speed: 128] [limit: 10000]") 
	os.Exit(0)

	return nil
}

func	(self *Environ) GetSpeed(argv []string) error {
	var result	int
	var err		error

	if len(argv) == 1 {
		return fmt.Errorf("usage: langton speed [step by millisecond]")
	}
	if result, err = strconv.Atoi(argv[1]); err != nil {
		return fmt.Errorf("error: speed \"%s\" is not supported", argv[1])
	}
	if result <= 0 {
		return fmt.Errorf("error: speed cannot be less or equal to 0")
	}

	(*self).Speed = result
	return nil
}

func	(self *Environ) GetLimit(argv []string) error {
	var result	int
	var err		error

	if len(argv) == 1 {
		return fmt.Errorf("usage: langton limit [number_of_step]")
	}
	if result, err = strconv.Atoi(argv[1]); err != nil {
		return fmt.Errorf("error: limit \"%s\" is not supported", argv[1])
	}
	if result <= 0 {
		return fmt.Errorf("error: limit cannot be less or equal to 0")
	}

	(*self).Limit = result
	return nil
}

func	(self *StateMachine) GetFunction(argv string) (int, error) {
	for index, _ := range (*self) {
		if argv == (*self)[index].Name {
			return index, nil
		}
	}

	return 0, fmt.Errorf("error: unknown argument \"%s\"", argv)
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	(self *Environ) Get() error {
	var length		int
	var num			int
	var err			error
	var state_machine = StateMachine{{"-l", (*self).GetLimit},
									 {"-s", (*self).GetSpeed},
									 {"help", help},
	}

	(*self).Limit = 10000
	(*self).Speed = 128
	length = len(os.Args)

	for index := 1; index < length; index += 2 {
		if num, err = state_machine.GetFunction(os.Args[index]); err == nil {
			if err = state_machine[num].Call(os.Args[index:]); err != nil {
				return err
			}
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	
	return nil 
}
