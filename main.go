package main

import (
	"fmt"
	"github.com/lozovoya/gohomework3_1/pkg/card"
	"time"
)

func main() {
	master := card.Card{
		Id:       1,
		Issuer:   "Master Card",
		Balance:  150_000_00,
		Currency: "rub",
		Icon:     "http://...",
		Number:   "4535 3454 4354 4353",
	}
	visa := card.Card{
		Id:       2,
		Issuer:   "Visa",
		Balance:  100_000_00,
		Currency: "rub",
		Icon:     "http://...",
		Number:   "2222 2222 2222 2222",
	}

	cards := make([]card.Card, 0, 10)

	transactions := []card.Transactions{
		{1, 735_55, "Expense", "5812", "Finished", time.Now().Unix()},
		{2, 2_000_00, "Income", "87658", "Processing", time.Now().Unix()},
		{3, 1_203_91, "Expense", "5411", "Finished", time.Now().Unix()},
		{4, 6_000_00, "Expense", "5411", "Finished", time.Now().Unix()},
		{5, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()},
		{6, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()},
	}

	cards = append(cards, visa)
	cards = append(cards, master)

	card.AddTransaction(&cards[0], &transactions[0])
	card.AddTransaction(&cards[0], &transactions[3])
	card.AddTransaction(&cards[0], &transactions[4])
	card.AddTransaction(&cards[1], &transactions[1])
	card.AddTransaction(&cards[1], &transactions[2])
	card.AddTransaction(&cards[1], &transactions[5])

	mccList := []string{"5812", "5411"}
	mccSum := card.SumByMCC(transactions, mccList)
	fmt.Println(mccSum)
}
