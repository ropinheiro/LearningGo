package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Hello, Tour!")

  	fmt.Println("The time is", time.Now())

    rand.Seed(int64(time.Now().Nanosecond()))
    fmt.Println("My favorite number is", rand.Intn(10))
}