package domain

import "fmt"

type Message struct {
	Id 						int64  			`pg:",notnull"`
	SenderId 			int64				`pg:",notnull,fk`
	TransactionId int64				`pg:",notnull,fk`
	Message				string			`pg:",notnull`
}

type Transaction struct {
	Id 						int64  			`pg:",notnull"`
	Buyer 				*Account 		`pg:",notnull"`
	BuyerId				int64				`pg:",fk"`
	Seller 				*Account		`pg:",notnull"`
	SellerId			int64				`pg:",fk"`
	Messages 			[]Message 		`pg:",notnull"`
	Bid			 			int64  			`pg:",use_zero"`
	Status				string			`pg:""`
}

func StringTransaction(a *Transaction) string {
	return fmt.Sprintf("Transaction<id(%d) Buyer(name: %s, id: %d) seller(name: %s, id: %d) messages(%s) bid(%d) status(%s)>", a.Id, a.Buyer.Username, a.Buyer.Id, a.Seller.Username, a.Seller.Id, a.Messages, a.Bid, a.Status)
}

func StringMessage(a *Message) string {
	return fmt.Sprintf("Message<id(%d) sender(name: %s, id: %d) transaction(%d) message(%s)>", a.Id, a.SenderId, a.TransactionId, a.Message)
}