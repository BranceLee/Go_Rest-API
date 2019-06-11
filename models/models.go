package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"strings"
	u "github.com/Go_Rest_Api/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email 		string		`json: 	"email"`
	Password	string		`json: 	"password"`
	Token		string		`json:	"token"; sql:"-" `
}

const (
	ERROR_EMAIL = "Email address is inlegal"
	ERROR_PASSWORD_LENTH="Password is required"
	ERROR_EMAIL_USED="Email address has been used"
	ERROR_CONNECTION="Connection error. Please try again"
	GET_SUCCESS="Requirement passed"
)

// Validate incoming user details...
func (account Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, ERROR_EMAIL), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, ERROR_PASSWORD_LENTH), false
	}

	// Email must be unique
	temp := &Account{}
	
	// 数据类型合法时，查询数据表,GetDB return the configed db
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, ERROR_CONNECTION), false
	}
	// 存在该email
	if temp.Email != "" {
		return u.Message(false, ERROR_EMAIL_USED), false
	}

	return u.Message(false, GET_SUCCESS), true
}

// Type Account can create User 
func (account *Account) Create() (map[string] interface{}){

	if res, ok := account.Validate(); !ok {
		return res
	}

	// 密码加密，并替换新输入的原密码（MinCost int = 4, MaxCost int = 32, DefaultCost int = 10)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)


	// Create new JWT token from the new Registered account, 对account.ID 进行加密
	tk := &Token{UserId: account.ID}

	// The hashed tk matchs the account.ID
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),tk)
	// 字符串签名
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	// 生成返回数据给前端
	account.Password = ""
	response := u.Message(true, "Account has been Created")
	response["account"] = account
	return response
}

