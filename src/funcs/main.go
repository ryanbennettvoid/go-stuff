
package main

import "fmt"

func voidFunc() {
  fmt.Println( "just do something" )
}

func add( a int, b int ) ( int ) {
  return a + b
}

func mult( a int, b int ) ( int ) {
  return a * b
}

// Passing a function as an arg.
// Arg describes func signature.
func runFunc( fn func(int, int) int, a int, b int ) ( int ) {
  return fn( a, b )
}

func multiReturn() (int, int, int) {
  return 1, 2, 3
}

func main() {

  voidFunc()

  var x int = add( 3, 7 )

  if ( x != 10 ) {
    fmt.Println( "assertion for `add` function failed" )
    return
  }

  var lambdaFunc = func() int {
    return 3
  }

  var y = lambdaFunc()

  if ( y != 3 ) {
    fmt.Println( "assertion for `lambdaFunc` function failed" )
    return
  }

  var sum = runFunc( add, 5, 11 )

  if ( sum != 16 ) {
    fmt.Println( "assertion for `runFunc` function failed for `sum`" )
    return
  }

  var product = runFunc( mult, 10, 4 )

  if ( product != 40 ) {
    fmt.Println( "assertion for `runFunc` function failed for `product`" )
    return
  }

  var a, b, c = multiReturn()

  if ( !( a == 1 && b == 2 && c == 3 ) ) {
    fmt.Println( "assertion for `multiReturn` function failed" )
    return
  }

}