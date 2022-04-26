package model

type Costs map[string]int

type Node struct {
	Name  string `json:"name"`
	Costs Costs  `json:"costs"`
}
