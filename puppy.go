package puppy

import "github.com/kindness07/Dog"

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}
func BigBark() string {
	return Dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return Dog.WhenGrownUp(Barks())
}
