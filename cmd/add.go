package cmd

import (
	"fmt"
	"food-helper/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var (
	elements []string
	price    []string
	num      []string
)

func NewAddCmd(dishService *services.DishService) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "添加菜名、需要的食材、食材的分量、价格",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Usage: add <菜名> -elements <材料1，材料2，...> -price <价格1，价格2，...> -num <度量1,度量2，...>")
				return
			}
			dishName := args[0]

			if len(elements) != len(price) || len(elements) != len(num) {
				fmt.Println("食材、价格、和度量的数量必须相同，同时注意需要使用英文下的逗号")
				return
			}
			_, err := dishService.GetDishByName(dishName)
			if err == nil {
				fmt.Println("该菜品已经存在，请重新添加！")
				return
			}

			err = dishService.AddDish(dishName, elements, price, num)
			if err != nil {
				fmt.Println("添加菜名失败:%v", err)
				return
			}

			fmt.Printf("添加菜名：%s \n", dishName)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"食材", "度量", "价格"})

			for i, _ := range elements {
				table.Append([]string{elements[i], price[i], num[i]})
			}
			table.Render()

		},
	}
}
