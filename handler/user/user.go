package userhandler

import (
	"encoding/json"
	"mydemo/handler"
	Muser "mydemo/model/user"
	"mydemo/service/user"
	"net/http"
)

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
		handler.BaseResponse(w, 200, "ok")
	} else {
		handler.BaseResponse(w, 500, "error")
	}
}
