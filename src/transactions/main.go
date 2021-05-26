package main

import (
	"fmt"
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/conf"
	"marketplace/transactions/internal/infrastructure/accounts"
	"marketplace/transactions/internal/infrastructure/ads"
	"marketplace/transactions/internal/transport/http"
	"marketplace/transactions/internal/usecase"
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
		(*domain.Transaction)(nil),
		(*domain.Message)(nil),
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
func initRoute(router *gin.Engine, db *pg.DB, config conf.Configuration, adsFetcher ads.Fetcher, accountsFetcher accounts.Fetcher) {
	router.POST("/", http.AuthMiddleware(db, config), http.CreateTransactionHandler(db, usecase.CreateTransaction(adsFetcher, accountsFetcher)))
	router.GET("/", http.AuthMiddleware(db, config), http.GetMyTransactionsHandler(db, usecase.GetMyTransactions()))
	router.POST("/message/:id", http.AuthMiddleware(db, config), http.PostMessageOnTransactionHandler(db, usecase.PostMessageOnTransaction()))
	router.POST("/:id/accept", http.AuthMiddleware(db, config), http.AcceptDeclineTransactionHandler(db, usecase.AcceptDeclineTransaction("accept", adsFetcher, accountsFetcher)))
	router.POST("/:id/decline", http.AuthMiddleware(db, config), http.AcceptDeclineTransactionHandler(db, usecase.AcceptDeclineTransaction("decline", adsFetcher, accountsFetcher)))
	router.POST("/:id/cancel", http.AuthMiddleware(db, config), http.CancelTransactionHandler(db, usecase.CancelTransaction()))
}

/**
Open the config file, decode the yaml content and return the configuration object
*/
func initConfig() (conf.Configuration, error) {
	f, err := os.Open("conf/dev.yaml")
	if err != nil {
		return conf.Configuration{}, errors.Wrap(err, "Impossible d'ouvrir le fichier de configuration.")
	}

	defer func() { f.Close() }()

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
		User:     config.DatabaseConfig.User,
		Database: config.DatabaseConfig.Database,
		Addr:     config.DatabaseConfig.Addr,
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

	adsFetcher := ads.NewAPI(config.AdsService)
	accountsFetcher := accounts.NewAPI(config.AccountsService)

	router := gin.Default()
	initRoute(router, db, config, adsFetcher, accountsFetcher)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}

	defer db.Close()
}
