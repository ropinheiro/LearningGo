package main

import (
	"fmt"
	"math"
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