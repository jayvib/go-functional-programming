package car

type Car struct {
	Vin string `json:"vin"`
	Make string `json:"make"`
	Model string `json:"model"`
	Options struct {
		Option1 string `json:"option_1"`
		Option2 string `json:"option_2"`
	} `json:"options"`
	Timestamp string `json:"timestamp"`
}
