package auth

var ProtectedMutationsAndQueries = []string{
	// Query
	"getUser",
	"getUsers",
	"loginAdminUser",
	"getAdminUsers",
	"getAdminUserAccessRequests",
	"getProjects",
	"getProjectCategories",
	// "authenticateClient",

	// Mutations
	"createAdminUser",
	"createProjectCategory",
	"createProject",
	"adminUserAccessRequest",
	// "createClient",
}
