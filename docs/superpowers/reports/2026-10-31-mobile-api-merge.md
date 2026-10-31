# Mobile API merge: A1, A2, A3 landed on main (2026-10-31)

The owner lifted the launch freeze. The three held API branches went to main one at a time.
Each was green before the next one started. Each branch merged origin/main; no commit already
on main was rewritten.

| Lane | Branch | Landed on main | Deployed |
|---|---|---|---|
| A1 app sign-in | `wt/mobile-auth` | `811e70081` | run 41215643651, success |
| A2 push | `wt/mobile-push` | `b8a3c4ee8` | in run 41221035618 (`abe32d07d`), success; the run for `b8a3c4ee8` itself was cancelled by later pushes |
| A3 home and config | `wt/mobile-home-api` | `98367fa0d`, then the cache fix below | run 41224168734, success; the fix in its own deploy |

## Migrations

Main's highest migration was 0059 (`fight_counts`), so the branches were renumbered in order:

- A1: `0059_app_sign_in` became **0060**.
- A2: `0060_notification_prefs` became **0061**, and `0061_push_tokens` became **0062**.
- A3 has no migrations.

The migration tests and the code comments now use the new numbers. The deployed API applied
the migrations itself. Nothing was run against production by hand.

## A1: app sign-in

- **Merge.** The openapi conflict takes main's version, since the contracts lane already
  documents `/v1/auth/app/*`. The handler's answers match `AppSession` and `Device`. The
  sitemap keeps `/report-shell` and `/app/auth` out.
- **Security review** (the checklist, applied to `app_auth*.go`, `session.go`, `store.go` and
  the billing gates). These hold up:
  - PKCE S256 only.
  - The 256-bit one-time code is stored as a hash, can be used once, and expires after
    2 minutes. It is spent before the verifier is checked.
  - The state is bound to the browser cookie, and the flow is bound to the state hash.
  - The callback target is fixed, so there is no open redirect.
  - The confirm page sends a strict CSP and a frame refusal.
  - App tokens lapse after 30 days idle and 180 days total.
  - Billing writes need a browser session. Moderator powers never go to an app token.
  - The down migration revokes app devices.
- **HIGH, fixed (`339138c3e`).** `RequireDevice` admitted any device token. A signed-in
  app's token could therefore do the companion's jobs: upload exports and logs, read and
  drain the addon inbox, heartbeat a live report, and Forget a device. `RequireDevice` is
  now companion-only. The app's logout moved behind a new `RequireApp` gate. A test covers
  both gates.
- **Still open, not blocking:**
  - L2: there is no cap on live app devices per account.
  - L4: `apple-app-site-association` and `assetlinks.json` still carry `TEAMID_TODO` and
    `SHA256_FINGERPRINT_TODO`. Universal and App Links will not verify until the owner
    fills `web/app-links.json` and runs `npm run sync:app-links`. Until then, the code
    lands on the `/app/auth` page, which hands it back to the app.

## A2: push

- **FCM through ADC.** `FCM_SERVICE_ACCOUNT_JSON` is gone. The FCM v1 sender uses
  Application Default Credentials, which on Cloud Run is `api-runtime@foreversixty`. That
  account holds `roles/firebasecloudmessaging.admin`, verified in the project IAM policy.
  `FCM_PROJECT_ID` is optional and defaults to the credentials' project (`foreversixty`,
  where Firebase and `fcm.googleapis.com` are enabled). If no credentials are found, the
  API logs "FCM off" and keeps serving.
- **APNs stays dark.** With no `APNS_*` keys, startup logs "APNs off: iOS push is skipped".
  The sink logs each iOS token it skips. Nothing fails. After the flag went on, production
  logged "APNs off" and did not log "FCM off".
- **Push tokens.** `/v1/me/push-tokens` takes app tokens only (`RequireApp`). This follows
  from the A1 fix: under the old `RequireDevice`, companions could register.
- **Raid sign-up PUT.** It took `auth.Require`, which let a companion token sign up. It now
  takes `RequireSession` (browser or app) while `APP_NOTIFICATIONS_ENABLED` is on, and a
  browser session only while it is off.
- **Discord output is unchanged by the flag.** `DiscordSink` still renders only
  `sim.finished` and `report.ready`.

## A3: home and config

