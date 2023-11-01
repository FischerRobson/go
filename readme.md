### Run go 

Build and execute go code.

```sh
go run [file]
```
### Package

Its like a workspace.

Executable type: generates a file that we can run.
Main is used for executable package.
Go build generates main file that is runnable.

Reusable type: code used as 'helpers', good for reusable logic.

### Var declaration

```go
var card string = "Ace of Spades"
card1 := "Ace of Spades" // short init
card1 = "Five of Diamonds"
```

### Custom types

```go
type deck []string
```

### Array X Slice

Array in go have fixed lenght, while slice can grow or shrink.

Slices element must be of same type.

```go
cards := []string{"Ace of Diamonds", "Five of Spades"} // slice declaration

cards = append(cards, "Ace of Spades") // adding element to slice, append creates a new slice

for i, card := range cards { // iterate in cards
	fmt.Println(i, card)
}
```

Dividing slice:

`slice[startIndexIncluding:upToNotIncluded]` 

Both values can be omited, so go will infer start and end of slice.

```go

fruits = []string{"Apple", "Banana", "Grape", "Orange"}

fmt.Println(fruits[0:2])

```

### Receiver functions

Any variable of type deck now access print method.

Use the first letter for name reference, like `d`.

```go
func (d deck) print() {}
```

### Multiple returning value from functions

The `(deck, deck)` represent that this function return 2 variables of deck type.

```go
// method declare
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// retriving the values
hand, remainingDeck := deal(cards, 5)
```