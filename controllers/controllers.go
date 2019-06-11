package controllers

import(
	"net/http"
	u "github.com/Go_Rest_Api/utils"
	"github.com/Go_Rest_Api/models"
	"encoding/json"
)

// 对应 url 的 handler
func CreateAccount(w http.ResponseWriter, r *http.Request){
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	
	u.Respond(w, resp)
}

func  Authenticate(){

}

func CreateContact(){

}
