# Secondary screens: every interaction and its capture

Lane `lane/mobile-rest`, 2026-11-10. Each row is one step a player takes on a secondary screen. **Before** and **After** name the capture that shows it. Android captures are the Roborazzi goldens under `apps/android/app/src/test/screenshots/` (the path below is relative to that folder). iOS columns are filled when the iOS half lands.

The **web** column names the live page, captured at 412 wide in `android/*-reference-*.png`, or the app mock when the page needs a signed-in guild.

**Pressed** is drawn by the primitives, not by each screen:
- `FsTopBar` and its icon buttons;
- the panel rows;
- `BrandButton`;
- the chips.

Their pressed states are captured once, in `design-system-sheet.png`, and cover every screen that uses them.

A step marked **GAP** has no capture yet. It blocks the gate until it has one.

## Shared chrome (every pushed screen)

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open a pushed screen from the menu or a link | the tab root (`parity/android-menu.png`) | the screen with the band bar: title, Back, actions | site header and page head |
| Back (the bar arrow or the system back) | any pushed screen | the tab root it was opened over | browser back |
| Text at the largest size | `*-large-*`, `*-xl-*`, `*-ax5*` | the same states at 2x | n/a |

## Guides (Guides home, guide reader, class page, tiers, what changed)

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open Guides from the menu | `parity/android-menu.png` | `guides/f9-guides-and.png` | `/guides` |
| Guides home without a character | | `guides/f9-guides-nochar-and.png` | `/guides` |
| Guides home offline | | `guides/f9-guides-offline-and.png` | n/a |
| Tap Your guide | `guides/f9-guides-and.png` | `guides/f9-guide-top-and.png` | `/guides/shaman/elemental` |
| Tap a class crest | `guides/f9-guides-and.png` | `guides/f9-class-offline-and.png`, `f9-class-large-and.png` | `/guides/hunter` |
| Guide loading, error, not cached | | `guides/f9-guide-loading-and.png`, `f9-guide-error-and.png`, `f9-guide-notcached-and.png` | n/a |
| Tap a section chip (Talents, Rotation, Stats, Races) | `guides/f9-guide-top-and.png` | `f9-guide-talents-and.png`, `f9-guide-rotation-and.png`, `f9-guide-stats-and.png`, `f9-guide-races-and.png` | Jump to |
| New data arrives, then Update, then the update fails | `guides/f9-guide-top-and.png` | `f9-newdata-and.png`, `f9-newdata-failed-and.png` | n/a |
| Offline with a saved guide | | `guides/f9-guide-offline-and.png` | n/a |
| Share the guide | `guides/f9-guide-top-and.png` | the system share sheet (platform; not drawn by us) | copy link |
| Tiers segment, then a role (DPS, Tank, Healer) | `guides/f9-guides-and.png` | `f9-tiers-dps-and.png`, `f9-tiers-tank-and.png`, `f9-tiers-healer-and.png` | `/tiers` Change |
| Tiers first open and without a character | | `f9-tiers-first-and.png`, `f9-tiers-nochar-and.png` | `/tiers` |
| What changed segment, then an entry, then an error | `guides/f9-guides-and.png` | `f9-changes-and.png`, `f9-change-entry-and.png`, `f9-changes-error-and.png` | `/changelog?show=game` |
| Search from Guides | `guides/f9-guides-and.png` | `system/f15-guides-search-and.png` | header search |

## Character and rankings

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open a character (Search, Activity, a link) | `search/f12-2-characters-and.png` | `rankings/character-gear-and.png`, `character-other-and.png`, `character-yours-and.png` | `/character/...` |
| Loading, error, not found, offline | | `character-loading-and.png`, `character-error-and.png`, `character-notfound-and.png`, `character-offline-and.png` | n/a |
| No gear export, no logs | | `character-nogear-and.png`, `character-nologs-and.png` | same states |
| Watch, then the watch fails | `rankings/character-other-and.png` | `character-watch-on-and.png`, `character-watch-failed-and.png` | Watch |
| Tap a gear item | `rankings/character-gear-and.png` | the item sheet (`tooltip-medium.png`) | the tooltip |
| Open in the planner | `rankings/character-gear-and.png` | `system/f15-char-open-planner-and.png` | n/a |
| Open Rankings, on page and off page | `parity/android-menu.png` | `rankings/rankings-onpage-and.png`, `rankings-offpage-and.png` | `/rankings` |
| Rankings menu, filters, encounters, Today on and off | `rankings/rankings-onpage-and.png` | `rankings-menu-and.png`, `filters-and.png`, `filters-encounters-and.png`, `filters-today-on-and.png`, `filters-today-off-and.png` | filter bar |
| Rankings empty, error, loading, offline, signed out | | `rankings-empty-and.png`, `rankings-error-and.png`, `rankings-loading-and.png`, `rankings-offline-and.png`, `rankings-signedout-and.png` | n/a |

