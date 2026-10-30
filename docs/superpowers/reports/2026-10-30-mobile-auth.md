# Mobile lane A1: app sign-in (report, 2026-10-30)

Branch `wt/mobile-auth`. Not pushed, not deployed. Migration **0059**.

## What landed

- **Routes, dark.** `GET /v1/auth/app/start`, `POST /v1/auth/app/exchange`, `POST /v1/auth/app/logout`. They are mounted only when `APP_SIGN_IN=true` (config.AppSignIn, default off) *and* Battle.net is configured. With the flag off they answer 404.
- **Shared Battle.net flow.** App start goes through the same `sendToBattleNet` as the site's start: same state cookie, same `BNET_REDIRECT_URL`, same callback. The callback is split into `checkBnetState` → `identifyBnet` → finish. The finish branches on the cookie's `next` value: an app flow carries the sentinel `app`, which `safeNext` can never produce. The web path is otherwise byte-for-byte the old one. **No new Battle.net redirect URI has to be registered** (lane O can drop that item).
- **Migration 0059.** It adds `app_auth_flows`: one row per sign-in, keyed by the SHA-256 of the Battle.net state. The row holds the platform, the S256 challenge and the app's state, then `code_hash`, `user_id`, `code_expires_at` and `used_at`. Rows older than an hour past expiry are swept when a new flow starts. It also adds `devices.kind` (`companion` default | `app`), because `platform` alone cannot tell the two apart: the companion names its own platform.
- **Policy change: `auth.RequireSession` now admits an app token.** Without this the app could not read `/v1/me`, guilds, sims, raid signups or anything else: almost every account route is RequireSession. A companion token is refused exactly as before, and the existing test pinning that still passes. Discord's browser-redirect connect flows still refuse every device. CSRF does not apply to bearer requests, as before.
- **Devices list.** `GET /v1/devices` gains `kind`. The site's `/account` already lists every device with Revoke (`DELETE /v1/devices/{id}`), so apps appear there as `name · ios`. The setup steps, the logs landing, the pairing poll, the "download the companion" prompt and the addon's ack (`addon/inbox_reads.go`) now count companions only. Before this, a signed-in phone would have read as a paired companion.
- **Universal links.**
  - The config is `web/app-links.json`. Placeholders: `TEAMID_TODO` and `SHA256_FINGERPRINT_TODO`, plus a `todo` field telling the owner where each value comes from.
  - `web/scripts/app-links.mjs` builds `public/.well-known/apple-app-site-association` and `assetlinks.json` from that config. It runs on `prebuild` and via `npm run sync:app-links`.
  - Only `/app/*` is listed. The script refuses any other path.
  - `_headers` serves both files as `application/json` with a 1 h cache.
  - A vitest holds the committed files to the config's output.

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

- `302` to Battle.net.
- `400 invalid`, with `error.fields` naming each bad parameter.
- `429 rate_limited`.

The app flow can sit at Battle.net for at most 10 minutes.

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
- After logout, or a Revoke on the site (`DELETE /v1/devices/{id}`), the token answers `401 unauthorized` everywhere.
- `GET /v1/devices` items are `{id, name, platform, kind: "companion"|"app", created_at, last_seen_at?}`.

## Tests

Go, `-race`, on a private database: auth, config, db, addon, server, cmd/api, discordbot and reports, all green. Web vitest for the touched files: 34 files and 548 tests green. ESLint and Prettier are clean on touched files.

The new auth tests cover:
- the happy path on both platforms
- a wrong verifier, which also burns the code
- a replayed code
- an expired code
- the two-minute lifetime
- a state mismatch
- an expired flow
- start and exchange validation
- the default device name
- logout, and a revoke from the site, each refusing the token afterwards
- the web flow unchanged with the flag on
- a companion token named `ios` still refused on `/v1/me`
- both rate limits
- the routes dark by default

There is also a migration up/down test, and the OpenAPI spec lists the three paths.

## Concerns and owner TODOs

1. **Owner TODO:** fill `appleTeamId` and `androidSha256Fingerprints` in `web/app-links.json`, run `npm run sync:app-links` and commit.
2. **`/app/auth` has no web page yet.** With the app not installed, or the link not verified, the browser lands on the 404 page. Lane D/F1 should add a small "Open the app" page there.
3. **Android redirect.** Chrome Custom Tabs does not always hand a server-side 302 to a verified App Link without a user gesture. ASWebAuthenticationSession on iOS 17.4+ handles it with an `https` callback (host `foreversixty.gg`, path `/app/auth`). Android F1 should test it on a device. If it fails, the `/app/auth` page from item 2 needs an explicit button.
4. **The RequireSession widening** (above) is the one behaviour change outside the new routes. It only takes effect once app devices exist, which needs the flag on.
5. **Not verified:** that Astro and Workers static assets publish the `.well-known/` dot-directory. Check `https://foreversixty.gg/.well-known/assetlinks.json` after the next site deploy.
6. **Launch freeze:** migration 0059 must not deploy from Nov 3 to Nov 5.
