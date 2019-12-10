package github_provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"
	"github.com/mgw2007/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func Test_getAuthorizationHeader(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test token", args: args{"1233333"}, want: "Token 1233333"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAuthorizationHeader(tt.args.token); got != tt.want {
				t.Errorf("getAuthorizationHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRepoFailedToConnect(t *testing.T) {
	//test request failed connect by make client time out < the server processing time out
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Microsecond * 50) // make server wait for 50 Microsecond more than the client need
		rw.WriteHeader(http.StatusGatewayTimeout)
	}))
	// Close the server when test finishes
	defer server.Close()
	c := server.Client()
	c.Timeout = time.Microsecond // cut connection after 1 Microsecond before server finishd process
	_, err := CreateRepo(restclient.APIClient{
		Client:  c,
		BaseURL: server.URL,
	}, "", github.CreateRepoRequest{})
	assert.Error(t, err)
}
func TestCreateRepo(t *testing.T) {
	type args struct {
		token string
		req   github.CreateRepoRequest
	}
	tests := []struct {
		name   string
		args   args
		want   *github.CreateRepoResponse
		want1  *github.ErrorResponse
		server *httptest.Server
	}{
		{
			name: "Test Create Repo Invalid response body",
			args: args{
				token: "",
				req: github.CreateRepoRequest{
					Name: "valid request",
				},
			},
			want: nil,
			want1: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid response body",
			},
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Header().Set("Content-Length", "1")
			})),
		},
		{
			name: "Test Create Repo Invalid request body",
			args: args{
				token: "",
				req: github.CreateRepoRequest{
					Name: "valid request",
				},
			},
			want: nil,
			want1: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response error body",
			},
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.WriteHeader(http.StatusNotFound)
				rw.Write([]byte(`not_valid_data`))
			})),
		},
		{
			name: "Test Create Repo Return valid json Error from repo Invalid Token ",
			args: args{
				token: "invalid token",
				req: github.CreateRepoRequest{
					Name: "valid request",
				},
			},
			want: nil,
			want1: &github.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Bad credentials",
			},
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
			name: "Test Create Repo in Valid response body",
			args: args{
				token: "",
				req: github.CreateRepoRequest{
					Name: "valid request",
				},
			},
			want: nil,
			want1: &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response success body",
			},
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.WriteHeader(http.StatusAccepted)
				rw.Write([]byte(`not_valid_data`))
			})),
		},
		{
			name: "Test Create Repo SUCCESS Valid Request Valid response body",
			args: args{
				token: "valid token",
				req: github.CreateRepoRequest{
					Name: "valid request",
				},
			},
			want: &github.CreateRepoResponse{
				ID:   int64(123),
				Name: "our_test_name",
			},
			want1: nil,
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.WriteHeader(http.StatusCreated)
				json.NewEncoder(rw).Encode(struct {
					Name string
					ID   int64
				}{
					ID:   int64(123),
					Name: "our_test_name",
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
			got, got1 := CreateRepo(restclient.APIClient{
				Client:  c,
				BaseURL: server.URL,
			}, tt.args.token, tt.args.req)
			if tt.want != nil {
				assert.Equal(t, tt.want, got)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRepo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateRepo() got1 = %v, want %v", got1, tt.want1)
			}

		})
	}
}
