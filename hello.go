package main

import (
	"fmt"
	"math"
)

//function for checking prime number
func isprime(b int) bool{
    for i:=2;i<=int(math.Sqrt(float64(b)));i++{
		if b%i==0{
		  return false
		}
	}
	return true
}




func main(){
	//hello world
	fmt.Println("hello world")
	fmt.Print("hello world")
	fmt.Printf("hello world")
	fmt.Println(isprime(24))
	//even and odd logic
	var a int =15
	var b int =8
	switch z:=a%2;z{
	case 0:
		fmt.Println(a,"is even")
		fmt.Println(b,"is odd")
	case 1:
		fmt.Println(a,"is odd")
		fmt.Println(b,"is even")
	
	}
	
	
}
