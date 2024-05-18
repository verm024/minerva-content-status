package tests

import (
	"minerva-content-status/controllers"
	"minerva-content-status/dto"
	helper_response "minerva-content-status/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterNewUser(param *dto.RegisterNewUserDTO) (string, error) {
	args := m.Called(param)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserUseCase) Login(param *dto.LoginDTO) (string, error) {
	args := m.Called(param)
	return args.Get(0).(string), args.Error(1)
}

func TestRegisterNewUser(t *testing.T) {
	ENDPOINT_PATH := "/user"
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, ENDPOINT_PATH, strings.NewReader(`{"username": "fikri", "password": "1234qwertyY", "email": "kevin@gmail.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(ENDPOINT_PATH)

	mockUserUseCase := new(MockUserUseCase)
	param := dto.RegisterNewUserDTO{Username: "fikri", Password: "1234qwertyY", Email: "kevin@gmail.com"}
	token := "dummytoken"
	mockUserUseCase.On("RegisterNewUser", &param).Return(token, nil)

	cont := controllers.InitializeUserController(mockUserUseCase)
	cont.RegisterNewUser(c)
	responseData := new(dto.RegisterNewUserResponseDTO)
	decoded, _ := helper_response.DecodeResponseJson(rec.Body.String(), &responseData)

	assert.Equal(t, token, responseData.Token)
	assert.Equal(t, http.StatusOK, int(decoded.Code))
	mockUserUseCase.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	ENDPOINT_PATH := "/user/login"
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, ENDPOINT_PATH, strings.NewReader(`{"username": "fikri", "password": "1234qwertyY"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(ENDPOINT_PATH)

	mockUserUseCase := new(MockUserUseCase)
	param := dto.LoginDTO{Username: "fikri", Password: "1234qwertyY"}
	token := "dummytoken"
	mockUserUseCase.On("Login", &param).Return(token, nil)

	cont := controllers.InitializeUserController(mockUserUseCase)
	cont.Login(c)
	responseData := new(dto.LoginResponseDTO)
	decoded, _ := helper_response.DecodeResponseJson(rec.Body.String(), &responseData)

	assert.Equal(t, token, responseData.Token)
	assert.Equal(t, http.StatusOK, int(decoded.Code))
	mockUserUseCase.AssertExpectations(t)
}
