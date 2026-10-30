# Mobile lane A1: app sign-in (report, 2026-10-30)

Branch `wt/mobile-auth`. Not pushed, not deployed. Migration **0059**.

## What landed

- **Routes, dark.** `GET /v1/auth/app/start`, `GET /v1/auth/app/continue`, `POST /v1/auth/app/exchange`, `POST /v1/auth/app/logout`. They are mounted only when `APP_SIGN_IN=true` (config.AppSignIn, default off) *and* Battle.net is configured. With the flag off they answer 404.
- **Shared Battle.net flow.** App start sets the same state cookie as the site's start (`setBnetState`) and uses the same `BNET_REDIRECT_URL` and callback. The callback is split into `checkBnetState` → `identifyBnet` → finish. The finish branches on the cookie's `next` value: an app flow carries the sentinel `app`, which `safeNext` can never produce. The web path is otherwise byte-for-byte the old one. **No new Battle.net redirect URI has to be registered** (lane O can drop that item).
- **Migration 0059.** It adds `app_auth_flows`: one row per sign-in, keyed by the SHA-256 of the Battle.net state. The row holds the platform, the S256 challenge and the app's state, then `code_hash`, `user_id`, `code_expires_at` and `used_at`. Rows older than an hour past expiry are swept when a new flow starts. It also adds `devices.kind` (`companion` default | `app`), because `platform` alone cannot tell the two apart: the companion names its own platform.
- **Policy change: `auth.RequireSession` now admits an app token.** Without this the app could not read `/v1/me`, guilds, sims, raid signups or anything else: almost every account route is RequireSession. A companion token is refused exactly as before, and the existing test pinning that still passes. Discord's browser-redirect connect flows still refuse every device. CSRF does not apply to bearer requests, as before.
- **Devices list.** `GET /v1/devices` gains `kind`. The site's `/account` already lists every device with Revoke (`DELETE /v1/devices/{id}`), so apps appear there as `name · ios`. The setup steps, the logs landing, the pairing poll, the "download the companion" prompt and the addon's ack (`addon/inbox_reads.go`) now count companions only. Before this, a signed-in phone would have read as a paired companion.
- **Universal links.**
  - The config is `web/app-links.json`. Placeholders: `TEAMID_TODO` and `SHA256_FINGERPRINT_TODO`, plus a `todo` field telling the owner where each value comes from.
  - `web/scripts/app-links.mjs` builds `public/.well-known/apple-app-site-association` and `assetlinks.json` from that config. It runs from `npm run sync`, which prebuild and predev run, and via `npm run sync:app-links`. (The first round added a second `prebuild` key, which silently lost to the existing one; fixed in the fix round.)
  - Only `/app/*` is listed. The script refuses any other path.
  - `_headers` serves both files as `application/json` with a 1 h cache.
  - A vitest holds the committed files to the config's output.

## Fix round (security review, 2026-10-30)

The review is `docs/superpowers/reports/2026-10-30-mobile-auth-security.md` (verdict MERGE). Fixed: M1, M2, L1, L3 and L6's replay test. Left: L2, L4 and L5 (see Concerns).

- **M1, start side.** `GET /v1/auth/app/start` no longer redirects.
  - It records the flow, sets the state cookie and answers a self-contained confirm page: "Sign in to the Forever Sixty app on your phone", a "Continue to Battle.net" link and "If you didn't start a sign-in from the Forever Sixty app on your phone, close this page."
  - The page has no script, form, image or meta refresh. CSP is `default-src 'none'; style-src 'unsafe-inline'; base-uri 'none'; form-action 'none'; frame-ancestors 'none'`, with `X-Frame-Options: DENY`, `Cache-Control: no-store` and `Referrer-Policy: no-referrer`.
  - The link is `GET /v1/auth/app/continue?state=…`. It goes on to Battle.net only when this browser's state cookie holds that state *and* marks an app flow. A continue link sent from another browser, a forged state, or a web sign-in's state all answer 400.
  - So a link alone cannot finish a flow: someone has to click on our page, and framing is refused.
  - Battle.net's support for a `prompt` parameter is not confirmed, so nothing relies on it. The interstitial is used every time.
