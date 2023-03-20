package pk

type Bank interface {
	bankAccount(a int, specific string) string
}

type OnlinePay struct {
	ammount int
}

func (o *OnlinePay) bankAccount(a int, specific string) string {
	o.ammount = a
	return "Successfully send"
}
func Phonepay(B Bank, amt int, specific string) string {
	return B.bankAccount(amt, specific)
}
