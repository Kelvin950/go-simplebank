package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kelvin950/simplebank/db"
	"github.com/kelvin950/simplebank/token"
	"github.com/kelvin950/simplebank/util"
)

type Server struct {
	store db.Store 
	config util.Config
	tokenMaker token.Maker
	router *gin.Engine 

}

func NewServer( config util.Config ,store db.Store )( *Server , error){

	
	tokenMaker , err := token.NewPasetoMaker(config.TokenSymmetricKey)
 
	if err!=nil{

		return  nil , fmt.Errorf("cannot create token maker: %w ", err)
	}

	server:= &Server{store: store , tokenMaker: tokenMaker , config: config}
	

	if v ,ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)
	}

	
	
	server.setUpRouter() 


	return server , nil
}

func (server *Server) Start (add string)error{
	return server.router.Run(add)
}

func errorResponse(err error)gin.H{
	return gin.H{
		"error":err.Error() ,
	}
}


func( server *Server)setUpRouter(){
	router := gin.Default()
	router.POST("/user" , server.createUser)
    router.POST("/user/login" , server.loginUser)
	router.POST("/accounts" , server.createAccount)
	router.GET("/accounts/:id", server.getAccount) 
	authRoutes := router.Group("/").Use(AuthMiddleware(server.tokenMaker))
	authRoutes.GET("/accounts", server.ListAccount) 
	 authRoutes.POST("/transfer" , server.createTransfer)

	server.router = router
}