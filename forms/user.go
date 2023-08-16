package forms

type UserLoginForm struct {
	Password string `form:"password" json:"password" binding:"required,max=32"`
	Username string `json:"username" form:"username" binding:"required max=32"`
}
type UserRes struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	Avater          string `json:"avater"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowerCount   int    `json:"follower_count"`
	FollowCount     int    `json:"follow_count"`
	Password        string `json:"password"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}
