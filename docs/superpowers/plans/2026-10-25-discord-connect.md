# Discord connect (one click) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A Premium member connects a Discord channel in one click: our API sends them to Discord's own server-and-channel picker (OAuth2 `webhook.incoming`), receives a fresh incoming webhook on the way back, and stores it exactly as a pasted one. The paste form stays as the secondary path.

**Architecture:** Two new API routes (`connect` → 302 to Discord with a cookie-held random state, the Battle.net pattern; `callback` → exchange the code form-encoded, read `webhook.{url, channel_id, guild_id}`, seal and save with `source = 'connected'`, 302 back to `/account?discord=<outcome>#discord`). Migration 0048 adds `source` and `guild_id`. The poster gains `Delete` so a replaced or removed connected webhook is deleted from Discord best effort. The web shows a connect state when `/v1/me` says `connect: true`, handles the three return flags once, shows `Connected to a Discord channel` + `Open in Discord` for a connected webhook, and switches remove/replace wording by source. Feature off (no Discord application configured) = the shipped behaviour, untouched.

**Tech Stack:** Go (net/http, pgx), Svelte 5, vitest, Playwright.

**Spec:** `design/specs/2026-10-25-discord-notifications.md` as amended 2026-10-25 (connect lane): amendment log, §1, §3.2, §4.1–4.3, §5, §6.1, §7, §8 (13–16), §9 (captures 24–31), §9.1 (`connect` hint), §10, §11 (21–25). Research: `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/discord-oauth-webhook-research.md` (copy its key facts into the Task 1 report).

## Global Constraints

