package Models

import (
	"fmt"

	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	cnfg "techtrain-mission/config"
)

// NOTE: This file contains all of the functions that
//       interact directly with database.
//       Additional service code/logic should be in controller files

// creates one chracter
func CreateCharacter(c *CreateCharacterRequest) (err error) {
	if err = cnfg.DB_CONN.Create(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// gets a list of probability distribution for character's rarity
func CreateRarityList(r *RarityList) (err error) {
	fmt.Println("invoked here")
	if err = cnfg.DB_CONN.Create(r).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// gets a list of probability distribution for character's rarity
func GetRarityList(r *RarityList) (err error) {
	fmt.Println("invoked here get rarity_list")
	if err = cnfg.DB_CONN.Find(r).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// gets a single characters based on the given probability
func GetCharacter(c *CreateCharacterRequest, rarity string) (err error) {

	if err = cnfg.DB_CONN.Raw("SELECT * FROM create_character_requests WHERE rarity = ?", rarity).Take(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// gets characters' list whose owner is a user of the given id
func GetCharacterList(c *[]UserCharacter, token string) (err error) {
	if err = cnfg.DB_CONN.Where("user_character_id = ?", token).Find(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// creates a character list with its owner's id
func CreateCharacterList(c []interface{}) (err error) {
	runtimeErr := gormbulk.BulkInsert(cnfg.DB_CONN, c, 100)

	if runtimeErr != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
