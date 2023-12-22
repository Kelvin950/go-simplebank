package token

import (
	
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)


type PasetoMaker  struct{

  paseto  *paseto.V2
	AsymmetricKey []byte
	
}

func NewPasetoMaker(asymmetricKey string)(Maker , error){
        if len(asymmetricKey) != chacha20poly1305.KeySize {

			return  nil , fmt.Errorf("invalid key size : %d", chacha20poly1305.KeySize)

			
		}
 

		maker := &PasetoMaker{
			AsymmetricKey: []byte(asymmetricKey),
			paseto:  paseto.NewV2(),
		}

	return maker ,   nil
}


func  (m PasetoMaker)  CreateToken(username  string , duration time.Duration)(string , error){
  

 
 
	  payload  , err:= NewPayload(username , duration)
	
	  if err!=nil{
		return  "" , fmt.Errorf("failed to create payload %s" , err) 
	  }

	  return m.paseto.Encrypt(m.AsymmetricKey , payload ,  nil) 
}

func (m PasetoMaker)  VerifyToken(token string)(*Payload , error){
   
	payload := &Payload{}

	err :=m.paseto.Decrypt(token  ,m.AsymmetricKey ,payload ,nil )

		if err!=nil{
		return nil , paseto.ErrInvalidTokenAuth
		}
	
		return payload , nil 
}