package discordbot

import (
	"context"
	"errors"
	"strings"
	"testing"
)

const registerGuild = "srvtest"

func putCalls(r *FakeREST) []RESTCall {
	var out []RESTCall
	for _, c := range r.Calls() {
		if c.Method == "PutCommands" {
			out = append(out, c)
		}
	}
	return out
}

func registerWorld(t *testing.T) *pressWorld {
	t.Helper()
	w := newPressWorld(t)
	w.svc.d.ApplicationID = "app-" + t.Name()
	w.svc.d.CommandsGuildID = registerGuild
	w.exec(t, `delete from discord_command_sets`)
	return w
}

func TestRegisterOnlyWhenChanged(t *testing.T) {
	w := registerWorld(t)
	ctx := context.Background()
	if err := w.svc.RegisterCommands(ctx); err != nil {
		t.Fatal(err)
	}
	calls := putCalls(w.rest)
	if len(calls) != 2 || calls[0].ServerID != "" || calls[1].ServerID != registerGuild {
		t.Fatalf("first run calls = %+v", calls)
	}
	if len(calls[0].Commands) != 4 || len(calls[1].Commands) != 5 {
		t.Fatalf("released %d, unreleased %d", len(calls[0].Commands), len(calls[1].Commands))
	}
	if err := w.svc.RegisterCommands(ctx); err != nil || len(putCalls(w.rest)) != 2 {
		t.Fatalf("second run: %v, %d calls", err, len(putCalls(w.rest)))
	}
	w.svc.commandSet = func() []Command {
		cmds := Commands()
		cmds[0].Description = "Changed."
		return cmds
	}
	if err := w.svc.RegisterCommands(ctx); err != nil {
		t.Fatal(err)
	}
	calls = putCalls(w.rest)
	if len(calls) != 3 || calls[2].ServerID != "" {
		t.Fatalf("after an edit: %+v", calls)
	}
}

func TestRegisterSkipsTheDevelopmentServerWhenUnset(t *testing.T) {
	w := registerWorld(t)
	w.svc.d.CommandsGuildID = ""
	if err := w.svc.RegisterCommands(context.Background()); err != nil {
		t.Fatal(err)
	}
	if calls := putCalls(w.rest); len(calls) != 1 || calls[0].ServerID != "" {
		t.Fatalf("calls = %+v", calls)
	}
}

func TestRegisterFailureDoesNotStopStartup(t *testing.T) {
	w := registerWorld(t)
	boom := errors.New("discord says no")
	w.rest.ErrFor = map[string]error{"PutCommands": boom}
	err := w.svc.RegisterCommands(context.Background())
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v", err)
	}
	var n int
	if e := w.st.Pool.QueryRow(context.Background(), `select count(*) from discord_command_sets where application_id = $1`,
		w.svc.d.ApplicationID).Scan(&n); e != nil || n != 0 {
		t.Fatalf("hash rows = %d, %v", n, e)
	}
	w.rest.ErrFor = nil
	if err := w.svc.RegisterCommands(context.Background()); err != nil || len(putCalls(w.rest)) != 4 {
		t.Fatalf("retry: %v, %d calls", err, len(putCalls(w.rest)))
	}
}

func TestCommandsHashIsStableAndSensitive(t *testing.T) {
	a, b := commandsHash(Commands()), commandsHash(Commands())
	cmds := Commands()
	cmds[0].Description += "!"
	if a != b || a == commandsHash(cmds) || len(a) != 64 || strings.Trim(a, "0123456789abcdef") != "" {
		t.Fatalf("hashes %s %s", a, b)
	}
}
