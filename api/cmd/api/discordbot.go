package main

import (
	"context"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
)

// botEnabled reports whether the bot's variables are all set and the guild connect it installs
// through is on. A bot that is off logs one warn line saying why.
func botEnabled(cfg config.Config, deps server.Deps, log *slog.Logger) bool {
	if !cfg.DiscordBotConfigured() {
		log.Warn("discordbot", "state", "bot off", "effect", "no interactions endpoint")
		return false
	}
	if deps.Notify == nil || !deps.Notify.GuildConnectEnabled() {
		log.Warn("discordbot", "state", "bot off", "effect", "guild connect is off, so the bot has nothing to install through")
		return false
	}
	return true
}

// botDeps is the wiring serve and the raid-posts job share; serve adds the character and coach
// views the interactions answer with.
func botDeps(cfg config.Config, pool *pgxpool.Pool, deps server.Deps, log *slog.Logger) discordbot.Deps {
	rest := discordbot.NewHTTPREST(cfg.DiscordApplicationID, cfg.DiscordBotToken)
	return discordbot.Deps{
		Store:         &discordbot.Store{Pool: pool, Log: log},
		REST:          rest,
		Emoji:         rest,
		OAuth:         discordbot.NewHTTPOAuth(cfg.DiscordClientID, cfg.DiscordClientSecret, notify.CallbackURL(cfg.APIBaseURL)),
		PublicKey:     cfg.DiscordPublicKey,
		ApplicationID: cfg.DiscordApplicationID, CommandsGuildID: cfg.DiscordCommandsGuildID,
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL,
		Guilds: deps.Guilds, Notify: deps.Notify, Log: log,
	}
}

// sheetRenderer draws Share in channel's character sheet from the active build's icons
// ({TreeDataDir}/{build}/icons, as the tiles route reads them). Nil, with one warn line, when
// no build is active: the share then posts its panel without the sheet.
func sheetRenderer(cfg config.Config, deps server.Deps, log *slog.Logger) discordbot.SheetRenderer {
	build, ok := deps.Guilds.ActiveBuild()
	if !ok {
		log.Warn("discordbot", "state", "no sheet", "effect", "shares post the panel without the sheet")
		return nil
	}
	return sheet.NewRenderer(sheet.NewFileIcons(filepath.Join(cfg.TreeDataDir, build, "icons")))
}

// raidRouterFor routes the raid posts: the bot's one edited post where it is installed, the
// webhook elsewhere. serve and the raid-posts job build the same router.
func raidRouterFor(bot *discordbot.Service, deps server.Deps) guilds.RaidEvents {
	return &discordbot.RaidRouter{Bot: bot, Webhook: deps.Notify.GuildDeliverer()}
}

// newDiscordBot builds the bot's interactions service. It is nil, with one
// warn line, unless the bot's variables are all set and the guild connect it
// installs through is on. A live bot also becomes the accounts service's Discord link view
// and routes and drains the raid posts.
// stop is serve's signal context: the drainer stops claiming when it ends.
func newDiscordBot(stop context.Context, cfg config.Config, pool *pgxpool.Pool, deps server.Deps, accounts *auth.Service, log *slog.Logger) *discordbot.Service {
	if !botEnabled(cfg, deps, log) {
		return nil
	}
	d := botDeps(cfg, pool, deps, log)
	d.Characters, d.Coach = deps.Rankings, coachLinesOf(deps.Coach)
	d.Sheets = sheetRenderer(cfg, deps, log)
	bot := discordbot.New(d)
	accounts.DiscordLinks = bot
	deps.Guilds.Bot = bot
	deps.Guilds.RaidEvents = raidRouterFor(bot, deps)
	deps.Guilds.RaidPosts = d.Store
	// The drainer runs in this process (one warm instance, always-on CPU). It stops claiming
	// at the stop signal itself, so it does not take rows late in the 9 s shutdown budget;
	// drainDiscordBot then waits for the row in flight within that budget.
	drainer := bot.NewDrainer(deps.Guilds)
	drainer.StopWithin = shutdownBudget
	bot.StartDrainer(stop, drainer)
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
