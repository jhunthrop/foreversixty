# Mobile backfill review brief (round 9)

Scope: the four tab roots (Home, Raid, Gear, Activity), the profile drawer, and the raid
notification, on iOS (SwiftUI, apps/ios) and Android (Compose, apps/android). The tabs are four by
design (docs/superpowers/plans/2026-11-09-redesign.md, Wave 3). Secondary screens (raid page, Roster,
History, Calendar, Character, Guides, Tiers, BiS, account, the planner) belong to another lane and are
out of scope here: mention them only if a main-tab interaction lands on them.

Branch: lane/mobile-review-backfill. Fixture character: Gorvash, Orc Fury Warrior 58, Officer of
Olympus XXVII; alt Yellfury, Hunter 31, no spec reported. The app is dark only, as the site.

What changed and what a ruling keeps: round-2- to round-9-decisions.md (read them
first; contest a KEEP only with a reason its ruling does not answer). Your earlier verdicts are in
`<persona>.md` beside this file.

## Captures (round 9)
Round 9 adds `android-r9/` (Gear tab and the three item sheets: crafted helm with a PvP runner-up, crafted Dawn's Edge, vendor Leggings) and `ios/r9-*` (the same three sheets). Round 9 changed only the sheet's runner-up row (and, through the shared `items.sameDps`, the item tooltip's "The next best" rows).

Round 8 adds `android-r8/` (Gear `No gain`, the item sheet for a crafted pick with its skill, a crafted upgrade, a vendor upgrade, the Mention and Decline failures) and `ios/r8-*` (the same, plus the drawer Notifications frames from the top). No capture shows a boss-drop runner-up: the fixture's first head runner-up is a PvP helm and the list order is kept as the data has it; the boss-drop row is in the item-sheet parity state (`Onyxia, Onyxia's Lair, 8%`).
Round 7 adds `android-r7/` (fresh debug build: Gear `182 DPS`, the Gear sheet with its Where line, both officer push failures, the drawer Notifications value some off / All on / All off) and `ios/r7-*` (the same set plus `r7-home-01-top`). Earlier sets stand for what round 7 did not change.
Round 6 adds `android-r6/` (fresh debug build: drawer, Activity in the house format, officer roster filtered, push failure results, a stale Gear frame) and `ios/r6-*` (the same set). The round-5 sets below still stand for everything round 6 did not change.

| Set | Where |
|---|---|
| iOS step captures (iPhone 17 Pro sim, fixture app, XCUITest; `-t0`, `-t400`, `-settled`, `-pressed`; `light-*`, `ax5-*`) | design/reviews/2026-11-10-mobile-backfill/ios/ (recaptured for round 5; `r4-*`/`r5-*` are the officer, filtered roster, Refreshing and push-result steps; the officer frames pin the clock to 17 Dec, so they read `Synced 37 d ago`, the one stale capture kept; earlier rounds in ios-r1/, ios-r2/) |
| Android step captures (Pixel 9 class emulator, fixture app, adb; `light-*`, `fs2-*` at font scale 2) | design/reviews/2026-11-10-mobile-backfill/android-r5/ (incl. the officer Raid tab and its filtered roster, officer pushes with their results, the drawer Refreshing frame; earlier rounds android-r2/ to android-r4/) |
| Android mocked states (Roborazzi, incl. `-x1_5`, `-x2_0`) | apps/android/app/src/test/screenshots/parity/android-*.png (+ `.txt` text goldens) |
| Android flow steps (real controller, a service that echoes the answer: sign-up, Home sign-up, deep links, saving, failed, offline) | apps/android/app/src/test/screenshots/flows/ |
| iOS flow and parity snapshots | round-5 goldens from record run 42953910502 are in apps/ios/AppTests/**/__Snapshots__ (incl. Flows/); the round-9 run (42968686168) re-records every golden that rounds 6 to 9 moved and is in flight; until it lands, judge iOS from the step captures (apps/ios/AppTests/Flows, the parity test data) |
| Round-5 app mocks (updated for round 5) | design/mocks/app/ios-*.html, android-*.html, regions/ |
| Web counterparts at 390 px | design/reviews/2026-11-10-mobile-backfill/web/ (README.md) |
| Push fixtures | docs/contracts/fixtures/app/push.*.json (raid-posted, raid-short with its Mention action, raid-starting, guild-applicant with Accept and Decline) |

## Fixture limits (not app behaviour)
- In the fixture APP the sign-up write answers one canned raid, so a tap on Late snaps back to In
  (with `Saved.`). The flow tests use a service that echoes the answer: judge sign-up from those.
- Feed filters and Show older are server-side; the fixture returns one page for every query.
- The fixture app's raid has no comp target, so its tallies read `Healers 6`; the mocked states carry
  a target and read `Healers 5 of 8 · 3 short`.

## Known in round 4 (do not re-report)
- The fixture build (1.60.1.70658) has no art for the Might and Wisdom spell icons (the names load on
  the live build), so those names draw no icon, by the no-empty-frame rule.
- Android notifications show three actions (In, Late, Bench) by platform limit; the tap-through lands on
  the raid page, a secondary screen in the old style (another lane's), as does the filtered roster.
- iOS parity lists dated known drifts (AppTests/ParityCheck.swift): Home's raid-line title wraps on the
  phone and not in the mock's Chromium (16 pt), raid-officer Before you pull at the 8 pt limit, and the
  mocks' 18.29 sub line against the app's 18 (Olympus XXVII, Activity This week).

## After the round-5 merge with main (not a regression)
Main brought the house format (apps/shared/formats.json, owner rulings 2026-11-10): figures grouped and
whole from 100, the 12-hour clock without `:00` in server text. Round-5 row 6 (one decimal on Activity
figures) is therefore reverted: Activity reads `204 DPS`; the guild-ladder lead reads `+12 over your 191
(Fury Warrior, 60).`; raid pushes read `7 pm CST`. Step captures taken before the merge may still show
`203.6 DPS` or `7:00 pm` in push text: judge by the house format. The apps' own raid clocks
(`7:00 pm`) move with the apps' format lane, not this one.
