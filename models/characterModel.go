package Models

type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}

type UserCharacter struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`

	User User `gorm:"foreignKey:UserCharacterID"`
}

type CreateCharacterRequest struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	CharacterID string  `json:"CharacterID" gorm:"unique"`
	Name        string  `json:"name" gorm:"unique"`
	Probability float32 `json:"probability" gorm:"unique"`
}
