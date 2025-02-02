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
func LCM(a int, b int) int {
	c := GCD(a,b)
	result := (a*b)/c
	return result 
}
func isPrime(a int) bool{
	for i := 1; i<int(a/2);i++ {
		if a % i == 0{
			return false
		}
	}
	return true
}
func palidrome(text string) bool{
	temp := []rune(text)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i],temp[j] = temp[j],temp[i]
	}
	if text == string(temp){
		return true
	}
	return false
}
func selection(array []int) []int {
	n := len(array)
	for i := 0; i < len(array);i++{
		index := i 
		for j := i + 1;j<n;j++{
			if array[j]<array[index]{
				index = j
			}
		}
		temp := array[index]
		array[index] = array[i]
		array[i] = temp
	}
	return array
}
func bubble(array []int) []int {
	n := len(array)
	for i := 0; i<n;i++{	
		for j := 1; j<n-i;j++{
			if array[j-1] > array[j] {
				temp := array[j-1]
				array[j-1] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}
func binarySearch(array []int, n int) {
	left := 0
	right := len(array)-1
	for left <= right {
		middle := int((left+right)/2)
		if array[middle] > n {
			right = middle - 1
		}else if array[middle] < n{
			left = middle + 1
		}else{
			fmt.Print("Number Found on ", middle,"\n")
			return
		}
	}
	fmt.Print("Number not Found\n")
	return
}

func main(){
	var option int
	var a int
	var b int
	var word string
	for {
		fmt.Print("-------------------------\nChose an algorithm \n1.Greatest common divisor\n2.Least common multiple\n3.Prime number\n4.Is a palindrome\n5.Bubble sort\n6.Selection sort\n7.Binary search\nChoose option: ")
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
			fmt.Print("Type numbers \n a =")
			fmt.Scan(&a)
			if isPrime(a) == true {
				fmt.Print("Number is prime\n")
			}else{
				fmt.Print("Number is not prime\n")
			}
		case 4:
			fmt.Print("Type a word: ")
			fmt.Scan(&word)
			if (palidrome(word)) {
				fmt.Print("Word is a palindrome")
			}else {
				fmt.Print("Word is not a palindrome")
			}
		case 5:
			fmt.Print("Create an array \n How long array should be = ")
			fmt.Scan(&a)
			array := []int{}
			b = 0
			for i := 1; i <= a;i++ {
				fmt.Print("\nAdd a number = ")
				fmt.Scan(&b)
				if b == int(b){
					array = append(array, b)
				}else{
					fmt.Print("Input must be a number");
					i--
				}
			}
			fmt.Print("Sorted array = ", bubble(array),"\n")
		case 6:
			fmt.Print("Create an array \n How long array should be = ")
			fmt.Scan(&a)
			array := []int{}
			b = 0
			for i := 1; i <= a;i++ {
				fmt.Print("\nAdd a number = ")
				fmt.Scan(&b)
				if b == int(b){
					array = append(array, b)
				}else{
					fmt.Print("Input must be a number");
					i--
				}
				
			}
			fmt.Print("Sorted array = ", selection(array),"\n")
		case 7:
			fmt.Print("Create an array \n How long array should be = ")
			fmt.Scan(&a)
			array := []int{}
			b = 0
			for i := 0; i <= a;i++ {
				fmt.Print("\nAdd a number = ")
				fmt.Scan(&b)
				if b == int(b){
					array = append(array, b)
				}else{
					fmt.Print("Input must be a number");
					i--
				}
				
			}
			fmt.Print("Number you want to search for: ")
			fmt.Scan(&b)
			binarySearch(array,b)
		default:
			fmt.Print("Option does not exists\n")
		}
	}
}