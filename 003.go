package main
import "fmt"
import "math"

func isPrime(n int) bool {
    if (n <= 1) {
        return false
    }
    root := int(math.Sqrt(float64(n)))
    for i:=2; i<=root; i++ {
        if n % i == 0 {
           return false
        }
    }
    return true
}

func main() {
    prime := 600851475143
    root := int(math.Sqrt(float64(prime)))
    res := 1
    for i:=2; i<=root; i++ {
        if prime % i == 0 && isPrime(i) {
            res = i
        }
    }
    fmt.Println(res)
}
