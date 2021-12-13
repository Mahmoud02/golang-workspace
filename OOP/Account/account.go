package Account

import "fmt"

type Account struct{}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds from composition")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment from composition")
	return true
}

type CreditAccount struct {
	Account
}

func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds from credit account")
	return 250
}

type CheckingAccount struct {
	Account
}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds from checking account")
	return 125
}

/*
here we solve problem of delegation of  AvailableFunds method
*/

type HybridAccount struct {
	creditAccount   CreditAccount
	CheckingAccount CheckingAccount
}

func (h *HybridAccount) AvailableFunds() float32 {
	return h.creditAccount.AvailableFunds()
}
