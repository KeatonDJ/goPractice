package main

import (
	"fmt"
	"testPackage"
)

func main(){
	fmt.Println("Hello World!");
	fmt.Println(testPackage.AddAll(4, 5, 6));
}