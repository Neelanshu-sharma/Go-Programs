//Using constant in place of a float variable, Cont by default treat a variable as float64

package main
import "fmt"
func  main(){
    const pi = 3.14
    const r = 200.00
    const areaofc = pi*r*r
    fmt.Println("Value of pi is: ", pi)
    fmt.Println("Value of radius is: ", r)
    fmt.Println("Area of Circle is: ", areaofc)
}