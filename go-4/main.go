package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type repoGit struct {
	Name        string `json:"name" `
	Url         string `json:"html_url"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
}

type response struct {
	Name        string `json:"name" `
	Description string `json:"description"`
	Url         string `json:"url"`
	Stars       int    `json:"stars"`
}

type repos struct {
	Items []repoGit `json:"items"`
}

func main() {
	_ = githubStars("go")
}

func githubStars(lang string) error {
	var gitRes repos
	var tenStars []response
	resp, _ := http.Get("https://api.github.com/search/repositories?q=language:" + lang + "&page=1&per_page=10&sort=stars")
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &gitRes)
	for _, item := range gitRes.Items {
		tenStars = append(tenStars, response{
			Name:        item.Name,
			Url:         item.Url,
			Stars:       item.Stars,
			Description: item.Description,
		})
	}
	file, _ := json.MarshalIndent(tenStars, "", " ")
	_ = ioutil.WriteFile("stars.json", file, 0644)
	return nil
}
