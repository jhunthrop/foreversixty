# Guild raids and sign-ups (Lane G1): ux-designer review

Reviewer: ux-designer, 2026-10-26. Spec: `design/specs/2026-10-26-guild-raids.md` §3–§7, §9. Evidence: the PNGs and JSON in this folder only.
Fixture names, dates and the `Sign in` header are not defects. The fix round already in flight covers these items: the ` · ` spaces from `md`, the notes bullets, the gold checkbox, the 16 px `Settings` hit area, the 77 → ≤ 72 planned row, the 38 → 36 roster row, and the loading reserves. Each is judged "fix in flight" below unless the PNG shows a further problem.

## Region by region

### Raids tab (01–04, 12-tab, 15, 18-tab)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| Header row, `Plan a raid` | One fill, right; read-only `disabled` with the gate sentence under it | Matches (01, 04-360). The gate sentence and link wrap cleanly at 360. | ok |
| Form open | §5: `Post the raid` is the only fill | `Plan a raid` stays gold above the open form (the 13 form panels, top edge), so the screen shows two fills. | must fix |
| Planned row, desktop | 64–72 px; count `23/40` with its 96 px bar; mock sets the bar on the count's line | The bar sits on its own line under `14/40`, giving three lines and 77 px (15-focus-planned-row-1440). | fix in flight; see ruling (c) for the exact change |
| Planned row, phone line 3 | Count and breakdown never collapse (§7) | `1 tank · 3 healers · 9 D…` is truncated at 360 (01-panel-360, 04-panel-360). | must fix |
| Cancelled row | Title and count `--muted`; chip `Cancelled` | Title and count are muted. The countdown is still `--gold` (`in 5 days 7 h`). The viewer chip shows beside `Cancelled`: `IN · TANK` in `--kill`, or the gold `NOT SIGNED UP` (12-tab-ready-1440, where the two chips stack and add height). At 360 the title truncates to `Hyjal Su…` (01-panel-360). | must fix; see ruling (d) |
| Viewer chip, countdown | One gold mark per row | `Not signed up` rows have a muted countdown (12-tab-ready). Correct. | ok |
| Hover, focus-visible | `--card-top` ground with title underline; 2 px gold at −2 px | Matches (15-hover, 15-focus-planned-row). | ok |
| Empty Planned | Member and officer lines verbatim | Verbatim (03). `PLANNED 0` shows a zero count. | polish: drop the count at 0 |

### Raid page band (05–11, 18-page)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| Back link, eyebrow, h1, chip | §3.2 lines 1–3 | Matches at every width. The officer row sits right on the eyebrow line from `md` and wraps to two rows at 390 (09-panel-390). | ok |
| When, count lines | ` · ` separated | Separators lose their spaces from `md` (05/08/09 at 1440 and 2000). | fix in flight |
| Countdown colour, cancelled/ended/done | Gold appears once, for a raid you can still act on (§5) | `in 2 days 7 h` stays gold on a cancelled raid, and the bar stays gold too (08-cancelled-panel-1440). `ended 1 h ago` is gold (08-ended, 08-done). | must fix |
| Done line | `Logged as a raid night.` + `Open the night` | Present (08-done-390). | ok |
| Cancel confirm | Replaces the officer row: question, checkbox, `Cancel the raid`, `Keep it` | Matches (09-confirm-390). The checkbox is native. | fix in flight (checkbox) |
| Copy failed | Line + selected 6-row textarea | Matches (16-panel-390). | ok |
| Not visible | Line + `SignInPrompt`; unverified sentence | Matches (11-visitor-panel-390). No raid data leaks. | ok |

### Your sign-up strip (05, 06, 07, 08, the 15 focus captures)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| Rows, desktop | 68 px grid row; `minmax(180px,1fr) auto 152px minmax(200px,1.4fr)` | 68 px measured. The columns match. At 1024 the note field is 200 px and its text clips inside the input (18-page-1024). That is acceptable. | ok |
| Rows, phone | ≤ 196 px; the role select on line 1, segments line 2, note line 3 | 183 px with the status line (05-360). The segments are 81 px each at 360. `Bench` fits. | ok |
| Segment focus-visible | 2 px gold at −2 px | Matches (15-focus-segment-390). | ok |
| Disabled (locked, ended, done) | Native `disabled`, described by the state line | The state lines show (07, 08). The segments, role select and note look identical to the live state, so nothing reads as disabled (08-ended-panel-1440, 07-lock-race-panel-390, 08-done-390). | must fix |
| Lock race line | `--wipe`, under the row | Matches (07-lock-race-panel-390). | ok |
| Cancelled | Strip replaced by `This raid was cancelled.` | Matches (08-cancelled). | ok |

