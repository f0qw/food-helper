package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Name     string
	ImageURL string
	Elements []Element `gorm:"foreignkey:DishID"`
}

type Element struct {
	gorm.Model
	DishID uint
	Name   string
	Price  string
	Num    string
}
