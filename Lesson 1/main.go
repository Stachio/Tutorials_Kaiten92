// Main package - IE the package to start in
package main

import (
	"log"
)

// Program starts here
func main() {
	var justin Character
	//justin.broke = true
	justin.name = "Justin"
	//justin.PrintBroke()
	PrintBroke(&justin)

	var c *Character
	c = NewCharacter("Bustin", true)
	//c := NewCharacter("Justin")
	//c.broke = false
	PrintBroke(c)
}

// Beginning character struct type
type Character struct {
	name  string
	broke bool
}

// Method to create a new character type
func NewCharacter(name2 string, broke2 bool) *Character {
	var charData Character
	charData.name = name2
	charData.broke = broke2
	//return &Character{name: name}
	return &charData
}

// Method to print whether or not the character sucks
func PrintBroke(c *Character) {
	if c.broke {
		log.Println(c.name + " is broke")
	} else {
		log.Println(c.name + " is not broke")
	}
}
