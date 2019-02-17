//Count 1-20. If the number is divisible by 3 print fizz. If num is divisible by 5 print buzz. If both, print fizz buzz.

package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("fizz")
		}

		if i%5 == 0 {
			fmt.Println("buzz")
		}

		fmt.Println(i)

	}

}
