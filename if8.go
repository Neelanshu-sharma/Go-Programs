//Write a Go program to calculate profit or loss.
package main
import "fmt"
func main(){
    var cp int
    var sp int
    var profit int
    var loss int
    fmt.Println("Please enter Cost Price")
    fmt.Scanln(&cp)
    fmt.Println("Please eneter Sales Price")
    fmt.Scanln(&sp)
    if(cp < sp){
        profit = sp - cp
        fmt.Println("Profit is: ", profit)
    }
    if(sp < cp){
        loss = cp - sp
        fmt.Println("Loss is: ", loss)
    }
}
