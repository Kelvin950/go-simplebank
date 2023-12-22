package db

import (
	"context"
	"testing"

	"github.com/kelvin950/simplebank/util"
)

func TestCreateAccount(t *testing.T) {
  
	 
	var testValues = []CreateAccountParams{
            
    {Owner: createRandomUser(t).Username ,Currency: "$" , Balance: util.RandomInt(0,1000)} ,
	    {Owner: createRandomUser(t).Username ,Currency: "$" , Balance: util.RandomInt(0,1000)} ,
		    {Owner: createRandomUser(t).Username ,Currency: "$" , Balance: util.RandomInt(0,1000)} ,
			    {Owner: createRandomUser(t).Username ,Currency: "$" , Balance: util.RandomInt(0,1000)} ,
				    {Owner: createRandomUser(t).Username ,Currency: "$" , Balance: util.RandomInt(0,1000)} ,
	}


	for _ , v := range testValues { 

 
		 result , err :=  testQueries.CreateAccount(context.Background()  , v) 

		 if err!=nil {
			t.Errorf("%s" , err) 

		 }


		 if result.Owner !=  v.Owner {
				t.Errorf("%s" ,"failed wrong owner")
		 }

		 if result.Balance != v.Balance{
			t.Errorf("%s" , " failed wrong balance")
		 }

		 if result.ID == 0 {
			 t.Errorf("%s %s" , v.Owner , "not created" )
		 }

		 if result.Currency!= v.Currency{
			t.Errorf("%s" , "currency")
		 }

	}
 
}

func createRandomAccount(t *testing.T) Account {
	
	arg := CreateAccountParams{
		Owner:   util.RandomString(6),
		Balance:  util.RandomInt(0 , 1000),
		Currency: "$",
	}
 

   acc , _ := testQueries.CreateAccount(context.Background()  , arg)
   return acc
}