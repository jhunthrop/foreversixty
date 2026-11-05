# Saved sim result view, run kind (experience spec)

Owner, 2026-11-04, on a competitor's saved sim report: "I like the sim view they have way better than ours."

This spec redesigns `/sim/<id>` for a `run` result: the page a player lands on from "Copy link", a Discord unfurl or their own history. It takes the competitor's information architecture and depth (every section visible on one scrolling page, a real fight timeline, a per-ability breakdown with outcome splits, a histogram, the build that produced the number), and goes past it on hierarchy, tooltips, honesty about old data and phone. It uses our design system and our voice. Nothing visual is copied from the competitor.

Bulk kinds (`gear`, `talents`, `drops`) and `weights` keep their own views (`SavedCombos`, `SavedWeights`). They change only where §4.9 says so.

## 0. Mocks

Folder: `design/mocks/2026-11-04-sim-results/`. Run `python3 design/mocks/2026-11-04-sim-results/gen.py` from the repo root. `gen.py` writes HTML under `html/` (git-ignored) and renders 2x PNGs with Playwright through `shoot.mjs`.

| Board | Shows |
|---|---|
| `result-1440.png`, `result-1024.png`, `result-2000.png` | The whole page, default state, desktop, tablet and wide |
| `result-390.png` | The whole page on a phone |
| `timeline-1440.png`, `timeline-390.png` | The timeline in Close up (2:30 to 3:00), Swings and ticks on, the crosshair at 2:50.0 with its readout, an Execute tick's tooltip (desktop) |
| `states-1440.png` | Loading skeletons at ready height, an older saved run's fallbacks, talents on an older layout, load failed |

**Data rules for the mocks.**

- Every figure, name and icon comes from the real run `design/mocks/data/simview/fury-result.json`: Fury warrior, level 60, raid buffs, 3:00 single target, 3,000 runs, 402.3 DPS ± 0.5, engine `d5c708834`.
- Names resolve in the order `sync-data.mjs` folds them: `spells.json`, then `data/curated/engine-spell-names.json`, then the class talent file. Icons come from `spellicons.json`, then the talent file. `data.py` lists the ten ids neither carries an icon for (`ICON_GAPS`). §11.4 makes the data lane add them.
- **Illustrative regions** draw data the engine does not return yet (§11). Each carries a dashed `ILLUSTRATIVE` tag on the board and nowhere in the build. They are:
  - the DPS spread bins: a normal curve at the run's own mean and spread;
  - the rage sources: modelled to sum to the run's real totals;
  - the target debuff column;
  - the timeline's On you and On the target bars: derived from cast times plus spell durations, and from Deep Wounds ticks.
- **Talents** draw the published Fury build from `bis/warrior-fury.json` (17/32/2). This run's stored string does not fit build 70790's trees (§4.7, open question 1).

## 1. The question and the answer in view

**The question:** "How much damage does this character do, and what produced it?"

**The answer that must be visible without scrolling:**

| Viewport | In the first screen |
|---|---|
| 1280 x 720 and up | The band: crest, spec, the DPS figure with its error, the facts row. The section nav. The top of Damage by ability (at least its first three rows at 1440 x 900). |
| 360 to 390 x 640 | Crest, spec, the DPS figure, "± 0.5 · 0.13% error" and "Run this yourself". |

The figure is the page's one primary element. Everything below it explains it.

## 2. Reference

| Surface players trust | What we borrow | What we do not |
|---|---|---|
| **Details! damage meter** (in game) | One row per ability: icon, name, bar, DPS, share. White swings are their own rows. Sorting by damage. | Its row-click drill-down. Our split is on the row. |
| **Warcraft Logs, Damage Done and Casts timeline** | The outcome split per ability (crit, miss, dodge, glance). One swimlane per ability with a tick per cast. Buff uptime bars on the same time axis. A hover readout. | Its multi-actor chrome. A sim has one player. |
| **SimulationCraft HTML report** | The DPS distribution histogram. Mean, error, min, max, iterations. Resource gained by source and spent by ability. | Its walls of tables. Its fixed-width layout. |
| **The in-game character panel** | The paperdoll (spec 2026-10-31) as the gear section, with the client's green enchant line under each item. | Nothing; it is the shared component. |
| **The in-game talent frame** | Three trees, ranks on the icon, read only. | Editing. That is the planner's job. |
| **Competitor saved sim** | The order "header, timeline, warnings, abilities, spread, talents, gear, resources, auras, rotation, settings, actions". All sections open on one page. The Recklessness-style warning. Copy JSON and download. | Its hero art, its about-this-snapshot prose, its early-rankings banner. Their label column truncates names on phone ("Bloodth..."); ours never truncates, it shows the icon only. |

Why ours is better than the competitor's (the review must be able to say this):

1. **Hierarchy.** The DPS is 64px gold. Every other fact is a 15px value in a quiet row. The competitor sets seven figures at the same size.
2. **Answer before diagnosis.** Damage by ability sits above the timeline. A player asks "where does my damage come from" before "when did I press it".
3. **Every ability and item is real.** Icon, rank and the shared spell or item tooltip, everywhere.
4. **Honest about old runs.** A saved sim missing new data says so in one line. The competitor's view assumes every run is current.
5. **Phone first.** The timeline keeps every lane on a phone: icons in a fixed column, a sideways-scrolling plot, a readout pinned under it. No truncated names.

## 3. Layout

### 3.1 Section order, top to bottom, and why

| # | Region | Why here |
|---|---|---|
| 1 | Result band | The answer: DPS, error and what the run was. The one primary action. |
| 2 | Section nav (sticky) | Replaces the tab strip. Depth stays visible; the nav only jumps. |
| 3 | Worth knowing (warnings) | Only when a rule fires (§3.5). It qualifies the number, so it sits right under it. |
| 4 | Damage by ability + DPS spread | What produced the number, then how much it varies. |
| 5 | Timeline | How one fight played out. Diagnosis after the answer. |
| 6 | Gear | This run's input, shown as the character screen. |
| 7 | Talents | This run's input, three trees. |
| 8 | Resource (Rage, Energy or Mana) | Where the resource came from and went. |
| 9 | Buffs and debuffs | Uptimes on you and on the target. |
| 10 | Rotation | What the engine was told to press, in words. |
| 11 | Settings and actions | Fight, buffs, consumables, engine; copy and download. |

**The tab strip does not survive.** `SimResults.svelte`'s eight tabs (Damage, Buffs, Debuffs, Casts, Resources, Timeline, One iteration, Distribution) hid seven of eight panels. Tenet 4 says the point is never behind a disclosure. Hierarchy (tenet 11) comes from order, card size and one figure per card instead.

**The sticky section nav earns its place.** The page is about five desktop screens and twelve phone screens. The nav is the only way back to Damage from Settings without a long scroll. It is the shared `TabStrip.svelte` in link mode (§3.4).

The "One iteration" tab's table becomes the timeline's cast list (§3.7). The "Debuffs" tab becomes the On the target column (§3.10). The old `summarySentence` line ("main-hand white hits and ... are 78% of your damage") is removed. The table says it better.

### 3.2 Page frame

- `SimTabs.astro` stays above the band, unchanged.
- `ScopeNote.astro` ("The simulator runs damage specs only...") is removed from `/sim/<id>`. A saved run is always a damage spec, so the note says nothing there. It stays on `/sim`.
- The band is full-bleed. Every region below it sits in the page column: max width 1344px, 48px gutters from 768px, 18px below, centred at every width including 2000px.
- Section gap 32px desktop, 22px phone. Cards are `card-surface` (12px radius, `--color-line-card` border, `--card-highlight` inset top line, `--shadow-card`). Warnings use `card-surface-quiet`.
- Every card opens with `CardHead.svelte`: a 36px `IconTile` glyph, an 11px gold micro label, the title in the heading face at 24px (20px below 768px) as an `h2`, and a right slot of 13px muted meta. The micro label and the meta for each card are in §6.
- Each card carries an `id` equal to its nav anchor: `damage`, `timeline`, `gear`, `talents`, `resource`, `buffs`, `rotation`, `settings`. `scroll-margin-top: 64px` keeps a jumped-to card clear of the sticky nav.

