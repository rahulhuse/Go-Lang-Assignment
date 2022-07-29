package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mob     string `json:"mob"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
