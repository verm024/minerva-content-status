package repository

import (
	"fmt"
	"minerva-content-status/dto"
	"minerva-content-status/models"

	"gorm.io/gorm"
)

type ContentManagementRepository struct {
	db *gorm.DB
}

func InitializeContentManagementRepository(db *gorm.DB) *ContentManagementRepository {
	contentManagementRepository := ContentManagementRepository{db}

	return &contentManagementRepository
}

func (repo *ContentManagementRepository) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) ([]models.ContentManagement, error) {
	contentManagement := []models.ContentManagement{}

	query := repo.db.Model(&contentManagement)

	if filter.Search != "" {
		query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", filter.Search))
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

func (repo *ContentManagementRepository) CreateContent(contentData *dto.CreateContentDTO) (*models.ContentManagement, error) {
	content := models.ContentManagement{
		Title:       contentData.Title,
		Description: contentData.Description,
	}

	result := repo.db.Create(&content)

	if result.Error != nil {
		return &models.ContentManagement{}, result.Error
	}

	return &content, nil
}

func (repo *ContentManagementRepository) UpdateContent(contentData *dto.UpdateContentDTO) (*models.ContentManagement, error) {
	contentManagement := models.ContentManagement{}

	result := repo.db.Model(&contentManagement).Where("content_management_id = ?", contentData.ContentManagementId).First(&contentManagement).Updates(models.ContentManagement{Title: contentData.Title, Description: contentData.Description})

	if result.Error != nil {
		return nil, result.Error
	}

	return &contentManagement, nil
}
