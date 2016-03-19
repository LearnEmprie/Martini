package main

import (

	"github.com/test/config"
	"fmt"
)

func main() {

	var result string
	result = config.GetCurrentDirectory()
	fmt.Println(result)
	
}
