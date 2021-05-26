package request

import (
	"marketplace/transactions/domain"
)

type GetAdsResponse struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	UserId      int64   `json:"userId"`
	Picture     string  `json:"-"`
	Sold        bool    `json:"sold"`
}

type GetMessageResponse struct {
	Id            int64               `json:"id"`
	SenderId      int64               `json:"senderId"`
	Sender        *domain.Account     `json:"-"`
	TransactionId int64               `json:"transactionId"`
	Transaction   *domain.Transaction `json:"-"`
	Message       string              `json:"message"`
}

type GetAccountResponse struct {
	Id       int64        `json:"id"`
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Password string       `json:"-"`
	Balance  float64      `json:"-"`
	Ads      []domain.Ads `json:"-"`
	Admin    bool         `json:"admin"`
}

type GetTransactionsResponse struct {
	Id       int64                `json:"id"`
	Buyer    GetAccountResponse   `json:"buyer"`
	BuyerId  int64                `json:"buyerId"`
	Seller   GetAccountResponse   `json:"seller"`
	SellerId int64                `json:"sellerId"`
	Ads      GetAdsResponse       `json:"ads"`
	AdsId    int64                `json:"adsId"`
	Messages []GetMessageResponse `json:"messages"`
	Bid      float64              `json:"bid"`
	Status   string               `json:"status"`
}

func ConvertToResponse(transacs []domain.Transaction) []GetTransactionsResponse {
	convertedTransacs := []GetTransactionsResponse{}

	for _, transac := range transacs {

		convertedMessages := []GetMessageResponse{}
		for _, message := range transac.Messages {
			convertedMessages = append(convertedMessages, GetMessageResponse{
				Id:            message.Id,
				SenderId:      message.SenderId,
				Sender:        message.Sender,
				TransactionId: message.TransactionId,
				Transaction:   message.Transaction,
				Message:       message.Message,
			})
		}

		convertedTransacs = append(convertedTransacs, GetTransactionsResponse{
			Id:       transac.Id,
			Buyer:    GetAccountResponse(*transac.Buyer),
			BuyerId:  transac.BuyerId,
			Seller:   GetAccountResponse(*transac.Seller),
			SellerId: transac.SellerId,
			Ads:      GetAdsResponse(*transac.Ads),
			AdsId:    transac.AdsId,
			Messages: convertedMessages,
			Bid:      transac.Bid,
			Status:   transac.Status,
		})
	}

	return convertedTransacs
}
