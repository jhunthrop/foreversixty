# Android captures, round 1 (2026-11-10)

Device: the `foreversixty-test` emulator (Pixel 9 class, 1080 x 2424, 420 dpi, API 36), the debug APK of
origin/main 5536d28d0 in fixture mode (contract fixtures from docs/contracts/fixtures/app). Driven with
`adb` (tap, swipe, screencap) step by step. The parity goldens in
apps/android/app/src/test/screenshots/parity/ hold the mocked states (loading, error, empty, signed out,
officer, live, 2x font) and are part of the review set.

## Fixture limits (not app behaviour)
- The sign-up write answers one canned raid (`In`, two sign-ups), so a tap on Late in the fixture app
  snaps back to In and Who is coming drops to two names. The real flow, with a server that echoes the
  answer, is in `apps/android/app/src/test/screenshots/flows/` (Roborazzi, driven by the new flow tests).
- The feed filters and Show older are server-side; the fixture returns the same page for every query.
- Portrait URLs do not resolve in fixture mode (no network), which is the same as a Battle.net render
  that fails in production.

## Steps
| File | Step |
|---|---|
| a-launch-dark | cold launch, signed in (Gorvash, Fury Warrior 58, Olympus XXVII) |
| home-01-top, home-02-scrolled, home-03-bottom | Home at rest and scrolled to the end |
| home-signup-02-late-tap-immediate, home-signup-03-late-settled | Home card: tap Late, at once and settled (fixture snaps back) |
| raid-01-top | Raid tab root |
| raid-signup-02-late-t0, -03-late-t400, -04-late-settled | Raid tab head: tap Late at 0, 400 ms, settled |
| raid-05-scrolled | Raid tab scrolled: Before you pull, the guild rows |
| raid-06-roster-pushed, raid-07-back-to-root | Roster row, then system Back (the root returns at the top, scroll lost) |
| raid-10-guild-sims-opens, raid-11-back-from-web | Guild sims row opens the site in Chrome, signed out there |
| raid-12-return-to-app | back in the app |
| gear-01-top | Gear tab (paperdoll list, band line with Change) |
| gear-04-item-sheet-opening, gear-05-item-sheet, gear-06-sheet-dismissed | tap a slot: sheet opening, open, dismissed with Back |
| gear-07-change-picker, gear-08-picker-scrolled | the band line's Change: the list picker sheet |
| activity-01-top .. -06 | Activity: All, then each filter chip, the chip row scrolled |
| activity-07-markread-t300, -08-markread-after | Mark all read: at 300 ms and after (badge clears) |
| activity-09-older-t300, -10-older-after | Show older |
| activity-11-pull-refresh, -12-refreshed | a pull-down on Activity: nothing happens (no pull to refresh) |
| drawer-01-opening, -02-open | the profile drawer opening and open |
| drawer-03-refresh-t300, -04-refresh-after | the stamp's Refresh: pressed, then after (no visible change) |
| drawer-05-scrolled | the drawer scrolled to its foot (content runs under the status bar) |
| drawer-06-switch-t300, -07-switched-home | Switch character to Yellfury (Hunter 31) |
| leveler-raid, leveler-gear, leveler-activity | the tabs as Yellfury (Gear still shows the Fury Warrior list) |
| drawer-10-switched-back | back to Gorvash |
| drawer-11-planner-pushed, -12-raid-tab-after-push, -13-home-tab-after-push | drawer Planner, then the Raid and Home tabs: Your builds covers every tab |
| light-home, light-raid, light-gear, light-activity, light-drawer | the system in light mode (the app is dark only, as the site) |
| fs2-home, fs2-home-scrolled, fs2-raid, fs2-gear, fs2-activity, fs2-drawer | system font scale 2.0 |
| notif-01-shade, -02-expanded | the raid-posted push (debug `fs.debug.raidPush`), collapsed and expanded |
| notif-03-late-action, -04-late-30 | the push's Late action, then How late? 30 min |
| notif-05-tap-through | tapping the answered notification: lands on the Raid tab root, not raid 7 |
| link-01-app-raid-2-7 | the app link /app/raid/2/7 opens the raid page |
