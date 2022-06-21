package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TakeshiHA/test-middleware/models"
)

func GenerateRandomJoke() (*models.Joke, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.chucknorris.io/jokes/random", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var randomJoke models.Joke
	err = json.Unmarshal(bodyBytes, &randomJoke)
	if err != nil {
		return nil, err
	}
	return &randomJoke, nil
}

func CheckingLenght(list []*models.Joke) []*models.Joke {
	for {
		fmt.Println(len(list))
		list = ValidateList(list)
		if len(list) < 25 {
			joke, _ := GenerateRandomJoke()
			list = append(list, joke)
		} else {
			break
		}
	}
	return list
}

func ValidateList(list []*models.Joke) []*models.Joke {
	keys := make(map[string]*models.Joke)
	jokeList := []*models.Joke{}
	for _, entry := range list {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = entry
			jokeList = append(jokeList, entry)
		}
	}
	return jokeList
}