### Notes and roster (05, 09-roster, 18-page)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| Notes list | Bullets (mock) | No bullets. | fix in flight |
| Role columns 1280+ | 4 columns, DPS across 3–4 | Matches at 1440 and 2000. The page stays in the centred 1344 column at 2000. | ok |
| Row | Crest 24, name, spec, note; 36 px | The spec sits inline after the name. It falls under the name only when the name is long (`Duskstrider / Marksmanship`, `Shadowquill / Assassination`, `Arcanemoor Arcane · / DPS`). Rows in one column therefore run 38 or 56 px, and the columns stop aligning (09-2000, 18-page-1440). | must fix; see ruling (c) |
| Late / Bench / Out | 24 px under the columns | Stacked one under another, each confined to column 1's ~300 px. The other three quarters stay empty at 1440 and 2000. | must fix; see ruling (a) |
| Group sub-heads | 15 px sub-heads allowed; mock draws `label` 11 px | Seven 15 px Cinzel heads compete with the h2 `Sign-ups`. | polish: `label` 11 px muted, count in mono, as `PLANNED 3` |
| Not signed up | Names for members, rows for officers | Members get one wrapped line of names (05-360). Officers get rows across the role grid's columns (09-2000). Both are good. | ok |
| Own-row wash, officer selects | `rgba(229,185,85,.06)`; 32 / 44 px, 96 px wide | Matches (09-roster-390, 09-2000). | ok |

### Form (13, the 15 form focus captures)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| Layout, defaults | Two columns from 1024; last zone, weekday, `20:00`, `3 h`, guild zone | Matches. The form is 726 px against the 760 cap and 989 px against 1180. | ok |
| Selects | `FilterSelect` chrome on every select | `Raid`, `Duration` and `Guild time zone` draw the browser's native arrow (13-form-default-panel-1440). The strip's role select has the shared chevron, so the two sets of selects look different. | must fix |
| `For you` | `For you: Thu 17 Dec, 9:00 pm.` | Verbatim (13-form-for-you-panel-1440). | ok |
| Errors | `--wipe` line under the field, `aria-invalid` | `Pick a date.` under Date (13-form-errors-390). | ok |
| Repeat at the cap | Helper `Only 2 more fit…` | The helper shows (13-cap-360). The box can still be ticked, though the API would refuse the save (422, nothing created). | must fix |
| No webhook | The line with the inline `Settings` link | Shows. The link is 16 px tall. | fix in flight |
| Checkboxes | Gold site checkbox | Native box. | fix in flight |

### Band line (14, 12-tab)
| Element | Spec | Shipped | Severity |
|---|---|---|---|
| In, unsigned | Sentence, segments (36 px from `md`, 44 phone), gold zone link | Matches. The line is 44 px at 1440 and 98 px at 360/390. The phone sentence takes two lines with the link beside it (14-band-in-panel-360, 14-band-unsigned-panel-390). | ok |
| Two own characters | The segments control the current character | 12-tab-ready reads `you are In on Duskstrider`, yet its segments show nothing checked, because they are bound to Zulmara, who is unsigned. A press of `In` there moves Duskstrider to Bench (R3) with no warning. | must fix |
| None planned | Member line; officer line + `Plan a raid` link | Verbatim (14-band-none-officer-panel-1440). | ok |
| Needs line | `before Thursday` / `before the next raid` | Matches. | ok |

## Rulings

### (a) Late, Bench, Out: three columns, as the mock
From 1024 the three groups form one grid: `grid-template-columns: repeat(3, minmax(0,1fr))`, with the role grid's column gap and `margin-top: 24px` under the role columns. Each group keeps its sub-head and the standard roster row. At 768–1023 the grid is `repeat(2, minmax(0,1fr))`, with Late and Bench in row 1 and Out in row 2, column 1. Below 768 every group stacks, in that order. An empty group is left out, and the rest flow left; no blank column remains. `Not signed up` stays full width under this grid, 24 px below it.

### (b) CLS: the cap stays at 0.02, measured from first paint
The band is this lane's surface, so its settling counts. Each reserve must equal the ready render. No reserve may be a guessed constant.
- Raid page band: 347 px at 390 and 311 px at 1440 for a member. The ready band runs 72→419 and 109→420. The spec's 300 / 232 is replaced. Draw the back link, the eyebrow (without its chip) and `GuildTabs` for real while loading, because they need only the guild. Skeleton only lines 3–5, at their ready line boxes: h1, the when line (three lines on phone, one on desktop) and the count line. When the cached role is officer, the reserve also holds the officer row: one row on desktop, two wrapped rows on phone.
- Strip: the panel padding plus the sub-head plus n × 183 px on phone, or n × 101 px on desktop (the row plus its 20 px status line, as measured). The spec's 196 and 68 are caps, not reserves.
- Raids tab: draw the header row for real. The planned reserve is 49 + n × 72 px on desktop and 60 + n × (the phone row after the line-4 fix below, about 115) px. n comes from `fs.guildPlanned.<guildId>`. This replaces 24 + n × 64 / 108. The loading capture puts the `Raid nights` skeleton at y 773 against 849 ready (1440), and at 920 against 934 (390).
- Guild band, any tab: when the viewer was last seen verified, the band's loading min height grows by the raid line plus its gap (44 + 12 desktop, 98 + 12 phone), because §3.4 added that line.