## BiS picker, search, item page and tooltip

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open the BiS list | | `bis_list_default.png` | `/bis` |
| Open the picker (spec, band, faction) | `bis_list_default.png` | `bis_picker_default.png`, `bis_picker_bottom_large.png` | ControlBar Change |
| Tap an item | `bis_list_default.png` | `bis_tooltip_default.png` | tooltip |
| Loading, error, not cached, offline | | `bis_loading_default.png`, `bis_error_default.png`, `bis_not_cached_default.png`, `bis_offline_default.png` | n/a |
| Open Search, type, results | | `search/f12-1-empty-and.png`, `f12-2-all-and.png`, `f12-3-items-and.png` | header search |
| No match, offline, index failed | | `f12-6-none-and.png`, `f12-6-offline-and.png`, `f12-index-failed-and.png` | n/a |
| Open an item: loading, pick, runner-up, not listed, uses failed | `search/f12-3-items-and.png` | `f12-4-loading-and.png`, `f12-4-pick-and.png`, `f12-4-runner-up-and.png`, `f12-4-not-listed-and.png`, `f12-4-uses-failed-and.png` | item tooltip |
| BiS for a spec from the item | `search/f12-4-pick-and.png` | `f12-5-bis-for-and.png` | n/a |
| Tooltip kinds (weapon, vendor, large) | | `tooltip-weapon.png`, `tooltip-vendor.png`, `tooltip-large.png` | tooltip |

## Account (You, plan, devices, delete, about, sign in, offline data)

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open Settings (You) from the menu, free and Premium | `parity/android-menu.png` | `account_you-free.png`, `account_you-premium.png`, `premium_you-seat.png` | `/account` |
| Sign out, then confirm | `account_you-premium.png` | `account_you-signout-confirm.png`, then `account_signedout-default.png` | Sign out |
| Plan (US link, GB, guild seat) | `account_you-premium.png` | `premium_plan-us.png`, `premium_plan-gb.png`, `premium_plan-seat.png` | `/premium` |
| Devices, revoke, confirm, revoke fails | `account_you-premium.png` | `account_devices-default.png`, `account_devices-revoke-confirm.png`, `account_devices-revoke-failed.png` | `/account` devices |
| Delete account: free, Premium, guild duty, seat | `account_you-premium.png` | `account_delete-free.png`, `account_delete-premium.png`, `account_delete-guild-obligation.png`, `account_delete-seat.png` | `/account` delete |
| Type the BattleTag: mismatch, ready, failed | `account_delete-free.png` | `account_delete-mismatch.png`, `account_delete-ready.png`, `account_delete-failed.png` | same |
| About | `account_you-premium.png` | `account_about-default.png` | footer |
| Signed out, sign in, browser open, failed, 401, deleted | | `account_signedout-default.png`, `account_signin-browser-open.png`, `account_signin-failed.png`, `account_signedout-401.png`, `account_signedout-deleted.png` | Sign in |
| Import and no characters | `account_signin-browser-open.png` | `account_import-default.png`, `account_home-no-characters.png` | setup |
| Offline data: list, remove, empty, offline | `parity/android-menu.png` | `system/f15-offline-data-and.png`, `f15-offline-data-remove-and.png`, `f15-offline-data-empty-and.png`, `f15-offline-data-offline-and.png` | n/a |
| Live surfaces settings | `account_you-premium.png` | `widgets/and-live-settings.png` | n/a |
| Update required | | `update_required.png` | n/a |

