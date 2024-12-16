package healthcheck

import (
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
