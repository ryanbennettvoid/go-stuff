
package main

import "fmt"

func doStuff( str string, ch chan string ) {
  ch <- str
}

func main() {

  ch := make( chan string )

  // pushing to channel must
  // occur inside of a function
  // called with `go`
  go doStuff( "foo", ch )
  go doStuff( "bar", ch )
  go doStuff( "cake", ch )

  // var str string
  for item := range ch {
    fmt.Printf( "%s\n", item )
  }

  close( ch )


}