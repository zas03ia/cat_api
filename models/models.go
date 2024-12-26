package models

type APIResponse struct {
	URL     string `json:"url"`
	ID      string `json:"id"`
	Message string ""
}



type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Origin string `json:"origin"`
	Description string `json:"description"`
	Wikipedia_url string `json:"wikipedia_url"`
}

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}


type FavouriteRequest struct {
	ImageID string `json:"image_id"`
}

type FavouriteResponse struct {
	ID int `json:"id"`
}


type Favourite struct {
	ID        int    `json:"id"`
	ImageID   string `json:"image_id"`
	SubID     string `json:"sub_id"`
	CreatedAt string `json:"created_at"`
	Image     struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}


type VoteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id,omitempty"`
	Value   int    `json:"value"`
}

type VoteAPIResponse struct {
	Message string `json:"message"`
}

type APIInterface interface {
	FetchImages() ([]APIResponse, error)
}

type BreedsInterface interface {
	FetchBreedsAndImages() ([]Breed, map[string][]CatImage, error)
}

type FavouritesInterface interface {
	FetchFavourites() ([]Favourite, error)
}