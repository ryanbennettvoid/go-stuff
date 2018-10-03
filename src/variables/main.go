
package main

import (
  "fmt"
)

func main() {

  const msg string = "hey"
  fmt.Println( msg )

  // different ways of assigning
  var x int = 3   // explicit type
  var y = 5       // implicit type
  z := 7          // shorthand

  result := x + y + z

  if ( result != 15 ) {
    fmt.Println( "bad assertion for x + y + z" )
    return
  }

}