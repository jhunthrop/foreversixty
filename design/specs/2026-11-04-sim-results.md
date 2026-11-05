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
| `timeline-1440.png`, `timeline-390.png` | The timeline in Close up (2:05 to 2:35, the Execute phase opening), Swings and ticks on, the crosshair at 2:25.0 with its readout, an Execute crit's tooltip (desktop), lane names inside the lanes (phone) |
| `states-1440.png` | Loading skeletons at ready height, an older saved run's fallbacks, an older run whose talents do not fit today's trees and the capstone warning it raises, load failed |

**Data rules for the mocks.**

- **The run** is `design/mocks/data/simview/fury-result.json`, a real forever-sim run of the published warrior-fury level 60 raid BiS entry (`fury-request.json`):
  - troll, the guide talents `30323013002000000-051531025120100231-200000000000000000` (17/32/2);
  - the raid preset, Windfury Totem included, and 8 consumables;
  - 3:00 ± 20% single target, 10,000 runs, 904.4 DPS ± 0.5, engine `d5c708834`;
  - Bloodthirst 22 casts per fight on average;
  - no enchants, because the ranker sims without them. The mock shows none, honestly.
- **Names** resolve in the order `sync-data.mjs` folds them: `spells.json`, then `data/curated/engine-spell-names.json`, then the class talent file. Items come from `items.json`.
- **Icons** come from `spellicons.json`, then the talent file, then `simbuffs.json` for a buff or consumable of the same name.
  - `data.py` `ICON_GAPS` lists the ids none of those carry an icon for, each with the client SpellIcon or the simbuffs icon of the same buff.
  - Two Forever-new ids have no icon anywhere in the repo: 1249110 Chthonic Power and 21919 Thrash. They draw the empty skeleton box, never an invented icon.
  - §11.4 puts all of these on the data lane.
