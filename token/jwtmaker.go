package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct{
	
	secretKey []byte
}

type MyCustomClaims struct {
	 Payload
	jwt.RegisteredClaims
}

func NewJwtMaker(secretKey string)(Maker , error){
 
	if len(secretKey) > minSecretKeySize {

	return nil , fmt.Errorf("invalid key size: must at be at least %d  characters" , minSecretKeySize)
	}
 
	return &JWTMaker{secretKey:[]byte(secretKey) } , nil 
}

func  (m JWTMaker)  CreateToken(username string ,duration time.Duration)(string ,error){
  payload , err :=  NewPayload(username , duration) 

  if err!=nil{
	return "" , err
  }

  claims := MyCustomClaims{
	 *payload,
       jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
	  IssuedAt: jwt.NewNumericDate(payload.IssuedAt),
	   },
  }

       token:=  jwt.NewWithClaims(jwt.SigningMethodHS256 , claims) 
           fmt.Println(m.secretKey)
	   ss , err := token.SignedString(m.secretKey)
           
     if err!=nil{
		return ""  , err
	 }
	    
	 return ss , nil 
}

func  (m JWTMaker) VerifyToken(token string)(  *Payload,error){
  
   	tok , err:=  jwt.ParseWithClaims(token , &MyCustomClaims{} ,func(t *jwt.Token) (interface{}, error) {

		return  m.secretKey , nil
	  })
   
	  
		if err!=nil{
			switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
	     return nil ,  fmt.Errorf("token malformed %s" , err)
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
	// Invalid signature
		return nil,	fmt.Errorf("token malformed %s" , err)
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
	// Token is either expired or not active yet
			return nil , errors.New(jwt.ErrTokenExpired.Error())
	default:
	return nil ,fmt.Errorf("couldn't handle this token: %s", err)
		}
		}
	  
     if claims, ok := tok.Claims.(*MyCustomClaims); ok {
	   
		return &claims.Payload , nil
	}else{
		 return nil , fmt.Errorf("could not decode token")
	}
}  