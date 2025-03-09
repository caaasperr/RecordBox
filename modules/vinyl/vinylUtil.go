package vinyl

import (
	"encoding/json"
	"errors"
	"fmt"
	"myvinyl/models"
	"myvinyl/modules/vinyl/controller"
	"myvinyl/utils"
	"net/http"
	"net/url"
)

type AlbumSearchResponse struct {
	Results struct {
		Albummatches struct {
			Album []struct {
				Name   string `json:"name"`
				Artist string `json:"artist"`
				Image  []struct {
					Text string `json:"#text"`
					Size string `json:"size"`
				} `json:"image"`
			} `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

type AlbumResponse struct {
	Album struct {
		Name  string `json:"name"`
		Image []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
	} `json:"album"`
}

func IsOwned(user models.User, vinylId uint) error {
	vinyl, err := controller.GetVinylByID(vinylId)
	if err != nil {
		return errors.New("check err")
	}
	if user.ID == vinyl.UserID {
		return nil
	} else {
		return errors.New("not owned vinyl")
	}
}

func GetAlbumCoverFromSpecificInformation(name string, artist string) (AlbumResponse, error) {
	apiURL := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=%s&album=%s&artist=%s&format=json",
		utils.LASTFM_API_KEY,
		url.QueryEscape(name),
		url.QueryEscape(artist),
	)
	resp, err := http.Get(apiURL)
	if err != nil {
		utils.Logger.Error("Failed to fetch data from Last.fm API", http.StatusInternalServerError)
		return AlbumResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		utils.Logger.Error("Last.fm API returned an error")
		return AlbumResponse{}, errors.New("lastfm returned an error")
	}
	var albumResponse AlbumResponse
	if err := json.NewDecoder(resp.Body).Decode(&albumResponse); err != nil {
		utils.Logger.Error("Failed to decode API response", http.StatusInternalServerError)
		return AlbumResponse{}, errors.New("failed to decode api response")
	}
	return albumResponse, nil
}

func GetAlbumCoversFromLastFmByName(name string) (AlbumSearchResponse, error) {
	apiURL := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=album.search&api_key=%s&album=%s&format=json",
		utils.LASTFM_API_KEY,
		url.QueryEscape(name),
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		utils.Logger.Error("Failed to fetch data from Last.fm API", http.StatusInternalServerError)
		return AlbumSearchResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		utils.Logger.Error("Last.fm API returned an error")
		return AlbumSearchResponse{}, errors.New("lastfm returned an error")
	}

	var albumResponse AlbumSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&albumResponse); err != nil {
		utils.Logger.Error("Failed to decode API response", http.StatusInternalServerError)
		return AlbumSearchResponse{}, errors.New("failed to decode api response")
	}
	return albumResponse, nil
}
