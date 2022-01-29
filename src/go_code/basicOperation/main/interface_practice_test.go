package main

import "testing"

func TestInterface(t *testing.T) {
	redDuck := RedDuck{
		Duck{"GaGa", 2},
	}
	rubberDuck := RubberDuck{
		Duck{"LaLa", 3},
	}

	action(redDuck)
	action(rubberDuck)
}
