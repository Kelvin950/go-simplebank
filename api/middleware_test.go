package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/simplebank/token"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T , 
	request *http.Request ,
	tokenMaker  token.Maker , 
	username string , 
	duration time.Duration ,

){
 
	   token , err := tokenMaker.CreateToken(username , duration)


	    require.NoError(t , err) 
 
		authorizationHeader := fmt.Sprintf("Bearer %s", token)
		request.Header.Set(authorizationHeaderKey , authorizationHeader  )
}


func TestAuthMiddleWare(t *testing.T) {
 
	testCases := []struct{

		name string  
        setupAuth  func(t *testing.T , request *http.Request , tokenMaker token.Maker)
		checkRes  func( t *testing.T , recorder *httptest.ResponseRecorder)
	}{  
		  
		{name:"ok" ,setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			addAuthorization(t  , request , tokenMaker ,"user" , time.Minute  )
		}  , 
	
	checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {

		require.Equal(t , http.StatusOK , recorder.Code)
	}} ,

		  
		{name:"noauth" ,setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {

            
		}  , 
	
	checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {

		require.Equal(t , http.StatusUnauthorized , recorder.Code)
	}} ,

		  {name:"expired" ,setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			addAuthorization(t  , request , tokenMaker ,"user" , -time.Minute  )
		}  , 
	
	checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {

		require.Equal(t , http.StatusOK , recorder.Code)
	}} ,

		


	}

	for  _, tc :=range  testCases{

		t.Run(tc.name , func(t *testing.T) {
			server := newTestServer(t , nil) 
			authPath :="/auth"

			server.router.GET(
				authPath , 
				AuthMiddleware(server.tokenMaker) ,
				func(ctx *gin.Context) {

				} ,
			)

			recorder :=httptest.NewRecorder()
			request ,err:= http.NewRequest(http.MethodGet , authPath , nil)
			 
			require.NoError(t , err) 
			tc.setupAuth(t , request , server.tokenMaker)
			server.router.ServeHTTP(recorder ,request )
			tc.checkRes(t , recorder)
		})  
	}
}