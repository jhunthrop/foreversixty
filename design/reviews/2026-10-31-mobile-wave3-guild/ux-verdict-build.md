# F16 Guild extras, build review on captures (2026-10-31)

Verdict: HOLD. Spec copy is exact on both platforms and Android is close to the board. iOS has one layout break at largest text, one accuracy contradiction, and states that cannot be judged.

Captures live in design/reviews/2026-10-31-mobile-wave3-guild/build/{ios,android}/.

## MUST (blocking)

1. iOS, `f16.3-history-page1-xl.png`: on every row with the `You were there` tag the 48 pt date block is pushed off the left edge (clipped to "U/1/C"). The Tue 29 row, with no tag, lays out correctly. Meta also breaks to one fragment per line, five lines. Fix: the date block keeps its width and sits above the text at AX sizes, as Android does in `android/f16.3-history-page1-xl.png`; keep the meta on `{a} · {b} · {c}` lines.
2. iOS, `f16.4-calendar-on-xl.png`: the Which raids and Alert rows are indented against the switch row. The chevrons collapse to stray dots. `1 hour before invites` wraps under its chevron. The link icon crosses the card edge. Reset link is shorter than Copy link. Fix: stack label above value at AX sizes, keep a 44 pt row and a 20 pt chevron.
3. iOS, `f16.4-calendar-on.png` and `f16.5-access-off.png`: the switch draws with no knob. On is a gold pill and off a dark pill, so state is colour only. `f16.4-reset-alert.png` draws the knob, so this is likely a snapshot artefact. Re-capture with the knob visible.
4. iOS, `f16.7-raids-tab.png`, and Android, `f16.7-raids-tab.png`: the guild head says `24 on the roster`, the Roster row directly under it says `38 on the roster`. One screen contradicts itself. Both must come from the API `counts`.
5. iOS, every roster capture and `f16.7-raids-tab.png`: spec icons are empty 32 pt tiles. Spec: a class crest when the spec is unknown, never a blank tile (tenet: no missing-icon flash). Android shows real icons. Fix the fallback to the crest, and re-capture with icons loaded from a fixture.
6. iOS, `f16.1-roster-all.png`, `f16.3-history-page1.png`, `f16.4-calendar-on.png`: the pushed-screen titles use the system large title. v1 §4 sets pushed titles to Barlow Condensed 800 at 34, caps, as on the board and on Android. Apply it to Roster, Raid history and Calendar.
7. Raids tab, both platforms, plus history captures: the fixture makes history run to 31 Dec while Next raid is Sat 19 Dec and the latest raid night is Mon 14 Dec. History cannot be later than "next". Re-capture with coherent dates.

## SHOULD

1. iOS roster has no trailing search button (`f16.1-roster-all.png`); the search field is shown in `f16.1-roster-search.png` with no visible entry point. Spec: a trailing search button, as on Android.
2. iOS roster and search captures show thin vertical lines at the left and right edges of the chips and the search field (`f16.1-roster-search.png`). Chip focus ring is clipped by the scroller.
3. Order inside a rank differs. iOS is Kraggor, Draknar, Mendara, Thokgar (name). Android is Kraggor, Mendara, Thokgar, Draknar. Pick one and state it in the spec.
4. Mono durations render `1  h  48  m` with wide gaps on both platforms (`android/f16.3-history-page1.png`). Board sets them tight.
5. `NO ACCOUNT` is a tag the spec does not list. `Waiting` renders correctly on Android. Add it to the copy table or drop it.
6. Android, `f16.1-roster-all-xl.png`: the Sort chip is clipped at 2.0 and only scrolls. Wrap the chips as iOS does, but give iOS XL chips 8 pt row spacing (`ios/f16.1-roster-all-xl.png`: rows touch).
7. iOS XL roster: the spec icon stays 32 pt, pinned to the row top. Centre it on the name line.
8. Guild head glyph differs: red shield on iOS, red star on Android. The board shows a Horde crest. One crest language.
9. Android reset alert has a heavy condensed title and 30 pt empty space under the buttons (`android/f16.4-reset-alert.png`). Use the body type size.
10. Missing acceptance captures: iOS DPS chip, sort menu, sort Name, Class, Level; Android F16.6 (known); a 40-member scroll on both; any iOS capture above the fold with the tab bar, so the answer visible without scrolling cannot be judged.

## Ticked against the spec
- [x] Counts card, role chips, rows, attendance bar and `{n} of 8` (Android exact; iOS ok).
- [x] `Member` as the plain rank word.
- [x] No match, error, empty and footer copy verbatim.
- [x] `You were there` tag under the meta on one line.
- [x] Access-off card, switch off, `Open Settings`.
- [x] iOS `Add to Calendar`, Android has none, as specified.
- [x] Event title, location, notes and alert match (`ios/f16.6-event.png`).
- [x] Reset alert copy and buttons.
- [ ] Largest text, iOS F16.3 and F16.4 (MUST 1, 2).
- [ ] Icons, switch states, header counts (MUST 3 to 5).
