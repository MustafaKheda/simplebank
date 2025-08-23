package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "github.com/MustafaKheda/simplebank/db/sqlc"
	"github.com/MustafaKheda/simplebank/token"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	log.Printf("createTransfer called with request: %v", ctx.Request.Body)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	authPlayload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)

	if fromAccount.Owner != authPlayload.Username {
		err := fmt.Errorf("account [%d] does not belong to the authenticated user [%s]", req.FromAccountID, authPlayload.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	// Check if the to account is valid
	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency)

	if !valid {
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	account, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch %s vs %s", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}
	return account, true
}
