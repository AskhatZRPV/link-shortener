package pwd

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	pwdBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}
