package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mgw2007/golang-microservices/src/api/clients/restclient"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepo(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		reqBody *strings.Reader
		server  *httptest.Server
	}{
		{
			name:    "Invalid json request",
			status:  http.StatusBadRequest,
			reqBody: strings.NewReader(""),
			server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Write([]byte{})
			})),
		},
		{
			name:    "valid json request BUT response has error",
			status:  http.StatusUnauthorized,
			reqBody: strings.NewReader(`{"name":"validNamxe"}`),
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
			name:    "valid json req",
			status:  http.StatusCreated,
			reqBody: strings.NewReader(`{"name":"validNamxe"}`),
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
			client := server.Client()
			res := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(res)
			req, _ := http.NewRequest(http.MethodGet, "/repositories", tt.reqBody)
			c.Request = req
			CreateRepo(restclient.APIClient{
				Client:  client,
				BaseURL: server.URL,
			}, c)
			assert.EqualValues(t, tt.status, res.Code)

		})
	}
}
