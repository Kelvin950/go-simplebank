package db

import (
	"context"
	"testing"

	"github.com/kelvin950/simplebank/util"
)

func TestCreateUser(t *testing.T) {

	var testValues = []CreateUserParams{

		{Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6)},
		{Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6)},
		{Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6)},
		{Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6)},
		{Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6)},
	}

	for _, v := range testValues {

		result, err := testQueries.CreateUser(context.Background(), v)

		if err != nil {
			t.Errorf("%s", err)

		}

		if result.Username != v.Username {
			t.Errorf("%s", "failed wrong owner")
		}

		if result.Email != v.Email {
			t.Errorf("%s", " failed wrong balance")
		}

		if result.Username == "" {
			t.Errorf("%s %s", v.Username, "not created")
		}

		if result.HashedPassword != v.HashedPassword {
			t.Errorf("%s", "currency")
		}

	}

}


func TestGetUser(t *testing.T){

	user :=  createRandomUser(t) 


	user2 ,err :=  testQueries.GetUser(context.Background() , user.Username) ; 

 if err!=nil{
	t.Errorf("%s" ,  err )
 }


  if user.Username != user2.Username {
	t.Error("username incorrect")
  }
 
}
func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Username: util.RandomString(10), HashedPassword: "$", FullName: util.RandomString( 10) , Email: util.RandomString(6),
	}

	user, _ := testQueries.CreateUser(context.Background(), arg)
	return user
}