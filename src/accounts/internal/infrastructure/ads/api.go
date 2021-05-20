package ads

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/conf"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type API struct {
	client  *http.Client
	address string
	apiKey  string
}

func NewAPI(config conf.AdsService) *API {
	client := &http.Client{}
	return &API{client, config.URL, config.ApiKey}
}

func (a *API) GetMyAds(c *gin.Context) ([]domain.Ads, error) {
	uri := fmt.Sprintf("%s/list", a.address)

	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return []domain.Ads{}, errors.Wrap(err, "Unable to build http request.")
	}

	request = request.WithContext(c)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := a.client.Do(request)

	if err != nil {
		return []domain.Ads{}, errors.Wrap(err, "Unable to handle the http request.")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []domain.Ads{}, errors.New("Can not read the body.")
		}
		userAds := []domain.Ads{}
		err = json.Unmarshal(data, &userAds)
		if err != nil {
			return []domain.Ads{}, errors.Wrap(err, "Can not json.Unmarshal body")
		}
		return userAds, nil
	default:
		return []domain.Ads{}, errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}

func (a *API) DeleteAllMyAds(c *gin.Context) (error) {
	uri := fmt.Sprintf("%s/delete/all", a.address)

	request, err := http.NewRequest(http.MethodDelete, uri, nil)
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
	default:
		return errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}

