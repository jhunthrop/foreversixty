package main

import (
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

func TestNewDiscordBotStaysOffWithoutConfig(t *testing.T) {
	if bot := newDiscordBot(config.Config{}, nil, server.Deps{}, slog.Default()); bot != nil {
		t.Fatal("bot built without its variables")
	}
}
