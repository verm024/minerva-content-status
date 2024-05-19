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

type GetContentManagementDashboardResponseDTO struct {
	ContentList []map[string]interface{} `json:"content_list"`
}

type GetContentManagementDashboardUseCaseOutputDTO struct {
	ContentList []map[string]interface{}
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

// * START UPDATE LINK DTO
type UpdateLinkRepoInputDTO struct {
	ContentManagementId uint64
	TiktokLink          string
	YoutubeLink         string
	IgLink              string
}

// * END UPDATE LINK DTO

// * START PUBLISH AND UPDATE LINK DTO

type PublishAndUpdateLinkUseCaseInputDTO struct {
	ContentManagementId uint64
	TiktokLink          string
	YoutubeLink         string
	IgLink              string
}

type PublishAndUpdateLinkRequestDTO struct {
	ContentManagementId uint64 `param:"content_management_id" validate:"required"`
	TiktokLink          string `json:"tiktok_link"`
	YoutubeLink         string `json:"youtube_link"`
	IgLink              string `json:"ig_link"`
}

// * END PUBLISH AND UPDATE LINK DTO

// * START UPDATE STATUS DTO

type UpdateStatusRepoInputDTO struct {
	ContentManagementId uint64
	Status              string
}

// * END UPDATE STATUS DTO

// * START DELETE CONTENT
type DeleteCMRequestDTO struct {
	ContentManagementId uint64 `param:"content_management_id" validate:"required"`
}

// * END DELETE CONTENT
