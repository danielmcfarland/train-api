package routes

import (
	"fmt"
	"net/http"
)

func (h *Handler) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Get Station: %s", id)
}
