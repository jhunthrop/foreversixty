# iOS secondary screens restyle (flows spec section 2, build step 8): handover

Branch `ios2-restyle`. What is on the flat system with a parity test, what is left, and why.

## Done
| Mock states | Screen | Notes |
|---|---|---|
| reports, reports-empty, reports-failed | `App/Reports/ReportsListView.swift` | PageHeader with the lead and count, Upload a log (opens `/logs`), one Plate of rows (state Tag and visibility glyph under the line), EmptyState with the sample report link, LoadError. The This week / Earlier split and the row crest are gone (the mock's `row nolead`, one list newest first). |
| report-error | `App/Reports/ReportView.swift` (`.failed`, `.notYours`) | Plain `Report` title, the Plate with an inline LoadError. The loaded report itself is still the old layout (see Left). |
| devices | `App/Account/DevicesView.swift` | Signed in section, Revoke as small danger buttons, the This device tag, **Sign out everywhere else** (new: `SessionStore.revokeOthers()`, `POST /v1/devices/revoke-others`, with a confirmation). |
| offline-data, offline-data-empty, offline-strip | `App/System/OfflineDataView.swift` | Sections On this phone and Updates, labelled switches, Update now (secondary) and Remove (danger); offline, the Notice under the title. |
| update-required | `App/UpdateRequiredView.swift` | Seal, centred PageHeader, the store button. |

Also: the Planner (`YourBuildsView`) carries **Your talents against the guide**, the TalentsScreen's entry point after Home's Next rows left (parent lane request).

## Primitives added (`apps/ios/Sources/DesignSystem`)
`Plate`, `EmptyState`, `LoadError`, `ActionLink`, `Notice`, `LabelledSwitch` (`Plate.swift`); `PageHeader` takes an `actions` slot (`.ph-actions`, 16 under the title block).

## Surface guard
`apps/ios/AppTests/SurfaceGuardTests.swift`: outside `Sources/DesignSystem`, a token colour as a background (`bg` aside), `.overlay(RoundedRectangle|Capsule(...).stroke...)`, `primaryCard`, `quietCard`, `surfaceBand`, `MicroLabel`, `TagView`, `ScaledType.title/h2/stat/sec/cardHead`, `BrandButtonStyle` and `Text(... .uppercased())` / `.textCase(.uppercase)` count per file against `allowance`. A file may not exceed its count, and a file below its count fails until its line is lowered (deleted at 0), so the list only shrinks. The retired modifiers and views stay in `Surfaces.swift` / `Buttons.swift` until the last allowance line that uses them is gone (Home, Raid, Gear and the widgets still use them).

## Mocks changed
- `design/mocks/ws_app.py` `page(..., flush=True)`: the account chores (devices, delete-*, offline-data*, update-required, expired) put the head on the 16 pt gutter instead of 32 (the flows screens already did, `.ap .pc`). Widgets, live-activity and raid-night are unchanged.
- Offline data: sizes over 10 MB are whole megabytes, as `OfflineWords.size` writes them (`On this phone, 27 MB`).
- Regions and captures from app-mocks runs 42466817180 and 42467887264 (Linux runner).

## Left, and why
| Mock states | Why not in this pass |
|---|---|
| sim, sim-you, sim-running, sim-result, sim-saved | The mocks are the web simulator (paste an addon export and run in the browser, gear and talent trees, timelines; 1,700 to 12,000 pt). The app has server runs and results, no in-app engine: this is new feature work (or a mock of the app's own Sims list), not a restyle. |
| guide, guide-you, guide-loading, guide-error, guides, guides-class, guides-you, tiers, tiers-you, changes | Web page ports (the guide is 16,600 pt with 266 counted items); each is a lane-sized build. `GuidesHomeView`, `GuideBody`, `GuideParts`, `TiersView`, `ChangesView` are on the allowance. |
| report, report-player, report-coach, report-loading, report-processing | The web report (tab strip, players, charts). report-processing needs the Summary / More tabs the app does not have. |
| rankings, rankings-you, rankings-empty | `rankings` is an encounter landing the app does not have (it opens a board); the board is the web RankList. |
| planner-you, planner-guide, planner-loading, planner-error, build-shared | The web planner pages (TalentEditor, PointOrder, share). |
| character-* | The web character page (paperdoll, Forever Score, Best per encounter). |
| search, search-results | The app's scope chips sit between the field and the sections; the mock has none. Add the chips to the mock (`ws_app.py search`) or drop them, then restyle. |
| account, account-premium, account-loading, account-error, premium | The web `/account` and `/premium` pages; the app's Settings / You and Plan screens are shaped differently. |
| delete-free, delete-premium, delete-failed, delete-done | The app keeps the web dialog's deleted / kept / you lose lists and the typed confirmation (App Store 5.1.1(v)); the mock has a 3-row What goes. The mock should take the app's lists and the field before the restyle. |
| expired | No such screen yet (R16 Sign in again): new work. |

## Notes for whoever lands this
- `RankingsStore.homeLine`, `ReportsStore.homeLine` and their Copy helpers only fed Home's old Next rows; delete them (and their tests) once the Home step that deletes `HomeAnswerView` is on the branch. They still compile here because `HomeAnswerView` uses them.
- `ReportsList.sections`, `ReportsList.crestClass` and `OfflineCategory.systemImage` are deleted with their tests.
- Copy keys added: `reports.lead`, `reports.count`, `reports.upload`, `reports.sample`, `devices.signedInTitle`, `devices.thisTag`, `devices.signOutOthers` (+ `.title`, `.body`, `.failed`), `offlineData.sectionEmpty`, `plan.builds.compare`, `plan.builds.compareSub`. `reports.thisWeek`, `reports.earlier`, `reports.footer`, `reports.getCompanion` are no longer read on iOS (Android still reads `reports.footer`).

## Goldens: still to record
The new parity tests (`ReportsParityTests`, `AccountParityTests`) pass their region checks locally (the iOS 26.5 shared simulator); they only lack their text and image goldens, and the restyled screens' older goldens (`ReportsSnapshotTests`, `AccountSnapshotTests` devices, `SystemSnapshotTests` offline data, `UpdateRequiredSnapshotTests`) need re-recording. Five record runs of ios.yml on 2026-11-08 (42470289875, 42470937721, 42471409917, 42471919826, 42472315927) failed before the tests ran: GitHub answered 504 for the Sentry xcframework and the XcodeGen release downloads. Next: `make -C apps/ios snapshots BRANCH=ios2-restyle`, commit `apps/ios/AppTests/**/__Snapshots__`, then a normal `gh workflow run ios.yml --ref ios2-restyle` that must be green.
