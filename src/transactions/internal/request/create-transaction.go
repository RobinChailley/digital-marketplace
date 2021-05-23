package request

type CreateTransactionRequest struct {
	AdsId					int64 				`json:"adsId"`
	Bid						float64				`json:"bid"`
}