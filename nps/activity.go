package nps

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Activities struct {
	Total int `json:"total,string"`
	Data  []Category
	Limit int `json:"Limit,string"`
	Start int `json:"Start,string"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Parks struct {
	Total int `json:"total,string"`
	Data  []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Parks []Park `json:"parks"`
	}
	Limit int `json:"Limit,string"`
	Start int `json:"Start,string"`
}

type Park struct {
	States      string `json:"states"`
	FullName    string `json:"fullName"`
	Url         string `json:"url"`
	ParkCode    string `json:"parkCode"`
	Designation string `json:"designation"`
	Name        string `json:"name"`
}

func GetCategories(apikey string) (Activities, error) {
	req, err := http.NewRequest("GET", "https://developer.nps.gov/api/v1/activities", nil)

	if err != nil {
		return Activities{}, errors.New("Failed to get list of available activity categories.")
	}

	req.Header.Set("x-api-key", apikey)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Print("error making request", err)
		return Activities{}, errors.New("Error making web request")
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	a := Activities{}

	if err := decoder.Decode(&a); err != nil {
		fmt.Print("Error decoding response body: ", err)
	}

	return a, nil
}

func GetParkByCategory(a string, apikey string) (Parks, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://developer.nps.gov/api/v1/activities/parks?id=%s", a),
		nil)

	if err != nil {
		return Parks{}, fmt.Errorf("Failed to create request: %w", err)
	}

	req.Header.Set("x-api-key", apikey)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return Parks{}, fmt.Errorf("HTTP ERROR:", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	p := Parks{}

	if err := decoder.Decode(&p); err != nil {
		return Parks{}, fmt.Errorf("Json Decode Error: %w", err)
	}

	return p, nil
}
