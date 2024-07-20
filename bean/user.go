package bean

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Rating   int    `json:"rating"`
	Ranking  int    `json:"ranking"`
}

func (User) TableName() string {
	return "user"
}
