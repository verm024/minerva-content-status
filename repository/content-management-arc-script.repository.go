package repository

import (
	"fmt"
	"minerva-content-status/dto"
	"minerva-content-status/models"

	"gorm.io/gorm"
)

type ContentManagementArcScriptRepoInterface interface {
	CreateCMAScript(data *dto.CreateCMAScriptRepoInputDTO) (*models.ContentManagementArcScript, error)
	UpdateCMAScript(data *dto.UpdateCMAScriptRepoInputDTO) error
	GetById(cmascriptId uint64) (*models.ContentManagementArcScript, error)
	DeleteCMAScript(scriptId uint64) error
	FindAllCMAScriptByArcId(arcId uint64) ([]models.ContentManagementArcScript, error)
}

type ContentManagementArcScriptRepository struct {
	Db *gorm.DB
}

func InitializeContentManagementArcScriptRepository(db *gorm.DB) *ContentManagementArcScriptRepository {
	repo := ContentManagementArcScriptRepository{db}
	return &repo
}

func (repo *ContentManagementArcScriptRepository) GetById(cmascriptId uint64) (*models.ContentManagementArcScript, error) {
	cmascript := models.ContentManagementArcScript{}
	result := repo.Db.Model(&cmascript).Where("content_management_arc_script_id = ?", cmascriptId).Find(&cmascript)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cmascript, nil
}

func (repo *ContentManagementArcScriptRepository) CreateCMAScript(data *dto.CreateCMAScriptRepoInputDTO) (*models.ContentManagementArcScript, error) {
	cmascript := models.ContentManagementArcScript{ContentManagementArcId: data.ContentManagementArcId, ArcScript: data.ArcScript}
	result := repo.Db.Create(&cmascript)
	fmt.Println(cmascript.ContentManagementArcId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &cmascript, nil
}

func (repo *ContentManagementArcScriptRepository) UpdateCMAScript(data *dto.UpdateCMAScriptRepoInputDTO) error {
	result := repo.Db.Model(&models.ContentManagementArcScript{}).Where("content_management_arc_script_id = ?", data.ContentManagementArcScriptId).Updates(models.ContentManagementArcScript{ArcScript: data.ArcScript})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ContentManagementArcScriptRepository) DeleteCMAScript(scriptId uint64) error {
	cmascript := models.ContentManagementArcScript{}
	result := repo.Db.Model(&cmascript).Where("content_management_arc_script_id = ?", scriptId).Delete(&cmascript)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ContentManagementArcScriptRepository) FindAllCMAScriptByArcId(arcId uint64) ([]models.ContentManagementArcScript, error) {
	cmascript := []models.ContentManagementArcScript{}

	result := repo.Db.Model(&models.ContentManagementArcScript{}).Where("content_management_arc_id = ?", arcId).Find(&cmascript)
	if result.Error != nil {
		return []models.ContentManagementArcScript{}, result.Error
	}

	return cmascript, nil
}
