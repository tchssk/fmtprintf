package a

import "fmt"

func main() {
	fmt.Printf("foo") // want `\Afmt\.Printf call which have one argument can be replaced with fmt\.Print\z`
}
