package main

import "fmt"

func main() {
	var greetingTheWorld = "THIS IS MY FIRST TIME USING GO!"

	fmt.Println("################################################################")
	fmt.Println(greetingTheWorld)

	var MalawiDiscricts [3]string

	// Elements are assigned using index
	MalawiDiscricts[0] = "Lilongwe"
	MalawiDiscricts[1] = "Blantrye"
	MalawiDiscricts[2] = "Zomba"

	fmt.Print(MalawiDiscricts[0])

}
