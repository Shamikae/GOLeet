package main
import (
	"fmt"
)

var fruit = [][]string{{"red","apple"},{"yellow","banana"},{"green","grapes"},{"blue","berry"}}

func main(){

	fmt.Println(fruit[0][0])
	fmt.Println(fruit[0][1])

	for a,value := range fruit {
		fmt.Print("return ",a +1, "\n")
		for _,cell := range value {
			fmt.Println(value)
			fmt.Println(cell)
		}
		fmt.Println("")	
	}

	for i:=0;i<len(fruit);i++{
		fmt.Println("line", i)
		for j:=0;j<2;j++{
			fmt.Print(fruit[i][j]," ")
		}
		fmt.Println(" ")
	}

}