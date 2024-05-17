package repository

import (
	"minerva-content-status/dto"
	"minerva-content-status/models"
)

func (repo *Repository) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardRepoFilter) ([]models.ContentManagement, error) {
	contentManagement := []models.ContentManagement{}

	query := repo.db.Model(&contentManagement)

	if filter.Search != "" {
		query.Where("title LIKE %?%", filter.Search)
	}

	if filter.Status != "" {
		query.Where("status = ?", filter.Status)
	}

	if filter.SortBy != "" {
		switch filter.SortBy {
		case "CREATED_AT_DESC":
			query.Order("created_at DESC")
		case "CREATED_AT_ASC":
			query.Order("created_at ASC")
		}
	}

	result := query.Find(&contentManagement)

	if result.Error != nil {
		return []models.ContentManagement{}, result.Error
	}

	return contentManagement, nil
}
