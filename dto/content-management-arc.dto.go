package dto

// * START CREATE DTO
type CreateCMArcRepoInputDTO struct {
	ContentManagementId uint64
	Title               string
	Description         string
}

type CreateCMArcUseCaseInputDTO struct {
	ContentManagementId uint64
	Title               string
	Description         string
}

type CreateCMArcRequestDTO struct {
	ContentManagementId uint64 `param:"content_management_id" validate:"required"`
	Title               string `json:"title" validate:"required"`
	Description         string `json:"description"`
}

// * END CREATE DTO

// * START UPDATE DTO
type UpdateCMArcRepoInputDTO struct {
	ContentManagementArcId uint64
	Title                  string
	Description            string
	IsFinal                bool
	IsVoiceRecorded        bool
	IsEdited               bool
}

type UpdateCmArcUseCaseInputDTO struct {
	ContentManagementArcId uint64
	Title                  string
	Description            string
	IsFinal                bool
	IsVoiceRecorded        bool
	IsEdited               bool
}

type UpdateCMArcRequestDTO struct {
	ContentManagementArcId uint64 `param:"content_management_arc_id" validate:"required"`
	Title                  string `json:"title" validate:"required"`
	Description            string `json:"description"`
	IsFinal                bool   `json:"is_final"`
	IsVoiceRecorded        bool   `json:"is_voice_recorded"`
	IsEdited               bool   `json:"is_edited"`
}

// * END UPDATE DTO

// * START GET LIST ARCS BY CM ID

type CMArcListByCMIdRequestDTO struct {
	ContentManagementId uint64 `param:"content_management_id" validate:"required"`
}

type CMArcListByCMIdUseCaseOutputDTO struct {
	ArcList []map[string]interface{}
}

type CMArcListByCMIdResponseDTO struct {
	ArcList []map[string]interface{} `json:"arc_list"`
}

// * END GET LIST ARCS BY CM ID

// * START DELETE DTO
type DeleteCMArcRequestDTO struct {
	ContentManagementArcId uint64 `param:"content_management_arc_id" validate:"required"`
}

// * END DELETE DTO
