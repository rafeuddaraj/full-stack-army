package main

import "fmt"

func main() {
	var names = [5]string{"Rafe Uddaraj", "Tasfiya", "Samiha", "Adiyan", "Nihan"}

	fmt.Println(len(names))
	var rafe = names[0]
	if names[0] == rafe {
		fmt.Println("Yes Founded")
	}
	fmt.Println((rafe))
}
