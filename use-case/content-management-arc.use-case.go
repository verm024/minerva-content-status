package usecase

import (
	"minerva-content-status/dto"
	"minerva-content-status/repository"
	"net/http"

	"gorm.io/gorm"
)

type ContentManagementArcUseCaseInterface interface {
	CreateCMArc(data *dto.CreateCMArcUseCaseInputDTO) error
	UpdateCMArc(data *dto.UpdateCmArcUseCaseInputDTO) dto.CustomErrorInterface
	DeleteCMArc(arcId uint64) error
	CMArcListByCMId(cmId uint64) (*dto.CMArcListByCMIdUseCaseOutputDTO, error)
}

type ContentManagementArcUseCase struct {
	repo   repository.ContentManagementArcRepositoryInterface
	db     *gorm.DB
	cmrepo repository.ContentManagementRepoInterface
}

func InitializeContentManagementArcUseCase(repo repository.ContentManagementArcRepositoryInterface, cmrepo repository.ContentManagementRepoInterface, db *gorm.DB) *ContentManagementArcUseCase {
	uc := ContentManagementArcUseCase{repo, db, cmrepo}
	return &uc
}

func (uc *ContentManagementArcUseCase) CreateCMArc(data *dto.CreateCMArcUseCaseInputDTO) error {

	tx := uc.db.Begin()
	_, err := uc.repo.CreateCMArc(&dto.CreateCMArcRepoInputDTO{Title: data.Title, Description: data.Description, ContentManagementId: data.ContentManagementId})

	if err != nil {
		tx.Rollback()
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcUseCase) UpdateCMArc(data *dto.UpdateCmArcUseCaseInputDTO) dto.CustomErrorInterface {
	tx := uc.db.Begin()

	oneCmArc, findOneErr := uc.repo.GetById(data.ContentManagementArcId)

	if findOneErr != nil {
		tx.Rollback()
		return findOneErr
	}

	if oneCmArc.ContentManagementArcId != data.ContentManagementArcId {
		tx.Rollback()
		return dto.NewCustomError("content management arc not found", http.StatusNotFound)
	}

	updateErr := uc.repo.UpdateCMArc(&dto.UpdateCMArcRepoInputDTO{Title: data.Title, Description: data.Description, IsFinal: data.IsFinal, IsVoiceRecorded: data.IsVoiceRecorded, IsEdited: data.IsEdited, ContentManagementArcId: data.ContentManagementArcId})
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcUseCase) DeleteCMArc(arcId uint64) error {
	tx := uc.db.Begin()

	oneCmArc, findOneErr := uc.repo.GetById(arcId)

	if findOneErr != nil {
		tx.Rollback()
		return findOneErr
	}

	if oneCmArc.ContentManagementArcId != arcId {
		tx.Rollback()
		return dto.NewCustomError("content management arc not found", http.StatusNotFound)
	}

	deleteErr := uc.repo.DeleteCMArc(arcId)

	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementArcUseCase) CMArcListByCMId(cmId uint64) (*dto.CMArcListByCMIdUseCaseOutputDTO, error) {
	// * validate the existence of content management id
	oneCmRepo, cmRepoFindError := uc.cmrepo.FindOneById(cmId)
	if cmRepoFindError != nil {
		return &dto.CMArcListByCMIdUseCaseOutputDTO{ArcList: []map[string]interface{}{}}, cmRepoFindError
	}

	if oneCmRepo.ContentManagementId != cmId {
		return &dto.CMArcListByCMIdUseCaseOutputDTO{ArcList: []map[string]interface{}{}}, dto.NewCustomError("content management not found", http.StatusNotFound)
	}

	arcList, arcListFindError := uc.repo.FindAllCmArcByCmId(cmId)

	if arcListFindError != nil {
		return &dto.CMArcListByCMIdUseCaseOutputDTO{ArcList: []map[string]interface{}{}}, arcListFindError
	}

	arcListMap := []map[string]interface{}{}
	for _, item := range arcList {
		computedItem := map[string]interface{}{"content_management_arc_id": item.ContentManagementArcId, "title": item.Title, "description": item.Description, "is_edited": item.IsEdited, "is_final": item.IsFinal, "is_voice_recorded": item.IsVoiceRecorded}
		arcListMap = append(arcListMap, computedItem)
	}

	return &dto.CMArcListByCMIdUseCaseOutputDTO{ArcList: arcListMap}, nil
}
