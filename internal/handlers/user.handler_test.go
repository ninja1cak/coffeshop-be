package handlers

import (
	"net/http"
	"net/http/httptest"
	"ninja1cak/coffeshop-be/internal/models"
	"ninja1cak/coffeshop-be/internal/repositories"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repositories.RepoUserMock{}
var reqBody = `{
		"user_id" : "1",
		"email" : "hauzan41200@gmail.com",
		"phone_number" : "09123",
		"password" : "abcd1234"
	}
`

var reqBodyUpdate = `{
	"user_id" : "1",
	"email" : "hauzan41200@gmail.com",
	"phone_number" : "09123",
	"image" : "google.drive.com"
}
`

func TestPostDataUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := "user success created"
	repoUserMock.On("CreateUser", mock.Anything).Return(expectedResult, nil)
	r.POST("/create-user", handler.PostDataUser)
	req := httptest.NewRequest("POST", "/create-user", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data":"user success created","message":"Created","status":200}`, w.Body.String())

}

func TestGetDataUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := &[]models.User{
		models.User{
			Email:        "hauzan41200@gmail.com",
			Phone_number: "08123741724",
		},
	}

	repoUserMock.On("GetUser", mock.Anything).Return(expectedResult, nil)
	r.GET("/get-user", handler.GetDataUser)
	req := httptest.NewRequest("GET", "/get-user", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"data": [
			{
				"email": "hauzan41200@gmail.com",
				"phone_number": "08123741724"
			}
		],
		"message":"Ok",
		"status":200
		}`, w.Body.String())

}

func TestUpdateDataUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := "update user success"
	repoUserMock.On("UpdateUser", mock.Anything).Return(expectedResult, nil)
	r.PATCH("/update-user", handler.UpdateDataUser)
	req := httptest.NewRequest("PATCH", "/update-user", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data":"update user success", "message":"Ok", "status":201}`, w.Body.String())

}

func TestDeleteDataUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := "Delete user success"
	repoUserMock.On("DeleteUser", mock.Anything).Return(expectedResult, nil)
	r.DELETE("/delete-user", handler.DeleteDataUser)
	req := httptest.NewRequest("DELETE", "/delete-user", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data":"Delete user success", "message":"Ok", "status":201}`, w.Body.String())

}
