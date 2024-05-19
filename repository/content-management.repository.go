package repository

import (
	"fmt"
	"minerva-content-status/dto"
	"minerva-content-status/models"

	"gorm.io/gorm"
)

type ContentManagementRepoInterface interface {
	GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) ([]models.ContentManagement, error)
	CreateContent(contentData *dto.CreateContentDTO) (*models.ContentManagement, error)
	UpdateContent(contentData *dto.UpdateContentDTO) (*models.ContentManagement, error)
	DeleteContent(contentId uint64) error
	UpdateLink(data *dto.UpdateLinkRepoInputDTO) error
	UpdateStatus(data *dto.UpdateStatusRepoInputDTO) error
}

type ContentManagementRepository struct {
	Db *gorm.DB
}

func InitializeContentManagementRepository(db *gorm.DB) *ContentManagementRepository {
	contentManagementRepository := ContentManagementRepository{db}

	return &contentManagementRepository
}

func (repo *ContentManagementRepository) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) ([]models.ContentManagement, error) {
	contentManagement := []models.ContentManagement{}

	query := repo.Db.Model(&contentManagement)

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

	result := repo.Db.Create(&content)

	if result.Error != nil {
		return &models.ContentManagement{}, result.Error
	}

	return &content, nil
}

func (repo *ContentManagementRepository) UpdateContent(contentData *dto.UpdateContentDTO) (*models.ContentManagement, error) {
	contentManagement := models.ContentManagement{}

	result := repo.Db.Model(&contentManagement).Where("content_management_id = ?", contentData.ContentManagementId).First(&contentManagement).Updates(models.ContentManagement{Title: contentData.Title, Description: contentData.Description})

	if result.Error != nil {
		return nil, result.Error
	}

	return &contentManagement, nil
}

func (repo *ContentManagementRepository) DeleteContent(contentId uint64) error {
	contentManagement := models.ContentManagement{}
	result := repo.Db.Model(&contentManagement).Where("content_management_id = ?", contentId).Delete(&contentManagement)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ContentManagementRepository) UpdateLink(data *dto.UpdateLinkRepoInputDTO) error {
	contentManagement := models.ContentManagement{}
	result := repo.Db.Model(&contentManagement).Where("content_management_id = ?", data.ContentManagementId).Updates(models.ContentManagement{TiktokLink: data.TiktokLink, YoutubeLink: data.YoutubeLink, IgLink: data.IgLink})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ContentManagementRepository) UpdateStatus(data *dto.UpdateStatusRepoInputDTO) error {
	contentManagement := models.ContentManagement{}
	result := repo.Db.Model(&contentManagement).Where("content_management_id = ?", data.ContentManagementId).Updates(models.ContentManagement{Status: data.Status})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ContentManagementRepository) FindOneById(contentId uint64) (*models.ContentManagement, error) {
	cm := models.ContentManagement{}
	result := repo.Db.Model(&cm).Where("content_management_id = ?", contentId).Find(&cm)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cm, nil
}
