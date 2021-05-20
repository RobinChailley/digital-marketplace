package ads

import (
	"fmt"
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

func (a *API) Test(c *gin.Context) (error) {
	uri := fmt.Sprintf("%s/list", a.address)

	request, err := http.NewRequest(http.MethodGet, uri, nil)
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
		return nil
	}
}