# Mobile Wave 1, lane F4: guild and raids (2026-10-31)

Branch `wt/mobile-f4-raids`. Not pushed. Built from `design/specs/2026-10-30-mobile-apps.md` §5.4 (revision 2) on both platforms.

## Commits

| Commit | What |
|---|---|
| cb95b5775 | `feat(contract)`: the raid fields the sign-up needs, Your raid night, F4 fixtures, and the `raids.*` copy |
| 290c5454a | `feat(contract)`: the officer's view after Lock (`guild.raid.officer-locked`), plus the `raids.when.*` copy |
| 038394fab | `feat(android)`: guild raids and one-tap sign-up |
| c35d297cc | `feat(ios)`: guild raids and one-tap sign-up |
| 79c9693c9 | `fix(android)`: the top bars no longer take the status-bar inset twice |
| d2c58d016 | `docs(design)`: review captures for both platforms |

The branch also carries `752720d42`, a merge of origin/main that brings in the shared device lock.

## What was built (both platforms)

- **Raids tab (F4.1, F4.13 to F4.17).**
  - The guild head, then the next raid card. The card has the date block, invites and pull in mono, the countdown to invites, the count and its breakdown, and your sign-up segments.
  - "Who is coming" pushes the raid page at the roster on iOS. On Android it opens the top of the page.
  - The officer's "waiting for approval" row, then Also planned with one chip per row, then up to three raid nights.
  - States: no guild, waiting, no raid planned, error, offline (cached on disk, shown under the strip), signed out, and the first-load skeleton.
- **Raid page (F4.2, F4.6).**
  - Eyebrow, status chip and guild name, then the title.
  - The long date, then `Invites · Pull · Ends`, then the countdown, which ticks each minute.
  - The guild-time line, shown only when the offsets differ.
  - The count block, then the sign-up card, the classes strip (5 + 4 at large text), and the notes as styled text (bold, https links).
  - The roster: Tanks, Healers, DPS, Late, Bench, Out, each row with its spec icon. Empty roles have their sentence. Then Not signed up.
  - Refetches every 30 s inside the raid's last 24 h. Add to calendar, Share, and Open on foreversixty.gg.
  - The done raid shows the report link and the Your raid night row.
- **Sign-up (F4.3 to F4.12).**
  - One tap, drawn optimistically. Success shows `Saved.` for 4 s; a failure restores the previous answer with the API's sentence (F4.11).
  - Every refusal maps to its sentence: locked race, cancelled race, rate limited, tools ended, forbidden, offline.
  - Haptics: selection on the press, success on the save, error on a refusal.
  - "Which spec tonight?" shows the dashed pending outline. Closing the sheet drops the press, Out never asks, and Change spec re-sends the current answer.
  - The `moved` field puts `{name} moved to Bench.` on the other character's block.
  - Locked, ended, cancelled and offline states; waiting characters get no controls; the note field with its counter from 120 characters.
  - The reminder ask comes after a first In or Late, at most twice, the second 30 days later.
  - The segments switch to 2 x 2 at AX1 on iOS and at 1.5x on Android.
- **Officer tools (F4.18 to F4.21),** drawn only from `viewer.can`.
  - Lock or Unlock, with the toast `Sign-ups locked. {n} In and {b} on the Bench were told.`
  - Copy the list, in the site's format.
  - A place button on each row, with the spec's menu table: Move to In or Bench, Remove, Add as In or Bench. A move is optimistic and restores the row on failure.
  - Officer read-only (`tools_ended`): the buttons draw disabled, with the plan sentence.
- **Your raid night (F4.22, F4.23).**
  - Kills, wipes and new bests, then the loot awarded to you with its quality border and colour, then personal bests with the parse colour.
  - Coach notes show only when the API sends them; a free player sees no section and no prompt.
