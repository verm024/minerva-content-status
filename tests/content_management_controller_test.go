package tests

import (
	"minerva-content-status/controllers"
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockContentManagementUseCase struct {
	mock.Mock
}

func (m *MockContentManagementUseCase) GetContentManagementDashboard(filter *dto.GetContentManagementDashboardDTO) (*dto.GetContentManagementDashboardUseCaseOutputDTO, error) {
	args := m.Called(filter)
	convertedFirstArg := args.Get(0).(dto.GetContentManagementDashboardUseCaseOutputDTO)
	return &convertedFirstArg, args.Error(1)
}

func (m *MockContentManagementUseCase) CreateContent(contentData *dto.CreateContentDTO) error {
	args := m.Called(contentData)
	return args.Error(0)
}

func (m *MockContentManagementUseCase) UpdateContent(contentData *dto.UpdateContentDTO) dto.CustomErrorInterface {
	args := m.Called(contentData)
	return args.Error(0)
}

func (m *MockContentManagementUseCase) DeleteContent(contentId uint64) dto.CustomErrorInterface {

	args := m.Called(contentId)
	return args.Error(0)
}

func (m *MockContentManagementUseCase) PublishAndUpdateLink(data *dto.PublishAndUpdateLinkUseCaseInputDTO) dto.CustomErrorInterface {
	args := m.Called(data)
	return args.Error(0)
}

func TestGetContentManagementDashboard(t *testing.T) {
	const ENDPOINT_PATH string = "/content-management"
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, ENDPOINT_PATH, nil)
	q := req.URL.Query()
	q.Add("search", "lowrence")
	q.Add("sort_by", "CREATED_AT_DESC")
	q.Add("status", "DRAFT")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(ENDPOINT_PATH)

	param := dto.GetContentManagementDashboardDTO{Search: "lowrence", Status: "DRAFT", SortBy: "CREATED_AT_DESC"}
	useCaseResult := dto.GetContentManagementDashboardUseCaseOutputDTO{ContentList: []map[string]interface{}{{"content_management_id": 1, "description": "this is desc", "title": "this is title"}}}
	mockContentManagementUseCase := new(MockContentManagementUseCase)
	mockContentManagementUseCase.On("GetContentManagementDashboard", &param).Return(useCaseResult, nil)

	cont := controllers.InitializeContentManagementController(mockContentManagementUseCase)
	cont.GetContentManagementDashboard(c)
	responseData := dto.GetContentManagementDashboardResponseDTO{}
	decodedResponse, _ := helper_response.DecodeResponseJson(rec.Body.String(), &responseData)

	assert.EqualValues(t, useCaseResult.ContentList[0]["content_management_id"], responseData.ContentList[0]["content_management_id"])
	assert.Equal(t, useCaseResult.ContentList[0]["description"], responseData.ContentList[0]["description"])
	assert.Equal(t, useCaseResult.ContentList[0]["title"], responseData.ContentList[0]["title"])
	assert.Equal(t, http.StatusOK, int(decodedResponse.Code))

	mockContentManagementUseCase.AssertExpectations(t)
}
