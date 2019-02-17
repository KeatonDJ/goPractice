//function to test how many combinations of 4 digit numbers, when multiplied together, have the same first and last number.

package main

import "fmt"

func main() {
	
	count := 0
	
	for i := 1000; i < 10000; i++ {
		for j := i; j < 10000; j++ {
		
			num := i * j
			
			str := fmt.Sprintf("%d", num)
			
			if str[0] == str[len(str) - 1] {
			
				count++
			}
		
		}
	}
	
	fmt.Println(count)

}