- Discord application: Client ID `1563973783985852506` (public; config `DISCORD_CLIENT_ID`), Client Secret only in Secret Manager as `DISCORD_CLIENT_SECRET` (never in chat, logs, files or tests — tests use a fake); redirect URI `${API_BASE_URL}/v1/notifications/discord/callback` (production `https://api.foreversixty.gg/...`; local `http://127.0.0.1:8080/...`). `connect` is on only when both values are set AND the notification key exists.
- The OAuth `code`, `access_token`, `refresh_token`, the webhook URL/token and the sealed box never reach a response, redirect, log line or test output. The access/refresh tokens are discarded unread.
- `state`: random (`auth.NewSessionID()`), held in an httpOnly `SameSite=Lax` cookie `fs_discord_state` scoped to the callback path, 10 minutes, compared with `auth.SameToken`, cleared on use; the callback also requires the same session cookie (RequireSession) so a stolen link cannot bind a webhook to another account.
- Discord's API is called only through the `notify` package's poster/OAuth client with 5 s timeouts and no redirects; every failure maps to `denied` or `failed`, never a 500 page for the member.
- Copy only from the copy modules, verbatim from §6.1; the connect button is plain text in our own style (no Discord logo, wordmark or blurple; acceptance capture 31 greps the build).
- Tests: API `-race -p 1` with `TEST_DATABASE_URL`; web vitest for touched files, `npm run check`; e2e `account-discord`, `account-reserve`, `account-billing`, `premium`, `layout`; the account page is a shared surface — lane close runs the full web suites.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. A forged or replayed callback (wrong state, reused state, another user's session, a code for a different redirect) stores nothing and lands on `failed` (Task 1 tests).
2. A member who cancels in Discord lands on `denied` with nothing stored and their previous webhook (if any) untouched (Task 1 + Task 3).
3. Replacing a connected webhook deletes the OLD one from Discord only after the new one is saved; a pasted webhook is never deleted from Discord; a lapsed member's remove deletes nothing in Discord (Task 1).
4. With the application unconfigured, `/v1/me` says `connect: false`, `GET /connect` is 404, and the page is pixel-identical to today (Task 2 + Task 3 capture).
5. The return flag is handled once: a reload after `?discord=connected` shows no second status line (Task 2 + Task 3).

---

### Task 1: API — connect and callback routes, source and guild id, Discord-side delete

**Files:**
- Create: `api/internal/db/migrations/0048_notification_webhook_source.up.sql` (`alter table notification_webhooks add column source text not null default 'pasted' check (source in ('pasted','connected')), add column guild_id text;`) / `.down.sql`; `api/internal/notify/oauth.go` (`OAuthClient` interface: `AuthURL(state string) string`, `Exchange(ctx, code string) (OAuthWebhook, error)`; `OAuthWebhook{ID, Token, URL, ChannelID, GuildID, Name string}`; `HTTPOAuthClient{ClientID, ClientSecret, RedirectURI, HTTP}` with form-encoded POST to `https://discord.com/api/oauth2/token`, 5 s, no redirects, body limit, errors with the status only (never the body — it could echo the code); `FakeOAuthClient`), `api/internal/notify/connect_handler.go` (the two routes).
- Modify: `api/internal/config/config.go` (`DiscordClientID`, `DiscordClientSecret` from env; both-or-neither, else startup error; README rows), `notify.go` (`Webhook.Source Source`, `Webhook.GuildID *string`; `Source` type with `SourcePasted`/`SourceConnected`), `store.go` (`Save(ctx, userID, kind, channelID, guildID *string, name *string, box, events, source)`; `Get` returns both; `TokenBox` also returns `source` + the parsed id/token? — no: `Delete`-from-Discord needs the URL, which `Open(box)` gives), `discord.go` (`Poster.Delete(ctx, url) error` → `DELETE` on the webhook URL, 2xx/404 = ok, 5 s; Fake records it), `handler.go` (PUT with a URL stores `source = pasted` + the lookup's `guild_id`; `DELETE` route: when the stored source is connected AND the member is entitled, open the box and `Poster.Delete` best effort BEFORE the row is removed; mount the two connect routes only when the OAuth client exists; `Service.OAuth OAuthClient` nil = connect off), `dto.go` + `auth` DTO (`connect bool` on the `NotificationsBlock` level — `{"connect": true, "discord": …}` — and `source`, `guild_id` on the block; `name` is null for connected), `api/cmd/api/notifyservice.go` (build the OAuth client from config), `openapi.yaml`, `README.md` (env rows + a short "Discord application" runbook: name, icon, redirect URIs, where the secret lives).
- Test: `connect_handler_test.go` (`TestConnectRedirectsToDiscordWithAStateCookie`, `TestConnectIs404WhenUnconfigured`, `TestConnectRedirectsToFailedWithoutEntitlement`, `TestCallbackStoresAConnectedWebhookAndRedirectsConnected` (sealed URL opens to the fake's url; `source = connected`; `name` nil; `guild_id` stored; no token/code in the Location header or logs), `TestCallbackWithoutCodeIsDenied`, `TestCallbackAccessDeniedIsDenied`, `TestCallbackWithWrongStateIsFailedAndStoresNothing`, `TestCallbackStateIsSingleUse`, `TestCallbackExchangeFailureIsFailed`, `TestReconnectDeletesTheOldWebhookFromDiscordAfterSaving`, `TestRemoveDeletesAConnectedWebhookFromDiscordButNotAPastedOne`, `TestLapsedRemoveDeletesNothingInDiscord`, `TestMeCarriesConnectSourceAndGuildID`), plus `oauth_test.go` (the HTTP client against `httptest`: form body fields, basic auth or body credentials per the research, 5 s timeout, a non-2xx → error without the body) and the migration test (an old row reads `pasted`).

- [ ] Commit `feat(notify): connect a Discord channel in one click — OAuth webhook.incoming, source and guild id, Discord-side delete`.

### Task 2: Web — connect state, return flags, connected destination

**Files:**
- Modify: `web/src/lib/account/api.ts` (`Me.notifications.connect: boolean`, `DiscordNotifications.source: 'pasted'|'connected'`, `guild_id: string | null`), `notifications-copy.ts` (every amended §6.1 string: the connect intro/explanation, `Connect a Discord channel`, `Opening Discord…`, `Paste a webhook URL instead`, `Connect a Discord channel instead`, the three return lines, `Connected to a Discord channel`, `Open in Discord`, `Connect a different channel` + its note, the two remove questions), `notifications.ts` (`connectAvailable(me)`, `discordReturnFlag(search) : 'connected'|'denied'|'failed'|null`, destination model by source), `layout.ts` + `Base.astro` + `global.css` (seventh `data-discord` value `connect` at 320/340 — the hint expression: `form` state with `connect: true` → `connect`; keep the parity test and the CSS/table test green), `web/src/components/account/discord/*` (`DiscordConnect.svelte`: the connect state; the paste form's back link; `DiscordSummary`: the connected destination + `Open in Discord` link; `DiscordActions`/`DiscordSaved`: replace link by source with the note; remove question by source; `AccountDiscordBlock.svelte`: the view switch (connect ↔ paste), the return-flag handling on mount (status line + focus + `history.replaceState` keeping `#discord`; after `connected` mark `/v1/me` stale via `fetchMeFresh`), the `pageshow` reset of `Opening Discord…`), fixtures (`test-support/notifications-api.ts`, `tests/e2e/support/discord-notifications.ts`: `connect`, `source`, `guild_id`).
- Test: copy verbatim table extended; `notifications.test.ts` (flag parsing, connect availability, hint `connect`); component tests per new state; `layout-reserve.test.ts` for the seventh value.

- [ ] Commit `feat(web): Connect a Discord channel — the connect state, the return flags, the connected destination`.

### Task 3: E2E, captures, changelog, runbook

- [ ] `tests/e2e/account-discord.spec.ts`: connect state when `connect: true` (button href, explanation, the paste disclosure and its back link); connect absent when `connect: false` (pixel-identical paste form — compare a screenshot against the first lane's capture or assert the DOM is unchanged); the three returns by visiting `/account?discord=<flag>#discord` (status line words, focus target, URL cleaned, no second line on reload); saved connected (`Connected to a Discord channel`, `Channel …`, `Open in Discord` href from the fixture's guild/channel ids, `Connect a different channel` + note); the two remove questions; `account-reserve.spec.ts` seventh state. Captures 24–31 via the capture spec (controller runs it). Changelog: one sentence appended to the 2026-10-25 entry. `web/README.md` or `api/README.md`: the Developer Portal runbook (name, icon, redirect URIs, secret in GSM, `DISCORD_CLIENT_ID` public).
- [ ] Commit `test(web): Discord connect e2e, captures, changelog`.

## Rulings made by the plan (owner to check)

- The state lives in a cookie (the Battle.net pattern) rather than a signed token; "single use" = the cookie is cleared on the first callback.
- The old connected webhook is deleted from Discord only after the new one is saved (a failed exchange keeps the old one posting).
- `guild_id` is stored for pasted webhooks too (the lookup returns it), so `Open in Discord` appears for both sources when known.
- No account-deletion handler exists today; a cascade cannot call Discord, so a deleted account leaves its connected webhook in Discord — noted in the runbook as a follow-up when account deletion ships.
- Events are not chosen before the redirect; a new connection starts on the defaults.
