package adminUserReusables

type AdminUsersPaginatedResponse struct {
	Data       []AdminUserInputMongo `json:"data"`
	Page       int32                 `json:"page"`
	TotalPages int32                 `json:"totalPages"`
	TotalItems int32                 `json:"totalItems"`
}
