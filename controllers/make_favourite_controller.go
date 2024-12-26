package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)

type MakeFavouriteController struct {
	beego.Controller
}



// Handle adding an image to favourites
func (c *MakeFavouriteController) Favourite() {
	// Get the API key from the config
	apiKey, _ := beego.AppConfig.String("api_key")

	// Get the image ID from the request
	imageID := c.GetString("image_id")
	if imageID == "" || imageID == " " {
		c.Data["json"] = map[string]string{"message": "Image ID is required"}
		c.ServeJSON()
		return
	}

	// Prepare the request payload
	favRequest := &models.FavouriteRequest{
		ImageID: imageID,
	}
	data, err := json.Marshal(favRequest)
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Failed to marshal data"}
		c.ServeJSON()
		return
	}

	// Send the POST request to The Cat API
	url := "https://api.thecatapi.com/v1/favourites"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Failed to create request"}
		c.ServeJSON()
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Failed to send request"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Data["json"] = map[string]string{"message": "Failed to add favourite"}
		c.ServeJSON()
		return
	}

	// Parse and log the response
	var favResponse models.FavouriteResponse
	err = json.Unmarshal(body, &favResponse)
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Failed to parse response"}
		c.ServeJSON()
		return
	}

	// Redirect to the root route after success
	c.Redirect("/", http.StatusFound)
}
