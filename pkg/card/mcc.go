package card

func TranslateMCC(code string) string {

	const notDefined = "Значение не определено"
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
		"8765": "Крокодильни",
	}

	if value, ok := mcc[code]; ok {
		return value
	} else {
		return notDefined
	}
}
