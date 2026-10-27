# Discord bot (Lane G5) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Put Forever Sixty inside a guild's Discord: one raid post per raid with `In`/`Late`/`Bench`/`Out` buttons that write the site's own sign-up row and stay current through durable coalesced edits, an account link, a one-flow bot install with a channel pick, and slash commands (`/raid`, `/signup`, `/link`, `/help` first; `/char`, `/bis`, `/report`, `/attendance`, `/readiness` after).

**Architecture:** A new Go package `api/internal/discordbot` owns everything the bot does: the Ed25519-verified HTTP interactions endpoint (no gateway), the bot-token REST client, the install and identify OAuth flows (riding notify's existing redirect URI by state kind), the stores, the post drainer and every command. It imports `guilds` and `notify`; `guilds` reaches it only through two small interfaces it declares (`RaidPostMarker`, `DiscordBot`), so there is no cycle. The raid embed stays one builder in `notify` (product §8.11), extended with post states and the two-line counts field; the bot adds components around it. Pending post edits live in Postgres and any instance drains them. The web gets two new components on guild settings and one on the account page, each with a copy module pinned to the experience spec by spec-table tests.

**Tech Stack:** Go 1.x stdlib (`crypto/ed25519`, `net/http`), pgx v5, Postgres migrations (golang-migrate files), Astro + Svelte 5 runes, Vitest, Playwright.

**Spec:** `docs/superpowers/specs/2026-10-27-discord-bot-design.md` (product, §8 decisions and §10 staging binding) and `design/specs/2026-10-27-discord-bot.md` (experience: copy verbatim, states, §9 contract and acceptance). Executors read both. Where this plan rules on a conflict, the ruling is in "Rulings" below.

## Global Constraints

- Credentials: the bot token, the client secret, OAuth codes and Discord access tokens are never logged, returned in a response, put in the DOM, in browser storage or in a database row. Access tokens are used inside one function and dropped. Every error type that can wrap a Discord answer is sanitised (notify's `sanitize`); tests capture slog output and assert the token and code strings are absent.
- Only the deployed API applies migrations. No task runs a migration against any shared database; no tool or script migrates ahead of the API deploy.
- API tests run with `-p 1` on a private database per agent: `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/<private db>?sslmode=disable`, created on the `forever_test` server (`psql postgres://forever:forever@localhost:5434/forever_test -c 'create database forever_g5_<agent>'`). A test that skips for a missing `TEST_DATABASE_URL` is not a pass.
- Scoped testing: each task runs the tests of the packages and files it touched; the full suite is CI's. Exception: a task touching `web/src/layouts/Base.astro` or `web/src/lib/account/layout.ts` runs the full web Vitest, the full Playwright suite and the lhci shard.
- Never chain a piped test command into a commit or push; gate on the test command's own exit code.
- Web copy goes through a copy module (`web/src/lib/guild/discord-bot-copy.ts`, `web/src/lib/account/discord-link-copy.ts`) whose tests read the strings out of `design/specs/2026-10-27-discord-bot.md` §6.1 with `web/src/test-support/spec-table.ts`. Discord copy lives in Go copy vars whose tests read §6.2 through the Go twin of that helper (`api/internal/discordbot/spectable_test.go`, Task 3). No number is typed into prose; no string uses `soon`, `!` or `unlock`.
- File size: Go files under 800 lines, Svelte and TS files under 400. `api/cmd/api/main.go` (816) and `api/internal/auth/handler.go` (835) are already over: a task that touches either moves code out so the file ends under 800.
- Discord messages: `allowed_mentions` `{"parse": []}` on every message; member text escaped with notify's escaper; links absolute; button styles only Success (3), Secondary (2), Link (5); no emoji; every limit in experience §3.3.0 holds and is asserted.
- Nothing in a custom id, link or message identifies another account or carries a credential. Custom ids: `raid:{raidID}:{state}`, `pick:{raidID}:{state}`, `share:{token}` (token from `auth.NewSessionID()`).
- Config per environment (product §8.10): `DISCORD_APPLICATION_ID`, `DISCORD_PUBLIC_KEY`, `DISCORD_BOT_TOKEN`, optional `DISCORD_COMMANDS_GUILD_ID` (the private test server for unreleased commands). No code path names a specific application or server.
- Commit trailer, exactly: `Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>`. Conventional commit types. One PR for the lane by default.

## Rulings (spec conflicts and silences decided here)

1. **Cloud Run CPU.** Durable drains (§8.9), deferred answers and the end-time edit need CPU outside a request. Task 3 adds `--no-cpu-throttling` to the `deploy-cloudrun` flags in `.github/workflows/api.yml`. Cost note for the owner: CPU is billed for each live instance's lifetime, not per request. With scale-to-zero, an end-time edit waits for the next request to wake an instance; the row stays due in Postgres and nothing is lost.
2. **Removed in Discord** is learned from a REST answer only (403 Missing Access or 404 Unknown Guild on a post, edit or channel read), because the lane has no gateway connection. No events endpoint.
3. **Stale-tab install.** Same guild and same server: `installed`, idempotent, channel kept. That covers two officers installing at once. A guild that already holds a server and installs into a different one gets `taken`, naming itself (one guild, one server: remove first). A server held by another guild: `taken`, naming that guild.
4. **Copy placement.** The raid post's strings stay in notify's `raidCopy` beside the one builder. The bot's own sentences (presses, commands, share) live in `discordbot` copy vars (`copy_*.go`), one per task, all tested against §6.2. Together they are the spec's `botCopy`.
5. **Registration.** API startup does it (each deploy starts new instances), under a Postgres advisory lock, with a hash per application and scope in `discord_command_sets`. A command is `Released` (registered globally) or not (registered only to `DISCORD_COMMANDS_GUILD_ID`). Tasks 13 and 14 flip the lookups to released. A failed registration is logged and never stops startup.
6. **`report_ready` with the bot active** goes through the bot to the raid channel as its own message: the officer chose to post it. Without the bot it goes through the webhook, as today.
7. **A post deleted in Discord** (404 Unknown Message) marks the row `gone` and is never re-posted. A moderator's delete is respected.
8. **Signed-out link start.** `login`'s `next` must be a site path, but the start route is on the API origin. So the start route answers 302 to `{PUBLIC_BASE_URL}/login?next=/account?discord_link=start#you`, and the account page, seeing `discord_link=start` with a session, navigates on to the start route (drawing `Opening Discord…`).
9. **`/raid` for a linked viewer who is not a verified member of the server's guild** answers the press step-3 sentence (`Sign-ups are for members of {Guild}. …`). R1 holds.
10. **Plan gate on new bot posts** is the officer route's own gate: every path that creates a post (raid create, a ticked time change or cancel, post report) already passed `officerRaidGate`. The router adds no plan check, so lapsed guilds keep live edits (§8.6).
11. **Ended is the bot's own refusal.** `ended` is a derived state (past its end, not done or cancelled) that the site enforces only by disabling its segments. The press flow refuses it before calling the service, and the service's semantics do not change.
12. **R3 and the main character** follow experience rulings 5 and 8 (a move, not a refusal; the main character first), which product §8.5 already agrees with.

## Review Focus

1. **Replayed, stale or tampered interactions:** an old timestamp, a valid signature over a different body, a missing header, an oversized body. Expected: 401, and nothing runs. Test: Task 3 `TestInteractionsRejectsReplayAndTamper`.
2. **A press that names another guild's raid:** an old post after a server was re-bound, or a forged `raid:` id. Expected: no write, and the generic failure sentence. Test: Task 9 `TestPressRefusesARaidOfAnotherGuild`.
3. **Hostile member text:** raid titles and character names with `@everyone`, `<@123>`, markdown, 80-rune titles, names that fill a select label. Expected: the text renders literally, pings no one, and every Discord limit holds. Test: Task 7 `TestPostSurvivesHostileTitles`; Task 9 `TestSelectOptionsStayInsideLimits`.
4. **Concurrency:** two drainers on one due row, several presses and a site sign-up inside one debounce window, and a mark that lands while an edit is in flight. Expected: one create, one edit carrying the final counts, and no lost mark. Test: Task 8 `TestTwoDrainersPostOnce` and `TestMarkDuringEditIsKept`.
5. **Discord pushing back:** 429 with `retry_after`, 403 Missing Access, 404 Unknown Message, a token past 15 minutes. Expected: back off by `retry_after`; mark the install removed and fall back to the webhook; stop on gone; drop the late answer with one log line. Test: Task 8 `TestDrainHonoursDiscordRefusals`; Task 3 `TestRespondDropsAnswersPastTokenLife`.

## Parallel waves

Disjoint files, safe to run concurrently in separate worktrees:

- **Wave 1:** Tasks 1, 2, 6, 11, 12.
- **Wave 2:** Tasks 3 and 7 (after 1 and 2).
- **Then in order:** 4 → 5 → 8 → 9 → 10 → 13 → 14.

11 and 12 depend on no API task (the API is stubbed in Playwright and the §9 contract is fixed). They meet the API only at merge.

---

### Task 1: Config, migration 0053, and the bot's stores

**Files:**
- Modify: `api/internal/config/config.go` (Discord bot fields and checks), `api/internal/config/config_test.go`
- Create: `api/internal/db/migrations/0053_discord_bot.up.sql`, `0053_discord_bot.down.sql`
- Modify: `api/internal/notify/guild_store.go` (state kinds), `api/internal/notify/guild_store_test.go`
- Create: `api/internal/discordbot/store_install.go`, `store_links.go`, `store_refusals.go`, `store_test.go`

**Interfaces:**
- Produces (config): `Config.DiscordApplicationID string`, `Config.DiscordPublicKey ed25519.PublicKey`, `Config.DiscordBotToken string`, `Config.DiscordCommandsGuildID string`, `func (c Config) DiscordBotConfigured() bool` (all three of id, key and token set).
- Produces (notify): `const StateKindGuildWebhook = "guild_webhook"`, `StateKindBotInstall = "bot_install"`, `StateKindLink = "link"`; `ConnectState.Kind string` (GuildID 0 when null); `func (s *GuildStore) CreateState(ctx context.Context, state, kind string, userID int64, guildID *int64, ttl time.Duration) error` (`CreateConnectState` delegates with `StateKindGuildWebhook`).
- Produces (discordbot): `type Store struct{ Pool *pgxpool.Pool }`;
  `type Install struct { GuildID int64; ServerID, ServerName string; ChannelID, ChannelName *string; Failures int; LastError *string; RemovedInDiscord bool; InstalledBy *int64; UpdatedAt time.Time }`;
  `type BindOutcome int` with `BindInstalled`, `BindTaken`; `type BindResult struct { Outcome BindOutcome; HeldBy int64 }`;
  `BindServer(ctx, guildID int64, serverID, serverName string, by int64) (BindResult, error)`, `InstallFor(ctx, guildID int64) (Install, bool, error)`, `InstallByServer(ctx, serverID string) (Install, bool, error)`, `SetChannel(ctx, guildID int64, channelID, channelName string) error`, `MarkRemovedInDiscord(ctx, guildID int64, lastErr string) error`, `RecordPostSuccess(ctx, guildID int64) error`, `RecordPostFailure(ctx, guildID int64, lastErr string) error`, `DeleteInstall(ctx, guildID int64) error`;
  `type Link struct { UserID int64; DiscordUserID, DisplayName string; LinkedAt time.Time }`, `var ErrLinkConflict`, `SaveLink(ctx, userID int64, discordUserID, displayName string) error`, `LinkByUser(ctx, userID int64) (Link, bool, error)`, `UserByDiscord(ctx, discordUserID string) (int64, bool, error)`, `DeleteLink(ctx, userID int64) error`, `LinkedMembers(ctx, guildID int64) (linked, members int, err error)`, `SuggestedGuild(ctx, userID int64) (*string, error)`;
  `SaveRefusal(ctx, userID, guildID int64, heldByName string, ttl time.Duration) error`, `TakeRefusal(ctx, userID, guildID int64) (*string, error)`.

Migration 0053, decided here:
- `notify_connect_states`: add `kind text not null default 'guild_webhook'` with check in the three kinds; `guild_id` drops `not null`; check `kind = 'link' or guild_id is not null`.
- `discord_links(user_id bigint pk → users on delete cascade, discord_user_id text not null unique, display_name text not null, linked_at timestamptz not null default now())`.
- `guild_discord_installs(guild_id bigint pk → guilds on delete cascade, server_id text not null unique, server_name text not null, channel_id text, channel_name text, failures int not null default 0, last_error text, removed_in_discord boolean not null default false, installed_by bigint → users on delete set null, created_at, updated_at)`.
- `guild_discord_refusals(user_id, guild_id, held_by_name text not null, expires_at timestamptz not null, pk (user_id, guild_id))`.
- `guild_raid_discord_posts(raid_id bigint pk → guild_raids on delete cascade, guild_id bigint not null → guilds on delete cascade, channel_id text not null, message_id text, moved_from timestamptz, dirty_seq bigint not null default 1, dirty_at timestamptz default now(), wake_at timestamptz, edited_at timestamptz, attempts int not null default 0, next_attempt_at timestamptz not null default now(), lease_until timestamptz, last_error text, gone boolean not null default false, created_at timestamptz not null default now())`. Partial index on `(next_attempt_at)` `where not gone and (dirty_at is not null or wake_at is not null)`.
- `discord_command_sets(application_id text, scope text, hash text not null, registered_at timestamptz not null default now(), pk (application_id, scope))`.
- `discord_shares(token text pk, guild_id bigint → guilds on delete cascade, command text not null, options jsonb not null, used_at timestamptz, expires_at timestamptz not null)`.
- The down file reverses every step. It restores `guild_id not null` after deleting the non-webhook state rows.

`SuggestedGuild` returns the name of a guild where the account is a verified member and the install has a channel and `removed_in_discord` is false. `LinkedMembers` counts verified accounts of the guild (`guild_members`, verified) and how many of them have a `discord_links` row.

- [ ] **Step 1: Write the failing tests**

```go
// config_test.go
func TestDiscordBotConfigIsAllOrNothing(t *testing.T) // id+key set, token empty → error naming DISCORD_BOT_TOKEN; message never contains the key or token values
func TestDiscordPublicKeyMustBe32HexBytes(t *testing.T) // "zz" and 31-byte hex → error naming DISCORD_PUBLIC_KEY
func TestDiscordApplicationIDMustMatchClientID(t *testing.T) // both set and different → error; DiscordBotConfigured() false when unset
func TestDiscordBotNeedsTheClientPair(t *testing.T) // bot vars set, DISCORD_CLIENT_ID empty → error

// guild_store_test.go
func TestCreateStateRoundTripsKind(t *testing.T) {
	// CreateState(ctx, "s1", StateKindLink, user, nil, time.Minute); TakeConnectState → cs.Kind == "link", cs.GuildID == 0
	// CreateConnectState(...) → Kind == StateKindGuildWebhook
}

// store_test.go (discordbot)
func TestBindServerOneServerOneGuild(t *testing.T) {
	// BindServer(a,"srv1") → {BindInstalled,0}; BindServer(b,"srv1") → {BindTaken, HeldBy:a}
	// BindServer(a,"srv2") → {BindTaken, HeldBy:a}; BindServer(a,"srv1") again after SetChannel → BindInstalled and ChannelID kept
}
func TestSaveLinkOneToOne(t *testing.T) // u2 saving u1's discord id → errors.Is(err, ErrLinkConflict); u1 relinking a new id replaces the row
func TestLinkedMembersCountsVerifiedOnly(t *testing.T) // 2 verified (1 linked) + 1 pending linked → (1, 2)
func TestSuggestedGuildNeedsAChannelAndAPresentBot(t *testing.T) // no channel → nil; channel → "Olympus XXVII"; removed → nil
func TestTakeRefusalReadsOnceAndExpires(t *testing.T) // Save, Take → "Mograine Raiders", Take → nil; ttl -1s → nil
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_g5_<agent>?sslmode=disable go test -p 1 ./internal/config/ ./internal/notify/ ./internal/discordbot/ -run 'Discord|CreateState|BindServer|SaveLink|LinkedMembers|SuggestedGuild|TakeRefusal' -v`
Expected: FAIL (undefined fields and functions; the discordbot package does not exist).

- [ ] **Step 3: Implement**

Write the migration pair, the config fields and the `checkDiscordBot(c Config) error` and `decodeDiscordPublicKey(raw string) (ed25519.PublicKey, error)` helpers. Every message names the variable, never the value. Write `GuildStore.CreateState` and the stores above. `BindServer` runs in one transaction: it reads the guild's row and the server's row `for update`, then inserts or answers the outcome. A race on the unique `server_id` reads as `BindTaken`.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command. Expected: PASS. Also `go test -p 1 ./internal/notify/` passes unchanged.

- [ ] **Step 5: Commit**

```bash
git add api/internal/config api/internal/db/migrations/0053_discord_bot.* api/internal/notify/guild_store*.go api/internal/discordbot/store*.go
git commit -m "feat(discordbot): config, migration 0053 and the install, link and refusal stores

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 2: Discord REST and OAuth clients, message wire types, channel permissions

**Files:**
- Create: `api/internal/discordbot/doc.go`, `message.go`, `command_wire.go`, `rest.go`, `oauth.go`, `permissions.go`, `fakes_test.go`, `rest_test.go`, `oauth_test.go`, `permissions_test.go`
- Modify: `api/internal/notify/discord.go` (`Embed.Author`, `Embed.MarshalJSON`), `api/internal/notify/text.go` (exported escape wrappers), `api/internal/notify/discord_test.go`

**Interfaces:**
- Produces (notify): `Embed.Author string` (wire `author: {name}`); `func (e Embed) MarshalJSON() ([]byte, error)` (delegates to `wire()`); `func EscapeWithin(raw string, rawMax, outMax int) string`; `func Collapse(s string) string`.
- Produces (discordbot):
  `type Message struct { Content string; Embeds []notify.Embed; Components []ActionRow; Ephemeral bool }` (MarshalJSON writes `flags: 64` when ephemeral and always `allowed_mentions: {"parse": []}`);
  `type ActionRow []Component`; `type Component struct { Kind ComponentKind; Style ButtonStyle; Label, CustomID, URL, Placeholder string; Disabled bool; Options []SelectOption }`; `ComponentKind` with `KindButton = 2` and `KindStringSelect = 3`; `ButtonStyle` with `StyleSecondary = 2`, `StyleSuccess = 3` and `StyleLink = 5`; `type SelectOption struct { Label, Value, Description string }`;
  `type Command struct { Name, Description string; Options []CommandOption; Released bool }` (`Released` not marshalled; contexts `[0]` and integration_types `[0]` always written); `type CommandOption struct { Name, Description string; Type int; Required, Autocomplete bool; Choices []Choice }`; `type Choice struct { Name, Value string }`;
  `type REST interface { CreateMessage(ctx context.Context, channelID string, m Message) (string, error); EditMessage(ctx context.Context, channelID, messageID string, m Message) error; ServerLayout(ctx context.Context, serverID string) (ServerLayout, error); LeaveServer(ctx context.Context, serverID string) error; PutCommands(ctx context.Context, serverID string, cmds []Command) error; EditOriginal(ctx context.Context, interactionToken string, m Message) error; FollowUp(ctx context.Context, interactionToken string, m Message) error }` (`serverID ""` in PutCommands means global);
  `func NewHTTPREST(applicationID, botToken string) *HTTPREST` (`BaseURL` field for tests; the no-redirect 5 s client);
  `type APIError struct { Status, Code int; RetryAfter time.Duration; Message string }` with `RateLimited() bool` (429), `MissingAccess() bool` (403 or code 50001/50013 or 10004), `Gone() bool` (code 10003 or 10008);
  `type ServerLayout struct { Channels []RawChannel; Roles []Role; BotRoleIDs []string }`, `type Channel struct { ID string \`json:"id"\`; Name string \`json:"name"\`; Category *string \`json:"category"\` }`, `func PostableChannels(serverID string, l ServerLayout) []Channel`;
  `type OAuth interface { InstallURL(state string) string; LinkURL(state string) string; ExchangeInstall(ctx context.Context, code string) (InstalledServer, error); ExchangeIdentity(ctx context.Context, code string) (Identity, error) }`, `type InstalledServer struct{ ID, Name string }`, `type Identity struct{ ID, DisplayName string }`, `func NewHTTPOAuth(clientID, clientSecret, redirectURI string) *HTTPOAuth`, `const InstallPermissions = 84992` (View Channel 1<<10, Send Messages 1<<11, Embed Links 1<<14, Read Message History 1<<16);
  `FakeREST` and `FakeOAuth` in `fakes_test.go` (they record calls and return scripted errors), used by every later discordbot test.

`PostableChannels` decides, because the signature does not:
- Only type-0 text channels count.
- Permissions follow Discord's order: @everyone (role id = serverID), then the OR of the bot's roles. Administrator (1<<3) grants everything. Then the channel's @everyone overwrite, then the role overwrites (deny all, then allow all), then the member overwrite.
- A channel needs all four bits of `InstallPermissions`. A category's overwrites do not cascade, because Discord syncs them onto the child.
- Order: uncategorised first by position, then by category position, then by child position. Ties break by id.

`ExchangeIdentity` exchanges the code, calls `GET /users/@me` with the bearer token, and returns `global_name`, falling back to `username`. The token never leaves the function. `ExchangeInstall` reads the token answer's `guild` object.

- [ ] **Step 1: Write the failing tests**

```go
func TestCreateMessageSendsBotAuthAndPingsNoOne(t *testing.T) // header "Authorization: Bot t0k"; body allowed_mentions.parse == []; returns "m1"
func TestEditOriginalUsesTheInteractionWebhookWithoutBotAuth(t *testing.T) // PATCH /webhooks/{app}/{token}/messages/@original; no Authorization header
func TestAPIErrorClassifies(t *testing.T) // 429 {"retry_after":1.5} → RateLimited, RetryAfter == 1500*time.Millisecond; 403 code 50001 → MissingAccess; 404 code 10008 → Gone; err.Error() never contains "t0k"
func TestPutCommandsGlobalAndServerScoped(t *testing.T) // "" → PUT /applications/{app}/commands; "srv" → /applications/{app}/guilds/srv/commands
func TestPostableChannels(t *testing.T) // admin role → all text channels; @everyone deny SendMessages on #general → excluded; bot-role allow on #raids → included; voice channel excluded; order uncategorised, then category position
func TestInstallAndLinkURLs(t *testing.T) // InstallURL scope "bot applications.commands", permissions "84992", state; LinkURL scope "identify"
func TestExchangeIdentityPrefersGlobalNameAndDropsToken(t *testing.T) // Identity{ID:"42",DisplayName:"zulmara"}; bearer used once on /users/@me
func TestExchangeInstallReadsTheTokenAnswersGuild(t *testing.T) // InstalledServer{ID:"srv1",Name:"Olympus Discord"}; error text never contains the code
func TestMessageJSON(t *testing.T) // ephemeral → flags 64; row → {"type":1,"components":[{"type":2,"style":3,...}]}; link button carries url, no custom_id
// notify/discord_test.go
func TestEmbedMarshalsLikeTheWebhookWireAndCarriesAuthor(t *testing.T)
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && go test ./internal/discordbot/ ./internal/notify/ -run 'CreateMessage|EditOriginal|APIError|PutCommands|Postable|URLs|Exchange|MessageJSON|EmbedMarshals' -v`
Expected: FAIL (undefined).

- [ ] **Step 3: Implement the types and clients above**

Use `httptest` base URLs, the API base `https://discord.com/api/v10`, and form-encoded token exchange as in `notify/oauth.go`.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot api/internal/notify/discord.go api/internal/notify/text.go api/internal/notify/discord_test.go
git commit -m "feat(discordbot): REST and OAuth clients, message wire types and channel permissions

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 3: Interactions endpoint, the answer-or-defer responder, and wiring

**Files:**
- Create: `api/internal/discordbot/verify.go`, `interaction.go`, `respond.go`, `service.go`, `guilds_port.go`, `routes.go`, `mount.go`, `copy_common.go`, `spectable_test.go`, `verify_test.go`, `interactions_test.go`, `respond_test.go`
- Create: `api/cmd/api/discordbot.go`
- Modify: `api/cmd/api/main.go` (move the notify wiring block into `notifyservice.go` as `wireNotify`, call `newDiscordBot`, drain it on stop; the file ends under 800), `api/cmd/api/notifyservice.go`, `api/internal/server/server.go` (`Deps.DiscordBot`, mount), `.github/workflows/api.yml` (`--no-cpu-throttling` beside `--memory=1Gi`)

**Interfaces:**
- Consumes: Task 1 `Store`, config fields; Task 2 `REST`, `OAuth`, `Message`.
- Produces:
  `func Verify(pub ed25519.PublicKey, signatureHex, timestamp string, body []byte, now time.Time) error`, `ErrBadSignature`, `ErrStale`, `const MaxClockSkew = 5 * time.Minute`, `const MaxInteractionBytes = 256 << 10`;
  `type Interaction struct { ID, ApplicationID, Token, GuildID string; Type int; Data json.RawMessage; MessageID string; DiscordUserID string; ReceivedAt time.Time }` with `CommandName() string`, `CustomID() string`, `Option(name string) (string, bool)`, `Values() []string`, `Focused() (name, value string)`;
  `type Reply struct { Message Message; Update bool }` (Update = type 7 / deferred type 6);
  `type Handler func(ctx context.Context, in Interaction) (Reply, error)`; `type Autocomplete func(ctx context.Context, in Interaction) ([]Choice, error)`;
  `type Deps struct { Store *Store; REST REST; OAuth OAuth; PublicKey ed25519.PublicKey; ApplicationID, CommandsGuildID, PublicBaseURL, APIBaseURL string; Guilds GuildService; Notify *notify.Service; Log *slog.Logger }`, `func New(d Deps) *Service`;
  `type GuildService interface { GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error) }` (in `guilds_port.go`; later tasks append methods; `*guilds.Service` satisfies it);
  `func (s *Service) serverGuild(ctx context.Context, in Interaction) (guildID int64, bound bool, err error)`, `func (s *Service) viewer(ctx context.Context, in Interaction) (userID int64, linked bool, err error)`;
  `func (s *Service) Drain(ctx context.Context) error`; `func Mount(mux *http.ServeMux, s *Service)` (`POST /v1/discord/interactions`);
  `routes.go`: `func (s *Service) commandRoutes() map[string]Handler`, `componentRoutes() map[string]Handler` (keyed by custom-id prefix), `autocompleteRoutes() map[string]Autocomplete`. Later tasks add one entry each.
  `copy_common.go`: `commonCopy.Unbound`, `commonCopy.Failure` (§6.2 `Unbound server`, `Any failure`);
  `spectable_test.go`: `func specCell(t *testing.T, heading, label string) []string`, `func specString(t *testing.T, heading, label string, index int, fill map[string]string) string`, reading `../../../design/specs/2026-10-27-discord-bot.md` exactly as `web/src/test-support/spec-table.ts` does;
  `cmd/api`: `func newDiscordBot(cfg config.Config, pool *pgxpool.Pool, deps server.Deps, log *slog.Logger) *discordbot.Service`. It returns nil, with one warn line, unless `cfg.DiscordBotConfigured()` and `deps.Notify != nil && deps.Notify.GuildConnectEnabled()`.

The responder, as an algorithm (the tests pin it):
1. Run the handler in a goroutine on `context.WithoutCancel(r.Context())`, bounded at 14 min 30 s from `ReceivedAt`, and tracked by the service's WaitGroup.
2. Wait up to `s.answerWithin`, default 2.5 s. If the handler finishes in time, write type 4, or type 7 when `Update`.
3. Otherwise write type 5 with flags 64, or type 6 when the route is a component update. When the handler finishes, call `REST.EditOriginal`.
4. A handler error answers `commonCopy.Failure` ephemerally. An answer ready after 15 minutes from `ReceivedAt` is dropped, with one log line `op=discord_answer outcome=expired`.
5. Autocomplete (type 4) is never deferred. It answers inside 2.5 s or answers an empty choice list.
6. Unknown command or custom id: `commonCopy.Failure`. A guild command in a server with no install: `commonCopy.Unbound`.

The verifier works on the raw body. It reads `X-Signature-Ed25519` (hex) and `X-Signature-Timestamp` (unix seconds), signs the message `timestamp+body`, and refuses `|now − ts| > MaxClockSkew`. The route sits behind the server's middleware with no session, so a test proves an Origin-less, cookie-less POST reaches it.

- [ ] **Step 1: Write the failing tests**

```go
func TestVerify(t *testing.T) // good → nil; flipped body byte → ErrBadSignature; ts now-6m → ErrStale; "" sig → ErrBadSignature; non-hex → ErrBadSignature
func TestInteractionsAnswersPing(t *testing.T) // signed {"type":1} through server.New → 200 {"type":1}
func TestInteractionsRejectsReplayAndTamper(t *testing.T) // Review Focus 1: same signed body with ts now-10m → 401; valid sig + different body → 401; no headers → 401; body of MaxInteractionBytes+1 → 401; handler never called (route spy)
func TestRespondAnswersInline(t *testing.T) // answerWithin 50ms, handler 1ms → body type 4, data.flags 64
func TestRespondDefersThenEditsOriginal(t *testing.T) // handler 200ms → body {"type":5,"data":{"flags":64}}; after Drain, FakeREST.EditOriginal called once with the reply
func TestRespondDefersComponentUpdatesAsType6(t *testing.T)
func TestRespondDropsAnswersPastTokenLife(t *testing.T) // Review Focus 5: ReceivedAt = now-15m-1s at completion → no EditOriginal; log has outcome=expired and no token
func TestUnboundServerAnswersTheSpecSentence(t *testing.T) // == specString(t,"### 6.2 In Discord","Unbound server",0,nil)
func TestNewDiscordBotStaysOffWithoutConfig(t *testing.T) // cmd/api: nil when token empty
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./cmd/api/ -run 'Verify|Interactions|Respond|Unbound|NewDiscordBot' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

Mount is wired from `server.go` only when `d.DiscordBot != nil`. `serve` calls `bot.Drain(stopCtx)` beside `drainCompletions`. Add `--no-cpu-throttling` to the deploy flags (Ruling 1), with a comment citing product §8.9.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, then `go build ./... && go vet ./...`. Expected: PASS, and `wc -l api/cmd/api/main.go` prints under 800.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot api/cmd/api api/internal/server/server.go .github/workflows/api.yml
git commit -m "feat(discordbot): signed interactions endpoint with PING and the answer-or-defer responder

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 4: Account link (`Link Discord`, `/link`, `/v1/me` fields)

**Files:**
- Modify: `api/internal/notify/guild_connect_handler.go` (dispatch by state kind), `api/internal/notify/guild_connect_handler_test.go`
- Create: `api/internal/discordbot/link.go`, `cmd_link.go`, `copy_link.go`, `link_test.go`
- Create: `api/internal/auth/me_view.go` (move `Me`, `NotificationsBlock` and `buildNotifications` here; add the link block); modify `api/internal/auth/handler.go` (removals, plus three lines in `me`), `api/internal/auth/handler_test.go`
- Modify: `api/internal/discordbot/routes.go` (the `link` command), `api/cmd/api/discordbot.go` (register the state handler and `accounts.DiscordLinks`)

**Interfaces:**
- Consumes: Task 1 `SaveLink`, `LinkByUser`, `DeleteLink`, `SuggestedGuild`, `GuildStore.CreateState`; Task 2 `OAuth.LinkURL`, `ExchangeIdentity`; Task 3 router.
- Produces (notify): `type StateHandler func(w http.ResponseWriter, r *http.Request, cs ConnectState)`, `func (s *Service) OnState(kind string, h StateHandler)`, `func (s *Service) StartState(ctx context.Context, kind string, userID int64, guildID *int64) (string, error)` (10 min TTL). `callbackEntry` sends a taken state of kind `guild_webhook` to `guildCallback` as today, and any other kind to its registered handler.
- Produces (auth): `type DiscordLinkBlock struct { DisplayName string \`json:"display_name"\` }`; `type DiscordLinkView interface { LinkAvailable() bool; DiscordLink(ctx context.Context, userID int64) (link *DiscordLinkBlock, suggested *string, err error) }`; `Service.DiscordLinks DiscordLinkView` (nil-safe); `Me.DiscordLinkAvailable bool \`json:"discord_link_available"\``, `Me.DiscordLink *DiscordLinkBlock \`json:"discord_link"\``, `Me.DiscordLinkSuggested *string \`json:"discord_link_suggested"\`` (suggested only while unlinked).
- Produces (discordbot): routes `GET /v1/discord/link/start` (IP rate limit 20/min via `httpx.RateLimitPer`), `DELETE /v1/discord/link` (RequireSession; 204 even when nothing is linked); `func (s *Service) linkCallback(w http.ResponseWriter, r *http.Request, cs notify.ConnectState)`; `func (s *Service) linkStartURL() string` (`{APIBaseURL}/v1/discord/link/start`); `func (s *Service) LinkAvailable() bool`, `DiscordLink(...)`; `func linkButton() Component` (Link style, label `Link Discord`, the start URL); `func (s *Service) cmdLink(ctx context.Context, in Interaction) (Reply, error)`.

Redirects, verbatim:
- Signed out at start: 302 to `{PUBLIC}/login?next=%2Faccount%3Fdiscord_link%3Dstart%23you` (Ruling 8).
- Callback outcomes: 302 to `{PUBLIC}/account?discord_link={linked|denied|failed|conflict}#you`.
- `failed` covers: a different session user, an expired state, a device actor, and an exchange error.
- `denied` covers: `error=access_denied` and an empty code.

- [ ] **Step 1: Write the failing tests**

```go
func TestLinkStartSignedOutGoesThroughSignIn(t *testing.T) // Location == PUBLIC+"/login?next=%2Faccount%3Fdiscord_link%3Dstart%23you"
func TestLinkStartSignedInOpensIdentify(t *testing.T) // Location == FakeOAuth.LinkURL(state); state row kind "link", guild_id null
func TestLinkCallbackOutcomes(t *testing.T) // table: linked (row saved, display "zulmara"), denied, failed (other user / expired / exchange error), conflict (id held by another account)
func TestLinkCallbackNeverLogsTheCode(t *testing.T) // slog buffer lacks "c0de"
func TestWebhookConnectStillReturnsThroughTheSharedCallback(t *testing.T) // notify: guild_webhook state → guildCallback outcome connected
func TestUnlinkIsIdempotent(t *testing.T) // DELETE → 204, again → 204, LinkByUser → false
func TestMeCarriesTheDiscordLink(t *testing.T) // auth: available true; unlinked + active bot guild → discord_link null, suggested "Olympus XXVII"; linked → {display_name}, suggested null; DiscordLinks nil → available false
func TestLinkCommand(t *testing.T) // unlinked → spec "/link" seg 0 + button "Link Discord" URL API+"/v1/discord/link/start"; linked → seg 2 + "Account page" URL PUBLIC+"/account#you"; ephemeral both
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./internal/notify/ ./internal/auth/ -run 'Link|Unlink|Webhook.*Callback|MeCarries' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

Move the `Me` block out of `handler.go` first, with no behaviour change, so `wc -l auth/handler.go` ends under 800.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, and `go test -p 1 ./internal/notify/ ./internal/auth/` in full (callback and `/v1/me` regressions). Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/notify/guild_connect_handler*.go api/internal/discordbot api/internal/auth api/cmd/api/discordbot.go
git commit -m "feat(discordbot): Discord account link, /link and the /v1/me link fields

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 5: Bot install, channel pick, test, remove, and the settings read

**Files:**
- Modify: `api/internal/notify/guild_handler.go` (export the officer gate), `guild_handler_test.go`
- Create: `api/internal/discordbot/install.go`, `install_channels.go`, `install_test.go`
- Modify: `api/internal/guilds/discord.go` (the bot fields in settings; ctx-based `SettingsDiscordFor`), `api/internal/guilds/settings_test.go`, `api/internal/guilds/handler.go` (`Service.Bot`)
- Modify: `api/internal/discordbot/guilds_port.go`, `api/cmd/api/discordbot.go`

**Interfaces:**
- Consumes: Task 1 install store and refusals; Task 2 `OAuth.InstallURL`, `ExchangeInstall`, `REST.ServerLayout`, `PostableChannels`, `CreateMessage`, `LeaveServer`; Task 4 `notify.OnState`, `StartState`.
- Produces (notify): `func GuildOfficerGate(w http.ResponseWriter, r *http.Request, access GuildAccess, needPlan bool, fail func(op string, err error)) (int64, GuildStanding, bool)` (the existing `guildOfficer` calls it, so behaviour is identical).
- Produces (guilds): `type SettingsBot struct { ServerID string \`json:"server_id"\`; ServerName string \`json:"server_name"\`; ChannelID *string \`json:"channel_id"\`; ChannelName *string \`json:"channel_name"\`; LinkedMembers int \`json:"linked_members"\`; Members int \`json:"members"\`; Failures int \`json:"failures"\`; LastError *string \`json:"last_error"\`; RemovedInDiscord bool \`json:"removed_in_discord"\` }`; `type BotRefusal struct { GuildName string \`json:"guild_name"\` }`; `SettingsDiscord.BotAvailable bool \`json:"bot_available"\``, `.Bot *SettingsBot \`json:"bot"\``, `.BotRefusal *BotRefusal \`json:"bot_refusal"\``; `type DiscordBot interface { BotAvailable() bool; SettingsBot(ctx context.Context, guildID, viewerID int64) (*SettingsBot, *BotRefusal, error) }`; `Service.Bot DiscordBot`; `func (s *Service) SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*SettingsDiscord, error)`.
- Produces (discordbot): routes under `/v1/guilds/{id}/discord/bot`: `GET …/install` (gate with plan; account rate limit 20/min), `GET …/channels`, `PUT …/channel` (body `{"channel_id": string}`, strict decode), `POST …/test` (gate with plan; 5/min), `DELETE …` (gate without plan; frozen still 409). `func (s *Service) installCallback(w, r, cs notify.ConnectState)`. `func (s *Service) BotAvailable() bool`, `SettingsBot(...)`. `func (s *Service) Active(ctx context.Context, guildID int64) (Install, bool, error)` is true when installed with a channel and not removed. GuildService gains `SettingsDiscordFor`.

Answers, decided:
- Callback: 302 to `{PUBLIC}{guildPath}/settings?discord_bot={installed|denied|failed|taken}#discord`.
  - `taken` saves a 10-minute refusal for the officer with the holder's name (Ruling 3).
  - `failed` covers a standing that is not officer, plan, unfrozen at callback time, a session mismatch, expiry, and an exchange error.
- Channels: 200 `{"channels":[Channel…]}`. On a Discord MissingAccess, call `MarkRemovedInDiscord` and answer 502 `discord_unavailable`. With no install, 404.
- PUT channel: the id must be in a fresh `PostableChannels`, else 422 `channel_unusable`. It saves the channel name, clears `removed_in_discord`, and answers 200 with `SettingsDiscordFor`'s object.
- Test: posts `notify.GuildTestEmbed` through `CreateMessage` and answers 204. On success it calls `RecordPostSuccess`. On a refusal it answers 502 with sanitised Discord words, the same shape as the webhook test. A test never counts as a failure.
- DELETE: `LeaveServer` best effort (logged), then delete the install and the guild's post rows; 204.
- Settings read: `bot_available` is `BotAvailable()`. `bot` is null without an install. `bot_refusal` comes from `TakeRefusal` (cleared on read). `settings.discord` is present when `discordOn() || Bot != nil && Bot.BotAvailable()`.

- [ ] **Step 1: Write the failing tests**

```go
func TestInstallStartGate(t *testing.T) // officer+plan → 302 FakeOAuth.InstallURL(state) with state kind bot_install; member → 404; no plan → 403 guild_plan_required; frozen → 409 claim_contested
func TestInstallCallbackOutcomes(t *testing.T) // table: installed; same guild same server again → installed, channel kept; other guild's server → taken + settings bot_refusal.guild_name == "Mograine Raiders" once then null; guild with srv1 installing srv2 → taken naming itself; denied; officer demoted before return → failed
func TestChannelsListAndRemovedInDiscord(t *testing.T) // 200 ordered; MissingAccess → 502 and settings bot.removed_in_discord true
func TestSaveChannelRefusesAnUnusableChannel(t *testing.T) // 422 channel_unusable; usable → 200 bot.channel_name "raid-signups"
func TestBotTestMessage(t *testing.T) // CreateMessage(channel, title "Forever Sixty will post raids for Olympus XXVII here."); 403 → 502 with sanitised body; success resets failures
func TestRemoveLeavesTheServerWithoutThePlan(t *testing.T) // no plan → 204, LeaveServer("srv1") called, InstallFor false; frozen → 409
func TestSettingsCarryTheBot(t *testing.T) // guilds: bot_available, bot fields incl linked_members/members, bot null when none, Bot nil → bot_available false
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./internal/notify/ ./internal/guilds/ -run 'Install|Channel|BotTest|Remove|SettingsCarry|GuildOfficer' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, plus `go test -p 1 ./internal/notify/ ./internal/guilds/ -run 'Guild|Settings'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/notify/guild_handler*.go api/internal/discordbot api/internal/guilds/discord.go api/internal/guilds/handler.go api/internal/guilds/settings_test.go api/cmd/api/discordbot.go
git commit -m "feat(discordbot): bot install, channel pick, test message, remove and the settings read

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 6: The sign-up as a service function, one shared limiter, and change marks

**Files:**
- Modify: `api/internal/httpx/middleware.go` (`KeyLimiter`), `middleware_test.go`
- Create: `api/internal/guilds/raids_signup_service.go`, `raids_signup_service_test.go`
- Modify: `api/internal/guilds/raids_signup_handler.go` (calls the service and maps errors to the same HTTP answers as today), `raids_signups_store.go` (keep the note when nil), `raids_store.go` (`EndsAt`, `Ended`), `raids_handler.go` and `raids_match.go` (marks), `handler.go` (fields and limiter wiring)

**Interfaces:**
- Produces (httpx): `type KeyLimiter`, `func NewKeyLimiter(n int, window time.Duration) *KeyLimiter`, `func (l *KeyLimiter) Allow(key string) bool`, `func RateLimitWith(l *KeyLimiter, keyOf func(*http.Request) (string, bool)) func(http.Handler) http.Handler` (`RateLimitByKey` becomes `RateLimitWith(NewKeyLimiter(…), keyOf)`).
- Produces (guilds):
  `type SignupCommand struct { GuildID, RaidID, ActorID int64; CharacterKey, State, Role string; Note *string }` (Role `""` = the character's default; Note nil keeps the stored note, `""` on insert);
  `type SignupCandidate struct { CharacterKey, Name, Class, Spec, Role string; Verified, Main bool }`;
  `type SignupResult struct { Raid RaidView; Character SignupCandidate; Moved *string; Unchanged bool }`;
  `func (s *Service) SignUp(ctx context.Context, cmd SignupCommand) (SignupResult, error)`;
  `func (s *Service) SignupCandidates(ctx context.Context, guildID, userID int64) ([]SignupCandidate, error)` (the account's characters in the guild, main first, then by name; pending rows included with `Verified false`);
  errors `ErrNotMember`, `ErrCharacterNotFound`, `ErrNotYourCharacter`, `ErrCharacterNotVerified`, `ErrSignupRateLimited`, `ErrPlanRequired`, `ErrClaimContested`, plus the existing `ErrRaidLocked`, `ErrRaidCancelled`, `ErrNotFound`;
  `type RaidPostMarker interface { MarkRaidDirty(ctx context.Context, raidID int64) }`, `Service.RaidPosts RaidPostMarker` (nil = no-op), `Service.SignupLimiter *httpx.KeyLimiter` (default `NewKeyLimiter(signupsPerMinute, time.Minute)`, key `"signup:" + userID`, used by the PUT route's middleware and by `SignUp`);
  `func (r Raid) EndsAt() time.Time`, `func (r Raid) Ended(now time.Time) bool` (past its end and status not done or cancelled).

`SignUp` runs `putSignup`'s checks in the same order, as errors:
1. The limiter.
2. The actor is a verified member (`ErrNotMember`, which the PUT maps to 404).
3. The raid exists in this guild.
4. The character is in the roster.
5. The standing rules of `signupStanding`.

`Unchanged` is true, with no write, when the character's stored row already has the asked state and role. After a write it calls `RaidPosts.MarkRaidDirty`. The PUT route keeps its answer shape and statuses byte for byte; the existing `raids_signup_handler_test.go` is the regression.

Mark points, each after its write commits: sign-up PUT and `SignUp`; create (no mark: Task 8's router posts); patch; lock, unlock, cancel; the G2 match that sets `report_id` (in `raids_match.go`'s commit path).

- [ ] **Step 1: Write the failing tests**

```go
func TestKeyLimiterAllowsNPerWindow(t *testing.T) // 30 Allow("k") true, 31st false, other key true
func TestSignUpIsThePutRoute(t *testing.T) // own verified char In → Raid.Counts.In == 1; second char In → Moved == &"Zulmara"
func TestSignUpUnchangedOnTheSameState(t *testing.T) // Unchanged true; updated_at unchanged; marker not called
func TestSignUpKeepsTheNoteWhenNil(t *testing.T) // site note "late 15m", then SignUp(Note:nil, State late) → note kept
func TestSignUpRefusals(t *testing.T) // pending → ErrCharacterNotVerified; non-member → ErrNotMember; locked → ErrRaidLocked; cancelled → ErrRaidCancelled; member on another's char → ErrNotYourCharacter
func TestSignupLimitIsSharedBySiteAndService(t *testing.T) // 30 PUTs then SignUp → ErrSignupRateLimited
func TestSignupCandidatesMainFirst(t *testing.T) // main "Zulmara" first, then by name, pending flagged
func TestEveryRaidChangeMarksThePost(t *testing.T) // fake marker sees raid id after PUT signup, patch, lock, unlock, cancel, match
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/httpx/ ./internal/guilds/ -run 'KeyLimiter|SignUp|SignupLimit|SignupCandidates|EveryRaidChange' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, then `go test -p 1 ./internal/guilds/ -run 'Signup|Raid'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/httpx api/internal/guilds
git commit -m "refactor(guilds): sign-up as a service function with a shared limiter and post change marks

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 7: The raid post's states, the two-line counts field, and its buttons

**Files:**
- Modify: `api/internal/notify/raid_embed.go`, `raid_embed_test.go`
- Create: `api/internal/discordbot/post_message.go`, `post_message_test.go`

**Interfaces:**
- Consumes: Task 2 `Message`, `Component`, `notify.EscapeWithin`.
- Produces (notify): `type RaidPostState string` with `PostOpen`, `PostLocked`, `PostMoved`, `PostCancelled`, `PostEnded`, `PostDone`, `PostWebhook`; `type RaidCounts struct { In, Late, Bench, Out, Tanks, Healers, DPS int }`; `RaidPosted.RaidID int64`, `RaidPosted.Counts RaidCounts` (it replaces `SignedUp`); `type KillLine struct { Boss string; Duration time.Duration }`; `type RaidPost struct { RaidPosted; State RaidPostState; Was *time.Time; Came, CameOf int; Kills []KillLine }`; `func RaidPostEmbed(p RaidPost, now time.Time) Embed`; `func SignedUpValue(c RaidCounts, size int, sizeKnown bool) string`. `RaidPostedEmbed` and `RaidChangedEmbed` stay as the webhook wrappers. They now carry the full-width two-line counts field (experience §0 amends G1 §6.2), and webhook-posted keeps `Sign up on Forever Sixty.`
- Produces (discordbot): `func PostState(status string, startsAt, endsAt time.Time, movedFrom *time.Time, now time.Time) notify.RaidPostState` (precedence: cancelled, done, ended at or after the end, locked, moved while before the start, open); `func PostMessage(p notify.RaidPost, raidURL, nightURL string, now time.Time) Message`; `func raidCustomID(raidID int64, state string) string`, `func parseRaidCustomID(id string) (raidID int64, state string, ok bool)`. The `raid:` and `pick:` prefixes are constants.

Pinned by experience §3.3.1 (the table is the test oracle):
- Colours: gold `15055189`, muted `10130564`, wipe `15758187`, Done green `8377482`.
- `When` value: `<t:{unix}:F> (<t:{unix}:R>)`; struck through when cancelled. `Guild time` only with a zone.
- Done adds `Came` `{came} of {of}` and `Kills` (`Onyxia 1:42`, at most 6, then `+{n} more`).
- Row: open and moved get `In` (Success), `Late`, `Bench`, `Out` (Secondary) live, plus `Raid page` (Link). Locked and ended get the four disabled. Cancelled gets the link only. Done gets `Open the night` (Link, nightURL) only.
- The first counts clause is bold (`**23 of 40**`). Every limit in §3.3.0 holds.

- [ ] **Step 1: Write the failing tests**

```go
func TestSignedUpValueGrammar(t *testing.T) // table: {In20,Late3,Tanks3,Healers6,DPS12,Bench3,Out3; 40} → "**23 of 40** · 3 tanks · 6 healers · 12 DPS · 3 late\nBench 3 · Out 3"; Tanks1/Healers1 → "1 tank"/"1 healer"; unknown size → "**23 signed up** · …"; bench/out 0 → one line; none → "None yet"
func TestRaidPostEmbedStates(t *testing.T) // per state: Color, title prefix ("Moved: ","Cancelled: ","Done: "), Description == specCell §6.2 "Post descriptions" segment, field names in order
func TestDoneKillsCapAtSix(t *testing.T) // 8 kills → 6 lines + "+2 more"
func TestWebhookPostKeepsCountsAndNoButtons(t *testing.T) // RaidPostedEmbed → Description "Sign up on Forever Sixty.", Signed up full width (Inline false), two lines
func TestPostStatePrecedence(t *testing.T) // locked+past end → ended; cancelled+done → cancelled; moved before start → moved, after start → open
func TestPostMessageRows(t *testing.T) // open: 5 comps, styles [3,2,2,2,5], ids "raid:42:in"…; locked/ended: four Disabled; cancelled: 1 link; done: "Open the night" → nightURL; every custom id ≤ 100
func TestPostSurvivesHostileTitles(t *testing.T) // Review Focus 3: title strings.Repeat("*",80)+"@everyone <@123>" → title ≤256 runes, escaped, total embed chars ≤6000; marshalled message allowed_mentions.parse == []
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && go test ./internal/notify/ ./internal/discordbot/ -run 'SignedUpValue|RaidPostEmbed|DoneKills|WebhookPost|PostState|PostMessage|Hostile' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

Post strings go into `raidCopy` (Ruling 4). Update `guilds.raidPost` to fill `RaidID` and `Counts` from `RaidView.Counts`; this is the only guilds line in this task, in `raids_discord.go`.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, plus `go test -p 1 ./internal/guilds/ -run Discord` (webhook post regressions). Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/notify/raid_embed*.go api/internal/discordbot/post_message*.go api/internal/guilds/raids_discord.go
git commit -m "feat(notify): raid post states, the two-line counts field and the post's buttons

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 8: One post per raid, durable coalesced edits, and the router

**Files:**
- Create: `api/internal/discordbot/store_posts.go`, `drain.go`, `router.go`, `drain_test.go`, `router_test.go`
- Create: `api/internal/guilds/raids_discord_facts.go`, `raids_discord_facts_test.go`
- Modify: `api/internal/discordbot/guilds_port.go`, `api/cmd/api/discordbot.go` (the router replaces `deps.Guilds.RaidEvents`; `deps.Guilds.RaidPosts = bot.Store`; `go drainer.Run(ctx)`; drain on stop)

**Interfaces:**
- Consumes: Task 5 `Active`, `RecordPost*`, `MarkRemovedInDiscord`; Task 6 `RaidPostMarker`; Task 7 `PostState`, `PostMessage`; Task 2 `REST`, `APIError`.
- Produces (guilds): `type RaidPostFacts struct { GuildID int64; Status string; Post notify.RaidPost; EndsAt time.Time; RaidURL, NightURL string }`; `func (s *Service) RaidPostFacts(ctx context.Context, raidID int64) (RaidPostFacts, bool, error)` (counts, the guild's name, crest and zone through `raidPost`; `Came/CameOf` from the G2 night; kills from the matched report's kill order; it ignores the viewer). GuildService gains it.
- Produces (discordbot store): `type Post struct { RaidID, GuildID int64; ChannelID string; MessageID *string; MovedFrom *time.Time; DirtySeq int64; Attempts int }`; `EnqueuePost(ctx, raidID, guildID int64, channelID string) error`; `func (s *Store) MarkRaidDirty(ctx context.Context, raidID int64)` (implements `guilds.RaidPostMarker`; `dirty_seq = dirty_seq + 1, dirty_at = coalesce(dirty_at, now())`; errors logged); `SetMovedFrom(ctx, raidID int64, was time.Time) error`; `PostExists(ctx, raidID int64) (bool, error)`; `ClaimDue(ctx context.Context, now time.Time, debounce, lease time.Duration, limit int) ([]Post, error)`; `FinishEdit(ctx context.Context, raidID, seq int64, messageID string, wakeAt *time.Time, now time.Time) error`; `FailEdit(ctx context.Context, raidID int64, retryAt time.Time, lastErr string) error`; `MarkGone(ctx, raidID int64) error`; `DeletePostsForGuild(ctx, guildID int64) error`.
- Produces (drain): `type Drainer struct { Store *Store; REST REST; Facts interface{ RaidPostFacts(context.Context, int64) (guilds.RaidPostFacts, bool, error) }; OnRemoved func(ctx context.Context, guildID int64, lastErr string); Now func() time.Time; Tick, Debounce, Lease time.Duration; Log *slog.Logger }` (defaults 1 s, 2 s, 30 s); `func (d *Drainer) Run(ctx context.Context)`; `func (d *Drainer) DrainOnce(ctx context.Context) error`; `func backoff(attempts int, retryAfter time.Duration) time.Duration` = max(retryAfter, min(2 s × 2^attempts, 5 min)).
- Produces (router): `type RaidRouter struct { Bot *Service; Webhook *notify.GuildDeliverer }` implementing `guilds.RaidEvents` (`RaidPosted`, `RaidChanged`, `ReportReady`).

The algorithms the signatures do not determine:
- **Claim.** One statement, `update … set lease_until = now + lease where raid_id in (select raid_id … where not gone and (lease_until is null or lease_until < now) and next_attempt_at <= now and ((dirty_at is not null and (edited_at is null or edited_at <= now − debounce)) or wake_at <= now) order by next_attempt_at limit $n for update skip locked) returning …`. Two instances never hold one row.
- **One row.**
  1. Read the facts. If the raid is gone, delete the row.
  2. Derive the state, then build `PostMessage`.
  3. With `message_id` null, `CreateMessage`; otherwise `EditMessage`.
  4. `FinishEdit` clears `dirty_at` only `where dirty_seq = seq`, so a mark during the edit is kept. It sets `edited_at = now`, `attempts = 0`, `lease_until = null`, and `wake_at` to the next time the state changes on its own: the start while moved, else the end while not ended, else null.
- **Refusals.** `Gone()`: `MarkGone` (Ruling 7). `MissingAccess()`: `OnRemoved` (marks the install removed, so the router falls back to the webhook) and stop retrying the row. `RateLimited()` or another error: `FailEdit` at `now + backoff`.
- **Router.**
  - `RaidPosted`: when `Active`, `EnqueuePost` and no webhook post; else the webhook.
  - `RaidChanged`: a row exists → `SetMovedFrom` when `Was != nil`, then mark dirty (no new message). No row and `Active` → `EnqueuePost` (a fresh post for a webhook-era raid, product §8.1) plus `SetMovedFrom`. Else the webhook.
  - `ReportReady`: when `Active`, `CreateMessage(channel, ReportReadyEmbed)` reports posted (Ruling 6); else the webhook.
  - No plan check (Ruling 10).

- [ ] **Step 1: Write the failing tests**

```go
func TestRouterPostsThroughTheBotWhenActive(t *testing.T) // RaidPosted → posts row (message_id null); FakePoster.Posts() empty
func TestRouterFallsBackToTheWebhook(t *testing.T) // no install / no channel / removed → FakePoster 1 post, no row
func TestRaidChangedEditsTheOnePost(t *testing.T) // existing row → dirty, moved_from set, no EnqueuePost; webhook-era raid → row created
func TestDrainCreatesThenCoalesces(t *testing.T) // fake clock: enqueue, DrainOnce → CreateMessage once; 5 marks within 1s, DrainOnce → nothing; +2s DrainOnce → EditMessage once with the final counts
func TestTwoDrainersPostOnce(t *testing.T) // Review Focus 4: two Drainers DrainOnce concurrently on one due row (real DB) → CreateMessage called exactly once
func TestMarkDuringEditIsKept(t *testing.T) // Review Focus 4: FakeREST hook marks dirty mid-edit → row still dirty after FinishEdit; next drain edits again
func TestDrainHonoursDiscordRefusals(t *testing.T) // Review Focus 5: 429 retry_after 7s → next_attempt_at ≥ now+7s; 403 50001 → install removed_in_discord true and router now uses the webhook; 404 10008 → gone, never re-posted
func TestThePostEndsOnItsOwn(t *testing.T) // no marks; clock past EndsAt → one edit, buttons disabled, description "This raid has ended."
func TestReportReadyGoesThroughTheBot(t *testing.T)
func TestRaidPostFactsCarryCountsNightAndKills(t *testing.T) // guilds
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./internal/guilds/ -run 'Router|RaidChanged|Drain|TwoDrainers|MarkDuring|EndsOnItsOwn|ReportReadyGoes|RaidPostFacts' -race -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command (with `-race`), then `go build ./...`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot api/internal/guilds/raids_discord_facts*.go api/cmd/api/discordbot.go
git commit -m "feat(discordbot): one post per raid with durable coalesced edits and the raid router

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 9: Button presses, the character select, and `/signup`

**Files:**
- Create: `api/internal/discordbot/press.go`, `press_select.go`, `cmd_signup.go`, `autocomplete.go`, `copy_press.go`, `press_test.go`, `cmd_signup_test.go`
- Modify: `api/internal/discordbot/routes.go` (`raid:`, `pick:`, `signup`, autocomplete `signup`), `guilds_port.go`

**Interfaces:**
- Consumes: Task 3 `viewer`, `serverGuild`, the responder; Task 6 `SignUp`, `SignupCandidates`, `Raid.Ended`, errors; Task 7 `parseRaidCustomID`; Task 4 `linkButton`.
- Produces: `type signupAsk struct { GuildID, RaidID, UserID int64; State, CharacterKey string; Update bool }`; `func (s *Service) signupFlow(ctx context.Context, a signupAsk) (Reply, error)`; `func (s *Service) press(ctx context.Context, in Interaction) (Reply, error)`; `func (s *Service) pick(ctx context.Context, in Interaction) (Reply, error)` (Update true: type 7 replaces the ephemeral); `func (s *Service) cmdSignup(ctx context.Context, in Interaction) (Reply, error)`; `func (s *Service) autocompleteSignup(ctx context.Context, in Interaction) ([]Choice, error)`. GuildService gains `SignUp`, `SignupCandidates`, `PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]guilds.Raid, error)`, `GetRaid(ctx context.Context, guildID, raidID int64) (guilds.Raid, bool, error)`, `GuildName(ctx context.Context, guildID int64) (string, error)`.

The flow runs in experience §3.3.2's order, and the copy is that section's table:
1. Link.
2. Server to guild. The raid must belong to it, else `pressCopy.Failure` (`That did not save. Try again.`).
3. Cancelled, ended (Ruling 11) or locked sentence.
4. Candidates: none, pending one, or pending several.
5. One verified character: sign it up with the spec role, `Role ""`.
6. Several: the select. The main first; option description `{Spec} {Class} · {Role}` with ` · your main`; at most 25, with the `Showing 25 of {n}.` line.
7. Map the outcome.

`pick` re-checks that the chosen key is one of the presser's verified characters. The press never edits the post; the mark from `SignUp` does. Role words: `Tank`, `Healer`, `DPS`.

- [ ] **Step 1: Write the failing tests**

```go
func TestPressFlowSentences(t *testing.T) // table over §3.3.2 steps 1–7, each compared to specString(...§6.2 "Press: …"): unlinked (+ "Link Discord" link), cancelled, ended, locked, not a member (+ "Get set up" → PUBLIC+"/setup"), one pending, several pending, one verified In ("Zulmara is In as DPS for Onyxia's Lair. 23 of 40 signed up."), Late (+ "Add a note" → raid URL+"#signup"), Bench, Out, R3 move (" Duskstrider moved to Bench."), already, rate limited, service ErrRaidLocked → locked sentence, other error → failure; every reply Ephemeral
func TestSelectOrderAndCap(t *testing.T) // 30 verified chars → 25 options, main first with " · your main", text ends "Showing 25 of 30. Sign up the others on the raid page."
func TestSelectOptionsStayInsideLimits(t *testing.T) // Review Focus 3: 100-rune names → label/description ≤100, placeholder "Choose a character"
func TestPickReplacesTheEphemeral(t *testing.T) // Reply.Update true; chosen key not owned → failure, nothing written
func TestPressRefusesARaidOfAnotherGuild(t *testing.T) // Review Focus 2: raid of guild B pressed in A's server → failure sentence, no sign-up row
func TestPressMarksThePostAndNeverEditsIt(t *testing.T) // posts row dirty; FakeREST.EditMessage not called by the press
func TestSignupCommand(t *testing.T) // no raid option → next planned; none planned → "No raid planned yet."; character option skips the select
func TestSignupAutocomplete(t *testing.T) // raids soonest first ≤25 named "Onyxia's Lair · Thu 10 Dec"; characters: own verified only
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ -run 'Press|Select|Pick|SignupCommand|SignupAutocomplete' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot
git commit -m "feat(discordbot): sign up from the raid post's buttons, the character select and /signup

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 10: `/raid`, `/help`, the command registry and registration at startup

**Files:**
- Create: `api/internal/discordbot/cmd_raid.go`, `cmd_help.go`, `commands.go`, `register.go`, `copy_commands.go`, `cmd_raid_test.go`, `commands_test.go`, `register_test.go`
- Modify: `routes.go`, `guilds_port.go`, `api/cmd/api/discordbot.go` (call `RegisterCommands` in a goroutine at startup)

**Interfaces:**
- Consumes: Task 7 `PostMessage`; Task 8 `RaidPostFacts`; Task 9 candidates and sentences; Task 2 `PutCommands`, `Command`.
- Produces: `func Commands() []Command` (all nine, in §6.2 order, with the options of experience §3.3.3's table: names, types (3 string, choices for `status`, `class`, `slot`, `phase`), required flags, autocomplete flags). Released: `raid`, `signup`, `link`, `help`. Unreleased: `char`, `bis`, `report`, `attendance`, `readiness`. Plus `func commandsHash(cmds []Command) string` (sha256 of the canonical JSON), `func (s *Service) RegisterCommands(ctx context.Context) error`, `func (s *Service) cmdRaid(ctx context.Context, in Interaction) (Reply, error)`, `func (s *Service) cmdHelp(ctx context.Context, in Interaction) (Reply, error)`. GuildService gains `NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*guilds.NextRaidView, error)`.

`RegisterCommands` takes `pg_advisory_lock` on a constant key and splits into two scopes: global (released) and `CommandsGuildID` (unreleased; skipped when empty). Per scope it compares `commandsHash` with `discord_command_sets` and calls `PutCommands` only on a change, then upserts the hash. Errors are logged; startup continues.

`/raid`:
- The answer is the post's embed for the guild's next planned raid, with the description replaced by the viewer's state sentence, plus the post's row.
- Unlinked: the link sentence.
- Linked non-member: the press step 3 sentence (Ruling 9).
- None planned: `No raid planned yet.`
- No Share.

`/help`: one gold embed, title `Forever Sixty commands`, one `**/name** line` per command from the descriptions, footer `Forever Sixty`, no components.

- [ ] **Step 1: Write the failing tests**

```go
func TestRaidCommandStates(t *testing.T) // In/Late/Bench/Out/unsigned descriptions == §6.2 "/raid" segments; row identical to PostMessage's; unlinked; non-member; none planned
func TestHelpListsEveryCommand(t *testing.T) // 9 lines in order, each "**/raid** Your guild's next raid. Sign up from the answer."…; Footer "Forever Sixty"; no Components
func TestCommandDefinitions(t *testing.T) // 9 commands; descriptions == §6.2 "Command descriptions" segments; option descriptions == "Option descriptions"; every description ≤100; status choices In/Late/Bench/Out; slot has 17 choices; contexts [0]
func TestRegisterOnlyWhenChanged(t *testing.T) // first run → PutCommands("",released) and PutCommands("srvtest",unreleased); second run → no calls; edit a description → global PUT only
func TestRegisterFailureDoesNotStopStartup(t *testing.T) // PutCommands error → RegisterCommands returns err, logged, hash row not written
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ -run 'RaidCommand|HelpLists|CommandDefinitions|Register' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot api/cmd/api/discordbot.go
git commit -m "feat(discordbot): /raid, /help, the command registry and hash-checked registration

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 11: Web: guild settings, the bot's five states and the channel picker

**Files:**
- Modify: `web/src/lib/guild/api.ts` (types `GuildSettingsDiscord.bot_available`, `.bot`, `.bot_refusal`; `GuildBotChannel`)
- Create: `web/src/lib/guild/discord-bot-copy.ts` + `.test.ts`, `discord-bot-api.ts` + `.test.ts`, `discord-bot-state.ts` + `.test.ts`
- Create: `web/src/components/guild/GuildDiscordBot.svelte` (states 2 and 3, the remove question), `GuildDiscordChannelPick.svelte` (state 1)
- Modify: `web/src/components/guild/GuildDiscordSection.svelte` (state 5's add and webhook link, state 4's offer, delegating 1–3), `web/src/lib/guild/discord-state.ts` (the reserve picks the bot state first)
- Create: `web/tests/e2e/guild-discord-bot.spec.ts`

**Interfaces:**
- Consumes: the experience §9 contract (settings `discord.bot_available`, `bot`, `bot_refusal`; the bot routes), shared parts named in experience §3 (`PRIMARY_BUTTON_FIXED`, `SECONDARY_BUTTON_FIXED`, `FilterSelect`, `DiscordConnectLink`, `DiscordStatusLine`, `DiscordNotice`, `GuildPlanGate`, `createFocusQueue`, `takeReturnFlag`-style reader).
- Produces: `guildBotCopy` (every §6.1 site string for settings; functions for `{Server}`, `{Guild}`, `{linked}`/`{members}` forms); `guildBotInstallUrl(guildId: number, apiBase?: string): string`, `fetchBotChannels(guildId: number): Promise<GuildBotChannel[]>`, `saveBotChannel(guildId: number, channelId: string): Promise<GuildSettingsDiscord>`, `testGuildBot(guildId: number): Promise<void>`, `removeGuildBot(guildId: number): Promise<void>` (each invalidates the guild's cached reads as `discord-api.ts` does); `type GuildBotState = 'picker' | 'installed' | 'removed' | 'webhook-only' | 'nothing'`, `guildBotState(discord: GuildSettingsDiscord, changingChannel: boolean): GuildBotState` (§3.1 precedence; with `bot_available` false it answers `webhook-only` or `nothing`, and the section draws G1 unchanged); `botReturnOutcome(flag: string, refusal: { guild_name: string } | null): BotReturnOutcome | null` (§3.1 return table: message, role, focus target, `holdsPicker`); `preselectChannel(channels: GuildBotChannel[], current: string | null): string | null`; `groupChannels(channels): { category: string | null; channels: GuildBotChannel[] }[]`; `GUILD_BOT_RESERVE` (heights from the §3.1 caps).

Behaviour pinned by experience §3.1 and §4.1:
- The return flag is `discord_bot`, read once and dropped with `history.replaceState`, keeping `#discord`.
- `installed` holds the picker height until the read answers.
- An empty `Save channel` press shows `Choose a channel.` with `aria-invalid`.
- 422 `channel_unusable` refetches, clears the select, and shows the race line, focused.
- Read-only: the add and webhook links are `aria-disabled` with no `href`, test, change and save are `disabled`, and `Remove` stays live. Frozen disables all.
- A 403 `guild_plan_required` gives `Not saved. This guild's officer tools have ended.`, focused.
- The `Open in Discord` URL is `https://discord.com/channels/{server_id}/{channel_id}`.

- [ ] **Step 1: Write the failing tests**

```ts
// discord-bot-copy.test.ts — each string === specString(SPEC, '### 6.1 On the site', <row>, i, fill)
it('pins every settings row', ...) // Add button, Add explanation, Webhook path, Bot offer (state 4), Picker, Picker list states, Picker race, Installed, Linked members (3 forms), Fallback note, Removed in Discord, Remove question, Read-only, Install returns
// discord-bot-state.test.ts
it('follows the §3.1 precedence', ...) // channel null → picker; channel + !removed → installed; removed → removed; bot null + webhook → webhook-only; none → nothing; bot_available false → never picker/installed/removed
it('preselects the current, then a raid channel, then none', ...) // current wins; "Raid-Signups" over "general"; none → null
it('maps each return flag', ...) // installed → status, focus select, holdsPicker; taken with refusal names "Mograine Raiders"; taken without → the expired sentence
// discord-bot-api.test.ts
it('builds the install URL on the API origin', ...) // `${api}/v1/guilds/7/discord/bot/install`
```
```ts
// guild-discord-bot.spec.ts — experience §9 items 1–7 and 9 for this surface, API stubbed
test('state 5 links to the install route and the webhook route', ...)
test('each install return shows its line, moves focus and leaves #discord with no query', ...)
test('picker loading, empty, error, ready, preselect, empty press, save, 422 race', ...)
test('installed: destination, linked-members forms, 3-failure warning, change channel, remove to state 4 and 5', ...)
test('webhook only shows the outline offer', ...); test('removed in Discord, both variants', ...)
test('officer read-only and frozen', ...); test('axe clean, 44 px controls at 360, heights inside the §3.1 caps', ...)
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd web && npx vitest run src/lib/guild/discord-bot-*.test.ts && npx playwright test tests/e2e/guild-discord-bot.spec.ts`
Expected: FAIL (modules missing).

- [ ] **Step 3: Implement the modules and components above**

Every Svelte file stays under 400 lines. The section imports the bot parts lazily, as it does its other chunks.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, plus `npx playwright test tests/e2e/guild-settings-discord.spec.ts tests/e2e/guild-settings.spec.ts` (G1 regressions) and `npx vitest run src/lib/guild`. Expected: PASS. Then take the §9 site captures 1–7, 9–11 for this surface: 360, 390, 1440 and 2000, side by side with `discord-bot-settings` at 1440.

- [ ] **Step 5: Commit**

```bash
git add web/src/lib/guild web/src/components/guild web/tests/e2e/guild-discord-bot.spec.ts
git commit -m "feat(web): guild settings adds the Forever Sixty bot with the channel picker

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 12: Web: account page, the Discord link row, its loading reserve, and the changelog

**Files:**
- Modify: `web/src/lib/account/api.ts` (`Me.discord_link_available`, `discord_link`, `discord_link_suggested`)
- Create: `web/src/lib/account/discord-link-copy.ts` + `.test.ts`, `discord-link-api.ts`, `discord-link-return.ts` + `.test.ts`
- Create: `web/src/components/account/AccountDiscordLink.svelte`
- Modify: `web/src/components/Account.svelte` (the `You` panel gets `id="you"` and `scroll-margin-top: 80px` and mounts the row after the pseudonym note; net lines do not grow beyond the mount), `web/src/layouts/Base.astro` (`data-discord-link`), `web/src/lib/account/layout.ts` + `layout-reserve.test.ts`
- Create: `web/src/content/changelog/2026-xx-xx-discord-bot.md` (dated the ship day; §6.1 Changelog string, verbatim)
- Create: `web/tests/e2e/account-discord-link.spec.ts`

**Interfaces:**
- Consumes: the §9 `/v1/me` contract; `DiscordConnectLink`, `DiscordStatusLine`, `opening-watchdog.ts` (the 10 s stuck line), the `pageshow` reset.
- Produces: `discordLinkCopy`; `discordLinkStartUrl(apiBase?: string): string` (`${api}/v1/discord/link/start`); `unlinkDiscord(): Promise<void>` (DELETE, then patch `me`); `type DiscordLinkFlag = 'linked' | 'denied' | 'failed' | 'conflict' | 'start'`; `discordLinkReturn(flag): DiscordLinkOutcome` (message, role, focus, help line, and `navigate: true` for `start` (Ruling 8)); `discordLinkHint(me): 'none' | 'unlinked' | 'suggested' | 'linked'`; `ACCOUNT_DISCORD_LINK_RESERVE` (1440/390: unlinked 115/135, suggested 120/140, linked 95/95).

The `Base.astro` inline script's copy of the hint function is kept verbatim with `layout.ts` (the existing `layout-reserve.test.ts` pattern evaluates the inline text against the same fixtures).

- [ ] **Step 1: Write the failing tests**

```ts
// discord-link-copy.test.ts — Account lead, Account controls, Account returns rows from §6.1
// discord-link-return.test.ts
it('maps each flag', ...) // linked → status, focus Unlink; conflict → alert + help line "Sign in to that account and unlink it on its account page, then try again."; start → navigate true
it('reads the hint from me', ...) // unavailable → none; suggested name → suggested; display_name → linked
// layout-reserve.test.ts: the inline Base.astro hint === discordLinkHint for each fixture
```
```ts
// account-discord-link.spec.ts — experience §9 items 8 and 9, API stubbed
test('unlinked default and suggested (the fill)', ...); test('each discord_link flag', ...)
test('discord_link=start with a session goes on to the start route', ...)
test('unlink asks in place, Escape cancels, success shows Unlinked. and focuses Link Discord', ...)
test('Opening Discord… and the 10 s stuck line under fake timers', ...)
test('absent when unavailable; #you scrolls the panel to the top; axe clean; 44 px at 360', ...)
test('warm CLS at 390 and 1440 ≤ 0.02', ...)
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd web && npx vitest run src/lib/account/discord-link-*.test.ts src/lib/account/layout-reserve.test.ts && npx playwright test tests/e2e/account-discord-link.spec.ts`
Expected: FAIL.

- [ ] **Step 3: Implement the modules, the component and the mount above**

- [ ] **Step 4: Run the full web suites (shared layout touched)**

Run: `cd web && npm test && npm run test:e2e && npm run lhci`
Expected: all PASS; the lhci `/account` and `/guild/` budgets unchanged. Then take the §9 site captures 8–11 for the account: 360, 390, 1440 and 2000, side by side with `discord-bot-account` at 1440.

- [ ] **Step 5: Commit**

```bash
git add web/src/lib/account web/src/components/account/AccountDiscordLink.svelte web/src/components/Account.svelte web/src/layouts/Base.astro web/src/content/changelog web/tests/e2e/account-discord-link.spec.ts
git commit -m "feat(web): link Discord from the account page

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 13: Share in channel, the item embed, `/char` and `/bis`

**Files:**
- Create: `api/internal/discordbot/share.go`, `item_embed.go`, `cmd_char.go`, `cmd_bis.go`, `copy_lookup.go`, `share_test.go`, `cmd_char_test.go`, `cmd_bis_test.go`
- Modify: `api/internal/rankings/handler.go` (extract `CharacterCard`), `api/internal/bis/band.go` (preset selection), `api/internal/guilds/readiness.go` (`CharacterReadinessWord`), `routes.go`, `commands.go` (`char`, `bis` released), `guilds_port.go`, `api/cmd/api/discordbot.go`

**Interfaces:**
- Consumes: Task 10 registry; Task 3 responder; Task 1 `discord_shares`; `trees.Data.Latest()` for the active build.
- Produces:
  `func (s *rankings.Service) CharacterCard(ctx context.Context, region, ruleset, name string) (rankings.Character, bool, error)` (the handler's body minus HTTP; the handler calls it);
  `func (s *guilds.Service) CharacterReadinessWord(ctx context.Context, guildID int64, characterKey string) (string, error)` (the board's word: `Stocked`, `Short: …`, `Not shared`);
  `func bis.LoadPresetBand(dataDir, build, fileSlug string, band int, faction, preset string) (Band, bool, error)` (mirrors `web/src/lib/bis/presets.ts`: `Phase 1 raid` is level 60 with the raid-ready preset; the bare levels use the bare preset);
  `type ItemFacts struct { ID int; Name string; Quality int; ItemLevel int; Icon, Source, URL string }`, `func ItemEmbed(it ItemFacts, slot string, delta string, author string) notify.Embed` (title the name linking to `it.URL`, colour the quality's game colour, thumbnail `{PUBLIC}/data/{build}/icons/{icon}.webp`, the description lines of experience §3.3.3, quality words `Poor`…`Legendary`, minus sign U+2212, one decimal);
  `func (s *Service) offerShare(ctx context.Context, guildID int64, command string, opts map[string]string) (Component, error)` (stores the token, 1 h expiry); `func (s *Service) share(ctx context.Context, in Interaction) (Reply, error)` (recompute as a visitor through `visitorAnswer(command)`; empty → `That is no longer public, so it was not shared.`; used → no second share; else `FollowUp` public plus `Reply{Update: true}` with the button disabled, labelled `Shared`);
  `func (s *Service) cmdChar(...)`, `cmdBis(...)`, `autocompleteChar(...)`, `autocompleteSpec(...)`.

The item facts source is the build's item table that the guild loot tab and the BiS page already read. Locate it from `guilds/loot.go` (`Store.Trees`) and add one exported reader beside it; do not add a second table.

`/char`:
- Card: name, link, class colour, crest `{PUBLIC}/icons/hd/crests/{class}.webp`, `Level {n} {Spec} {Class} · {Guild}`.
- Inline fields `Item level`, `Readiness`, `Synced`, then `Top 3 upgrades`. Tank specs say `By score gained…`.
- Up to three item embeds by gap, then the row `Share in channel`, `Character page`.
- No gear shared, or private or unknown: the §6.2 sentences.

`/bis`:
- Defaults from the linked main: class, spec, band; faction is the guild's.
- With a slot: the pick plus two alternatives. Without one: five slots by the first alternative's `dps_delta` (experience ruling 4).
- Footer `{Spec} {Class} · {band label} · Forever Sixty`.

- [ ] **Step 1: Write the failing tests**

```go
func TestItemEmbedGrammar(t *testing.T) // Color == 10696174 for Epic (--color-rarity-epic #a335ee, never the -text variant), Thumbnail absolute icon URL, Description "Head · Epic · item level 61\n−15.0 DPS · PvP rank 14 · …"
func TestCharCard(t *testing.T) // spec strings; 3 upgrade embeds ordered by gain "+9.2 DPS"; no gear → one embed + "No gear shared yet…"; private/unknown → "No public character by that name."; ≤10 embeds, ≤6000 chars
func TestBisWithAndWithoutSlot(t *testing.T) // text line "Head for a Fury Warrior, level 60 Phase 1 raid:"; author "Best in slot"; without slot five embeds ordered by first-alternative delta; unlinked w/o class → link sentence + button
func TestShareRecomputesAsAVisitorOnce(t *testing.T) // FollowUp public (no flags); ephemeral updated with disabled "Shared"; second press → no FollowUp; character turned private → "That is no longer public, so it was not shared."
func TestCharacterCardMatchesTheRoute(t *testing.T) // rankings: handler output unchanged
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./internal/rankings/ ./internal/bis/ ./internal/guilds/ -run 'ItemEmbed|CharCard|Bis|Share|CharacterCard|Preset|ReadinessWord' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

Flip `char` and `bis` to `Released`.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, plus `go test -p 1 ./internal/rankings/`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/discordbot api/internal/rankings api/internal/bis api/internal/guilds/readiness.go api/cmd/api/discordbot.go
git commit -m "feat(discordbot): /char and /bis with item embeds and Share in channel

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

### Task 14: `/report`, `/attendance`, `/readiness`

**Files:**
- Create: `api/internal/guilds/discord_reads.go`, `discord_reads_test.go` (service functions extracted from the report, attendance and readiness handlers)
- Create: `api/internal/discordbot/cmd_report.go`, `cmd_attendance.go`, `cmd_readiness.go`, `copy_officer.go`, matching `_test.go` files
- Modify: `routes.go`, `commands.go` (the three released), `guilds_port.go`

**Interfaces:**
- Consumes: Task 13 `ItemEmbed`, `offerShare`; Task 3 `viewer`, `serverGuild`.
- Produces (guilds):
  `type ReportAnswer struct { Raid RaidView; Report *ReportView; Officer, Public bool }`, `func (s *Service) ReportFor(ctx context.Context, guildID, viewerID int64, raidID *int64) (ReportAnswer, bool, error)` (the last done raid, or the named one; `Public` when the matched log's visibility is Everyone);
  `type AttendanceRecord struct { Name, Class, Chip string; Last []AttendanceLine }`, `func (s *Service) AttendanceFor(ctx context.Context, guildID, viewerID int64, name string) ([]AttendanceRecord, error)` (own characters, main first, at most 5; `ErrOfficerOnly` for a member naming another, `ErrNoSuchMember`);
  `type ReadinessAnswer struct { Raid *Raid; Signed, Failing int; Rows []string; More int }`, `func (s *Service) NextRaidReadiness(ctx context.Context, guildID, viewerID int64, now time.Time) (ReadinessAnswer, error)` (rank, plan and freeze on every call: `ErrNotOfficer`, `ErrPlanRequired`, `ErrClaimContested`);
  each existing handler calls its extracted function, with its HTTP answers unchanged.
- Produces (discordbot): `cmdReport`, `cmdAttendance`, `cmdReadiness`, `autocompleteReport`, `autocompleteMember`.

Pinned by experience §3.3.3:
- `/report`: officers get `Most deaths` and `Consumables`. At most 8 loot embeds, most valuable quality first, and the whole message is at most 9 embeds. Share only when `Public`, and the shared answer is the night line and kill order, linking the public report.
- `/attendance`: G2's chip verbatim, outcome words (never "no-show"), no Share.
- `/readiness`: at most 10 failing rows, then `+{n} more on the board`. The plan-lapsed answer carries a `Subscribe the guild` link to `{guildPath}/settings#billing`. No Share.

- [ ] **Step 1: Write the failing tests**

```go
func TestReportMemberAndOfficer(t *testing.T) // member: no Most deaths/Consumables; officer: both; ≤9 embeds; private log → no Share; Everyone → Share present; none → "No raid logged yet. A raid shows here once its log is matched."
func TestAttendanceRules(t *testing.T) // own two chars main first; member naming another → "Only officers can look up another member's attendance."; officer by name; unknown → "No member of Olympus XXVII by that name."; "Not in the log" wording
func TestReadinessGate(t *testing.T) // member, lapsed (with Subscribe link), frozen, none planned, unlinked sentences; officer: 12 failing → 10 lines + "+2 more on the board"
func TestExtractedReadsMatchTheRoutes(t *testing.T) // guilds: report/attendance/readiness HTTP answers unchanged
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=… go test -p 1 ./internal/discordbot/ ./internal/guilds/ -run 'Report|Attendance|Readiness|ExtractedReads' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the signatures above**

Flip the three commands to `Released`.

- [ ] **Step 4: Run them to verify they pass**

Run: the Step 2 command, plus `go test -p 1 ./internal/guilds/` and `go vet ./...`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add api/internal/guilds api/internal/discordbot
git commit -m "feat(discordbot): /report, /attendance and /readiness

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

## After merge (owner and capture checklist, not tasks)

1. Owner, product §5 steps 1–4: create the bot user, store `DISCORD_BOT_TOKEN` in Secret Manager, set `DISCORD_PUBLIC_KEY`, `DISCORD_APPLICATION_ID` and `DISCORD_COMMANDS_GUILD_ID` on the service, deploy, then set the Interactions Endpoint URL `https://api.foreversixty.gg/v1/discord/interactions`. Set DISCORD_* env and the always-on CPU flags in the same api.yml deploy commit.
2. Experience §9 Discord captures D1–D12 in the private test server, desktop and phone clients, beside the mock frames; this includes G1's rows 17 and 24.

## Self-review

- **Spec coverage.**
  - Product §4.1 → T5/T11. §4.2 → T4/T12. §4.3 → T7/T8/T9. §4.4 → T9/T10/T13/T14. §4.5 and §8.6 → Ruling 10, T8.
  - §8.1 → T8. §8.2/§8.3 → T1/T5. §8.4 → T1/T4. §8.5 → T6/T9. §8.7 → T13. §8.8 → T13/T14 extractions. §8.9 → T8 + T3 responder. §8.10 → T1 config, T10 registry. §8.11 → T7. §8.12 → task order.
  - Experience §3.1 → T11. §3.2 → T12. §3.3.0–3.3.4 → T7, T9, T10, T13, T14. §4 → T11/T12 e2e. §6 → copy tests in every task. §9 contract → T4/T5. §9 E2E 1–9 → T11/T12. §9 item 10 → T7/T9/T13 tests.
  - Changelog → T12. Owner steps and Discord captures → after merge.
- **Step scan.** Each step names its test with assertions, its run command and expected output, and the signature with its file. Algorithms the signatures do not fix are spelled out once: permissions, the responder, claim/finish/backoff, the router, registration. Copy is never transcribed; tests read it from the spec.
- **Type consistency.**
  - `RaidPostMarker.MarkRaidDirty` (T6) is implemented by `Store.MarkRaidDirty` (T8).
  - `RaidPosted.Counts` (T7) replaces `SignedUp`, and T7 updates `guilds.raidPost`.
  - `GuildService` grows only in sequential tasks (3, 5, 8, 9, 10, 13, 14).
  - `Active` (T5) is used by T8. `linkButton` (T4) is used by T9, T10, T13 and T14. `Choice` (T2) serves both options and autocomplete.
- **Review Focus.** Five lines, each with a named test in its owning task (T3, T7, T8, T9).
- **Proportion.** Signatures, test names and assertions only. No function bodies beyond the three algorithms.
