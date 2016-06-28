package main

import (
	"fmt"
	"illmenu"
)

func images() {
	links, e := illmenu.ImageSearch("dogs")
	if e != nil {
		panic(e.Error())
	}

	fmt.Println(links)
}

func server() {
	illmenu.Serve()
}

func main() {
	server()
}
