package services

import (
	"food-helper/dao"
	"food-helper/models"
)

type DishService struct {
	DishDAO *dao.DishDAO
}

func NewDishService(dishDAO *dao.DishDAO) *DishService {
	return &DishService{dishDAO}
}

func (s *DishService) AddDish(name string, elements, prices, nums []string) error {
	dish := models.Dish{Name: name}
	for i := range elements {
		element := models.Element{
			Name:  elements[i],
			Price: prices[i],
			Num:   nums[i],
		}
		dish.Elements = append(dish.Elements, element)
	}

	return s.DishDAO.CreateDish(&dish)
}

func (s *DishService) ListDishes() ([]models.Dish, error) {
	return s.DishDAO.GetAllDishes()
}

func (s *DishService) DeleteDish(name string) error {
	dish, err := s.DishDAO.GetDishByName(name)
	if err != nil {
		return err
	}
	return s.DishDAO.DeleteDish(dish)
}

func (s *DishService) OrderDish(name string) (*models.Dish, error) {
	return s.DishDAO.GetDishByName(name)
}