- **Deep links and notifications.**
  - `/app/raid/…`, `/app/raid/…/night` and `/app/guild/…` select the Raids tab and push onto a fresh stack. `spec=ask` presses In, which opens the spec sheet when no spec is saved.
  - iOS: `RaidNotificationCenter` registers `RAID_SIGNUP` (In, Late, Out), `RAID_ASK_SPEC` and `RAID_CHOOSE`. Actions are answered in the background.
  - Android: `RaidActionReceiver` uses `goAsync`, and `RaidNotificationActions.actions(…)` builds the buttons for lane F7's notification builder.
  - On both platforms, an action reads the raid, keeps the character's stored note (the PUT clears it otherwise), sends the PUT, and posts the server's answer text in place of the question.
  - Which character answers: the payload's `character_key`, else the character already signed, else the only verified one. Otherwise the action opens the raid page.
- **Feature gate.** Both apps gate the screens on `raids`, which is on in debug and fixture builds.

## Contract added (openapi.yaml + fixtures, regenerated): what each handler must add

The guild contract tests list these paths as pending (`raidViewPending` in `api/internal/guilds/app_contract_options_test.go`). Drop a path there when its handler starts sending it.

1. **`GET /v1/guilds/{id}/raids/{raid}`, every raid view** (the list's `planned[]`, the sign-up PUT's `raid`, and lock, unlock and cancel). Lane A3 / the guilds API:
   - `invite_at` (nullable instant). This needs a raids column, the web form's field, and the scheduler: `raid.starting` fires at `invite_at` minus the lead.
   - `signups[].spec_name` and `signups[].spec_icon`, from specs.json `name` and `icon`.
   - `not_signed[].role`, `spec_name` and `spec_icon`.
   - `viewer.characters[].spec_name` and `spec_icon`, of `raid_spec` or else of `spec`.
   - `viewer.characters[].spec_options[]` (`{spec, name, role, icon}`): the class's specs in catalogue order.
   - `viewer.can`: `lock` for a verified officer or leader who may run status actions, and `move` for one who may PUT for others. Take both from `guilds/capabilities.go`, never from a rank.
   - `viewer.tools_ended`: true for an officer of a guild whose plan ended.
   - `viewer.raid_night`: true when `/v1/me/raid-nights/{raid}` has a summary for this viewer.
2. **`GET /v1/guilds/{id}/raids`.** Add `waiting_for_approval`, sent only when the viewer's `can` includes `approve`.
3. **New `GET /v1/me/raid-nights/{raid}`** (schema `RaidNightSummary`; fixtures `me.raid-night`, `me.raid-night.free` and `me.raid-night.error-not-found`).
   - The night's kills and wipes, the viewer's loot with its awarding officer and what it replaces, personal bests (boss, DPS, percentile, kill time), and `coach_notes` (null without `log_coach`).
   - 404 when there is no matched log or no viewer character in it.
4. **Raid push payloads (lane A2).**
   - `data.character_key`, using the spec's pick rule.
   - `data.ask_spec`, as "true" or "false".
   - `data.answers`: a JSON string of `{in, late, out, moved, locked, cancelled, failed}`, each `{title, body}` in the player's locale. `moved.body` carries `{name}`.
   - Category `RAID_ASK_SPEC` when `ask_spec` is true, and `RAID_CHOOSE` when no character can be picked.
   - The device token must be admitted on the sign-up PUT (`APP_NOTIFICATIONS_ENABLED`) and on lock and unlock (A1).

New fixtures:
- **Raid states:**
  - `guild.raid.officer`
  - `.officer-locked`
  - `.locked`
  - `.cancelled`
  - `.no-invite`
  - `.ask-spec`
  - `.two-characters`
- **Raids list:** `guild.raids.empty`
- **Sign-up refusals:**
  - `guild.signup.error-raid-cancelled`
  - `.error-rate-limited`
  - `.error-plan-required`
- **Account:** `me.waiting`

