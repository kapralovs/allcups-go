package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Scanln(&input)
	strNum := strconv.Itoa(input)
	strCode := strNum[1:4]
	code, _ := strconv.Atoi(strCode)
	fmt.Print(code)
}
