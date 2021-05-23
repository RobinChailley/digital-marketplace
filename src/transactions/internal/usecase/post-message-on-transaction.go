package usecase

import (
	"errors"
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/request"

	"github.com/go-pg/pg/v10"
)

type PostMessageOnTransactionCmd func (db *pg.DB, messageRequest *request.MessageRequest, user domain.Account) (domain.Transaction, error)

func PostMessageOnTransaction() PostMessageOnTransactionCmd {

	return func (db *pg.DB, messageRequest *request.MessageRequest, user domain.Account) (domain.Transaction, error) {

		transac := domain.Transaction{}

		err := db.Model(&transac).
			Column("transaction.*").
			Relation("Buyer").
			Relation("Ads").
			Relation("Seller").
			Relation("Messages").
			Where("transaction.seller_id = ?", user.Id).
			WhereOr("transaction.buyer_id = ?", user.Id).
			Select()
			
		if err != nil {
			return domain.Transaction{}, err
		}

		if user.Id != transac.SellerId && user.Id != transac.BuyerId {
			return domain.Transaction{}, errors.New("unauthorized : this transaction is not yours")
		}

		message := domain.Message{
			SenderId: user.Id,
			Sender: &user,
			TransactionId: transac.Id,
			Transaction: &transac,
			Message: messageRequest.Message,
		}

		_, err = db.Model(&message).Insert()



		if err != nil {
			return domain.Transaction{}, err
		}
		message.Transaction = nil
		transac.Messages = append(transac.Messages, message)

		return transac, nil
	}

}