## Guild and raid sub-screens

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open a raid (member, officer, not signed up) | `parity/android-raid.png` | `raids/raid-page-member-and.png`, `raid-page-officer-and.png`, `raid-page-unsigned-and.png` | guild raid page |
| Sign up: the spec sheet, then a spec change | `raids/raid-page-unsigned-and.png` | `raids/spec-sheet-ask-and.png`, `spec-sheet-change-and.png` | sign-up |
| Save fails, rate limited, offline | `raids/raid-page-member-and.png` | `raid-page-save-failed-and.png`, `raid-page-rate-limited-and.png`, `raid-page-offline-and.png` | n/a |
| Reminder ask after sign-up | `raids/raid-page-member-and.png` | `raid-page-reminder-ask-and.png` | n/a |
| Raid states: locked, full, moved, cancelled, done, no invite, race locked | | `raid-page-locked-and.png`, `raid-page-full-and.png`, `raid-page-moved-and.png`, `raid-page-cancelled-and.png`, `raid-page-done-and.png`, `raid-page-no-invite-and.png`, `raid-page-race-locked-and.png` | same |
| Officer: move a member, locked | `raids/raid-page-officer-and.png` | `officer-move-menu-and.png`, `raid-page-officer-locked-and.png` | roster tools |
| Raid night (free, Premium) | `raids/raid-page-member-and.png` | `raid-night-free-and.png`, `raid-night-premium-and.png`, `parity/android-raid-night-live.png`, `parity/android-raid-night-signup.png` | none (app only) |
| Applicant: open, accept | `parity/android-raid-officer.png` | `parity/android-raid-applicant.png`, `parity/android-raid-applicant-accepted.png` | guild applications |
| Not answered list | `parity/android-raid-officer.png` | `parity/android-raid-unanswered.png` | none |
| Roster: filters (tanks, healers, dps), sort menu, sorts, loading, error, offline, no match | `guild/f16.7-raids-tab.png` | `guild/f16.1-roster-*.png`, `guild/f16.2-no-match.png` | `/guild` Roster |
| History: page 1, page 2, more loading, more failed, empty, error, offline | `guild/f16.7-raids-tab.png` | `guild/f16.3-history-*.png` | `/guild` Loot |
| Calendar: on, off, which character, reset alert, link copied, elsewhere, failed, offline, removed, sync failed | `guild/you-calendar-row.png` | `guild/f16.4-*.png` | calendar link |
| Officer access off | | `guild/f16.5-access-off.png` | n/a |

## Sims, reports, talents and the planner

| Step | Android before | Android after | Web |
|---|---|---|---|
| Open Sims (Premium, free, lapsed) | `parity/android-menu.png` | `sims/sims-premium-and.png`, `sims-free-and.png`, `sims-lapsed-and.png` | `/sim` |
| Sims loading, error, empty, offline, page failed | | `sims-loading-and.png`, `sims-error-and.png`, `sims-empty-and.png`, `sims-offline-and.png`, `sims-page-failed-and.png` | n/a |
| Run again: running, done, failed | `sims/sims-premium-and.png` | `sims-rerun-states-and.png` | Run |
| Open a result: loading, done, gear, failed, offline | `sims/sims-premium-and.png` | `sims/result-loading-and.png`, `result-run-and.png`, `result-gear-and.png`, `result-failed-and.png`, `result-offline-and.png` | result |
| Open Logs: list, empty, error, loading, offline, signed out | `parity/android-menu.png` | `reports/reports-list-and.png`, `reports-empty-and.png`, `reports-error-and.png`, `reports-loading-and.png`, `reports-offline-and.png`, `reports-signedout-and.png` | `/logs` |
| Open a report: kill, wipes, two zones, live, processing, failed, not yours, deep link | `reports/reports-list-and.png` | `reports/report-*.png` | `/reports/...` |
| Coach: Premium, free, writing, parses failed | `reports/report-kill-and.png` | `report-premium-and.png`, `report-free-and.png`, `report-coachwriting-and.png`, `report-parsesfailed-and.png` | Coach tab |
| Meters sheet: damage, healing, taken, trash, flagged, loading, failed | `reports/report-kill-and.png` | `reports/meters-*.png` | report tabs |
| Talents: matches, differs, beyond, no build, no export, no spec, guide, restoration, loading, error, offline | `parity/android-menu.png` | `talents_*_default.png` | `/planner` |
| Talent rank sheet | `talents_matches_default.png` | `talents_rank_*.png` | tooltip |
| Planner: fresh, plan, zero, reset, undo, refusal, share | `talents_matches_default.png` | `f13-1-*.png`, `f13-2-refusal_default.png`, `f13-3-share_default.png` | `/planner` |
| Your builds, empty; shared build, not found, offline, compare | `f13-1-plan_default.png` | `f13-6-builds_default.png`, `f13-6-empty_default.png`, `f13-5-*.png`, `system/f15-shared-build-compare-*.png` | `/b/:id` |

## Gaps

These are left after this inventory. Each is either filled before the gate or named in the verdicts:
- **Pressed states.** They are drawn per primitive, so one capture per primitive (`design-system-sheet.png`) stands for every screen. Reviewers judge the primitive.
- **The system share sheet** and **the browser hand-off** (plan billing, the web breakdown, privacy and terms). These are platform surfaces, not ours.
