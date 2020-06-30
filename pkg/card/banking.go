package card

type Transactions struct {
	Id     int64
	Amount int64
	Type   string
	Mcc    string
	Status string
	TrTime int64
}

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Icon         string
	Number       string
	Transactions []Transactions
}

func AddTransaction(card *Card, transaction *Transactions) {
	card.Transactions = append(card.Transactions, *transaction)
	//fmt.Println(card, card.Transactions, *transaction)
	return
}

func SumByMCC(transactions []Transactions, mcc []string) int64 {
	var summ int64 = 0
	for i := range transactions {
		for j := range mcc {
			if transactions[i].Mcc == mcc[j] {
				summ += transactions[i].Amount
			}
		}
	}
	return summ
}
