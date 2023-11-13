package food

import (
	entity "GladiResik/Entity"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	FoodData *Foods
}

func InitializeFoodHandler(router *gin.Engine, foods *Foods) {
	handler := &FoodHandler{
		FoodData: foods,
	}

	router.GET("/food/view", handler.GetFuncs)
	// router.GET("/food/view/:id", handler.GetOneFuncs)
}

func (h *FoodHandler) GetFuncs(c *gin.Context) {

	search := c.Query("search")

	foodData, err := h.FoodData.ViewAll(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}

	if search != "" {
		var temp []entity.Food
		for _, food := range foodData {
			if strings.Contains(strings.ToLower(food.Name), strings.ToLower(search)) {
				temp = append(temp, food)
			}
		}
		foodData = temp
	}

	c.JSON(200, foodData)
}

func GetOneFuncs(router *gin.Engine, foods *Foods) {
	router.GET("/food/view/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Bad Request",
			})
			return
		}

		foodData, err := foods.ViewAll(c)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}

		for _, food := range foodData {
			if food.ID == id {
				c.JSON(200, food)
				return
			}
		}

		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	})
}

func PostFuncs(router *gin.Engine) {

}

func PatchFuncs(router *gin.Engine) {

}

func DeleteFuncs(router *gin.Engine) {

}

func Setup(router *gin.Engine, foods *Foods) {
	// GetFuncs(router, foods)
	GetOneFuncs(router, foods)
	PostFuncs(router)
	PatchFuncs(router)
	DeleteFuncs(router)
}
