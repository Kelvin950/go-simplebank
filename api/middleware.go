package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/simplebank/token"
)


const (
	authorizationHeaderKey = "authorization"
	authorizationPayloadKey =  "authorizationPayloadKey"
)
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc{

	return   func(ctx *gin.Context) {
          
      authroizationHeader := ctx.GetHeader(authorizationHeaderKey) 

	  if len(authroizationHeader) ==0{
		err := errors.New("authorization header not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized , errorResponse(err))
	  }
          

	  fields := strings.Split(authroizationHeader, " ")
       
	   if len(fields)  <2 {
		 
		ctx.AbortWithStatusJSON(http.StatusUnauthorized , errorResponse(errors.New("unauthorized")))
	   }
	    
	   token := fields[1]
	 payload , err := tokenMaker.VerifyToken(token)
 
	   if err!=nil{
		 
		ctx.AbortWithStatusJSON(http.StatusUnauthorized , errorResponse(err))
	   }


    
	   ctx.Set(authorizationPayloadKey ,payload)
	    ctx.Next()

	}
}