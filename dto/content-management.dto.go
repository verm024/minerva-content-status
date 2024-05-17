package dto

type GetContentManagementDashboardRepoFilter struct {
	SortBy string
	Search string
	Status string
}

type GetContentManagementDashboardUseCaseFilter struct {
	SortBy string
	Search string
	Status string
}

type GetContentManagementDashboardRequestQuery struct {
	SortBy string `query:"sort_by" validate:"oneof=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string `query:"search"`
	Status string `query:"status" validate:"oneof=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}
