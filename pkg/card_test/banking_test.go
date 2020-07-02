package card_test

import (
	"fmt"
	"github.com/lozovoya/gohomework3_1/pkg/card"
	"time"
)

//transactionsCheck := []card.Transactions{
//	{4, 6_000_00, "Expense", "5411", "Finished", time.Now().Unix()},
//	{5, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()},
//	{6, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()},
//	}

//mccList := []string{"5812", "5411"}
func ExampleSumByMCC() {
	fmt.Println(card.SumByMCC(
	[{
		{
			4, 6_000_00, "Expense", "5411", "Finished", time.Now().Unix()
		},
		{
			5, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()
		},
		{
			6, 16_000_00, "Expense", "5812", "Finished", time.Now().Unix()
		},
	}], ["5812", "5411"] ))

	// Output:
	// 32000000
}