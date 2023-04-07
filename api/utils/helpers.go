package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gevs90/crypto-api/api/models"
)

func NewKey() (string, error) {
	url := "https://random-data-api.com/api/color/random_color"

	client := &http.Client{Timeout: 20 * time.Second}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating request", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error makin request", err)
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("errro response status code", response.StatusCode)
		return "", err
	}

	var responseData models.ResponseRandomData
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		fmt.Println("error decoding response", err)
		return "", err
	}

	return responseData.Uid, nil
}
