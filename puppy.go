package puppy

import (
	"fmt"

	"github.com/DJgit1/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGorwnUp(Bark())
}

func BigBarks() string {
	return dog.WhenGorwnUp((Barks()))
}

func From11() {
	fmt.Println("Hi, I'm from version 1.1.0")
}
