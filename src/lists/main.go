
package main

import "fmt"

func printSlice( items [] string ) {
  fmt.Println()
  // for each
  for idx, val := range items {
    fmt.Printf( "%d: %s\n", idx, val )
  }
}

func main() {

  // slice, not array- has dynamic size
  items := [] string { "foo", "bar", "cake" }

  printSlice( items )

  // push
  items = append( items, "hey" )

  printSlice( items )

  firstTwo := items[:2]

  printSlice( firstTwo )

  lastTwo := items[2:]

  printSlice( lastTwo )


}