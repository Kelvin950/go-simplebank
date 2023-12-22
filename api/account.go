package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/simplebank/db"
	"github.com/lib/pq"
)

type CreateAccountRequest  struct{
	Owner    string  	`json:"owner" binding:"required"`
	 
	Currency string		`json:"currency" binding:"required,oneof=currency"`
}

func (server *Server) createAccount(ctx *gin.Context){
  var req CreateAccountRequest 
  
  if err := ctx.ShouldBindJSON(&req) ; err!=nil{
	 ctx.JSON(http.StatusBadRequest , errorResponse(err))
	 return 
  }

  arg :=  db.CreateAccountParams{

	Owner: req.Owner , 
	Currency: req.Currency, 
	Balance: 0,

  }


   account , err:= server.store.CreateAccount(ctx , arg) 
    
   if err!=nil{

	  if pqErr , ok:= err.(*pq.Error) ; ok {
		 ctx.JSON(http.StatusForbidden , errorResponse(pqErr))

	return 

	  }
	ctx.JSON(http.StatusInternalServerError , errorResponse(err))

	return 
   }  
   

   ctx.JSON(http.StatusOK , account)


}


type getAccountReq struct{
	Id int64   `uri:"id" bindng:"required ,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){

	var req getAccountReq 

	if err:= ctx.ShouldBindUri(&req) ; err!=nil{
		ctx.JSON(http.StatusBadGateway , errorResponse(err))
		return  
	}

 
	account , err := server.store.GetAccount(ctx , req.Id)
  
	if err!=nil {
        if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound , errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError , errorResponse(err)) 

		return 
	}
   
	 ctx.JSON(http.StatusOK , gin.H{

		"data":account , 
		
	 })
	

}


type listAccountsRequest struct{
	PageId int32   `form:"page_id" bindng:"required,min=1"`
	PageSize int32   `form:"page_id" bindng:"required,min=5 , max=10"`
}

func (server *Server) ListAccount(ctx *gin.Context){

	var req listAccountsRequest

	if err:= ctx.ShouldBindUri(&req) ; err!=nil{
		ctx.JSON(http.StatusBadGateway , errorResponse(err))
		return  
	}
 

	arg:= db.ListAccountsParams{
		Limit: int64(req.PageSize),
		Offset:  (int64(req.PageId)-1)*int64(req.PageSize)  ,
	}
 
	accounts , err := server.store.ListAccounts(ctx ,arg)
  
	if err!=nil {
        
		ctx.JSON(http.StatusInternalServerError , errorResponse(err)) 

		return 
	}
   
	 ctx.JSON(http.StatusOK , gin.H{

		"data":accounts , 
		
	 })
	

}