- **Illustrative regions** draw data the engine does not return yet (§11.2). Each carries a dashed `ILLUSTRATIVE` tag on the board and nowhere in the build. Their values come from the web lanes' hand fixture `web/src/fixtures/sim/result-full.json`, so the mock and the lanes agree, and the fixture's values are derived from this run. They are:
  - the DPS spread bins;
  - the rage sources (they sum to the run's 2,127 gained and 2,116 spent);
  - the target debuff tracks and the Deep Wounds uptime;
  - the timeline's On you and On the target bars, and the hit and crit shapes of the ticks;
  - the character stats (§3.8). Hit and crit are read off this run's own hit table; attack power and haste show a dash until the engine reports them;
  - the Saved date.

## 1. The question and the answer in view

**The question:** "How much damage does this character do, and what produced it?"

**The answer that must be visible without scrolling:**

| Viewport | In the first screen |
|---|---|
| 1280 x 720 and up | The band: crest, spec, the DPS figure with its error, the facts row. The section nav. The top of Damage by ability (at least its first three rows at 1440 x 900). |
| 360 to 390 x 640 | Crest, spec, the DPS figure, "± 0.5 · 0.05% error" and "Open in the sim". |

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
  - The `h1` (`heading-page`, 44/34px) in the class colour. It reads the first of these that applies:
    - `result.title`, if set;
    - the character's name, if set and not an engine placeholder. `NAMELESS_REQUEST_NAME` ("Sim") and the ranker's "breakdown" are placeholders; the set lives in `lib/sim/character.ts`;
    - the spec label ("Fury Warrior").
  - The descriptor at 13/22px muted:
    - race, spec and class in `--text` 600, then "Level 60", then the talent split: "Troll Fury Warrior · Level 60 · 17/32/2 talents";
    - the split is computed once, by `talentSplit(request.character.talents)` (the digit sum of each tree). The Talents card's meta reads the same value. No other split appears on the page;
    - each piece appears only when known;
    - separators are drawn pseudo-elements, as in the profile header, so no line starts with "·".
  - The next-step line at 14px `--text`:
    - `afterSimSentence(topUpgrade)` when a Top Gear upgrade is on record;
    - otherwise "Next: find the upgrade that adds the most." and an `ActionLink` "Run Top Gear" to `/sim/gear`.

    This is tenet 14's "the next step that improves it".
- **The answer** (right from 1024px, under Who below that):
  - The figure: heading face, 800 weight, 64px (56px phone), the gold gradient clipped to text, `tabular-nums`, one decimal. "DPS" follows at 18px heading face, `--text-muted`, letter-spacing .04em.
  - Under it at 15px: "± 0.5 · 0.05% error". The margin is `--text` 600; the rest is muted.
  - Then the actions, right-aligned with a 14px top gap. Both are `button()` md, 44px. Below 768px they stack full width, primary first.
    - The **primary** button "Open in the sim" (play glyph).
    - The **secondary** button "Copy link" (link glyph).
- **"Open in the sim" destination:** `/sim` with this run's setup loaded and editable, ready to press Sim. The setup is:
  - race, level and talents;
  - gear with enchants and suffixes;
  - buffs, consumables and the fight (`encounter`, preset).

  It reuses the existing `onrerun` handoff, which already rebuilds a store from the stored request. The page never runs the sim on the visitor's behalf. On `/sim` the run control is focused, so Enter starts it.
- **Facts row** (full width under both, 24px top gap, a 1px `rgba(229,185,85,.14)` top rule):
  - Each fact is a `quiet-label` (13px muted) over a 15px 600 `--strong` `numeric` value. Facts are divided by 1px `--border-soft` verticals with 24px padding.
  - Facts, in order: Fight, Targets, Buffs, Consumables, Range, Runs, Saved.
  - Below 768px the facts form a two-column grid with 1px row rules.
  - "Saved" is the page's one date stamp. There are no pills and no engine sha in the band; the engine lives in Settings.
- Values:
  - **Fight:** "3:00 ± 20%" from `encounter.duration_sec` and `variation`.
  - **Targets:** "1".
  - **Buffs:** "Raid" when the request's buff list equals the raid preset, "None" when empty, else "Custom".
  - **Consumables:** the count of `character.consumes` ("8"), or "None". Consumables move a warrior's DPS more than any other input here, so they sit beside Buffs.
  - **Range:** "731 to 1,116" from `dps.min` and `dps.max`.
  - **Runs:** `iterations_run`.
  - **Saved:** the sim's `created_at` (new, §11.2) as "4 Nov 2026". Until the field exists, the fact is left out. `source.captured_at` is the character's capture time, not the save time, so it is never labelled "Saved".
- Actions per minute moves to Settings (§3.12).

### 3.4 Section nav

Reference: the report's underline tabs, which players already use on `/reports/<id>`.

- `TabStrip.svelte` in link mode (no `onSelect`): a `nav` of `#anchor` links. The strip is 48px tall and sticky at `top: 0`. Its ground is `rgba(7,9,13,.94)` with a 1px `--border-soft` bottom rule, at `z-index` 5.
- Labels: Damage, Timeline, Gear, Talents, the resource name (Rage, Energy or Mana from the player's power type), Buffs, Rotation, Settings. A section that does not render (no gear on file) drops its link.
- **Current section:** an `IntersectionObserver` marks the link whose card crosses 30% of the viewport from the top. It gets `aria-current="location"`, `--strong` ink and the 2px gold underline. Clicking a link scrolls smoothly (instantly under reduced motion) and moves focus to the card's `h2` (`tabindex="-1"`).
- **Compact figure:** from 1024px, once the band's figure has scrolled out, the strip's right end shows the figure: "904.4" in the heading face at 18px gold, then "DPS" 13px muted. It fades in over 120ms. It is `aria-hidden`, since the `h1` region already states it.
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
| W1 stale | `isStale(result.engine_version)` | "This run used an older engine. A run today may give a different number." + ActionLink "Open in the sim" |
| W2 capstone never pressed | the spec's capstone spell (`capstone_spell` in `data/curated/specs.json`, §11.4; Bloodthirst for Fury, Mortal Strike for Arms) has 0 casts in `summary.casts`. This does **not** depend on `talentsFit` (§4.7): it fires for a talent string that does not decode too, which is exactly when it matters. | "{Bloodthirst} was never pressed in this run. Check the talents in Open in the sim." + ActionLink "Open in the sim" |
| W3 missing talent | the talents fit (§4.7), W2 did not fire, and a rotation line casts a talent ability other than the capstone with 0 points in this build | "{Reckless Swings} is in the {Fury} rotation, but these talents do not take it." + ActionLink "Open in planner" |
| W4 long cooldown | a cast spell has `cooldown_ms` greater than the longest possible fight (`duration_sec x (1 + variation)`) and was cast at least once | "{Recklessness} has a {30 minute} cooldown, longer than this fight, so the result counts one use every pull. In a raid it is back only every {30 minutes}." |
| W5 resource at cap | `wasted / gained ≥ 5%` for rage or energy | "{n} {rage} ({p}%) was lost at the cap." Rage adds "Queue Heroic Strike sooner so swings do not land at full rage." Energy adds "Spend before the bar fills." |
| W6 wide error | `dps.error / dps.mean > 0.5%` | "This run's error is ± {x} DPS. A difference smaller than that between two runs is noise." + ActionLink "Open in the sim" |

This run fires W4 only, for Recklessness: a 30 minute cooldown against a longest fight of 3:36. The others do not fire:

- W2: Bloodthirst has 22 casts.
- W5: 57 of 2,127 rage is 2.7%.
- W6: ± 0.5 is 0.05%.

The previous fixture (0 Bloodthirst casts, talents not fitting) fires W2. `states-1440.png` draws it.

### 3.6 Damage by ability + DPS spread

Grid from 1280px: `minmax(0,8fr) minmax(0,4fr)`, a 16px gap, top-aligned. From 1024 to 1279px both cards are full width, and the spread card puts its tiles to the right of its histogram (7fr and 5fr). Below 1024px both cards are full width and stacked.

**Damage by ability.** A new `AbilityTable.svelte` in `components/report/`, extracted from `ActorRow.svelte`'s expanded ability table so the logs report gets the same rows; `ActorRow` mounts it. Reference: Details! rows plus WCL's Avg Hit, Crit % and outcome split.

**The adapter's counts:** `hits` already includes crits and glances; `misses` holds MISS, DODGE, PARRY, BLOCK and GLANCING; a glance is in both (`sim/adapter/adapter.go` `ability()`). So:

- **attempts** = `hits + MISS + DODGE + PARRY` (BLOCK and GLANCING are already in `hits`);
- **landed** = `hits`.

The header row and columns:

- CardHead: label "THIS RUN", title "Damage by ability", meta "904.4 DPS over 2:59 on average".
- A table (`table`, real `th` with `scope="col"`). The header row is 40px with `quiet-label` headers; numeric columns are right-aligned.
- **Columns at 1024px and up:**
  - **Ability** (1fr, min 240): a 28px icon (6px radius, 1px black border, 1px `rgba(229,185,85,.18)` outer ring), then the name at 15px 600 `--strong` with " Rank 5" in 13px muted.
  - **DPS:** 15px 600 `--strong`, one decimal.
  - **Share:** a 120 by 8px bar (track `--border-soft`, fill `--gold-deep` to `--gold` gradient) scaled to the top row, then the percentage at one decimal, 48px wide.
  - **Casts:** a whole number. A white row counts swings (attempts). A periodic row shows a muted "–".
  - **Avg hit:**
    - normally damage divided by **landed** hits (Execute: 34,337 / 23 = 1,493);
    - a periodic row is damage per tick (Deep Wounds: 8,315 / 59 = 141).
    - Never per attempt: a miss is not a hit, and the old "Average" (per attempt) read like a weak weapon.
  - **Crit:**
    - normally crits divided by **landed** hits, a glance counted as landed. This is what WCL's Crit % shows. Main hand: 20 / 47 = 42.6%;
    - a periodic row is crits per tick ("tick crit").
    - **Deviation from the round 2 ruling, for the coordinator:** the ruling asked for crits over landed *non-glancing* hits. On this run that reads Main hand 90.9%, Off hand 73.0% and Extra attacks 100.0%. Each is true (a dual wielder near the crit cap lands almost every non-glancing white swing as a crit), but each reads as a bug, and WCL does not count that way. Yellow rows never glance, so the two definitions differ only on white rows. Switching back is one line in `AbilityTable`'s view model.
  - **Biggest hit:** the largest single landed hit, from the ability's `max`. The column renders only when at least one row has `max > 0`. The engine returns 0 today, so the column is absent on this run (lane E, §11.2).
  - **Uptime:** for a periodic row, the matching target DEBUFF track's uptime (Deep Wounds 100.0%). Otherwise a muted "–". The column is dropped when no row has a value.
- **The sub-line** under each name, at 13/18px muted with the figures in `--text`. Each "Label value" pair is one `white-space: nowrap` unit, so a wrap never splits "Glance" from "40.2%".
  - The outcome split, as a share of **attempts**, in the order Miss, Dodge, Parry, Glance, Block: "Miss 8.5% · Dodge 6.1% · Glance 40.2%". A zero outcome is left out.
  - A periodic row reads "59 ticks · tick crit 37.3%".
  - The main-hand row adds "32 swings became Heroic Strike": the Heroic Strike landing count (`spell:11567`), the swings it replaced.
  - No row carries an explanatory note about a figure the accuracy loop has not verified. The off-hand miss gap stays unexplained on the page (§13).
- **White swing rows:**
  - `other:attack/1` is "Main hand" and `other:attack/2` is "Off hand". The icon is the equipped weapon's icon, as the in-game Attack button shows it. After the name comes " · " and the weapon's name in its quality text colour, with the shared item tooltip on hover or focus.
  - **Extra attacks:** `other:attack/3` (tag 3, the engine's extra-attack hand, confirmed by lane E against the adapter's tag allocation) is its own row, "Extra attacks", never folded into Main hand. It uses the main-hand weapon's icon and name, and its sub-line ends with the sources: "From Windfury Totem".
    - With `extra_attack_sources` (§11.2) the sources carry counts: "From Hand of Justice 6 · Windfury Totem 21".
    - Without it, the sources are named from the request, with no counts. The source list is curated in `lib/sim/extra-attacks.ts`: the buff `windfury_totem`, plus worn items whose client effect is an extra attack (Hand of Justice 11815, Ironfoe 11684, Thrash Blade 17705, Flurry Axe 871, and any Forever item lane E adds).
- **Pet rows** (hunter, warlock): indented 24px under a "Pet" sub-heading row with the pet's name. Each row reads as above.
- **Sorting:** DPS descending by default. The headers DPS, Share, Casts, Avg hit, Crit and Biggest hit are `button`s inside the `th` with `aria-sort`. A first press sorts descending, a second press ascending. The sorted header is `--strong` with a gold ▾ or ▴.
- **Row states:**
  - Hover (fine pointer) gives the row `--row-hover`.
  - The icon and name are one focusable trigger (`data-spell-key`, 44px tall on phone). It opens the shared spell tooltip (`GameTooltip.svelte` through `BisTooltipHost`'s game-tip mode) on hover, focus or tap. Escape closes it, as everywhere else.
- **Total row:** muted 13px: "Total, 9 sources", the summed DPS, "100%", "162,715 damage a fight".
- **Phone (below 768px):**
  - Two columns, Ability and DPS. The DPS cell stacks the DPS over the share percentage, muted.
  - The sub-line gains a second line: "24 casts · 1,493 avg hit · crit 52.2%". Every pair is nowrap.
  - The share bar is not shown on phone; the percentage carries it.

**DPS spread.** `DpsDistribution.svelte`, rewritten to draw the engine's real bins (§11.2). The current normal-curve drawing is removed, because a drawn shape the run did not produce breaks tenet 8.

- CardHead: label "10,000 RUNS", title "DPS spread", no meta.
- **Histogram:** 168px tall, at most 40 bars with 2px gaps. The engine's 1 DPS bins are grouped client-side into `ceil(bins / 40)` DPS buckets.
  - Bars inside the 16th to 84th percentile are `rgba(229,185,85,.7)`; the rest are `--border`.
  - A 2px `--gold-light` mean line with a 13px label above it ("904.4").
  - Under the bars, the axis: min, "DPS per run", max.
- **Readout:** hovering or focusing a bar shows `ChartTooltip`: "900 to 910 DPS: 712 runs". The histogram is one `role="img"` with an `aria-label`, plus arrow-key stepping when focused.
- **Caption** at 13px muted: "Two runs in three land between 851 and 961 DPS." The figures are the 16th and 84th percentiles of the real bins. Nothing more: the sim does not know how real pulls spread (movement, buff timing and latency are not in it).
- **Six StatTiles** in three columns (two on phone): Mean, Error, Spread (the standard deviation, as "± 49.5"), Lowest, Highest, Runs.

### 3.7 Timeline (new `SimTimeline.svelte` in `components/sim/`, lazy chunk)

Reference: WCL's casts timeline and the competitor's lanes, built on the report's own chart parts: `lib/chart/time-axis.ts` (`timeTicks`, `clockLabel`), `SyncedCrosshair.svelte`, `OverviewBrush.svelte`, `ChartTooltip.svelte`, `PresetChips.svelte`, `Switch.svelte` and `ChartSrTable.svelte`.

`TimelinesView.svelte` stays the logs view. It draws one lane per player; this view draws one lane per ability. They share axis, crosshair and brush code, not a component.

**Source:** `result.sample`, the median-DPS run's events, plus `sample_auras` and `sample_duration_ms` when present (§11.2). The axis runs from the first event (pre-pull, -2.5 s here) to `sample_duration_ms`. Without that field, the axis ends at the last event plus 1 s. CardHead meta: "The median run · 2:53 long" (`m:ss` of the sample's own length, which differs from the 3:00 average fight because fights vary ± 20%).

**Toolbar** (min height 56px, 10px by 20px padding, a 1px top rule, wraps):

1. **Zoom:** `PresetChips` with two chips, "Whole fight" and "Close up". Close up opens a 30 s window centred on the crosshair, or at the pull when there is none.
2. **Swings and ticks:** a `Switch`, off by default. On, it adds the group "Swings and ticks" (Main hand, Off hand and each periodic effect's ticks).
3. **Legend** (right-aligned, 13px muted): Damage (3 by 12px `--color-school-<school>` tick), Crit (4 by 16px `--gold-light`, only with outcome data), Missed (3 by 12px `--text-muted` outline, only with outcome data), Cooldown (3 by 12px `--text` tick), On you (16 by 8px `--text` at 35%), On the target (16 by 8px `--gold-deep` at 78%), Rage (16 by 2px `--gold` line). The legend's resource word follows the player.

**Plot** (desktop): a 200px label column and the plot filling the rest.

- **Axis row**, 30px: 13px 600 muted labels every 15 s (whole fight, 1280px and up), 30 s (below 900px of plot), 5 s (Close up). There is a dashed `--border-warm` line at 0:00 labelled "Pull" in gold, and 1px `--border-soft` gridlines through every lane.
- **Groups**, each with a 30px heading row (13px 600 muted): Casts; Swings and ticks (when on); On you; On the target. A group with no lanes does not draw.
- **Lanes**, 24px each. The label is an 18px icon and the 14px name. A name never truncates: the label column is 200px, and a longer name wraps the column to 240px for that page.
- **Lane order:** Casts by first use in the sample. On you by first application. On the target by first application.
- **Lane membership:**
  - Casts: one lane per base spell id. A tagged variant folds into its base: `spell:11567/1`, the Heroic Strike queue press, draws as a hollow tick (1px outline) on the Heroic Strike lane in Close up only, and the landing `spell:11567` draws as a filled tick at every zoom. In Whole fight only landings draw, so the lane never reads as twice the presses.
  - Swings and ticks: `other:attack/1` Main hand, `other:attack/2` Off hand, `other:attack/3` Extra attacks, and any key whose summary row has `ticks > 0` ("Deep Wounds ticks").
  - **Overpower window:** when the rotation casts Overpower (`summary.casts` has `spell:11585` above 0), one 16px lane under Casts named "Overpower window" draws the `spell:11585` aura's segments as a 6px `--text` 35% bar. It is absent for Fury at 60, which never casts it.
- **Ticks:** 3 by 14px, 1px radius, centred in the lane.
  - Fill: a damaging cast is its school token (`--color-school-physical` for a warrior). A cast with no damage row (Bloodrage, Death Wish) is `--text`.
  - Swings and periodic ticks are 2 by 10px at 80%.
  - **With outcome data** (§11.2): a crit is 4 by 18px in `--gold-light`; a miss, dodge or parry is a 1px `--text-muted` outline; a glance is the fill at 60%.
- **Bars:** 10px tall, 3px radius, centred. On you is `--text` at 35% (the report's own aura bar). On the target is `--gold-deep` at 78%. A stacking aura prints its stack count at 11px inside the bar when the bar is at least 18px wide.
- **Execute phase:** from the first Execute in the sample to the end, a `rgba(214,110,40,.07)` wash over every lane, with "EXECUTE PHASE" at 11px 700, .14em tracking, `--ember`, 6px from its left edge. It draws only for a class with an execute spell in its rotation (warrior Execute; later hunter Kill Command if added). It is labelled from the first use, so it never claims a health threshold.
- **Resource lane:** 72px. The label is "Rage" (14px 600) over "peak 100 of 100" (13px muted). The plot is a step line, 1.5px `--gold`, over a `rgba(229,185,85,.12)` area, with a dashed `--border-warm` line at the cap. Values are the sample's `resources` after each event, held until the next event.

**Brush:** `OverviewBrush` under the plot, 44px. It draws the whole fight's resource line at 1px `#5d5849`. In Close up it draws the window: a `rgba(229,185,85,.10)` fill, a 1px gold border and two 6 by 20px gold handles. Dragging the window pans it; dragging a handle resizes it (10 s minimum). Arrow keys pan 5 s; Shift with an arrow key resizes. Whole fight hides the window.

**Crosshair and readout:**

- Moving a fine pointer over the plot draws a 1px `--gold-hover` vertical line. A 52 by 20px gold time chip sits on the axis ("2:25").
- The readout is a 280px `--bg-raised` panel, 12px radius, `--border`, the tooltip shadow. It sits 12px right of the line, or flips left near the right edge. It holds:
  - the time ("2:25.0", heading face 18px);
  - "Rage 0 · On you: Death Wish, Recklessness, Berserker Rage" at 13px muted with the figure strong (Flurry is left out of this line: it is up most of the fight);
  - a rule;
  - the five most recent casts, each a 20px icon, the name, and "0.5 s ago" right-aligned, muted.
- Swings and ticks are not listed in the readout.
- Click or Enter pins the readout. A pinned readout gains a close button (44px hit, × glyph). Escape or a second click unpins it.

**Tick tooltip:** hover or focus on a tick opens `ChartTooltip`, 250px wide:

- line 1: the 24px icon, the name and the rank;
- line 2: "2:13.0 · rage after 12";
- line 3, muted: "Next Execute 1.5 s later", or "Last Execute of the fight";
- with outcome data, a line 4: "Crit · 2,275 damage".

Without outcome data, line 4 is absent.

**Cast list** (under the brush): the old One iteration table, rebuilt on `SampleLog.svelte`.

- A 52px foot row: the quiet label "Casts: 120 of 120. Swings and Deep Wounds ticks are on the lanes above." In Close up it reads "Casts in view: 19 of 120". A secondary sm button follows: "Show all 120" or "Show all 19".
- Rows are 40px: time ("2:41.2", numeric muted, 72px), a 24px icon, the name ("Heroic Strike queued" for the queue press), and "Rage 100" right-aligned, muted.
- Collapsed, the list shows 5 rows: in Close up the five nearest the window's start; in Whole fight the first five. Expanded, it shows every row in view. "Show fewer" collapses it again.
- The list is the timeline's accessible table: it is what a screen reader reads (§4.3).

**Keyboard:**

- The plot is one tab stop: `role="application"` with `aria-roledescription="timeline"` and an `aria-label` of "Fight timeline, 2:53 long, 120 casts. Arrow keys move between casts.".
- Left and Right move the crosshair to the previous or next cast across visible lanes. Up and Down move to the nearest cast in the lane above or below. Home and End go to the first and last.
- Enter pins the readout. `+` and `-` switch Close up and Whole fight. Escape unpins.
- Each move announces the readout's first two lines through a polite live region.

**Touch:**

- Tap a tick for its tooltip. Tap the empty plot to place the crosshair; this pins the readout.
- The plot never captures vertical swipes (`touch-action: pan-x pan-y`).

**Phone (below 768px):**

- The label column is 44px: the 18px icon, `aria-label` set to the name. Tapping the icon opens the spell tooltip.
- Each lane also prints its name at 11px `--text-muted` inside the plot, on a `#0d0f15` 88% backing chip (3px radius, 14px tall), `position: sticky; left: 4px`. It stays readable while the plot scrolls, so five red-orange warrior icons are never the only label.
- Group headings show as a 12px muted label over the plot's left edge, on the card ground. The resource lane shows "Rage" at 11px.
- The plot **scrolls sideways** inside the card. The label column and the axis row stay fixed: the label column is `position: sticky; left: 0`, and the axis lives inside the scroller with `position: sticky; top: 0`. A 24px fade marks the clipped right edge.
- Scale: Whole fight is 4px per second (this run's 173.9 s is 696px). Close up is 12px per second.
- The readout does not float on a phone. It is pinned **under the plot**, full width, after the brush, and updates on tap.
- The tick tooltip opens as the same pinned panel. Nothing is cut; the toolbar wraps to three lines.

**Loading placeholder:** the chunk loads with `createLazyComponent` once the card nears the viewport (`IntersectionObserver`, 600px root margin) or on the first nav click. Until then the card renders its CardHead and a `Skeleton` sized to the ready height. That height is computed in the main chunk from data it already has:

`56 (toolbar) + 30 (axis) + 30 x groups + 24 x lanes + 72 (resource) + 44 (brush) + 52 (cast foot) + 40 x 5 (cast rows)`

The lane count follows §3.7's membership rules with Swings and ticks off. The swap is the one 160ms `.reveal`. Nothing moves.

### 3.8 Gear

`Paperdoll.svelte` (spec 2026-10-31), wrapped in a Card.

- CardHead: shield glyph, label "THIS RUN'S INPUT", title "Gear", meta "17 of 17 slots · no enchants" (or "· 11 enchants").
- **One change to the shared component: an enchant line.** A slot whose gear entry has an `enchant` gains a line between the name and the slot line.
  - Style: 13/18px in `--color-rarity-uncommon` (the client's green enchant line), led by a 12px enchant glyph from `glyphs.ts` (four-point spark, 2.4px stroke). The glyph is there because an uncommon item's name is the same green.
  - **Text: the client's own enchant name**, the `SpellItemEnchantment` name the item tooltip prints. It is never the recipe name. Examples: "+8 Strength" (Lesser Arcanum of Voracity), "+7 Strength", "+9 Strength", "+4 All Stats", "+3 Agility", "Minor Speed Increase", "Crusader".
  - The data lane adds that name to `enchants.json` as `display` (§11.4). `enchants.json`'s `name` is the recipe ("Enchant Gloves - Greater Strength") and is never shown.
  - A slot with no enchant draws no line, so the published BiS entry (ranked without enchants) shows none. An imported character's enchants show on their slots.
- The paperdoll's other parts work as specced there: hover, focus and tap open the shared item tooltip, the right column opens toward the centre, and the centre shows the crest and the average item level.
- **Character stats** (new data `final_stats`, §11.2): one inset block under the stage. It has the same `#0b0e14` ground, 1px `--border-soft` and 12px radius, and the stage's 960px max width. It opens with a `quiet-label` head: "Character stats in the fight, against a level 63 boss" (the boss level is the request's `target_level`). Five cells follow, each a 13px muted label over a 15px 600 `--strong` numeric value:
  - **Attack power:** `final_stats.attack_power`, whole number.
  - **Hit:** `final_stats.melee_hit_pct`, one decimal ("8.0%"). Under it, a 220 by 6px bar to 27% (fill `--gold`), a 2px `--gold-light` tick at 8%, and the 13px muted reference "8% specials · 27% dual-wield swings". The references are for a level 63 target. A one-hander drops the 27% and its tick.
  - **Crit on the boss:** `final_stats.melee_crit_pct_vs_target`, one decimal. This is the crit chance after the level 63 suppression, not the character sheet's.
  - **Weapon skill:** `final_stats.weapon_skill.main_hand · off_hand` ("300 · 300"); one figure without an off hand.
  - **Haste:** `final_stats.melee_haste_pct`, one decimal, at the pull. Flurry, Berserking and other procs are uptimes in Buffs, not a stat.
  - Every figure is one the engine computed for this run. None is recomputed on the page.
  - Without `final_stats` (an older run) the block is not drawn; the paperdoll ends at the stage.
  - Phone: two columns, Hit full width first.
- **Foot** (52px, 1px top rule):
  - One 13px muted line about sets. With active bonuses: "Set: {name} 2 of 5 · (2) {bonus text}", one line per set, the active bonus text in `--color-rarity-uncommon` and inactive ones muted. With none: "No set bonus is active: no two pieces share a set."
  - Then, right-aligned, the ActionLink "Upgrades for this character on Best in slot" to `/bis/<class>/<spec>`. This card never shows a gain or a pick. That is /bis's job.
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
- **Three StatTiles:** Gained, Spent, and "Lost at the cap" with the share ("57 · 2.7%"). For mana the third tile is "Ended with" (mana left at the end), since mana has no cap loss.
- **Two lists** side by side (stacked on phone), "Gained from" and "Spent on", each headed by a `quiet-label` row with the list total.
  - Rows are 40px: a 24px icon, the name, a 96px bar (track `--border-soft`, fill `--gold`, scaled to the list's top row) and the amount (600 `--strong`, whole numbers; one decimal under 10).
  - Rows sort by amount, descending.
  - A gain row with `wasted > 0` adds the overflow after the amount, on the same line, at 13px in `--ember`: "−22.8 lost". The value column is 136px (112px on phone) so it never wraps.
  - Rows come from `resource_sources` (§11.2). The lists sum to the tiles; lane E's fixture test asserts it within 0.5.
- `ResourceGraphs.svelte` is not used here; the timeline's resource lane is the graph.

**Buffs and debuffs card.** `AuraTable.svelte`, given a `variant="bars"` that renders the rows below. The logs report keeps the default variant.

- CardHead: spark glyph, label "AVERAGE FIGHT", title "Buffs and debuffs", meta "Uptime, then times applied".
- **Two columns** (stacked on phone), "On you" and "On the target":
  - each column heads with a `quiet-label` row and a count ("10 that come and go");
  - rows are 40px: a 24px icon, the name, a 96px bar (On you: `--text` at 55%; On the target: `--gold-deep`), the uptime at one decimal (600 `--strong`) and the count "91×" muted;
  - rows sort by uptime, descending.
- **Up all fight:** auras at 99.9% or more leave the bar list. They gather under a 13px muted "Up all fight" label as a wrapping row of 20px icon plus 14px name. A stacking debuff adds its stacks in muted text: "Sunder Armor × 5".
- **Left out of the card:**
  - The Heroic Strike queue track (`spell:11567/1`) is not a buff. It is drawn on the timeline.
  - The **Overpower window** (`spell:11585`) is never a buff row. The client has no such aura; it lights the Overpower button. It appears only as a timeline lane, and only when the rotation casts Overpower (§3.7). Fury at 60 never does, so this run shows it nowhere.
- Two tracks of one spell with hand tags (`spell:20007/1`, `/2`, a Crusader proc on each weapon) read "Holy Strength (main hand)" and "(off hand)", once lane E confirms the tags (§13). The client shows two plain "Holy Strength" buffs; the hand is a helpful addition.
- A Forever-new aura with no icon in the build (Chthonic Power, Thrash on this run) draws the empty skeleton box beside its real name. It never borrows another spell's icon. It is on the data lane's list (§11.4).
- **Phone:** the bar is 56px wide.

### 3.11 Rotation

`RotationCard.svelte`, rebuilt from a disclosure into an always-open list (tenet 4). `RotationDisclosure` stays for the settings bar on `/sim`.

- CardHead: list glyph, label "{SPEC}, AT {LEVEL}" ("FURY, AT 60"), title "Rotation", meta "12 lines · casts are per fight".
- **"Before the pull"** sub-heading (`quiet-label`, 14px by 20px padding), then one row per pre-pull action:
  - its time, 14px 600 muted, nowrap ("-2.5 s");
  - a 28px icon;
  - the name with its rank;
  - on the right: "Cast" when the sample has that spell before 0:00, else "Not cast in this run". A row not cast is dimmed exactly like a priority line. Battle Shout here is not cast: the raid's shout is already up.
- **"Then, from the top at every decision"**, then one row per priority line. The row grid is `44px 28px 1fr 140px`, 12px gap, 12px by 20px padding, a 1px top rule. It holds:
  - the line number;
  - the icon;
  - the name and rank (15px 600), with the line's note under it in 14/21px `--text`, max 860px;
  - the average casts per fight from `summary.casts` ("24 casts", "1 cast"), right-aligned.
- The notes are the curated APL's own `notes` (`rotations.json`, the guide voice), verbatim. A line with no note shows the name only.
- **One name for the Execute stand-in.** The sim has no health bar: it runs Execute in the last `execute_ratio` of the fight, standing in for "the target under 20% health". The page and the notes call that window "the Execute phase", always, with a capital E. The data lane rewrites the curated Fury notes, and every APL with an execute line, so they say it the same way. This ships with the guide text, under the sim-changes-ship-their-guide-text rule:
  - Execute: "First, below 20%: it takes ..." becomes "First, in the Execute phase: it takes ...";
  - Heroic Strike: "... once you have Execute never from 25% down: hold that rage for it." becomes "... once you have Execute never in the Execute phase: hold that rage for it.";
  - Death Wish and Recklessness: "the execute phase" becomes "the Execute phase".

  Settings states the window once (§3.12).
- **The "Other cooldowns as they come up" line** (`autocastOtherCooldowns`) lists this run's casts that no other line names, items included: "This run: Mighty Rage Potion, Berserking, Chthonic Power.".
- **Not cast:** a line whose spell has 0 casts is drawn at 50% (icon, name and note), and its count reads "Not cast in this run" in 13px muted. W2 and W3 (§3.5) explain a missing capstone or talent.
- **Phone:** the count moves to the name line, floated right, 13px muted.
- **Foot:** "The {Fury} rotation the engine follows. A line with no cast in this run is dimmed." and the ActionLink "Read the {Fury} guide" to `/guides/<class>/<spec>`.
- **Rotation drift:** when `rotation_sha` (§11.2) differs from today's curated APL hash, the meta adds " · today's rotation". This fires exactly when the rotation changed, not only when the engine is stale. Without the field, the meta adds it only under W1.

### 3.12 Settings and actions

`DetailsCard.svelte`, rebuilt.

- CardHead: gear glyph, label "THIS RUN'S INPUT", title "Settings".
- **Left column:** key/value rows, 44px, 1px `--border-soft` rules, the key muted and the value 600 `--strong` right-aligned.
  - **Fight length:** "3:00, ± 20% (2:24 to 3:36)".
  - **Target:** "Level 63 boss" from `encounter.target_level`.
    - When the engine reports `target_stats` (§11.2), it adds the armor: "Level 63 boss · {armor} armor, {after} after debuffs".
    - With an explicit `target_armor` in the request, that figure is used and "after debuffs" still comes from the engine.
  - **Execute phase:** "Last 25% of the fight, standing in for the boss under 20% health", from `encounter.execute_ratio`. Why 25% is the default is the sim-accuracy lane's to state in `/sim/specs`, not this page's.
  - **Race:** "Troll".
  - **Runs:** "10,000".
  - **Actions per minute:** "40.4". It counts player button presses: a `summary.casts` row whose key is not `other:*` and has no `/<tag>` suffix (121 per average fight of 2:59.9).
  - **Engine:** the 7-character sha in the mono face, muted (`--text-muted`, tenet 7), with the stale `Tag` (muted, "Older engine") when W1 fires.
  - There is no "Ran on" row: where the run happened changes nothing a player does.
- **Right column:** three groups, each headed by a `quiet-label` row with a count: On you (14), On the target (8), Consumables (8).
  - **Chips** are 30px: 4px 10px padding, a 22px icon, the 13px name, a 1px `--border` and a 6px radius. Names and icons come from `simbuffs.json`.
    - A weapon imbue names its hand: "Elemental Sharpening Stone (main hand)", "Dense Sharpening Stone (off hand)".
    - An improved grade reads "Mana Spring Totem, improved". A raw key never shows.
  - **Damage first.** In On you and On the target, buffs that change this spec's damage come first. The rest follow under a 13px muted sub-label "No effect on your damage", their chips at 62% opacity.
    - The rule is data, not a hand list: a buff has no effect when every stat in its `simbuffs.json` `stats` (§11.4) has a 0 scale factor in the spec's weights (`stat-weights.json`), and it carries no armor, damage-taken or attack-speed effect.
    - For this Fury run: Arcane Brilliance, Power Word: Fortitude, Divine Spirit, Blessing of Wisdom, Moonkin Aura, Mana Spring Totem, Curse of the Elements and the three Judgements.
  - **A debuff the game suppresses** carries its reason in the chip, in muted text after the name: "Sunder Armor · replaced by Expose Armor". The rule comes from `simbuffs.json`'s `exclusive_with` (§11.4): two armor debuffs in one exclusive group where the stronger one wins. This run has no such pair. Whether the raid preset should ever carry both is the sim-accuracy loop's question, not this page's.
  - An empty group reads "None in this run.".
  - A buff's side (on you or on the target) comes from `simbuffs.json`'s `target: true` flag (§11.4); there is no hand-kept list in the web.
- **Phone:** each group shows its first 6 chips and a secondary sm button "Show all 14". This is the one disclosure on the page, allowed because the buff list is reference, not the point.
- **Foot actions** (52px, wraps; on phone a two-column grid with "Copy for the game" full width first):
  - "Copy for the game": `CopyForGame.svelte`, unchanged.
  - "Copy JSON": copies `JSON.stringify(result, null, 2)`.
  - "Download JSON": downloads `forever-sixty-sim-<sim_id>.json`.
  - The 13px muted line "The JSON is this whole result, request included."

## 4. States

### 4.1 Every interactive element

| Element | Default | Hover (fine pointer) | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Open in the sim (primary) | gold gradient | brightness 1.06 | 2px gold ring, 2px offset | `--gold-hover` ground | never | — | — | — |
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
| `dps_histogram` | The spread card draws a range strip instead of bars: an 8px `--border-soft` track from min to max, a `rgba(229,185,85,.35)` band at mean ± one spread, a 2px `--gold-light` mean tick. Axis "731 / mean 904.4 / 1,116". Caption: "This run kept the range, not the shape. The shaded part is one spread either side of the mean." The six tiles stay. |
| `sample_auras` | The On you and On the target groups are replaced by one 13px muted row inside the plot area: "Buff and debuff bars need a run from {launch date of the field} or later. The casts and rage here are from this one." and the ActionLink "Open in the sim". The Execute wash and the resource lane still draw. |
| `sample_duration_ms` | The axis ends at the last event plus 1 s. No message. |
| `outcome` on sample events | Ticks are one colour; the tooltip has no outcome line. No message. |
| `sample` entirely | The timeline card renders its CardHead and one line: "This run has no recorded fight to draw." The nav drops "Timeline". |
| `resource_sources` | The three tiles stay. The lists are replaced by "Where the {rage} came from and went needs a newer run." |
| Target debuff tracks | The On the target column shows the raid debuffs from the request's buff list as the Up all fight row, under the line "Uptime on the target needs a newer run." The Uptime column of Damage by ability is dropped. |
| `created_at` | The Saved fact is left out. |
| `final_stats` | The character stats block under the paperdoll is not drawn. No message. |
| `target_stats` | The Target row reads "Level 63 boss" with no armor. No message. |
| `extra_attack_sources` | The Extra attacks sub-line names the request's sources without counts. |
| `rotation_sha` | The rotation meta adds " · today's rotation" only under W1. |
| `max` above 0 | The Biggest hit column is not drawn. |

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

There is no planner link, and W3 is not evaluated. W2 still is, because it reads casts, not talents. The published run passes the check. The previous fixture failed it: its Fury part has 17 digits against 18 talents, and Arms digit 5 sits on a 3-rank talent (open question 1).

### 4.8 Stale engine

W1 fires (§3.5). The Settings engine row carries the "Older engine" Tag. The rotation meta adds " · today's rotation". The figure is not dimmed: it is what that engine produced.

### 4.9 Other kinds

`gear`, `talents`, `drops` and `weights` keep `SavedCombos` and `SavedWeights`. They take two parts of this spec, so the site reads as one:

- the result band, with the kind title as `h1`, the kind's own answer as the figure (the top combination's gain, or "Weights" with no figure), the primary button "Open in the sim", and only the Fight, Targets, Buffs, Runs and Saved facts;
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

There is exactly one filled gold element on the page: "Open in the sim". Whitespace: 32px between cards; 16 to 20px inside them; no card holds two figures at the band's size.

## 6. Copy (verbatim)

No em dashes anywhere. Figures are examples from this run.

**Band:**
- Eyebrow "SAVED SIM"; h1 per §3.3.
- Descriptor "Troll Fury Warrior · Level 60 · 17/32/2 talents".
- Next step "Next: find the upgrade that adds the most." and "Run Top Gear".
- Figure "DPS", "± 0.5 · 0.05% error".
- Buttons "Open in the sim", "Copy link", "Copied", "Copy failed".
- Facts "Fight", "Targets", "Buffs", "Consumables", "Range", "Runs", "Saved".
- Values "3:00 ± 20%", "Raid" / "None" / "Custom", "8" / "None", "731 to 1,116", "4 Nov 2026".

**Nav:** "Damage", "Timeline", "Gear", "Talents", "Rage" / "Energy" / "Mana", "Buffs", "Rotation", "Settings". Nav `aria-label` "Sections of this result".

**Worth knowing:** label "THIS RUN", title "Worth knowing", meta "1 note" / "{n} notes". W1 to W6 as in §3.5.

**Damage by ability:**
- Label "THIS RUN", title "Damage by ability", meta "{904.4} DPS over {2:59} on average".
- Headers "Ability", "DPS", "Share", "Casts", "Avg hit", "Crit", "Biggest hit", "Uptime".
- Sub-lines "Miss {x}%", "Dodge {x}%", "Parry {x}%", "Glance {x}%", "Block {x}%", "{n} ticks · tick crit {x}%", "{n} swings became Heroic Strike", "From {Windfury Totem}" / "From {Hand of Justice} {6} · {Windfury Totem} {21}".
- Phone sub-line "{n} casts · {avg} avg hit · crit {x}%" ("swings" for a white row).
- Names "Main hand", "Off hand", "Extra attacks". Pet sub-heading "Pet · {name}".
- Total "Total, {n} sources" and "{162,715} damage a fight".

**Spread:**
- Label "{10,000} RUNS", title "DPS spread", axis middle "DPS per run".
- Caption "Two runs in three land between {851} and {961} DPS."
- Tiles "Mean", "Error", "Spread", "Lowest", "Highest", "Runs".
- Bar tooltip "{901} to {910} DPS: {800} runs".
- Image label "DPS of each of {10,000} runs, from {731} to {1,116}, most near {904}".

**Timeline:**
- Label "ONE FIGHT", title "Timeline", meta "The median run · {2:53} long".
- Chips "Whole fight", "Close up". Switch "Swings and ticks".
- Legend "Damage", "Crit", "Missed", "Cooldown", "On you", "On the target", "Rage".
- Groups "Casts", "Swings and ticks", "On you", "On the target". Lane "Overpower window" (§3.7). Lane names as resolved; the periodic lane "{Deep Wounds} ticks".
- Axis "Time", "Pull". Wash "EXECUTE PHASE". Resource "Rage" / "peak {100} of {100}".
- Readout "{2:25.0}", "Rage {0} · On you: {Death Wish, Recklessness, Berserker Rage}" (with nothing: "On you: nothing extra"), "{0.5} s ago".
- Tooltip "{2:13.0} · rage after {12}", "Next {Execute} {1.5} s later", "Last {Execute} of the fight", "{Crit} · {2,275} damage" ("Hit", "Crit", "Glance", "Miss", "Dodge", "Parry", "Block").
- Cast list "Casts: {120} of {120}. Swings and {Deep Wounds} ticks are on the lanes above.", "Casts in view: {19} of {120}", "Show all {n}", "Show fewer", "{Heroic Strike} queued", "Rage {100}".
- Unpin button `aria-label` "Close the readout".

**Gear:**
- Label "THIS RUN'S INPUT", title "Gear", meta "{17} of 17 slots · {11} enchants" / "· no enchants".
- Enchant lines as the client prints them ("+7 Strength", "Crusader").
- Stats head "Character stats in the fight, against a level {63} boss"; labels "Attack power", "Hit", "Crit on the boss", "Weapon skill", "Haste"; reference "8% specials · 27% dual-wield swings".
- Foot "No set bonus is active: no two pieces share a set.", "Set: {name} {2} of {5}", "Upgrades for this character on Best in slot", "This run had no gear on file."

**Talents:** label "THIS RUN'S INPUT", title "Talents", "Read only. Hover or focus a talent for its tooltip.", phone "Read only. Tap a talent for its tooltip.", "Open in planner", §4.7's line, "Talents didn't load."

**Resource:** label "AVERAGE FIGHT", title "Rage", tiles "Gained", "Spent", "Lost at the cap" (mana: "Ended with"), lists "Gained from", "Spent on", overflow "−{22.8} lost", §4.3's line.

**Buffs:** label "AVERAGE FIGHT", title "Buffs and debuffs", meta "Uptime, then times applied", "On you", "On the target", "{10} that come and go", "Up all fight", "× {5}", "{91}×", §4.3's line.

**Rotation:**
- Label "{FURY}, AT {60}", title "Rotation", meta "{12} lines · casts are per fight" (+ " · today's rotation").
- "Before the pull", "Cast", "Then, from the top at every decision".
- "Other cooldowns as they come up", "This run: {Mighty Rage Potion, Berserking, Chthonic Power}.".
- "{24} casts" / "1 cast", "Not cast in this run".
- Foot "The {Fury} rotation the engine follows. A line with no cast in this run is dimmed.", "Read the {Fury} guide".

**Settings:**
- Label "THIS RUN'S INPUT", title "Settings".
- Keys "Fight length", "Target", "Execute phase", "Race", "Runs", "Actions per minute", "Engine".
- Values "3:00, ± 20% (2:24 to 3:36)", "Level {63} boss", "Level {63} boss · {armor} armor, {after} after debuffs", "Last {25}% of the fight, standing in for the boss under 20% health".
- Tag "Older engine".
- Groups "On you", "On the target", "Consumables"; sub-label "No effect on your damage"; suppressed "{Sunder Armor} · replaced by {Expose Armor}"; hands "(main hand)", "(off hand)"; grade ", improved"; "None in this run."; "Show all {14}".
- Actions "Copy for the game", "Copy JSON", "Download JSON", "The JSON is this whole result, request included."

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
| A5 | Worth knowing with W1, W2 and W5 forced (fixture variant) | 1440, 390 |
| A6 | Damage table: hover on a row, focus ring on a sort header, sorted by Crit ascending, spell tooltip open on Execute | 1440, 390 (tooltip by tap) |
| A7 | Timeline: Whole fight; Close up with the brush window; Swings and ticks on; crosshair readout; pinned readout; tick tooltip; keyboard focus on a tick | 1440, 390 |
| A8 | Timeline scrolled to the end on a phone (sticky labels and axis visible) | 390 |
| A9 | Histogram bar tooltip | 1440 |
| A10 | Gear: an item tooltip open (left and right column); a set bonus active (fixture variant) | 1440, 390 |
| A11 | Talents: three trees; phone tabs; a talent tooltip; §4.7 fallback (the previous fixture's string 33305013002000000-15353100051010501-000000000000000000) | 1440, 390 |
| A11b | Extra attacks row with its sources; the character stats block; an imported character with enchants (client enchant text) | 1440, 390 |
| A11c | Settings: Target row with armor, the "No effect on your damage" group, a suppressed-debuff chip (fixture variant) | 1440, 390 |
| A12 | Rotation with a dimmed line; Settings with Show all expanded on a phone | 1440, 390 |
| A13 | Copy JSON copied state; Copy failed state | 1440 |
| A14 | Loading: the prerendered skeleton before hydration, and each lazy card's skeleton (throttled) | 390, 1440 |
| A15 | Older result: `web/src/fixtures/sim/result.json` (no new fields), full page | 390, 1440 |
| A16 | Load failed (API 500) and a lazy chunk failure | 390, 1440 |
| A17 | Stale engine variant | 1440 |
| A18 | Lighthouse report for `/sim/simfixtureab.html` and the island size line from the build | — |

## 10. Accuracy checks (tenet 8; recorded in each lane report)

1. Every figure in the band, table, tiles and lists matches a hand calculation from the fixture JSON. The fixture test pins:
   - Execute: 190.8 DPS, 21.1%, 24 casts, 1,493 avg hit (34,337 / 23 landed), 52.2% crit (12 / 23).
   - Main hand: 63 swings, 506 avg hit (23,802 / 47 landed), 42.6% crit (20 / 47), miss 19.0%, dodge 6.3%, glance 39.7% (each of 63 attempts).
   - Extra attacks: 21 swings, 557 avg hit.
   - Deep Wounds: 59 ticks, 141 per tick, tick crit 37.3%.

   Attempts are `hits + MISS + DODGE + PARRY`, because the adapter counts glances and blocks inside `hits` (`sim/adapter/adapter.go` `ability()`).
2. Actions per minute: 40.4, from 121 presses over 179.9 s.
3. The talent split shown in the band equals the Talents card's meta, and both equal `talentSplit(request.character.talents)`: 17/32/2.
4. Every action key in the fixture resolves to a name and an icon. A test fails on "An unnamed spell", or on a missing icon for any key not on §11.4's documented gap list.
5. On the fixture, W4 fires for Recklessness and nothing else. On a variant of the fixture with the `spell:23894` rows removed, W2 fires.
6. The histogram counts sum to `iterations_run` (10,000).
7. Resource sources sum to `gained` (2,127) and `spent` (2,116) within 0.5.
8. Illustrative fixture values must not contradict the run's own summary. Two checks:
   - The sample's outcome mix per ability must land within 15 points of the summary's crit and miss rates.
   - A DEBUFF track's uptime must agree with its periodic row's ticks times the tick period.

   Lane E runs both as a fixture test.

## 11. Data

### 11.1 Have today (verified on the real run)

- `dps` {mean, stddev, error, min, max}; `iterations_run`; `request` (gear with enchants, talents, buffs, consumes, encounter including `target_level`, `target_armor` when set, and `execute_ratio`); `engine_version`; `lane`.
- `summary.damage_done[0].abilities[]` {spell_id, name key, school, total, hits (including crits and glances), crits, ticks, misses map, max (0 today)}. These are per-iteration averages.
- White swings are split by tag: `other:attack/1` main hand, `/2` off hand, `/3` extra attacks (Windfury Totem on this run).
- `summary.casts[]`; `summary.auras[]` (player BUFF tracks: applications, uptime_ms); `summary.resources[0]` {gained, spent, wasted}.
- `sample[]` {at_ms, action, target, resources}: the median-DPS run's events, white swings and periodic ticks included.

### 11.2 New fields (engine and adapter lane)

All are additive on `SimResult` (`sim/api/envelope.go`, mirrored in `web/src/lib/sim/types.ts`). All are optional (`omitempty`). Stored results need no migration (`sims.result` is jsonb).

The hand fixture `web/src/fixtures/sim/result-full.json` carries most of these shapes today. The ones it does not carry yet are `final_stats`, `target_stats`, `extra_attack_sources` and `rotation_sha`; lane E adds them to the fixture before the web lanes need them.

```jsonc
// SimResult additions
{
  // Engine DistributionMetrics.hist, rounded-DPS bins, contiguous from `start`.
  // counts sum to iterations_run. Absent for weights and bulk kinds.
  "dps_histogram": { "start": 731, "bin_dps": 1, "counts": [1, 0, 0 /* ... 385 bins */] },

  // Per resource, the average per fight, one decimal. Keys are action keys
  // (same grammar as summary rows), so names and icons resolve the same way.
  "resource_sources": [
    {
      "resource": "rage",
      "gained": [
        { "action": "other:attack/1", "amount": 850.8, "wasted": 22.8 },
        { "action": "other:attack/3", "amount": 255.2, "wasted": 11.4 },
        { "action": "item:13442", "amount": 148.9, "wasted": 5.7 },
        { "action": "spell:18499", "amount": 85.1, "wasted": 0 }
      ],
      "spent": [
        { "action": "spell:11567", "amount": 634.8 },
        { "action": "spell:20662", "amount": 529.0 }
      ]
    }
  ],

  // The sample run's own length; the timeline axis ends here.
  "sample_duration_ms": 173950,

  // Aura on/off segments inside the sample run, player and target.
  // `stacks` is optional: [at_ms, stacks] change points for a stacking aura.
  "sample_auras": [
    { "action": "spell:12328", "on": "player", "segments": [[130000, 160000]] },
    { "action": "spell:11597", "on": "target", "segments": [[0, 173950]], "stacks": [[0, 5]] }
  ],

  // The character as the engine computed it at the pull, with every buff and consumable,
  // against the request's target. Source: the engine's own stat computation for the
  // player unit (lane E confirms the call; wowsims computes it on the unit before the
  // first iteration). Percentages are percent points, one decimal.
  "final_stats": {
    "attack_power": 2210,
    "melee_hit_pct": 8.0,
    "melee_crit_pct_vs_target": 30.9,
    "melee_haste_pct": 0.0,
    "weapon_skill": { "main_hand": 300, "off_hand": 300 }
  },

  // The target as the engine fought it. armor_after_debuffs only when the engine reports it.
  "target_stats": { "level": 63, "armor": 3731, "armor_after_debuffs": 1471 },

  // Extra attacks by source, average per fight. Keys: a buff id from simbuffs or item:<id>.
  "extra_attack_sources": [ { "source": "windfury_totem", "count": 21.0 } ],

  // Content hash of the APL the run used (sha256 of the curated rotation JSON, first 12 hex).
  "rotation_sha": "3f9a0c1b7d22",

  // Saved time, a sibling like `title` on GET /v1/sims/{id} (api/internal/sims/handler.go GetOutput).
  "created_at": "2026-11-04T21:04:11Z"
}
```

The figures inside `final_stats`, `target_stats`, `extra_attack_sources` and `rotation_sha` are shape examples only. No board or fixture may show them as this run's values.

```jsonc
// SampleCast additions (each sample[] row), optional
{ "at_ms": 133000, "action": "spell:20662", "target": "Target 1", "resources": { "rage": 12 },
  "outcome": "crit",   // hit | crit | miss | dodge | parry | glance | block | tick | tick_crit
  "amount": 2275 }     // damage of this event, whole number; absent for a non-damage cast
```

**Target debuff tracks:** no new shape.

- The adapter adds `summary.auras[]` rows with `"type": "DEBUFF"`, `target_guid` `"target-1"`, `target_name` `"Target 1"`, applications, `uptime_ms` and `max_stacks`. They come from the engine's target unit metrics (`encounter_metrics.targets[].auras`).
- Damage by ability's Uptime column reads the DEBUFF track whose spell id matches a periodic row.

**Fixtures for the web lanes:**

- (a) `web/src/fixtures/sim/result-full.json`: the real published Fury run plus illustrative new fields. Lane E regenerates it from a real engine run before merge. Two values in it contradict the run (§10 check 8), and lane E fixes them first:
  - every one of the sample's 93 `spell:12867` events is `tick_crit` for 377, against the summary's 59 ticks, 37.3% tick crit and 141 per tick;
  - the sample has 93 Deep Wounds events against the summary's 59 ticks per fight. Lane E says which events the sample records (ticks only, or applications too) before the timeline counts ticks.
- (b) `web/src/fixtures/sim/result.json` stays the older shape. It is the fallback fixture for A15. The W2 state uses a variant of (a) with the `spell:23894` rows removed, made in lane A's test.

### 11.3 Request addition

`request.data_build` (string, e.g. "1.60.1.70790"): the talent and item build the request was made against. The web sets it when it builds a request. `talentsFit` decodes against that build's trees when it is not the active one.

### 11.4 Data files (data lane, inside lane E)

- **Icons**, in `data/curated/extra-spell-icons.json`: the client SpellIcon for 12319 Flurry, 20554 Berserking, 12867 Deep Wounds, 462891 Whirlwind (off-hand hit), 20906 Trueshot Aura, 10610 Windfury Totem, 409583 Heart of the Lion and 11722 Curse of the Elements. Also 1249110 Chthonic Power and 21919 Thrash, which have no icon anywhere in the repo today. Each is verified against the client's SpellIcon table, not copied from the mock.
- **Names**, in `data/curated/engine-spell-names.json`:
  - 11722 "Curse of the Elements";
  - 462891: the 1.60 client's own name if it has one. Otherwise "Whirlwind (off hand)", cited to the engine file that casts it. Confirm that the Forever Whirlwind tooltip says both weapons.
- **Thrash (21919)** is up all fight on this run, but no worn item names it. Lane E says where it comes from before it ships as a buff row.
- **`enchants.json`:** add `display`, the client's `SpellItemEnchantment` name ("+7 Strength", "Crusader"), for every enchant.
- **`simbuffs.json`:** add, per entry:
  - `target: true` for anything that lands on the target;
  - `stats` (the buff's stat block, for the "No effect on your damage" rule);
  - `exclusive_with` (the group of debuffs that do not stack, with the winner by strength).
- **`data/curated/specs.json`:** add `capstone_spell` per dps spec (the 31-point ability where the spec has one: warrior-fury 23894 Bloodthirst, warrior-arms 12294 Mortal Strike, and so on), for W2.
- **Curated APL notes:** the Execute-phase wording in §3.11, shipped with the guide text.
- **A test:** every action key in `result-full.json` resolves to a name and an icon. The only exceptions are ids on this list while their icon is pending.

## 12. Lane split

Three web lanes plus the engine and adapter lane. One integrator lands them (the batch-landings rule).

| Lane | Owns | Files | Starts from |
|---|---|---|---|
| **E: engine, adapter, data** | §11 | fork: `DistributionMetrics.hist` kept, sample aura segments, sample outcome and amount, per-action resource metrics, `final_stats`, `target_stats`, `extra_attack_sources`, `rotation_sha`, ability `max`; `sim/adapter` (Summarize, Sample, new envelope fields, DEBUFF tracks); `sim/api/envelope.go`; `api/internal/sims/handler.go` `created_at`; `web/src/lib/sim/types.ts`; `web/src/fixtures/sim/result-full.json`; §11.4 data files and test | now |
| **A: frame and damage** | §3.2 to §3.6, §3.12, §4.2 to §4.9, §8 skeleton and reservation | `SavedSim.svelte` (the only lane that edits it; it mounts B's and C's components), new `ResultBand.svelte`, nav, `lib/sim/warnings.ts`, `report/AbilityTable.svelte` (extracted from `ActorRow`), `DpsDistribution.svelte`, `DetailsCard.svelte`, `lib/sim/skeleton.ts`, `pages/sim/[id].astro`, `lib/sim/copy.ts` | E's hand fixture |
| **B: timeline** | §3.7 | new `SimTimeline.svelte` plus `lib/sim/timeline.ts` (lanes, segments, keyboard model), `SampleLog.svelte` rework; reuses `chart/*` | E's hand fixture |
| **C: inputs, resource, buffs, rotation** | §3.8 to §3.11, §4.7 | `Paperdoll.svelte` enchant line (and `lib/gear/paperdoll.ts`), a talents card over `GuideTreeColumns` plus `lib/sim/talents-fit.ts`, a resource card, `AuraTable.svelte` `variant="bars"`, `RotationCard.svelte` ; `lib/sim/extra-attacks.ts` is lane A's (the table), the stats block under the paperdoll is lane C's | E's hand fixture |

`SimResults.svelte` (the tab strip) stays only for the live `/sim` page until that page adopts this layout. That adoption is a follow-up spec, not this one.

Each lane runs only its touched tests. The integrator runs the full web vitest, the full Playwright suite and the lhci shard, because this rebuilds a page (the shared-layout rule).

## 13. Open questions

1. **Talent strings saved on an older layout** (the previous fixture: 17 Fury digits for 18 talents) still need lane E's answer: is the engine's talent order the client's tab order? The published run decodes, so the mock draws its trees. §4.7 stays the fallback, and W2 now catches the case that matters (a capstone never pressed) without depending on the decode.
2. **Off hand misses less than main hand** (this run 8.5% against 19.0%). The Heroic Strike queue removes the off hand's dual-wield penalty while it is queued, so the direction is right. The size goes to the sim-accuracy loop. The page carries no note about it until the loop signs off.
3. **`rotation_sha`:** accepted in §11.2. Lane E decides which hash.
4. **Holy Strength hand tags** (`/1`, `/2`): confirm from the adapter's tag allocation. They are not in this run (no Crusader).
5. **Overpower window:** settled. It is never a buff row; it is a timeline lane only when the rotation casts Overpower.
6. **`created_at`:** accepted in §11.2.
7. **Crit definition on white rows** (§3.6): this spec uses crits over landed hits with glances counted, as WCL does. The round 2 ruling said non-glancing, which reads 90.9% and 100.0% here. The coordinator confirms or overrules.
8. **`final_stats` source:** lane E confirms the engine call. The page shows nothing it would have to compute itself.
9. **Fixture contradictions** (§11.2 (a)): Deep Wounds sample events and outcomes.

## 14. Round 2: what changed per finding

Answers `design/reviews/2026-11-04-sim-results-spec-wow-player.md`, with the coordinator's rulings.

| # | Finding | Change |
|---|---|---|
| 1 | 0 Bloodthirst, no warning | The mock is rebuilt from the published BiS run (Bloodthirst 22 casts). New **W2 capstone never pressed** (§3.5), independent of `talentsFit`, with `capstone_spell` data (§11.4). The old run's W2 state is on `states-1440.png`. |
| 2 | Test items in the gear | Solved by the new run: real launch kit, no QA items. |
| 3 | Recipe names as enchant text | The enchant line prints the client's `SpellItemEnchantment` name ("+7 Strength"), via a new `display` field (§3.8, §11.4). The new run has no enchants, and the mock shows none. |
| 4 | Average per attempt | "Avg hit" = damage over landed hits; Crit = crits over landed hits (§3.6). The column "Biggest hit" appears when the engine returns `max`. On non-glancing crit, see open question 7. |
| 5 | Three Execute thresholds | One wording, "the Execute phase". Settings: "Last 25% of the fight, standing in for the boss under 20% health". The curated notes are rewritten by the data lane (§3.11). The timeline wash reads "EXECUTE PHASE". |
| 6 | No character stats | `final_stats` (§11.2) drawn as a stats block under the paperdoll (§3.8): Attack power, Hit against the 8% and 27% references, Crit on the boss, Weapon skill per hand, Haste. |
| 7 | No target | Settings "Target" row from `target_level`, plus armor and armor after debuffs from `target_stats` (§3.12). Suppressed-debuff chips via `exclusive_with`. Preset stacking goes to the accuracy loop. |
| 8 | Overpower as a buff | Never a buff row. An "Overpower window" timeline lane appears only when the rotation casts Overpower (§3.7, §3.10). |
| 9 | Off-hand miss note | No row note until the accuracy loop verifies the figure (§3.6, §13). |
| 10 | Phone lanes indistinguishable | Each lane prints its name at 11px muted inside the plot, sticky at its left edge (§3.7). |
| 11 | "Open in the sim" | Renamed "Open in the sim", with the destination specified (§3.3). |
| 12 | Extra attacks | Their own "Extra attacks" row and timeline lane (`other:attack/3`), with sources (§3.6). |
| 13 | Pre-pull rows | Dimmed with "Not cast in this run" like priority lines (§3.11). |
| 14 | Spread caption overclaims | The second sentence is dropped (§3.6). |
| 15 | Illustrative numbers contradict the run | The mock reads the hand fixture, whose sources sum to the run's totals. Two remaining fixture contradictions are listed for lane E, with a fixture test (§10 check 8, §11.2). |
| 16 | Two talent splits | One `talentSplit()` value in the band and the Talents meta (§3.3). |
| 17 | Heroic Strike doubled in Whole fight | Whole fight draws landings only; Close up adds the hollow queue ticks (§3.7). |
| 18 | Phone split breaks inside a pair | Each "Label value" pair is nowrap (§3.6). |
| 19 | "Ran on", engine sha | "Ran on" is gone. The engine sha is muted (§3.12). |
| 20 | Consumables in the band | The Consumables fact replaces Actions per minute, which moves to Settings (§3.3). |
| 21 | Buffs that do nothing | Damage first, then "No effect on your damage", decided from `simbuffs.json` `stats` against the spec's weights (§3.12). |
| 22 | Holy Strength labels | Kept, pending the tag check (§13). |
| 23 | Whirlwind (off hand) | Use the client's name if it has one; confirm the tooltip (§11.4). |
