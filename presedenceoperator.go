package main
import "fmt"
func main(){
   var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Println("Value of (a + b) * c / d is : ",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Println("Value of ((a + b) * c) / d is  : " ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Println("Value of (a + b) * (c / d) is  : ",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Println("Value of a + (b * c) / d is  : " ,  e );  
}