package controllers


import (
	"encoding/json"
	"sync"
	"log"
	"time"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)



type AggregateController struct {
	beego.Controller
	apiController APIController
	breedsController BreedsController
	favouritesController ShowFavouritesController
}

func (c *AggregateController) Index() {
	// Initialize controllers
	apiController := &APIController{}
	breedsController := &BreedsController{}
	favouritesController := &ShowFavouritesController{}

	// Fetch data concurrently
	var (
		images     []models.APIResponse
		breeds     []models.Breed
		breedImages map[string][]models.CatImage
		favourites []models.Favourite
		errs       []error
	)
	
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Fetch images
	go func() {
		defer wg.Done()
		var err error
		images, err = apiController.FetchImages()
		if err != nil {
			errs = append(errs, err)
		}
	}()

	// Fetch breeds and images
	go func() {
		defer wg.Done()
		
		// Define retry parameters (Sometimes cat api does not respond properly)
		retryCount := 3 
		retryDelay := 2 * time.Second 
		var err error
	
		// Attempt to fetch breeds and breed images with retries
		for attempt := 1; attempt <= retryCount; attempt++ {
			breeds, breedImages, err = breedsController.FetchBreedsAndImages()
			if err == nil {
				// Success, break out of the loop
				return
			}
	
			// If error occurs, log it and retry
			log.Printf("Error fetching breeds and images (attempt %d/%d): %v", attempt, retryCount, err)
	
			if attempt < retryCount {
				// Delay before retrying
				time.Sleep(retryDelay)
			}
		}
	
		// After retrying, if the error persists, append to the error slice
		if err != nil {
			errs = append(errs, err)
		}
	}()
	

	// Fetch favourites
	go func() {
		defer wg.Done()
		var err error
		favourites, err = favouritesController.FetchFavourites()
		if err != nil {
			errs = append(errs, err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Handle errors
	if len(errs) > 0 {
		c.Data["json"] = map[string]interface{}{"message": "Failed to fetch data", "errors": errs}
		c.ServeJSON()
		return
	}

	breedsJSON, _ := json.Marshal(breeds)
	breedImagesJSON, _ := json.Marshal(breedImages)

	// Pass data to the template
	c.Data["imageURL"] = images[0].URL
	c.Data["imageID"] = images[0].ID
	c.Data["breeds"] = breeds
	c.Data["Breeds_json"] = string(breedsJSON)
	c.Data["breedImages"] = string(breedImagesJSON)
	c.Data["favourites"] = favourites
	c.TplName = "index.tpl"
}
