package models

type Laptop struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`

	Weight int `json:"weight"`
	RAM    int `json:"ram"`
	Cores  int `json:"cores"`
}
