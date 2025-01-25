package cmd

import (
	"fmt"
	"food-helper/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

func NewListCmd(dishService *services.DishService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "列出所有的菜名",
		Run: func(cmd *cobra.Command, args []string) {
			dishes, err := dishService.ListDishes()
			if err != nil {
				fmt.Println("列出菜名失败：%v", err)
				return
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"食材", "度量", "价格"})

			for _, dish := range dishes {
				fmt.Printf("\n菜名: %s\n", dish.Name)

				for _, e := range dish.Elements {

					table.Append([]string{e.Name, e.Num, e.Price})
					//fmt.Printf("  食材：%s,价格：“%s, 度量：%s", e.Name, e.Price, e.Num)
				}
				table.Render()
			}

		},
	}
}
