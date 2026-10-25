package notify

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteSendsDELETEAndTreats2xxAnd404AsSuccess(t *testing.T) {
	for status, wantErr := range map[int]bool{204: false, 200: false, 404: false, 401: true, 500: true} {
		var method string
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method = r.Method
			w.WriteHeader(status)
		}))
		err := NewHTTPPoster("").Delete(context.Background(), srv.URL)
		srv.Close()
		if method != http.MethodDelete || (err != nil) != wantErr {
			t.Fatalf("status %d: method %s, err %v", status, method, err)
		}
		var discord *DiscordError
		if wantErr && !errors.As(err, &discord) {
			t.Fatalf("status %d: err %v is not a *DiscordError", status, err)
		}
	}
}

func TestLookupReadsTheGuildID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"channel_id":"5","guild_id":"6","name":"n"}`))
	}))
	defer srv.Close()
	got, err := NewHTTPPoster("").Lookup(context.Background(), srv.URL)
	if err != nil || got.GuildID != "6" || got.ChannelID != "5" {
		t.Fatalf("lookup = %+v, %v", got, err)
	}
}
