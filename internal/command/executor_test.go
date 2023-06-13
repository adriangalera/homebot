package command

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	commandExecutor := Executor{}

	ok := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, "")
	}
	fail := func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "force error", http.StatusBadRequest)
	}

	http.HandleFunc("/ok", ok)
	http.HandleFunc("/fail", fail)

	ts := httptest.NewServer(nil)

	port, _ := strconv.Atoi(strings.Split(ts.URL, ":")[2])

	defer ts.Close()

	type args struct {
		command Command
	}
	type expects struct {
		error    bool
		response string
	}
	okCommand := remoteCommand{Command: "/ok", Path: "/ok", Port: port}
	failCommand := remoteCommand{Command: "/fail", Path: "/fail", Port: port}
	notFoundCommand := remoteCommand{Command: "/notfound", Path: "/notfound", Port: port}
	noConnectionCommand := remoteCommand{Command: "/notfound", Path: "/notfound", Port: 10000}
	local := &HelpCommand{wrapper: nil}

	tests := []struct {
		name    string
		args    args
		expects expects
	}{
		{name: "Server 200 OK", args: args{command: okCommand}, expects: expects{error: false}},
		{name: "Server 500", args: args{command: failCommand}, expects: expects{error: true}},
		{name: "Server Not found", args: args{command: notFoundCommand}, expects: expects{error: true}},
		{name: "No connection", args: args{command: noConnectionCommand}, expects: expects{error: true}},
		{name: "local command", args: args{command: local}, expects: expects{error: false, response: local.Execute()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, errExecutingCommand := commandExecutor.Execute(tt.args.command)
			if tt.expects.error {
				assert.NotNil(t, errExecutingCommand)
			} else {
				assert.Nil(t, errExecutingCommand)
			}
			assert.Equal(t, tt.expects.response, response)
		})
	}
}
