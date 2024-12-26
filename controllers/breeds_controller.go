package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	beego.Controller
}

// Helper function to fetch breeds from TheCatAPI
func fetchBreeds(apiKey string) ([]models.Breed, error) {
	url := "https://api.thecatapi.com/v1/breeds"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set API Key
	req.Header.Set("x-api-key", apiKey)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []models.Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

// Helper function to fetch images by breed ID from TheCatAPI
func fetchImagesByBreed(apiKey string, breedID string) ([]models.CatImage, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=8&size=med&sub_id=demo-0.471510602413433444&breed_ids=%s", breedID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set API Key
	req.Header.Set("x-api-key", apiKey)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var images []models.CatImage
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (c *BreedsController) FetchBreedsAndImages() ([]models.Breed, map[string][]models.CatImage, error) {
	apiKey, _ := beego.AppConfig.String("api_key")

	// Fetch breeds
	breeds, err := fetchBreeds(apiKey)
	if err != nil {
		return nil, nil, err
	}

	// Create a map to store images for each breed
	breedImages := make(map[string][]models.CatImage)

	// Mutex to protect concurrent access to breedImages
	var mu sync.Mutex

	// Channel to collect errors from goroutines
	errCh := make(chan error, len(breeds))

	// Wait group to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Fetch images for each breed concurrently
	for _, breed := range breeds {
		wg.Add(1)
		go func(breed models.Breed) {
			defer wg.Done()

			// Fetch images for the current breed
			images, err := fetchImagesByBreed(apiKey, breed.ID)
			if err != nil {
				errCh <- err
				return
			}

			// Store the images in the map with mutex lock
			mu.Lock()
			breedImages[breed.ID] = images
			mu.Unlock()
		}(breed)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check for any errors during goroutine execution
	select {
	case err := <-errCh:
		return nil, nil, err
	default:
		// No errors, return the breeds and their images
		return breeds, breedImages, nil
	}
}
