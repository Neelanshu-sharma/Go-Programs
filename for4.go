//Program to display alphabets using ASCII values
package main
import "fmt"
func main(){
    var i int
    fmt.Println("alphabet characters from a to z are: ")
    for i = 97; i<=122; i++{
        fmt.Printf("%c \n", i) // Integer i with %c will convert integer to character before printing. %c will take ascii from i and display its character equivalent.
    }
    
}