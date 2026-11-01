# Mobile Wave 3 F16: guild extras (2026-10-31)

Branch `wt/mobile-guild-extras`, lane `mobile-guild-extras`. One lane, both platforms and the API. Spec: `design/specs/2026-10-31-mobile-wave3.md` §5.8, contract W3-C16 and W3-C17 (§10). Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f16-guild-*`.

## API (W3-C16, W3-C17)

- **Roster spec words.** `GET /v1/guilds/{id}/home`: `roster[]` gains `spec_name` and `spec_icon`; `members[]` gains `role`, `spec_name` and `spec_icon` (nullable), from the spec catalogue (`raidspec`), so the apps say "Holy Priest", never a bare "Holy", and never derive a role from a spec word. They are read only from fields the row already shows, so consent withholds them too (`api/internal/guilds/roster_spec.go`). `members[].character_key` is now optional in the spec, as the handler always sent it.
- **Calendar link.** `GET /v1/me/calendar-feed` (makes the link, shows its url once, then `url: null`), `POST /v1/me/calendar-feed/reset` (revoke and reissue in one transaction), public `GET /v1/calendar/{token}.ics` (RFC 5545, `{zone} · {answer}`, `{guild} · {zone}`, `Pull {time}. {n} of {size} signed up.`, raid link, no VALARM; planned, locked and done raids of every guild where the account is a verified member, 30 days back; cancelled raids dropped; membership read per fetch). Files `calendar_handler.go`, `calendar_store.go`, `calendar_ics.go`; migration `0064_calendar_feeds` (renumbered after main's 0063).
- **Security design.** Token: 256 random bits (`auth.NewSessionID`), base64url, shown once, stored only as SHA-256 (`token_hash bytea unique`, length-checked), one live row per account (partial unique index), reset under a per-account advisory lock. Budgets: 60 reads an hour per live link (keyed after the lookup, so random tokens cannot fill the limiter), 30 misses an hour per address (only a link that does not exist spends it, so calendar services fetching many real links from one address are never refused), 10 resets an hour per account. One 404 body for a malformed, unknown or revoked token. Headers: `private, max-age=300`, `nosniff`, `Referrer-Policy: no-referrer`; the account routes are `private, no-store`. The request and panic logs write the route as `/v1/calendar/:redacted` (`httpx.LogPath`, matched on the cleaned path so `//`, `/./` and `/../` spellings are redacted too). ICS text escapes `\ ; , CR LF` and drops control characters and U+2028/2029. The feed carries raid times, the guild name, the zone, a count and the owner's own answer: no notes, no other player's name.
- **Security pass.** A `security-reviewer` read the commit: no CRITICAL or HIGH. Fixed from it: the path-redaction bypass (MEDIUM), the limiter fill-up and shared-address refusals (LOW), the concurrent-reset unique violation (LOW), control characters in ICS text (LOW).
- OpenAPI, `docs/contracts/app-v1.md` (W3-C16, Raid history, Calendar link), fixtures `calendar.feed`, `calendar.feed.shown`, `calendar.ics.error-not-found`, the golden `calendar.ics`, `guild.home.roster` (a 38-person Horde guild), `guild.raids.history` and `-2`; generator `design/mocks/2026-10-31-mobile-wave3/gen_f16_fixtures.py`. Regenerated Swift and Kotlin models, endpoints, copy and the `guild_extras` switch.

## Apps (both platforms)

- Raids tab: `Roster` (`{n} on the roster · attendance`) and `Raid history` rows under the guild head.
- Roster: counts card, role chips from the server's role, sort menu (rank, class, name, attendance, level; ties by name), spec icon (class crest fallback), full spec name, rank word, attendance bar over `{n} of {nights} nights`, `Waiting` and `No account` tags, search, every state; a row opens the character's web page until F11 lands.
- Raid history: date block, bosses, kills and wipes, raiders (or players for 10 or fewer), night length, `You were there` under the meta for the account's own characters, Load more by cursor; a row opens Your raid night when you were there, else the report on the web.
- You › Calendar: an app-owned local calendar `Forever Sixty raids` (EventKit full access / CalendarContract), asked only after the switch, synced on switch-on, every Raids refresh and your own answer; Which raids and Alert; access-off and removed states. The calendar link with Copy (sensitive / local-only clipboard), Reset (confirmed) and, on iOS, Add to Calendar (`webcal://`); the url kept in the Keychain / the encrypted store, cleared at sign-out; `url: null` with nothing stored offers Reset.
- iOS `apps/ios/App/Guild/` (Logic, Calendar, Views); Android `.../app/guild/` (logic, platform, ui), `core/auth/SecretStore.kt`. Pure, tested logic for the roster model, history paging, the event builder and diff, and the link state machine; the system calendar sits behind a thin adapter with a fake in tests.

## Reviews

Captures (goldens and running-app captures on the fixture clock, default and largest text): `design/reviews/2026-10-31-mobile-wave3-guild/build/{ios,android}/`.
- ux-designer: HOLD, HOLD, **SHIP** (round 3) — `ux-verdict-build.md`.
- wow-player: HOLD, HOLD, HOLD, **SHIP** (round 4) — `player-verdict-build.md`.
The rounds fixed: iOS blank spec icons, the switch knob, condensed titles, XL history, calendar and Raids-tab layouts, a 24/38 roster count contradiction, fixture realism (Onyxia only on its five-day lockout, every night 9-18 Dec before the next planned raid, raid sizes by instance, varying lengths, one viewer everywhere, attendance counted from the history itself, the golden feed equal to the Raids tab's planned raids).

## Gates (after merging origin/main)

- `go test -p 1 ./internal/guilds ./internal/httpx ./internal/contract ./internal/server ./internal/db ./internal/auth ./cmd/api` (private DB): all ok.
- Codegen re-run, no drift; `npx vitest run src/lib/mobile`: 30 passed.
- Android `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`: exit 0.
- iOS `make -C apps/ios test`: see the coordinator reply.

## Gaps

1. **Switch not turned on.** Both reviewers SHIP, but the owner reverted F14's bundled-default flip (c6c402eee), so `guild_extras` stays default false; it turns on through the API's `APP_FEATURES` in the deployment, which is the owner's call.
2. **Logs outside the app.** Cloud Run's request logs (`httpRequest.requestUrl`) and any Cloudflare logs record the full feed URL, token included. App-level redaction cannot reach them: add a Cloud Logging exclusion for `requestUrl:"/v1/calendar/"` (and Logpush redaction if the path is proxied).
3. **`GET` creates the link** (spec W3-C17). A cross-site navigation with a cookie session could create a link the player never sees; no data is exposed and Reset recovers it.
4. **Feed start time.** The API has no `invite_at` yet (F4 contract pending), so the feed's `DTSTART` is the pull while the app's own calendar uses invites.
5. **Attendance counts every logged night,** dungeon runs included (the API's existing rule), so most raiders read 0-3 of 8 in the fixture. A raids-only window is a product follow-up (wow-player SHOULD).
6. **F16.6** is captured on iOS only (Google Calendar on the emulator needs a Google account). No background calendar sync on either platform (neither app has a background refresh yet).
7. History rows that you were not in open the report on the web; F10's native report screen is now on main and could take that tap.
