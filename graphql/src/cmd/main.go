package main

import (
	"log"
	"os"

	gmGraphql "github.com/tanaka-takuto/goal-minder/adapter/graphql"
	"github.com/tanaka-takuto/goal-minder/adapter/graphql/resolver"
	"github.com/tanaka-takuto/goal-minder/infra/db"
	"github.com/tanaka-takuto/goal-minder/infra/echo/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

const defaultPort = "3000"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.SetRequestID())
	e.Use(middleware.Authentication())
	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.POST("/query", func(c echo.Context) error {
		srv := handler.NewDefaultServer(gmGraphql.NewExecutableSchema(gmGraphql.Config{Resolvers: &resolver.Resolver{
			DB: db.Con,
		}}))
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
