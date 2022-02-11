package server

import (
	"encoding/json"
	"net/http"

	"tecnobelvi-api/info"
)

func GetBuilds(responseWriter http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(responseWriter).Encode(info.GetBuildsInfo())
}
