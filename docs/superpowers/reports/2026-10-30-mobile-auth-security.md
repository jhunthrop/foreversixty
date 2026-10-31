# Security review: native app sign-in (wt/mobile-auth), 2026-10-30

Scope: `git diff origin/main...HEAD` (7f75a8fb3, b8b653108, a34b31128). Read only.
Tests run against a scratch database: `internal/db`, `internal/auth`, `internal/config` all pass.

**Verdict: MERGE.** The flag defaults off. No CRITICAL or HIGH findings. Fix the MEDIUMs before `APP_SIGN_IN=true`.

## What holds

- **PKCE.** `code_challenge_method` must be exactly `S256`, so `plain` and a missing method get 400. The challenge must be 43 base64url characters, checked by regex and by a database CHECK. The verifier follows RFC 7636 §4.1 (`[A-Za-z0-9._~-]{43,128}`). The S256 recompute is compared in constant time (`SameToken`).
- **State.** The Battle.net state is 256 random bits. It is bound to the browser by the HttpOnly, SameSite=Lax `fs_bnet_state` cookie, which the callback checks in constant time. It is also bound to the flow row by `state_hash`. `IssueAppCode` only fires when `code_hash is null and expires_at > now()`, so a replayed Battle.net callback mints nothing.
- **App/web separation.** The cookie's `next` is `"app"` only when `appStart` sets it. `safeNext` always returns a path that starts with `/`, so a web sign-in cannot become an app flow. The app branch returns before `StartSession`, so an app callback never sets a web session. A web flow that reaches the app branch with no flow row gets a 400.
- **Code.** The one-time code is 256 bits and only its SHA-256 is stored. It is looked up by hash, so the lookup leaks no timing. It is spent atomically (`UPDATE … WHERE used_at IS NULL AND code_expires_at > now()`) *before* the verifier check, so a wrong verifier burns it, and of two racing exchanges only one wins. It expires after 2 minutes. Unknown, spent, expired and wrong-verifier codes all get the same `invalid_grant` answer.
- **Redirects.** The callback redirects to `PublicBaseURL + "/app/auth"`, a fixed host and path. `code` and `state` are passed through `url.Values.Encode`, and the app's state is limited by regex to 16–128 unreserved characters. I found no open redirect or header injection.
- **Tokens.** The token is `fsd_` plus 32 base32 characters (160 bits from crypto/rand), stored as SHA-256 and shown once. `kind` is set when the device is created, and the connect/pair paths still insert companions, so I found no route that turns a companion token into an app token.
- **Revocation.** It takes effect on the next request, because `DeviceByToken` checks `revoked_at is null` on every call. `appLogout` revokes only `ActorFrom(ctx).DeviceID`. `DELETE /v1/devices/{id}` is scoped to the owner.
- **CSRF.** Cookie sessions are unchanged: bearer requests skip double-submit, and that is safe because a bearer header is not ambient. CORS sends `Allow-Origin` only for the site origin.
- **Discord.** Linking still refuses every device actor in `linkStart`/`linkCallback`, and the install and notify callbacks do too. Widening `RequireSession` does not open them.
- **Logging.** The request log records `r.URL.Path` only, with no query. The code travels in the JSON body to the API and in the query string to the static site (Cloudflare Pages, no analytics, default referrer policy).

## MEDIUM

### M1. An attacker-started app flow puts a victim's code in the victim's browser
`api/internal/auth/app_auth.go:89` (`appStart`) and `:119` (redirect)

Scenario:
1. The attacker makes their own verifier V and challenge C.
2. They send the victim `https://api.foreversixty.gg/v1/auth/app/start?platform=ios&code_challenge=C&code_challenge_method=S256&state=xxxxxxxxxxxxxxxx`.
3. If the victim is signed in to Battle.net and has already consented to this client ID, which the web flow shares, Battle.net returns at once. The state cookie checks out because it is the victim's own browser.
4. A code for the *victim's* account lands at `https://foreversixty.gg/app/auth?code=…`.
5. On a desktop, or a phone without the app, there is no page there: the site's 404/catch-all renders and the code stays in the address bar and history.

PKCE does not help, because the attacker holds V. The attacker still needs the code. Ways to get it: talk the victim into pasting the URL ("send me the link you landed on"), any future script or XSS on the site, a shared machine's history, or a browser extension. With the code, the attacker gets a non-expiring app token with full account powers (see M2).

