package main
import (
	"fmt"
	"strings"
)
//Syntax: strings.Replace()
//func Replace(s, old, new string, n int) string
// s = string, 
// old = string to be replaced 
// new string
// n = number of times to be replaced if n< 0 unlimeted times



var address = "255.100.50.0"

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func main(){
	fmt.Print(defangIPaddr(address),"\n")
}