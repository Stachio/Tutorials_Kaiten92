# Lesson 1: As Little as Possible

## Running your go program
Navigate to the lesson1 folder through either git bash or cmd and type: `go run main.go`

## functions
All functions begin with the **func** keyword.  Functions are designed to help reduce repeatable code and decrease troubleshooting time.

```go
func boo() {
    // Do stuff
}
```

## The main function
The main function is where your golang program starts.  From there your program reads left to right, top to bottom until all execution is complete.  Execution generally ends at the end of the main function.

```go
func main() {
    // Your code starts here
}
```

## Function WANT variables
When you write a function you have the option to designate what data the function *wants*.  Functions are allowed to make demands because they can't see the data (variables) inside each other.  The main function has all the character data but the PrintBroke function can't see that data unless we give it over.

Recall the function below.

```go
func NewCharacter(name2 string, broke2 bool) *Character {
	var charData Character
	charData.name = name2
	charData.broke = broke2
	//return &Character{name: name}
	return &charData
}
```

NewCharacter *wants* name2 and broke2 because those are the parts we need to set within the character variable.

Here's an example where we avoid using NewCharacter for the same results.

```go
var justin Character
justin.name = "Justin"
justin.broke = true

var bustin Character
bustin.name = "Bustin"
bustin.broke = false

var stache Character
stache.name = "Stache"
stache.broke = false
//stache.lowtier = true

var eddie Character
eddie.name = "Eddie"
eddie.broke = true
```

Notice all the repetition.  There's no harm in using this approach but there are surely better ways.

```go
var justin Character
justin = NewCharacter("Justin", true)

var bustin Character
bustin = NewCharacter("Bustin", false)

var stache Character
stache = NewCharacter("Stache", false)

var eddie Character
eddie = NewCharacter("Eddie", true)
```

Reading wise you can already see a big difference.  This get's even simplier once we dig deeper into variable declaration and construction.

```go
var justin Character = NewCharacter("Justin", true)
var bustin Character = NewCharacter("Bustin", false)
var stache Character = NewCharacter("Stache", false)
var eddie Character = NewCharacter("Eddie", true)

PrintBroke(justin)
PrintBroke(bustin)
PrintBroke(stache)
PrintBroke(eddie)
```

If you're up to the task, try to get this code to compile.  A key hint is to remove the "&" and "*" values where you see them.

## Next lesson
In the next lesson, we'll proceed to rip out all "&" and "*" keywords from the code.  After that, we'll begin setting pieces in stone for the base design of characters.  If needed, we'll branch into functions even further.  If not, we'll continue on with the basics for reinforcement purposes.