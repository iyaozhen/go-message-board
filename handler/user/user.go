package userhandler

import (
	"encoding/json"
	"mydemo/handler"
	Muser "mydemo/model/user"
	"mydemo/service/user"
	"net/http"
)

var sessionStore = handler.SessionStart()
var SessionName = "user"

func SingUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userInfo Muser.User
	err := decoder.Decode(&userInfo)
	if err != nil {
		handler.BaseResponse(w, 400, "params error")
	}
	res := user.Signup(userInfo.UserName, userInfo.PassWord)
	if res == true {
		handler.BaseResponse(w, 200, "ok")
	} else {
		handler.BaseResponse(w, 500, "error")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userInfo Muser.User
	err := decoder.Decode(&userInfo)
	if err != nil {
		handler.BaseResponse(w, 400, "params error")
	}
	loginUser := user.Login(userInfo.UserName, userInfo.PassWord)
	if loginUser.Id > 0 {
		session, _ := sessionStore.Get(r, SessionName)
		// Set some session values.
		session.Values["user_id"] = loginUser.Id
		session.Values["user_name"] = loginUser.UserName
		// Save it before we write to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handler.BaseResponse(w, 200, "ok")
	} else {
		handler.BaseResponse(w, 500, "error")
	}
}
