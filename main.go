package main

import (
	"food-helper/cmd"
	"food-helper/config"
	"food-helper/dao"
	"food-helper/services"
)

func main() {
	// food-helper dish -name "菜名1" -elements "食材1 50g"
	// food-helper  -name "菜名1" -

	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	dishDAO := dao.NewDishDAO(db)
	dishService := services.NewDishService(dishDAO)
	cmd.Execute(dishService)

}
