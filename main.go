package main

import (
	"fmt"
)
func GCD(a int, b int) int {
	var c int
	for {
		if a % b == 0 {
			return b
		}else{
			c = b
			b = a % b
			a = c
		}
	}
}
// func LCM(a int, b int){
// 	c := GCD(a,b)

// }
func main(){
	var option int
	var a int
	var b int
	for {
		fmt.Print("-------------------------\nChose an algorithm \n1.Greatest common divisor\n2.Least common multiple\n")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Print("Type numbers \n a =")
			fmt.Scan(&a)
			fmt.Print("b = ")
			fmt.Scan(&b)
			fmt.Print("GCD equals = ", GCD(a,b))
		case 2:
			fmt.Print("Type numbers \n a =")
			fmt.Scan(&a)
			fmt.Print("b = ")
			fmt.Scan(&b)
			//LCM(a,b)
		}
	}
}