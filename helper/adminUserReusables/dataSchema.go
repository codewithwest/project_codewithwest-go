package adminUserReusables

type AdminUser struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  *string `json:"password"`
	Role      string  `json:"role"`
	Type      string  `json:"type"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	LastLogin string  `json:"last_login"`
}
type AdminUserInput struct {
	ID       int     `json:"id"`
	UserName string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type AdminUserRequest struct {
	ID        int    `json:"id" bson:"id"`
	Email     string `json:"email" bson:"email"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}

type AdminUserInputMongo struct {
	ID        int     `json:"id" bson:"id"`
	UserName  string  `json:"name" bson:"username"`
	Email     string  `json:"email" bson:"email"`
	Password  *string `json:"password" bson:"password"`
	Role      string  `json:"role" bson:"role"`
	Type      string  `json:"type" bson:"type"`
	Status    string  `json:"status" bson:"status"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	LastLogin *string `json:"last_login" bson:"last_login"`
}

type AdminUserIsAdministrator struct {
	ID   int    `json:"id" bson:"id"`
	Role string `json:"role" bson:"role"`
}
