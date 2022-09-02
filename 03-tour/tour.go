package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
    fmt.Print("\n")
    fmt.Println(" ====================")
    fmt.Println("| Hello, Tour of Go! |")
    fmt.Println(" ====================")

    fmt.Println("\nPlaying with math, time, rand...")
  	fmt.Println("> The time is", time.Now())
    rand.Seed(int64(time.Now().Nanosecond()))
    fmt.Println("> My favorite number is", rand.Intn(10))
   	fmt.Printf("> Now you have %g problems.\n", math.Sqrt(7))
    fmt.Println("> The value of Pi is", math.Pi)
    fmt.Println("> The sum of 3 with 7 is", add(3, 7))
    fmt.Println("> The factorial of 5 is", factorial(5))

    fmt.Println("\nPlaying with swaps and splits...")
    var a string = "hello"
    var b string = "world"
    fmt.Println("> Before swap: ", a, b)
    var c, d = swap("hello", "world")
    fmt.Println("> After swap: ", c, d)
    fmt.Print("> Splitting 17 in parts: ")
    fmt.Println(split(17))

    fmt.Println("\nShowing off types...")
    showOffTypes()

    fmt.Println("\nShowing off zero values for unitialized variables...")
    showOffZeroValues()

    fmt.Println("\nShowing off type conversions...")
    showOffTypeConversion()

    fmt.Println("\nShowing off type inferences...")
    showOffTypeInference()

    fmt.Println("\nShowing off constants usage...")
    showOffConstants()

    fmt.Print("\n")
  }

func add(x, y int) int {
  k := 1 // Assignment with implicit type.
	return x + y + (k - k)
}

func factorial(n int) int {
   if (n == 1) { return 1 }
   return n * factorial(n-1)
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return (avoid usage)
}

func showOffTypes() {
  var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
  )
  fmt.Printf("> Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("> Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("> Type: %T Value: %v\n", z, z)
}

func showOffZeroValues() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("> Zero value for int: %v\n", i)
  fmt.Printf("> Zero value for float64: %v\n", f)
  fmt.Printf("> Zero value for bool: %v\n", b)
  fmt.Printf("> Zero value for string: %q\n", s)
}

func showOffTypeConversion() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Printf("> Value for X: %v and Y: %v\n", x, y)
  fmt.Printf("> Value for F: %v\n", f)
  fmt.Printf("> Value for Z: %v\n", z)
}

func showOffTypeInference() {
  a := 42 // int
  fmt.Printf("> Value A = %v is of type: %T\n", a, a)
  b := 42.1 // float64
  fmt.Printf("> Value B = %v is of type: %T\n", b, b)
  c := "XPTO" // string
  fmt.Printf("> Value C = %v is of type: %T\n", c, c)
}

const Pi = 3.14
const (
  // Create a huge number by shifting a 1 bit left 100 places.
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
  // Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
  return x * 0.1
}
func showOffConstants() {
  const HelloWorld = "Hello World"
  fmt.Printf("> I say %q!\n", HelloWorld)
  fmt.Printf("> And Pi is around: %v.\n", Pi)

  const Truth = false
  fmt.Printf("> This sentence is %v.\n", Truth)

  // Numeric constants
  fmt.Printf("> needInt(Small): %v.\n", needInt(Small))
  // fmt.Printf("> needInt(Big): %v.\n", needInt(Big)) // this gives an error, number is too big for an int type
  fmt.Printf("> needFloat(Small): %v.\n", needFloat(Small))
  fmt.Printf("> needFloat(Big): %v.\n", needFloat(Big))
}