package api

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

// func (server *Server) CreateTransfer(ctx *gin.Context) {
// 	var req transferRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.TransferTxParams{
// 		FromAccountID: req.FromAccountID,
// 		ToAccountID:   req.ToAccountID,
// 		Amount:        req.Amount,
// 	}

// 	//result, err := server.store.TransferTx(ctx, arg)

// }
