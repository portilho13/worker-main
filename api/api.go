package api

import (
	"encoding/json"
	"net/http"
)

func Api(ip string, SERVERS_TCP_IP *[]string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /servers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		if err := json.NewEncoder(w).Encode(SERVERS_TCP_IP); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("POST /servers/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		*SERVERS_TCP_IP = append(*SERVERS_TCP_IP, id)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Sucessfully Added"))
	})

	if err := http.ListenAndServe(ip, mux); err != nil {
		return err
	}

	return nil
}