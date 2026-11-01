# Mobile Wave 3 F15: system integration (2026-11-01)

One lane, both apps. Spec: `design/specs/2026-10-31-mobile-wave3.md` §5.7, W3-C15. Boards:
`design/mocks/2026-10-31-mobile-wave3/renders/f15-system-{ios,and}`. It also carries the F9, F11 and F12 follow-ups and two F13 follow-ups.

Feature switch: **`system`** (`apps/shared/features.json`, default off). It gates Siri and App Shortcuts answers, the Control, Spotlight indexing, Android dynamic shortcuts and the Assistant features, and You › Offline data. Share actions, the strips, haptics, the one static store and the follow-up wiring are not gated by `system`; each follows its own screen's switch.

## API

No new endpoint and no migration. Intents read `GET /v1/me/widget` (F14) and the raid view, and sign up through the existing `PUT`. The only API changes:
- `api/internal/mail/resend_test.go` dials a closed local server, so it no longer depends on DNS. A proxy answered the old unresolvable host with a 502.
- `api/openapi.yaml` claims the new compare fixture `builds.get.compare.json` under `x-fixtures`.

## Apps

| | iOS | Android |
|---|---|---|
| Voice | App Intents in one `AppShortcutsProvider`: sign up In/Late/Out, next upgrade snippet, next raid, who is coming, open BiS. The voice phrases preset `askFirst`, so only a voice run confirms. `Which raid?` and the spec question are disambiguations. A Control Center control `In for the next raid`. One sign-up path with F14's widget intent | Four static launcher shortcuts. `Sign up In` is a hidden, non-exported activity that toasts. Up to four dynamic shortcuts: the current character, the next raid (the widget's), two watched. Assistant `OPEN_APP_FEATURE` |
| Search | `IndexedEntity` characters (own up to 9, first 50 watched), the next 5 raids and saved guides, with circular class crests. The index is deleted at sign-out | Dynamic shortcuts in launcher search, removed at sign-out |
| Share | Character, guide, tier list and public or unlisted report, each with its `share.*` sentence. Build shares keep F13's text | the same |
| Strip | One strip, three states (offline, stale with `Retry`, new data with `Update`). One pure decision function. Stale shows on every timed cache: Home and raids 5 min; rankings, character, reports and sims 10 min; roster and history 5 min. New data is extended to BiS, Talents and the item page | the same |
| Haptics | One injectable facade with the spec's table. F13's planner haptics go through it. A recording test covers each flow | the same |
| Store | F12's `StaticDataStore` and F9's cached client are one store, with one folder, one build swap, one delete and sizes per category. F13's planner reads through it | the same |
| Offline data | You › Offline data: sizes per category, `Keep my classes ready` and `Download on Wi-Fi only`, `Update now` (disabled offline, with its reason), `Remove downloaded data` with confirmation, the empty state and the footer | the same |

The copy is in `apps/shared/copy/en.json`: `offline.stale`, `offline.retry`, `offline.needsConnection`, `offlineData.*`, `intent.*`, `shortcut.*`, `spotlight.*`, `share.*` and the shared-build compare `plan.cmp.*`. The keys added in the build are listed in the spec's F15 copy keys.

## Follow-ups closed

- **F11:**
  - Rankings and player deep links open signed out and with no character.
  - `Open in planner` is one chevron row at the foot of Builds seen. It opens `/app/planner/{class}` with the class only.
- **F12:**
  - Search's `Rankings` link and its Watching group open the F11 screens.
  - `Search again` on not found reopens Search with the name filled in.
  - Search names never break mid-word at the largest text.
- **F9:**
  - The search button is on the Guides root bar.
  - The guides cache and F12's store are one store.
- **F13:** `Compare with mine` on a shared build is shown when you have a talent export of that class. It reuses F6's compare, worded by a "compare against" value (`plan.cmp.*`).

## Merge with F13

origin/main at a63ea8da7 was merged as 99bbaacc7. `features.json` keeps both `system` and `planner`. `en.json` keeps both, and the generated copy, features and links were regenerated with the scripts. F13's planner then moved onto the one store and the one haptics facade.

## Reviews

The captures are in `design/reviews/2026-11-01-mobile-integration/build/{ios,and}/`, at default and largest text (iOS AX5, Android 2.0). They include real system captures: iOS share sheets, the Shortcuts run, the Control and the remove alert; the Android launcher menu, the toast and the share chooser.

- **Round 1:** both HOLD. The MUSTs:
  - the raid count disagreed (24 against 25);
  - names, guide dates and saved sizes differed between surfaces;
  - text clipped or broke mid-word at AX5;
  - Spotlight used a person glyph instead of the crest;
  - the iOS Guides capture had no search button;
  - `Update now` did not look disabled offline;
  - acceptance captures were missing.
- **Round 2:** both HOLD on cross-platform fixture mismatches: the character page's gear and gap, the compare story, and the iOS gear icons in the snapshot harness. All three are fixed; the compare now uses a full 51-point shared build.
- **Rulings:**
  - The confirm keeps the platform order (Cancel, then Sign up); the board and spec are updated.
  - me.json's "Obnoxious Yell" stays; it is the owner's real alt.
  - The Item icons row shows its size only, as the board does.
  - Android primary labels stay in capitals.
- **Round 3:** both **SHIP** (`ux-verdict-build.md`, `player-verdict-build.md`). The open SHOULDs are listed there: the Builds seen row layout differs between the apps, the iOS AX5 Search chips, and the Compare button padding at AX5.

## Gates

These were run on the tree merged with origin/main (bef2cb599), each on its own with its exit code checked:
- `make -C apps/ios test`: TEST SUCCEEDED, exit 0.
- `./gradlew test :app:verifyRoborazziDebug lint assembleDebug`: exit 0.
- `npx vitest run src/lib/mobile src/styles/mobile-tokens.test.ts` (codegen drift): 39 passed, exit 0.
- `go test -count=1 -p 1 ./...` in `api/` against a fresh private database (`forever_test_f15`): exit 0.

## Turning it on

Add `system` to `APP_FEATURES` on the API deploy. No migration and no API change are needed. Siri phrases and the static launcher shortcuts are declared in the binaries either way; with `system` off they answer `Forever Sixty cannot do that yet.` or open the app.

## Gaps for the on-device pass

1. Siri itself, the Assistant, and the Android emulator launcher's app-shortcut search cannot be captured on the simulator or emulator. Our own views and the registered capability (unit test) stand in.
2. Signed in, the Pixel launcher gives two of its menu slots to dynamic shortcuts, so two static shortcuts drop out of the long-press menu. This is the platform's limit.
3. There are no VoiceOver or TalkBack recordings.
4. Background calendar or index refresh only runs when the app runs.
5. `gen_wave3.py` stops on the F9 board, because Flame Shock entered the Elemental loop on 70707. The F15 board's confirm swap was applied in the generator source and in its render.
6. An orphaned `static-store` folder may remain on development installs after the store merge.
