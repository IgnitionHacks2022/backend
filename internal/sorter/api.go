package sorter

import (
	"fmt"
	"net/http"
)

func ClassifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Classify")
}
