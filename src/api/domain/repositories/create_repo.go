package repositories

//CreateRepoRequest for used in CreateRepo
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//CreateRepoResponse for return from  CreateRepo
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
