package main
import (
	"fmt"
)

type Employee struct {
	ID int;
	NAME string;
	PAY float32;
}



func main(){
	var name string
	var emp2 Employee

	emp2.ID = 44
	emp2.NAME = "Foo"
	emp2.PAY = 60.2

	emp1:= Employee {
		ID: 45,
		NAME: "SUE",
		PAY: 55.50,
		}	

	var myPointer *int
	var num = 12
	myPointer = & num

	fmt.Println("Value of num:", num)
	fmt.Println("Value of num:", *myPointer)
	fmt.Println("Address of num:", &num)
	fmt.Println("Address of num:", myPointer)
	fmt.Println("Address of pointer myPointer:", &myPointer)
	fmt.Println("Hello world")
	fmt.Println("#ID",emp1.ID)
	fmt.Println("#ID",emp2.ID,"Name:", emp2.NAME, "Pay:",emp2.PAY)
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	if name == "Winner"{
		fmt.Println("Ok, winner")
	} else  {
		fmt.Println("You entered:",name)	
	} 

}