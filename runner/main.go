package main

import (
	"fmt"
	"illmenu"
)

func main() {
	links, e := illmenu.ImageSearch("dogs")
	if e != nil {
		panic(e.Error())
	}

	fmt.Println(links)
}
