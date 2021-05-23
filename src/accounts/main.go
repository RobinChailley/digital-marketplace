package main

import (
	"fmt"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/conf"
	"marketplace/accounts/internal/infrastructure/ads"
	"marketplace/accounts/internal/transport/http"
	"marketplace/accounts/internal/usecase"
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
		(*domain.Account)(nil),
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
func initRoute(router *gin.Engine, db *pg.DB, config conf.Configuration, adsFetcher ads.Fetcher) {
	router.GET("/ping", func (c *gin.Context) { c.JSON(200, "pong")})
	router.POST("/sign-up", http.SignUpHandler(db, usecase.SignUp()))
	router.POST("/sign-in", http.SignInHandler(db, usecase.SignIn(config)))
	router.DELETE("/delete-me", http.AuthMiddleware(db, config), http.DeleteMeHandler(db, usecase.DeleteMe(adsFetcher)))
	router.PATCH("/update-me", http.AuthMiddleware(db, config), http.UpdateMeHandler(db, usecase.UpdateMe()))
	router.GET("/get-me", http.AuthMiddleware(db, config), http.GetMeHandler(db, usecase.GetMe(adsFetcher)))
	router.POST("/add-funds", http.AuthMiddleware(db, config), http.AddFundsHandler(db, usecase.AddFunds()))
	router.GET("/info/:email", http.AuthMiddleware(db, config), http.GetUserHandler(db, usecase.GetUser(adsFetcher)))
	router.GET("/info/byId/:id", http.AuthMiddleware(db, config), http.GetUserByIdHandler(db, usecase.GetUserById(adsFetcher)))
	router.POST("/update-balance/byId/:id", http.AuthMiddleware(db, config), http.UpdateBalanceByIdHandler(db, usecase.UpdateBalanceById()))
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

	adsFetcher := ads.NewAPI(config.AdsService)

	router := gin.Default()
	initRoute(router, db, config, adsFetcher)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}

	defer db.Close()
}