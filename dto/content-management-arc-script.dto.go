package dto

// * START CREATE
type CreateCMAScriptRepoInputDTO struct {
	ContentManagementArcId uint64
	ArcScript              string
}

type CreateCMAScriptUseCaseInputDTO struct {
	ContentManagementArcId uint64
	ArcScript              string
}

type CreateCMAScriptRequestDTO struct {
	ContentManagementArcId uint64 `param:"content_management_arc_id" validate:"required"`
	ArcScript              string `json:"arc_script" validate:"required"`
}

// * END CREATE

// * START UPDATE

type UpdateCMAScriptRepoInputDTO struct {
	ArcScript                    string
	ContentManagementArcScriptId uint64
}

type UpdateCMAScriptUseCaseInputDTO struct {
	ContentManagementArcScriptId uint64
	ArcScript                    string
}

type UpdateCMAScriptRequestDTO struct {
	ContentManagementArcScriptId uint64 `param:"content_management_arc_script_id" validate:"required"`
	ArcScript                    string `json:"arc_script"`
}

// * END UPDATE

// * START DELETE
type DeleteCMAScriptRequestDTO struct {
	ContentManagementArcScriptId uint64 `param:"content_management_arc_script_id" validate:"required"`
}

// * END DELETE

// * START GET CMA Script List
type CMAScriptListByCMAIdUseCaseOutputDTO struct {
	Script []map[string]interface{}
}

type CMAScriptListByCMAIdRequestDTO struct {
	ContentManagementArcId uint64 `param:"content_management_arc_id" validate:"required"`
}

type CMAScriptListByCMAIdResponseDTO struct {
	Script []map[string]interface{} `json:"script"`
}

// * END GET CMA SCRIPT LIST
