package main

import (
	"fmt"
	"marketplace/ads/domain"
	"marketplace/ads/internal/conf"
	"marketplace/ads/internal/transport/http"
	"marketplace/ads/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

/**
Create entity schemas in the database
*/
func createSchema(db *pg.DB) error {
	models := []interface{}{
			(*domain.Ads)(nil),
	}

	for _, model := range models {
			err := db.Model(model).CreateTable(&orm.CreateTableOptions{
				IfNotExists: true,
			})
			if err != nil {
				return err
			}
	}
	return nil
}

/**
Assign function to endpoints
*/
func initRoute(router *gin.Engine, db *pg.DB, config conf.Configuration) {
	router.POST("/create", http.AuthMiddleware(db, config), http.CreateAdsHandler(db, usecase.CreateAds()))
	router.GET("/list", http.AuthMiddleware(db, config), http.ListMyAdsHandler(db, usecase.ListUserAds()))
	router.GET("/list/:id", http.AuthMiddleware(db, config), http.ListUserAdsHandler(db, usecase.ListUserAds()))
	router.GET("/search", http.AuthMiddleware(db, config), http.SearchAdsHandler(db, usecase.SearchAds()))
	router.DELETE("/delete/:id", http.AuthMiddleware(db, config), http.DeleteOwnAdsHandler(db, usecase.DeleteOwnAds()))
	router.DELETE("/delete/all", http.AuthMiddleware(db, config), http.DeleteAllMyAdsHandler(db, usecase.DeleteAllMyAds()))
	router.PATCH("/:id", http.AuthMiddleware(db, config), http.UpdateMyAdsHandler(db, usecase.UpdateMyAds()))
}


/**
 Open the config file, decode the yaml content and return the configuration object
 */
func initConfig() (conf.Configuration, error) {
	f, err := os.Open("conf/dev.yaml")
	if err != nil {
		return conf.Configuration{}, errors.Wrap(err, "Impossible d'ouvrir le fichier de configuration.")
	}

	defer func() {f.Close()}()

	var config = &conf.Configuration{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)

	if err != nil {
		return conf.Configuration{}, errors.Wrap(err, "Impossible de décoder le fichier de configuration")
	}
	return *config, nil
}


/**
 Connect to the database, try a simple query to check out the connexion
 */
func logOnDb(config *conf.Configuration) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User: config.DatabaseConfig.User,
		Database: config.DatabaseConfig.Database,
		Addr: config.DatabaseConfig.Addr,
		Password: config.DatabaseConfig.Password,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return db, errors.Wrap(err, "Impossible de se connecter à la base de donné.")
	} else {
		fmt.Println("Successfully logged to the database.")
	}
	return db, nil
}

/**
 The main function
 */
func main() {	
	config, err := initConfig()
	if err != nil {
		panic(err)
	}

	db, err := logOnDb(&config)
	if err != nil {
		panic(err)
	}

	err = createSchema(db)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	initRoute(router, db, config)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}

	defer db.Close()
}