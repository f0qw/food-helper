package cmd

import (
	"fmt"
	"food-helper/services"
	"github.com/spf13/cobra"
)

func NewOrderCmd(dishService *services.DishService) *cobra.Command {
	return &cobra.Command{
		Use:   "order",
		Short: "模拟点菜操作",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Usage: order <菜名>")
				return
			}
			dishName := args[0]

			dish, err := dishService.OrderDish(dishName)

			if err != nil {
				fmt.Printf("点菜失败：%v，\n", err)
				return
			}
			fmt.Printf("点菜成功： %v\n", dish.Name)
		},
	}
}
