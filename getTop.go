package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name            string   `json:"name"`
	StargazersCount int      `json:"stargazers_count"`
	Language        string   `json:"language"`
	Topics          []string `json:"topics"`
	Description     string   `json:"description"`
}

func searchRepositories(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=ROS+user:%s&sort=stars&order=desc", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repositories: %s", resp.Status)
	}

	var result struct {
		Items []Repository `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

func main() {
	username := "ARTI-Robots" // Замените на нужное имя пользователя GitHub
	repos, err := searchRepositories(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("Repo: %s\nStars: %d\nLanguage: %s\nTopics: %v\nDescription: %s\n\n",
			repo.Name, repo.StargazersCount, repo.Language, repo.Topics, repo.Description)
	}
}
