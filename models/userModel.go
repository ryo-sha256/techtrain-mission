package Models

// uppercase so that the other service codes can have
// an access to the struct (e.g. ../models/user.go)
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserResponse struct {
	Name string `json:"name"`
}

func (b *User) TableName() string {
	return "user"
}
