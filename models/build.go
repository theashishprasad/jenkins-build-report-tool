package models

type Build struct {
	FullDisplayName string `json:"fullDisplayName"`
	Number          int    `json:"number"`
	Result          string `json:"result"`
	Duration        int    `json:"duration"`
}
