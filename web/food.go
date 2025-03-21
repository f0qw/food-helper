package web

import (
	"fmt"
	"food-helper/models"
	"food-helper/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type DishHandler struct {
	dishService *services.DishService
}

func NewDishHandler(s *services.DishService) *DishHandler {
	return &DishHandler{
		dishService: s,
	}
}

func (h *DishHandler) RegisterRoutes(server *gin.Engine) {
	dg := server.Group("/dishes")
	dg.GET("", h.ListDishes)         // GET /dishes 获取菜品列表
	dg.POST("", h.AddDish)           // POST /dishes 添加新菜品
	dg.DELETE("/:id", h.DeleteDish)  // DELETE /dishes/:id 根据ID删除菜品
	dg.POST("/order", h.OrderDishes) // POST /dishes/order 下单菜品
	dg.PUT("/:id", h.UpdateDish)     // 新增编辑路由
}

// DishRequest 添加菜品请求结构体
type DishRequest struct {
	Name     string         `json:"name"`
	ImageURL string         `json:"image_url"` // 新增图片URL字段
	Elements []ElementInput `json:"elements"`
}

type ElementInput struct {
	Name  string `json:"name"`
	Num   string `json:"quantity"`
	Price string `json:"price"`
}

// DishResponse 统一响应格式
type DishResponse struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	ImageURL string            `json:"image_url"` // 新增图片URL字段
	Elements []ElementResponse `json:"elements"`
}

type ElementResponse struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

// ListDishes 获取菜品列表
func (h *DishHandler) ListDishes(c *gin.Context) {
	dishes, err := h.dishService.ListDishes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50001,
			"message": "获取菜品列表失败",
		})
		return
	}

	var response []DishResponse
	for _, dish := range dishes {
		var elements []ElementResponse
		for _, e := range dish.Elements {
			elements = append(elements, ElementResponse{
				Name:     e.Name,
				Quantity: e.Num,
				Price:    e.Price,
			})
		}

		response = append(response, DishResponse{
			ID:       dish.ID,
			Name:     dish.Name,
			ImageURL: dish.ImageURL,
			Elements: elements,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": response,
	})
}

// AddDish 添加菜品
func (h *DishHandler) AddDish(c *gin.Context) {
	var req DishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "请求参数错误",
		})
		return
	}

	// 转换请求参数到服务层需要的格式
	var elements, prices, nums []string
	for _, e := range req.Elements {
		elements = append(elements, e.Name)
		prices = append(prices, e.Price)
		nums = append(nums, e.Num)
	}

	err := h.dishService.AddDish(req.Name, req.ImageURL, elements, prices, nums)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50002,
			"message": "添加菜品失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "菜品添加成功",
	})
}

// DeleteDish 删除菜品
func (h *DishHandler) DeleteDish(c *gin.Context) {
	id := c.Param("id")
	err := h.dishService.DeleteDishByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50003,
			"message": "删除菜品失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": fmt.Sprintf("菜品ID %s 已删除", id),
	})
}

// 新增UpdateDish处理函数
// UpdateDish 更新菜品
// @Summary 更新菜品信息
// @Description 根据ID更新菜品详细信息
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param id path string true "菜品ID"
// @Param dish body DishRequest true "菜品更新信息"
// @Success 200 {object} gin.H "更新成功"
// @Failure 400 {object} gin.H "请求参数错误"
// @Failure 500 {object} gin.H "服务器内部错误"
func (h *DishHandler) UpdateDish(c *gin.Context) {
	id := c.Param("id")
	var req DishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "请求参数错误",
		})
		return
	}

	// 转换请求参数
	var elements, prices, nums []string
	for _, e := range req.Elements {
		elements = append(elements, e.Name)
		prices = append(prices, e.Price)
		nums = append(nums, e.Num)
	}

	err := h.dishService.UpdateDish(id, req.Name, req.ImageURL, elements, prices, nums)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50004,
			"message": "更新菜品失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "菜品更新成功",
	})
}

// OrderRequest 下单请求结构体
type OrderRequest struct {
	DishIDs []string `json:"dish_ids"` // 菜品ID列表
}

// OrderResponse 下单响应结构体
type OrderResponse struct {
	TotalPrice float64 `json:"total_price"` // 总价格
	Details    []struct {
		Ingredient string  `json:"ingredient"` // 食材名称
		Quantity   float64 `json:"quantity"`   // 总数量
		Price      float64 `json:"price"`      // 总价格
	} `json:"details"`
}

// OrderDishes 下单菜品
func (h *DishHandler) OrderDishes(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "请求参数错误",
		})
		return
	}

	if len(req.DishIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "请传入需要下单的菜品ID",
		})
		return
	}

	dishMap := make(map[string]*models.Dish)
	for _, id := range req.DishIDs {
		dish, err := h.dishService.GetDishByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    40401,
				"message": fmt.Sprintf("获取菜品ID %s 失败", id),
			})
			return
		}
		dishMap[dish.Name] = dish
	}

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

	var details []struct {
		Ingredient string  `json:"ingredient"`
		Quantity   float64 `json:"quantity"`
		Price      float64 `json:"price"`
	}
	totalPrice := 0.0
	for name, values := range elementMap {
		details = append(details, struct {
			Ingredient string  `json:"ingredient"`
			Quantity   float64 `json:"quantity"`
			Price      float64 `json:"price"`
		}{
			Ingredient: name,
			Quantity:   values["num"],
			Price:      values["price"],
		})
		totalPrice += values["price"]
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": OrderResponse{
			TotalPrice: totalPrice,
			Details:    details,
		},
	})
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
