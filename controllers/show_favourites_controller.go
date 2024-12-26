package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bee_project/models"
	beego "github.com/beego/beego/v2/server/web"
)

type ShowFavouritesController struct {
	beego.Controller
}



func (c *ShowFavouritesController) FetchFavourites() ([]models.Favourite, error) {
	apiKey, _ := beego.AppConfig.String("api_key")
	url := "https://api.thecatapi.com/v1/favourites"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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

	var favourites []models.Favourite
	err = json.Unmarshal(body, &favourites)
	if err != nil {
		return nil, err
	}

	return favourites, nil
}
