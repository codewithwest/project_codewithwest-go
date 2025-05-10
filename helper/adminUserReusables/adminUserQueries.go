package adminUserReusables

import (
	"github.com/graphql-go/graphql"
	"go_server/helper"
)

type AdminUserRequestMongo struct {
	ID        int    `json:"id" bson:"id"`
	Email     string `json:"email" bson:"email"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
type AdminUsersResponse struct {
	Data       []AdminUserInputMongo `json:"data"`
	Pagination helper.Pagination     `json:"pagination"`
}

type AdminUsersRequestResponse struct {
	Data       []AdminUserRequestMongo `json:"data"`
	Pagination helper.Pagination       `json:"pagination"`
}

var AdminUserQueriesInput = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type:         graphql.Int,
		DefaultValue: 10,
		Description:  "Number of items per page",
	},
	"page": &graphql.ArgumentConfig{
		Type:         graphql.Int,
		DefaultValue: 1,
		Description:  "Page number",
	},
}
