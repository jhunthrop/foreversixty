# F15 system integration: UX build verdict

## Round 1

Reviewed: the iOS and Android captures in `build/ios/` and `build/and/`, default and largest text, against spec `design/specs/2026-10-31-mobile-wave3.md` §5.7 F15 (copy keys, Offline data, strips, share table, Haptics), §6 Strip and Search-button rows, and boards `f15-system-ios.png` and `f15-system-and.png`. The planner follow-ups (F11.4, F13.5) are not in this round.

### What matches (no action)
- Voice copy is verbatim: confirm `Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?`, signed-up answer, refusal `Sign-ups are locked. An officer can still change yours.`, next raid, who is coming (`24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.`).
- Upgrade snippet: icon, name in rare blue, `{Slot} · item level {n} · {source}`, `+3.8 DPS for Lightningbolltz. Replaces Vest of Elements.` Layout follows the board at default text.
- Three strips carry the exact spec strings; `Retry` and `Update` are gold, trailing, and not clipped at the largest text on both platforms (`f15-strip-*-large-*`). The new-data strip sits on the gold ground and reaches the guide page.
- Offline data: title, section label `On this phone, 48 MB`, five rows, both switches with their sub-lines, `Update now` (primary), `Remove downloaded data` (wipe colour), footer, and the confirmation sentence all match. No build ids or hashes. Android parity with the board is close.
- Android toast `You are In for Onyxia's Lair.` appears over the launcher; the app does not open (`f15-signup-toast-emulator-and.png`).
- Share text: `Lightningbolltz, Elemental Shaman, level 60`; `Elemental Shaman guide`; `Healer tier list`; `Onyxia's Lair · 14 Dec`, each with the right link.
- Search's Watching group shows crests, class-coloured names and `Level 60 · US PvP`; Rankings row present; Guides root bar on Android has the Search button.

### MUST (block ship)

1. **iOS confirm buttons clip at the largest text.** `f15-siri-confirm-large-ios.png` (iOS, spec F15.1 and v1 §7 Dynamic Type). `Sign up` and `Cancel` sit in fixed 100 px pills; the glyphs are cut at the bottom edge of both. Buttons must grow with text (minimum height 48 pt, no fixed height).
2. **iOS Remove confirmation breaks mid-word.** `f15-offlinedata-remove-large-ios.png` (iOS, spec `offlineData.removeConfirm`). The custom dialog wraps `Remove` as `Remov` / `e` and the two buttons have unequal heights. Stack the buttons vertically at large sizes, or use the system alert (it is correct at default size in `f15-offlinedata-remove-system-ios.png`).
3. **`Update now` does not look disabled offline.** `f15-offlinedata-offline-ios.png` and `f15-offline-data-offline-and.png` (both, spec "Writes ... stay disabled with their reason"). The reason line `Update now needs a connection.` is there; the button is still full gold. Dim it to the disabled token and mark it disabled for VoiceOver and TalkBack.
4. **iOS Spotlight characters use a generic person glyph, not the crest.** `f15-spotlight-entries-ios.png` and `-large-ios.png` (iOS, spec F15.3, tenet "one crest language"). Android's launcher shortcuts carry the crest. Every character entry on iOS must use the circular HD crest as its thumbnail.
5. **Android launcher menu shows two of the four static shortcuts.** `f15-launcher-shortcuts-emulator-and.png` (Android, spec "Static shortcuts: Next raid, Sign up In, My upgrades, Search"). `My upgrades` and `Search` are missing; two dynamic entries took their place. The current character (Lightningbolltz) is not among the dynamic entries either; `Voltaris` and `Stormwrath` appear, and the next raid is not offered as a dynamic one. Register all four statics and keep dynamic ones to what the launcher has room for after them.
6. **iOS Guides root capture shows no Search button.** `f15-guides-searchbutton-ios.png` (iOS, §6 Search button row `Search`, follow-up F15). The capture has no top bar at all: no character selector, no trailing search icon. Android (`f15-guides-search-and.png`) has both. Fix and recapture with the bar visible.
7. **Acceptance captures missing.** Spec lists: the share sheet for a character, a guide and a report on both platforms; Shortcuts-app run without confirmation; launcher search; Assistant (F15.2). Present: iOS share text screens only (no sheet), no Android share capture, no Android launcher search or Assistant capture, no iOS Shortcuts-run or Control Center capture, no `Which raid?` or spec-question snippet. Attach them, or state per item which the simulator cannot render and replace it with our own view as was done for Siri.

