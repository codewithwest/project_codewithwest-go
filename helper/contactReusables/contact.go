package contactReusables

import "github.com/graphql-go/graphql"

var ContactMessageInputDef = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "ContactMessageInput",
	Description: "Input for creating a new contact message",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Name of the sender",
		},
		"email": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Email of the sender",
		},
		"message": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The message content",
		},
	},
})

type ContactMessageMongo struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Message   string `json:"message" bson:"message"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
