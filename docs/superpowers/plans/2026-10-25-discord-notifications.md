# Discord notifications Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A Premium member pastes a Discord incoming-webhook URL on the account page, picks which events post (a server sim finished; an upload finished parsing), sends a test, and then gets one embed per event in the channel; a webhook that fails five times in a row is switched off and the page says so.

**Architecture:** A new `api/internal/notify` package owns the `notification_webhooks` table (one row per user and channel kind, the URL's token encrypted with AES-256-GCM under a key from Secret Manager), URL validation, the embed builder, a Discord poster with a five-second timeout, and the failure counter. The two jobs that produce events (`sims.Run` on `Finish`, `parse.Report` on `StatusComplete`) call `notify.Deliver` from the job, never from a request handler. `/v1/me` carries the block's state; three routes save, test and delete. The web block lives on the account page and is gated by `entitlements.notifications`.

**Tech Stack:** Go (pgx v5, `crypto/aes`+`cipher.NewGCM`, `net/http` with a timeout), Svelte 5, vitest, Playwright.

**Spec:** Product: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.6, §5. Experience: `design/specs/2026-10-25-discord-notifications.md` (ux-designer; §3 layout, §4 states, §6 copy incl. the two embeds, §10 contract bind — filled into Tasks 4–6 when it lands).

## Global Constraints

- Only incoming webhook URLs on `discord.com` or `discordapp.com` with path `/api/webhooks/<id>/<token>` are accepted; host and path shape validated; nothing else.
- The URL is a credential: stored encrypted (token) with the channel id in clear; never returned by any route, never logged; deleted with the account (`on delete cascade`).
- Posts happen from the producing job with a 5 s timeout and no retry beyond the next event; the fifth consecutive failure disables the webhook (`disabled_reason = "failed five times"`, `last_error` kept); a success resets the counter.
- Gating by `Can(FeatureNotifications)` on save/test and on delivery; a lapsed plan keeps the row and stops posts (`disabled_reason = "plan lapsed"` computed at read time, never stored).
- The encryption key is `NOTIFY_ENCRYPTION_KEY` (32 bytes, base64) from Secret Manager, loaded by `config` like the Stripe keys; its absence turns the feature off (routes unmounted, jobs skip delivery) and is logged once at startup; no key value in any log, file or chat.
- Copy only from the copy module, verbatim from the experience spec §6 (incl. the embed texts, which the API builds from one Go copy table mirrored by a Go test).
- Tests: API with `TEST_DATABASE_URL` and `-p 1`; web `npm run check`, vitest for touched files, the lane's e2e plus `account-billing.spec.ts`; the account page is a shared surface — lane close runs the full web suites.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. A pasted URL on another host, or a Discord URL that is not an incoming webhook, is refused with the spec's message and never stored (Task 1 validator table test, Task 2 route test).
2. A finished server sim for a member with `sim_finished` on posts exactly one embed; a member with it off, or lapsed, or disabled, gets none (Task 3 job tests).
3. Five consecutive failures disable the webhook and the sixth event does not post; a success before the fifth resets the count (Task 1 store test, Task 3).
4. The URL never appears in `/v1/me`, in any route answer, or in logs; the page shows the channel id only (Task 2, Task 5).
5. Deleting the webhook removes the row; deleting the account removes it by cascade (Task 1 migration test).

---

### Task 1: Migration, store, validator, crypto

**Files:**
- Create: `api/internal/db/migrations/0046_notification_webhooks.up.sql`, `.down.sql`
- Create: `api/internal/notify/notify.go` (types, `Kind`, `Event`), `store.go`, `url.go` (`ParseWebhookURL`), `crypto.go` (`Sealer`), tests beside each
- Modify: `api/internal/config/config.go` (`NotifyEncryptionKey []byte` from `NOTIFY_ENCRYPTION_KEY`, decoded, 32 bytes or empty), `api/README.md` (the env var)

**Interfaces:**
```go
const KindDiscord = "discord"
type EventName string; const EventSimFinished EventName = "sim_finished"; EventUploadParsed EventName = "upload_parsed"
const MaxFailures = 5
type Webhook struct { UserID int64; Kind string; ChannelID string; Events map[EventName]bool; Failures int; LastError string; DisabledAt *time.Time; UpdatedAt time.Time }
type ParsedURL struct { ChannelID, Token string }
func ParseWebhookURL(raw string) (ParsedURL, error)   // host ∈ discord.com|discordapp.com (any subdomain? no: exact or www/ptb/canary), path /api/webhooks/<digits>/<token>, https only; ErrNotAWebhook otherwise
type Sealer struct{ key []byte }; func NewSealer(key []byte) (*Sealer, error); func (s *Sealer) Seal(plain string) ([]byte, error); func (s *Sealer) Open(box []byte) (string, error) // AES-256-GCM, random nonce prefixed
func (s *Store) Save(ctx, userID int64, kind string, p ParsedURL, tokenBox []byte, events map[EventName]bool) error  // upsert on (user_id, kind), resets failures/disabled/last_error
func (s *Store) Get(ctx, userID int64, kind string) (Webhook, bool, error)             // never the token
func (s *Store) TokenBox(ctx, userID int64, kind string) ([]byte, error)              // the job's read only
func (s *Store) Delete(ctx, userID int64, kind string) error
func (s *Store) RecordSuccess(ctx, userID int64, kind string) error                    // failures=0, last_error=''
func (s *Store) RecordFailure(ctx, userID int64, kind string, lastErr string) (disabled bool, err error) // failures+1; at MaxFailures sets disabled_at
func (s *Store) Reenable(ctx, userID int64, kind string) error                        // on re-save (Save does it) — expose only if the spec adds a Re-enable control
```
Migration: `notification_webhooks(user_id bigint references users(id) on delete cascade, kind text, channel_id text not null, token_box bytea not null, events jsonb not null default '{}', failures int not null default 0, last_error text not null default '', disabled_at timestamptz, created_at, updated_at, primary key (user_id, kind))`.

- [ ] Tests: `TestParseWebhookURL` table (accepts `https://discord.com/api/webhooks/123/abc`, `discordapp.com`, `ptb.discord.com`/`canary.discord.com` — ruling: the three Discord hosts and `www` only; refuses http, other hosts, `/api/webhooks/` without a token, query strings, a path with extra segments, an id that is not digits); `TestSealerRoundTripsAndRefusesAForeignBox`; `TestSaveUpsertsAndResetsFailures`; `TestGetNeverReturnsTheToken` (Get's struct has no token field; TokenBox is separate); `TestRecordFailureDisablesAtTheFifthAndSuccessResets`; `TestDeleteCascadesWithTheUser`. Commit `feat(notify): the webhook store, URL rule and sealer`.

### Task 2: Routes and `/v1/me`

**Files:** `api/internal/notify/handler.go` (`Service{Store, Sealer, Entitlements Entitlementer, Poster Poster, Log}`; `Mount`: `PUT /v1/notifications/discord`, `DELETE /v1/notifications/discord`, `POST /v1/notifications/discord/test`, all `auth.RequireSession`, mounted only when the sealer exists), `dto.go` (the block DTO per §10: `channel_id, events{sim_finished, upload_parsed}, disabled_reason, last_error, failures`), `api/internal/auth/handler.go` (`Me.Notifications *NotificationsView` via a `NotificationsView` interface declared in `auth` + an adapter in `api/cmd/api/notifyview.go`, same pattern as lane 3's `SimsView`), `api/cmd/api/main.go` wiring, `api/openapi.yaml`, `api/README.md`.
Rules: 402 `premium_required` without `FeatureNotifications`; `PUT` body `{webhook_url?, events}` → 400 `invalid` with the spec's message for a bad URL (never echoing the input); with a URL the handler calls `Poster.Lookup(ctx, url)` (Discord `GET /webhooks/{id}/{token}`, 5 s) to read the channel id and prove the webhook exists (422 `invalid` with the spec's message when it fails, per spec §10; nothing stored); without a URL it updates the events of the stored row (404 when none); answers the block DTO; `POST …/test` posts the spec's test embed through `Poster` and answers 204 or 502 `upstream` with Discord's status text (truncated, no URL); `disabled_reason` is `"failed five times"` when `disabled_at` is set, else `"plan lapsed"` when the entitlement is false, else null.
- [ ] Tests: save → `/v1/me` shows the channel id and events and no URL anywhere in the body (assert the token string is absent); bad host 400; non-payer 402; test posts once through a fake poster and 502 on failure; delete → `/v1/me` block null. Commit `feat(notify): save, test and delete the Discord webhook; /v1/me carries its state`.

### Task 3: Delivery from the jobs

**Files:** `api/internal/notify/deliver.go` (`Deliverer{Store, Sealer, Entitlements, Poster, Log}`; `Deliver(ctx, userID, event EventName, embed Embed) error` → reads the row, skips when absent/disabled/event off/lapsed, opens the token, posts with the 5 s timeout, records success/failure, logs only the user id, event and status), `embed.go` (`Embed{Title, Fields [][2]string, URL}`; `SimFinishedEmbed(spec, dps, title, link)`, `UploadParsedEmbed(reportTitle, fights, link)` from a Go copy table matching spec §6; Discord JSON `{embeds:[{title, url, fields:[{name,value,inline}], footer}]}`), `discord.go` (`Poster` interface — `Post(ctx, url string, embed Embed) error` and `Lookup(ctx, url string) (channelID string, err error)` — plus `HTTPPoster` with `http.Client{Timeout: 5s}`, the `mail/resend.go` pattern; errors carry Discord's status and a sanitised body; a `FakePoster` for tests recording calls and scripting failures/429s), `api/internal/sims/job.go` (after `Finish`: `d.Notify.SimFinished(ctx, userID, …)` when `JobDeps.Notify != nil`), `api/internal/parse/parse.go` (after `SetStatus(StatusComplete)`: the upload-parsed event for the report's owner when `Deps.Notify != nil`), `api/cmd/api/main.go` (`runSim`/`runParse` build the deliverer when the key is configured).
- [ ] Tests: `TestDeliverPostsOnceAndRecordsSuccess`, `TestDeliverSkipsWhenOffLapsedOrDisabled`, `TestFifthFailureDisablesAndTheNextEventIsSkipped`, `TestEmbedsMatchTheCopyTable` (golden JSON), a sims job test and a parse job test with the fake poster proving one post each. Commit `feat(notify): the jobs post one embed per event; five failures switch the webhook off`.

### Task 4–6: Web (filled from the experience spec)

- Task 4: `web/src/lib/account/notifications-api.ts` (types mirroring §10; `saveDiscordWebhook`, `testDiscordWebhook`, `deleteDiscordWebhook` through `requestEnvelope`, each followed by `invalidateMe()`), `notifications-copy.ts` (every §6 string), `notifications.ts` (view model: block state from `me.notifications` and `me.entitlements.notifications` → `free | empty | saved | disabled | lapsed`; URL pre-check mirroring the API rule for the inline message), tests; the `test-support` route stubs.
- Task 5: `web/src/components/account/AccountDiscordBlock.svelte` (+ small children under `components/account/discord/` per §3; a change updates `me` in place with `setQueryData` and one background `refetchMe`, never `session.refresh()`), mounted in `Account.svelte` where §3.1 says; `AccountPremiumBlock`'s notifications row goes live (`On`/`Off` per §6; `shipsMonth` retired entirely); `/premium`'s Discord line loses its month; tests per state.
- Task 6: `web/tests/e2e/account-discord.spec.ts` (save valid → channel id shown; invalid host → message and no save request; test → sent / failed; disabled after five failures → reason + last error + re-save; lapsed → why; free member → sentence; URL never in the DOM), changelog, lane close (full suites, captures per §9, ux-designer + wow-player, merge, deploy; after deploy: set `NOTIFY_ENCRYPTION_KEY` in Secret Manager and the Cloud Run service + both jobs — a controller step, key generated locally and never printed).

## Rulings made by the plan (owner to check)

- Encryption over hashing: the job must recover the token to post, so AES-256-GCM with a Secret Manager key; the channel id stays in clear for display.
- Hosts accepted: exactly `discord.com` and `discordapp.com` (the experience spec §11.11); https only.
- "Plan lapsed" is computed at read and delivery time from `Can(FeatureNotifications)`, never stored, so a renewal resumes posts without a re-save.
- Re-enable after five failures is `Test and switch on` (experience spec §11.9): a successful test POST resets the counter and clears `disabled_at`; a re-save of the URL does the same.
- A manual test and a Discord 429 do not count toward the five failures (§11.12); every other failed event post does.
- The channel id is not in the URL (§11.1): on save the API calls Discord's `GET /webhooks/{id}/{token}` once (5 s timeout) to read `channel_id` and prove the webhook exists; a failure to read it refuses the save with the spec's message; the page shows `channel <id>`.
- `PUT /v1/notifications/discord` takes `webhook_url` as optional: absent means "keep the stored URL, change the events only" (the page never holds the URL).
- `POST /v1/notifications/discord/test` is rate-limited to 5 per minute per account (429 beyond), through `httpx.RateLimitByKey` like `POST /v1/sims/run`.
- Discord's error text is sanitised by the API before it reaches the page or the row: any URL or token-like segment stripped, truncated to 200 characters.
- The embeds never mention anyone (`allowed_mentions: {parse: []}`) and member-supplied text is escaped (§11.6); failed sims do not post (§11.4); private uploads post with a `Visible to` field (§11.5).
- The spec's twenty §11 decisions are otherwise accepted as written; the Premium row wording (`On · 2 events`, `Off · failed five times`, `Paused · plan lapsed`) is the spec's; `shipsMonth` is deleted with this lane.
