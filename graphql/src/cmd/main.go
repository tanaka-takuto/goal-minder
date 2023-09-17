package main

import (
	"fmt"
	"log"

	gmGraphql "goal-minder/adapter/graphql"
	"goal-minder/adapter/graphql/directive"
	"goal-minder/adapter/graphql/resolver"
	"goal-minder/config"
	"goal-minder/infra/echo/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

func main() {
	port := config.Port()

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.SetRequestID())
	e.Use(middleware.Authentication())
	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.POST("/query", func(c echo.Context) error {
		srv := handler.NewDefaultServer(gmGraphql.NewExecutableSchema(gmGraphql.Config{
			Resolvers: &resolver.Resolver{},
			Directives: gmGraphql.DirectiveRoot{
				Authorization: directive.AuthorizationDirective,
			},
		}))
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	log.Printf("connect to http://localhost:%v/ for GraphQL playground", port)
	err := e.Start(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
}
