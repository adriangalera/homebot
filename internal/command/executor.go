package command

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Executor struct {
}

func (e *Executor) Execute(command Command) (string, error) {
	switch command.(type) {
	case remoteCommand:
		statusCode := e.post(e.buildUrl(command.(remoteCommand)))
		if statusCode != http.StatusOK {
			return "", errors.New(fmt.Sprintf("cannot execute command %s", command))
		}
	case localCommand:
		return command.(localCommand).Execute(), nil
	}
	return "", nil
}

func (e *Executor) buildUrl(command remoteCommand) string {
	return fmt.Sprintf("http://localhost:%d%s", command.Port, command.Path)
}

func (e *Executor) post(url string) int {
	response, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Cannot do post request to %s. Error: %s", url, err)
		return -1
	}
	return response.StatusCode
}
