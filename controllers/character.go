package Controllers

import (
	"fmt"
	"net/http"

	Models "techtrain-mission/models"
	Utils "techtrain-mission/utils"

	"github.com/gin-gonic/gin"
)

// creates a list of drop rate for each rarity, such as common, uncommon, and rare
// @param: float number between 0 and 1 for each rarity
func CreateRarityList(c *gin.Context) {
	fmt.Println("invoked in CreateRarityList")
	var rarityList Models.RarityList
	c.BindJSON(&rarityList)

	err := Models.CreateRarityList(&rarityList)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println("setting up the list of drop rates....")
		c.JSON(http.StatusOK, rarityList)
	}
}

func GetRarityList(c *gin.Context, r *Models.RarityList) {
	err := Models.GetRarityList(r)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

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

	var characterList []interface{}
	var probablityList = Utils.RandGenerator(times)
	var res Models.GachaDrawResponse
	var rarityList Models.RarityList

	err := Models.GetRarityList(&rarityList)
	if err != nil {
		fmt.Println("ERROR: error occured when retrieving rarity list")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for i := 0; i < times; i++ {
		var character Models.CreateCharacterRequest

		rarity := Utils.RarityTable(probablityList[i], rarityList)

		// randomly retrieve a character from the database
		err := Models.GetCharacter(&character, rarity)

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

	err = Models.CreateCharacterList(characterList)

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
