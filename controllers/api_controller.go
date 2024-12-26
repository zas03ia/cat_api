package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) FetchImages() ([]models.APIResponse, error) {
	apiKey, _ := beego.AppConfig.String("api_key")
	apiURL, _ := beego.AppConfig.String("get_images_api_url")

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []models.APIResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