`guild.raid` and `guild.raids` now carry the boards' full roster: 3 tanks, 6 healers, 14 DPS, 1 Late, 1 Bench and 3 Out, plus three planned raids.

## Tests

| Check | Exit | Result |
|---|---|---|
| `make -C apps/ios test` (shared simulator) | 0 | TEST SUCCEEDED: 85 app-hosted tests in 14 suites, with every package suite passing. F4 adds 46 logic tests and 31 snapshots: iPhone 17 at the default size and at AX5. |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | 0 | 46 F4 logic tests (Robolectric), 31 F4 goldens (Pixel 9, default and font scale 2.0), lint clean, debug APK built |
| `go test ./internal/contract` and `-run AppContract ./internal/guilds ./internal/auth ./internal/server` (private DB `forever_gc_f4`, since dropped) | 0 | ok |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 0 | 28 passed. `web/node_modules` was symlinked for the run and removed before each commit. |

The test list is shared across the platforms: test names match one for one, covering times, facts, roster, sign-up, officer, tab, notifications, links and notes.

The fixture helpers throw. A missing fixture fails only its own test, with no `try!`, force unwrap or `fatalError` in the iOS test support. This fixes the host crashes of 17:20 to 17:28: the fixtures were being read from a bundle folder named `app`, not `Fixtures`.

## Captures

`design/reviews/2026-10-31-mobile-raids/`: 66 files.
- `<screen>-<state>-ios.png` from the snapshot suite on the iPhone 17 simulator.
- `-and.png` from Roborazzi at Pixel 9.
- `emulator-*-and.png`: the debug app on the shared emulator, with real spec and item icons from the CDN.

Screens and states:
- Raids tab: loaded, loaded-xl, no guild, waiting, no raid, error, offline, signed out, loading.
- Raid page:
  - member, full, xl, no-invite
  - unsigned, spec pending, reminder ask
  - locked, cancelled, offline
  - save failed, rate limited, race locked, moved
  - done, officer, officer locked (with the toast)
- Spec sheet: ask, ask-xl, change.
- Raid night: premium, free.
- Android only: the move menu.

## Concerns

1. **Icons in the snapshots.** The snapshot suites draw spec and item icons as their bordered tiles, because tests have no network. The emulator captures show the real icons. There is no live simulator capture on iOS: universal links need the Team ID (lane O), so a raid link cannot be opened in the simulator yet.
2. **No captures yet:**
   - the calendar sheet;
   - the iOS move menu (SwiftUI `Menu` does not snapshot);
   - the More menu;
   - keyboard focus on a segment;
   - the VoiceOver and TalkBack recordings.

   Wave 2 or the reviewers need to drive these by hand.
3. **Shared files touched additively.** Other lanes will merge against these:
   - `RootView.swift` (the raids case, `onOpenURL` and the notification hookup);
   - `ForeverSixtyApp.swift` (the notification delegate install);
   - `project.yml` (the crests folder and the calendar usage string);
   - `ForeverAppScaffold.kt`, `MainActivity.kt`, `AppGraph.kt` and `AndroidManifest.xml` (the receiver and `POST_NOTIFICATIONS`);
   - `libs.versions.toml` (Coil 3.0.4) and build-logic (a `forever.crests` plugin).

   Both lane F7 and F4 need the notification delegate (iOS) and the receiver (Android). F7 should call into `RaidNotificationCenter` and `RaidNotificationActions` rather than replace them.
4. **Design-system additions.** These may collide with F2 and F3:
   - `Sources/DesignSystem/Theme+Game.swift` holds the class, quality and parse colours, plus the type roles;
   - Android's `RaidType` and `RaidColors` live in the raids package.

   If F2 or F3 adds the same names, fold them into one design-system file at merge.
