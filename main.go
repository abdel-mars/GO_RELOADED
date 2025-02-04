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

	content := tools.ReadFileString(inputFile)

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		line = tools.HexToDecimal(line)
		line = tools.BinToDecimal(line)
		line = tools.TransformAToAn(line)

		startBrac := tools.IndexOfStartBrac(line)
		endBrac := tools.IndexOfEndBrac(line)

		if len(startBrac) != len(endBrac) {
			fmt.Println("Mismatched brackets in line:", line)
			continue
		}

		for j := len(startBrac) - 1; j >= 0; j-- {
			start, end := startBrac[j], endBrac[j]
			if start+1 >= len(line) || end >= len(line) {
				fmt.Println("Bracket index out of range in line:", line)
				continue
			}

			cmdStr := line[start+1 : end]
			command, num := tools.SeturnSubStrAndNum(cmdStr)
			if num == -1 {
				fmt.Printf("Invalid flag: (%s)\n", cmdStr)
				continue
			}

			preCommand := line[:start]
			postCommand := line[end+1:]

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

			line = strings.TrimSpace(preCommand) + postCommand
		}

		line = tools.TransformAToAn(line)
		line = tools.FormatWithRegex(line)
		lines[i] = line
	}

	finalStr := strings.Join(lines, "\n")
	fmt.Println(finalStr)
	tools.StringToWriteFile(outputFile, finalStr)
}
