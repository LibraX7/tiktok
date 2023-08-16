package models

type User struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	Avater          string `json:"avater"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowerCount   int    `json:"follower_count"`
	FollowCount     int    `json:"follow_count"`
	Password        string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
