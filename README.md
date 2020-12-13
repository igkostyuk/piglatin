# piglatin

Package for translating from English to Pig Latin.

## Example

   package main

    import (
	    "fmt"

	    "github.com/igkostyuk/piglatin"
    )

    func main() {
	    w := "Hello World!"
	    pw := piglatin.Translate(w)
	    fmt.Printf("%s is %s in pig latin.\n", w, pw)
    }
