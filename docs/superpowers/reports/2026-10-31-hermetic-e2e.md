# Hermetic e2e, raid 429 copy, neutral Discord answers (2026-10-31)

Lane `wt/hermetic-e2e`. Trigger: the owner's raid save got `429 rate_limited` in production because
local Playwright runs had spent the owner's per-IP budget. The fixture site was built against
`https://api.foreversixty.gg`, so every API call a spec didn't stub reached production. Before this
change, local runs, CI's `e2e` job and CI's `lhci` job all did this.

## 1. The e2e suite is hermetic

The three layers are all defined in `web/tests/hermetic.ts`:

1. **Dead API base.** `playwright.config.ts`'s `webServer` env and web.yml's fixture build (the
   dist that `e2e` and `lhci` serve) set `PUBLIC_API_BASE_URL=http://127.0.0.1:9`. Only deploy's
   real-data build keeps the production URL. `build-real` keeps the workflow default but serves
   nothing.
2. **Production host unresolvable.** Every test browser launches with
   `--host-resolver-rules=MAP api.foreversixty.gg ~NOTFOUND`. That covers `playwright.config.ts`,
   `playwright.journey.config.ts`, all 10 capture configs, `scripts/capture-guild.mjs`, and
   Lighthouse's Chrome through `lighthouserc.json`'s `chromeFlags`. A dist built with the
   production default therefore cannot reach the real API either.
3. **Guard fixture.** `web/tests/e2e/support/test.ts` exports the `test` that all 166 specs now
   import. Its context route aborts any API request a spec didn't stub and fails the test with the
   list. When the served dist still points at production, the failure message says to rebuild with
   the dead base. Stubs a spec adds, on the page or the context, take priority over the guard.
   - `AMBIENT_READS` gives a few site-wide background GETs a default answer instead of a failure:
     - `/v1/me` answers 401 (signed out).
     - The rest are aborted as unreachable: character `sim-input` and `rating`, `specs`, `phases`,
       `rankings/percentile`, per-fight `ratings`, and the reports, recent reports, top-guilds,
       `sims` and coach-recent feeds.
   - Before this change those reads reached production but failed CORS in the browser (the comment
     in sim-talents.spec.ts confirms localhost has no CORS grant). The default therefore matches
     what the specs already saw, without spending the rate limit.
   - Any other read, and every write, fails the test.
   - `apiUnreachable(page, ...globs)` lets a spec declare a read unanswered explicitly.

`web/src/e2e-hermetic.test.ts` enforces all of this:
- every spec imports `./support/test`;
- nothing under `tests/e2e` takes `test` straight from `@playwright/test`;
- every Playwright config carries `HERMETIC_LAUNCH_ARGS`;
- `lighthouserc.json` carries the same rule;
- web.yml's fixture build uses `DEAD_API_BASE`.

The README's e2e section documents the setup.

### Specs that had been reaching production

The first full run with the guard on failed 1,603 tests, and every one listed unstubbed reads. No
writes leaked. The site-wide reads are now covered by `AMBIENT_READS`. These specs had additional
unstubbed reads and now declare them with `apiUnreachable`:

- first-run-signed-in-empty, logs-recent-reports (`/v1/devices`)
- guild-billing (`/v1/guilds/*/billing/upgrade-preview`)
- guild-cache-invalidation, guild-member-view (guild `loot`, `progression`, `raids`, `readiness`)
- guild-centre (those four plus `settings`)
- guild-claim (`/v1/devices`, `/v1/guilds/*/home`)
- guild-settings, guild-settings-discord (`/v1/guilds/*/billing`)
- logs-companion (`/v1/characters/{region}/{ruleset}/{name}`)
- logs-guild-tabs (character, `/v1/devices`)
- nav-selector (character, `/v1/encounters`)
- rankings, rankings-phone, sim-execution, spine-bar (`/v1/encounters`)
- report-phone (`/v1/rankings`)
- sim-drops, sim-gear, sim-tabs, sim-talents, sim-tools-phone (`/v1/builds?mine=1`)

Every report, rankings and character spec was also hitting production through the ambient reads.
That includes the sample-report character Elyra Duskvale (`/sim-input`, `/rating`), which matches the
production logs.

Specs that asserted the production URL itself now read it from the configured base (`E2E_API_BASE`,
which equals `DEAD_API_BASE`):
- account-discord-link (start route)
- support/discord-notifications.ts, which covers account-discord, guild-discord-bot and
  guild-settings-discord
- report-queries (first-party origins)

## 2. A 429 on an officer write says to wait

New copy key `raidCopy.rateLimited` in `web/src/lib/guild/raid-copy.ts`: "Too many requests right
now. Wait a minute and try again." `raids-tab-controller.ts` gains `isRateLimited` and
`writeFailedLine`.

- `submitRefusal` (raid form create, `POST`, and edit, `PATCH`) returns the new line on a 429.
- The raid page controller's shared `refusalLine` uses it for lock, unlock and cancel, for
  match/clear match, and for the attendance night actions through `NightActions`' host.
