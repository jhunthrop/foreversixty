# Selected, current and own marks: inventory (2026-11-09)

Owner finding: on /bis the viewer's own spec row wore a gold bar and a gold "Yours" pill, while the nav, the tabs and the select rows mark selection in the round-5 accent (`#0093b2`). The basis is design/specs/2026-11-09-directions.md (section 2: "blue, active tab `#0093b2`"), which supersedes the 2026-11-05 web layout spec's "gold means the selected state" (P2, the segmented control and paperdoll rows).

## The one treatment (tokens, design/tokens/tokens.json)

| Token | Web value | Use |
|---|---|---|
| `--color-selected` | `rgba(0, 147, 178, 0.16)` (was white at 7 %) | The ground under anything selected, current or own |
| `--color-selected-mark` | `var(--color-accent)`, `#0093b2` | The mark: a chip's edge, a slot's ring, a dot, a rail |
| `--shadow-selected-bar` | `inset 3px 0 0 #0093b2` | Rows, cells, list items |
| `--shadow-selected-underline` | `inset 0 -3px 0 #0093b2` | Tabs and segments |
| `Tag tone="yours"` | accent edge, the tint, white ink | "Yours", "Your spec", "You" |

Unchanged on purpose: keyboard focus is the 2 px gold ring (global.css `:focus-visible`; the directions spec says nothing new about focus), so focus never reads as selected. Hover is `--color-hover` with no mark. Gold stays for links, the answer number and the text accents. The apps keep their gold `selected` value until the mobile lane adopts the palette.

## Inventory

"Before" is the colour on main at 52748165d; "After" is this lane.