### 3.3 Result band (new `ResultBand.svelte` in `components/sim/`)

Reference: the Character header's three regions (who, the page's answer, the page's one action), on the `surface-band` ground the site uses for page heads.

- Ground `--surface-band`, full bleed. Inner wrapper 1344px with 48px gutters (18px phone). Padding 28px top (20px phone).
- **Who** (left):
  - `ClassCrestRing.svelte` at 64px (56px below 768px), ringed in the class colour.
  - Beside it, stacked: the eyebrow "SAVED SIM" (`label`, gold).
  - The `h1` (`heading-page`, 44/34px) in the class colour. It reads `result.title` if set, else the character's name if set, else the spec label ("Fury Warrior").
  - The descriptor at 13/22px muted: race, spec and class in `--text` 600, then "Level 60", then the talent split ("Orc Fury Warrior · Level 60 · 20/31/0 talents"). Each piece appears only when known. Separators are drawn pseudo-elements, as in the profile header, so no line starts with "·".
  - The next-step line at 14px `--text`: `afterSimSentence(topUpgrade)` when a Top Gear upgrade is on record. Otherwise "Next: find the upgrade that adds the most." and an `ActionLink` "Run Top Gear" to `/sim/gear`. This is tenet 14's "the next step that improves it".
- **The answer** (right from 1024px, under Who below that):
  - The figure in the heading face, 800 weight, 64px (56px phone), the gold gradient clipped to text, `tabular-nums`, one decimal. "DPS" follows at 18px heading face, `--text-muted`, letter-spacing .04em.
  - Under it at 15px: "± 0.5 · 0.13% error". The margin is `--text` 600; the rest is muted.
  - Then the actions, right-aligned with a 14px top gap: the **primary** button "Run this yourself" (play glyph) and the **secondary** "Copy link" (link glyph). Both are `button()` md, 44px. Below 768px they stack full width, primary first.
- **Facts row** (full width under both, 24px top gap, a 1px `rgba(229,185,85,.14)` top rule):
  - Each fact is a `quiet-label` (13px muted) over a 15px 600 `--strong` `numeric` value. Facts are divided by 1px `--border-soft` verticals with 24px padding.
  - Facts, in order: Fight, Targets, Buffs, Actions per minute, Range, Runs, Saved.
  - Below 768px the facts form a two-column grid with 1px row rules.
  - "Saved" is the page's one date stamp. There are no pills and no engine sha in the band; the engine lives in Settings.
- Values:
  - **Fight:** "3:00 ± 20%" from `encounter.duration_sec` and `variation`.
  - **Targets:** "1".
  - **Buffs:** "Raid" when the request's buff list equals the raid preset, "None" when empty, else "Custom".
  - **Actions per minute:** player button presses divided by the average fight length. A press is a `summary.casts` row whose key is not `other:*` and has no `/<tag>` suffix. This run: 75 presses over 2:59.5 is "25.1".
  - **Range:** "311 to 511" from `dps.min` and `dps.max`.
  - **Runs:** `iterations_run`.
  - **Saved:** the sim's `created_at` (new, §11.2) as "18 Sep 2026". Until the field exists, the fact is left out. `source.captured_at` is the character's capture time, not the save time, so it is never labelled "Saved".

### 3.4 Section nav

Reference: the report's underline tabs, which players already use on `/reports/<id>`.

