package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// newDiscordBot builds the bot's interactions service. It is nil, with one
// warn line, unless the bot's variables are all set and the guild connect it
// installs through is on. A live bot also becomes the accounts service's Discord link view
// and routes and drains the raid posts.
// stop is serve's signal context: the drainer stops claiming when it ends.
func newDiscordBot(stop context.Context, cfg config.Config, pool *pgxpool.Pool, deps server.Deps, accounts *auth.Service, log *slog.Logger) *discordbot.Service {
	if !cfg.DiscordBotConfigured() {
		log.Warn("discordbot", "state", "bot off", "effect", "no interactions endpoint")
		return nil
	}
	if deps.Notify == nil || !deps.Notify.GuildConnectEnabled() {
		log.Warn("discordbot", "state", "bot off", "effect", "guild connect is off, so the bot has nothing to install through")
		return nil
	}
	store := &discordbot.Store{Pool: pool, Log: log}
	bot := discordbot.New(discordbot.Deps{
		Store:         store,
		REST:          discordbot.NewHTTPREST(cfg.DiscordApplicationID, cfg.DiscordBotToken),
		OAuth:         discordbot.NewHTTPOAuth(cfg.DiscordClientID, cfg.DiscordClientSecret, notify.CallbackURL(cfg.APIBaseURL)),
		PublicKey:     cfg.DiscordPublicKey,
		ApplicationID: cfg.DiscordApplicationID, CommandsGuildID: cfg.DiscordCommandsGuildID,
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL,
		Guilds: deps.Guilds, Notify: deps.Notify, Log: log,
	})
	accounts.DiscordLinks = bot
	deps.Guilds.Bot = bot
	// The raid posts: the bot's one edited post where it is installed, the webhook elsewhere.
	deps.Guilds.RaidEvents = &discordbot.RaidRouter{Bot: bot, Webhook: deps.Notify.GuildDeliverer()}
	deps.Guilds.RaidPosts = store
	// The drainer runs in this process (one warm instance, always-on CPU). It stops claiming
	// at the stop signal itself, so it does not take rows late in the 9 s shutdown budget;
	// drainDiscordBot then waits for the row in flight within that budget.
	bot.StartDrainer(stop, bot.NewDrainer(deps.Guilds))
	go registerDiscordCommands(stop, bot, log)
	return bot
}

// commandsRegisterWithin bounds the startup registration (two Discord calls and a lock wait).
const commandsRegisterWithin = time.Minute

// registerDiscordCommands pushes the slash commands to Discord off the serving path. A failure
// is logged (the error carries no token) and the next start tries again.
func registerDiscordCommands(stop context.Context, bot *discordbot.Service, log *slog.Logger) {
	ctx, cancel := context.WithTimeout(stop, commandsRegisterWithin)
	defer cancel()
	if err := bot.RegisterCommands(ctx); err != nil {
		log.Error("discordbot", "op", "register_commands", "outcome", "failed", "err", err)
	}
}
