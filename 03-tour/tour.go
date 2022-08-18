package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Hello, Tour!")

  	fmt.Println("The time is", time.Now())

    rand.Seed(int64(time.Now().Nanosecond()))
    fmt.Println("My favorite number is", rand.Intn(10))
    
   	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    fmt.Println("The value of Pi is", math.Pi)

    fmt.Println("The sum of 3 with 7 is", add(3, 7))
    fmt.Println("The factorial of 5 is", factorial(5))
}

func add(x, y int) int {
	return x + y
}

func factorial(n int) int {
   if (n == 1) { return 1 }
   return n * factorial(n-1)
}