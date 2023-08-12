package main

import (
	"strings"
)

func SplitWordsBySeparator(words []string, separator byte) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		res = append(res, strings.Split(words[i], string(separator))...)
	}
	var b []string
	for _, v := range res {
		if v == "" {
			continue
		}
		b = append(b, v)
	}
	return b
}

func Reverse(a string) string {
	b := []byte(a)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func FinalString(s string) (a string) {
	for _, v := range s {
		if v == 'i' {
			a = Reverse(a)
		} else {
			a += string(v)
		}
	}
	return a
}
