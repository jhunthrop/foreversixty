package main

import (
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

func TestNewDiscordBotStaysOffWithoutConfig(t *testing.T) {
	accounts := &auth.Service{}
	if bot := newDiscordBot(config.Config{}, nil, server.Deps{}, accounts, slog.Default()); bot != nil {
		t.Fatal("bot built without its variables")
	}
	if accounts.DiscordLinks != nil {
		t.Fatal("a bot that is off still answers the account's link fields")
	}
}
