package request

type UpdateAdsRequest struct {
	Title 				string		`json:"title"`
	Description 		string		`json:"description"`
	Price		 		float64		`json:"price"`
	Picture	 			string		`json:"picture"`
}