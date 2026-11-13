# Addon source filter, 2026-11-13

Status: written by ux-designer before the build. Owner, in game today: "a lot of the leveling BiS are expensive world drops that most players won't be able to buy. We should have a toggle." The copy rule of `2026-11-08-addon-ux.md` (owner, 2026-11-08) binds every string: a number gets its unit and nothing else, no explanation the context already gives, game words only. The surface passes ai-slop-reviewer, wow-player, info-density-reviewer and brand-consistency-reviewer against the web BiS page. Boards and verdicts: `design/reviews/2026-11-13-addon-sources/` (boards from `design/mocks/addon_sources.py`).

## 1. The question and the answer

**The question:** which of these upgrades can I actually get?

**The answer, visible on opening the panel:** the three Upgrades rows already re-picked from the sources I can use, and one line under them that says which sources those are. Nothing scrolls: the foot line sits directly under the rows, inside the panel's existing height.

**Reference:** the web BiS Sources group (`BisSources.svelte`, `source-choice.ts`, `source-filter.ts`, `page-state.ts` defaultSources). The addon copies its eight kinds, their order, its labels, its default and its field wording exactly. From the addon, the Talents row's `<build name> · Change` pattern (gold text button) and the Settings page's switch rows.

## 2. Vocabulary (identical to the web)

Eight kinds in this order: Dungeon, Raid, World / BoE, Quest, Crafted, Reputation, Vendor, PvP.

- Default with no saved choice, below level 60: every kind on except Reputation, and standings up to Friendly still count. At 60: every kind on.
- Filtering re-picks, never hides. A slot with a kind off shows the best pick from the kinds still on.
- The state words, from `sourcesFieldLabel`: `All sources`, `No rep above Friendly` (the leveling default only), `Pre-raid` (the set the web names so), `Dungeon only` (one kind), `6 of 8 sources`, `No sources` (none).
- The state is one module, `Sources.lua`: `Sources.state(band)` returns the kinds on and `Sources.words` the words; `Gear.pick` resolves every pick through `Sources.resolve`. Panel, flyout, Settings and tooltips read it. No surface restates the rule. Another player's gear (group panel, loot council, the hello's gap) is judged on every source (`Sources.ALL`).
- Choice: `ForeverSixtyDB.ui.gearSources.byCharacter[key].on`, a list of kind ids, or no entry. No entry means the site's default for the band looked at. A saved list is literal and never rewritten.

## 3. The panel foot line

Region: directly under the Upgrades rows and their rest line, above the Talents row. One line, 12 px body font, row height 18, no divider.

| Part | Style |
|---|---|
| State words | muted |
| ` · ` | muted |
| `Change` | gold text button, the Talents row's button (`button(content, L.panelChange, ...)`) |

Always shown, also at `All sources`. The state is part of the answer, and the player at 60 must be able to find the switch. Text: `All sources · Change`, `No rep above Friendly · Change`, `Dungeon only · Change`, `6 of 8 sources · Change`.

Hover on the state words: a tooltip titled `Sources`, one muted line per kind that is off, in the web's order, for example `Off: Reputation, PvP`. At the leveling default the line is `Off: Reputation above Friendly`. At `All sources` there is no tooltip.

Elements: `Change` has default (gold), hover (gold, underlined), focus-visible (1 px gold ring), active (dimmed gold), disabled in combat (muted, no hover). Keyboard: Tab reaches it, Enter and Space open the flyout, as for the Talents row.

## 4. The Sources flyout

Opens beside the panel on `Change`: attached to the panel's right edge with a 4 px gap, or its left edge when the panel is within 260 px of the screen's right edge. 232 px wide, the panel's backdrop, border and padding. It is built on first open, never at login.

Regions, top to bottom:

1. Title row: `Sources` in the panel heading style, a close X at the right.
2. Eight rows, 22 px high, the Settings checkbox at the left, the kind's label, the count right-aligned in muted tabular figures. Order and labels as in section 2. The count is gear slots with an option from that source at the character's band, independent of what is on.
3. A text button `Default` at the left of the foot, shown only when a choice is saved; muted, hover gold. It clears the saved choice.
4. Foot line, muted: `Counts are gear slots with an option from that source.`

