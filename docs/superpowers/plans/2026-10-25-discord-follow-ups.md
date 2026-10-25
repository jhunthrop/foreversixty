# Discord notifications follow-ups Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Close every item deferred from the Discord notifications lane: the webhook's name on the destination line, a private-logs opt-in, the clearer switched-off and warning copy, numbered how-to steps, the Premium row's new words, `Run time`/`Boss fights`/the character name in the embeds, the account page's loading reserve, gold checkboxes site-wide, and the in-place `/v1/me` patch surviving an inflight read.

**Architecture:** One API migration (0047) adds `webhook_name` and the third event flag; the poster's lookup returns Discord's webhook `name`; the parse job gates private logs on the flag. The web mirrors the new contract in its lib, copy and components; the account page's loading state becomes the ready page's skeleton at real positions, keyed by a third Base.astro pre-paint block and a stored pair of column heights; one `global.css` rule makes every checkbox gold; `query.ts` discards an inflight answer that predates a `setQueryData` write.

**Tech Stack:** Go (pgx v5), Svelte 5, vitest, Playwright.

**Spec:** `design/specs/2026-10-25-discord-notifications.md` as amended 2026-10-25 (follow-up lane): §3.2, §3.4, §3.7, §4, §6.1, §6.2, §7, §9.1, §9.2, §10. Mock: `design/mocks/gen_discord_notifications.py`.

## Global Constraints

- Copy only from the copy modules, verbatim from §6.1 (web) and §6.2 (Go `embedCopy`), mirrored by verbatim tests.
- The webhook URL, token, box and webhook id never reach a response, a log, the DOM or storage; the webhook `name` is member-controlled text: trimmed, control characters removed, ≤ 80 runes stored, escaped and ≤ 40 characters shown.
- `/v1/me`'s `notifications.discord` gains `name` (string|null) and `events.upload_private`; `PUT` validates all three event keys; existing rows get `upload_private = true` once (migration), new rows default `false`.
- The parse job posts a private report only when `upload_parsed` and `upload_private` are both on.
- Premium row: `Off · needs attention`, `Off · Premium`, no row at all when the deployment has notifications off; the event count is `sim_finished + upload_parsed`; non-breaking spaces per §3.7.
- The account loading state reserves the ready layout (§9.1); warm CLS ≤ 0.05 at 390 and ≤ 0.03 at 1440; cold recorded.
- Tests: API `TEST_DATABASE_URL`, `-race -p 1`; web `npm run check`, vitest for touched files; e2e `account-discord`, `account-billing`, `premium`, `layout`; the account page and `global.css` are shared surfaces — lane close runs the full web suites.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. A webhook saved before this lane keeps posting private logs (`upload_private` backfilled true); a new webhook does not until the member turns the third box on (Task 1 migration + delivery tests).
2. A webhook name containing markdown, HTML or a URL renders as text, never as a link, never longer than 40 characters (Task 1 sanitiser, Task 3 escaping test).
3. The loading skeleton never shows for a visitor without the session hint; a warm visitor's footer moves ≤ 8 px (Task 3b, Task 4 CLS spec).
4. A second action within the first action's background revalidation keeps the newer state (Task 2 `query.ts` test).
5. Every checkbox on the site is gold, its disabled state muted, its focus ring visible; no local override draws a different box (Task 3b grep + capture).

---

### Task 1: API — webhook name, private-logs flag, embed words

**Files:**
- Create: `api/internal/db/migrations/0047_notification_webhook_name_private.up.sql` / `.down.sql`
- Modify: `api/internal/notify/notify.go` (`Webhook.Name *string`, `Events.UploadPrivate bool`), `store.go` (`Save(ctx, userID, kind, channelID string, name *string, box []byte, events Events)`; `SaveEvents` unchanged signature), `discord.go` (`Poster.Lookup(ctx, url) (LookupResult, error)` with `LookupResult{ChannelID, Name string}`; `HTTPPoster` decodes `name`; `FakePoster.LookupName`), `handler.go` (`putRequest.Events` requires `upload_private *bool`; `cleanWebhookName(raw string) *string`), `dto.go` + `api/internal/auth` DTO (`Name *string`, `Events.UploadPrivate`), `embed.go` (`SimTitleNamed "%s's %s sim finished"` with the escaped character name, `SimTookName "Run time"`, `LogFightsName "Boss fights"`, value `9 · 7 kills` and the field omitted when the count is 0, `TestPrivateClause ", private logs included."`), `deliver.go` (`SimFinished.Character string`; `UploadParsed.Private bool`; `Deliverer` skips a private upload unless `upload_private`), `api/internal/sims/job.go` (`Character: res.Request.Character.Name`), `api/internal/parse/parse.go` (`Private: report visibility == private`), `api/openapi.yaml`, `api/README.md`.
- Test: the package tests beside each file; goldens in `embed_test.go` updated.

