package food

import (
	"github.com/gin-gonic/gin"
)

func GetFuncs(router *gin.Engine, foods *Foods) {
	router.GET("/food/view", func(c *gin.Context) {
		foods, err := foods.ViewAll(c)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		c.JSON(200, foods)
	})
}

func PostFuncs(router *gin.Engine) {

}

func PatchFuncs(router *gin.Engine) {

}

func DeleteFuncs(router *gin.Engine) {

}

func Setup(router *gin.Engine, foods *Foods) {
	GetFuncs(router, foods)
	PostFuncs(router)
	PatchFuncs(router)
	DeleteFuncs(router)
}
