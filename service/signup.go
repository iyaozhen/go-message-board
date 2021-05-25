package service

import (
	"crypto/sha256"
	"fmt"
	"mydemo/config"
	"mydemo/model"
)

func Signup(userName string, password string) bool {
	conf := config.Config().Security
	salt := conf["salt"]
	passwordHash := sha256.Sum256([]byte(fmt.Sprintf("%s%s%s", userName, salt, password)))
	return model.Signup(userName, fmt.Sprintf("%x", passwordHash))
}
