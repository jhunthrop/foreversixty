# Mobile Wave 3 F13: The talent planner (2026-11-01)

One lane, both platforms, the API change and the shared rules. Spec: `design/specs/2026-10-31-mobile-wave3.md`
§5.5 and §10 (W3-C9 to W3-C11). Boards: `design/mocks/2026-10-31-mobile-wave3/renders/f13-planner-{ios,and}`.

## Shared pieces

| Piece | Where |
|---|---|
| Feature switch `planner` (default false) | `apps/shared/features.json`, generated into both apps |
| Copy `plan.*` (editor, refusals, share confirm, viewer, Your builds) | `apps/shared/copy/en.json`; the refusals are the site planner's own `messages` (`web/src/lib/planner/rules.ts`), plus the app's level-budget line `No points left. Take one off first.` |
| The tier refusal, changed at the source | Review round 1 (wow-player MUST): the site said `Tier 4 of Elemental needs 20 points in Elemental first`, a 0-based row index a player reads as tier 5. Now `Needs 20 points in Elemental. You have 14.` (the spec's sentence) in `rules.ts`, the API's validator (`api/internal/builds/validate.go`), both apps and the vectors. Because the sentence now carries the tree's running count, `canRemovePoint` compares violations by rule (`tier <tree> <need>`, else the sentence), not by text; both apps port the same |
| Golden vectors | `apps/shared/fixtures/planner-rules.json`, written by `web/scripts/gen-planner-rules.mjs` (bundles `web/scripts/mobile-codegen/planner-golden.ts` with esbuild) from the site planner's `canAddPoint`, `canRemovePoint`, `validateOrder`, `pointsPerTree`, `pointsAvailableAt`, `decodeFS1` and `orderFromRanks` over the active build's shaman, warrior and druid trees. Seeded walks plus a targeted prerequisite case per talent; the generator fails if a refusal kind goes missing. `src/lib/mobile/planner-rules-codegen.test.ts` fails CI on drift |
| Deep links (W3-C11) | `apps/shared/links.json`: `/app/planner/{class}?code=&race=` (`planner`), `/app/builds` (`builds`), and `/b/{id}` (`sharedBuild`), the first route outside `/app/`. A route marked `site: true` is a real site page; the generator emits it as `site_pages` in `web/app-link-paths.json`, and `web/scripts/app-links.mjs` admits exactly those besides `/app/*` (`web/app-links.json` paths `["/app/*", "/b/*"]`, the AASA regenerated) |

**F11's planner link.** Rankings' `Open in planner` opens `DeepLink.planner(className: <class slug>, code: nil, race: nil)` on iOS
and `DeepLink.Planner(className, null, null)` on Android: the editor on that class (its saved draft, else empty).
`builds_seen` carries only a tree split, which is not a build, so there is no code to pass. A row that ever carries an
FS1 code passes it as `code` and the editor loads it as one legal order.

## API

- **W3-C9 `DELETE /v1/builds/{id}`** (`api/internal/builds`): session or device token; `204`; `404 not_found` when the
  build is not on the caller's list (another player's, nobody's, unknown), never `403`. It clears `user_id` on the
  caller's own row only: the record and its `/b/{id}` link stay, and saving the same build again puts it back on the
  list (the apps' Undo). `builds.user_id` is already nullable (migration 0015), so **no migration**. Store, handler
  and contract tests; OpenAPI `removeBuild`; fixture `builds.delete.error-not-found` (a 204 has no body, so there is
  no `builds.delete` fixture, the convention every other 204 follows).
- **W3-C10** `/data/{build}/races.json` and `/data/{build}/combos.json` as `app-static` operations (`RaceInfo`,
  `RaceClassCombo`; `ForeverChange` hoisted to a shared schema for classes and races), fixtures `cdn.races`,
  `cdn.combos` held to the generator's files by `cdn_test.go`.
- Contract doc: `docs/contracts/app-v1.md` (DELETE, the planner's rules, races and combos).

## Apps

| | iOS (`apps/ios/App/Planner`) | Android (`app/.../planner`) |
|---|---|---|
| Rules | `PlannerRules.swift`: the site's rules and FS1 decode, refusals from `plan.refuse.*` | `logic/PlannerRules.kt`, `Fs1.kt`, `PlannerCopy.kt` |
| Editor (F13.1, F13.2) | `PlannerState`, `PlannerEditorModel`, `PlannerEditorView/Screen`, `PlannerSheets` | `PlannerEditor`, `PlannerEditorModel`, `ui/PlannerEditorScreen`, `ui/PlannerSheets` |
| Share (F13.3) | confirm sheet, `POST /v1/builds`, the system share sheet | the same, `ACTION_SEND` |
| Viewer (F13.5) | `SharedBuildView`: the build's own `tree_version`, cached for offline | `ui/SharedBuildScreen` |
| Your builds (F13.6) | `YourBuildsView`: swipe, `DELETE`, `Removed.` + Undo (re-POSTs the build) | `ui/YourBuildsScreen` |
| Shared tree | F6's `TalentTreeView` takes an optional editing handler and a gold `available` state; one component, two call sites | F6's `TalentTree` the same |

- Tap adds, long press removes, `+`/`−` in the strip, accessibility actions `Add a point` / `Remove a point`, arrow
  keys and `+`/`-`. Haptics: tick, long press, reject, confirm (unit tested). Undo keeps 100 steps. A level holds
  `level - 9` points and never drops below spent + 9. The draft is saved per class on every change and restored after a
  force quit (unit tested). A shared build is copied, never edited.
- Stepper strip: a fixed 96 minimum, rank text clamped to 3 lines and refusals to 2 (tap opens F6's rank sheet with the
  full text); at the largest text it leads the scroll with Reset and Share pinned, and a refusal scrolls it into view.
  `+` dims (still tappable) when maxed or none left, `−` at rank 0. No cell is gold at 0 left. iOS tightened the
  planner's row pitch to 60 pt so all seven shaman rows fit above the strip on iPhone 17.
- Entry points, all behind `planner`: Talents (F6) `Plan a build` (replacing the web link) and `Your builds` (signed
  in); a guide's `Open this build in the planner` (the raid build's FS1 code, else the leveling build's); the links.
  With the switch off a `/b/` link opens the browser.
- iOS also: `UITests/PlannerCaptureTests` (the Reset alert from the running app, `ForeverSixtyCaptures` scheme); the
  debug `-FSOpenLink` launch argument opens planner links.
- Android also: core `BrandButton` gained an opt-in `iconScalesWithText`; the Reset dialog is drawn in a
  `BasicAlertDialog` (no scrim dismiss); the commented App Links note in the manifest names `/b/` beside `/app/`.

## Reviews

Captures (default and largest text, both platforms, fixture-driven; the apps are dark-only) in
`design/reviews/2026-11-01-mobile-planner/build/{ios,and}/`.

- Round 1: ux-designer HOLD (strip height moved under the thumb; largest text hid the strip; no Reset alert; no F13.4),
  wow-player HOLD (0-based `Tier 4`; largest text; iOS last row under the strip). Fixed as above.
- Round 2: both HOLD only on two iOS captures still recording (AX5 refusal, Reset alert); Android polish done.
- Round 3: ux-designer **SHIP** (`ux-verdict-build.md`), wow-player **SHIP** (`player-verdict-build.md`). Earlier
  rounds are kept beside them (`*-r1.md`, `*-r2.md`).
- Open SHOULDs: the iOS Your builds AX5 separator wrap; the loaded-build note in the strip clearing after the first
  edit; `Share anyway` naming unspent points; button casing across platforms; faint edge lines on iOS chips in
  snapshots; the viewer's share toolbar item absent from iOS snapshots (iOS 26 toolbar items do not render in view
  snapshots, as F9 found).

## Gates

On the tree merged with origin/main, each run on its own with its exit code checked:

- `npx vitest run src/lib/mobile src/lib/app-links src/lib/planner/rules.test.ts src/lib/planner/share.test.ts src/lib/planner/fs1.test.ts`: exit 0 (codegen drift for API, copy, features, links and the planner vectors; app links).
- `go test -p 1 ./internal/builds/ ./internal/contract/ ./internal/server/ ./internal/app/ ./internal/rankings/` on a private database (`forever_test_f13`): exit 0.
- `make -C apps/ios test`: TEST SUCCEEDED, exit 0 (741 app tests in 66 suites plus the package suites; every planner golden replayed).
- `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`: exit 0 (769 app tests).

## Turning it on

The API change ships with the next API deploy (no migration). Add `planner` to `APP_FEATURES` after it. Universal
links and App Links for `/b/*` start working when the owner fills in the Team ID and signing fingerprints
(`web/app-links.json`); until then the links route inside the app.

## Gaps and follow-ups

1. **`Compare with mine` is hidden** on both platforms: F6's compare copy and status line are written for a guide band,
   and a shared build on that side needs new copy and a status variant (spec F13.5; both reviewers accept it as a
   follow-up).
2. **F13.4, the system share sheet**, is system UI a snapshot cannot draw; it shares `https://foreversixty.gg/b/{id}`
   with the title. No `card.png` preview is attached yet. Capture it in the on-device acceptance pass.
3. **On-device acceptance pass:** VoiceOver and TalkBack recordings (a locked talent must speak its requirement), the
   draft after a force quit, F13.5 with an old `tree_version`, the swipe states, a build over a legal 51, and the
   §11 budgets.
4. **A guide's FS1 code is read against the active trees** without the site's remap by talent id
   (`remapTreeRanksByTalentId`). Guides are written for the active build, so this holds until a tree changes shape
   between builds.
5. **Data for the accuracy reviewer:** Elemental Mastery's rank text reads `... by 0.` at 0/1 and names Lava Burst;
   check it against the Forever client's own string.
6. **Strip clamping vs spec §9** ("nothing is cut with an ellipsis"): the ux-designer asked for the clamp so the strip
   holds 96; the full text is one tap away on the rank sheet. Below the accessibility sizes only.
