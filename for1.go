// For loop in Go, first example
package main
import "fmt"
func main(){
    var b int = 15
    var a int
    fmt.Println("Value of a is: ")
    for a := 0; a < 10; a++ {
        fmt.Println(a)
    }
    fmt.Println("Value of a in second case is: ")
    for a<b{
        a++
        fmt.Println(a)
            }
}