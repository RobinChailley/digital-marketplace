package ads

import (
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

func (a *API) GetAdsById(c *gin.Context, adsId int64) (domain.Ads, error) {
	uri := fmt.Sprintf("%s/%d", a.address, adsId)

	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return domain.Ads{}, errors.Wrap(err, "Unable to build http request.")
	}

	request = request.WithContext(c)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := a.client.Do(request)

	if err != nil {
		return domain.Ads{}, errors.Wrap(err, "Unable to handle the http request.")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return domain.Ads{}, errors.New("Can not read the body.")
		}
		ads := domain.Ads{}
		err = json.Unmarshal(data, &ads)

		if err != nil {
			return domain.Ads{}, errors.Wrap(err, "Can not json.Unmarshal body")
		}
		return ads, nil
	default:
		return domain.Ads{}, errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}

func (a *API) SetSoldToAds(c *gin.Context, adsId int64) (domain.Ads, error) {
	uri := fmt.Sprintf("%s/set-sold/%d", a.address, adsId)

	request, err := http.NewRequest(http.MethodPatch, uri, nil)
	if err != nil {
		return domain.Ads{}, errors.Wrap(err, "Unable to build http request.")
	}

	request = request.WithContext(c)
	request.Header.Set("api-key", a.apiKey)
	request.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := a.client.Do(request)

	if err != nil {
		return domain.Ads{}, errors.Wrap(err, "Unable to handle the http request.")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return domain.Ads{}, errors.New("Can not read the body.")
		}
		ads := domain.Ads{}
		err = json.Unmarshal(data, &ads)

		if err != nil {
			return domain.Ads{}, errors.Wrap(err, "Can not json.Unmarshal body")
		}
		return ads, nil
	default:
		return domain.Ads{}, errors.New(fmt.Sprintf("Unable to handle the http request. Code : %d", resp.StatusCode))
	}
}