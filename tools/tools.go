package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func StringToWriteFile(filename, myString string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer f.Close()
	_, err2 := f.WriteString(myString)
	if err2 != nil {
		fmt.Println("Error writing to file:", err2)
		os.Exit(1)
	}
}

func FormatPunctuation2(input string) string {
	var output []rune
	var prevRune rune

	for _, r := range input {
		if r == '\'' || r == ' ' {
			if prevRune == ' ' && r == '\'' {
				output = output[:len(output)-1]
				output = append(output, r)
			} else if prevRune != '\'' || r != ' ' {
				output = append(output, r)
			}
		} else {
			output = append(output, r)
		}
		prevRune = r
	}
	return string(output)
}

func FormatPunctuation(input string) string {
	var output []rune
	var prevRune rune

	for index, r := range input {
		if strings.ContainsRune(".,!?;:", r) {
			if prevRune == ' ' {
				output = output[:len(output)-1]
			}
			output = append(output, r)
			if index != len(input)-1 && !unicode.IsPunct(rune(input[index+1])) && input[index+1] != ' ' {
				output = append(output, ' ')
			}
		} else {
			output = append(output, r)
		}
		prevRune = r
	}
	return string(output)
}

func IsVowelOrH(char rune) bool {
	return strings.ContainsRune("aeiouh", unicode.ToLower(char))
}

func TransformAToAn(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		if strings.EqualFold(words[i], "a") {
			firstChar := rune(words[i+1][0])
			if IsVowelOrH(firstChar) {
				if words[i] == "a" {
					words[i] = "an"
				} else if words[i] == "A" {
					words[i] = "An"
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func HexToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexStr := words[i-1]
			decimalNum, err := strconv.ParseInt(hexStr, 16, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")
}

func BinToDecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binStr := words[i-1]
			decimalNum, err := strconv.ParseInt(binStr, 2, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")
}

func RemoveBrac(s string) string {
	var res string
	inBrac := false

	for _, char := range s {
		if char == '(' {
			inBrac = true
		} else if char == ')' {
			inBrac = false
		} else if !inBrac {
			res += string(char)
		}
	}
	return res
}

func ReadFileString(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return string(data)
}

func IndexOfStartBrac(s string) []int {
	var indexes []int
	for i, char := range s {
		if char == '(' {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func IndexOfEndBrac(s string) []int {
	var indexes []int
	for i, char := range s {
		if char == ')' {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func HasComma(s string) bool {
	return strings.Contains(s, ",")
}

func SeturnSubStrAndNum(newStr string) (string, int) {
	newStr = strings.TrimSpace(newStr)
	if HasComma(newStr) {
		parts := strings.SplitN(newStr, ",", 2)
		if len(parts) != 2 {
			return newStr, -1
		}
		command := strings.TrimSpace(parts[0])
		numStr := strings.TrimSpace(parts[1])
		if numStr == "" {
			return command, -1
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return command, -1
		}
		return command, num
	}
	return newStr, 1
}

func Up(s string, num int) string {
	if num <= 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) < num {
		num = len(words)
	}
	for i := len(words) - num; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}

func Cap(s string, num int) string {
	if num <= 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) < num {
		num = len(words)
	}
	for i := len(words) - num; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, " ")
}

func Low(s string, num int) string {
	if num <= 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) < num {
		num = len(words)
	}
	for i := len(words) - num; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, " ")
}
