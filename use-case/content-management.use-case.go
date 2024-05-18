package usecase

import (
	"minerva-content-status/dto"
)

func (uc *UseCase) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardUseCaseFilter) ([]map[string]interface{}, error) {
	result, err := uc.repo.GetContentManagementDashboard(&dto.GetContentManagementDashboardRepoFilter{Search: filter.Search, Status: filter.Status, SortBy: filter.SortBy})

	if err != nil {
		return make([]map[string]interface{}, 0), err
	}

	returnedData := make([]map[string]interface{}, len(result))

	for index, item := range result {
		returnedItem := map[string]interface{}{"title": item.Title, "description": item.Description, "content_management_id": item.ContentManagementId}
		returnedData[index] = returnedItem
	}
	return returnedData, err
}

func (uc *UseCase) CreateContent(contentData *dto.CreateContentDTO) error {
	_, err := uc.repo.CreateContent(contentData)
	return err
}

func (uc *UseCase) UpdateContent(contentData *dto.UpdateContentDTO) error {
	_, err := uc.repo.UpdateContent(contentData)
	return err
}
