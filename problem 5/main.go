package main
import "fmt"

func main() {
    for i := 0; i < 1; i++ {
        n := 20
        answer := lcm(1, 2)
        for j := 2; j <= n - 1; j++ {
            answer = lcm(answer, j + 1)
        }
        fmt.Println(answer)
    }
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}