package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/TakeshiHA/test-middleware/models"
	"github.com/TakeshiHA/test-middleware/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetRandomJokes).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetRandomJokes(w http.ResponseWriter, r *http.Request) {

	list := []*models.Joke{}
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	wg.Add(25)
	for i := 0; i < 25; i++ {
		go func() {
			joke, err := utils.GenerateRandomJoke()
			defer wg.Done()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			list = append(list, joke)
		}()
	}
	wg.Wait()
	wg2.Add(1)
	go func() {
		wg2.Done()
		list = utils.CheckingLenght(list)
	}()
	wg2.Wait()

	fmt.Println(len(list))

	json.NewEncoder(w).Encode(list)
	return
}
