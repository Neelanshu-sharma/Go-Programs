//Write a Go program to print all odd number between 1 to 100.
package main
import "fmt"
func main(){
    var i int
    for i=0; i<=100; i++{
        if(i%2 != 0){
            fmt.Println(i)
        }
    }
}