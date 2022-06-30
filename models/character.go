package Models

import (
	"fmt"

	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	cnfg "techtrain-mission/config"
)

// creates one chracter
func CreateCharacter(c *CreateCharacterRequest) (err error) {

	if err = cnfg.DB_CONN.Create(c).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// gets a single characters based on the given probability
func GetCharacter(c *CreateCharacterRequest, probability float32) (err error) {
	// the logic for 'gacha' are the followings:
	// 1, grab the list of all the characters in probability's ascending order
	// 2, return the first character whose probability is lower than the given probability
	// FIXME: this is still a very naive way of giving each character probability distribution
	//        since the query below has to assume there is always a lower bound "0.01"
	if err = cnfg.DB_CONN.Where("probability BETWEEN ? AND ?", 0, probability).Order("probability desc").Limit(1).Find(c).Error; err != nil {
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
	runtimeErr := gormbulk.BulkInsert(cnfg.DB_CONN, c, 1000)

	if runtimeErr != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
