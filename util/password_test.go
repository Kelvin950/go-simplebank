package util

import "testing"

func TestPass(t *testing.T) {

	password := RandomString(10)

	hashPassword, err := HashedPassword(password)
	if err != nil {
		t.Error("failed to hash")
	}

	 isValid :=  CheckPassword(hashPassword  , password) ;
 

	 if !isValid{
		t.Error("password not correct") 
	 }
}