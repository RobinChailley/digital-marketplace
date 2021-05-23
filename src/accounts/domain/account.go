package domain

import "fmt"

type Account struct {
	Id 					int64  	`pg:",notnull"`
	Email 			string 	`pg:",notnull,unique"`
	Username 		string 	`pg:",notnull,unique"`
	Password 		string 	`pg:",notnull"`
	Balance 		float64	`pg:",use_zero"`
	Ads					[]Ads   `pg:"rel:has-many"`
}

func StringAccount(a *Account) string {
	return fmt.Sprintf("Account<id(%d) email(%s) username(%s) password(%s) balance(%d)>", a.Id, a.Email, a.Username, a.Password, a.Balance)
}