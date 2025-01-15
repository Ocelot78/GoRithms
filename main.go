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
func LCM(a int, b int) int{
	c := GCD(a,b)
	result := (a*b)/c
	return result 
}
func main(){
	var option int
	var a int
	var b int
	for {
		fmt.Print("-------------------------\nChose an algorithm \n1.Greatest common divisor\n2.Least common multiple\n3.Bubble sort\nChoose option: ")
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
			fmt.Print("LCM equals = ",LCM(a,b))
		case 3:
			fmt.Print("Create an array \n How long array should be = ")
			fmt.Scan(&a)
			array := []int{}
			b = 0
			for i := 0; i <= a;i++{
				fmt.Print("\nAdd a number = ")
				fmt.Scan(&b)
				array = append(array, b)
			}
			fmt.Print(array)
		}
	}
}