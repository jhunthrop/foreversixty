# UX verdict: mobile BiS tab (F3), iOS and Android, 2026-10-31

Reviewed: live list captures (both devices), rendered picker, offline, not-cached, error, large-text boards, against spec section 5.3 and renders f3-bis-{ios,and}. Read only; nothing re-run.

## Verdict: HOLD

The list is right and ships as built. The hold is two unproven states, not a rework.

## Blockers

1. **iOS chips and pills show stray vertical hairlines at both edges.** Seen in render-picker-default-ios (level range, faction pills), render-list-not-cached-default-ios (all three chips). Spec: a chip is a clean pill, 4.5% fill, `--color-line-card` border. The live iOS list (list-default-ios) is clean, so this may be an ImageRenderer artifact of `pillSurface`, but the iOS picker has never been captured on a device (only rendered; simctl cannot tap). Fix: capture the iOS picker and the chips on a device or simulator through the real view (drive it with a launch argument like `-FSTab`, e.g. `-FSPicker open`), at default and AX5. If the hairlines are on screen, fix `pillSurface`; if not, say so in the report.
2. **Android picker at font scale 2.0 is unproven.** render-picker-large-and shows the shell tab bar drawn over the sheet and cutting off the Level range pills; Faction and `Back to {name}'s list` are not visible. Spec: the sheet stays open and every choice is reachable. Fix: a large-text picker capture where the sheet content scrolls to Faction and Back (live on device, or a scrolled render), with the tab bar not over the sheet.

## Shoulds

Density and site match
- List density matches the mock: 3-line rows, about 90 pt, icon 44 with quality ring, name in quality colour, order as published. Yours tag, chips, set line, `Updated` stamp all match. Good.
- iOS live and Android live differ only in system chrome. Parity holds on text, spacing and colours.
- The site's gear row prints a source-kind icon before the source line (`SourceKindIcon`; rep gets the faction mark). The app prints text only. Add the kind icon at 16 or accept the gap in writing.
- The iOS live capture has no row with `You wear it` and none with `Gain not modelled yet`; the offline render shows `You wear it` on Android only. Capture each tag once per platform.

Picker
- Class grid, spec buttons (48 tall, icon, name, role), six range pills, two faction pills, gold selected state: all match the mock. Android live picker is the reference; iOS render lacks spec icons because the render has no icon source (fine, but capture it live per blocker 1).
- `Back to {name}'s list` is uppercase on iOS (render-picker-default-ios). Spec and mock: sentence case. Use sentence case.
- The picker sheet is fine as a large detent; the list redraw behind it was not captured. One capture of the list changing under the open sheet.

Offline
- Strip copy and placement match the mock (`Offline. Showing what you had at Thu 4:41 pm.`). Offline exists only as a render; the report admits a fixture build cannot reach it. Accept for this pass, but one device capture with airplane mode is owed before the wave closes.
- Not-cached state (render-list-not-cached-default-ios) is verbatim but a dead end: one plain sentence at body weight, 1400 pt of empty screen. Add `Try again` (secondary button, same as the error state) and use secondary text colour with a 24 pt top margin so it reads as a state, not a stray paragraph.

Large text
- iOS AX5: the title wraps to two lines and the `YOURS` tag centres between them. Align the tag to the first baseline. Chips stack one per row (correct, none cut). The first list row starts below the fold at AX5 (title, three chips, four-line set line); shorten the set line to the figure plus `DPS at best in slot` and move `Updated` to its own meta line only if it still fits. Acceptable at AX5 as is.
- Android 2.0: list rows and chips reflow cleanly, nothing is cut. The tab labels break mid-word; known, F2.

Departures from the spec
- **Unit words (HPS for healers, points for tanks): accept.** Printing DPS on a healer or tank band is wrong, and the site already does this for healers (`unit` prop, `HPS`). Amend the spec. Tanks: confirm the word `points` against the site's tank band wording; the site's `RateUnit` today has no tank value, so the two surfaces must agree before launch.
- **Raw source strings (`Stratholme: Balzaphon`, `Warsong Outriders`): accept.** The site prints `sourceDetail` as published. The app now matches the site; the mock's dot form and `(Exalted)` should be dropped from the mock and spec. Do not re-derive.
- **`Ring` for both rings: fix.** Two rows reading `Ring` cannot be told apart, and a player cannot tell which ring slot the pick is for. The character frame says Finger 1 and Finger 2; the site's list word is `Ring`. Use `Ring 1` and `Ring 2` (and `Trinket 1`, `Trinket 2` if the same collapse applies). Take the numbering from row order.

Known gaps (tooltip, selector, Android tab labels) are not counted.

## Checklist
- [x] 17 rows in character-frame order, quality-coloured names, 44 icons
- [x] Three chips wrap, none cut (both platforms, both text sizes)
- [x] Yours tag; set line with figure in mono
- [x] Next line, `same DPS`, unit word
- [ ] Source-kind icon as on the site
- [x] Picker content and states on Android
- [ ] Picker on iOS captured live, clean pills (blocker 1)
- [ ] Picker at large text reaches Faction and Back (blocker 2)
- [x] Offline strip copy and placement (render only)
- [ ] Not-cached has a way out
- [x] Error and loading states rendered both platforms
- [ ] Ring and trinket rows distinguishable
