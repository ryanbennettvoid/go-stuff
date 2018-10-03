
package main

import "fmt"

// To build, simply run `go build`
// and a binary appears.

// Similar to GCC, the `-o` argument
// indicates the output path.

// `go build -o <output binary> <src dir>`

// The compiler is strict with bracket
// placement; can't do this:

// func main()
// {

// }

func main() {
  fmt.Println( "The obligatory Hello World thingy... In Golang!" )
}