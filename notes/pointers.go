package main
import(
	"fmt"
)

func main(){
	var numb1,numb2 int
	var ptr1,ptr2 *int

	numb1 = 20
	numb2 = 35

	fmt.Println("Before")
	fmt.Println("Numb1:", numb1)
	fmt.Println("Numb2:", numb2)
	fmt.Println("")

	ptr1 = &numb1
	ptr2 = &numb2
	fmt.Println(*ptr2)
	fmt.Println(ptr1)

	fmt.Println("This is the address of num1:", &numb1)
	fmt.Println("This is the address of num2:", &numb2)
	fmt.Println("This is a pointer to the value of num1:", *ptr1)
	fmt.Println("This is a pointer to the value of num2:", *ptr2)
	fmt.Println("This is a pointer to the address of num1:", ptr1)
	
	pointChange(&numb1)
	fmt.Println(numb1)
	swap(&numb1, &numb2)
	fmt.Println("Num1 After", numb1 )
	fmt.Println("Num2 After", numb2 )


}

func pointChange(p *int){
	*p += 1 
	fmt.Println("Point func",p, *p)
}
func swap(x,y *int){
	temp := *x
	*x = *y
	*y = temp
	fmt.Println("Swapped",*x,*y)

}