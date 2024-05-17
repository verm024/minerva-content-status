package dto

type GetContentManagementDashboardRepoFilter struct {
	SortBy string `validate:"enum_validator=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string
	Status string `validate:"enum_validator=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}

type GetContentManagementDashboardUseCaseFilter struct {
	SortBy string `validate:"enum_validator=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string
	Status string `validate:"enum_validator=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}

type GetContentManagementDashboardRequestQuery struct {
	SortBy string `query:"sort_by" validate:"enum_validator=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string `query:"search"`
	Status string `query:"status" validate:"enum_validator=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}
