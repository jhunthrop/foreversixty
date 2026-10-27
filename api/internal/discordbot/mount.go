package discordbot

import "net/http"

// InteractionsPath is the one endpoint Discord posts interactions to.
const InteractionsPath = "/v1/discord/interactions"

// Mount registers the interactions endpoint. It carries no session: Discord's
// signature is its only credential.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("POST "+InteractionsPath, s.serveInteraction)
}
