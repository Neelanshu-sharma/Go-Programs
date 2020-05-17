package main
import "fmt"
func main(){
    var a = 10
    var c int
    c = a 
    fmt.Println("Value of c = a, ", c)
    c += a
    fmt.Println("Value of c = c+a, ", c)
    c -= a
    fmt.Println("Value of c = c-a, ", c)
    c *= a
    fmt.Println("Value of c = c*a, ", c)
    c /= a
    fmt.Println("Value of c = c/a, ", c)
    c %= a
    fmt.Println("Value of c = c%a, ", c)    
}