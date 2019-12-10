package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken) //169d7a3fa8cc9c9c6bd8e6bef1f8db0a6378d642
)

//GetGithubAccessToken return token
func GetGithubAccessToken() string {
	return githubAccessToken
}

//GetGithubRepoURL return url for repo rest api
func GetGithubRepoURL() string {
	return "https://api.github.com/user/repos"
}
