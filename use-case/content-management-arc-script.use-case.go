package usecase

import (
	"minerva-content-status/dto"
	"minerva-content-status/repository"
	"net/http"

	"gorm.io/gorm"
)

type ContentManagementArcScriptUseCaseInterface interface {
	CreateCMAScript(data *dto.CreateCMAScriptUseCaseInputDTO) dto.CustomErrorInterface
	UpdateCMAScript(data *dto.UpdateCMAScriptUseCaseInputDTO) dto.CustomErrorInterface
	DeleteCMAScript(scriptId uint64) dto.CustomErrorInterface
	CMAScriptListByCMAId(cmaId uint64) (*dto.CMAScriptListByCMAIdUseCaseOutputDTO, dto.CustomErrorInterface)
}

type ContentManagementArcScriptUseCase struct {
	repo    repository.ContentManagementArcScriptRepoInterface
	cmarepo repository.ContentManagementArcRepositoryInterface
	db      *gorm.DB
}

func InitializeContentManagementArcScriptUseCase(repo repository.ContentManagementArcScriptRepoInterface, cmarepo repository.ContentManagementArcRepositoryInterface, db *gorm.DB) *ContentManagementArcScriptUseCase {
	uc := ContentManagementArcScriptUseCase{repo, cmarepo, db}
	return &uc
}

func (uc *ContentManagementArcScriptUseCase) CreateCMAScript(data *dto.CreateCMAScriptUseCaseInputDTO) dto.CustomErrorInterface {

	tx := uc.db.Begin()
	oneCma, cmRepoFindError := uc.cmarepo.GetById(data.ContentManagementArcId)
	if cmRepoFindError != nil {
		return cmRepoFindError
	}

	if oneCma.ContentManagementArcId != data.ContentManagementArcId {
		return dto.NewCustomError("content management arc not found", http.StatusNotFound)
	}

	_, err := uc.repo.CreateCMAScript(&dto.CreateCMAScriptRepoInputDTO{ArcScript: data.ArcScript, ContentManagementArcId: data.ContentManagementArcId})

	if err != nil {
		tx.Rollback()
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcScriptUseCase) UpdateCMAScript(data *dto.UpdateCMAScriptUseCaseInputDTO) dto.CustomErrorInterface {
	tx := uc.db.Begin()

	oneCmascript, findOneErr := uc.repo.GetById(data.ContentManagementArcScriptId)

	if findOneErr != nil {
		tx.Rollback()
		return findOneErr
	}

	if oneCmascript.ContentManagementArcScriptId != data.ContentManagementArcScriptId {
		tx.Rollback()
		return dto.NewCustomError("content management arc script not found", http.StatusNotFound)
	}

	updateErr := uc.repo.UpdateCMAScript(&dto.UpdateCMAScriptRepoInputDTO{ContentManagementArcScriptId: data.ContentManagementArcScriptId, ArcScript: data.ArcScript})
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcScriptUseCase) DeleteCMAScript(scriptId uint64) dto.CustomErrorInterface {
	tx := uc.db.Begin()

	oneCmascript, findOneErr := uc.repo.GetById(scriptId)

	if findOneErr != nil {
		tx.Rollback()
		return findOneErr
	}

	if oneCmascript.ContentManagementArcScriptId != scriptId {
		tx.Rollback()
		return dto.NewCustomError("content management arc script not found", http.StatusNotFound)
	}

	deleteErr := uc.repo.DeleteCMAScript(scriptId)

	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcScriptUseCase) CMAScriptListByCMAId(cmaId uint64) (*dto.CMAScriptListByCMAIdUseCaseOutputDTO, dto.CustomErrorInterface) {
	oneCma, cmRepoFindError := uc.cmarepo.GetById(cmaId)
	if cmRepoFindError != nil {
		return &dto.CMAScriptListByCMAIdUseCaseOutputDTO{Script: []map[string]interface{}{}}, cmRepoFindError
	}

	if oneCma.ContentManagementArcId != cmaId {
		return &dto.CMAScriptListByCMAIdUseCaseOutputDTO{Script: []map[string]interface{}{}}, dto.NewCustomError("content management arc not found", http.StatusNotFound)
	}

	scriptList, scriptListFindErr := uc.repo.FindAllCMAScriptByArcId(cmaId)

	if scriptListFindErr != nil {
		return &dto.CMAScriptListByCMAIdUseCaseOutputDTO{Script: []map[string]interface{}{}}, scriptListFindErr
	}

	scriptListMap := []map[string]interface{}{}
	for _, item := range scriptList {
		computedItem := map[string]interface{}{"content_management_arc_script_id": item.ContentManagementArcScriptId, "arc_script": item.ArcScript}
		scriptListMap = append(scriptListMap, computedItem)
	}

	return &dto.CMAScriptListByCMAIdUseCaseOutputDTO{Script: scriptListMap}, nil
}
