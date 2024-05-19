package usecase

import (
	"errors"
	"minerva-content-status/dto"
	"minerva-content-status/repository"
	"net/http"
)

type ContentManagementUseCaseInterface interface {
	GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) (*dto.GetContentManagementDashboardUseCaseOutputDTO, error)
	CreateContent(contentData *dto.CreateContentDTO) error
	UpdateContent(contentData *dto.UpdateContentDTO) dto.CustomErrorInterface
	DeleteContent(contentId uint64) dto.CustomErrorInterface
	PublishAndUpdateLink(data *dto.PublishAndUpdateLinkUseCaseInputDTO) dto.CustomErrorInterface
	UpdateContentStatusProgress(data *dto.UpdateContentStatusProgressUseCaseInputDTO) dto.CustomErrorInterface
}

type ContentManagementUseCase struct {
	repo *repository.ContentManagementRepository
}

func InitializeContentManagementUseCase(repo *repository.ContentManagementRepository) *ContentManagementUseCase {
	contentManagementUseCase := ContentManagementUseCase{repo}

	return &contentManagementUseCase
}

func (uc *ContentManagementUseCase) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) (*dto.GetContentManagementDashboardUseCaseOutputDTO, error) {
	result, err := uc.repo.GetContentManagementDashboard(&dto.GetContentManagementDashboardDTO{Search: filter.Search, Status: filter.Status, SortBy: filter.SortBy})

	if err != nil {
		return &dto.GetContentManagementDashboardUseCaseOutputDTO{ContentList: make([]map[string]interface{}, 0)}, err
	}

	returnedData := make([]map[string]interface{}, len(result))

	for index, item := range result {
		returnedItem := map[string]interface{}{"title": item.Title, "description": item.Description, "content_management_id": item.ContentManagementId}
		returnedData[index] = returnedItem
	}
	return &dto.GetContentManagementDashboardUseCaseOutputDTO{ContentList: returnedData}, err
}

func (uc *ContentManagementUseCase) CreateContent(contentData *dto.CreateContentDTO) error {
	tx := uc.repo.Db.Begin()
	_, err := uc.repo.CreateContent(contentData)

	if err != nil {
		tx.Rollback()
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementUseCase) UpdateContent(contentData *dto.UpdateContentDTO) dto.CustomErrorInterface {
	tx := uc.repo.Db.Begin()
	cm, findCmError := uc.repo.FindOneById(contentData.ContentManagementId)

	if findCmError != nil {
		tx.Rollback()
		return findCmError
	}

	if cm.ContentManagementId != contentData.ContentManagementId {
		tx.Rollback()
		return dto.NewCustomError("content not found", http.StatusNotFound)
	}

	_, updateErr := uc.repo.UpdateContent(contentData)

	if updateErr != nil {
		tx.Rollback()
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}
	return updateErr
}

func (uc *ContentManagementUseCase) DeleteContent(contentId uint64) dto.CustomErrorInterface {
	tx := uc.repo.Db.Begin()
	cm, findCmError := uc.repo.FindOneById(contentId)

	if findCmError != nil {
		tx.Rollback()
		return findCmError
	}

	if cm.ContentManagementId != contentId {
		tx.Rollback()
		return dto.NewCustomError("content not found", http.StatusNotFound)
	}

	deleteErr := uc.repo.DeleteContent(contentId)

	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

func (uc *ContentManagementUseCase) PublishAndUpdateLink(data *dto.PublishAndUpdateLinkUseCaseInputDTO) dto.CustomErrorInterface {
	// * initiate transaction
	tx := uc.repo.Db.Begin()

	cm, findCmError := uc.repo.FindOneById(data.ContentManagementId)

	if findCmError != nil {
		tx.Rollback()
		return findCmError
	}

	if cm.ContentManagementId != data.ContentManagementId {
		tx.Rollback()
		return dto.NewCustomError("content not found", http.StatusNotFound)
	}

	if cm.Status != "WAIT_PUBLISH" {
		tx.Rollback()
		return errors.New("cannot publish content that does not have status 'WAIT_PUBLISH'")
	}

	updateStatusErr := uc.repo.UpdateStatus(&dto.UpdateStatusRepoInputDTO{ContentManagementId: data.ContentManagementId, Status: "PUBLISHED"})
	updateLinkErr := uc.repo.UpdateLink(&dto.UpdateLinkRepoInputDTO{ContentManagementId: data.ContentManagementId, TiktokLink: data.TiktokLink, YoutubeLink: data.YoutubeLink, IgLink: data.IgLink})

	if updateStatusErr != nil {
		tx.Rollback()
		return updateStatusErr
	}

	if updateLinkErr != nil {
		tx.Rollback()
		return updateLinkErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}
	return nil
}

func (uc *ContentManagementUseCase) UpdateContentStatusProgress(data *dto.UpdateContentStatusProgressUseCaseInputDTO) dto.CustomErrorInterface {
	tx := uc.repo.Db.Begin()

	cm, findCmError := uc.repo.FindOneById(data.ContentManagementId)

	if findCmError != nil {
		tx.Rollback()
		return findCmError
	}

	if cm.ContentManagementId != data.ContentManagementId {
		tx.Rollback()
		return dto.NewCustomError("content not found", http.StatusNotFound)
	}

	updateStatusErr := uc.repo.UpdateStatus(&dto.UpdateStatusRepoInputDTO{ContentManagementId: data.ContentManagementId, Status: data.Status})
	if updateStatusErr != nil {
		tx.Rollback()
		return updateStatusErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}
	return nil
}
