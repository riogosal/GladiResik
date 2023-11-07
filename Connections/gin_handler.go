package connection

import (
	food "GladiResik/Food"

	"github.com/gin-gonic/gin"
)

func GetFuncs(router *gin.Engine, MDB *MongoDB, foods *food.Foods) {
	router.GET("/food/view", func(c *gin.Context) {
		foods.ViewAll(c)
	})
}

func PostFuncs(router *gin.Engine) {

}

func PatchFuncs(router *gin.Engine) {

}

func DeleteFuncs(router *gin.Engine) {

}

func Setup(router *gin.Engine, MDB *MongoDB, foods *food.Foods) {
	GetFuncs(router, MDB, foods)
	PostFuncs(router)
	PatchFuncs(router)
	DeleteFuncs(router)
}
