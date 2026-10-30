package discordbot

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// defaultAnswerWithin is how long a handler gets before the request
	// answers with a deferral (Discord's own limit is 3 s).
	defaultAnswerWithin = 2500 * time.Millisecond
	// interactionTokenLife is how long Discord accepts follow-ups.
	interactionTokenLife = 15 * time.Minute
	// workBudget is how long deferred work may run from ReceivedAt, leaving
	// the last 30 s of the token's life for the edit.
	workBudget = 14*time.Minute + 30*time.Second
)

// Deps is everything the service needs, injected.
type Deps struct {
	Store                                                     *Store
	REST                                                      REST
	OAuth                                                     OAuth
	PublicKey                                                 ed25519.PublicKey
	ApplicationID, CommandsGuildID, PublicBaseURL, APIBaseURL string
	Guilds                                                    GuildService
	// Characters answers /char: the public character page.
	Characters CharacterCards
	// Sheets draws the character sheet Share in channel posts for /bis and /char. Nil posts
	// the panel without it.
	Sheets SheetRenderer
	// Coach words the log coach's fix-first line in /report. Nil: no coach lines.
	Coach  CoachLines
	Notify *notify.Service
	// Emoji is the application-emoji API the roster's spec marks are read from; nil draws the
	// roster as text lines.
	Emoji EmojiAPI
	Log   *slog.Logger
	// Background runs work that outlives a request. Nil means in-process
	// goroutines tracked by Drain; see Background.
	Background Background
}

// Background is the seam for deferred answers: the work that runs after the
// request has already answered with a deferral. The default runs it in this
// process, which needs CPU outside a request (Cloud Run: always-on CPU). A
// separate worker would implement this instead and run the work elsewhere.
type Background interface {
	// Start runs work without blocking the caller. It reports false, and
	// runs nothing, once Drain has begun.
	Start(work func()) bool
	// Drain refuses new work, then waits for started work, until ctx ends.
	Drain(ctx context.Context) error
}

// inProcess is the default Background.
type inProcess struct {
	mu       sync.Mutex
	draining bool
	wg       sync.WaitGroup
}

func (p *inProcess) Start(work func()) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.draining {
		return false
	}
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		work()
	}()
	return true
}

func (p *inProcess) Drain(ctx context.Context) error {
	p.mu.Lock()
	p.draining = true
	p.mu.Unlock()
	done := make(chan struct{})
	go func() { p.wg.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Service is the bot's interactions endpoint and what hangs off it.
type Service struct {
	d             Deps
	log           *slog.Logger
	bg            Background
	answerWithin  time.Duration
	now           func() time.Time
	commands      map[string]Handler
	components    map[string]Handler
	autocompletes map[string]Autocomplete
	updating      map[string]bool
	// commandSet replaces Commands() when set (tests).
	commandSet func() []Command
	// candidates caches each account's sign-up candidates briefly (candidates.go).
	candidates candidateCache
	// sources is the BiS build's loot.json, for item source lines (item_source.go).
	sources sourceBook
	// emoji holds the spec emoji ids the roster draws (emoji.go).
	emoji *EmojiRegistry
	// stopDrainer ends the post drainer StartDrainer began; nil when none runs.
	stopDrainer func(ctx context.Context) error
}

// New builds the service.
func New(d Deps) *Service {
	s := &Service{d: d, log: d.Log, bg: d.Background, answerWithin: defaultAnswerWithin, now: time.Now}
	if s.log == nil {
		s.log = slog.Default()
	}
	if s.bg == nil {
		s.bg = &inProcess{}
	}
	if d.Emoji != nil {
		s.emoji = &EmojiRegistry{API: d.Emoji, Log: s.log}
	}
	s.commands, s.components = s.commandRoutes(), s.componentRoutes()
	if d.Notify != nil {
		d.Notify.OnState(notify.StateKindLink, s.linkCallback)
		d.Notify.OnState(notify.StateKindBotInstall, s.installCallback)
	}
	s.autocompletes, s.updating = s.autocompleteRoutes(), updateComponents()
	return s
}

// Drain stops the post drainer, then waits for deferred answers still running, until ctx ends.
func (s *Service) Drain(ctx context.Context) error {
	var drainerErr error
	if s.stopDrainer != nil {
		drainerErr = s.stopDrainer(ctx)
	}
	return errors.Join(drainerErr, s.bg.Drain(ctx))
}

// StartDrainer runs d in this process until stop ends (the stop signal) or Drain, whichever
// comes first. Call it once, before serving.
func (s *Service) StartDrainer(stop context.Context, d *Drainer) {
	ctx, cancel := context.WithCancel(stop)
	done := make(chan struct{})
	go func() {
		defer close(done)
		d.Run(ctx)
	}()
	s.stopDrainer = func(wait context.Context) error {
		cancel()
		select {
		case <-done:
			return nil
		case <-wait.Done():
			return fmt.Errorf("discordbot: post drainer: %w", wait.Err())
		}
	}
}

// NewDrainer is a drainer over the service's store and REST client, reading facts from facts.
func (s *Service) NewDrainer(facts PostFacts) *Drainer {
	return &Drainer{Store: s.d.Store, REST: s.d.REST, Facts: facts, Marks: s.emoji, Log: s.log}
}

// serverGuild is the guild the interaction's server is installed for.
func (s *Service) serverGuild(ctx context.Context, in Interaction) (guildID int64, bound bool, err error) {
	if in.GuildID == "" {
		return 0, false, nil
	}
	install, found, err := s.d.Store.InstallByServer(ctx, in.GuildID)
	if err != nil || !found {
		return 0, false, err
	}
	return install.GuildID, true, nil
}

// viewer is the site account linked to the interaction's Discord user.
func (s *Service) viewer(ctx context.Context, in Interaction) (userID int64, linked bool, err error) {
	if in.DiscordUserID == "" {
		return 0, false, nil
	}
	return s.d.Store.UserByDiscord(ctx, in.DiscordUserID)
}
