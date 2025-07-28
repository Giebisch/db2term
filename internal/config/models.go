package config

type Config struct {
	Trips []Trip `json:"trips"`
}

type Trip struct {
	From     string `json:"from"`
	To       string `json:"to"`
	FromCode string `json:"fromCode"`
	ToCode   string `json:"toCode"`
}