The implementer re-prints `cls-390.json` and `cls-1440.json` and attaches loading beside ready, matching to ±1 px.

### (c) Caps confirmed: planned row 64–72 desktop, roster row 36 desktop
- Planned row: set the bar on the count's line (`14/40` then the 96 px bar, gap 8), with the breakdown under it, as the mock draws. The count cell then has two lines and the row lands at ~66 px.
- Roster row: the name (14 px, line 18) sits over the spec (12 px muted, line 16) at every width, both one line with an ellipsis, and the full text goes in `title`. Late and Bench rows read `Arcane · DPS` on that second line. Total 34 px, min 36. Every row in a column is then the same height until a note adds lines.
- Phone planned row: the breakdown moves to its own line 4 (12 px muted, 16 px line). The phone cap of 96–120 holds at ~115.

### (d) Chip order: the status chip first, then the viewer chip; a cancelled row shows `Cancelled` alone
- On a cancelled row your sign-up no longer matters. `IN · TANK` in `--kill` or a gold `NOT SIGNED UP` beside `Cancelled` gives the row two marks and makes it taller. The countdown on that row turns `--muted`, as the mock draws.
- Locked and ended rows show both chips on desktop, status first.
- On phone, line 1 holds one chip, chosen in this order: `Cancelled`, `Locked`, `Ended`, then the viewer chip. This lets `Hyjal Summit` keep its title.

### (e) At 360
The strip and the band line read correctly; their heights are 183 and 98 px. Two defects at 360 are already listed above: the band segments when the sentence names another character, and the planned row's truncated breakdown.

## Checklist (§9)
- [x] Rows 1–16 at their widths; row 18 at 1024, 1440 and 2000. Row 17 belongs to G1b.
- [x] Copy verbatim on every captured string. One exception: the ` · ` spacing, fix in flight.
- [x] 2000 px: the 1344 column is centred and the band is full-bleed.
- [x] Focus-visible on row, segment, role, note, officer button and form controls; hover at 1440.
- [x] Not-visible states leak nothing (R1).
- [ ] Heights within caps: planned row 77 and roster row 38 (fix in flight, ruling c).
- [ ] CLS ≤ 0.02 at 390 and 1440 on both surfaces (ruling b).
- [ ] One gold per row or region (cancelled and ended countdowns, the open form).
- [ ] Disabled controls read as disabled.
- [ ] Every select in `FilterSelect` chrome.

## Verdict: SHIP WITH FIXES
Each fix below must land with re-captures attached.
1. Disabled look (08-ended-panel-1440, 07-lock-race-panel-390): the segment group, role select and note take the shared disabled recipe, `opacity-50` and `cursor-not-allowed`, whenever they are `disabled`.
2. Gold on past and cancelled raids (08-cancelled-panel-1440, 08-ended, 08-done-390, 15-focus-planned-row-1440): the countdown is `--muted` when the raid is cancelled or its time reads `started…` or `ended…`. The page's bar fill is `--muted` when the raid is cancelled.
3. Cancelled planned row (12-tab-ready-1440, 01-panel-360): show `Cancelled` alone and make the countdown muted. On phone, follow the chip priority in ruling (d).
4. Planned row breakdown at 360 (01 and 04 panels at 360): put the breakdown on line 4, per ruling (c).
5. Roster rows (09-2000, 18-page-1440): name over spec at a fixed 34/36 px, per ruling (c).
6. Late, Bench, Out grid (18-page-1440, 09-2000): per ruling (a).
7. Reserves (the 12 captures): the heights in ruling (b); re-print the CLS.
8. Form selects (13-form-default-panel-1440): `FilterSelect` chrome on `Raid`, `Duration` and `Guild time zone`.
9. Repeat at the cap (13-form-cap-no-webhook-360): with fewer than four slots, the repeat box is `disabled`, with `aria-describedby` pointing at the `Only 2 more fit…` helper.
10. Two fills with the form open (the 13 form panels): while the form is open, `Plan a raid` is hidden. On close it returns and takes focus, per §3.1.
11. Band segments with two characters (12-tab-ready at 1440 and 390): the segments bind to the character the sentence names when that character is In or Late. Otherwise they bind to the header's current character. Their `aria-label` names the bound character.
12. Polish (03, 09): drop the count from `PLANNED 0`; set the roster group heads in `label` 11 px.
