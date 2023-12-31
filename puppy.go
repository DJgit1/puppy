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

func From12() {
	fmt.Println("Hi, I'm from version 1.2.0")
}

func From13() {
	fmt.Println("Hi, I'm from version 1.3.0")
}

func From14() {
	fmt.Println("Hi, I'm from version 1.4.0")
}

func From15() {
	fmt.Println("Hi, I'm from version 1.5.0")
}
