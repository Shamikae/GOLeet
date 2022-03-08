package main
import (
	"fmt"
)



func main(){

	fruit:= []string{"apple","pear","banana","mango","cherry"}

	  //int, string
	for index, value :=range fruit {
		fmt.Printf( "%v\t%v\n",index, value) /* space= \t newline= \n value= %v, %#v type= %t, %T */
	}
	for _, value :=range fruit {
		fmt.Printf( "%T\n",value)
	}

	fmt.Println("Hello")
}