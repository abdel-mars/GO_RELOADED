package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {

	if len(os.Args) == 3 {
		str := ReadFileString(os.Args[1])
		startArrBrac := indexOfStartBrac(str)
		endArrBrac := indexOfEndBrac(str)
		lenStartArrBrac := len(startArrBrac)
		

	}
	

}

func indexOfEndBrac(s string) []int {
	var res []int
	for i, x := range s {
		if x == ')' {
			res = append(res, i)
		}
	}
	return res
}

func ReadFileString(s string) string {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func indexOfStartBrac(s string) []int {
	var res []int 
	for i, x := range s {
		if x == '(' {
			res = append(res, i)
		}
	}
	return res
}

func Low(S string) string {
	words := string.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(low)") 

	}
}