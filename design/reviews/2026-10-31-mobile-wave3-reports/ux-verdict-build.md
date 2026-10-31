# F10 Your reports, build verdict (iOS and Android)

Reviewer: ux-designer, 2026-10-31. Reviewed what renders: the 57 captures in `build/` against `design/specs/2026-10-31-mobile-wave3.md` §5.2 (F10.1 to F10.8), §4, §6 to §9, §11 and the boards `f10-reports-ios.png`, `f10-reports-and.png`. Known and intentional items (fixture story, missing iOS 26 glass buttons, non-tappable names, roster head count) are not counted.

## Verdict: HOLD

One MUST. Everything else is SHOULD or polish. Fix the MUST, re-capture `reports-list-large-ios.png`, and this ships.

## MUST

1. **`reports-list-large-ios.png`: the list loses its side margin at AX5.** The large title, the `This week` and `Earlier` labels, the cards and the footer sit about 6 pt from the screen edge (x of about 18 px on a 1206 px capture). Default is 16 pt, and `reports-list-large-and.png`, `report-kill-large-ios.png` and `meters-damage-large-ios.png` keep 15 to 17 pt. Card borders and row text nearly touch the bezel. Cause is likely an inset that scales or a width taken from the wrong container. The 16 pt gutter must hold at every text size (§11 largest text).

## SHOULD

1. **`reports-list-large-ios.png` and `reports-list-large-and.png`: row furniture does not scale.** The 36 pt crest drops to roughly 24 pt next to a three-line title. The 16 pt visibility glyph is unreadable and floats alone on its own line (Onyxia, Monday). The chevron is tiny and top-aligned. Spec §9 asks for rows that stay legible. Fix: scale crest and glyph with text up to a cap, centre the chevron, and keep the glyph beside the tag or meta, not orphaned.
2. **`reports-list-large-ios.png`: title weight inversion.** `YOUR REPORTS` renders no larger than the row titles (`Blackrock Spire`). The large title must stay the biggest thing on the screen. Cap the body scale or raise the title.
3. **`report-kill-large-ios.png`: `FIGHTS` and `3 BOSS PULLS` break onto two lines**, and the count wraps to `3 BOSS / PULLS`. Spec F10.2 puts the count right-aligned on the label line. At AX5 stack the label over the count and never split the count.
4. **`report-kill-large-ios.png` and `report-kill-large-and.png`: fight-row dots and chevrons are tiny and top-aligned** against two- and three-line text. Centre them on the first text line or the row. `Kill · 4:28 · 40 players · 3 deaths` leaves an orphan `deaths` on Android.
5. **`meters-damage-large-ios.png`: the spec icon stays 28 pt-ish small and sits above the name's cap height**, not centred on the two-line row (§9: value under the name). The head dot is 8 pt beside a 2-line title. Scale the icon with text up to 40 pt and centre it on the row.
6. **`meters-damage-large-and.png`: the `Damage taken` segment wraps to two lines and is taller than its siblings**, so the control is lopsided. Give all three segments the tallest height, or stack as iOS does.
7. **`meters-flagged-and.png` (and iOS): the flagged sheet still shows the three segments.** They do nothing. Spec F10.3 says the flagged report is `This report cannot be shown.` Remove the segments in that state. The failed state keeps them correctly (Try again retries the active segment).
8. **Acceptance coverage gaps (spec §5.2 Acceptance captures).** Missing: a report opened from a deep link; `report-coachwriting-ios`, `report-error-ios`, `reports-offline-ios` (parity with Android); the trash view on Android (`meters-trash-and` exists, but no Android healing or taken view is among the files beyond the three listed, confirm); pressed row state on either platform.

## Polish

