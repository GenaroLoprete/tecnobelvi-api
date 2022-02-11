package server

import (
	"encoding/json"
	"net/http"
)

func Ping(responseWriter http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(responseWriter).Encode("pong")
}
