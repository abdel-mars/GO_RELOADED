package tools

import (
	"fmt"
	"os"
	"regexp"
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
	_, err2 := f.WriteString(myString + "\n")
	if err2 != nil {
		fmt.Println("Error writing to file:", err2)
		os.Exit(1)
	}
}

func IsVowelOrH(char rune) bool {
	return strings.ContainsRune("aeiouh", unicode.ToLower(char))
}

func TransformAToAn(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		words := strings.Split(line, " ")
		for j := 0; j < len(words)-1; j++ {
			if strings.EqualFold(words[j], "a") {
				firstChar := rune(words[j+1][0])
				if IsVowelOrH(firstChar) {
					if words[j] == "a" {
						words[j] = "an"
					} else if words[j] == "A" {
						words[j] = "An"
					}
				}
			}
		}
		lines[i] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}

func HexToDecimal(s string) string {
	lines := strings.Split(s, "\n")
	for idx, line := range lines {
		words := strings.Split(line, " ")
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
		lines[idx] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}

func BinToDecimal(s string) string {
	lines := strings.Split(s, "\n")
	for idx, line := range lines {
		words := strings.Split(line, " ")
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
		lines[idx] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
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

func FormatWithRegex(final_res string) string {
	rgx := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	final_res = rgx.ReplaceAllString(final_res, "$1 ")

	rgx = regexp.MustCompile(`\.\s*\.\s*\.`)
	final_res = rgx.ReplaceAllString(final_res, "...")

	rgx = regexp.MustCompile(`\.\.\.(\S)`)
	final_res = rgx.ReplaceAllString(final_res, "... $1")

	rgx = regexp.MustCompile(`\s+([.,!?:;])`)
	final_res = rgx.ReplaceAllString(final_res, "$1")

	rgx = regexp.MustCompile(`'\s*(.*?)\s*'`)
	final_res = rgx.ReplaceAllString(final_res, "'$1'")

	lines := strings.Split(final_res, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}
	return strings.Join(lines, "\n")
}
