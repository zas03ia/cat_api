package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)

type VoteController struct {
	beego.Controller
}



// Handle voting on images (Up or Down)
func (c *VoteController) Vote() {
	// Get the API key from the config
	apiKey, _ := beego.AppConfig.String("api_key")

	// Retrieve parameters
	imageID := c.GetString("image_id")
	value, err := strconv.Atoi(c.GetString("vote"))
	if err != nil {
		c.Data["json"] = &models.VoteAPIResponse{Message: "Invalid vote value"}
		c.ServeJSON()
		return
	}

	// Create the vote request payload
	voteRequest := &models.VoteRequest{
		ImageID: imageID,
		SubID:   "user123",
		Value:   value,
	}

	// Serialize the payload to JSON
	data, err := json.Marshal(voteRequest)
	if err != nil {
		c.Data["json"] = &models.VoteAPIResponse{Message: "Failed to serialize request payload"}
		c.ServeJSON()
		return
	}

	// Prepare the HTTP client and request
	client := &http.Client{}
	url := "https://api.thecatapi.com/v1/votes"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		c.Data["json"] = &models.VoteAPIResponse{Message: "Failed to create HTTP request"}
		c.ServeJSON()
		return
	}

	// Set the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = &models.VoteAPIResponse{Message: "Failed to send HTTP request"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Parse the response
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.Data["json"] = &models.VoteAPIResponse{Message: fmt.Sprintf("Failed to register vote: %s", string(body))}
		c.ServeJSON()
		return
	}

	c.Redirect("/", http.StatusFound)
}
