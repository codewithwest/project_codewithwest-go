package helper

type ProjectCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProjectMongo struct {
	ID                int      `json:"id" bson:"id"`
	ProjectCategoryId int      `json:"project_category_id" bson:"project_category_id"`
	Name              string   `json:"name" bson:"name"`
	Description       string   `json:"description" bson:"description"`
	TechStacks        []string `json:"tech_stacks" bson:"tech_stacks"`
	GithubLink        string   `json:"github_link" bson:"github_link"`
	LiveLink          string   `json:"live_link" bson:"live_link"`
	TestLink          string   `json:"test_link" bson:"test_link"`
	CreatedAt         string   `json:"created_at" bson:"created_at"`
	UpdatedAt         string   `json:"updated_at" bson:"updated_at"`
}

type ProjectCategoryMongo struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
