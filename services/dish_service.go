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

func (s *DishService) AddDish(name, imageURL string, elements, prices, nums []string) error {
	dish := models.Dish{
		Name:     name,
		ImageURL: imageURL,
	}
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

func (s *DishService) DeleteDishByID(id string) error {
	dish, err := s.DishDAO.GetDishByID(id)
	if err != nil {
		return err
	}
	return s.DishDAO.DeleteDish(dish)
}

func (s *DishService) OrderDish(name string) (*models.Dish, error) {
	return s.DishDAO.GetDishByName(name)
}

func (s *DishService) GetDishByName(name string) (*models.Dish, error) {
	return s.DishDAO.GetDishByName(name)
}

func (s *DishService) GetDishByID(id string) (*models.Dish, error) {
	return s.DishDAO.GetDishByID(id)
}

// UpdateDish 更新菜品信息
func (s *DishService) UpdateDish(id, name, imageURL string, elements, prices, nums []string) error {
	// 获取现有菜品
	dish, err := s.DishDAO.GetDishByID(id)
	if err != nil {
		return err
	}

	// 更新基本信息
	dish.Name = name
	dish.ImageURL = imageURL

	// 清空原有元素
	if err := s.DishDAO.DB.Model(dish).Association("Elements").Delete(dish.Elements); err != nil {
		return err
	}

	// 添加新元素
	for i := range elements {
		element := models.Element{
			Name:  elements[i],
			Price: prices[i],
			Num:   nums[i],
		}
		dish.Elements = append(dish.Elements, element)
	}

	return s.DishDAO.UpdateDish(dish)
}