| # | Mark | Where | Before | After |
|---|---|---|---|---|
| 1 | Own spec row (BiS landing, guide class and spec pickers) | components/guides/GuideSpecRow.astro | white 7 % ground, 3 px **gold** bar | accent tint, 3 px accent bar |
| 2 | "Yours" / "Your spec" pill on that row | GuideSpecRow.astro, ui/Tag.svelte | Tag tone gold: **gold** ink and edge on gold 10 % | Tag tone yours: accent edge, accent tint, white ink |
| 3 | "You" pill on the rankings | components/Rankings.svelte | Tag tone gold | Tag tone yours |
| 4 | `row-band-own` (guild roster own row) | ui/row.css, lib/ui/row.ts | white 5 % ground, no mark | accent tint, accent bar |
| 5 | `row-band-mark`, `.row-band-target:target` (marked or addressed fight) | ui/row.css | white 7 % ground, 2 px accent bar | accent tint, 3 px accent bar |
| 6 | Rank row `aria-current` (tier list own spec, rankings) | ui/rank-list.css | white 7 % ground, **gold** rail slice | accent tint, accent rail |
| 7 | Compact rank row `aria-current` (tier list beside its rail) | ui/rank-list.css | accent 16 %, accent rail (own rule) | the shared rule (duplicate deleted) |
| 8 | Paperdoll grid cell `aria-selected` | ui/Paperdoll.svelte | white 7 %, 3 px **gold** bar | accent tint, accent bar |
| 9 | Inspector expanded in place under the selected cell | ui/Inspector.svelte | white 7 %, 3 px **gold** bar | accent tint, accent bar |
| 10 | Paperdoll frame selected slot (BiS spec, character) | ui/PaperdollFrame.svelte | 2 px **cream** ring (same as focus) | 3 px accent ring with an accent glow; focus keeps the cream ring |
| 11 | Paperdoll phone list selected item | ui/PaperdollFrame.svelte | 2 px **cream** outline | 2 px accent outline |
| 12 | Segmented control `aria-checked` (Buffs Raid/None/Custom, level range, faction) | ui/segmented.css | white text, 3 px accent underline | unchanged look, now the token |
| 13 | Stacked segmented `aria-checked` (phone) | ui/segmented.css | 2 px **gold** bar | accent bar |
| 14 | Tab strip current (`aria-current="page"`, `aria-selected`) | lib/ui/tabs.ts | white text, accent underline | unchanged look, now the token |
| 15 | Tab strip overflow menu current | ui/TabStripMore.svelte | white, 3 px **gold** bar | accent bar |
| 16 | Talent editor tree tabs `aria-selected` | ui/TalentEditor.svelte | **gold** text, 2 px **gold** underline, white 7 % | white text, accent underline, accent tint |
| 17 | Contents rail current section (`aria-current="location"`) | ui/Toc.svelte | **gold** text, 2 px **gold** bar | white text, 2 px accent bar (the rail is 2 px) |
| 18 | Select option chosen (single select) | ui/select.css | gold check only; hover and keyboard-active rows wore the accent tint and bar, so moving looked chosen | chosen row: accent tint, bar and check; hover and active: the hover ground |
| 19 | Multi-select box checked | ui/select.css | **gold** box, dark tick | accent box, white tick |
| 20 | `sel-you` ("You" inside a select row) | ui/select.css | **gold** ink, gold 38 % edge | the yours pill |
| 21 | `chip-preset` pressed (chart players presets) | ui/select.css | **gold** edge, gold 10 %, gold ink | accent edge, accent tint, white ink |
| 22 | Menu item keyboard focus | ui/select.css | **gold** 10 % wash and gold bar (read as selected) | the gold focus ring inside, hover ground |
| 23 | BiS sources chip pressed | components/bis/BisSources.svelte | **gold** edge and ink | accent edge, accent tint, white ink |
| 24 | Guild filter and sort chips pressed (roster sort, readiness Next raid) | lib/guild/filter-styles.ts | **gold** edge, ink, gold 8 % | accent edge, accent tint, white ink |
| 25 | Button `pressed` toggles (a raid sign-up's I'm in / Late / Can't come) | lib/ui/button.ts | white edge and ring | accent edge and ring |
| 26 | Item row `aria-pressed` | ui/ItemRow.svelte | accent tint, 3 px accent bar | the tokens |
| 27 | Boss picker tab (character parse column) | components/character/BossPicker.svelte | **gold** text, 2 px **gold** underline | white text, accent underline |
| 28 | Coach owned unit | components/coach/CoachTab.svelte | 2 px **gold** ring | accent ring |
| 29 | Raid match radio dot | components/guild/RaidMatchPicker.svelte | **gold** dot | accent dot |
| 30 | Premium feature row addressed by `#anchor` | components/premium/PlanFeatures.astro | gold tile, 3 px **gold** bar | accent tint, accent bar |
| 31 | Point order "Felsnap is level 34" cap line (planner) | ui/PointOrder.svelte | **gold** text and dash | white text, accent dash |
| 32 | Character selector current row (`csel-row.is-current`) | styles/nav.css (selector part only), nav/SelectorRow.svelte | accent 16 %, accent rail, accent check | the tokens; its focus ring was accent, now the site's gold |
| 33 | Overlay row `aria-selected` (typeahead, lookup) | ui/overlay.css | hover ground | unchanged: the keyboard cursor, not a choice |

### Left alone, on purpose

| Mark | Where | Colour | Why |
|---|---|---|---|
| Talent maxed / part ring, talent arrow on | ui/TalentTrees.svelte, SimTalentsSection.svelte, TalentEditor `path.on` | gold / green | The game's own talent frame |
| Ratings current tier | score/RatingsPage.svelte | the tier's colour | Data colour |
| Typeahead active crest | styles/typeahead.css | class colour | The keyboard cursor in a class list |
| Stat table sorted head | ui/stat-table.css | white | Sort state, with its arrow |
| Chart legend chip on/off | chart/ChartLegend.svelte | series colour vs dimmed | Visibility, not selection |
| Next point in the point order | ui/PointOrder.svelte `.em.nx` | dashed gold | A call to act, not a selection |
| Vote tags | ui/Votes.svelte | the vote's tone and dot | The vote's meaning; the button edge follows #25 |
| Checkboxes, switches | global.css, lib/guild/field-styles.ts | teal | Already the round-5 teal |

### For the integrator (batch E files)

- styles/nav.css `.site-group:has(> .site-door[aria-current='page'])`: the nav's current-page underline is `inset 0 -3px 0 var(--color-accent)`. Same look as the token; switching it to `var(--shadow-selected-underline)` is a one-line change for batch E's nav lane.
- styles/nav.css `.site-door-premium[aria-current='page']`: white ink on the teal Subscribe. Not a selection mark; leave it.
- This lane edited only the character selector's part of nav.css (`.csel-*`) and SelectorPanel.svelte, as the coordinator asked. Expect a small rebase there.

## Overlays: one edge (coordinator, 2026-11-09)

| Overlay | Before | After |
|---|---|---|
| Character and profile menu (SelectorPanel, `csel-panel`) | `ui-overlay` plus `ui-framed`: the outer edge and a cream inset line 6 px in (on the phone sheet too) | `ui-overlay` only, one edge |
| Nav door panels, header search, lookup typeahead, Select popovers, legend sheet, account delete dialog, UiSheet | one edge | unchanged |
| Client tooltip | black outer, grey edge, blue-grey inner line | unchanged: it is the client's tooltip, exactly |

ui/overlay.css now drops the inset frame (`.ui-overlay:is(.ui-framed, .ui-lead)::before { content: none }`), so an overlay cannot take a featured panel's frame again.
