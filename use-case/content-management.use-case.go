package usecase

import (
	"minerva-content-status/dto"
	"minerva-content-status/repository"
)

type ContentManagementUseCaseInterface interface {
	GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) (*dto.GetContentManagementDashboardUseCaseOutputDTO, error)
	CreateContent(contentData *dto.CreateContentDTO) error
	UpdateContent(contentData *dto.UpdateContentDTO) error
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
	_, err := uc.repo.CreateContent(contentData)
	return err
}

func (uc *ContentManagementUseCase) UpdateContent(contentData *dto.UpdateContentDTO) error {
	_, err := uc.repo.UpdateContent(contentData)
	return err
}
