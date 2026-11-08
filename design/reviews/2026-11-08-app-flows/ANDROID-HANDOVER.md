# Android build lane: handover

State of `apps/android` against `design/specs/2026-11-08-app-flows.md` (with `HANDOFF.md` beside this file), for the next Android lane.

## Done (on main)
| Step | What a player gets | Where |
|---|---|---|
| 1 Push | FCM on sign-in, seven channels, RAID_SIGNUP answered in the background, Late's minutes, the server's confirmation | `push/`, `raids/push/` |
| 2 Navigation | Four tabs with badges (Raid: `RaidsFeature.badge` from `GET /v1/guilds/{id}/raid-badge`; Activity: unread), the profile button and menu | `navigation/`, `menu/` |
| 3 Home | One head (crest, name in class colour, facts; under 60 `6 levels to band 40`), then `/v1/home` `sections` in order: Next upgrades (top three, Best run line from the cached BiS and loot files, left out beside a raid card), the next raid card or line answered in place (Late opens 15/30/60 min, `SignupController.pressLate`), raid live, Last night (the activity rows the day-after preview excludes), New for you (three rows, `All activity (N) ›`), Find a guild, notice; loading, error, no character; the offline strip | `home/flows/` (`HomeComposer`, `RaidCardWords`, `HomeFlowState`, `HomeFlowRoute`, `HomeFlowScreen`), `gear/WhereToGo.kt` |
| 4 Raid (root) | Next raid card, Needs you (officer, `6 open`, the newest pending applicant named and opened), raid night (live card, roles with names, Ready check from `GET .../ready?raid=`), no raid, no guild (guilds recruiting your spec on your faction, Apply on the site; Claim), under 60, signed out | `raids/flows/` (`RaidTabPresenter` + `RaidTabExtras`, `RaidTabScreen`), `RaidsFeature.refreshExtras` |
| 4 Not answered | The officer's raid page (`/app/raid/{g}/{r}?view=unanswered`, from Preview mention and the raid.short/roster.pending/raid.late_change pushes): filters, Mention the short role / Mention all with the preview (both open the raid page on the site until api-ping-role), rows healers first with `last_week`, Mark In / Bench held 4 s on the row with Undo, Discord DM (Discord's DMs until `discord_dm_url`), Healers on the bench with what fills the gap | `raids/flows/NotAnswered*.kt`, `raids/logic/ShortRole.kt` |
| 4 Applicant | Score, sim, item level, Raid ready, guild history, their note, Accept into guild / Decline; after Accept `/ginvite` with Copy | `raids/flows/Applicant*.kt`, `RaidFlowService` |
| 4 Raid night | `GET /v1/me/raid-nights/{raid}`: phase before/live draws the raid page on raid night (countdown with your answer, Your sign-up with Late's minutes, who is coming against the comp target, Before you pull); after, Your raid night with nullable log fields | `raids/flows/RaidNightPage*.kt`, `RaidNightRoute.kt`, `RaidNightModel.answer` |
| Ruling | Talents opens from the menu's Talent planner row; Home's Your gear page is deleted (the Gear tab replaces it) | `menu/MenuScreens.kt` |
| 5 Activity | The inbox, filters, day groups, unread, Mark all read | `activity/` |
| Links | App Links for every F6 page | `AndroidManifest.xml`, `MainActivity.kt` |

The flows primitives every new screen uses: `flows/FlowLayout.kt` (FlowPage, FlowHead with a crest or a `lead` such as `FlowHeadIcon`, facts that wrap, FlowSection, FlowBlock), `flows/FlowRows.kt` (FlowRow with leads Icon/Item/Check/Crest, FlowItemRow, FlowButton, FlowButtonRow, FlowWideButton, FlowLink), `flows/FlowCountdown.kt` (countdown, tally, `lineBox`), `flows/FlowFigures.kt`, `flows/FlowBackBar.kt` (a pushed screen's back and share). Home's `RaidAnswerButtons` and `LateChips`, `RaidCardWords.countdown`/`roleWord`, `RaidIcons`, `ReadyRows`, `ReadyList` are shared. Measured against the mocks with `design/mocks` boxes; parity tests `home/flows/HomeParityTest.kt` (13 states), `raids/flows/RaidTabParityTest.kt` (8), `raids/flows/RaidPagesParityTest.kt` (raid-unanswered, raid-applicant, -accepted, raid-night-live, -signup). A sideways scroller wraps its items in `ParityScrollRegion` so every one counts.

## Next
6. Gear: summary line with Change, Where to go tonight (`WhereToGo.picks` grouped by `source.kind`/`name`: instance, profession, reputation), Every slot (`home/GearModel` and `GearScreen` hold the old Your gear rows; reuse or delete them, `HomeScreenshotTest` still draws them), the item sheet with `GET /v1/guilds/{id}/items/{item}/needs` (`ApiEndpoints.getGuildItemNeeds`). Mocks `gear`, `gear-out`, `item-sheet`.
7. Glance widgets, LiveNotifier, the QS tile, restyled (mocks `widgets-guild`, `-noguild`, `-open`).
8. The secondary screens restyled, among them the in-app Notification settings (mock `notification-settings`; the menu row opens Android's own channel screen today) with the new `officer_dms` toggle ("Let officers DM me on Discord").

## Coming in land/api-ping-role (the coordinator sends the final shapes)
- `POST /v1/guilds/{id}/raids/{raid}/ping` with optional `{"role"}`, the count mentioned: wire `NotAnsweredActions.mentionRole`/`mentionAll` in `NotAnsweredRoute` (today both open the raid page on the site) and confirm `Mentioned 2 healers`.
- `ApplicantCharacter.race` and `.faction`: pass `origin` (`Alliance Dwarf`) to `ApplicantPresenter.content` (the parity fixture already shows it).
- `RaidNotSignedRaider.discord_dm_url` (opted-in members only): the row's Discord DM opens it; hide the button when null.
- `officer_dms` on `GET/PATCH /v1/me/notifications`: the toggle above.

## Open
- The old Raids root views in `raids/ui/RaidsTab.kt` (`Loaded`, `NextRaidCard`, NoGuild/SignedOut) are reached only for Loading, Failed and Waiting; the old raid page `raids/ui/RaidPage.kt` still serves a plain raid link. Delete what the restyle does not reuse in step 8.
- Raid night's times use the player's zone (the night summary carries no guild time zone).
- 37 copy keys of the old Home answer are unused on Android but still used by iOS (`home.answer.*`, `home.next.*`, `home.logs*`, `home.rankings*`, ...): delete them from `en.json` once iOS Home is rebuilt.
- `home-60-noavatar` and `home-out` have no Android parity test yet (the first differs only in the profile button; the second is the unchanged signed-out panel).

## Gotchas
- **Toolchain**: `JAVA_HOME=/opt/homebrew/opt/openjdk@21 ANDROID_HOME=/Users/jh/Library/Android/sdk ./gradlew ...`; run only touched tests; `:app:lintDebug` is strict (NewApi: `Duration.toMinutesPart` is API 31; unused drawables).
- **Measuring a mock**: render the HTML in `design/mocks/app/` with Playwright and dump every element's box relative to `<main>`; that is how the primitives' paddings were set (rows 10/6, item rows 8, section head 24 + 10 + rule, block 24 above, section 28 above). Never commit a Mac render.
- **Trailing lambdas**: `FlowRow`'s last parameter is `onClick`; pass a trailing control as `trail = { ... }`.
- **Text taller than its line** (Barlow 40/42 clock): use `Modifier.lineBox(style.lineHeight)`.
- **Goldens are recorded on CI, never locally** (`gh workflow run android.yml --ref <branch> -f record=true`, `gh run download <run> -n android-snapshots -D apps/android/app/src/test/screenshots`). `ParityCheck` now captures the image before comparing, so a local `recordRoborazziDebug` leaves the capture to read even when the regions fail; delete those PNGs before committing (the `.txt` goldens are renderer independent).
- **Copy**: append keys to `apps/shared/copy/en.json` keeping its hand formatting (a new blank-line group), no `%` (pass `8%` as an argument), then `node web/scripts/gen-mobile-copy.mjs`.
- **Scratchpad**: prefix files with `android3-lane-`; earlier lanes used `android-lane-` and `android2-lane-`.
- **Gradle in a worktree**: pass `--no-watch-fs`, or a new source file can be missed (`No tests found`).
- **core:auth on CI**: `SessionManagerTest`'s dead-token watcher runs on `Dispatchers.Default` (UNDISPATCHED); a watcher on the test scheduler starved under the wall-clock wait and timed out on every CI verify.
- **Launch mode** skips `android.yml` on `land/**`; run the record on a non-land branch first and land only a green tree.
