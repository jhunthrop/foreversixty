package discordbot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// pgUniqueViolation is Postgres's unique_violation SQLSTATE.
const pgUniqueViolation = "23505"

// Store is every read and write of the bot's tables.
type Store struct{ Pool *pgxpool.Pool }

// Install is a guild's one Discord server and its raid channel.
type Install struct {
	GuildID          int64
	ServerID         string
	ServerName       string
	ChannelID        *string
	ChannelName      *string
	Failures         int
	LastError        *string
	RemovedInDiscord bool
	InstalledBy      *int64
	UpdatedAt        time.Time
}

// BindOutcome is the answer to BindServer.
type BindOutcome int

// The outcomes of BindServer.
const (
	BindInstalled BindOutcome = iota
	BindTaken
)

// BindResult carries HeldBy, the guild that holds the conflicting install, when taken.
type BindResult struct {
	Outcome BindOutcome
	HeldBy  int64
}

const installColumns = `guild_id, server_id, server_name, channel_id, channel_name, failures,
	last_error, removed_in_discord, installed_by, updated_at`

func scanInstall(row pgx.Row) (Install, error) {
	var in Install
	err := row.Scan(&in.GuildID, &in.ServerID, &in.ServerName, &in.ChannelID, &in.ChannelName,
		&in.Failures, &in.LastError, &in.RemovedInDiscord, &in.InstalledBy, &in.UpdatedAt)
	return in, err
}

// BindServer ties serverID to guildID: one guild, one server. The same pair again is
// idempotent (the channel is kept); any other pairing is taken, naming the holding guild.
func (s *Store) BindServer(ctx context.Context, guildID int64, serverID, serverName string, by int64) (BindResult, error) {
	res, err := s.bindOnce(ctx, guildID, serverID, serverName, by)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
		return s.bindConflict(ctx, guildID, serverID)
	}
	return res, err
}

func (s *Store) bindOnce(ctx context.Context, guildID int64, serverID, serverName string, by int64) (BindResult, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return BindResult{}, fmt.Errorf("discordbot: bind begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var heldServer string
	err = tx.QueryRow(ctx, `select server_id from guild_discord_installs where guild_id = $1 for update`, guildID).Scan(&heldServer)
	switch {
	case err == nil && heldServer == serverID:
		_, err = tx.Exec(ctx, `update guild_discord_installs set server_name = $2, installed_by = $3,
			failures = 0, last_error = null, removed_in_discord = false, updated_at = now() where guild_id = $1`,
			guildID, serverName, by)
		return s.commit(ctx, tx, BindResult{Outcome: BindInstalled}, err)
	case err == nil:
		return s.commit(ctx, tx, BindResult{Outcome: BindTaken, HeldBy: guildID}, nil)
	case !errors.Is(err, pgx.ErrNoRows):
		return BindResult{}, fmt.Errorf("discordbot: bind read guild %d: %w", guildID, err)
	}

	var holder int64
	err = tx.QueryRow(ctx, `select guild_id from guild_discord_installs where server_id = $1 for update`, serverID).Scan(&holder)
	switch {
	case err == nil:
		return s.commit(ctx, tx, BindResult{Outcome: BindTaken, HeldBy: holder}, nil)
	case !errors.Is(err, pgx.ErrNoRows):
		return BindResult{}, fmt.Errorf("discordbot: bind read server: %w", err)
	}
	_, err = tx.Exec(ctx, `insert into guild_discord_installs (guild_id, server_id, server_name, installed_by)
		values ($1, $2, $3, $4)`, guildID, serverID, serverName, by)
	return s.commit(ctx, tx, BindResult{Outcome: BindInstalled}, err)
}

func (s *Store) commit(ctx context.Context, tx pgx.Tx, res BindResult, execErr error) (BindResult, error) {
	if execErr != nil {
		return BindResult{}, fmt.Errorf("discordbot: bind write: %w", execErr)
	}
	if err := tx.Commit(ctx); err != nil {
		return BindResult{}, fmt.Errorf("discordbot: bind commit: %w", err)
	}
	return res, nil
}

// bindConflict answers a lost race on the unique server or guild key as taken.
func (s *Store) bindConflict(ctx context.Context, guildID int64, serverID string) (BindResult, error) {
	in, ok, err := s.InstallByServer(ctx, serverID)
	if err != nil {
		return BindResult{}, err
	}
	if ok && in.GuildID == guildID {
		return BindResult{Outcome: BindInstalled}, nil
	}
	if ok {
		return BindResult{Outcome: BindTaken, HeldBy: in.GuildID}, nil
	}
	own, ok, err := s.InstallFor(ctx, guildID)
	if err != nil {
		return BindResult{}, err
	}
	if ok && own.ServerID != serverID {
		return BindResult{Outcome: BindTaken, HeldBy: guildID}, nil
	}
	return BindResult{}, fmt.Errorf("discordbot: bind guild %d: conflicting install vanished, retry", guildID)
}

// InstallFor returns the guild's install; ok is false when none exists.
func (s *Store) InstallFor(ctx context.Context, guildID int64) (Install, bool, error) {
	return s.oneInstall(ctx, `select `+installColumns+` from guild_discord_installs where guild_id = $1`, guildID)
}

// InstallByServer returns the install holding the Discord server; ok is false when none does.
func (s *Store) InstallByServer(ctx context.Context, serverID string) (Install, bool, error) {
	return s.oneInstall(ctx, `select `+installColumns+` from guild_discord_installs where server_id = $1`, serverID)
}

func (s *Store) oneInstall(ctx context.Context, query string, arg any) (Install, bool, error) {
	in, err := scanInstall(s.Pool.QueryRow(ctx, query, arg))
	if errors.Is(err, pgx.ErrNoRows) {
		return Install{}, false, nil
	}
	if err != nil {
		return Install{}, false, fmt.Errorf("discordbot: read install: %w", err)
	}
	return in, true, nil
}

// SetChannel saves the raid channel and clears the failure state.
func (s *Store) SetChannel(ctx context.Context, guildID int64, channelID, channelName string) error {
	return s.exec(ctx, `update guild_discord_installs set channel_id = $2, channel_name = $3,
		failures = 0, last_error = null, updated_at = now() where guild_id = $1`, guildID, channelID, channelName)
}

// MarkRemovedInDiscord records that Discord says the bot is gone from the server.
func (s *Store) MarkRemovedInDiscord(ctx context.Context, guildID int64, lastErr string) error {
	return s.exec(ctx, `update guild_discord_installs set removed_in_discord = true, last_error = $2,
		updated_at = now() where guild_id = $1`, guildID, lastErr)
}

// RecordPostSuccess clears the failure counter and last error.
func (s *Store) RecordPostSuccess(ctx context.Context, guildID int64) error {
	return s.exec(ctx, `update guild_discord_installs set failures = 0, last_error = null,
		updated_at = now() where guild_id = $1`, guildID)
}

// RecordPostFailure counts a failed post and keeps its sanitised message.
func (s *Store) RecordPostFailure(ctx context.Context, guildID int64, lastErr string) error {
	return s.exec(ctx, `update guild_discord_installs set failures = failures + 1, last_error = $2,
		updated_at = now() where guild_id = $1`, guildID, lastErr)
}

// DeleteInstall removes the guild's install; an absent one is not an error.
func (s *Store) DeleteInstall(ctx context.Context, guildID int64) error {
	return s.exec(ctx, `delete from guild_discord_installs where guild_id = $1`, guildID)
}

func (s *Store) exec(ctx context.Context, query string, args ...any) error {
	if _, err := s.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("discordbot: store write: %w", err)
	}
	return nil
}
