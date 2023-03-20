package main

import (
	"fmt"
)

type Tranver interface {
	Credited(a int) string
	Debited(a int) string
	BankBalance2() string
}

type Bank1 struct {
	name          string
	bankaccountNo string
	IFSCCODE      string
	BankBalance   int
}
type Bank2 struct {
	name          string
	bankaccountNo string
	IFSCCODE      string
	BankBalance   int
}

func (B *Bank1) Credited(a int) string {
	B.BankBalance = B.BankBalance + a
	if B.BankBalance < 0 {
		v := fmt.Sprintln(" Your Bank Balanse is Negative", B.BankBalance)
		return v
	} else {
		v := fmt.Sprintln("Your Bank Balanse is ", B.BankBalance)
		return v
	}
}
func (B *Bank1) Debited(a int) string {
	charge := 0.000001
	if B.BankBalance > 0 {
		B.BankBalance = B.BankBalance - a - int(charge)
		if B.BankBalance < a {
			v := fmt.Sprintln("Your Bank Balanse is Less the current balance ", B.BankBalance)
			return v
		} else {
			v := fmt.Sprintln("Your Bank Balanse is  ", B.BankBalance)
			return v
		}
	} else {
		v := fmt.Sprintln("Your Bank Balanse is Negative ", B.BankBalance)
		return v
	}
}

func (B *Bank1) BankBalance2() string {
	return fmt.Sprintln("your Bank Balance is ", B.BankBalance)
}

func (B *Bank2) Credited(a int) string {
	B.BankBalance = B.BankBalance + a
	if B.BankBalance < 0 {
		v := fmt.Sprintln("Your Bank Balanse is Negative", B.BankBalance)
		return v
	} else {
		v := fmt.Sprintln("Your Bank Balanse is ", B.BankBalance)
		return v
	}
}
func (B *Bank2) Debited(a int) string {
	charge := 0.0000002
	if B.BankBalance > 0 {
		B.BankBalance = B.BankBalance - a - int(charge)
		if B.BankBalance < a {
			v := fmt.Sprintln("Your Bank Balanse is Less the current balance ", B.BankBalance)
			return v
		} else {
			v := fmt.Sprintln("Your Bank Balanse is  ", B.BankBalance)
			return v
		}
	} else {
		v := fmt.Sprintln("Your Bank Balanse is Negative ", B.BankBalance)
		return v
	}
}

func (B *Bank2) BankBalance2() string {
	return fmt.Sprintln("your Bank Balance is ", B.BankBalance)
}

func Workingwithall(T Tranver, a int, ammount int) string {
	switch a {
	case 1:
		return T.Credited(ammount)
	case 2:
		return T.Debited(ammount)
	case 3:
		return T.BankBalance2()
	default:
		return "Tank you"
	}
}

func Phonepay(t Tranver, a int, ammont int) string {
	return Workingwithall(t, a, ammont)
}

func main() {

	// a := Bank1{}
	a := Bank2{}
	result := Workingwithall(&a, 1, 1)
	result = Workingwithall(&a, 1, 100)
	result = Workingwithall(&a, 2, 50)
	fmt.Println(result)

}

// package main

// import "fmt"

// type Cart interface {
// 	Add(produts string) []string
// }

// type Vegitabls struct {
// 	name []string
// }

// type Electronics struct {
// 	name []string
// }

// func (v Vegitabls) Add(produts string) []string {
// 	v.name = append(v.name, produts)
// 	return v.name
// }

// func (v *Electronics) Add(produts string) []string {
// 	v.name = append(v.name, produts)
// 	return v.name
// }

// func Pr(C []Cart) []Cart {
// 	var pr []Cart
// 	pr = append(pr, C...)
// 	return pr
// }

// func main() {
// 	a := []Cart{
// 		Vegitabls{[]string{"banana", "tomato", "hfhfb"}},
// 		&Electronics{[]string{"battery"}},
// 	}
// 	fmt.Println(Pr(a))

// 	//
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Funtion(a, b int, ch chan int, wg *sync.WaitGroup) {
// 	//defer close(ch)
// 	defer wg.Done()
// 	for i := a; i <= b; i++ {
// 		ch <- i
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {

// 	ch := make(chan int)
// 	ch2 := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go Funtion(1, 6, ch, &wg)
// 	go Funtion(6, 11, ch2, &wg)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// 	// _ = ch2
// 	// for v := range ch2 {
// 	// 	fmt.Println(v)
// 	// }
// 	wg.Wait()

// 	// var a any = 10
// 	// var s []int
// 	// // xt := reflect.TypeOf(a).Kind()
// 	// // if xt == reflect.String {
// 	// val, ok := a.(string)
// 	// if ok {
// 	// 	fmt.Println(val)
// 	// }
// 	// val2, ok2 := a.(int)
// 	// if ok2 {
// 	// 	s = append(s, val2)
// 	// }
// 	// fmt.Println(xt)
// 	// }
// 	// if xt == reflect.Int {
// 	// 	switch temp.data {
// 	// 	case temp.data.(int):
// 	// 		val, _ := temp.data.(int)
// 	// 		a = append(a, val)
// 	// 	}
// 	// }
// 	// if xt == reflect.Float64 {
// 	// 	switch temp.data {
// 	// 	case temp.data.(float64):
// 	// 		val, _ := temp.data.(float64)
// 	// 		b = append(b, val)
// 	// 	}
// 	// }

// 	// sort.Ints(a)
// 	// sort.Strings(c)
// 	// sort.Float64s(b)

// }