Row states: default, hover (row background lifts one step, tooltip below), focus-visible (1 px gold ring on the row), active (checkbox pressed), disabled (a kind with a count of 0 is still toggleable, its count reads `0` muted; nothing else disables). At the leveling default the Reputation row is unchecked and its hover tooltip reads `Standings up to Friendly still count.`

What happens at once: each click saves the list, then every surface re-picks in the same frame: Upgrades rows, foot line, slot marks, chip line, open tooltips, the Builds dialog's gear rows. There is no Apply. Unchecking the last kind is allowed and reads `No sources`.

How it closes: its X, Escape, `Change` again, opening the Builds dialog or the tour, and combat start (hidden, not reopened afterwards). Clicking elsewhere does not close it.

Loading: none, the counts ship with the data. Error: if the data has no counts for the band, the count column shows `-` in muted and the toggles still work. Empty: not reachable.

## 5. Settings page

Decision: eight toggles under a heading, not one row that opens the flyout. The Settings canvas has no panel to anchor a flyout to, and a row that opens a floating window from a page is a second control to learn. The group reuses the page's switch row style (label, hint under it) and is generated from the same kind list as the flyout, so the two cannot drift.

- Heading `Gear sources`, in the page's existing group heading style, placed after the tooltip switches. At its right, the `Default` text button, same rule as the flyout.
- One hint line under the heading, muted: `Upgrades are picked from the sources you leave on.`
- Eight switch rows, labels as section 2, no per-row hint, counts at the right as in the flyout.
- A change here updates the open panel at once, and the reverse.

`SettingsPage.SWITCHES` stays booleans of `Prefs`. The group is a second list, `SettingsPage.SOURCE_ROWS`, built from `Gear.sourceKinds()`; `layout` measures it like the other rows.

## 6. Item tooltip

Plain hover: no new line. The `Next` line and the muted source line name the item the filter picked. A rep or vendor item that needs a standing ends its source line with it: `Silverwing Sentinels · Friendly`.

Behind Shift, one muted line, shown only when the slot's pick differs from the pick with every source on, placed directly under the source line:

`L.tipSetAside = "Best with all sources: %s, %s DPS more"`

First `%s` the item name in its quality colour, second the gap to one decimal, priced in game from the two items' stats as every other gain in the addon (Data.lua does not carry the sim's `dps_delta`); with no price, `Best with all sources: %s` alone. For example `Best with all sources: Hawkeye's Cloak, 0.3 DPS more`. The line never shows at `All sources`, and never when the unfiltered pick is the same item.

On an item that is itself the unfiltered best but is not in the chosen sources (a world drop in a bag), the verdict is unchanged and the Shift line is not shown.

## 7. Upgrades row detail

Detail stays `<slot> · <source>`. Add nothing for a re-picked slot: the gap is behind Shift in the tooltip, and a row that carries it is a fourth fact on a 3-fact row. For a reputation or vendor item that needs a standing, the source ends with it: `Back · Silverwing Sentinels · Friendly`. The standing is a gate, so it shows always, not only when re-picked.

## 8. States of the whole feature

