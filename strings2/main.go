package main

import (
	"fmt"
	"strconv"
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


func minimizedStringLength(s string) int {
	hash := make(map[rune]struct{})
	for _, v := range s {
		hash[v] = struct{}{}
	}
	return len(hash)

}


func minLength(s string) int {
	b := []byte(s)
	var a bool
	for i := 0; i < len(b)-1; i++ {
		if b[i] == 'A' && b[i+1] == 'B' || b[i] == 'C' && b[i+1] == 'D' {
			a = true
			b = append(b[:i], b[i+2:]...)
		}
	}
	if a == true {
		return minLength(string(b))
	} else {
		return len(b)
	}
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

func reverseVowels(s string) {
	hashmap := map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	var buf string
	for _, v := range s {
		_, ok := hashmap[v]
		if ok {
			buf = string(v) + buf
		}
	}
	i1 := 0
	var a string
	for _, v := range s {
		_, ok := hashmap[v]
		if ok {
			a += string(buf[i1])
			i1++
		} else {
			a += string(v)
		}
	}
	fmt.Println(a)

	vowels := "aeiouAEIOU"
	vowelsIdx := make([]int, 0)

	// Find the indices of vowels in the string
	for i, v := range s {
		if strings.ContainsRune(vowels, v) {
			vowelsIdx = append(vowelsIdx, i)
		}
	}

	// Convert string to []rune for manipulation
	runeStr := []rune(s)
	vowelsCount := len(vowelsIdx)

	// Swap the vowels using the indices
	for i := 0; i < vowelsCount/2; i++ {
		runeStr[vowelsIdx[i]], runeStr[vowelsIdx[vowelsCount-i-1]] = runeStr[vowelsIdx[vowelsCount-i-1]], runeStr[vowelsIdx[i]]
	}

	// Print the result
	fmt.Println(string(runeStr))
}
func ReverseStr(s string, k int) string {
	s1 := []rune(s)
	i := 0
	for i < len(s1) {
		// Calculate the end index of the current segment
		end := i + k
		if end > len(s1) {
			end = len(s1)
		}
		// Reverse the characters in the current segment
		for i2, j := i, end-1; i2 < j; i2, j = i2+1, j-1 {
			s1[i2], s1[j] = s1[j], s1[i2]
		}
		// Move to the next 2k segment
		i += 2 * k
	}
	return string(s1)

}

func Add(a string) string {

	var right int
	var left int
	b := ""
	// "abb24ccc8ddbbca1"
	for right < len(a)-1 {
		if a[right] == a[right+1] {
			left++
			right++
		}
		if a[right] != a[right+1] {
			if !(a[right] >= '1' && a[right] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + (string(a[right]) + val)
			} else {
				b = b + string(a[right])
			}
			left = 0
			right++
		}
		if right == len(a)-1 {
			b = b + string(a[len(a)-1])
		}
	}
	// ’a1b224c3d2b2c1a11’
	return b
}

//	s := "abb24ccc8ddbbca1"
//
// fmt.Println(encode(s)) // output: "a1b224c3d2b2c1a11"
func Encode(a string) string {
	var right int
	var left int
	b := ""
	for right < len(a)-1 {
		if a[right] == a[right+1] {
			left++
		}
		if a[right] != a[right+1] {
			if !(a[right] >= '1' && a[right] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + (string(a[right]) + val)
			}
			left = 0
		}
		right++
		if right == len(a)-1 {
			if a[right-1] == a[len(a)-1] && !(a[right-1] >= '1' && a[right-1] <= '9') {
				val := strconv.Itoa(left + 1)
				b = b + string(a[right-1]) + val
			}
		}
		if a[right] >= '1' && a[right] <= '9' {
			b = b + string(a[right])
		}
	}
	return b
}

func IsAcronym(words []string, s string) bool {
	var a bool
	if len(words) == len(s) {
		for i, v := range words {
			if s[i] == v[0] {
				a = true
			} else {
				a = false
				break
			}
		}
	}
	return a
}

// func main() {

// 	fmt.Println(IsAcronym([]string{"never", "gonna", "give", "up", "on", "you"}, "ngguoy"))

// 	// fmt.Println("Hello, world!")
// }
