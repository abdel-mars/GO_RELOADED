package main

import (
	"fmt"
	"os"
	"strings"
	"tools/tools"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: program <input_file> <output_file>")
		os.Exit(1)
	}

	if !strings.HasSuffix(os.Args[2], ".txt") {
		fmt.Println("Error: input filename should have a .txt")
		os.Exit(1)
	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Error: input filename should have a .txt")
		os.Exit(1)
	}

	str := tools.ReadFileString(os.Args[1])
	str = tools.HexToDecimal(str)
	str = tools.BinToDecimal(str)
	str = tools.TransformAToAn(str)

	startBrac := tools.IndexOfStartBrac(str)
	endBrac := tools.IndexOfEndBrac(str)

	if len(startBrac) != len(endBrac) {
		fmt.Println("Mismatched brackets in input")
		os.Exit(1)
	}

	// avoid position shifts
	for i := len(startBrac) - 1; i >= 0; i-- {
		start, end := startBrac[i], endBrac[i]
		if start+1 >= len(str) || end >= len(str) {
			fmt.Println("Bracket index out of range")
			os.Exit(1)
		}

		cmdStr := str[start+1 : end]
		command, num := tools.SeturnSubStrAndNum(cmdStr)

		if num == -1 {
			fmt.Print("Invalid flag \n")
			os.Exit(1)
		}

		preCommand := str[:start]
		postCommand := str[end+1:]

		switch command {
		case "up":
			preCommand = tools.Up(preCommand, num)
		case "cap":
			preCommand = tools.Cap(preCommand, num)
		case "low":
			preCommand = tools.Low(preCommand, num)
		default:
			fmt.Print("Unknown flag \n")
		}

		str = strings.TrimSpace(preCommand) + postCommand
	}

	str = tools.FormatPunctuation(str)
	str = tools.FormatPunctuation2(str)

	fmt.Println(str)
	tools.StringToWriteFile(os.Args[2], str)
}
