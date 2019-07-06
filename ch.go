// go语言运行shell命令
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Sex  bool
}

func main() {

	var ps []Person
	var count = len(ps)
	fmt.Println(count)
}
