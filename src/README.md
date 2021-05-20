# Golang Microservices

## Accounts
```
port: 8080
```

#### The entity
```
- Id (int64, primary key)
- Email (string, unique, not null)
- Username (string, unique, not null)
- Password (string, not null)
- Balance (int64)
```


## Ads
```
port: 8081
```

#### The entity
```
- Id (int64, primary key)
- Title (string, not null)
- Description (string, not null)
- Price (float64)
- Picture (string)
```



## Transactions


## Bonus
