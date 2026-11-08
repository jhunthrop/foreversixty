# iOS build lane 2: handover

State of `apps/ios` against `design/specs/2026-11-08-app-flows.md` after lane 2's step 3, for whoever continues steps 4 and 6 (read `HANDOFF.md` and `ANDROID-HANDOVER.md` beside this file for the shared gotchas).

## Done
| Step | What a player gets | Where |
|---|---|---|
| 3 Home | `/v1/home` sections in the server's order under one head (crest, name, `Level 60 Fury Warrior · Olympus XXVII`, a leveler's levels to the next band): raid live, the raid-day card, Last night, Next upgrades with Best run, the next raid line with its answers, New for you, Find a guild, notices; loading, error, no character | `App/Home/Flow/` (HomeBlocks, HomeFlowContent, HomeFlowPresenter, HomeFlowModel, HomeFlowView, HomeFlowLive), `App/Home/HomeScreen.swift` |
| 3 the next raid card | countdown, `as Thandrel, Damage`, In/Late/Bench/Out with the current one selected, Late's chips (`late_minutes`), the tally against the comp target, `N things before you pull` (ready check), an officer's unanswered (short role) and applicants | `App/Raids/Views/NextRaidBlock.swift` (NextRaidCardView, NextRaidLineView, RaidAnswerButtons, LateMinutesChips); content type `NextRaidContent` |
| 3 Best run | computed from the cached BiS and loot files (no endpoint) | `App/Bis/BestRun.swift` |
| 3 captures | `CaptureLaunch` signs in and opens a tab plus a menu row (DEBUG `-FSMenu <row>`) | `UITests/CaptureLaunch.swift` |

New design-system pieces: `FSCountdown` + `FSClockPart`, `FSTally`, `FSMeter`, `FSSkeleton` (`Sources/DesignSystem/Countdown.swift`), `WrapLayout(lineSpacing:)`, `ItemLineView` (`App/Items`, `ui/ItemRow`).

Parity: `AppTests/Home/HomeFlowTests.swift` (13 home states, no EXPECTED_DRIFT) with `HomeFlowTestData` (the mocks' scenario as content values).

## Next, in the spec's order
4. **Raid tab** (mocks ios-raid, -officer, -live, -no-raid, -none, -leveler, -out, -unanswered, -applicant, -applicant-accepted, -night-live, -night-signup): replace `RaidsTabContentView`'s guild head with Next raid (reuse `NextRaidCardView` and `HomeFlowPresenter.nextRaid`), Needs you (`Section('Needs you', meta: '6 open')`, the badge's number), the guild pages (Roster, Calendar, Loot history, Guild sims) as rows; raid night (`Raid night` with roles named, `Ready check` checklist from `GET .../ready?raid=`, generated `ReadyCheck`/`ReadyItem.whereToGet`); Not answered (filters, `Mention the 2 healers`, rows with Mark In + Undo, Bench, Discord DM, `Healers on the bench`; `RaidNotSignedRaider.role` and `last_week`); the applicant (`GuildOfficerApplication.sheet`, Accept / Decline, `/ginvite` with Copy); no guild: `GET /v1/guilds/recruiting` with Apply, Claim your guild; under 60: `Raids start at 60`. Home's raid actions (`HomeScreen.flowActions`) all open the Raid tab today: route Preview mention / the applicant / Open raid night to their screens once they exist.
6. **Gear** (mocks ios-gear, -gear-out, -item-sheet): summary line with Change (the picker, preselected), Where to go tonight (generalise `BestRun`'s dungeon index to every source kind: instance, profession, reputation), Every slot (Worn or the gain), the item sheet with `GET /v1/guilds/{id}/items/{item}/needs`. Delete `GearScreen` / `GearView` / `GearModel` in `App/Home` once Gear replaces them; `HomeSnapshotTests`' tooltip cases go with the old sheet.

## Open items
- `answer_by` is always null (no guild stores a cutoff): Home prints the raid time alone; the mock fixture dropped "answer by" (spec appendix A still quotes it).
- `zone_icon` and Last night's loot `gain`/`unit` arrive with api-raidnight (optional); until then `RaidIcons` mirrors `api/internal/zoneicons` and the loot line stops at the slot.
- `RankingsStore.homeLine`, `ReportsStore.homeLine` and their Copy helpers are dead now Home's rows are gone: delete them once the restyle branch and this one are both on main (see IOS-RESTYLE-HANDOVER.md).
- `SurfaceGuardTests` (restyle sub-lane) ratchets per file: lower or delete the Home/Raid/Bis allowance lines when those files drop a retired style.

## Gotchas
- Local tests: `apps/ios/build/ios2-test.sh LOG Suite...` (git-ignored) builds once and runs the suites under the simulator lock; other lanes queue on the same lock, so a run can wait 20 minutes.
- Goldens only from CI: push a non-land branch, `gh workflow run ios.yml --ref <branch> -f record=true`, download `ios-snapshots` into `apps/ios/AppTests`, commit, then a plain run must be green. ios.yml and android.yml are skipped on `land/**` in launch mode.
- A copy placeholder named after a Swift keyword (`{as}`) breaks the generated Copy.swift; `%` is not allowed in copy (pass it in the value).
