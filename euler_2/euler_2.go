package main

import (
    "fmt"
    "math"
)

func main() {
    fib1, fib2, sum := 1, 2, 2

    for fib2 <= 4000000 {
        newfib := fib1 + fib2
        if math.Mod(float64(newfib), 2) == 0 {
            sum += newfib
        }

        fib1, fib2 = fib2, newfib
    }

    fmt.Println(sum)
}
