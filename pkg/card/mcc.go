package card

func TranslateMCC(code string) string {

	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
		"8765": "Крокодильни",
	}
	description := ""

	for key, value := range mcc {
		if key == code {
			description = value
		}
	}
	if description == "" {
		description = "Категория не указана"
	}
	return description
}
