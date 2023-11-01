package main

import "fmt"

type Person struct {
	name    string
	age     int
	country string
}

func main() {
	var rafe Person

	rafe.name = "Rafe Uddaraj"
	rafe.age = 19
	rafe.country = "Bangladesh"

	fmt.Println(rafe)

}
