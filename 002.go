package main
import "fmt"

func main() {
    n1, n2 := 1, 2 
    lim := 4000000
    sum := 0 + 2
    for true {
        n1, n2 = n2, n1+n2
        if n2 % 2 == 0 {
            sum += n2
        }
        if n2 > lim {
            break
        }
    }
    fmt.Println(sum)
}
