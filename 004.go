package main
import "fmt"
import "math"

func main() {
    maxlen := len(fmt.Sprintf("%v", 999*999))
    halflen := -(-maxlen / 2)
    fmt.Println(halflen)
    // loop through possible palindrome combinations from longest
    part := 0.0
    part2:= 0.0
    num := 0.0
    for true {
        for i:=halflen; i>0; i-- {
            for j:=9; j>=0; j-- {
                part = part*10 + float64(j)
                part2 = (part2+float64(j)) / 10
            }
        }
        fmt.Println(part, part2)
        num = (part+part2) * math.Pow(10, float64(halflen))
        fmt.Println(num)
        part = 0
        part2 = 0
        num = 0
        break
    }
}
