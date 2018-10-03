
package main

import "fmt"

// "class"
type Position struct {
  x float64
  y float64
}

// receiver (class method)
func ( p Position ) print() {
  fmt.Printf( "x: %f, y: %f\n", p.x, p.y )
}

// use pointer if mutating
func ( p *Position ) shift( xOffset float64, yOffset float64 ) {
  p.x += xOffset
  p.y += yOffset
}

func main() {

  var pos Position = Position { 3.7, 4.9 }
  pos.print()

  pos.shift( 1, 2 )
  pos.print()

}