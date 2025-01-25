package dao

import (
	"food-helper/models"
	"gorm.io/gorm"
)

type DishDAO struct {
	DB *gorm.DB
}

func NewDishDAO(db *gorm.DB) *DishDAO {
	return &DishDAO{DB: db}
}

func (dao *DishDAO) CreateDish(dish *models.Dish) error {
	return dao.DB.Create(&dish).Error
}

func (dao *DishDAO) GetAllDishes() ([]models.Dish, error) {
	var dishes []models.Dish
	err := dao.DB.Preload("Elements").Find(&dishes).Error
	return dishes, err
}

func (dao *DishDAO) GetDishByName(name string) (*models.Dish, error) {
	var dish models.Dish
	err := dao.DB.Where("name = ?", name).Preload("Elements").First(&dish).Error
	return &dish, err
}

func (dao *DishDAO) DeleteDish(dish *models.Dish) error {
	return dao.DB.Delete(dish).Error
}
