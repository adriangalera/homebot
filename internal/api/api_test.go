package api

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"homebot/internal/bot"
	"io"
	"net/http/httptest"
	"testing"
)

var filename = "mock-filename"
var textPayload = "this is a test"
var validFilenamePayload = []byte(`{"filename": "` + filename + `"}`)
var emptyFilenamePayload = []byte(`{"filename": ""}`)
var invalidPayload = []byte(`{`)
var validTextPayload = []byte(`{"text": "` + textPayload + `"}`)
var emptyTextPayload = []byte(`{"text": ""}`)

func TestApi_healthcheck(t *testing.T) {
	api := Api{}
	req := httptest.NewRequest("GET", "http://localhost/health-check", nil)
	rec := httptest.NewRecorder()
	api.healthcheck(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "hello! I'm alive!", string(body))
}

func TestApi_audio(t *testing.T) {
	okSender := &MockHomeBot{}
	okSender.On("SendAudio", filename).Return(nil).Once()
	noopSender := &MockHomeBot{}
	failingSender := &MockHomeBot{}
	failingSender.On("SendAudio", mock.Anything).Return(errors.New("failing bot"))

	tests := []struct {
		name               string
		sender             bot.HomeBot
		payload            []byte
		expectedStatusCode int
	}{
		{name: "Successfully send audio", sender: okSender, payload: validFilenamePayload, expectedStatusCode: 200},
		{name: "Send audio fails parsing invalid JSON", sender: noopSender, payload: invalidPayload, expectedStatusCode: 400},
		{name: "Send audio fails parsing empty filename", sender: noopSender, payload: emptyFilenamePayload, expectedStatusCode: 400},
		{name: "Send audio fails sending", sender: failingSender, payload: validFilenamePayload, expectedStatusCode: 500},
	}
	for _, tt := range tests {
		api := Api{bot: tt.sender}
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://localhost/audio", bytes.NewReader(tt.payload))
			rec := httptest.NewRecorder()
			api.audio(rec, req)
			resp := rec.Result()
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}

func TestApi_image(t *testing.T) {
	okSender := &MockHomeBot{}
	okSender.On("SendImage", filename).Return(nil).Once()
	noopSender := &MockHomeBot{}
	failingSender := &MockHomeBot{}
	failingSender.On("SendImage", mock.Anything).Return(errors.New("failing bot"))

	tests := []struct {
		name               string
		sender             bot.HomeBot
		payload            []byte
		expectedStatusCode int
	}{
		{name: "Successfully send image", sender: okSender, payload: validFilenamePayload, expectedStatusCode: 200},
		{name: "Send image fails parsing invalid JSON", sender: noopSender, payload: invalidPayload, expectedStatusCode: 400},
		{name: "Send image fails parsing empty filename", sender: noopSender, payload: emptyFilenamePayload, expectedStatusCode: 400},
		{name: "Send image fails sending", sender: failingSender, payload: validFilenamePayload, expectedStatusCode: 500},
	}
	for _, tt := range tests {
		api := Api{bot: tt.sender}
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://localhost/image", bytes.NewReader(tt.payload))
			rec := httptest.NewRecorder()
			api.image(rec, req)
			resp := rec.Result()
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}

func TestApi_text(t *testing.T) {
	okSender := &MockHomeBot{}
	okSender.On("SendText", textPayload).Return(nil).Once()
	noopSender := &MockHomeBot{}
	failingSender := &MockHomeBot{}
	failingSender.On("SendText", mock.Anything).Return(errors.New("failing bot"))

	tests := []struct {
		name               string
		sender             bot.HomeBot
		payload            []byte
		expectedStatusCode int
	}{
		{name: "Successfully send text", sender: okSender, payload: validTextPayload, expectedStatusCode: 200},
		{name: "Send text fails parsing invalid JSON", sender: noopSender, payload: invalidPayload, expectedStatusCode: 400},
		{name: "Send text fails parsing empty text", sender: noopSender, payload: emptyTextPayload, expectedStatusCode: 400},
		{name: "Send text fails sending", sender: failingSender, payload: validTextPayload, expectedStatusCode: 500},
	}
	for _, tt := range tests {
		api := Api{bot: tt.sender}
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "http://localhost/text", bytes.NewReader(tt.payload))
			rec := httptest.NewRecorder()
			api.text(rec, req)
			resp := rec.Result()
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}
