package Controllers

import (
	"net/http"

	Models "techtrain-mission/models"
	Utils "techtrain-mission/utils"

	"github.com/gin-gonic/gin"
)

// CreateUser: creates a user and returns the token
func CreateUser(c *gin.Context) {

	// initialize a variable using an User model
	var user Models.User

	// Bind checks the Method and Content-Type
	// to select a binding engine automatically,
	c.BindJSON(&user)

	user.Token = Utils.GenerateToken()
	err := Models.CreateUser(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var res Models.CreateUserResponse
		res.Token = user.Token
		c.JSON(http.StatusOK, res)
	}
}

// GetUser: gets a user data with Models' module.
//          expects c (=*gin.Context) to have token
func GetUser(c *gin.Context) {
	token := Utils.ExtractAuthToken(c)

	var user Models.User
	c.BindJSON(&user)

	err := Models.GetUserByToken(&user, token)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// bind the result with the gin.Context here
		var res Models.GetUserResponse
		res.Name = user.Name
		c.JSON(http.StatusOK, res)
	}
}

// UpdateUser: updates a user's name with Models' module
//             expects c (=*gin.Context) to have token
func UpdateUser(c *gin.Context) {
	token := Utils.ExtractAuthToken(c)

	var user Models.User
	c.BindJSON(&user)

	name := user.Name

	err := Models.UpdateUser(&user, name, token)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var res Models.UpdateUserResponse
		res.Name = user.Name
		c.JSON(http.StatusOK, res)
	}
}