5. **Departures from the spec:**
   - **Mono font.** JetBrains Mono is not bundled, so the system mono stands in on both platforms. It reads wider than the boards.
   - **Tabs.** The tab bar is still the shell's five tabs, and the screen title reads "Raids" instead of the character selector. Both belong to lane F1.
   - **Raids badge dot.** Not built; it needs the shared tab bar.
   - **Contract over spec on push ids.** The push category and action ids follow the contract (`RAID_SIGNUP`, `signup_in`, …), not the spec's §6.3 names (`RAID_ANSWER`, `IN`).
   - **Background PUT on Android.** It runs inside `goAsync` with an 8 s budget, not a WorkManager job.
6. **Answer posting needs lane F7.**
   - Android posts its answers on channel id `raids`, which lane F7 must create.
   - iOS asks for notification permission from the reminder ask, but token registration is lane F7's.
7. **Layout gap.** The next raid card leaves a visible gap under the segments when no status line shows: the reserved 20 pt line plus the 48 pt target of "Who is coming". The reviewers may want it tightened.
8. **Captures folder size.** It is 27 MB, mostly the tall full-page and xl PNGs.
9. **Fixture folder in the iOS bundle (P1 bug, not fixed here).**
   - XcodeGen copies `docs/contracts/fixtures/app` into the app as a folder named `app`, not `Fixtures`; the `name:` key has no effect on a folder reference.
   - So `FixtureTransport.bundled(in:)` finds nothing and falls back to the bundle root, and fixture mode in the running iOS app likely answers `missingFixture`.
   - The crests folder lands as `crests` for the same reason; `CrestView` reads that name.

## Fix round (2026-10-31): the reviewers' HOLD

The verdicts are `design/reviews/2026-10-31-mobile-raids/ux-verdict.md` and `player-verdict.md`, committed with this round. The coordinator accepted player #1 (API fields pending) and #4 (notification answers need lane F7) as sequenced. Everything else is fixed here.

1. **iOS fixture mode works.** Lane F3's bundle-name fix (aedc28eb6: the fixtures land at `app/`) is taken as `defc1c7b9`, with its iOS fixture-mode hunks only.
2. **Moved to Bench draws its state (F4.12).**
   - `guild.signup.moved` now answers with the real R3 view: Yellfury signs In and the server benches Obnoxious Yell.
   - The moved character's block draws Bench checked, with `Obnoxious Yell moved to Bench.` under it, on both platforms.
   - A test on each platform checks both blocks.
3. **Large text (iOS AX5, Android 2.0).** Nothing breaks inside a word.
   - Names stay on one line and shrink rather than split.
   - Caps labels drop their tracking and shrink.
   - The status chip never wraps, and the guild name drops to its own line.
   - The who line stacks the crest, then the name, then Change spec.
   - The count sits above a full-width bar.
   - On the next raid card, the date block stacks above the title.
   - In Also planned rows, the date block stays 56 wide and the chip moves under the title.
   - The guild head and Your raid night's figures stack.
   - New large-text captures cover the officer view and Your raid night on both platforms.
