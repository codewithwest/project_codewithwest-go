package projectCategoryReusables

import "go_server/helper"

var ProjectCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProjectCategoryMongo struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

func NewProjectCategory(id int, name string) ProjectCategoryMongo {
	return ProjectCategoryMongo{
		ID:        id + 1,
		Name:      name,
		CreatedAt: helper.GetCurrentDateTime(),
		UpdatedAt: helper.GetCurrentDateTime(),
	}
}
