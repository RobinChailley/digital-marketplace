package domain

import "fmt"

type Message struct {
	Id 						int64  			`pg:",notnull"`
	SenderId 			int64				`pg:""`
	Sender   			*Account		`pg:"rel:has-one"`
	TransactionId int64				`pg:""`
	Transaction 	*Transaction`pg:"rel:has-one"`
	Message				string			`pg:",notnull`
}

type Transaction struct {
	Id 						int64 			`pg:",notnull"`
	Buyer 				*Account 		`pg:"rel:has-one"`
	BuyerId				int64				`pg:""`
	Seller 				*Account		`pg:"rel:has-one"`
	SellerId			int64				`pg:""`
	Ads						*Ads				`pg:"rel:has-one"`
	AdsId					int64				`pg:""`
	Messages 			[]Message 	`pg:"rel:has-many"`
	Bid			 			float64  		`pg:",use_zero"`
	Status				string			`pg:""`
}

func StringTransaction(a *Transaction) string {
	return fmt.Sprintf("Transaction<id(%d) Buyer(name: %s, id: %d) seller(name: %s, id: %d) Ads(title: %s, id: %d) messages(%s) bid(%d) status(%s)>", a.Id, a.Buyer.Username, a.Buyer.Id, a.Seller.Username, a.Seller.Id, a.Ads.Title, a.Ads.Id, a.Messages, a.Bid, a.Status)
}

func StringMessage(a *Message) string {
	return fmt.Sprintf("Message<id(%d) sender(name: %s, id: %d) transaction(%d) message(%s)>", a.Id, a.SenderId, a.TransactionId, a.Message)
}