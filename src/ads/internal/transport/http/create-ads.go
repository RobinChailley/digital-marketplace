package http

import (
	b64 "encoding/base64"
	"io/ioutil"
	"marketplace/ads/domain"
	"marketplace/ads/internal/request"
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/sirupsen/logrus"
	//"net/http"
)


func CreateAdsHandler(db *pg.DB, cmd usecase.CreateAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")
		price := c.PostForm("price")
		file, _ := c.FormFile("file")

		floatPrice, err := strconv.ParseFloat(price, 10)

		if err != nil {
			logrus.WithError(err).Error("The price field must be a float")
			c.Status(http.StatusBadRequest)
			return
		}

		openedFile, err := file.Open() // TODO : petit probleme ici parfois?

		if err != nil {
			logrus.WithError(err).Error("Can not open the file")
			c.Status(http.StatusInternalServerError)
			return
		}
	
		defer openedFile.Close()

		content, err := ioutil.ReadAll(openedFile)

		if err != nil {
			logrus.WithError(err).Error("Can not read the file")
			c.Status(http.StatusInternalServerError)
			return
		}


		sEnc := b64.StdEncoding.EncodeToString(content)
				


		createAdsRequest := &request.UpdateAdsRequest{
			Title: title,
			Description: description,
			Price: floatPrice,
			Picture: sEnc,
		}

		user := c.MustGet("acc").(domain.Account)

		ads, err := cmd(db, createAdsRequest, user.Id)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, ads)
	}
}
