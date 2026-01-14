package types

import (
	"go_server/helper"
)

var IntegrationRequestQueryType = helper.GlobalPaginatedQueryResolver(
	IntegrationType,
	"IntegrationRequestQuery",
)
