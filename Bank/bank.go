package Bank

import "fmt"

type TransactionType string

const (
	DepositTx  TransactionType = "DEPOSIT"
	WithDrawTx TransactionType = "WITHDRAW"
)

type Transaction struct {
	Type   TransactionType
	Amount float64
	Note   string
}
type BankAccount struct {
	ID           int
	Owner        int
	Balance      float64
	Transactions []Transaction
}

func (b *BankAccount) Deposit(amount float64, note string) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	b.Balance += amount
	b.Transactions = append(b.Transactions, Transaction{
		Type:   DepositTx,
		Amount: amount,
		Note:   note,
	})
	return nil
}

func (b *BankAccount) WithDraw(amount float64, note string) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	b.Balance += amount
	b.Transactions = append(b.Transactions, Transaction{
		Type:   WithDrawTx,
		Amount: amount,
		Note:   note,
	})
	return nil
}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}
