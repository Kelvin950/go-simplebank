package db

import (
	"context"

	"testing"


)

func TestTransferTx(t *testing.T){

  store:= NewStore(testDb) 
   

            
   account1 := createRandomAccount(t)
	   account2  := createRandomAccount(t) 
		     
	
	   n:=5 
	   amount := int64(10) 
	 
	    errs:= make(chan error) 

		results := make(chan  TransferTxResult)

	   for i:= 0 ;  i< n ; i++{
		go func(){
            
			 result , err:= store.TransferTx(context.Background() , TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId: account2.ID, 
				Amount: amount,
			 })

			 errs <-err 
			 results <- result
 		}()
	   }
	
 
	    for i := 0 ;  i<n ; i++{

			err:= <-errs 
			if err!=nil{
				t.Errorf("%s" , err)
			}
			result:=<-results 
 
			if &result ==nil{
				 
				t.Errorf("result is empty")
			}
            
		}
}