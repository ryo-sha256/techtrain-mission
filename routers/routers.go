package Routers

import (
	"techtrain-mission/controllers"

	"github.com/gin-gonic/gin"
)

func RouterBuilder() *gin.Engine {
	r := gin.Default()

	r.POST("/user/create", Controllers.CreateUser)
	r.GET("/user/get", Controllers.GetUser)
	r.PUT("/user/update", Controllers.UpdateUser)

	r.POST("/gacha/create", Controllers.CreateCharacter)
	r.GET("/gacha/draw", Controllers.DrawCharacters)

	r.GET("/character/list", Controllers.GetCharacterList)

	return r
}
