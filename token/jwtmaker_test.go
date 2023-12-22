package token

import (
	
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kelvin950/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {

	maker, err := NewJwtMaker(util.RandomString(23))

	if err!=nil{
		
		t.Errorf("failed %s" , err)
	}

	username:= util.RandomString(23) 
	duration := time.Minute

	token , err:= maker.CreateToken(username , duration) 
		expiredAt :=  time.Now().Add(duration)
	if err!=nil{
		t.Errorf("failed %s" , err)
	}
     
	if token ==""{
		t.Error("token failed to create")
	}
 
	 Payload , err:=maker.VerifyToken(token) 

	 if err!=nil{
		t.Errorf("failed to verify %s" , err)
	 }

	 if Payload ==nil{
		t.Error("failed to verify")
	 }

       require.NotEmpty(t ,Payload)
	   require.Equal(t , username , Payload.Username)
      require.WithinDuration(t , expiredAt , Payload.ExpiredAt ,time.Second)

	  
}

func TestExpiredJwtToken(t *testing.T){

	maker , err:=NewJwtMaker(util.RandomString(32)) ;

	require.NoError(t , err) 

	token , err := maker.CreateToken(util.RandomString(21) , -time.Minute)
	require.NoError(t, err) 

	payload , err :=  maker.VerifyToken(token)

	require.Error(t , err) 
	require.EqualError(t , err ,jwt.ErrTokenExpired.Error())

	require.Nil(t,payload )
}



// func TestInvalidJwtAlgNone(t *testing.T){


// }