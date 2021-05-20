package usecase

import (
	"errors"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/conf"
	"time"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type SignInCmd func(db *pg.DB, acc *domain.Account) (string, error)

type Claims struct {
	Email 			string 		`json:"email"`
	Id 				int64 		`json:"id"`
	jwt.StandardClaims
}

func SignIn(config conf.Configuration) SignInCmd {

	return func(db *pg.DB, acc *domain.Account) (string, error) {

		var accs []domain.Account;
		err := db.Model(&accs).Where("account.email = ? ", acc.Email).Select()

		if err != nil {
			return "", err
		}

		if len(accs) < 1 {
			return "", errors.New("There is not any account that corresponds with this email.")
		}

		err = bcrypt.CompareHashAndPassword([]byte(accs[0].Password),[]byte(acc.Password), )


		if err != nil {
			return "", errors.New("Incorrect password or email");
		}

		var jwtKey = []byte(config.Credentials.JwtSecret)
		expirationTime := time.Now().Add(60 * 24 * time.Minute)

		claims := &Claims{
			Email: acc.Email,
			Id: accs[0].Id,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			logrus.WithError(err).Error();			
			return "", err
		}

		return tokenString, nil
	}
}