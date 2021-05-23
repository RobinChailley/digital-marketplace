package accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/conf"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type API struct {
	client  *http.Client
	address string
	apiKey  string
}

func NewAPI(config conf.Service) *API {
	client := &http.Client{}
	return &API{client, config.URL, config.ApiKey}
}

func (a *API) GetUserById(c *gin.Context, userId int64) (domain.Account, error) {
	uri := fmt.Sprintf("%s/info/byId/%d", a.address, userId)

	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return domain.Account{}, errors.Wrap(err, "Unable to build http request.")
	}

	request = request.WithContext(c)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := a.client.Do(request)

	if err != nil {
		return domain.Account{}, errors.Wrap(err, "Unable to handle the http request.")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return domain.Account{}, errors.New("Can not read the body.")
		}

		user := domain.Account{}
		err = json.Unmarshal(data, &user)

		if err != nil {
			return domain.Account{}, errors.Wrap(err, "Can not json.Unmarshal body")
		}
		return user, nil
	default:
		return domain.Account{}, errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}


func (a *API) UpdateUserBalanceById(c *gin.Context, userId int64, deltaBalance float64) (error) {
	uri := fmt.Sprintf("%s/update-balance/byId/%d", a.address, userId)

	jsonString := []byte(fmt.Sprintf(`{"balance": %f}`, deltaBalance))
	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonString))
	if err != nil {
		return errors.Wrap(err, "Unable to build http request.")
	}

	request = request.WithContext(c)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := a.client.Do(request)

	if err != nil {
		return errors.Wrap(err, "Unable to handle the http request.")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New("The user balance is too low to make this update.")
	default:
		return errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}