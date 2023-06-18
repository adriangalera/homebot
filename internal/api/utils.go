package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func parseFileNameRequest(r *http.Request) (inputFilename, error) {
	var fileNameRequest inputFilename

	errorParsingBody := json.NewDecoder(r.Body).Decode(&fileNameRequest)
	if errorParsingBody != nil {
		return inputFilename{}, errorParsingBody
	}

	if fileNameRequest.Filename == "" {
		return inputFilename{}, errors.New("invalid payload")
	}
	return fileNameRequest, nil
}
