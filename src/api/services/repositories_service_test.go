package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"

	"github.com/mgw2007/golang-microservices/src/api/domain/repositories"
	"github.com/mgw2007/golang-microservices/src/api/utils/errors"
)

func Test_repoService_CreateRepo(t *testing.T) {

	tests := []struct {
		name   string
		input  repositories.CreateRepoRequest
		want   *repositories.CreateRepoResponse
		want1  errors.ApiError
		server *httptest.Server
	}{
		{
			name: "Test Invalid input name",
			input: repositories.CreateRepoRequest{
				Name:        "",
				Description: "desc",
			},
			want:  nil,
			want1: errors.NewBadRequestError("Invalid Repository Name"),
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Write([]byte{})
			})),
		},
		{
			name: "Test Invalid Token Create Repo Return valid json Error from repo Invalid Token ",
			input: repositories.CreateRepoRequest{
				Name:        "validName",
				Description: "validDesc",
			},
			want:  nil,
			want1: errors.NewAPIError(http.StatusUnauthorized, "Bad credentials"),
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Header().Set("Content-Type", "application/json")
				json.NewEncoder(rw).Encode(struct {
					StatusCode int
					Message    string
				}{
					StatusCode: http.StatusUnauthorized,
					Message:    "Bad credentials",
				})
			})),
		},
		{
			name: "Test valid Token Create Repo Return valid json from repo ",
			input: repositories.CreateRepoRequest{
				Name:        "validName",
				Description: "validDesc",
			},
			want: &repositories.CreateRepoResponse{
				ID:    123,
				Name:  "validName",
				Owner: "OwnerId",
			},
			want1: nil,
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.WriteHeader(http.StatusCreated)
				rw.Header().Set("Content-Type", "application/json")
				json.NewEncoder(rw).Encode(struct {
					ID       int64  `json:"id"`
					Name     string `json:"name"`
					FullName string `json:"full_name"`
					Owner    struct {
						ID    int64  `json:"id"`
						Login string `json:"login"`
					} `json:"owner"`
				}{
					ID:       123,
					Name:     "validName",
					FullName: "valid Name",
					Owner: struct {
						ID    int64  `json:"id"`
						Login string `json:"login"`
					}{
						ID:    123,
						Login: "OwnerId",
					},
				})
			})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.server
			// Close the server when test finishes
			defer server.Close()
			c := server.Client()
			got, got1 := RepositoryService.CreateRepo(restclient.APIClient{
				Client:  c,
				BaseURL: server.URL,
			}, tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repoService.CreateRepo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("repoService.CreateRepo() got1 = %v, want %v", got1.Message(), tt.want1.Message())
			}
		})
	}
}
