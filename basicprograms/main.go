package main

import (
	"fmt"
	"math"
	"strconv"
)

func Positiveornegative(a int) string {
	switch {
	case a == 0:
		return "ZERO"
	case a < 0:
		return "Negative"
	case a > 0:
		return "Positive"
	default:
		return "error"
	}
}

func EvenOrOdd(a int) string {
	switch {
	case a%2 == 0:
		return "Even"
	default:
		return "Odd"
	}
}

func NthNaturalNumber(a int) int {
	return (a * (a + 1) / 2)
}

func Sumofgivenrange(a int, b int) int {
	// var val int
	// for i := a; i <= b; i++ {
	// 	val += i
	// }
	// return val
	s := a * (a + 1) / 2
	s1 := b * (b + 1) / 2
	return s1 - s + a
}

func GretestoftwoNumer(a, b int) int {
	switch {
	case a > b:
		return a
	case b < a:
		return b
	default:
		return 0
	}
}

func GretestofThreeNumber(a int, b, c int) int {
	return int(math.Max(float64(a), math.Max(float64(b), float64(c))))
	// switch {
	// case a >= b && a >= c:
	// 	return a
	// case b >= c && b >= a:
	// 	return b
	// default:
	// 	return c
	// }
}

func LeapyearorNot(a int) string {
	if a%4 == 0 {
		if a%100 == 0 {
			if a%400 == 0 {
				return "LeapYear"
			} else {
				return "Not leap year"
			}
		} else {
			return "Leap year"
		}
	} else {
		return "Not leap year"
	}
}

func PrimeNuner(a int) string {
	var ok bool
	for i := 2; i < a; i++ {
		if a%i == 0 {
			ok = true
		}
	}
	fmt.Println(ok)
	return ""
}

func sieveOfEratosthenes(min int, max int) []int {
	// Initialize the sieve
	sieve := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		sieve[i] = true
	}

	// Mark non-prime numbers as false in the sieve
	// 20
	for i := 2; i*i <= max; i++ {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}
	// fmt.Println(sieve)
	// Collect the prime numbers from the sieve
	primes := []int{}
	for i := min; i <= max; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func PrimeNumerrange(min, max int) []int {
	sieve := make([]bool, max+1)

	for i := 2; i <= max; i++ {
		sieve[i] = true
	}

	for i := 2; i*i <= max; i++ {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}
	var primes []int
	for i := min; i < max; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes

}

func SumofDigit(a int, rever int) int {
	if a > 0 {
		rem := a % 10
		rever = rever*10 + rem
		return SumofDigit(a/10, rever)

	}
	return rever
	// var count int
	// temp := a
	// for temp > 0 {
	// 	rem := temp % 10
	// 	count = count*10 + rem
	// 	temp /= 10
	// }
	// return count
}

func Amstrong(a int) {
	i := strconv.Itoa(a)
	temp := a
	var count int
	for temp > 0 {
		rem := temp % 10
		count += int(math.Pow(float64(rem), float64(len(i))))
		temp /= 10
	}
	fmt.Println(count)
}

func Bott0m_Up_Fib(n int) []int {
	ans := make([]int, n+1)
	ans[0], ans[1] = 0, 1
	for i := 2; i <= n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	fmt.Println(ans)
	return ans
}

type Person struct {
	name   string
	age    int
	height float64
}

func Init(name string, age int, height float64) *Person {
	return &Person{name: name, age: age, height: height}
}

type MethodOver interface {
	add(a, b int) int
}

type number struct {
	a int
	b int
}

func (n *number) add(a, b int) int {
	n.a = a
	n.b = b
	return n.a + n.b
}

type Number2 struct {
	number
}

func (n *Number2) add(a, b int) int {
	n.a = a
	n.b = b
	return n.a + n.b
}

func main() {
	b := number{
		a: 10,
		b: 29,
	}
	s := Number2{
		b,
	}
	var c MethodOver = &s
	fmt.Println(c.add(10, 20))

	// fmt.Println(v.height)
	// Amstrong(a)
	// fmt.Println(a == SumofDigit(121, 0))
	// fmt.Println(PrimeNumerrange(1, 100))
	// fmt.Println(sieveOfEratosthenes(1, 100))

	// for i := 1; i < 100; i++ {
	// 	if big.NewInt(int64(i)).ProbablyPrime(i) {
	// 		fmt.Print(i, " ")
	// 	}
	// }

	// fmt.Println(Positiveornegative(10))

	// fmt.Println(EvenOrOdd(2))
	// fmt.Println(NthNaturalNumber(8))
	// fmt.Println(Sumofgivenrange(3, 6))
	// fmt.Println(GretestofThreeNumber(30, 40, 60))

	// fmt.Println(PrimeNuner(9))
	// fmt.Println(LeapyearorNot(2000))
}