1. **`meters-taken-ios.png` and `meters-taken-and.png`:** bars scale to the top value, so the 1,210 DTPS tank makes rows 3 to 11 sliver bars. The slivers (rank 3 to 11) end under the rank numeral and clip it (`4`, `5`, `6`). This is spec-correct (value over the top value) and Details! behaves the same. Option: give the numeral a 28 pt gutter before the bar starts.
2. **`meters-damage-ios.png`:** on bar-end rows (rank 5, 9, 11) the value numeral starts exactly at the bar edge, so the numeral straddles bar and ground. Legible, but a 4 pt right inset on the bar would settle it.
3. **`reports-list-large-*`:** `LIVE` leading dot touches the letter (`·LIVE`). Keep the 6 pt gap at every size.
4. **`report-kill-ios.png` and siblings:** the gold `REPORT` kicker is not in spec F10.2 (meta line plus tag only). It is harmless and consistent, but it adds an element. Keep or drop in the spec, not in code.
5. **`home-logs-ios.png`:** the `Logs` row, like `Talents` and `Sims`, shows no chevron; spec §5.2 and the board draw one. Shared Home pattern, fix for all three together.
6. **Sheets (`meters-*-and.png`):** no drag handle on the Android sheet.

## Region by region, what matched

- **F10.1 list** (`reports-list-ios/and`): large title, `This week` and `Earlier` labels, one quiet card per section, 36 pt ringed crest, zone in 600, `{d Mon} · {n} boss(es) · {n} kill(s)`, `Live` with dot and `Updated 20 s ago`, gold `Processing`, red `Failed` with `could not be read`, glyphs lock, users, globe, link in the right order, chevron, footer sentence verbatim. Matches.
- **F10.6 empty, F10.7 loading, error, signed out** (`reports-empty-*`, `-loading-*`, `-error-*`, `-signed-out`): copy verbatim, one secondary button, four skeleton rows, gold primary only on sign-in. Matches.
- **Offline** (`reports-offline-and`, `report-offline-ios`): the v1 strip sits above the content and the cached data shows. Matches.
- **F10.2 report** (`report-kill-*`): zone title, `Monday 14 December 2026 · logged by Obnoxious Yell`, `GUILD` tag, the one warm `Your parses` card with crest, `Onyxia`, `Kill 4:28 · Obnoxious Yell, Fury`, `712 DPS` in mono, `91` in purple (75 to 94 floor correct) with `percentile`, gold `New best`. Fights card: dots red, red, green, grey; `3 boss pulls` excludes trash; trash folded to one row per zone (`Trash · Onyxia's Lair`, `6 pulls · 7:40`). Footer verbatim. Matches.
- **Two zones, wipes, parses failed, live, processing, failed, not available, free, premium, coach writing** (`report-two-zones-*`, `-wipes-*`, `-parses-failed-*`, `-live-*`, `-processing-*`, `-failed-*`, `-not-yours-*`, `-free-*`, `-premium-*`, `-coachwriting-and`): all copy verbatim. Free player sees no coach card and no prompt. Premium shows the five sections, tag and footer, the note printed as given. `2 boss pulls` with only wipes is correct.
- **F10.3 meters** (`meters-damage-*`, `-healing-*`, `-taken-*`, `-trash-*`): title follows the segment, head row dot, name and meta, 11 rows then `Show all 40`, spec icon, class-coloured name, mono value with `DPS`, `HPS`, `DTPS` unit in micro, 22% class bars (checked: yellow bar reads rgb 69,70,47), gold 3 pt edge and `--row` ground on your row, sorted high to low, trash title `Trash · Onyxia's Lair` with no parse numbers, loading is 8 skeleton rows, failed words and `Try again` verbatim. Matches.
- **§9 largest text, meters:** 4 rows with the value under the name, as specified. Matches.

## Checklist (spec §5.2)

- [x] F10.1 states: list, live, processing, failed, empty, loading, error, signed out
- [x] F10.2 kill, only wipes, two zones
- [x] F10.3 each segment and the trash view; loading, failed, flagged copy
- [x] F10.4 Premium and free (no card)
- [x] F10.5 live, processing, failed, not available, offline
- [x] Units on every figure; parse ladder colours
- [ ] Largest text on F10.1 iOS (MUST 1; SHOULD 1, 2)
- [~] Largest text on F10.2 and F10.3 (SHOULD 3 to 6)
- [ ] Deep-link capture
