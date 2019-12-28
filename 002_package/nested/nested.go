package main

import (
	"github.com/UryKA/002_package/nested/say"
	"github.com/UryKA/002_package/nested/hello"
	"fmt"
)

func main() {

	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())

}
