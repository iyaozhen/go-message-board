package user

import (
	"crypto/sha256"
	"fmt"
	"mydemo/config"
	"mydemo/model/user"
)

/**
password hash
*/
func hashPassword(userName string, password string) string {
	conf := config.Config().Security
	salt := conf["salt"]
	passwordHash := sha256.Sum256([]byte(fmt.Sprintf("%s%s%s", userName, salt, password)))
	return fmt.Sprintf("%x", passwordHash)
}

func Signup(userName string, password string) bool {
	return user.Signup(userName, hashPassword(userName, password))
}

func Login(userName string, password string) user.User {
	return user.Login(userName, hashPassword(userName, password))
}