- The roster state select now goes through `refusalLine` too.

New tests are in `raids-tab-controller.test.ts` and `raid-page-controller-officer.test.ts`. Members'
own sign-up presses still say "That did not save. Try again." because they are not officer writes.

## 3. Discord answer buttons are all Secondary

- `answerButton` in `api/internal/discordbot/post_message.go` builds In, Late, Bench and Out for the
  channel raid post. All four are Secondary, and Raid page stays a Link button.
- The ephemeral `/signup` card in `signup_card.go` had the same green In before anything was picked.
  It uses the same helper now.
- Tests updated: `post_message_test.go` and `signup_card_test.go`.
- Committed payload frames updated:
  - `design/mocks/2026-10-28-discord-guided-cards/payloads/post-*.json`
  - `design/mocks/2026-10-28-discord-guided-cards/payloads/signup-card.json`
  - `web/tests/captures/discord-bot/payloads.json`
- The two bot specs are amended with the owner's ruling.

The remaining Success buttons are `Show BiS` (bis card) and `Type a name` (char card). Both are the
card's one forward action, not a preselected answer, so they stay.

## Gates

| Gate | Result |
| :--- | :--- |
| web vitest (full) | exit 0, 691 files, 7,792 tests |
| `astro check` | exit 0, 0 errors |
| lint | exit 0 |
| format:check | exit 0 |
| Playwright full suite, desktop and mobile (dead-base fixture build, port 4711) | 3,392 passed, 410 skipped, 6 failed (see below); `--last-failed` rerun left 4 |
| `go vet ./internal/discordbot/ ./internal/notify/` | exit 0 |
| `go test -p 1` discordbot, notify (private database `forever_gc_btn`, dropped) | exit 0 |
| `actionlint .github/workflows/web.yml` | exit 0 |

The suite's remaining failures are pre-existing flakes. Each spec's original copy from `HEAD`, with
no guard, fails the same way when run side by side with `--repeat-each=3`:

- `account-discord.spec.ts` "turning an event off…" (the switch's uncheck race)
- `guild-raid-spec-ask.spec.ts` "pressing a state shifts nothing…"
- `home-panel.spec.ts` "signed in at 390/1440 … nothing moving". Under load the account-panel island
  hydrates after the 300 ms settle window and the hero changes height by 36 px.

`account-load-error` failed once and passed on rerun.

## Concerns

- A dist built by hand for e2e (the lanes' `npm run build` on 4711) must now set
  `PUBLIC_API_BASE_URL=http://127.0.0.1:9`. Without it the href-asserting Discord specs fail, though
  the resolver rule still keeps production unreachable. Lane instructions should add the variable.
- `lhci` and `lhci-nightly` now audit a dist whose API calls are refused at once instead of failing
  CORS against production. The scores could move slightly; neither job gates deploy in build mode.
- `go test` on discordbot and notify in parallel against one database deadlocks. Run them with `-p 1`.
- The design mock PNGs under `design/mocks/2026-10-28-discord-guided-cards/` still draw a green In.
  Only the payload JSON was updated; regenerating the boards needs `gen_guided_cards.mjs`.

## Follow-up (same day): reviews, merge, landing

- The guided-cards boards were redrawn with all four answers neutral (`FRAMES=post-,signup-card node
  gen_guided_cards.mjs`). The generator now leaves Components V2 rows alone and reads spec icons from
  `data/builds/`, so a fixture sync no longer blanks them. A fresh payload dump confirmed the real
  handlers emit Secondary everywhere, differing from the committed frames only in ids and nonces.
- The raid form's 429 line is captured at 1440 and 390 in `design/reviews/2026-10-31-raid-429/`
  (`tests/captures/guild-raids/capture-429.spec.ts`) and is covered by a new e2e test in
  guild-raids.spec.ts.
- ux-designer: SHIP, no blockers. Nits: the error red is small and low contrast (a site-wide token),
  and the line sits close to the card edge. wow-player: SHIP on both, no blockers. Nits: Raid page
  shares the grey (that is Discord's own style), and the line could quote Retry-After.
- The guard now routes by origin glob, not a function matcher. A function matcher sends every request
  through Node, which slowed hydration enough to move layouts: home-panel failed 5 of 12 runs with
  it and 0 of 12 with globs.
- After merging origin/main: spine-bar.spec.ts stays deleted. selector-restored.spec.ts (new) imports
  the guarded test and stubs character and encounters. handoffs-setup (builds) and guild-landing
  (guild page and home) gained stubs.
- Gates after the merge, all exit 0 unless noted: vitest 691 files / 7,785 tests, astro check, lint,
  format, actionlint, go vet, go test discordbot and notify. Full Playwright: 3,397 passed, 3 failed
  (account-discord "turning an event off", header-finished-sim 1440 mobile, home-panel 390 mobile).
  All 3 passed on `--last-failed`. A clean origin/main built with the dead base and run the same way
  failed home-panel and guild-raid-spec-ask too, plus the 23 href and unstubbed failures this lane fixes.