### SHOULD

1. **Large-text snippet breaks words.** `f15-siri-upgrade-large-ios.png`: `Leatherworki / ng` and `Lightningbo / lltz` in a narrow column beside the icon. Stack the icon above the text at accessibility sizes. Same pattern in `f15-search-watching-large-ios.png` (`Lightningboll / tz`).
2. **iOS Offline data rows at AX5 stack a 210 px tile above the text** (`f15-offlinedata-data-large-ios.png`), so two rows fill the screen. Keep the tile at 40 pt beside the text as Android does (`f15-offline-data-large-and.png`).
3. **Spotlight at largest text:** icons are vertically misaligned with the two-line titles (`f15-spotlight-entries-large-ios.png`). The board also shows a Shortcuts group (`Sign me up for tonight's raid`, `What's my next upgrade`); the build omits it.
4. **Switch knobs are missing in the iOS snapshot captures** (`f15-offlinedata-data-ios.png`, `-empty-ios.png`, `-uptodate-ios.png`, `-offline-ios.png`): blank gold pills. The real-app capture (`f15-offlinedata-app-ios.png`) is correct. Recapture so on and off are readable.
5. **Sign-up count fixture disagrees.** Next raid says `24 of 40` (`f15-siri-nextraid-ios.png`); the spec's example and the story say 25. The spec itself is inconsistent (who-is-coming 24, signed-up 25 after your In). Pick one number set: 24 before you sign up, 25 after, and use it in every capture.
6. **Confirm button order.** Build puts `Cancel` left and `Sign up` right; the board has `Sign up` left. Primary-right is the iOS norm; accept it and update the board, or match the board.
7. **Android empty state at AX scale hides `Remove downloaded data`** (`f15-offline-data-empty-large-and.png`) while iOS shows it. Show it on both or neither, disabled when nothing is stored.
8. **Android Remove dialog at 2.0** (`f15-offline-data-remove-large-and.png`) runs edge to edge with no side margin and a condensed display face for body copy. Use 24 dp side margins and the body face.
9. **Android search field at 2.0** collapses the placeholder to `Search items,` (`f15-search-watching-large-and.png`). Allow two lines or shorten at large sizes.
10. **Copy not in the spec:** `Nothing is saved on this phone yet.` and `Guides and lists save here the first time you open them.`, `Everything on this phone is up to date.`, `Not checked yet.`, `Saved on this phone: ... Guides update by themselves when the game data does.` Fine in tone; add to the spec's copy keys so tests pin them.
11. **Android dynamic shortcut labels** do not show `Forever Sixty · character` in the menu; confirm it renders in launcher search once captured (MUST 7).

Verdict (round 1): HOLD

## Round 2

Reviewed: the round-2 captures in `build/ios/` and `build/and/` against spec §5.7 F15, §6, F11.4 (`Open in planner`), F13.5 (`Compare with mine`) and `plan.cmp.*` in `apps/shared/copy/en.json`. Rulings from the brief are applied and not re-raised.

### Round 1 MUSTs

1. iOS confirm buttons clip at AX5: FIXED. `f15-siri-confirm-large-ios.png` shows full-width pills that grow with the text, `Sign up` primary above `Cancel`, no clipping.
2. iOS Remove confirmation mid-word break: FIXED. `f15-offlinedata-remove-large-ios.png` wraps whole words and stacks `Remove` (wipe colour) over `Cancel` (gold) at equal height. `f15-offlinedata-remove-system-large-ios.png` exists as the system alert.
3. `Update now` not disabled offline: FIXED on both. `f15-offlinedata-offline-ios.png` and `f15-offline-data-offline-and.png` dim the button to the disabled token, with `Update now needs a connection.` beneath it.
4. iOS Spotlight generic person glyph: FIXED. `f15-spotlight-entries-ios.png` shows circular class crests, guide and raid tiles, and the Shortcuts group.
5. Android launcher menu: FIXED for the statics. `f15-launcher-shortcuts-static-emulator-and.png` shows `Next raid`, `Sign up In`, `My upgrades`, `Search`. The signed-in menu with two dynamic slots is a platform ruling (see SHOULD 1 for ordering).
6. iOS Guides root search button: FIXED. `f15-guides-searchbutton-ios.png` and `-large-ios.png` show the character selector and a trailing 44 pt search button. The AX5 bar holds.
7. Acceptance captures: MET with documented gaps. Present: share text and system sheets (iOS character, guide), Android share sheets for character, guide and report, `f15-shortcuts-run-ios.png`, `f15-control-ios.png`, `f15-siri-which-ios.png`, `f15-siri-spec-ios.png`, launcher menu and launcher search. Siri, Assistant and per-shortcut launcher search are on-device gaps per the rulings.

