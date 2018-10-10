
package main

import "fmt"

func main() {

  words := [] string { "foo", "bar", "cake" }

  // synchronization method
  done := make( chan bool )

  // buffered channel with specified length
  ch := make( chan string, len( words ) )

  go func() {
    for {
      word, more := <- ch
      if ( more ) {
        fmt.Printf( "received word: %s\n", word )
        fmt.Println( "still more..." )
      } else {
        done <- true
        fmt.Println( "done" )
        return
      }
    }
  }()

  for _, word := range words {
    ch <- word
  }
  close( ch )

  // will wait until a message is received
  <- done



}