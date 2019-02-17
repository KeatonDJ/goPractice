//Find min/max value in a slice

package main

import "fmt"

func main(){

	nums := []int{28, 64, 128, 20, 5, 0, 45, 97}
	
	min, max := nums[0], nums[0]
	
	for _, num := range nums {
		if (num < min){
			min = num
		}
		
		if (num > max) {
			max = num
		}
	
	}
	
	fmt.Printf("Minimal: %d, Maximum: %d", min, max);
}