- **M1, landing side.** `web/src/pages/app/auth.astro` (Base layout, noindex, kept out of the sitemap).
  - Its `<head>` script (`web/src/lib/app-links/handoff.ts`, inlined) strips `code` and `state` with `history.replaceState` before paint and keeps them in a closure.
  - "Open Forever Sixty" is a `<button>` (nothing in the DOM) that navigates back to the app link. On Android it uses `intent://…#Intent;scheme=https;package=gg.foreversixty.app;end`, which Chrome needs from a Custom Tab.
  - The page says "If you didn't start a sign-in on your phone, close this page." With no code it shows "Start again from the Forever Sixty app" and no button.
  - `_headers` serves `/app/*` with `Referrer-Policy: no-referrer` and `Cache-Control: no-store`. The page loads no third-party script.
- **M2, expiry.**
  - An app token lapses 30 days after last use (`last_seen_at`, or `created_at` before first use) or 180 days after sign-in. Both are config: `APP_TOKEN_IDLE_DAYS` and `APP_TOKEN_MAX_AGE_DAYS`, whole days ≥ 1, idle ≤ max age, wired into `auth.Store`.
  - The hot path stays one query: the live-device clause is in the `DeviceByToken` update. Only on a miss does a second query revoke a still-live app row.
  - A lapsed token answers `401 session_expired` once. It is revoked at that moment, so after that it answers `401 unauthorized` and leaves the device list. `Devices()` hides lapsed apps too.
  - Companion tokens never lapse.
- **M2, browser-only powers.**
  - `auth.RequireBrowserSession` answers an app token `403 browser_session_required` ("this can only be done on foreversixty.gg in a browser"). It is on every billing write: checkout, portal, guild upgrade, transfer, transfer withdraw and transfer decline. Billing reads stay open.
  - Moderator powers are stripped centrally: `Actor.IsModerator()` is false for an app token, so every moderator gate across guilds, reports and coach treats an app as a non-moderator. The hidden-standing routes keep their deliberate 404 rather than a new error, so an app learns nothing about them. `/v1/me`'s guild `can` lists follow the same rule.
  - Seat give, move and accept stay on RequireSession: they assign a paid seat but are not a payment.
- **L1.** `DELETE /v1/sessions` and `POST /v1/auth/logout` called with an app token revoke that app's device. Before, they answered success and did nothing.
- **L3.** The 0059 down migration revokes every live app device before dropping `kind`, so no phone is left reading as a companion.
- **L6.** There is now a test that a replayed Battle.net callback on an app flow mints no second code.

## Contract (for lane C's app-v1.md)

All bodies use the API envelope: `{ok, data, request_id}` on success and `{ok:false, error:{code, message, fields?}, request_id}` on failure.

### `GET /v1/auth/app/start`

The app opens this in the system browser: ASWebAuthenticationSession on iOS, Custom Tabs on Android. Rate limit: 60 per hour per IP.

| query | rule |
|---|---|
| `platform` | `ios` \| `android` |
| `code_challenge` | base64url SHA-256 of the verifier, no padding, exactly 43 chars `[A-Za-z0-9_-]` |
| `code_challenge_method` | `S256` (only) |
| `state` | the app's own CSRF state, 16–128 chars `[A-Za-z0-9._~-]` |

- `200` with the confirm page (`text/html`): the person must tap "Continue to Battle.net", which is `GET /v1/auth/app/continue?state=…`. That answers `302` to Battle.net, or `400 invalid` from a browser that did not start the flow. Continue shares start's budget.
- `400 invalid`, with `error.fields` naming each bad parameter.
- `429 rate_limited`.

The app flow can sit at Battle.net for at most 10 minutes.

**Redirect URI.** `https://foreversixty.gg/app/auth` is the one canonical landing path, and the apps must register exactly it (not `/app/auth/callback`). The server sends Battle.net its own configured callback and always hands the code to `/app/auth`. **Any `redirect_uri` the client sends is ignored and never trusted.** The Android lane's `redirect_uri` parameter has no effect; it should drop it and match `https://foreversixty.gg/app/auth`.

### Battle.net callback (app flow)

`302` to `https://foreversixty.gg/app/auth?code=<one-time code>&state=<the app's state, verbatim>`.

- No `fs_session` cookie is set.
- The code is 43 base64url chars, stored only as SHA-256. It is single use and expires 2 minutes after the callback.
- The app **must** compare `state` with what it sent.
- A forged state answers `400 invalid` and mints nothing. So does an expired or already-issued flow: `400 invalid`, "start again in the app", with no web session as a fallback.

### `POST /v1/auth/app/exchange`

No cookies. Rate limit: 60 per hour per IP.

```json
{ "code": "…", "code_verifier": "43–128 chars [A-Za-z0-9._~-]", "device_name": "Baelgrim's iPhone" }
```

`device_name` is optional and trimmed to 60. When empty it defaults to `iOS` or `Android`.

