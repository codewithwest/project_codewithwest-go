package contact

import (
	"go_server/helper/contactReusables"
	"go_server/resolver/mutations"

	"github.com/graphql-go/graphql"
)

func CreateContactMessageMutation(contactMessageType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        contactMessageType,
		Description: "Create a new contact message",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(contactReusables.ContactMessageInputDef),
			},
		},
		Resolve: mutations.CreateContactMessage,
	}
}
