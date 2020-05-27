// Write a Go program to find prime numbers up to n number

package main
import "fmt"
func main(){
    var n int
    var i int
    var j int
    fmt.Println("Please enter the num")
    fmt.Scan(&n)
    fmt.Println("Prime numbers up to n number ", n, " is ")
    for i=2; i<=n; i++{
        var c = 0
        for j= 1; j<=i; j++{
            if(i%j == 0){
                c = c+1
            }
        }
        if(c == 2){
            fmt.Printf("%d ", i)
        }
    }
}