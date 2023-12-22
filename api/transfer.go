package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/simplebank/db"
)

type transferRequest struct {
	FrromAccountID int64 `json:"fromaccountid" binding:"required , min=1"`
     ToAccountID  int64  `json:"toaccountid" binding:"required ,min=1"`  
	 Amount int64 `json:"amount" binding:"required , gt=1"` 
	Currency string `json:"currency" binding:"required,oneof=currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req  transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
 
	if !server.validAccount(ctx , req.FrromAccountID ,req.Currency){
		
		return
	}

	 
	if !server.validAccount(ctx , req.ToAccountID ,req.Currency){
		
		return
	}

	
	arg := db.TransferTxParams{

	FromAccountId: req.FrromAccountID,
	ToAccountId: req.ToAccountID,
	Amount: req.Amount,
	
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK,result)

}


func (server  *Server) validAccount(ctx *gin.Context ,accountId int64 , currency string)bool{
	account , err := server.store.GetAccount(ctx, accountId) 

	
	if err!=nil {
        if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound , errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError , errorResponse(err)) 

		return  false
	}
 

	if account.Currency!=currency {

		err :=  fmt.Errorf("account [%d]  currency mismatch : %s vs %s", account.ID , account.Currency , currency)
        
		 	ctx.JSON(http.StatusBadRequest , errorResponse(err))
			return false
	}


	 
return true
	
}