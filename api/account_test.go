package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kelvin950/simplebank/db"
	mockdb "github.com/kelvin950/simplebank/db/mock"
	"github.com/kelvin950/simplebank/util"
	"go.uber.org/mock/gomock"
)

func TestGetAccountAPi(t *testing.T){
  
	account := createRandomAccount() 

	ctrl :=gomock.NewController(t) 

	defer  ctrl.Finish()

	store:= mockdb.NewMockStore(ctrl)
	

   store.EXPECT().GetAccount(gomock.Any() , gomock.Eq(account.ID)).Return(account, nil)

    
   server := newTestServer(t , store) 
   recorder := httptest.NewRecorder()

   url:= fmt.Sprintf("/account/%d", account.ID)
 request , err :=  http.NewRequest(http.MethodGet , url , nil) 
          
 if err!=nil{
	t.Error(err)
 }
    
 server.router.ServeHTTP(recorder , request)

  if recorder.Code != http.StatusOK{

	t.Errorf("failed  %d" , recorder.Code)
  }
   
}


func createRandomAccount()db.Account{

  return db.Account{
	Owner: util.RandomString(10),
	ID: util.RandomInt(0 ,100),
	Balance: util.RandomInt(0 ,200),
	Currency: "$",
  }

}