### Round 1 SHOULDs

1. Large-text snippet breaks words: FIXED. `f15-siri-upgrade-large-ios.png` stacks the icon above the text. `Lightningbolltz` no longer splits. Search Watching at AX5 on iOS was not rechecked; `f15-search-watching-large-and.png` is clean.
2. iOS Offline data rows at AX5: FIXED. `f15-offlinedata-data-large-ios.png` keeps the 40 pt tile beside the text. Polish: a wrapped line can start with a bare dot separator (`· 2.4 MB ·`, `· Updated today`). Drop the leading separator when the line wraps.
3. Spotlight alignment and Shortcuts group: FIXED (group present).
4. Switch knobs in iOS snapshots: FIXED (`f15-offlinedata-offline-ios.png` shows knobs).
5. Count fixture: FIXED. 25 of 40 in `f15-shortcuts-run-ios.png`, per the ruling.
6. Confirm button order: closed by ruling.
7. Android empty state hides Remove: FIXED. It shows, dimmed. See SHOULD 2.
8. Android Remove dialog at 2.0: FIXED. Side margins and body face are right, buttons stack.
9. Android search field at 2.0: FIXED. The placeholder wraps to three lines in full (`f15-search-watching-large-and.png`).
10. Copy not in the spec: still open. Add `Nothing is saved on this phone yet.`, `Guides and lists save here the first time you open them.`, `Everything on this phone is up to date.`, `Not checked yet.` and the Guides footer to the spec's copy keys so tests pin them.
11. Dynamic shortcut labels: not visible in the launcher capture (documented gap).

### MUST (block ship)

1. **iOS character page gear has no icons and no quality colours.** `f15-char-open-planner-ios.png` and `-large-ios.png` (iOS, F11.4 and F11.5, tenet "icons, quality colours, real tooltips everywhere"). Every slot is an empty dark square and every item name is grey. Android on the same fixture (`f15-char-open-planner-large-and.png`) shows icons with quality borders and names in green, blue and purple. This is the only iOS capture of the new `Open in planner` in context, and grey names are not an image-loading limit. Either the snapshot harness is missing the item metadata, or the app ships without quality colour. Recapture with the real icon and quality data, or show a real-app capture. If the app renders this way, it is a blocker.

### SHOULD

1. **Dynamic shortcut order (Android).** `f15-launcher-shortcuts-signedin-emulator-and.png`. With two dynamic slots the menu shows `Voltaris` and `Stormwrath`. Spec order is current character first, then the first two watched, then the next raid. Rank the current character first so the two slots are Lightningbolltz and Voltaris. Confirm `Lightningbolltz` is rank 0.
2. **Android disabled `Remove downloaded data` at 2.0.** `f15-offline-data-empty-large-and.png`. The label wraps to two lines inside a box with no vertical padding, so the glyphs touch the top and bottom border. Let the button grow to the text, minimum 48 dp, with 12 dp vertical padding.
3. **Compare fixtures disagree.** `f15-shared-build-compare-ios.png` shows `3 of the shared build's points still to take.` with Elemental 30, Enhancement 3, Restoration 15. `f15-shared-build-compare-and.png` shows `2 of your 48 points differ from the shared build.` with 28, 3, 17. Pick one fixture so both platforms tell the same story. Copy matches `plan.cmp.toTake` and `plan.cmp.differ` verbatim.
4. **Gear gap fixture.** `Gear gap 36.0 points` (iOS) against `3.0 points` (Android) for the same character. At AX5 the iOS value wraps to `36.0` over `points` (`f15-char-open-planner-large-ios.png`). Use one number and, at accessibility sizes, stack the label above the value.
5. **`Open in planner` differs per platform.** iOS puts one gold chevron card below the Builds seen card. Android makes it a row inside the card. F11.4 puts it on the build row. Match the spec: a row per build, so a second build gets its own `Open in planner`.
6. **Android shared-build title.** `f15-shared-build-compare-button-large-and.png` runs `Orc Shaman 30 / 3 / 18` on one line with wide gaps around the slashes. iOS breaks the totals onto a second line. Put the totals on their own line on both.
7. **Android share capture context.** `f15-share-character-emulator-and.png` shows the signed-out welcome screen behind the sheet. Recapture from the character page. The text and link are right: `Lightningbolltz, Elemental Shaman, level 60` and the `/character/us/normal/lightningbolltz` link.

