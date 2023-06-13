package api

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiE2e(t *testing.T) {
	type request struct {
		method  string
		path    string
		payload []byte
	}
	type expectedResponse struct {
		expectedStatusCode int
		expectedResponse   string
	}
	tests := []struct {
		name             string
		request          request
		expectedResponse expectedResponse
	}{
		{name: "healthcheck", request: request{method: "GET", path: "/health-check"},
			expectedResponse: expectedResponse{expectedStatusCode: http.StatusOK, expectedResponse: "hello! I'm alive!"}},
		{name: "sendText", request: request{method: "POST", path: "/v1/text", payload: validTextPayload},
			expectedResponse: expectedResponse{expectedStatusCode: http.StatusOK}},
		{name: "sendImage", request: request{method: "POST", path: "/v1/image", payload: validFilenamePayload},
			expectedResponse: expectedResponse{expectedStatusCode: http.StatusOK}},
		{name: "sendAudio", request: request{method: "POST", path: "/v1/audio", payload: validFilenamePayload},
			expectedResponse: expectedResponse{expectedStatusCode: http.StatusOK}},
	}
	okSender := &MockHomeBot{}
	okSender.On("SendText", mock.Anything).Return(nil)
	okSender.On("SendAudio", mock.Anything).Return(nil)
	okSender.On("SendImage", mock.Anything).Return(nil)

	RegisterApiEndpoints(okSender)
	ts := httptest.NewServer(nil)
	defer ts.Close()
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			client := ts.Client()
			var resp *http.Response
			var err error
			url := ts.URL + tt.request.path
			if tt.request.method == "GET" {
				resp, err = client.Get(url)
			} else {
				resp, err = client.Post(url, "application/json", bytes.NewReader(tt.request.payload))
			}

			if err != nil {
				log.Fatal(errors.New("error making request"))
				return
			}

			textResponse, errorReadingBody := io.ReadAll(resp.Body)
			errClosingBody := resp.Body.Close()
			if errorReadingBody != nil || errClosingBody != nil {
				log.Fatal(errors.New("error with body"))
				return
			}

			assert.Equal(t, tt.expectedResponse.expectedResponse, string(textResponse))
			assert.Equal(t, tt.expectedResponse.expectedStatusCode, resp.StatusCode)
		})
	}
}
