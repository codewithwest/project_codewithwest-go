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
	"getContactMessages",
	"getIntegrations",
	// "authenticateClient",

	// Mutations
	"createAdminUser",
	"createProjectCategory",
	"createProject",
	"adminUserAccessRequest",
	"createIntegration",
	"revokeIntegration",
	// "createClient",
}
