package	main

////////////////////////////////////////////////////////////////////////////////
/// STRUCTURES 
////////////////////////////////////////////////////////////////////////////////

type Ant			struct {
	X				int
	Y				int
	Direction		uint8
}

type Environ		struct {
	Limit			int
	Speed			int
}

type Function		struct {
	Name			string
	Call			func([]string) error
}

////////////////////////////////////////////////////////////////////////////////
/// TYPES 
////////////////////////////////////////////////////////////////////////////////

type StateMachine	[]Function

////////////////////////////////////////////////////////////////////////////////
/// CONST 
////////////////////////////////////////////////////////////////////////////////

const WIDTH_MIN = 62
const HEIGHT_MIN = 25