- **Merge.** The Discord slot choices keep A3's single source, `upgrades.Slots`, and take
  main's "Finger 1/2" labels. A second merge combined A3's `App` config with A2's `Push`
  config and `Deps`.
- **Fixtures.** A3's duplicate fixtures (`characters-home*.json`, `app-config.json`) are
  deleted. Main's contract test fails on any fixture no response claims. A3's contract test
  now reads `character.home*` and `app.config`. `drop_chance` stays optional on the home
  pick, as the spec allows. The handler does not send it yet.
- **Decision L.** `web_checkout_regions` depends on the platform in the `X-App` header:
  - `APP_WEB_CHECKOUT_REGIONS_IOS` defaults to `US`.
  - `APP_WEB_CHECKOUT_REGIONS_ANDROID` defaults to `none`, which gives `[]`.
  - A missing, malformed or unknown platform gets `[]`.

  The response sets `Vary: X-App`.

  **Found after the deploy and fixed.** Cloudflare ignores `Vary: X-App`. Once
  `APP_API` was on, the `public, max-age=300` answer was a cache HIT for every platform:
  ios, android and no header all got the same `request_id` and `["US"]`. That broke
  decision L for Android. The answer is now `private, max-age=300` through a new
  `httpx.CachePrivateFor`, so the app may keep it and no shared cache stores it. The spec
  and both contract docs say so.

  When the Play enrolment comes through, set
  `APP_WEB_CHECKOUT_REGIONS_ANDROID=US` on the service. No code change or app release is
  needed. Tests cover ios, android, no header, an unknown platform and a malformed header.
- **`features`.** The spec requires this field, and A3 did not send it. It is now filled
  from `APP_FEATURES` (a comma list of keys, each sent as true). When the variable is unset,
  the field is `{}`.
- **Docs.** `docs/contracts/app-v1.md`, `app-v1-home-config.md` and the spec description
  are updated. The generated endpoints changed only in their doc comment, so no model
  changed and the iOS and Android suites were not needed. The site's parity test got a type
  fix that `astro check` flagged after the merge.

## Switches turned on (Cloud Run service `api`)

| Switch | State | Why it is safe |
|---|---|---|
| `APP_SIGN_IN=true` | on after A1's deploy | Behind the confirm page and the M1/M2 mitigations. A bad start returns 400; a good start returns the 200 confirm page. |
| `APP_NOTIFICATIONS_ENABLED=1` | on after A2's deploy | No app is in a store, so no push tokens exist. Discord output is unchanged. Anonymous calls to the routes get 401. |
| `APP_API=true` | on after A3's deploy | Public reads only: the character page's visibility, and the config. |

Left alone: `APP_FEATURES` stays unset, so the apps keep their bundled feature defaults. The
`APNS_*` keys stay unset until the Apple key exists. The jobs share the image but not the
service's environment, so the `raid-posts` job keeps push off until the flag is set on it too.

## Gates (each run on its own, exit code checked)

- **Go tests.** `go test -p 1 ./...` ran on a private database (`forever_test_apimerge`).
  It was green for every branch except `internal/mail`'s
  `TestResendTransportErrorIsWrapped`, which fails only because of the sandbox HTTP proxy.
  It passes with the proxy unset, and A3's full run was clean that way. The database was
  dropped afterwards.
- `go vet ./...` and `gofmt -l` were clean. CI's Go lint is `go vet`.
- The contract tests (`internal/contract` and `internal/app`) passed.
- Mobile codegen drift (`vitest src/lib/mobile src/styles/mobile-tokens.test.ts`) passed,
  and the generator was re-run after every merge.
- Web: `astro check` returned 0 errors. Vitest passed for the touched areas: account,
  app-links, `/app/auth`, setup, logs, home.

## Left for later

- The owner needs to fill in the Apple Team ID and the Play signing SHA-256 in
  `web/app-links.json`, then sync and deploy.
- The APNs `.p8` key and `APNS_KEY_ID`/`APNS_TEAM_ID` go into GSM and the service.
- FCM is wired but has not sent a real message yet, because no app token exists. The first
  Android token will exercise it.
- The 426 `update_required` gate that the spec describes for `X-App` is not built.
- `drop_chance` on the home pick is not sent yet.
- L2: there is no cap on app devices per account.
