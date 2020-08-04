import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(s, "")
	processedString = strings.ToLower(processedString)

	lengthString := len(processedString)
	for i := 0; i < lengthString/2; i++ {
		if rune(processedString[i]) != rune(processedString[lengthString-1-i]) {
			return false
		}
	}
	return true
}