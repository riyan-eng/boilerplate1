package util

import "golang.org/x/crypto/bcrypt"

func GenerateHash(str string) (value string) {
	hashedStr, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	value = string(hashedStr)
	return
}

func VerifyHash(hashedStr, candidateStr string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(candidateStr)); err == nil {
		return true
	} else {
		return false
	}
}
