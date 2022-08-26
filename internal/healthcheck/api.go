package healthcheck

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health Check!")
	w.WriteHeader(http.StatusOK)
}
