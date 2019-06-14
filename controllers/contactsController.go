package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Go_Rest_Api/models"
	u "github.com/Go_Rest_Api/utils"
	"net/http"
)

// func CreateContact(w http.ResponseWriter, r*http.Request){
// }

func CreateContacts(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		resp := u.Message(false, "Error Params.")
		u.Respond(w, resp)
		return
	}
	contact.UserId = user

	resp := contact.Create()
	u.Respond(w, resp)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	fmt.Println("id:", id)

	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["contacts"] = data
	u.Respond(w, resp)
}
