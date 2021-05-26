# digital-marketplace

A simple Golang project to learn the language with micro-services. Educational project (Epitech 4rd year)

## Accounts

### Entity

```
# Account

Id            int64
Email         string
Username      string
Password      string
Balance       float64
Ads           []Ads
Admin         bool
```

### Endpoints

- GET("/ping") : Check if the accounts microservices is up
  - OK 200
- POST("/sign-up") : Create an account
  - CREATED 201
  - BAD REQUEST 400
  - CONFLICT 409
  - INTERNAL ERROR 500
- POST("/sign-in") : Sign in the API
  - OK 200
  - BAD REQUEST 400
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : DELETE("/delete-me") : Delete the user's account
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : PATCH("/update-me") : Update the user's account
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : GET("/get-me") : Get informations about the user's account
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : POST("/add-funds") : Add some funds to the user's balance
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
- (AUTH) : GET("/info/:email") : Get informations about a user (specified by email)
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : GET("/info/byId/:id") : Get informations about a user (specified by ID)
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : POST("/update-balance/byId/:id") : Update the user's balance with a delta (user specified by ID)
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500

### BONUS

- (ADMIN AUTH) : GET("/") : Get the full content of every users
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (ADMIN AUTH) : GET("/:id") : Get the full content of a user's account
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (ADMIN AUTH) : DELETE("/:id") : Delete an account
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (ADMIN AUTH) : PATCH("/:id") : Update an account
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500

## Ads

### Entity

```
# Ads

Id                int64
Title             string
Description       string
Price             float64
UserId            int64
Picture           string
Sold              bool
```

### Endpoints

- (AUTH) : POST("/create") : Create an ads
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : GET("/list") : List the user's ads (specified by JWT Token)
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : GET("/list/:id") : List the user's ads (specified by ID)
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : GET("/search?keyword=xxx") : Search an ads by keyword
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : DELETE("/delete/:id") : Delete an ads of the user
  - OK 200
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : DELETE("/delete/all") : Delete all ads of the user
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : PATCH("/:id") : Update an ads
  - OK 200
  - UNAUTHORIZED 401
  - BAD REQUEST 404
  - INTERNAL ERROR 500
- (AUTH) : PATCH("/set-sold/:id") : Set 'true' to the sold field of the add (specified by ID)
  - OK 200
  - UNAUTHORIZED 401
  - BAD REQUEST 404
  - INTERNAL ERROR 500
- (AUTH) : GET("/:id") : Get an ads
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500

## Transactions

### Entities

```
# Transaction

Id                int64
Buyer             *Account
BuyerId           int64
Seller            *Account
SellerId          int64
Ads               *Ads
AdsId             int64
Messages          []Message
Bid               float64
Status            string


# Message

Id                int64
SenderId          int64
Sender            *Account
TransactionId     int64
Transaction       *Transaction
Message           string
```

### Endpoints

- (AUTH) : POST("/") : Create a transaction
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : GET("/") : Get all transactions of the current user
  - OK 200
  - UNAUTHORIZED 401
  - INTERNAL ERROR 500
- (AUTH) : POST("/message/:id") : Post a message in the transaction's conversation
  - CREATED 201
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : POST("/:id/accept") : Accept a transaction
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : POST("/:id/decline") : Decline a transaction
  - CREATED 201
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
- (AUTH) : POST("/:id/cancel") : Cancel a transaction
  - OK 200
  - BAD REQUEST 400
  - UNAUTHORIZED 401
  - NOT FOUND 404
  - INTERNAL ERROR 500
