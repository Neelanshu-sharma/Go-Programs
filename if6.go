// Write a Go program to check whether a year is leap year or not.

package main
import "fmt"
func main(){
    var year int
    fmt.Println("Please enter the year: ")
    fmt.Scanln(&year)
    if((year%4 == 0 && year%100 != 0)||(year%400 == 0)){
        fmt.Println(year, "Year is leap year")
    }else{
        fmt.Println(year, "Year is a normal year and not a leap year")
    }
}
