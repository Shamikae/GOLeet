package main
import (
	"fmt"
)

var fruit = [][]string{{"red","apple"},{"yellow","banana"},{"green","grapes"},{"blue","berry"}}

func main(){

	fmt.Println(fruit[0][0])
	fmt.Println(fruit[0][1])

	for a,value := range fruit {
		fmt.Println("return",a +1)
		for _,cell := range value {
			fmt.Println(cell)
		}
		
	}

}