package api

import (
	"encoding/json"
	"fmt"
	"homebot/internal/bot"
	"net/http"
)

type Api struct {
	bot bot.HomeBot
}

type textMessage struct {
	Text string `json:"text"`
}

type inputFilename struct {
	Filename string `json:"filename"`
}

// GET /health-check
func (a *Api) healthcheck(w http.ResponseWriter, _ *http.Request) {
	responseString := "hello! I'm alive!"
	_, _ = fmt.Fprint(w, responseString)
}

// POST /v1/text
func (a *Api) text(w http.ResponseWriter, r *http.Request) {
	var newTextMessage textMessage

	errorParsingBody := json.NewDecoder(r.Body).Decode(&newTextMessage)
	if errorParsingBody != nil {
		http.Error(w, errorParsingBody.Error(), http.StatusBadRequest)
		return
	}

	if newTextMessage.Text == "" {
		http.Error(w, "empty text", http.StatusBadRequest)
		return
	}

	errorSendingMessage := a.bot.SendText(newTextMessage.Text)
	if errorSendingMessage != nil {
		http.Error(w, errorSendingMessage.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /v1/image
func (a *Api) image(w http.ResponseWriter, r *http.Request) {

	fileNameRequest, errorParsing := parseFileNameRequest(r)
	if errorParsing != nil {
		http.Error(w, errorParsing.Error(), http.StatusBadRequest)
		return
	}

	errorSending := a.bot.SendImage(fileNameRequest.Filename)
	if errorSending != nil {
		http.Error(w, errorSending.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /v1/audio
func (a *Api) audio(w http.ResponseWriter, r *http.Request) {
	fileNameRequest, errorParsing := parseFileNameRequest(r)
	if errorParsing != nil {
		http.Error(w, errorParsing.Error(), http.StatusBadRequest)
		return
	}

	errorSending := a.bot.SendAudio(fileNameRequest.Filename)
	if errorSending != nil {
		http.Error(w, errorSending.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterApiEndpoints(bot bot.HomeBot) {
	api := Api{bot: bot}
	http.HandleFunc("/health-check", api.healthcheck)
	http.HandleFunc("/v1/text", api.text)
	http.HandleFunc("/v1/image", api.image)
	http.HandleFunc("/v1/audio", api.audio)
}
