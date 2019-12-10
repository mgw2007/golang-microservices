package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

//GetGithubAccessToken return token
func GetGithubAccessToken() string {
	return githubAccessToken
}

//GetGithubRepoURL return url for repo rest api
func GetGithubRepoURL() string {
	return "https://api.github.com/user/repos"
}
