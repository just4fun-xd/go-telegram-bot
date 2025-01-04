package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DogImageResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetRandomDogImage() (string, error) {
	url := "https://dog.ceo/api/breeds/image/random"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("не удалось получить изображение собаки: %s", resp.Status)
	}

	var dogResp DogImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&dogResp); err != nil {
		return "", err
	}

	if dogResp.Status != "success" {
		return "", fmt.Errorf("API вернул ошибку: %s", dogResp.Status)
	}

	return dogResp.Message, nil
}

func GetDogImageByBreed(breed string) (string, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("не удалось получить изображение собаки: %s", resp.Status)
	}

	var dogResp DogImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&dogResp); err != nil {
		return "", err
	}

	if dogResp.Status != "success" {
		return "", fmt.Errorf("API вернул ошибку: %s", dogResp.Status)
	}
	return dogResp.Message, nil
}
