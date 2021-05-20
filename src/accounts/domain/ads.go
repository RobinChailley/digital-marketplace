package domain

import "fmt"

type Ads struct {
	Id 				int64  	`pg:",notnull"`
	Title 			string 	`pg:",notnull"`
	Description 	string 	`pg:",notnull"`
	Price 			float64 `pg:",notnull"`
	UserId			int64	`pg:",notnull,fk"`
	Picture 		string  `pg:",notnull"`
}

func StringAds(a *Ads) string {
	return fmt.Sprintf("Ads<id(%d) title(%s) description(%s) price(%f) picture(%s)>", a.Id, a.Title, a.Description, a.Price, a.Picture)
}