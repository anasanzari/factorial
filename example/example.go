package main

import(
	"fmt"
	"github.com/anasanzari/factorial"
)

func main(){
	var i int
	fmt.Print("Enter an integer: ");
    _, err := fmt.Scanf("%d", &i)
    if err != nil {
    	fmt.Print("parse error")
    }
	fmt.Println("Factorial : ",factorial.Factorial(i))  
}