### Checklist

- Voice, snippet, refusal and who-is-coming copy verbatim: yes.
- Strips, three states, `Retry` and `Update` not clipped at the largest text: yes (round 1, unchanged).
- Offline data data, empty, offline, remove: yes on both platforms.
- Spotlight crests: yes. Launcher statics: yes. Dynamic order: no (SHOULD 1).
- Share text for character, guide, tier list, report: yes.
- Planner follow-ups render and wrap at the largest text: yes. iOS gear quality and icons: no (MUST 1).
- Both platforms tell the same story: no on the compare and gear-gap fixtures (SHOULD 3, 4).

Verdict: HOLD
Verdict (round 2): HOLD

## Round 3

Reviewed: the round-3 captures in `build/ios/` and `build/and/` against spec §5.7 F15, §6, F11.4, F13.5 and the F15 Copy keys. Rulings and lead notes applied, not re-raised.

### Round 2 MUST

1. iOS character page gear icons and quality colours: FIXED. `f15-char-open-planner-ios.png` and `-large-ios.png` show real icons, quality borders (green, blue, purple) and quality-coloured names, with gold better-item dots. Matches Android (`f15-char-open-planner-and.png`).

### Round 2 SHOULDs

1. Dynamic shortcut order (Android): FIXED. `f15-launcher-shortcuts-signedin-emulator-and.png` shows the current character first, then `Onyxia's Lair · Thu 8 pm`, under the two static slots.
2. Android disabled `Remove downloaded data` at 2.0: FIXED. `f15-offline-data-empty-large-and.png` grows to the two-line label with padding; the glyphs clear the border.
3. Compare fixtures: FIXED. `f15-shared-build-compare-ios.png` and `-and.png` both read `3 of your 49 points differ from the shared build.` and `The shared build takes 5 points you have not.`, with the same Arms 17 / Fury 32 tabs and the same dashed -1, -3, +1, +1 nodes.
4. Gear gap: FIXED. 36.0 points on both; iOS stacks label over value at AX5 (`f15-char-open-planner-large-ios.png`), Android keeps one row and fits.
5. `Open in planner`: FIXED. One chevron row inside the Builds seen card on both platforms, under the build it opens.
6. Shared-build title: FIXED. `f15-shared-build-compare-button-large-and.png` and `-large-ios.png` put the totals on their own line on both.
7. Android share capture context: FIXED. `f15-share-character-emulator-and.png` is taken over the character page; text and link are right for the fixture user (per lead note).

### MUST (block ship)

None.

### SHOULD (polish, not blocking)

1. **Builds seen row differs.** `f15-char-open-planner-ios.png` shows a 40 pt build glyph tile and spaced `30 / 0 / 21`; `f15-char-open-planner-and.png` shows no tile and `30/0/21`. Use one row anatomy. Same story, different dress.
2. **Segment selected mark.** `f15-shared-build-compare-and.png` puts a check on `Your build`; `f15-shared-build-compare-ios.png` does not. Platform convention is acceptable; confirm iOS selection does not rely on fill alone (VoiceOver trait selected).
3. **iOS Compare with mine at AX5.** `f15-shared-build-compare-button-large-ios.png` wraps to two lines with about 8 pt vertical padding. Raise to 12 pt, as the Android button.
4. **Open items carried from round 2.** Wrapped lines that start with a bare dot separator in `f15-offlinedata-data-large-ios.png`; the empty and up-to-date copy keys are now in the spec's F15 Copy keys, so tests must pin them.
5. **On-device gaps.** Siri, Assistant, per-shortcut launcher search stay on the device pass, per ruling.

### Checklist

- Siri, snippet, refusal and who-is-coming copy verbatim: yes.
- Strips, three states, `Retry` and `Update` at the largest text: yes.
- Offline data data, empty, offline, remove: yes on both.
- Spotlight crests, launcher statics and dynamic order: yes.
- Share text for character, guide, tier list, report: yes.
- Planner follow-ups render, wrap and match across platforms: yes.
- iOS gear icons and quality colours: yes.
- Both platforms tell the same story: yes.

Verdict: SHIP
