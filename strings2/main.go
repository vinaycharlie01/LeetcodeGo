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
