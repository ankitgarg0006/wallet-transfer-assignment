package requests

type TransferRequest struct {
	SourceWalletID string `json:"source_wallet_id" binding:"required,uuid"`
	DestWalletID   string `json:"dest_wallet_id" binding:"required,uuid"`
	Amount         int64  `json:"amount" binding:"required,min=100,max=1000000"`
}
