package tests

import (
	"net/http"
	"testing"
	"strings"

	"bee_project/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"fmt"
)

// Mock HTTP Client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// Test for APIController FetchImages
func TestAPIController_FetchImages(t *testing.T) {
	// Mock the HTTP client to simulate an external API call
	mockClient := new(MockHTTPClient)

	// Set the mock client in the APIController
	apiController := &controllers.APIController{}
	// Mock API Key and URL configuration
	api_key := "your api key"
	beego.AppConfig.Set("api_key", api_key)
	api_url := "https://api.thecatapi.com/v1/images/search"
	beego.AppConfig.Set("get_images_api_url", api_url)

	// Call the method to fetch images
	images, err := apiController.FetchImages()


	// Extract the URL up to the last slash
    expectedURL := "https://cdn2.thecatapi.com/images"
    imageURL := images[0].URL
    lastSlashIndex := strings.LastIndex(imageURL, "/")
    if lastSlashIndex != -1 {
        imageURL = imageURL[:lastSlashIndex]
    }
    // Assert that the method works as expected, only checking up to the last slash
    assert.Equal(t, err, nil)
    assert.Equal(t, expectedURL, imageURL)

	// Check mock expectations
	mockClient.AssertExpectations(t)
}
