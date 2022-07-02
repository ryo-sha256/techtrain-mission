package Models

// TODO: not null
type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}

type UserCharacter struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type CreateCharacterRequest struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	CharacterID string `json:"CharacterID" gorm:"unique"`
	Name        string `json:"name" gorm:"unique"`
	Rarity      string `json:"rarity" gorm:"type:enum('common', 'uncommon', 'rare', 'ultraRare');default:'common'"`
}

type RarityList struct {
	Common    float32 `json:"common"`
	Uncommon  float32 `json:"uncommon"`
	Rare      float32 `json:"rare"`
	UltraRare float32 `json:"ultraRare"`
}
