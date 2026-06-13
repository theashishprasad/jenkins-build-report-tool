package models

type Build struct {
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Result   string `json:"result"`
	Duration int    `json:"duration"`
}