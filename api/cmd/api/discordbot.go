package main

import (
	"log/slog"

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
func newDiscordBot(cfg config.Config, pool *pgxpool.Pool, deps server.Deps, accounts *auth.Service, log *slog.Logger) *discordbot.Service {
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
	// The drainer runs in this process (one warm instance, always-on CPU) until the stop's
	// drainDiscordBot, inside the shared shutdown budget.
	bot.StartDrainer(bot.NewDrainer(deps.Guilds))
	return bot
}
