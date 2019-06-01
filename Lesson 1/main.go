// Main package - IE the package to start in
package main

import "fmt"

// Program starts here
func main() {
	c := NewCharacter("Kid Goku")
	c.broke = true
	c.PrintBroke()
}

// Beginning character struct type
type Character struct {
	name  string
	broke bool
}

// Method to create a new character type
func NewCharacter(name string) *Character {
	return &Character{name: name}
}

// Method to print whether or not the character sucks
func (c *Character) PrintBroke() {
	if c.broke {
		fmt.Println(c.name + " is broke")
	} else {
		fmt.Println(c.name + " is not broke")
	}
}
