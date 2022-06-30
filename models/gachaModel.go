package Models

type GachaDrawRequest struct {
	Times int `json:"times"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}
