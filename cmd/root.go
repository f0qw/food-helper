package cmd

import (
	"food-helper/services"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "food-helper",
	Short: "Food helper tool",
}

func Execute(dishService *services.DishService) {
	addCmd := NewAddCmd(dishService)
	addCmd.Flags().StringSliceVarP(&elements, "elements", "e", []string{}, "食材名")
	addCmd.Flags().StringSliceVarP(&price, "price", "p", []string{}, "食材价格")
	addCmd.Flags().StringSliceVarP(&num, "num", "n", []string{}, "食材度量")
	listCmd := NewListCmd(dishService)
	deleteCmd := NewDeleteCmd(dishService)
	orderCmd := NewOrderCmd(dishService)
	//orderCmd.Flags().StringSliceVarP(&orderNums,"n","n",[]string{},"菜单编号")

	rootCmd.AddCommand(addCmd, listCmd, deleteCmd, orderCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
