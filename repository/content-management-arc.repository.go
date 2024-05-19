package repository

import (
	"minerva-content-status/dto"
	"minerva-content-status/models"

	"gorm.io/gorm"
)

type ContentManagementArcRepositoryInterface interface {
	CreateCMArc(data *dto.CreateCMArcRepoInputDTO) (*models.ContentManagementArc, error)
	UpdateCMArc(data *dto.UpdateCMArcRepoInputDTO) error
	GetById(arcId uint64) (*models.ContentManagementArc, error)
	DeleteCMArc(arcId uint64) error
	FindAllCmArcByCmId(contentId uint64) ([]models.ContentManagementArc, error)
}

type ContentManagementArcRepository struct {
	Db *gorm.DB
}

func InitializeContentManagementArcRepository(db *gorm.DB) *ContentManagementArcRepository {
	return &ContentManagementArcRepository{db}
}

func (repo *ContentManagementArcRepository) CreateCMArc(data *dto.CreateCMArcRepoInputDTO) (*models.ContentManagementArc, error) {
	cmArc := models.ContentManagementArc{Title: data.Title, Description: data.Description, ContentManagementId: data.ContentManagementId}
	result := repo.Db.Create(&cmArc)

	if result.Error != nil {
		return nil, result.Error
	}

	return &cmArc, nil
}

func (repo *ContentManagementArcRepository) UpdateCMArc(data *dto.UpdateCMArcRepoInputDTO) error {
	result := repo.Db.Model(&models.ContentManagementArc{}).Where("content_management_arc_id = ?", data.ContentManagementArcId).Updates(models.ContentManagementArc{Title: data.Title, Description: data.Description, IsFinal: data.IsFinal, IsVoiceRecorded: data.IsVoiceRecorded, IsEdited: data.IsEdited})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ContentManagementArcRepository) GetById(arcId uint64) (*models.ContentManagementArc, error) {
	cmArc := models.ContentManagementArc{}
	result := repo.Db.Model(&cmArc).Where("content_management_arc_id = ?", arcId).Find(&cmArc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cmArc, nil
}

func (repo *ContentManagementArcRepository) DeleteCMArc(arcId uint64) error {
	cmArc := models.ContentManagementArc{}
	result := repo.Db.Model(&cmArc).Where("content_management_arc_id = ?", arcId).Delete(&cmArc)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ContentManagementArcRepository) FindAllCmArcByCmId(contentId uint64) ([]models.ContentManagementArc, error) {
	cmArcs := []models.ContentManagementArc{}

	result := repo.Db.Model(&models.ContentManagementArc{}).Where("content_management_id = ?", contentId).Find(&cmArcs)
	if result.Error != nil {
		return []models.ContentManagementArc{}, result.Error
	}

	return cmArcs, nil
}
