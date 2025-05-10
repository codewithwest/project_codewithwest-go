package projectCategoryReusables

import "go_server/helper"

type ProjectCategoryMongo struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

type ProjectCategoryResponse struct {
	Data       []ProjectCategoryMongo `json:"projectCategory"`
	Pagination helper.Pagination      `json:"pagination"`
}

func NewProjectCategory(id int, name string) ProjectCategoryMongo {
	return ProjectCategoryMongo{
		ID:        id + 1,
		Name:      name,
		CreatedAt: helper.GetCurrentDateTime(),
		UpdatedAt: helper.GetCurrentDateTime(),
	}
}
