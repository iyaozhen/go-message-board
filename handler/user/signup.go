package userhandler

import (
	"encoding/json"
	"mydemo/handler"
	"mydemo/service"
	"net/http"
)

func SingUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user handler.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	res := service.Signup(user.UserName, user.PassWord)
	if res == true {
		handler.BaseResponse(w, 200, "ok")
	} else {
		handler.BaseResponse(w, 500, "error")
	}
}
