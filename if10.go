/*Write a C program to input basic salary of an employee and calculate its Gross salary according to following:
Basic Salary <= 10000 : HRA = 20%, DA = 80%
Basic Salary <= 20000 : HRA = 25%, DA = 90%
Basic Salary > 20000 : HRA = 30%, DA = 95%*/

package main
import "fmt"
func main(){
    var basicsal float64
    fmt.Println("Please enter basic salary of an Employee")
    fmt.Scanln(&basicsal)
    var grosssal1 float64
    var grosssal2 float64
    var grosssal3 float64
    grosssal1 = basicsal+(basicsal*20/100)+(basicsal*80/100)
    grosssal2 = basicsal+(basicsal*25/100)+(basicsal*90/100)
    grosssal3 = basicsal+(basicsal*30/100)+(basicsal*95/100)
    if(basicsal<=10000){
         fmt.Println("Gross salary for employees which basic salary<=10000 is: ", grosssal1)
    }
    if(basicsal<=20000 && basicsal>10000){
         fmt.Println("Gross salary for employees which basic salary<=20000 is: ", grosssal2)
    }
    if(basicsal>20000){
         fmt.Println("Gross salary for employees which basic salary>20000 is: ", grosssal3)
    }
}