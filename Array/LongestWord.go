
import (
	"fmt"
	"strings"
	"unicode"
)

func Longest(sen string) string {
	best := ""
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	result := strings.FieldsFunc(sen, f)
	
	for _, v := range(result) {
		if len(v) > len(best) {
			best = v
		}
	}
	
	return best
}