package iteration

import "strings"

func Repeat (character string, repeatCount int) (repeated string) {
	repeated = strings.Repeat(character, repeatCount)
	return
}