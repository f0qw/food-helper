package cmd

import (
	"fmt"
	"food-helper/models"
	"food-helper/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strconv"
)

//var orderNums []string

func NewOrderCmd(dishService *services.DishService) *cobra.Command {
	return &cobra.Command{
		Use:   "order",
		Short: "模拟点菜操作",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("请传入需要下单的菜品ID编号，可以通过<food-helper list>获得")
				fmt.Println("Usage: order <菜单编号1,菜单编号2...>")
				return
			}

			dishMap := make(map[string]*models.Dish)

			for _, id := range args {
				dish, err := dishService.GetDishByID(id)
				if err != nil {
					fmt.Printf("获取菜品ID %s 失败： %v \n", id, err)
					continue
				}
				dishMap[dish.Name] = dish
			}

			// 每个食材的名字 以及该名字对应的价格和度量
			elementMap := make(map[string]map[string]float64)
			for _, dish := range dishMap {
				for _, element := range dish.Elements {
					if _, exists := elementMap[element.Name]; !exists {
						elementMap[element.Name] = map[string]float64{"price": 0, "num": 0}
					}
					elementMap[element.Name]["price"] += extractNumber(element.Price)
					elementMap[element.Name]["num"] += extractNumber(element.Num)
				}
			}

			table := tablewriter.NewWriter(os.Stdout)
			var recipeName string
			for _, n := range dishMap {
				recipeName = recipeName + n.Name + "  "
			}
			fmt.Println("总共点了下面的菜:", recipeName)
			table.SetHeader([]string{"菜名", "食材", "总价格", "总度量"})

			allPrice := 0.0
			for name, values := range elementMap {
				table.Append([]string{"", name, fmt.Sprintf("%.2f", values["num"]), fmt.Sprintf("%.2f", values["price"])})
				allPrice += values["price"]
			}
			table.Render()

			fmt.Printf("点菜成功,总共消费价格为：%v \n", allPrice)
		},
	}
}

func extractNumber(input string) float64 {
	re := regexp.MustCompile(`\d+(\.\d+)?`)
	match := re.FindString(input)
	if match == "" {
		return 0
	}
	value, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0
	}
	return value
}
