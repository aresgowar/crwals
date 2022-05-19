package entity

import (
)

type Questions struct {
	ID      int		`json:"id"`
	Author  string 	`json:"author"`
	Title   string 	`json:"title"`
	Link    string 	`json:"link"`
	Votes   string 	`json:"votes"`
	Answers string 	`json:"answers"`
	Views   string 	`json:"views"`
}
