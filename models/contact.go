package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/Go_Rest_Api/utils"
	"fmt"
)

type Contact struct{
	gorm.Model
	Name		string		`json:"name"`
	Phone		string		`json:"phone"`
	UserId		uint		`json:"userid"`
}

func (contact *Contact) Validate () (map[string]interface{},bool) {

	if contact.Name == "" {
		return u.Message(false,"Name can not be null"), false
	}
	if contact.Phone == "" {
		return u.Message(false,"Phone can not be null"), false
	}
	if contact.UserId < 0 {
		return u.Message(false,"User can not be recognized"), false
	}

	// Everything is ok
	return u.Message(true,"success"), true
}

func (contact *Contact) Create() (map[string]interface{}){

	if resp,ok:=contact.Validate(); !ok{
		return resp
	}
	
	GetDB().Create(contact)

	resp := u.Message(true,"success")
	resp["contact"] = contact
	return resp
}

func GetContacts(user uint )([]*Contact){
	var contacts []*Contact
	err := GetDB().Table("contacts").Where("user_id = ?",user).Find(&contacts).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}