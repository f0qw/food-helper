package cmd

import (
	"fmt"
	"food-helper/services"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(dishService *services.DishService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a dish",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Usage: delete <菜名>")
				return
			}
			dishName := args[0]

			err := dishService.DeleteDish(dishName)
			if err != nil {
				fmt.Printf("删除菜名失败：%v\n", err)
				return
			}
			fmt.Printf("菜名 %s 已经删除\n", dishName)
		},
	}
}