`201`:

```json
{ "token": "fsd_…", "device": { "id": "…", "name": "…", "platform": "ios" }, "me": { …the GET /v1/me body… } }
```

The token is shown only here. Send it as `Authorization: Bearer fsd_…`.

| status | code | when |
|---|---|---|
| `400` | `invalid` | malformed body, no code, or a verifier that is not well-formed |
| `400` | `invalid_grant` | the code is unknown, spent, expired or has the wrong verifier |
| `429` | `rate_limited` | over the per-IP budget |

All `invalid_grant` cases give one answer. **Any attempt spends the code**: a stolen code tried without the verifier is dead, and the real app must start again.

### `POST /v1/auth/app/logout`

Bearer device token. The calling device revokes itself.

- `200 {signed_out: true}`.
- `401` with no token, or a dead one.

### Afterwards

- The app token reads every RequireSession and Require route, `/v1/me` included.
- It does not carry moderator powers.
- Billing writes answer `403 browser_session_required`; the app should send the person to the site.
- After logout, `DELETE /v1/sessions`, `POST /v1/auth/logout` or a Revoke on the site (`DELETE /v1/devices/{id}`), the token answers `401 unauthorized` everywhere.
- **Expiry.** After 30 days unused or 180 days from sign-in (config), the next request answers `401 session_expired` once, and `401 unauthorized` after that. On either 401 the app clears its token and signs in again.
- `GET /v1/devices` items are `{id, name, platform, kind: "companion"|"app", created_at, last_seen_at?}`.

## Tests

**Fix round.** Go, `-race`, on a private database: auth, config, db, billing, server, cmd/api, addon, guilds, seats, coach, reports and discordbot, all green. The guilds, coach and reports packages are included because of the moderator change. Web vitest for `src/lib/app-links`, `src/pages`, `src/layouts` and `src/lib/account`: 46 files and 695 tests green. ESLint and Prettier are clean on touched files.

The fix round adds tests for:
- the confirm page: no redirect, no script, form or refresh, plus its headers and CSP
- continue: refused from another browser, with a forged state, or with a web sign-in's state
- a client `redirect_uri` ignored
- a replayed callback minting nothing
- expiry: idle, too old, limits taken from the Store, a recently used token, and a companion that never lapses
- `RequireBrowserSession`, and every billing write refusing an app token
- no moderator powers on an app token
- both generic sign-outs revoking an app
- the lifetime config's defaults, overrides and refusals
- the down migration revoking app devices
- the `/app/auth` handoff in jsdom: the URL stripped, the button's link on iOS and Android, no button without a code
- the page: script in `<head>`, copy, no third-party script, `_headers` rule

**First round.** Go, `-race`: auth, config, db, addon, server, cmd/api, discordbot and reports. Web: 34 files and 548 tests. It covered:
- the happy path on both platforms
- a wrong verifier, which also burns the code
- a replayed code
- an expired code and the two-minute lifetime
- a state mismatch and an expired flow
- start and exchange validation
- the default device name
- logout, and a revoke from the site
- the web flow unchanged
- a companion token named `ios` refused on `/v1/me`
- both rate limits
- the routes dark by default
- the migration up and down
- OpenAPI paths

## Concerns and owner TODOs

1. **Owner TODO:** fill `appleTeamId` and `androidSha256Fingerprints` in `web/app-links.json`, run `npm run sync:app-links` and commit. Do not turn `APP_SIGN_IN` on until the real values are deployed (review L4). The generator does not yet refuse `TODO` values.
2. **Android hand-off** needs a test on a real device: the 302 into a verified App Link from a Custom Tab, and the `/app/auth` button's intent.
3. **iOS hand-off.** ASWebAuthenticationSession on iOS 17.4+ catches the `https` callback (host `foreversixty.gg`, path `/app/auth`) before the page loads. Re-opening the same domain's link from the `/app/auth` page in Safari stays in Safari, which Apple's universal link rules require. So on iOS the page mostly helps people who do not have the app.
4. **The RequireSession widening** is the one behaviour change outside the new routes. It only takes effect once app devices exist, which needs the flag on.
5. **Not verified:** that Astro and Workers static assets publish the `.well-known/` dot-directory. Check `https://foreversixty.gg/.well-known/assetlinks.json` after the next site deploy.
6. **Open review LOWs:**
   - L2: no cap on live app devices per account. Expiry now bounds the damage.
   - L5: the start rate limit is per IP. The flow table is swept on start.
7. **Launch freeze:** migration 0059 must not deploy from Nov 3 to Nov 5.
