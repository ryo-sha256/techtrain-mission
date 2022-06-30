package Controllers

import (
	"fmt"
	"net/http"

	Models "techtrain-mission/models"
	Utils "techtrain-mission/utils"

	"github.com/gin-gonic/gin"
)

// creates a character and write it to database table
// @param :none
func CreateCharacter(c *gin.Context) {
	var character Models.CreateCharacterRequest
	c.BindJSON(&character)

	err := Models.CreateCharacter(&character)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, character)
	}
}

// draws N random characters
// @param : N as times to draw a character from
// the pool of registered character list
func DrawCharacters(c *gin.Context) {
	token := Utils.ExtractAuthToken(c)

	var gachaRequest Models.GachaDrawRequest
	c.BindJSON(&gachaRequest)

	times := int(gachaRequest.Times)
	if times >= 101 {
		fmt.Println("ERROR: the max number of gacha/draw is 100")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// var characterList Models.CharacterListResponse
	// make goes here
	var characterList []interface{}
	var probablityList = Utils.RandGenerator(times)
	var res Models.GachaDrawResponse

	for i := 0; i < times; i++ {
		var character Models.CreateCharacterRequest

		// randomly retrieve a character from the database
		err := Models.GetCharacter(&character, probablityList[i])

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {

			// list to write to the table
			characterList = append(characterList, Models.UserCharacter{
				UserCharacterID: token,
				CharacterID:     character.CharacterID,
				Name:            character.Name,
			})

			// list to show the response
			res.Results = append(res.Results, Models.GachaResult{
				CharacterID: character.CharacterID,
				Name:        character.Name,
			})
		}
	}

	err := Models.CreateCharacterList(characterList)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}

}

// gets a character list by a given user auth token
func GetCharacterList(c *gin.Context) {
	token := Utils.ExtractAuthToken(c)

	var characterList []Models.UserCharacter

	err := Models.GetCharacterList(&characterList, token)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Models.CharacterListResponse{Characters: characterList})
	}
}
