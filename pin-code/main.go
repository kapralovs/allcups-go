package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		input     int
		count     = 6
		strCombo  = ""
		nextCombo = make([]string, 3)
		combo     = 0
	)
	fmt.Scanln(&input)
	strInput := strconv.Itoa(input)
	splited := strings.Split(strInput, "")
	for i := 1; i <= count; i++ {
		if i == 1 {
			fmt.Println(input)
			for idx := range splited {
				switch idx{
				case 0:
					continue
				case 1:
					nextCombo[idx]=splited[idx+1]
				case 2:
					nextCombo[idx]=splited[idx-1]
				}
			}
			continue
		}
		combo, _ := strconv.Atoi(strings.Join(nextCombo, ""))
		fmt.Println(combo)
		if i%2 == 0 {
			nextCombo[0]=nextCombo[]
			for idx := range splited {
				nextCombo[]
			}
		}
	}
}
