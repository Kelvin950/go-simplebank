package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/simplebank/db"
	"github.com/kelvin950/simplebank/util"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type CreateUserRequest  struct{
	Username    string  	`json:"username" binding:"required ,alphanum"`
	 
	 Password string		`json:"password" binding:"required, min=6"`
	 FullName    string  	`json:"fullname" binding:"required"`
	 Email  string 			`json:"email" binding:"required ,email"`
}

func (server *Server) createUser(ctx *gin.Context){
  var req CreateUserRequest
  
  if err := ctx.ShouldBindJSON(&req) ; err!=nil{
	 ctx.JSON(http.StatusBadRequest , errorResponse(err))
	 return 
  }

 
  hashPassword , err := util.HashedPassword(req.Password)
if  err!=nil{
	 ctx.JSON(http.StatusBadRequest , errorResponse(err))
	 return 
  }

  arg :=  db.CreateUserParams{

	 Username: req.Username, 
	 FullName: req.FullName,
	 Email: req.Email,
	 HashedPassword: hashPassword,
  }


   user , err:= server.store.CreateUser(ctx , arg) 
    
   if err!=nil{

	  if pqErr , ok:= err.(*pq.Error) ; ok {
		 ctx.JSON(http.StatusForbidden , errorResponse(pqErr))

	return 

	  }
	ctx.JSON(http.StatusInternalServerError , errorResponse(err))

	return 
   }  
   
    rsp := NewUserResponse(user)
   ctx.JSON(http.StatusOK , rsp)


}


type LoginUserRequest  struct{
	Username    string  	`json:"username" binding:"required ,alphanum"`
	 
	 Password string		`json:"password" binding:"required, min=6"`

}

type  UserResponse struct{

	Username  string `json:"username"`
	Fulname string `json:"fullname"`
	Email string `json:"email"`
	PasswordChangedAt time.Time  `json:"passwordChangedAt"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type  LoginUserResponse struct{
	 AccessToken string `json:"accessToken"` 
      User UserResponse `json:"user"`
}


func NewUserResponse(user db.User)UserResponse{

	return  UserResponse{

		Username: user.Username,
		Fulname: user.FullName,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:  user.CreatedAt ,
	}
}


func  (server *Server) loginUser(ctx *gin.Context){

	var req LoginUserRequest 
 
	if err:= ctx.ShouldBindJSON(req); err!=nil{
		ctx.JSON(http.StatusBadRequest , errorResponse(err))

		return 
	}

	user , err := server.store.GetUser(ctx , req.Username) 
	
	  if err!=nil{
		if err== sql.ErrNoRows{
		 ctx.JSON(http.StatusNotFound , errorResponse(err))
		 return
	}

	ctx.JSON(http.StatusInternalServerError , errorResponse(err))

	return
	  }


	
	  isValid :=util.CheckPassword(user.HashedPassword , req.Password)
 
	   if !isValid{
			ctx.JSON(http.StatusInternalServerError , errorResponse(errors.New("password is invalid")))

			return

	   }


	   accessToken , err:= server.tokenMaker.CreateToken(user.Username , server.config.AcessTokenDuration)

	   if err== sql.ErrNoRows{
		 ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		 return
	}



	rsp := LoginUserResponse{
		AccessToken: accessToken,
		User:    NewUserResponse(user),
	} 

	ctx.JSON(http.StatusOK , rsp)

}
