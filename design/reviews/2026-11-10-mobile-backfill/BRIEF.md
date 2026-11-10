# Mobile backfill review brief (round 2)

Scope: the four tab roots (Home, Raid, Gear, Activity), the profile drawer, and the raid
notification, on iOS (SwiftUI, apps/ios) and Android (Compose, apps/android). The tabs are four by
design (docs/superpowers/plans/2026-11-09-redesign.md, Wave 3). Secondary screens (raid page, Roster,
History, Calendar, Character, Guides, Tiers, BiS, account, the planner) belong to another lane and are
out of scope here: mention them only if a main-tab interaction lands on them.

Branch: lane/mobile-review-backfill. Fixture character: Gorvash, Orc Fury Warrior 58, Officer of
Olympus XXVII; alt Yellfury, Hunter 31, no spec reported. The app is dark only, as the site.

What changed since round 1, and what a ruling keeps: round-2-decisions.md (read it first; contest a
KEEP if you disagree, with the reason). Your round-1 verdict is in `<persona>.md` beside this file.

## Captures (round 2)
| Set | Where |
|---|---|
| iOS step captures (iPhone 17 Pro sim, fixture app, XCUITest; `-t0`, `-t400`, `-settled`, `-pressed`; `light-*`, `ax5-*`) | design/reviews/2026-11-10-mobile-backfill/ios/ (round 1 kept in ios-r1/) |
| Android step captures (Pixel 9 class emulator, fixture app, adb; `light-*`, `fs2-*` at font scale 2) | design/reviews/2026-11-10-mobile-backfill/android-r2/ |
| Android mocked states (Roborazzi, incl. `-x2_0`) | apps/android/app/src/test/screenshots/parity/android-*.png (+ `.txt` text goldens) |
| Android flow steps (real controller, a service that echoes the answer: sign-up, Home sign-up, deep links, saving, failed, offline) | apps/android/app/src/test/screenshots/flows/ |
| iOS flow and parity snapshots | recorded on CI; until they land, judge iOS from the step captures and the code (apps/ios/AppTests/Flows) |
| Round-5 app mocks (updated this round) | design/mocks/app/ios-*.html, android-*.html, regions/ |
| Web counterparts at 390 px | design/reviews/2026-11-10-mobile-backfill/web/ (README.md) |
| Push fixtures | docs/contracts/fixtures/app/push.*.json (raid-posted, raid-short, raid-starting, guild-applicant) |

## Fixture limits (not app behaviour)
- In the fixture APP the sign-up write answers one canned raid, so a tap on Late snaps back to In
  (with `Saved.`). The flow tests use a service that echoes the answer: judge sign-up from those.
- Feed filters and Show older are server-side; the fixture returns one page for every query.
- The fixture raid has no comp target, so the app tallies read `Healers 6`; the mocked states
  (parity goldens) carry a target and read `Healers 5 of 8 · 3 short`.
