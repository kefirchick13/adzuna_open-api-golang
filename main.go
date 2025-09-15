package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Company struct {
	CanonicalName string `json:"canonical_name"`
	Count         int64  `json:"count"`
}

type Category struct {
	Tag   string `json:"tag"`
	Label string `json:"label"`
}

// Определяем тип для страны
type TCountry string

// Определяем допустимые значения через константы
const (
	CountryGB TCountry = "gb"
	CountryUS TCountry = "us"
	CountryFR TCountry = "fr"
	CountryAU TCountry = "au"
)

const (
	clienttime_out = 20 * time.Second
	app_id         = "7b272609"
	app_key        = "36ddf2c24ea04d56ffb97730e2b5a188"
	base_url       = "http://api.adzuna.com/v1/api/"
)

func main() {
	client := &http.Client{
		Timeout: clienttime_out,
	}

	resp, err := getTopCompanies(*client, CountryGB)
	if err != nil {
		log.Fatalf(err.Error())
	}
	topCompanies, err := json.MarshalIndent(resp, " ", " ")
	if err != nil {
		log.Fatal("Не удалось распарсить ответ")
	}
	log.Println(string(topCompanies))

	respCategories, err := getCategories(*client, CountryGB)
	if err != nil {
		log.Fatalf(err.Error())
	}
	categories, err := json.MarshalIndent(respCategories, " ", " ")
	if err != nil {
		log.Fatal("Не удалось распарсить ответ")
	}
	log.Println(string(categories))

}

type TopCompoaniesResponse struct {
	Leaderboard []Company `json:"leaderboard"`
}

func getTopCompanies(client http.Client, country TCountry) ([]Company, error) {
	url_str := base_url + "jobs/" + string(country) + "/top_companies" + "?app_id=" + app_id + "&app_key=" + app_key + "&content-type=application/json"

	resp, err := client.Get(url_str)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var c TopCompoaniesResponse
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}
	return c.Leaderboard, nil
}

type CategoriesResponse struct {
	Results []Category `json:"results"`
}

func getCategories(client http.Client, country TCountry) ([]Category, error) {
	url_str := base_url + "jobs/" + string(country) + "/categories" + "?app_id=" + app_id + "&app_key=" + app_key + "&content-type=application/json"

	resp, err := client.Get(url_str)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var c CategoriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}
	return c.Results, nil
}
