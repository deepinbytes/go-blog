package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/deepinbytes/go-blog/apis"
	"github.com/deepinbytes/go-blog/app"
	"github.com/deepinbytes/go-blog/daos"
	"github.com/deepinbytes/go-blog/errors"
	"github.com/deepinbytes/go-blog/services"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.Open("postgres", app.Config.DSN)
	if err != nil {
		fmt.Errorf("db error : %s", err)
	}

	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/v1")

	rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	//rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
	//	SigningMethod: app.Config.JWTSigningMethod,
	//	TokenHandler:  apis.JWTHandler,
	//}))

	// wire up resource APIs here
	articleDAO := daos.NewArticleDAO()
	apis.ServeArticleResource(rg, services.NewArticleService(articleDAO))

	return router
}
