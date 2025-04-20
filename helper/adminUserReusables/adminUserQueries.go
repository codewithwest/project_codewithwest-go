package adminUserReusables

import "github.com/graphql-go/graphql"

type AdminUserRequestMongo struct {
	ID        int    `json:"id" bson:"id"`
	Email     string `json:"email" bson:"email"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
type AdminUsersResponse struct {
	Data       []AdminUserInputMongo `json:"data"`
	Page       int32                 `json:"page"`
	TotalPages int32                 `json:"totalPages"`
	TotalItems int32                 `json:"totalItems"`
}

type AdminUsersRequestResponse struct {
	Data       []AdminUserRequestMongo `json:"data"`
	Page       int32                   `json:"page"`
	TotalPages int32                   `json:"totalPages"`
	TotalItems int32                   `json:"totalItems"`
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
