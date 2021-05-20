package request

type UpdateMeRequest struct {
	Password string         		`json:"password"`
	Email 	 string					`json:"email"`
	Username string					`json:"username"`
}

type UpdateMeResponse struct {
	Email 	 string					`json:"email"`
	Username string					`json:"username"`
	Balance  int64 					`json:"balance"`
	Id 		 int64 					`json:"id"`
}