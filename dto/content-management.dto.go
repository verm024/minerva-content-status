package dto

// * START GET CONTENT MANAGEMENT DASHBOARD DTO
type GetContentManagementDashboardDTO struct {
	SortBy string `validate:"enum_validator=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string
	Status string `validate:"enum_validator=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}

type GetContentManagementDashboardRequestDTO struct {
	SortBy string `query:"sort_by" validate:"enum_validator=CREATED_AT_DESC CREATED_AT_ASC"`
	Search string `query:"search"`
	Status string `query:"status" validate:"enum_validator=DRAFT WORKING_ON WAIT_PUBLISH PUBLISHED"`
}

// * END GET CONTENT MANAGEMENT DASHBOARD DTO

// * START CREATE CONTENT DTO
type CreateContentDTO struct {
	Title       string `validate:"required"`
	Description string
}

type CreateContentRequestDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

// * END CREATE CONTENT DTO

// * START UPDATE CONTENT DTO
type UpdateContentDTO struct {
	ContentManagementId uint64
	Title               string
	Description         string
}

type UpdateContentRequestDTO struct {
	ContentManagementId uint64 `param:"content_management_id" validate:"required"`
	Title               string `json:"title" validate:"required"`
	Description         string `json:"description"`
}

// * END UPDATE CONTENT DTO
