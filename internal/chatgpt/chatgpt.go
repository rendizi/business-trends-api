package chatgpt

import (
	"bta/internal/db"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Created int64 `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func Send(user db.User, apiKey string) (string, error) {
	url := "https://api.openai.com/v1/images/generations"

	data := map[string]interface{}{
		"model": "dall-e-3",
		"prompt": fmt.Sprintf("Busines in %s direction. Description: %s. It is located in %s. It's revenue in last 1 year in format 1;2;3;6;12 month ago: %s. Make a prediction for it's income in next few month and build chart",
			user.Direction, user.Description, user.City, user.Revenue),
		"n":    1,
		"size": "1024x1024",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return "", err

	}

	if len(responseData.Data) > 0 {
		return responseData.Data[0].URL, nil
	}
	return "", errors.New("no url")

}
