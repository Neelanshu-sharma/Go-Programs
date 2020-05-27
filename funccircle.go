//Write a Go program to find diameter, circumference and area of circle using functions.
package main
import "fmt"
func main(){
    var r float64
    fmt.Println("Radius of a circle is: ")
    fmt.Scanln(&r)
    fmt.Println("Diameter of Circle is: ", cdia(r))
    fmt.Println("Curcumfrance of a circle is: ", ccur(r))
    fmt.Println("Area of a circle is: ", acircle(r))
}
func cdia(num1 float64)float64{
    var dia float64
    dia = 2 * num1
    return dia;
}
func ccur(num2 float64)float64{
    var curcum float64
    var pi = 3.14
    curcum = 2*pi*num2
    return curcum;
}
func acircle(num3 float64)float64{
    var area float64
    var pi = 3.14
    area = pi*num3*num3
    return area;
}
