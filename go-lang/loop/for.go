package main

import "fmt"

func main() {
	var names = [5]string{"Rafe Uddaraj", "Tasfiya", "Samiha", "Adiyan", "Nihan"}
	for index, value := range names{
		fmt.Println(index , value)
	}
}
