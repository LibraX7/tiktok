package models

type Relation struct {
	Id          int `json:"id" gorm:"primaryKey"`
	FollowerId  int `json:"follower_id"`
	UserId      int `json:"user_id"`
	Friend_flag int `json:"friend_flag"`
}

func (Relation) TableName() string {
	return "relation"
}
