package main

import (
	"fmt"
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

// mylist = list(s)

//         # Iterate through the list with a step size of 2k
//         for i in range(0, len(mylist), k + k):
//             # Reverse the sublist of characters within the range of i to i + k
//             mylist[i:k + i] = reversed(mylist[i:k + i])

//         # Join the reversed list of characters back into a string and return it
//         return "".join(mylist)

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

// // type Vector = cmp.Ordered
// func main() {

// 	// fmt.Println(Hello(19.20, 10))

// 	reverseStr("hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl", 39)

// }