- `TabStrip.svelte` in link mode (no `onSelect`): a `nav` of `#anchor` links. The strip is 48px tall and sticky at `top: 0`. Its ground is `rgba(7,9,13,.94)` with a 1px `--border-soft` bottom rule, at `z-index` 5.
- Labels: Damage, Timeline, Gear, Talents, the resource name (Rage, Energy or Mana from the player's power type), Buffs, Rotation, Settings. A section that does not render (no gear on file) drops its link.
- **Current section:** an `IntersectionObserver` marks the link whose card crosses 30% of the viewport from the top. It gets `aria-current="location"`, `--strong` ink and the 2px gold underline. Clicking a link scrolls smoothly (instantly under reduced motion) and moves focus to the card's `h2` (`tabindex="-1"`).
- **Compact figure:** from 1024px, once the band's figure has scrolled out, the strip's right end shows the figure: "402.3" in the heading face at 18px gold, then "DPS" 13px muted. It fades in over 120ms. It is `aria-hidden`, since the `h1` region already states it.
- Phone: the strip scrolls sideways with TabStrip's 24px edge fades. The compact figure is not shown.

### 3.5 Worth knowing (warnings)

Reference: the competitor's "the sim emitted warnings" box, minus its hide control. Notes are few and each one is true, so none is dismissible.

- `card-surface-quiet` with a 3px `--ember` left edge. CardHead: the alert glyph, label "THIS RUN", title "Worth knowing", meta "1 note" or "2 notes".
- Each note is one row, at least 56px, with 12px by 20px padding and 1px `--border-soft` rules between rows. The row holds:
  - a 28px spell icon, or the alert `IconTile` when no spell is involved;
  - one 15px sentence;
  - an optional trailing `ActionLink`.
- Computed in a new `lib/sim/warnings.ts` from data the result already has. At most three, in this priority order. With none, the card does not render (no "all clear" line).

| Id | Fires when | Copy |
|---|---|---|
| W1 stale | `isStale(result.engine_version)` | "This run used an older engine. A run today may give a different number." + ActionLink "Run this yourself" |
| W2 missing talent | the talents validate (§4.7) and a rotation line casts a talent ability with 0 points in this build | "{Bloodthirst} is in the {Fury} rotation, but these talents do not take it." + ActionLink "Open in planner" |
| W3 long cooldown | a cast spell has `cooldown_ms` greater than the longest possible fight (`duration_sec x (1 + variation)`) and was cast at least once | "{Recklessness} has a {30 minute} cooldown, longer than this fight, so the result counts one use every pull. In a raid it is back only every {30 minutes}." |
| W4 resource at cap | `wasted / gained ≥ 5%` for rage or energy | "{n} {rage} ({p}%) was lost at the cap." Rage adds "Queue Heroic Strike sooner so swings do not land at full rage." Energy adds "Spend before the bar fills." |
| W5 wide error | `dps.error / dps.mean > 0.5%` | "This run's error is ± {x} DPS. A difference smaller than that between two runs is noise." + ActionLink "Run this yourself" |

This run fires W3 only: Recklessness has a 30 minute cooldown against a longest fight of 3:36. W4 does not fire: 14 of 1,109 rage is 1.3%.

### 3.6 Damage by ability + DPS spread

Grid from 1280px: `minmax(0,8fr) minmax(0,4fr)`, a 16px gap, top-aligned. From 1024 to 1279px both cards are full width, and the spread card puts its tiles to the right of its histogram (7fr and 5fr). Below 1024px both cards are full width and stacked.

**Damage by ability.** A new `AbilityTable.svelte` in `components/report/`. It is extracted from `ActorRow.svelte`'s expanded ability table so the logs report gets the same rows; `ActorRow` mounts it. Reference: Details! rows plus the WCL outcome split.

- CardHead: label "THIS RUN", title "Damage by ability", meta "402.3 DPS over 2:59 on average".
- A table (`table`, real `th` with `scope="col"`). The header row is 40px with `quiet-label` headers; numeric columns are right-aligned.
- **Columns at 1024px and up:**
  - **Ability** (1fr, min 240): a 28px icon (6px radius, 1px black border, 1px `rgba(229,185,85,.18)` outer ring), then the name at 15px 600 `--strong` with " Rank 5" in 13px muted.
  - **DPS:** 15px 600 `--strong`, one decimal.
  - **Share:** a 120 by 8px bar (track `--border-soft`, fill `--gold-deep` to `--gold` gradient) scaled to the top row, then the percentage at one decimal, 48px wide.
  - **Casts:** a whole number. A white swing row counts swings. A periodic row shows a muted "–".
  - **Average:** the average per cast, swing or tick, as a whole number with thousands separators.
  - **Crit:** a percentage, one decimal.
  - **Uptime:** for a periodic row, the target debuff's uptime. Otherwise a muted "–". The column is dropped when no row has a value.
- **The sub-line** under each name, at 13/18px muted with the figures in `--text`:
  - The outcome split, as a share of attempts, in the order Miss, Dodge, Parry, Glance, Block: "Miss 24.6% · Dodge 6.6% · Glance 39.3%". A zero outcome is left out.
  - A periodic row reads "58 ticks · tick crit 32.8%".
  - The main-hand row adds "28 swings became Heroic Strike". The figure is the Heroic Strike landing count (`spell:11567`), the swings it replaced.
- **White swing rows:** the name is "Main hand" or "Off hand". The icon is the equipped weapon's icon, as the in-game Attack button shows it. After the name comes " · " and the weapon's name in its quality text colour, with the shared item tooltip on hover or focus.
- **Pet rows** (hunter, warlock): indented 24px under a "Pet" sub-heading row with the pet's name. Each row reads as above.
- **Sorting:** DPS descending by default. The headers DPS, Share, Casts, Average and Crit are `button`s inside the `th` with `aria-sort`. Sorting is descending first, a second press sorts ascending. The sorted header is `--strong` with a gold ▾ or ▴.
- **Row states:**
  - Hover (fine pointer) gives the row `--row-hover`.
  - The icon and name are one focusable trigger (`data-spell-key`, 44px tall on phone). It opens the shared spell tooltip (`GameTooltip.svelte` through `BisTooltipHost`'s game-tip mode) on hover, focus or tap. Escape closes it, as everywhere else.
- **Total row:** muted 13px: "Total, 7 sources", the summed DPS, "100%", and "72,222 damage a fight".
- **Phone (below 768px):** two columns, Ability and DPS. The DPS cell stacks the DPS over the share percentage, muted. The sub-line gains a second line: "19 casts · 1,172 avg · crit 47.4%". The share bar is not shown on phone; the percentage carries it.

**DPS spread.** `DpsDistribution.svelte`, rewritten to draw the engine's real bins (§11.2). The current normal-curve drawing is removed, because a drawn shape the run did not produce breaks tenet 8.

- CardHead: label "3,000 RUNS", title "DPS spread", no meta.
- **Histogram:** 168px tall, at most 40 bars with 2px gaps. The engine's 1 DPS bins are grouped client-side into `ceil((max - min) / 40)` DPS buckets.
  - Bars inside the 16th to 84th percentile are `rgba(229,185,85,.7)`; the rest are `--border`.
  - A 2px `--gold-light` mean line carries a 13px label above it ("402.3").
  - Under the bars, the axis: min, "DPS per run", max.
- **Readout:** hovering or focusing a bar (the histogram is one `role="img"` with an `aria-label`, plus arrow-key stepping when focused) shows `ChartTooltip`: "400 to 405 DPS: 212 runs".
- **Caption** at 13px muted: "Two runs in three land between 371 and 431 DPS. Your pulls will spread the same way." The two figures are the 16th and 84th percentiles of the real bins.
- **Six StatTiles** in three columns (two on phone): Mean, Error, Spread (the standard deviation, as "± 28.2"), Lowest, Highest, Runs.

### 3.7 Timeline (new `SimTimeline.svelte` in `components/sim/`, lazy chunk)

Reference: WCL's casts timeline and the competitor's lanes, built on the report's own chart parts: `lib/chart/time-axis.ts` (`timeTicks`, `clockLabel`), `SyncedCrosshair.svelte`, `OverviewBrush.svelte`, `ChartTooltip.svelte`, `PresetChips.svelte`, `Switch.svelte` and `ChartSrTable.svelte`.

`TimelinesView.svelte` stays the logs view. It draws one lane per player; this view draws one lane per ability. They share axis, crosshair and brush code, not a component.

**Source:** `result.sample`, the median-DPS run's events, plus `sample_auras` and `sample_duration_ms` when present (§11.2). The axis runs from the first event (pre-pull, -2.5 s here) to `sample_duration_ms`. Without that field, the axis ends at the last event plus 1 s. CardHead meta: "The median run · 3:34 long" (`m:ss` of the sample's own length, which differs from the 3:00 average fight because fights vary ± 20%).

**Toolbar** (min height 56px, 10px by 20px padding, a 1px top rule, wraps):

1. **Zoom:** `PresetChips` with two chips, "Whole fight" and "Close up". Close up opens a 30 s window centred on the crosshair, or at the pull when there is none.
2. **Swings and ticks:** a `Switch`, off by default. On, it adds the group "Swings and ticks" (Main hand, Off hand and each periodic effect's ticks).
3. **Legend** (right-aligned, 13px muted): Damage (3 by 12px `--color-school-<school>` tick), Cooldown (3 by 12px `--text` tick), On you (16 by 8px `--text` at 35%), On the target (16 by 8px `--gold-deep` at 78%), Rage (16 by 2px `--gold` line). The legend's resource word follows the player.

**Plot** (desktop): a 200px label column and the plot filling the rest.

- **Axis row**, 30px: 13px 600 muted labels every 15 s (whole fight, 1280px and up), 30 s (below 900px of plot), 5 s (Close up). There is a dashed `--border-warm` line at 0:00 labelled "Pull" in gold, and 1px `--border-soft` gridlines through every lane.
- **Groups**, each with a 30px heading row (13px 600 muted): Casts; Swings and ticks (when on); On you; On the target. A group with no lanes does not draw.
- **Lanes**, 24px each. The label is an 18px icon and the 14px name. A name never truncates: the label column is 200px, and a longer name wraps the column to 240px for that page.
- **Lane order:** Casts by first use in the sample. On you by first application. On the target by first application.
- **Lane membership:**
  - Casts: one lane per base spell id. A tagged variant folds into its base: `spell:11567/1`, the Heroic Strike queue press, draws as a hollow tick (1px outline) on the Heroic Strike lane, and the landing `spell:11567` draws as a filled tick.
  - Swings and ticks: `other:attack/1`, `other:attack/2`, and any key whose summary row has `ticks > 0` (Deep Wounds here).
- **Ticks:** 3 by 14px, 1px radius, centred in the lane.
  - Fill: a damaging cast is its school token (`--color-school-physical` for a warrior). A cast with no damage row (Bloodrage, Death Wish) is `--text`.
  - Swings and periodic ticks are 2 by 10px at 80%.
  - **With outcome data** (§11.2): a crit is 4 by 18px in `--gold-light`; a miss, dodge or parry is a 1px `--text-muted` outline; a glance is the fill at 60%.
- **Bars:** 10px tall, 3px radius, centred. On you is `--text` at 35% (the report's own aura bar). On the target is `--gold-deep` at 78%. A stacking aura prints its stack count at 11px inside the bar when the bar is at least 18px wide.
- **Execute phase:** from the first Execute in the sample to the end, a `rgba(214,110,40,.07)` wash over every lane, with "EXECUTE" at 11px 700, .14em tracking, `--ember`, 6px from its left edge. It draws only for a class with an execute spell in its rotation (warrior Execute; later hunter Kill Command if added). It is labelled from the first use, so it never claims a health threshold.
- **Resource lane:** 72px. The label is "Rage" (14px 600) over "peak 100 of 100" (13px muted). The plot is a step line, 1.5px `--gold`, over a `rgba(229,185,85,.12)` area, with a dashed `--border-warm` line at the cap. Values are the sample's `resources` after each event, held until the next event.

**Brush:** `OverviewBrush` under the plot, 44px. It draws the whole fight's resource line at 1px `#5d5849`. In Close up it draws the window: a `rgba(229,185,85,.10)` fill, a 1px gold border and two 6 by 20px gold handles. Dragging the window pans it; dragging a handle resizes it (10 s minimum). Arrow keys pan 5 s; Shift with an arrow key resizes. Whole fight hides the window.

**Crosshair and readout:**

- Moving a fine pointer over the plot draws a 1px `--gold-hover` vertical line. A 52 by 20px gold time chip sits on the axis ("2:50").
- The readout is a 280px `--bg-raised` panel, 12px radius, `--border`, the tooltip shadow. It sits 12px right of the line, or flips left near the right edge. It holds:
  - the time ("2:50.0", heading face 18px);
  - "Rage 7 · On you: Recklessness" at 13px muted with the figure strong;
  - a rule;
  - the five most recent casts, each a 20px icon, the name, and "3.8 s ago" right-aligned, muted.
- Swings and ticks are not listed in the readout.
- Click or Enter pins the readout. A pinned readout gains a close button (44px hit, × glyph). Escape or a second click unpins it.

**Tick tooltip:** hover or focus on a tick opens `ChartTooltip`, 250px wide:

- line 1: the 24px icon, the name and the rank;
- line 2: "2:46.2 · rage after 0";
- line 3, muted: "Next Execute 4.0 s later", or "Last Execute of the fight";
- with outcome data, a line 4: "Crit · 1,284 damage".

Without outcome data, line 4 is absent. The board's line "Hit or crit shows here once the run records it." is a mock note, not copy.

**Cast list** (under the brush): the old One iteration table, rebuilt on `SampleLog.svelte`.

- A 52px foot row: the quiet label "Casts: 86 of 86. Swings and Deep Wounds ticks are on the lanes above." In Close up it reads "Casts in view: 12 of 86". A secondary sm button follows: "Show all 86" or "Show all 12".
- Rows are 40px: time ("2:41.2", numeric muted, 72px), a 24px icon, the name ("Heroic Strike queued" for the queue press), and "Rage 100" right-aligned, muted.
- Collapsed, the list shows 5 rows: in Close up the five nearest the window's start; in Whole fight the first five. Expanded, it shows every row in view. "Show fewer" collapses it again.
- The list is the timeline's accessible table: it is what a screen reader reads (§4.3).

**Keyboard:**

- The plot is one tab stop: `role="application"` with `aria-roledescription="timeline"` and an `aria-label` of "Fight timeline, 3:34 long, 86 casts. Arrow keys move between casts.".
- Left and Right move the crosshair to the previous or next cast across visible lanes. Up and Down move to the nearest cast in the lane above or below. Home and End go to the first and last.
- Enter pins the readout. `+` and `-` switch Close up and Whole fight. Escape unpins.
- Each move announces the readout's first two lines through a polite live region.

**Touch:**

- Tap a tick for its tooltip. Tap the empty plot to place the crosshair; this pins the readout.
- The plot never captures vertical swipes (`touch-action: pan-x pan-y`).

**Phone (below 768px):**

- The label column is 44px: the 18px icon only, `aria-label` set to the name. Group headings show as a 12px muted label in the plot's top-left, on the card ground, so lanes keep their labels. The resource lane shows "Rage" at 11px.
- The plot **scrolls sideways** inside the card. The label column and the axis row stay fixed: the label column is `position: sticky; left: 0`, and the axis lives inside the scroller with `position: sticky; top: 0`. A 24px fade marks the clipped right edge.
- Scale: Whole fight is 4px per second (214.5 s is 868px). Close up is 12px per second.
- The readout does not float on a phone. It is pinned **under the plot**, full width, after the brush, and updates on tap.
- The tick tooltip opens as the same pinned panel. Nothing is cut; the toolbar wraps to three lines.

**Loading placeholder:** the chunk loads with `createLazyComponent` once the card nears the viewport (`IntersectionObserver`, 600px root margin) or on the first nav click. Until then the card renders its CardHead and a `Skeleton` sized to the ready height. That height is computed in the main chunk from data it already has:

`56 (toolbar) + 30 (axis) + 30 x groups + 24 x lanes + 72 (resource) + 44 (brush) + 52 (cast foot) + 40 x 5 (cast rows)`

The lane count follows §3.7's membership rules with Swings and ticks off. The swap is the one 160ms `.reveal`. Nothing moves.

### 3.8 Gear

`Paperdoll.svelte` (spec 2026-10-31), wrapped in a Card.

- CardHead: shield glyph, label "THIS RUN'S INPUT", title "Gear", meta "17 of 17 slots · 11 enchants".
- **One change to the shared component:** an enchant line. A slot whose gear entry has an `enchant` gains a line between the name and the slot line: 13/18px in `--color-rarity-uncommon` (the client's green enchant line), led by a 12px enchant glyph from `glyphs.ts` (four-point spark, 2.4px stroke). The glyph is there because an uncommon item's name is the same green. The text is the enchant's display name: `enchants.json` `name` after " - " ("Crusader", "Greater Strength"); an arcanum keeps its full name ("Lesser Arcanum of Voracity").
- The paperdoll's other parts work as specced there: hover, focus and tap open the shared item tooltip, the right column opens toward the centre, and the centre shows the crest and the average item level.
- **Foot** (52px, 1px top rule):
  - one 13px muted line about sets. With active bonuses: "Set: {name} 2 of 5 · (2) {bonus text}", one line per set, the active bonus text in `--color-rarity-uncommon` and inactive ones muted. With none: "No set bonus is active: no two pieces share a set."
  - then, right-aligned, the ActionLink "Upgrades for this character on Best in slot" to `/bis/<class>/<spec>`. This card never shows a gain or a pick. That is /bis's job.
- Phone: the paperdoll's own under-640 fold (identity on top, two mirrored columns). The enchant line wraps under the name and is never cut.

### 3.9 Talents

`GuideTreeColumns.svelte` (the guides' read-only three-column trees) over `guideTreeStore`, fed the stored talents string. The run's code comes from `codeForCharacterSpec`, as `plannerHref` already does.

- CardHead: label "THIS RUN'S INPUT", title "Talents", meta "17 / 32 / 2" (numeric).
- **Desktop and tablet:** three tree panels in a row, 12px gap. Each panel:
  - shows the tree's art behind a `rgba(7,9,13,.55)` to `.88` vertical scrim;
  - has a heading row with the tree name (heading face 18px) and its points in gold;
  - draws 40px cells: a full talent with a 2px gold ring, a partial one with a 2px `--kill` ring, an untaken one greyscale at 45% brightness, and the rank badge "3/3" at the bottom right.
- **Phone:** `TreeTabs.svelte` above one tree. Three 44px tabs ("Arms 17", "Fury 32", "Protection 2"); the tab with the most points is selected and ringed in the class colour. Arrow keys move between tabs (TabStrip's roving tabindex).
- Hover, focus or tap on a cell opens the talent's `GameTooltip`, the planner's own.
- **Foot:** "Read only. Hover or focus a talent for its tooltip." (on phone: "Read only. Tap a talent for its tooltip.") and the ActionLink "Open in planner" to `plannerHref`.
- §4.7 covers a talent string that does not fit today's trees.

### 3.10 Resource and auras

**Resource card** (title by power type: "Rage", "Energy", "Mana"). Reference: SimulationCraft's resource tables, our bars.

- CardHead: bolt glyph, label "AVERAGE FIGHT", meta empty.
- **Three StatTiles:** Gained, Spent, and "Lost at the cap" with the share ("14 · 1.3%"). For mana the third tile is "Ended with" (mana left at the end), since mana has no cap loss.
- **Two lists** side by side (stacked on phone), "Gained from" and "Spent on", each headed by a `quiet-label` row with the list total.
  - Rows are 40px: a 24px icon, the name, a 96px bar (track `--border-soft`, fill `--gold`, scaled to the list's top row) and the amount (600 `--strong`, whole numbers; one decimal under 10).
  - A gain row with `wasted > 0` adds the overflow after the amount at 13px in `--ember`: "−1.3 lost".
  - Rows come from `resource_sources` (§11.2).
- `ResourceGraphs.svelte` is not used here; the timeline's resource lane is the graph.

**Buffs and debuffs card.** `AuraTable.svelte`, given a `variant="bars"` that renders the rows below. The logs report keeps the default variant.

- CardHead: spark glyph, label "AVERAGE FIGHT", title "Buffs and debuffs", meta "Uptime, then times applied".
- **Two columns** (stacked on phone), "On you" and "On the target":
  - each column heads with a `quiet-label` row and a count ("9 that come and go");
  - rows are 40px: a 24px icon, the name, a 96px bar (On you: `--text` at 55%; On the target: `--gold-deep`), the uptime at one decimal (600 `--strong`) and the count "68×" muted;
  - rows sort by uptime, descending.
- **Up all fight:** auras at 99.9% or more leave the bar list. They gather under a 13px muted "Up all fight" label (On the target: "Up all fight, from the raid") as a wrapping row of 20px icon plus 14px name.
- The Heroic Strike queue track (`spell:11567/1`) is not a buff and is left out. It is drawn on the timeline.
- Two tracks of one spell with hand tags (`spell:20007/1`, `/2`) read "Holy Strength (main hand)" and "(off hand)". This is verified in §11.4.
- **Phone:** the bar is 56px wide.

### 3.11 Rotation

`RotationCard.svelte`, rebuilt from a disclosure into an always-open list (tenet 4). `RotationDisclosure` stays for the settings bar on `/sim`.

- CardHead: list glyph, label "{SPEC}, AT {LEVEL}" ("FURY, AT 60"), title "Rotation", meta "12 lines · casts are per fight".
- **"Before the pull"** sub-heading (`quiet-label`, 14px by 20px padding), then one row per pre-pull action:
  - its time at 14px 600 muted, nowrap ("-2.5 s");
  - a 28px icon;
  - the name with its rank.
- **"Then, from the top at every decision"**, then one row per priority line. The row grid is `44px 28px 1fr 140px`, 12px gap, 12px by 20px padding, a 1px top rule. It holds:
  - the line number;
  - the icon;
  - the name and rank (15px 600), with the line's note under it in 14/21px `--text`, max 860px;
  - the average casts per fight from `summary.casts` ("19 casts", "1 cast"), right-aligned.
- The notes are the curated APL's own `notes` (`rotations.json`, the guide voice), verbatim. A line with no note shows the name only.
- **The "Other cooldowns as they come up" line** (`autocastOtherCooldowns`) lists this run's casts that no other line names: "This run: Blood Fury.".
- **Not cast:** a line whose spell has 0 casts is drawn at 50% (icon, name and note) and its count reads "Not cast in this run" in 13px muted. W2 (§3.5) explains the case where a talent is missing.
- **Phone:** the count moves to the name line, floated right, 13px muted.
- **Foot:** "The {Fury} rotation the engine follows. A line with no cast in this run is dimmed." and the ActionLink "Read the {Fury} guide" to `/guides/<class>/<spec>`.
- When the result is stale (W1), the meta adds " · today's rotation", because the stored result does not record the rotation it ran (open question 3).

### 3.12 Settings and actions

`DetailsCard.svelte`, rebuilt.

- CardHead: gear glyph, label "THIS RUN'S INPUT", title "Settings".
- **Left column:** key/value rows, 44px, 1px `--border-soft` rules, the key muted and the value 600 `--strong` right-aligned.
  - Fight length: "3:00, ± 20% (2:24 to 3:36)".
  - Targets: "1".
  - Execute phase: "Last 25%".
  - Race: "Orc".
  - Runs: "3,000".
  - Ran on: "Our server" or "Your browser".
  - Engine: the 7-character sha in the mono face, with the stale `Tag` (muted, "Older engine") when W1 fires.
- **Right column:** three groups, each headed by a `quiet-label` row with a count: On you (26), On the target (20), Consumables (0).
  - Chips are 30px: 4px 10px padding, a 22px icon, the 13px name and a 1px `--border`, 6px radius. Names and icons come from `simbuffs.json`.
  - An empty group reads "None in this run.".
  - A buff's group (on you or on the target) comes from a new `target: true` flag in `simbuffs.json` (§11.4); there is no hand-kept list in the web.
- **Phone:** each group shows its first 6 chips and a secondary sm button "Show all 26". This is the one disclosure on the page, allowed because the buff list is reference, not the point.
- **Foot actions** (52px, wraps; on phone a two-column grid with "Copy for the game" full width first):
  - "Copy for the game": `CopyForGame.svelte`, unchanged.
  - "Copy JSON": copies `JSON.stringify(result, null, 2)`.
  - "Download JSON": downloads `forever-sixty-sim-<sim_id>.json`.
  - the 13px muted line "The JSON is this whole result, request included."

## 4. States

### 4.1 Every interactive element

| Element | Default | Hover (fine pointer) | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Run this yourself (primary) | gold gradient | brightness 1.06 | 2px gold ring, 2px offset | `--gold-hover` ground | never | — | — | — |
| Copy link, Copy JSON | secondary | border 30% gold | ring | 10% gold ground | — | — | — | clipboard refused: label "Copy failed", 3 s, then back |
| Copied state | label "Copied", check glyph, 2 s | — | ring kept | — | — | — | — | — |
| Download JSON | secondary | as above | ring | as above | — | — | — | — |
| Copy for the game | as `CopyForGame` | — | — | — | talents not loaded: disabled 45% | — | — | as `CopyForGame` |
| Section nav link | `--nav` | `--strong` | ring inset 2px | `--gold-hover` | — | — | — | — |
| Current nav link | `--strong`, gold underline | — | ring | — | — | — | — | — |
| Sort header | `--muted` | `--strong` | ring | — | — | — | — | — |
| Ability icon and name | as drawn | underline name | ring | — | — | — | — | tooltip data failed: the name shows plain, no tooltip, no error line |
| Zoom chips | `PresetChips` states | | | | | | | |
| Swings and ticks | `Switch` states | | | | | | | |
| Plot | — | crosshair | 2px inset gold ring on the plot | pinned readout | — | skeleton (§3.7) | §4.3 | §4.4 |
| Tick | as drawn | 1.15 brightness + tooltip | 2px gold outline + tooltip | — | — | — | — | — |
| Brush window and handles | as drawn | handles brighten | ring | grabbing cursor | Whole fight: window hidden | — | — | — |
| Histogram | — | bar `--gold-light` + tooltip | ring + arrow keys step bars | — | — | skeleton 168px | §4.3 | — |
| Talent cell | as drawn | brightness 1.12 | ring | — | — | `GuideTreeColumns` skeleton | §4.7 | `LoadError` "Talents didn't load." + Try again |
| Paperdoll slot | per paperdoll spec | | | | | | | |
| Show all (cast list, chips) | secondary sm | as secondary | ring | — | — | — | hidden when everything shows | — |
| ActionLinks | `ActionLink` states | | | | | | | |

### 4.2 Page loading

- The prerendered shell's `SIM_SAVED_SKELETON_HTML` (`lib/sim/skeleton.ts`) is redrawn to the new order at ready heights:
  - the band (desktop 268px, phone 492px);
  - the nav (48px);
  - the Damage card (64px head + 40px header row + 8 x 56px rows + 44px total: 596px);
  - the spread card at the same height;
  - the timeline's head and toolbar.
- `[id].astro`'s reservation becomes the measured ready height of the fixture page at 360 and 1280 plus the existing 24px CI-fonts buffer. The comment block records both numbers.
- The band, nav and damage card fill in place with one `.reveal`. Cards below the fold render from the same data in the stages `SavedSim` already uses (`yieldToMain`).

### 4.3 An older saved result (missing the new fields)

Each section with new data has one quiet, honest fallback. Never an empty chart, never a broken one.

| Missing | Fallback |
|---|---|
| `dps_histogram` | The spread card draws a range strip instead of bars: an 8px `--border-soft` track from min to max, a `rgba(229,185,85,.35)` band at mean ± one spread, a 2px `--gold-light` mean tick. Axis "311 / mean 402.3 / 511". Caption: "This run kept the range, not the shape. The shaded part is one spread either side of the mean." The six tiles stay. |
| `sample_auras` | The On you and On the target groups are replaced by one 13px muted row inside the plot area: "Buff and debuff bars need a run from {launch date of the field} or later. The casts and rage here are from this one." and the ActionLink "Run this yourself". The Execute wash and the resource lane still draw. |
| `sample_duration_ms` | The axis ends at the last event plus 1 s. No message. |
| `outcome` on sample events | Ticks are one colour; the tooltip has no outcome line. No message. |
| `sample` entirely | The timeline card renders its CardHead and one line: "This run has no recorded fight to draw." The nav drops "Timeline". |
| `resource_sources` | The three tiles stay. The lists are replaced by "Where the {rage} came from and went needs a newer run." |
| Target debuff tracks | The On the target column shows the raid debuffs from the request's buff list as the Up all fight row, under the line "Uptime on the target needs a newer run." The Uptime column of Damage by ability is dropped. |
| `created_at` | The Saved fact is left out. |

### 4.4 Errors

- **The result does not load** (404): the existing not-found view, unchanged.
- **The result does not load** (5xx or network): one Card in the page column. CardHead: alert glyph, label "SAVED SIM", title "This sim did not load". Body: "The server did not answer. Your link is fine; try again in a moment." and a secondary "Try again" button. Nothing else renders.
- **A lazy chunk fails** (timeline, talents): that card shows `LoadError` with "This part of the page didn't load." and "Try again". The rest of the page stands.
- **The items file fails:** the paperdoll's own failed-file state.
- **The names file fails:** rows read "An unnamed spell" as today (`resolveActionName`). Icons fall back to the skeleton box, never a broken image.

### 4.5 No gear on file

The Gear card shows its CardHead and "This run had no gear on file." The nav drops "Gear".

### 4.6 Partial or aborted result

A saved result is never partial: saving refuses an aborted run. No state is needed.

### 4.7 Talents that do not fit today's trees

`talentsFit(string, trees)` is a new `lib/sim/talents-fit.ts`. It requires all of:

- one dash-separated part per tree;
- one digit per talent in each part;
- every digit at most the talent's `max_rank`;
- every prerequisite rank met;
- every tier's 5-points-per-tier gate met (the planner's `rules.ts`).

On failure, the card keeps its head with the stored split as meta ("20 / 31 / 0") and shows one 15px line:

"These talents were saved on an older talent layout. The point split above is as saved; the trees have changed since, so they are not drawn."

There is no planner link, and W2 is not evaluated. **This run fails the check**: its Fury part has 17 digits against 18 talents, and Arms digit 5 sits on a 3-rank talent. See open question 1.

### 4.8 Stale engine

W1 fires (§3.5). The Settings engine row carries the "Older engine" Tag. The rotation meta adds " · today's rotation". The figure is not dimmed: it is what that engine produced.

### 4.9 Other kinds

`gear`, `talents`, `drops` and `weights` keep `SavedCombos` and `SavedWeights`. They take two parts of this spec, so the site reads as one:

- the result band, with the kind title as `h1`, the kind's own answer as the figure (the top combination's gain, or "Weights" with no figure), the primary button "Run this yourself", and only the Fight, Targets, Buffs, Runs and Saved facts;
- the Settings card.

They take no section nav and no timeline.

## 5. Hierarchy

| Region | The one primary element | Secondary, visibly so |
|---|---|---|
| Band | The DPS figure (64px gold gradient) | Error 15px muted; facts 15px values under 13px labels; the h1 in class colour at 44px is identity, not answer |
| Nav | The current link | Other links `--nav` |
| Worth knowing | The sentence | The icon; the link |
| Damage by ability | The DPS column (600 `--strong`) and the share bar | Sub-lines 13px muted; casts, average, crit at 15px regular |
| Spread | The histogram | The caption; tiles at 24px (smaller than the band's figure) |
| Timeline | The cast lanes | Bars (35 to 78% alpha); gridlines `--border-soft`; the brush muted |
| Gear | The items (quality colours) | Enchant 13px; slot and item level 13px muted |
| Talents | The taken cells (gold rings) | Untaken cells greyscale |
| Resource | The three tiles | List rows |
| Auras | The bars that come and go | Up all fight as icon text only |
| Rotation | Line names | Notes in `--text`; counts right; untaken lines at 50% |
| Settings | The values | Chips |

There is exactly one filled gold element on the page: "Run this yourself". Whitespace: 32px between cards; 16 to 20px inside them; no card holds two figures at the band's size.

## 6. Copy (verbatim)

No em dashes anywhere. Figures are examples from this run.

**Band:** "SAVED SIM" · h1 per §3.3 · "Orc Fury Warrior · Level 60 · 20/31/0 talents" · "Next: find the upgrade that adds the most." · "Run Top Gear" · "DPS" · "± 0.5 · 0.13% error" · "Run this yourself" · "Copy link" · "Copied" · "Copy failed" · facts "Fight", "Targets", "Buffs", "Actions per minute", "Range", "Runs", "Saved"; values "3:00 ± 20%", "Raid" / "None" / "Custom", "311 to 511", "18 Sep 2026".

**Nav:** "Damage", "Timeline", "Gear", "Talents", "Rage" / "Energy" / "Mana", "Buffs", "Rotation", "Settings". Nav `aria-label` "Sections of this result".

**Worth knowing:** label "THIS RUN", title "Worth knowing", meta "1 note" / "{n} notes". W1 to W5 as in §3.5.

**Damage by ability:** label "THIS RUN", title "Damage by ability", meta "{402.3} DPS over {2:59} on average". Headers "Ability", "DPS", "Share", "Casts", "Average", "Crit", "Uptime". Sub-lines "Miss {x}%", "Dodge {x}%", "Parry {x}%", "Glance {x}%", "Block {x}%", "{n} ticks · tick crit {x}%", "{n} swings became Heroic Strike". Phone sub-line "{n} casts · {avg} avg · crit {x}%" ("swings" for a white row). Names "Main hand", "Off hand". Pet sub-heading "Pet · {name}". Total "Total, {n} sources" and "{72,222} damage a fight".

**Spread:** label "{3,000} RUNS", title "DPS spread", axis middle "DPS per run", caption "Two runs in three land between {371} and {431} DPS. Your pulls will spread the same way." Tiles "Mean", "Error", "Spread", "Lowest", "Highest", "Runs". Bar tooltip "{400} to {405} DPS: {212} runs". Image label "DPS of each of {3,000} runs, from {311} to {511}, most near {402}".

**Timeline:** label "ONE FIGHT", title "Timeline", meta "The median run · {3:34} long". Chips "Whole fight", "Close up". Switch "Swings and ticks". Legend "Damage", "Cooldown", "On you", "On the target", "Rage". Groups "Casts", "Swings and ticks", "On you", "On the target". Axis "Time", "Pull". Execute wash "EXECUTE". Resource "Rage" / "peak {100} of {100}". Readout "{2:50.0}", "Rage {7} · On you: {Recklessness}" (with nothing: "On you: nothing extra"), "{3.8} s ago". Tooltip "{2:46.2} · rage after {0}", "Next {Execute} {4.0} s later", "Last {Execute} of the fight", "{Crit} · {1,284} damage". Cast list "Casts: {86} of {86}. Swings and Deep Wounds ticks are on the lanes above." (the periodic name is the run's own), "Casts in view: {12} of {86}", "Show all {n}", "Show fewer", "{Heroic Strike} queued", "Rage {100}". Unpin button `aria-label` "Close the readout".

**Gear:** label "THIS RUN'S INPUT", title "Gear", meta "{17} of 17 slots · {11} enchants", "No set bonus is active: no two pieces share a set.", "Set: {name} {2} of {5}", "Upgrades for this character on Best in slot", "This run had no gear on file."

**Talents:** label "THIS RUN'S INPUT", title "Talents", "Read only. Hover or focus a talent for its tooltip.", phone "Read only. Tap a talent for its tooltip.", "Open in planner", §4.7's line, "Talents didn't load."

**Resource:** label "AVERAGE FIGHT", title "Rage", tiles "Gained", "Spent", "Lost at the cap" (mana: "Ended with"), lists "Gained from", "Spent on", overflow "−{1.3} lost", §4.3's line.

**Buffs:** label "AVERAGE FIGHT", title "Buffs and debuffs", meta "Uptime, then times applied", "On you", "On the target", "{9} that come and go", "Up all fight", "Up all fight, from the raid", "{68}×", §4.3's line.

**Rotation:** label "{FURY}, AT {60}", title "Rotation", meta "{12} lines · casts are per fight" (+ " · today's rotation"), "Before the pull", "Then, from the top at every decision", "Other cooldowns as they come up", "This run: {Blood Fury}.", "{19} casts" / "1 cast", "Not cast in this run", foot "The {Fury} rotation the engine follows. A line with no cast in this run is dimmed.", "Read the {Fury} guide".

**Settings:** label "THIS RUN'S INPUT", title "Settings", keys "Fight length", "Targets", "Execute phase", "Race", "Runs", "Ran on", "Engine"; values "3:00, ± 20% (2:24 to 3:36)", "Last 25%", "Our server" / "Your browser"; Tag "Older engine"; groups "On you", "On the target", "Consumables"; "None in this run."; "Show all {26}"; "Copy for the game", "Copy JSON", "Download JSON", "The JSON is this whole result, request included."

**Errors:** "This sim did not load", "The server did not answer. Your link is fine; try again in a moment.", "Try again", "This part of the page didn't load."

## 7. Phone, tablet, desktop, wide

| Width | Behaviour |
|---|---|
| 360 to 767 | 18px gutters. **Band:** crest 56px, h1 34px, figure 56px; buttons full width and stacked; facts in 2 columns. **Nav:** scrolls sideways, no compact figure. **Damage table:** 2 columns. **Spread:** tiles in 2 columns. **Timeline:** 44px icon column, sticky; plot scrolls sideways at 4 px/s; readout pinned under it; toolbar wraps. **Gear:** paperdoll fold. **Talents:** TreeTabs. **Resource and auras:** columns stack; bars 56px. **Rotation:** counts on the name line. **Settings:** columns stack; chips capped at 6 + Show all; actions in a 2-column grid. All hit targets at least 44px. |
| 768 to 1023 | 48px gutters. **Band:** stacked (answer under who, left-aligned). **Damage table:** full columns; the Uptime column is dropped when the table would exceed the card. **Timeline:** 200px labels, whole fight fits the width (no scroll), 30 s gridlines. Two-column cards stay two columns. |
| 1024 to 1279 | **Band:** two regions (answer right). The nav's compact figure appears. **Damage and Spread:** stacked full width; spread is histogram 7fr + tiles 5fr. |
| 1280 to 1440 | As the 1440 board. Damage 8fr and Spread 4fr side by side. |
| 1441 to 2000+ | Every region stays in the 1344px column, centred. The band's ground runs full bleed; its inner content stays at 1344. The timeline does not grow past the column. A 1-up card keeps the column width. |

**What collapses:** the share bar (phone), the label column (to icons), chips past 6, the answer region under who.

**What never collapses:**

- the DPS figure and its error;
- any ability name in the damage table;
- any lane of the timeline;
- the outcome split;
- any rotation line or note;
- the enchant line;
- any warning.

## 8. Performance and polish

- **Budgets** (`web/lighthouserc.json`, the `/sim/[a-z2-7]{12}` matrix): performance ≥ 0.90, accessibility ≥ 0.95, SEO ≥ 0.95, total blocking time ≤ 200 ms, CLS ≤ 0.05, all on the prerendered fixture `/sim/simfixtureab.html` at 360px mobile, 4x CPU.
- **Island size** (`scripts/check-island-size.mjs`): `dist/sim-island.js` ≤ 90 KB gzipped. The boot chunk gains only the parts above the fold:
  - `ResultBand`;
  - the nav;
  - `warnings.ts`;
  - `AbilityTable`;
  - the new `DpsDistribution`.
  Each lane records the build's printed gzip size before and after in its report.
- **Lazy chunks** (`createLazyComponent`, each its own file): `SimTimeline` (on approach), `GuideTreeColumns` with the talent data (on approach), `Paperdoll` (already its own CSS; on approach), the Resource and Buffs cards, `RotationCard` with `rotations.json` (already lazy), and the tooltip panels (already lazy via the host). No chart library is added: the timeline is one SVG per group, drawn from arrays.
- **Main thread:** the timeline builds its lanes in one pass over `sample` (404 events here; cap 5,000) and renders one SVG with at most `lanes x events` rects. Ticks under 1px apart in a lane merge into one 3px tick. TBT for the timeline mount must stay under 50 ms at 4x CPU; measure it in the lane.
- **No layout shift:**
  - every card's skeleton is its ready height (§4.2, §3.7);
  - fonts are the size-adjusted fallbacks;
  - icons have fixed boxes with the skeleton fill until they load, never a missing-icon glyph;
  - the compact nav figure is absolutely positioned, so it moves nothing.
- **Reduced motion:** no smooth scroll, no `.reveal`, no shimmer.

## 9. Acceptance captures (the implementer attaches all)

Widths 390, 1024, 1440, 1920 and 2000, full page, at 2x, against the real-run fixture (§11.3) unless a state says otherwise.

| # | Capture | Widths |
|---|---|---|
| A1 | Default, full page | 390, 1024, 1440, 1920, 2000 |
| A2 | Side by side with `result-1440.png` and `result-390.png` (the owner's rule: same scale, full page) | 1440, 390 |
| A3 | First screen only (viewport 1440 x 900, 390 x 844), showing §1's answer | 1440, 390 |
| A4 | Nav stuck mid-page with the compact figure, current link on Timeline | 1440 |
| A5 | Worth knowing with W1, W3 and W4 forced (fixture variant) | 1440, 390 |
| A6 | Damage table: hover on a row, focus ring on a sort header, sorted by Crit ascending, spell tooltip open on Execute | 1440, 390 (tooltip by tap) |
| A7 | Timeline: Whole fight; Close up with the brush window; Swings and ticks on; crosshair readout; pinned readout; tick tooltip; keyboard focus on a tick | 1440, 390 |
| A8 | Timeline scrolled to the end on a phone (sticky labels and axis visible) | 390 |
| A9 | Histogram bar tooltip | 1440 |
| A10 | Gear: an item tooltip open (left and right column); a set bonus active (fixture variant) | 1440, 390 |
| A11 | Talents: three trees; phone tabs; a talent tooltip; §4.7 fallback (this run's own string) | 1440, 390 |
| A12 | Rotation with a dimmed line; Settings with Show all expanded on a phone | 1440, 390 |
| A13 | Copy JSON copied state; Copy failed state | 1440 |
| A14 | Loading: the prerendered skeleton before hydration, and each lazy card's skeleton (throttled) | 390, 1440 |
| A15 | Older result: `web/src/fixtures/sim/result.json` (no new fields), full page | 390, 1440 |
| A16 | Load failed (API 500) and a lazy chunk failure | 390, 1440 |
| A17 | Stale engine variant | 1440 |
| A18 | Lighthouse report for `/sim/simfixtureab.html` and the island size line from the build | — |

## 10. Accuracy checks (tenet 8; recorded in each lane report)

1. Every figure in the band, table, tiles and lists matches a hand calculation from the fixture JSON. The fixture test pins Execute 124.0 DPS, 30.8%, 19 casts, 1,172 average, 47.4% crit, and Main hand miss 24.6%, dodge 6.6%, glance 39.3%. Attempts are `hits + misses` excluding glancing, because the adapter counts a glance inside `hits`; see `sim/adapter/adapter.go` `ability()`.
2. APM 25.1 from 75 presses over 179.5 s.
3. Every action key in the fixture resolves to a name and an icon. A test fails on "An unnamed spell" or a missing icon for any key the fixture carries.
4. W3 fires for Recklessness and nothing else on the fixture.
5. The histogram counts sum to `iterations_run`.
6. Resource sources sum to `gained` and `spent` within 0.5.

## 11. Data

### 11.1 Have today (verified on the real run)

- `dps` {mean, stddev, error, min, max}; `iterations_run`; `request` (gear with enchants, talents, buffs, consumes, encounter); `engine_version`; `lane`.
- `summary.damage_done[0].abilities[]` {spell_id, name key, school, total, hits (including crits and glances), crits, ticks, misses map}. These are per-iteration averages.
- `summary.casts[]`; `summary.auras[]` (player BUFF tracks: applications, uptime_ms); `summary.resources[0]` {gained, spent, wasted}.
- `sample[]` {at_ms, action, target, resources}: the median-DPS run's events, white swings and periodic ticks included.

### 11.2 New fields (engine and adapter lane)

All are additive on `SimResult` (`sim/api/envelope.go`, mirrored in `web/src/lib/sim/types.ts`). All are optional (`omitempty`). Stored results need no migration (`sims.result` is jsonb).

```jsonc
// SimResult additions
{
  // Engine DistributionMetrics.hist, rounded-DPS bins, contiguous from `start`.
  // counts sum to iterations_run. Absent for weights and bulk kinds.
  "dps_histogram": { "start": 311, "bin_dps": 1, "counts": [1, 0, 0, 2, 1, 3 /* ... */] },

  // Per resource, the average per fight, one decimal. Keys are action keys
  // (same grammar as summary rows), so names and icons resolve the same way.
  "resource_sources": [
    {
      "resource": "rage",
      "gained": [
        { "action": "other:attack/1", "amount": 512.3, "wasted": 3.1 },
        { "action": "spell:2687", "amount": 80.0, "wasted": 0 }
      ],
      "spent": [
        { "action": "spell:20662", "amount": 408.4 },
        { "action": "spell:11567", "amount": 336.0 }
      ]
    }
  ],

  // The sample run's own length; the timeline axis ends here.
  "sample_duration_ms": 214545,

  // Aura on/off segments inside the sample run, player and target.
  "sample_auras": [
    { "action": "spell:12328", "on": "player", "segments": [[0, 30000], [183000, 213000]] },
    { "action": "spell:12867", "on": "target", "segments": [[1810, 16400]], "stacks": [[1810, 1]] },
    { "action": "spell:7386",  "on": "target", "segments": [[3000, 214545]], "stacks": [[3000, 1], [4500, 2]] }
  ],

  // Saved time, a sibling like `title` on GET /v1/sims/{id} (api/internal/sims/handler.go GetOutput).
  "created_at": "2026-09-18T21:04:11Z"
}
```

```jsonc
// SampleCast additions (each sample[] row), optional
{ "at_ms": 166250, "action": "spell:20662", "target": "Target 1", "resources": { "rage": 0 },
  "outcome": "crit",   // hit | crit | miss | dodge | parry | glance | block | tick | tick_crit
  "amount": 1284 }     // damage of this event, whole number; absent for a non-damage cast
```

**Target debuff tracks:** no new shape. The adapter adds `summary.auras[]` rows with `"type": "DEBUFF"`, `target_guid` `"target-1"`, `target_name` `"Target 1"`, applications, `uptime_ms` and `max_stacks`, from the engine's target unit metrics (`encounter_metrics.targets[].auras`). The web's Damage by ability Uptime column reads the DEBUFF track whose key matches a periodic row.

**Fixtures for the web lanes:**

- (a) `web/src/fixtures/sim/result-full.json`. Lane E writes it first, by hand, from the real run with every new field filled in plausibly, within the lane's first step, so lanes A to C build against it. Then it is regenerated from a real engine run before merge.
- (b) `web/src/fixtures/sim/result.json` stays the older shape. It is the fallback fixture for A15.

### 11.3 Request addition

`request.data_build` (string, e.g. "1.60.1.70790"): the talent and item build the request was made against. The web sets it when it builds a request. `talentsFit` decodes against that build's trees when it is not the active one (open question 1).

### 11.4 Data files (data lane)

- `data/curated/extra-spell-icons.json`: add the client SpellIcon for 12319 Flurry, 20572 Blood Fury, 20007 Holy Strength, 12867 Deep Wounds, 462891 Whirlwind (off-hand hit), 11585 Overpower, 9910 Thorns, 20906 Trueshot Aura, 20218 Sanctity Aura and 10301 Retribution Aura. Each is verified against the client's SpellIcon table, not copied from this mock.
- `data/curated/engine-spell-names.json`: 462891 "Whirlwind (off hand)", cited to the engine file that casts it.
- `simbuffs.json`: a `target: true` flag on every entry that lands on the target.
- A test: every action key in `result-full.json` resolves to a name and an icon.

## 12. Lane split

Three web lanes plus the engine and adapter lane. One integrator lands them (the batch-landings rule).

| Lane | Owns | Files | Starts from |
|---|---|---|---|
| **E: engine, adapter, data** | §11 | fork: `DistributionMetrics.hist` kept, sample aura segments, sample outcome and amount, per-action resource metrics; `sim/adapter` (Summarize, Sample, new envelope fields, DEBUFF tracks); `sim/api/envelope.go`; `api/internal/sims/handler.go` `created_at`; `web/src/lib/sim/types.ts`; `web/src/fixtures/sim/result-full.json`; §11.4 data files and test | now |
| **A: frame and damage** | §3.2 to §3.6, §3.12, §4.2 to §4.9, §8 skeleton and reservation | `SavedSim.svelte` (the only lane that edits it; it mounts B's and C's components), new `ResultBand.svelte`, nav, `lib/sim/warnings.ts`, `report/AbilityTable.svelte` (extracted from `ActorRow`), `DpsDistribution.svelte`, `DetailsCard.svelte`, `lib/sim/skeleton.ts`, `pages/sim/[id].astro`, `lib/sim/copy.ts` | E's hand fixture |
| **B: timeline** | §3.7 | new `SimTimeline.svelte` plus `lib/sim/timeline.ts` (lanes, segments, keyboard model), `SampleLog.svelte` rework; reuses `chart/*` | E's hand fixture |
| **C: inputs, resource, buffs, rotation** | §3.8 to §3.11, §4.7 | `Paperdoll.svelte` enchant line (and `lib/gear/paperdoll.ts`), a talents card over `GuideTreeColumns` plus `lib/sim/talents-fit.ts`, a resource card, `AuraTable.svelte` `variant="bars"`, `RotationCard.svelte` | E's hand fixture |

`SimResults.svelte` (the tab strip) stays only for the live `/sim` page until that page adopts this layout. That adoption is a follow-up spec, not this one.

Each lane runs only its touched tests. The integrator runs the full web vitest, the full Playwright suite and the lhci shard, because this rebuilds a page (the shared-layout rule).

## 13. Open questions

1. **This run's talent string does not fit build 70790.** Fury has 17 digits for 18 talents; Arms digit 5 lands on Improved Tactical Mastery (max 3). Either the engine's talent order is not the client's tab order, or the request predates a tree change. Lane E answers which, from the engine's talent table. Without that answer, every saved sim's tree drawing is suspect. Until it is answered the page shows §4.7.
2. **Main hand misses 24.6% against the off hand's 14.8%** in this run. A dual wielder's main and off hand share the same miss chance on white swings. Is the engine charging the Heroic Strike queue's resolution to the white row, or is this real? The sim accuracy loop checks this before the fixture pins the figures.
3. **The stored result does not record the rotation it ran.** The page shows today's rotation and says so only when the engine is stale. Should the envelope carry the APL's content hash (`rotation_sha`) so the page can say "today's rotation" exactly when it differs?
4. **Hand tags on Holy Strength** (`spell:20007/1`, `/2`) are read as main hand and off hand. This needs confirming from the adapter's tag allocation before the copy ships.
5. **"Overpower" as a buff track** (`spell:11585`, 16 applications, 35.7% uptime) is the engine's "Overpower is usable" window. Should it read "Overpower ready" in the Buffs card? That is a name the client does not use.
6. **Saved date:** does the API add `created_at` (§11.2), or does the page leave the fact out for good?