Fix:
- Add a real `/app/auth` page. Serve it with `Referrer-Policy: no-referrer` and `Cache-Control: no-store`. Have it run `history.replaceState` at once to drop `code` from the URL, and say "open this in the Forever Sixty app; if you did not start a sign-in, ignore this page." Load no third-party resources on it.
- Optionally show a confirmation step in the API before the Battle.net redirect for app flows ("Sign in to the Forever Sixty app on iPhone?"). That way a silent Battle.net consent cannot finish the flow without a user action.
- Optionally shorten `AppCodeTTL` to 60 s.

### M2. App tokens never expire, yet now hold every session-only power
`api/internal/auth/session.go:249-252` (`RequireSession`), `api/internal/auth/store.go:445` (`DeviceByToken`)

Web sessions expire after 30 days (`SessionTTL`). Device tokens have no idle or absolute expiry. With this change an app token reaches about 60 `RequireSession` routes, including:
- billing checkout and portal
- guild upgrade and transfer (`billing/guild_routes.go`)
- guild claim, resolve and settings, raid officer actions, loot awards and CSV exports
- approving companion connects (`/v1/devices/connect/approve`) and pairing new companions
- revoking every other device
- moderator powers for moderator or admin accounts (`guilds/contest.go`, `officer_gate.go`)

Role and officer checks still run inside the handlers, so this is not a privilege bypass. But one token leaked from a phone backup, a log line, or the M1 path is valid for the account's lifetime.

Scenario: a moderator's phone token leaks from an unencrypted Android backup. Six months later it still resolves guild claim contests and approves a new companion on the account, and nothing expires it.

Fix:
- Give app devices an idle expiry, for example `last_seen_at < now() - 90 days`, and an absolute expiry. Check them in `DeviceByToken` for `kind = 'app'`.
- Or issue short-lived access tokens with a rotating refresh token.
- Consider keeping moderator-only actions and billing writes on browser sessions (`RequireBrowserSession`) until the apps need them.
- Make sure the apps store the token in Keychain/Keystore with `ThisDeviceOnly` and keep it out of backups.

## LOW

### L1. `DELETE /v1/sessions` and `POST /v1/auth/logout` answer success to an app token and revoke nothing
`api/internal/auth/handler.go:390`, `:382`

With no cookie, `EndSession` is a no-op that returns 204/200. If an app client calls the generic sign-out, it will believe it signed out while the token stays live. Fix: for `a.IsApp()`, revoke the device in these handlers too, or answer 400 and point to `/v1/auth/app/logout`.

### L2. No cap on app devices per account; every sign-in adds a row
`api/internal/auth/app_auth.go:174` (`issueAppDevice`)

Repeated sign-ins and reinstalls pile up live tokens. The device list grows, and old tokens on wiped phones stay valid (this makes M2 worse). Fix: cap live app devices, say 10, and revoke the oldest. Or reuse a device row when the app sends a stable install id.

### L3. The down migration turns app devices into companions
`api/internal/db/migrations/0059_app_sign_in.down.sql`

Dropping `kind` leaves app tokens valid, and they then read as companions. The `acks` "paired companion" check and companion-only routes will treat phones as companions. Fix: `update devices set revoked_at = now() where kind = 'app' and revoked_at is null;` before the column drop.

### L4. Placeholder Team ID and fingerprint in the AASA and assetlinks files
`web/public/.well-known/apple-app-site-association`, `web/public/.well-known/assetlinks.json`, `web/app-links.json`

With `TEAMID_TODO` and `SHA256_FINGERPRINT_TODO`, verification fails and every code lands in the browser, which feeds M1. Fix: make `app-links.mjs` refuse values containing `TODO` once `APP_SIGN_IN` is on, or at least in the release build. Do not flip the flag until the real values are deployed.

### L5. Start rate limit is per IP, and each start writes a row
`api/internal/auth/app_auth.go:42`

60/hour/IP is fine against guessing (256-bit secrets). A distributed caller can still insert flows quickly. The sweep only runs on start and only deletes rows older than expiry + 1 h. Fix (optional): a global ceiling on live flows, or sweep in the daily job.

### L6. Missing tests
- There is no test that a replayed Battle.net callback on an app flow (second `IssueAppCode`) is refused.
- There is no test that an app token is refused on `RequireDevice`-free companion assumptions elsewhere.
- `/v1/devices/self` also accepts an app token. That is harmless (same user), but it is undocumented.

## Notes (not findings)
- A malformed verifier returns 400 *without* spending the code (`app_auth.go:152`). Such a verifier can never match, so nothing is lost. The brief's "every attempt spends" holds for every well-formed attempt.
- Migration 0059 is additive. `add column … default 'companion' not null` is a metadata-only change on PG11+. The currently deployed code keeps working ahead of the deploy, since its inserts default to companion.