| State | Foot line | Rows and tooltips |
|---|---|---|
| All on (level 60 default, or chosen) | `All sources · Change` | Unfiltered picks. No Shift set-aside line. |
| Leveling default | `No rep above Friendly · Change` | Picks exclude standings above Friendly. |
| One kind | `Dungeon only · Change` | Picks from dungeon drops only. |
| Mixed | `6 of 8 sources · Change` | As chosen. |
| None on | `No sources · Change` | Every slot empty (below). |
| A slot has nothing from the kinds on | unchanged | The slot has no Upgrades row, no mark, and its empty-slot hover reads `No %s from these sources at this level` (the web's wording; `%s` the slot name in lower case). |
| Every slot empty | unchanged | The Upgrades region reads `No upgrades from these sources` in muted body text where the rows were; the rest line is hidden. |
| Level 60 crossing, no choice saved | becomes `All sources · Change` | Moves by itself on the level-up, no message, no toast. |
| Level 60 crossing, choice saved | unchanged | The saved list stays literal; `Default` clears it. |
| Combat | `Change` disabled | Flyout hidden. |

## 9. Locale keys

| Key | Text |
|---|---|
| `sourcesTitle` | `Sources` |
| `sourcesChange` | reuse `panelChange` (`Change`) |
| `sourcesDefault` | `Default` |
| `sourcesAll` | `All sources` |
| `sourcesLevelingDefault` | `No rep above Friendly` |
| `sourcesPreRaid` | `Pre-raid` |
| `sourcesNone` | `No sources` |
| `sourcesOne` | `%s only` |
| `sourcesSome` | `%d of %d sources` |
| `sourcesOffTip` | `Off: %s` |
| `sourcesOffLevelingTip` | `Off: Reputation above Friendly` |
| `sourcesRepTip` | `Standings up to Friendly still count.` |
| `sourcesFlyoutFoot` | `Counts are gear slots with an option from that source.` |
| `sourcesKindDungeon` ... `sourcesKindPvp` | `Dungeon`, `Raid`, `World / BoE`, `Quest`, `Crafted`, `Reputation`, `Vendor`, `PvP` |
| `sourcesSettingsHeading` | `Gear sources` |
| `sourcesSettingsHint` | `Upgrades are picked from the sources you leave on.` |
| `tipSetAside` | `Best with all sources: %s, %s DPS more` |
| `tipNoSlotFromSources` | `No %s from these sources at this level` |
| `panelNoUpgradesFromSources` | `No upgrades from these sources` |

Source detail with a standing: the place, ` · `, the standing as `Friendly`, `Honored`, `Revered`, `Exalted`.

## 10. Limits

- No layout shift: the foot line's 18 px is reserved in `Panel.upgrades` layout in every state, so toggling a kind never moves the Talents row.
- The flyout builds on first open; no frames or data work at login.
- A toggle re-picks in one frame. `Gear.pick` caches per slot, band and state; the state change clears the cache once.
- No new textures: the checkbox, X and backdrop are the existing ones.

## 11. Boards the generator (`design/mocks/addon_ux.py`) must render

- `panel-sources-leveling` : panel, level 24, `No rep above Friendly · Change`, rows re-picked
- `panel-sources-all` : panel, level 60, `All sources · Change`
- `panel-sources-empty` : panel, `No sources`, empty Upgrades region
- `flyout-sources-default` : flyout open beside the panel, leveling default, `Default` hidden
- `flyout-sources-chosen` : flyout, Dungeon only, `Default` shown, one row focused
- `panel-sources-state-tooltip` : the hover on the state words
- `settings-gear-sources` : the Settings page with the group
- `tooltip-set-aside-shift` : item tooltip behind Shift with the set-aside line
- `tooltip-rep-standing` : item tooltip with `Silverwing Sentinels · Friendly`

Acceptance: each board at 1x, with the panel at its real width, and the flyout board next to the panel edge at the screen's right side (flyout on the left).

## 12. Data (generator, 2026-11-13)

Data.lua carries the filter's facts per pick after its alternatives, one empty field, then one packed tail (`pipeline.addonlua._bis_tail`): a head `*!RhVh,19513` (two-hander mark, the pick's own kinds with the standing each asks, the ids its pair mate must not hold) and the best item of every other kind in rank order, `Dh!26686,12,1234+5678` (kind letter, standing letter, two-hander mark, id, an index into `bisSources`, the not-with ids on a pair's second slot only). The sim's `dps_delta` is not carried. addon-data.json does not mirror the tail. Measured 2026-11-13 on build 1.60.1.70790: Data.lua 1,290,915 to 1,541,776 bytes; Lua 5.1 heap 3,260 to 3,670 KB of the 4,096 KB ceiling.

Not in this pass: the Builds dialog's gear rows (owned by the Builds lane) and closing the flyout when the tour opens (owned by the tour lane) still show and behave as before.
