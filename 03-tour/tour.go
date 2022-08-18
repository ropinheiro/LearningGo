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
}