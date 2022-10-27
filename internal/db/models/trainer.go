package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Name    string `json:"name"`
	Price   uint16 `json:"price"`
	HatesGO bool   `json:"islucid"`
}
