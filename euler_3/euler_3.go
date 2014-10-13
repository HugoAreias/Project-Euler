package main

import (
    "fmt"
    "math"
)

func main() {
    var number, prime, largest uint64 = 600851475143, 0, 0
    result := number

    // TODO: tweak the function to create a factor tree to find the prime 
    // factorisation of the given number
    prime = nextPrime(0) // get first prime number
    for {
        if result % prime == 0 { // check if this is a prime factor
            result = result / prime;
            if isPrime(result) { // stop when we find the largest prime factor
                if result > largest {
                    largest = result
                }
                break
            }
            prime = nextPrime(0) // restart
        } else {
            prime = nextPrime(prime) // get the least prime factor
        }
    }
    fmt.Println(largest)
}

func nextPrime(prime uint64) uint64 {
    var next uint64 = 0
    for i := prime + 1 ; next == 0 ; i++ {
        if isPrime(i) {
            next = i
            return next
        }
    }

    return next
}

func isPrime(number uint64) bool {
    if number == 2 {
        return true
    }
    // all even numbers are divisible by 2
    if number < 2 || number !=2 && number % 2 == 0 {
        return false
    }

    // to verify if a number is prime we just need to go until its square root
    squareRoot := math.Sqrt(float64(number))
    squareRoot = math.Floor(squareRoot)
    for i := 2 ; uint64(i) <= uint64(squareRoot) ; i++ {
        if number % uint64(i) == 0 {
            return false
        }
    }

    return true
}