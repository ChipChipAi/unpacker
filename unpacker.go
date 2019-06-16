package unpacker

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) string {
	sliceRunes := []rune(str)

	if len(sliceRunes) < 1 {
		log.Printf("string is have 0 len\n")
		return ""
	}

	isBackslash := false
	strBuilder := strings.Builder{}

	for i, symb := range sliceRunes {
		// 92 is /
		if symb == 92 && !isBackslash {
			isBackslash = true
			continue
		}
		if unicode.IsLetter(symb) || isBackslash {
			isBackslash = false

			countChar := countLetter(sliceRunes, i)
			for countChar > 0 {
				strBuilder.WriteString(string(symb))
				countChar--
			}

		}
	}
	return strBuilder.String()
}

func countLetter(s []rune, index int) int {
	builder := strings.Builder{}
	index++
	if index+1 > len(s) || !unicode.IsDigit(s[index]) {
		return 1
	}
	for {
		if index+1 > len(s) {
			break
		}
		nRune := s[index]
		if !unicode.IsDigit(nRune) {
			break
		}
		builder.WriteRune(nRune)
		index++
	}
	count, err := strconv.Atoi(builder.String())
	if err != nil {
		log.Print("func countLetter: ", err)
	}
	return count
}
