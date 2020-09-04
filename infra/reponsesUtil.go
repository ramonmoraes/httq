package infra

import "net/http"

func validResponse(w http.ResponseWriter, content []byte) {
	w.WriteHeader(200)
	w.Write(content)
}

func invalidResponse(w http.ResponseWriter, content []byte) {
	w.WriteHeader(500)
	w.Write(content)
}

func nilResponse(w http.ResponseWriter) {
	msg := "Empty"
	w.WriteHeader(204)
	w.Write([]byte(msg))
}
