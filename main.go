package main

import (
	"fmt"
	"os"
	"strings"
	"tools/tools"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: program <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile, outputFile := os.Args[1], os.Args[2]
	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: Input/output files must have .txt extension")
		os.Exit(1)
	}

	str := tools.ReadFileString(inputFile)
	str = tools.HexToDecimal(str)
	str = tools.BinToDecimal(str)
	str = tools.TransformAToAn(str)

	startBrac := tools.IndexOfStartBrac(str)
	endBrac := tools.IndexOfEndBrac(str)

	if len(startBrac) != len(endBrac) {
		fmt.Println("Mismatched brackets in input")
		os.Exit(1)
	}

	for i := len(startBrac) - 1; i >= 0; i-- {
		start, end := startBrac[i], endBrac[i]
		if start+1 >= len(str) || end >= len(str) {
			fmt.Println("Bracket index out of range")
			continue
		}

		cmdStr := str[start+1 : end]
		command, num := tools.SeturnSubStrAndNum(cmdStr)

		if num == -1 {
			fmt.Printf("Invalid flag: (%s)\n", cmdStr)
			continue
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
			fmt.Printf("Unknown flag: (%s)\n", cmdStr)
		}

		str = strings.TrimSpace(preCommand) + postCommand
	}

	str = tools.FormatWithRegex(str)
	fmt.Println(str)
	tools.StringToWriteFile(outputFile, str)
}
