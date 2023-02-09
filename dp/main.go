package main

func Recone(a string, dp map[string]int) map[string]int {
	if len(a) > 0 {
		dp[a[len(a)-1:]]++
		a = a[:len(a)-1]
		return Recone(a, dp)
	}
	return dp
}