**Interfaces:** Produces the §10 contract above; consumed by Task 2's types and Task 4's stubs.

- [ ] Migration: `alter table notification_webhooks add column webhook_name text;` `update notification_webhooks set events = events || '{"upload_private": true}'::jsonb;` down: drop the column (the flag stays in the jsonb; harmless). Test `TestExistingWebhooksKeepPostingPrivateLogs` (a row saved before the migration's backfill has `UploadPrivate == true`; a new `Save` with `Events{SimFinished:true, UploadParsed:true}` reads back `UploadPrivate == false`).
- [ ] `TestLookupReturnsTheWebhookName`; `TestSaveStoresACleanName` (trim, `\x00`–`\x1f` removed, 81 runes → 80, empty → nil); `TestPutWithoutUploadPrivateIs400`; `TestMeCarriesNameAndThreeEvents`; `TestEventsOnlyPutKeepsTheName`.
- [ ] `TestEmbedsMatchTheCopyTable` goldens: `Run time`, `Boss fights` `9 · 7 kills`, omitted at 0; sim title named/unnamed; test message with the private clause. `TestPrivateUploadPostsOnlyWhenOptedIn` (deliver), a parse job test with a private report and the flag off → no post; on → one post with `Visible to: Only you`.
- [ ] Commit `feat(notify): the webhook's name, a private-logs opt-in, Run time and Boss fights in the embeds`.

### Task 2: Web lib — types, copy, state, reserve table, inflight fix

**Files:**
- Modify: `web/src/lib/account/api.ts` (`DiscordNotifications.name: string | null`, `DiscordEvents.upload_private`), `notifications-copy.ts` (every amended §6.1 string: the three how-to steps, the four-line switched-off notice incl. `Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.`, destination `Posts through “%s”` / `Posts to a Discord channel` / `Channel %s`, `Include private logs` + its note, the warning's `Send a test message to check.` sentence, Premium row values `Off · needs attention`, `Off · Premium` with the §3.7 non-breaking spaces), `notifications.ts` (`premiumRowValue` counts two events; `off` → no row: a new `premiumRowVisible(me)`), `web/src/lib/data/query.ts` (an `Entry.generation` counter bumped by `setQueryData`; `revalidate` captures it before `load()` and, when it changed by the time the answer arrives, drops that answer — keeps the newer data, clears `inflight`, returns the current data; a background re-read is NOT started automatically), `web/src/lib/account/layout.ts` (new: the §9.1 tables — header row, left floor, rail floors at 1440 and 390, the Discord floors by state; `ACCOUNT_RESERVE_KEY = 'fs.accountReserve'`, `readAccountReserve()`, `writeAccountReserve(bucket, left, rail)` on `lib/layout-reserve.ts`'s helpers; `discordHintState(me): 'none'|'free'|'form'|'saved'|'off'|'lapsed'` — the pure function the Base.astro expression mirrors), `web/src/test-support/notifications-api.ts` + `web/tests/e2e/support/discord-notifications.ts` (`name`, `upload_private` in the fixtures; `CHANNEL_ID` stays distinct from `WEBHOOK_ID`).
- Test: the files' tests; `query.test.ts` gains `TestAnInflightAnswerNeverOverwritesANewerWrite` (start a slow `query`, `setQueryData` a newer value, resolve the slow load, assert the newer value stands and `status === 'ready'`).

**Interfaces:** Consumes Task 1's contract (build from §10 if Task 1 is not merged yet). Produces `premiumRowVisible`, `discordHintState`, `layout.ts` constants for Tasks 3a/3b.

- [ ] Verbatim copy table test against §6.1 (character for character, incl. the nbsp characters).
- [ ] `blockState` unchanged keys; `premiumRowValue` for a webhook with `upload_private: true` and both events → `On · 2 events`.
- [ ] Commit `feat(web): Discord follow-up copy, name and private flag types, account reserve table, inflight-safe query writes`.

### Task 3a: Web components — the amended block

**Files:** `web/src/components/account/AccountDiscordBlock.svelte` and `components/account/discord/*` (paste form with the three numbered steps above the field and the error line directly under the input; the destination's two lines with the escaped name ≤ 40 chars; the three checkbox rows with the indented, dependent `Include private logs` and its always-shown note; the switched-off notice's four lines with `Use a different webhook` inside it; the warning's two paragraphs; the replace link hidden while the remove question is open), `AccountPremiumBlock.svelte` (no notifications row when `premiumRowVisible` is false), `web/src/lib/account/discord-actions.ts` (events-only save sends three keys).
Owns ONLY those files plus their tests. Does not touch `Account.svelte`, `Base.astro`, `global.css`, `layout.ts`.

- [ ] Component tests per state (static render) + action tests (three-key events body; the private box disabled while the log box is off, value kept).
- [ ] Commit `feat(web): Discord block — webhook name, private-logs row, numbered steps, clearer switched-off words`.

### Task 3b: Web shared surfaces — loading reserve and gold checkboxes

**Files:** `web/src/components/Account.svelte` (the loading branch per §9.1: header row reserve, `account-left-skeleton`, `account-rail-skeleton` with `account-devices/you/plan/discord-skeleton`; test ids `account-left-col`/`account-rail-col` on the ready columns; the height writer two frames after mount and on `pagehide`; sign-out removes the key), `web/src/layouts/Base.astro` (the third inline block: `/account` only, `data-session="1"` only; sets `data-discord` and `--account-left-reserve`/`--account-rail-reserve` from `fs.accountReserve` by the `lg`/`stack` bucket), `web/src/styles/global.css` (the `.session-reserve` interplay: hint-less visitors stay collapsed; the §9.2 checkbox rules in the base layer), `web/src/lib/account/layout.ts` (measured floors at 1440 AND 390 replace the targets; `SIGNED_OUT_MIN_H` by the same arithmetic).
Owns ONLY those files plus their tests. Does not touch the Discord components.

- [ ] Unit test: `discordHintState` fixtures vs the Base.astro expression (extract the expression into a tiny string the test evals, or mirror it as a function in `layout.ts` the inline script copies verbatim with a comment pointing both ways).
- [ ] Grep every `type="checkbox"` surface for local `accent-*`/`appearance-none`/hand-drawn boxes; remove non-gold overrides; keep the 44 px rows.
- [ ] Measure the floors on the fixture build at 1440 and 390 and write them into `layout.ts`.
- [ ] Commit `feat(web): the account page reserves its ready layout while loading; gold checkboxes site-wide`.

### Task 4: E2E, captures, changelog

- [ ] `web/tests/e2e/account-discord.spec.ts`: the name line, the private row (disabled while the log box is off; the three-key PUT), `Off · needs attention`, `Off · Premium`, no row when the deployment is off, the four-line notice, the steps; `web/tests/e2e/account-reserve.spec.ts` (new): §9.1 acceptance — warm and cold, six `data-discord` states, 390 and 1440, CLS and footer movement, with `/v1/me` delayed 400 ms; a checkbox capture-style assertion that `accent-color` resolves to the gold token on three surfaces.
- [ ] Changelog: one sentence appended to `web/src/content/changelog/2026-10-25-discord-notifications.md` (private logs are opt-in; the webhook's name shows on the account page).
- [ ] Captures 14–23 (spec §9) into `design/reviews/2026-10-25-discord-notifications/` via the committed capture spec (`web/tests/captures/discord-notifications/`), then ux-designer and wow-player re-checks (controller).
- [ ] Commit `test(web): Discord follow-ups e2e, the account reserve check, changelog line`.

## Rulings made by the plan (owner to check)

- Existing webhooks keep posting private logs (the backfill sets the new flag true once); only new webhooks default to off — nobody's behaviour changes silently (spec §10).
- The inflight fix is "drop the stale answer", not "re-read": the page already fires one background read after every patch, so a second read would double traffic for a ~200 ms window.
- The loading reserve rebuilds the whole `/account` loading state (the designer's scope), not a Discord-only reserve.
- The changelog entry of the day is amended rather than a second entry written.
