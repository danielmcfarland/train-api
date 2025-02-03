package routes

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) handleOther(w http.ResponseWriter, r *http.Request) {
	log.Println("Received an unresolvable request")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
