# Android build lane: handover

State of `apps/android` against `design/specs/2026-11-08-app-flows.md` (with `HANDOFF.md` beside this file), for the next Android lane.

## Done (on main)
| Step | What a player gets | Where |
|---|---|---|
| 1 Push | FCM on sign-in, seven channels, RAID_SIGNUP answered in the background, Late's minutes, the server's confirmation | `push/`, `raids/push/` |
| 2 Navigation | Four tabs with badges (Raid: `RaidsFeature.badge` from `GET /v1/guilds/{id}/raid-badge`; Activity: unread), the profile button and menu | `navigation/`, `menu/` |
| 3 Home | One head (crest, name in class colour, facts; under 60 `6 levels to band 40`), then `/v1/home` `sections` in order: Next upgrades (top three, Best run line from the cached BiS and loot files, left out beside a raid card), the next raid card or line answered in place (Late opens 15/30/60 min, `SignupController.pressLate`), raid live, Last night (the activity rows the day-after preview excludes), New for you (three rows, `All activity (N) ›`), Find a guild, notice; loading, error, no character; the offline strip | `home/flows/` (`HomeComposer`, `RaidCardWords`, `HomeFlowState`, `HomeFlowRoute`, `HomeFlowScreen`), `gear/WhereToGo.kt` |
| 4 Raid (root) | Next raid card, Needs you (officer, `6 open`), raid night (live card, roles with names, Ready check), no raid, no guild (Claim), under 60, signed out | `raids/flows/` (`RaidTabPresenter`, `RaidTabScreen`), wired in `RaidsRoute`'s root |
| 5 Activity | The inbox, filters, day groups, unread, Mark all read | `activity/` |
| Links | App Links for every F6 page | `AndroidManifest.xml`, `MainActivity.kt` |

The flows primitives every new screen uses: `flows/FlowLayout.kt` (FlowPage, FlowHead, FlowSection, FlowBlock), `flows/FlowRows.kt` (FlowRow with leads Icon/Item/Check, FlowItemRow, FlowButton, FlowButtonRow, FlowWideButton, FlowLink), `flows/FlowCountdown.kt` (countdown, tally, `lineBox`). Measured against the mocks with `design/mocks` boxes; parity tests `home/flows/HomeParityTest.kt` (13 states), `raids/flows/RaidTabParityTest.kt` (7 states).

## Next
4. The rest of Raid: the raid page restyled as `raid-unanswered` (filters Not answered/In/Late/Bench/Out/All, Mention the short role, Mention all, the preview, Mark In with Undo on the row, Healers on the bench), the applicant (`raid-applicant`, `-accepted`), raid night (`raid-night-live`, `-signup`). The old root views in `raids/ui/RaidsTab.kt` (`Loaded`, `NextRaidCard`, the NoGuild/SignedOut messages) are now only reached for Loading, Failed and Waiting: delete what the restyled page does not reuse.
6. Gear: summary line with Change, Where to go tonight (`WhereToGo.picks` grouped by `source.kind`/`name`: instance, profession, reputation), Every slot, the item sheet with `GET /v1/guilds/{id}/items/{item}/needs` (endpoint generated: `ApiEndpoints.getGuildItemNeeds`).
7. Glance widgets, LiveNotifier, the QS tile, restyled.
8. The secondary screens restyled.

## Waiting on the API (asked 2026-11-08; the coordinator routed them to app-api)
- Guilds recruiting a spec (Raid, no guild): `RaidTabView.NoGuild.guilds` is empty until it lands; the section hides.
- Applications in the app client (list, accept, decline): Home's and Needs you's applicant line (`OfficerItems.applicantLine`) is null; the row opens the guild page.
- Ready check: `GET /v1/characters/{region}/{ruleset}/{name}/ready?raid={id}` (also in `/v1/me/raid-nights/{raid}`). `RaidCardModel.ready` and `RaidTabView.Guild.ready` are empty until then.
- Not answered's history line (`last_week`) and role on `RaidNotSignedRaider`.

## Open
- Home no longer reaches the Talents and Your gear pages (`showTalents`, `showGear` in `HomeCharacterScreen` are never set): Gear replaces Your gear in step 6; Talents needs a menu row or a decision to fold into the Planner.
- 37 copy keys of the old Home answer are unused on Android but still used by iOS (`home.answer.*`, `home.next.*`, `home.logs*`, `home.rankings*`, ...): delete them from `en.json` once iOS Home is rebuilt.
- `home-60-noavatar` and `home-out` have no Android parity test yet (the first differs only in the profile button; the second is the unchanged signed-out panel).

## Gotchas
- **Toolchain**: `JAVA_HOME=/opt/homebrew/opt/openjdk@21 ANDROID_HOME=/Users/jh/Library/Android/sdk ./gradlew ...`; run only touched tests; `:app:lintDebug` is strict (NewApi: `Duration.toMinutesPart` is API 31; unused drawables).
- **Measuring a mock**: render the HTML in `design/mocks/app/` with Playwright and dump every element's box relative to `<main>`; that is how the primitives' paddings were set (rows 10/6, item rows 8, section head 24 + 10 + rule, block 24 above, section 28 above). Never commit a Mac render.
- **Trailing lambdas**: `FlowRow`'s last parameter is `onClick`; pass a trailing control as `trail = { ... }`.
- **Text taller than its line** (Barlow 40/42 clock): use `Modifier.lineBox(style.lineHeight)`.
- **Goldens are recorded on CI, never locally** (`gh workflow run android.yml --ref <branch> -f record=true`, `gh run download <run> -n android-snapshots -D apps/android/app/src/test/screenshots`). `ParityCheck` now captures the image before comparing, so a local `recordRoborazziDebug` leaves the capture to read even when the regions fail; delete those PNGs before committing (the `.txt` goldens are renderer independent).
- **Copy**: append keys to `apps/shared/copy/en.json` keeping its hand formatting (a new blank-line group), no `%` (pass `8%` as an argument), then `node web/scripts/gen-mobile-copy.mjs`.
- **Scratchpad**: prefix files with `android2-lane-`; the earlier lane used `android-lane-`.
- **Launch mode** skips `android.yml` on `land/**`; run the record on a non-land branch first and land only a green tree.
