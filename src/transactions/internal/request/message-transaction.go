package request

type MessageRequest struct {
	TransactionId		int64 				`json:"transactionId"`
	Message					string				`json:"message"`
}