4. **Composition for a raid leader (player #5).** The ruling's composition block sits at the top of the roster:
   - each class by name with its count, with zeros muted as "0";
   - role counts, shown against the raid's targets when it sets them (`Tanks 3/4`);
   - the Classic buff and utility row (Fortitude, Mark of the Wild, Arcane Intellect, Blessings, Shadow Protection, Soulstone, Rebirth) as icons, with missing ones muted and named (`Missing: Blessings`).

   It is a computed `composition` block on the raid view (`RaidComposition` in openapi.yaml, on every raid fixture). The apps only word it. The app-side classes strip is removed, so nothing computes composition in the app.
5. **Card gap.** The next raid card draws the status line only when it has text, and the link sits just under the segments. The raid page still reserves the line.
6. **Type.** The count sentence and the countdown are in Barlow; only clock digits stay mono (`6:45` mono, `pm` Barlow). The iOS Raids title is inline.
7. **Fixture accuracy.**
   - The raid-night loot is now Sapphiron Drape (17078, epic, back), an Onyxia drop in `loot.json` (`raid:onyxias-lair`). Dragon's Blood Cape is a real item but not an Onyxia drop.
   - It replaces Cape of the Black Baron (13340, a real rare cape).
   - The coach notes are reworded in raiders' terms.
   - The notes drop "Phase 2 prep".
   - Barrow Deeps stays (a real Forever raid).
8. **Captures that needed real taps.**
   - iOS: a debug-only UI test target, `ForeverSixtyUITests`, run with `make -C apps/ios captures` and never by `make test`. It drives the app in fixture mode through debug-only launch arguments (`App/Raids/Debug/RaidDebugLaunch.swift`) and captures the More menu, the calendar sheet, the officer move menu, and the raid notification answered with In from the long-press.
   - Android: a debug-only hook (`RaidDebugPush`) posts the fixture raid notification with its actions. On the shared emulator, In from the shade sent the sign-up PUT and replaced the notification with `You are In for Onyxia's Lair` (`notification-answer-in-and.png`). The More menu is captured live too.

### Added for the post-launch API work

- **Raid view `composition`** (on every raid view): `classes[]` (all nine, counting In and Late), `roles` (In counts), `targets` (the raid's role targets, or null; this needs a raid field and the web form when wanted), and `buffs[]` (`key, name, icon, present, count`) from the classes signed In or Late.
- **Icons the CDN export needs:** the icon set must include `spell_nature_regeneration` (Mark of the Wild) and `spell_holy_fistofjustice` (Blessings), which today's export lacks.

### Fix round tests

Each command ran on its own, with its exit code checked.

| Check | Exit | Result |
|---|---|---|
| `make -C apps/ios test` | 0 | 87 tests in 14 suites. F4 has 46 logic tests and 33 snapshots, including the new AX5 snapshots of the officer view and Your raid night. |
| `make -C apps/ios captures` | 0 | 3 UI tests passed: the More menu and calendar sheet, the officer move menu, and In from the notification long-press |
| `./gradlew test :app:verifyRoborazziDebug lint assembleDebug` (JDK 21) | 0 | 46 F4 logic tests, 33 goldens, lint clean, debug APK built |
| Go `-run 'AppContract\|OpenAPI\|Fixture'` on guilds, auth, contract and server (private DB `forever_gc_f4b`, since dropped) | 0 | ok |
| `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` | 0 | 28 passed |

### Fix round captures (in `design/reviews/2026-10-31-mobile-raids/`)

- **Moved state:** `raid-page-moved-{ios,and}` shows Bench checked on Obnoxious Yell, with the moved line.
- **Large text:** `*-xl-{ios,and}` covers the Raids tab, the raid page, the officer view, Your raid night and the spec sheet.
- **Tap-driven, iOS live simulator:**
  - `raid-page-more-menu-ios`, `raid-page-calendar-sheet-ios` and `officer-move-menu-ios`;
  - the notification path, `notification-delivered-ios`, `notification-actions-ios` (the long-press) and `notification-answer-in-ios` (the posted `You are In for Onyxia's Lair`).
- **Android emulator:**
  - `emulator-raid-page-more-menu-and`;
  - `notification-delivered-and` and `notification-answer-in-and`.

### Fix round concerns

1. **Icons on the live iOS simulator.** The game icons did not load there (the composition buffs and the roster's spec icons draw as their tiles), while the Android emulator loads them from the CDN. The iOS app reads the active build from the live site, so this needs a look on a device with the live data. The snapshots draw tiles by design.
2. **Not done this round** (UX shoulds outside the coordinator's list):
   - JetBrains Mono is still not bundled.
   - The toast inset is unchanged.
   - Keyboard focus and VoiceOver/TalkBack recordings still need a hand-driven pass.
3. **Two buff icons the export lacks.** `spell_nature_regeneration` (Mark of the Wild) and `spell_holy_fistofjustice` (Blessings) are true game icons that today's icon export lacks. The API lane must export them with the composition handler.
