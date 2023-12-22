package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {

	hashpassword,  err  := bcrypt.GenerateFromPassword([]byte(password) , 10) ;

	if err!=nil {
		return "" , fmt.Errorf("failed to hash password: %s",  err) 
	}
	return  string(hashpassword) ,err 
}

func CheckPassword(hashpassword , password string)bool{

	err :=  bcrypt.CompareHashAndPassword([]byte(hashpassword) , []byte(password))

	 
	return err== nil 
}