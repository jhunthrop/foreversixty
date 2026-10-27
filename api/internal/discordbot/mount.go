package discordbot

import "net/http"

// InteractionsPath is the one endpoint Discord posts interactions to.
const InteractionsPath = "/v1/discord/interactions"

// Mount registers the interactions endpoint, which carries no session (Discord's signature
// is its only credential), and the account link routes, which do.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("POST "+InteractionsPath, s.serveInteraction)
	s.mountLink(mux, trustedProxyHops)
}
