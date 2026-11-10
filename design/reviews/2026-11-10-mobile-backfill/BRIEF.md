# Mobile backfill review brief (round 1)

Scope: the four tab roots (Home, Raid, Gear, Activity), the profile drawer, and the raid
notification, on iOS (SwiftUI, apps/ios) and Android (Compose, apps/android). The tabs are four by
design (docs/superpowers/plans/2026-11-09-redesign.md, Wave 3). Secondary screens (raid page, Roster,
History, Calendar, Character, Guides, Tiers, BiS, account, the planner) belong to another lane and are
out of scope here: mention them only if a main-tab interaction lands on them.

Branch: lane/mobile-review-backfill. Fixture character: Gorvash, Orc Fury Warrior 58, Officer of
Olympus XXVII; alt Yellfury, Hunter 31, no spec reported. The app is dark only, as the site.

## Captures
| Set | Where |
|---|---|
| iOS step captures (iPhone 17 Pro sim, fixture app, XCUITest) | design/reviews/2026-11-10-mobile-backfill/ios/ (`<step>-t0`, `-t400`, `-settled`, `-pressed`) |
| Android mocked states (Roborazzi, Pixel 9, incl. `-x2_0` font scale) | apps/android/app/src/test/screenshots/parity/android-*.png |
| Android flow steps (sign-up, Home sign-up, deep links; real controller, echoing service) | apps/android/app/src/test/screenshots/flows/ |
| iOS mocked states (CI-recorded goldens) | apps/ios/AppTests/**/__Snapshots__/ (Home, Raids, Gear, Activity, Menu suites) |
| Web counterparts at 390 px, dark and light, plus web interaction steps | design/reviews/2026-11-10-mobile-backfill/web/ (README.md lists each) |
| Round-5 app mocks | design/mocks/app/ (m-ios-*, m-android-*), the redesign plan above |
| Interaction inventory | design/reviews/2026-11-10-mobile-backfill/interactions.md |
| Android emulator walk notes (captures being re-taken after fixes) | design/reviews/2026-11-10-mobile-backfill/android/README.md |

## Fixture limits (not app behaviour)
- In the fixture APP (emulator, simulator) the sign-up write answers one canned raid, so a tap on
  Late snaps back to In. The Android flow steps use a service that echoes the answer; judge the
  sign-up behaviour from those and from the code.
- Feed filters and Show older are server-side; the fixture returns one page for every query.

## Already found, fix in progress (do not re-report; do say if a fix looks wrong)
Android: a drawer screen covered every tab; the answered push opened the tab root; Gear kept the
previous character's list; empty portrait ring; no Saved./error/offline line on the Raid tab head
and Home card; offline answers look enabled; Home's Late minutes not shown after choosing; push
Late/Bench needing a spec became In; no pull to refresh on Home, Raid, Activity; drawer under the
status bar; Raid root scroll lost after Back; Gear Change link target 20 dp; Gear item names
ellipsized; drawer Refresh gives no feedback; 1.5x Home rows keep the trail at the right.
iOS: the Raid tab showed `Raids start at 60` to a level 58 guild member; the profile button drew the
person glyph instead of the class crest; Home said `Every slot matches` for a character with no spec.
