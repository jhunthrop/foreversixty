# Mobile backfill review brief (round 3)

Scope: the four tab roots (Home, Raid, Gear, Activity), the profile drawer, and the raid
notification, on iOS (SwiftUI, apps/ios) and Android (Compose, apps/android). The tabs are four by
design (docs/superpowers/plans/2026-11-09-redesign.md, Wave 3). Secondary screens (raid page, Roster,
History, Calendar, Character, Guides, Tiers, BiS, account, the planner) belong to another lane and are
out of scope here: mention them only if a main-tab interaction lands on them.

Branch: lane/mobile-review-backfill. Fixture character: Gorvash, Orc Fury Warrior 58, Officer of
Olympus XXVII; alt Yellfury, Hunter 31, no spec reported. The app is dark only, as the site.

What changed and what a ruling keeps: round-2-decisions.md and round-3-decisions.md (read them
first; contest a KEEP only with a reason its ruling does not answer). Your earlier verdicts are in
`<persona>.md` beside this file.

## Captures (round 3)
| Set | Where |
|---|---|
| iOS step captures (iPhone 17 Pro sim, fixture app, XCUITest; `-t0`, `-t400`, `-settled`, `-pressed`; `light-*`, `ax5-*`) | design/reviews/2026-11-10-mobile-backfill/ios/ (earlier rounds in ios-r1/, ios-r2/) |
| Android step captures (Pixel 9 class emulator, fixture app, adb; `light-*`, `fs2-*` at font scale 2) | design/reviews/2026-11-10-mobile-backfill/android-r3/ (round 2 in android-r2/) |
| Android mocked states (Roborazzi, incl. `-x1_5`, `-x2_0`) | apps/android/app/src/test/screenshots/parity/android-*.png (+ `.txt` text goldens) |
| Android flow steps (real controller, a service that echoes the answer: sign-up, Home sign-up, deep links, saving, failed, offline) | apps/android/app/src/test/screenshots/flows/ |
| iOS flow and parity snapshots | the round-3 CI record run is in flight; judge iOS from the step captures, the round-3 mocks and the code (apps/ios/AppTests/Flows, the parity test data) |
| Round-5 app mocks (updated for round 3) | design/mocks/app/ios-*.html, android-*.html, regions/ |
| Web counterparts at 390 px | design/reviews/2026-11-10-mobile-backfill/web/ (README.md) |
| Push fixtures | docs/contracts/fixtures/app/push.*.json (raid-posted, raid-short with its Mention action, raid-starting, guild-applicant with Accept and Decline) |

## Fixture limits (not app behaviour)
- In the fixture APP the sign-up write answers one canned raid, so a tap on Late snaps back to In
  (with `Saved.`). The flow tests use a service that echoes the answer: judge sign-up from those.
- Feed filters and Show older are server-side; the fixture returns one page for every query.
- The fixture app's raid has no comp target, so its tallies read `Healers 6`; the mocked states carry
  a target and read `Healers 5 of 8 · 3 short`. Its stamps read `12 d ago` (stale, warning tone) because
  the fixture's sync date is fixed.

## Known in round 3 (being fixed, do not re-report)
- iOS Gear in the fixture app reads `0 of 17 slots match the list` where Android reads `3 of 17` with
  Head, Neck and Main hand `Equipped` (the fixture's `already_best`); iOS is being aligned to Android.
- Android notifications show three actions (In, Late, Bench) by platform limit; the tap-through lands on
  the raid page, a secondary screen in the old style (another lane's).
