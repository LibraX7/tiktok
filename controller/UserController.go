package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"go_gin/forms"
	"go_gin/models"
	"go_gin/response"
	"go_gin/utils"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *gin.Context) {
	UserLoginForm := forms.UserLoginForm{}
	if err := c.ShouldBind(&UserLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//登录时将密码转化为hash加密后的密码
	Password, err := EncodePassword(UserLoginForm.Password)
	if err != nil {
		panic(err)
		return
	}
	UserLoginForm.Password = Password

	user, ok := dao.UserLogin(UserLoginForm.Username, UserLoginForm.Password)
	//用户未注册直接返回
	if !ok {
		response.Err(c, 401, 401, "该用户未注册", "")
		return
	}
	//为用户生成token
	token := utils.CreateToken(c, user.Id)
	//整理用户信息为统一返回封装的map格式（JSON）
	userinfoMap := map[string]interface{}{
		"userid": user.Id,
		"token":  token,
	}
	response.Success(c, 200, "success", userinfoMap)

}
func UserRegister(c *gin.Context) {
	UserLoginForm := forms.UserLoginForm{}
	if err := c.ShouldBind(&UserLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	Password, err := EncodePassword(UserLoginForm.Password)
	if err != nil {
		panic(err)
		return
	}
	UserLoginForm.Password = Password
	User := &models.User{}
	_, bool := dao.UserCreate(User)
	if !bool {
		response.Err(c, 500, 500, "插入失败", "")
		return
	}

	response.Success(c, 200, "创建成功", User)

}

// 功能函数，密码加密存储
func EncodePassword(password string) (string, error) {
	//加密方式：hash(password+randomnum)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return password, err
	}
	encodePWD := string(hash)
	return encodePWD, nil
}

// 功能函数:处理登录后的返回值
func HandleUserModelToMap(user *models.User) map[string]interface{} {

	userItemMap := map[string]interface{}{
		"id":              user.Id,
		"name":            user.Name,
		"avater":          user.Avater,
		"backgroundImage": user.BackgroundImage,
		"signature":       user.Signature,
		"followerCount":   user.FollowerCount,
		"followCount":     user.FollowCount,
		"password":        user.Password,
	}
	return userItemMap
}

type User struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	Name             string `json:"name"`
	Avater           string `json:"avater"`
	background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	FollowerCount    int    `json:"follower_count"`
	FollowCount      int    `json:"follow_count"`
	Password         string `json:"password"`
}
