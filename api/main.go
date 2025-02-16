package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type MainController struct {
	Schema graphql.Schema
}

func (mainController *MainController) GetData() http.Handler {
	return handler.New(&handler.Config{
		Schema:   &mainController.Schema,
		Pretty:   true,
		GraphiQL: false,
	})
}
