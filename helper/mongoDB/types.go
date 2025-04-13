package mongoDB

var NameType struct {
	Name string `bson:"name"`
}

var UserNameType struct {
	UserName string `bson:"username"`
}

var EmailType struct {
	Email string `bson:"email"`
}
