package model

func Signup(userName string, passWord string) bool {
	db := Conn()

	stmtIns, err := db.Prepare("INSERT INTO user VALUES(NULL, ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err2 := stmtIns.Exec(userName, passWord)
	if err2 != nil {
		return false
	}

	return true
}
