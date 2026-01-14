package contact

import (
	"go_server/helper"
	"go_server/resolver/queries"

	"github.com/graphql-go/graphql"
)

func GetContactMessagesSchema(contactMessageType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        contactMessageType,
		Description: "Get all contact messages",
		Args:        helper.GlobalPaginatedQueriesInput,
		Resolve:     queries.GetContactMessages,
	}
}
