# Simulator parity with Raidbots

**Date:** 2026-09-19
**Builds on:** `2026-09-14-simulator-design.md` (the simulator) and
`2026-09-14-simulator-interfaces.md` (the contract). This document changes
neither retroactively; where it amends them, the amendment is named.
**Source of the gap list:** a walk through every Raidbots tool, option
panel and report on 2026-09-19 (Top Gear, Droptimizer, Quick Sim,
Advanced, Stat Weights, Gear Compare, a finished Quick Sim report, the
developer page).

## Decisions

- **Scope is everything Raidbots offers a player, minus its developer
  surface.** Static data dumps, per-report raw files, talent-tree iframes
  and URL prefill are out of this round. Everything a player touches is in.
- **Free in the browser, bigger on the server.** Top Gear, Droptimizer,
  talent compare and stat weights all run on the browser lane under a cap.
  Premium lifts the cap and runs the same request on Cloud Run. Raidbots
  gives free users a queue; we give them their own CPU.
- **Candidates come from the addon and from search.** The companion's
  export string grows a bag-and-bank section. The page also searches the
  build's item database. Droptimizer sources feed the same candidate list.
- **One engine call shape for every combination tool.** Top Gear,
  Droptimizer and talent compare are the same thing: a base character, a
  set of substitutions, a ranked result. They share one request kind, one
  planner, one results view.
- **Gear Compare is not a tool.** Raidbots marks it legacy and steers people
  to Top Gear. Its one distinct use, "these two whole sets side by side",
  becomes named sets inside Top Gear.
- **Stat weights ship with the warning.** The engine already computes them.
  The page carries Raidbots' own caution in our words and links to Top Gear.
- **Advanced is a JSON editor, not a script language.** SimulationCraft
  input is Raidbots' escape hatch. Ours is the request envelope itself,
  editable and shareable, with the same validation the page uses.
- **Every setting Raidbots exposes has an equivalent here or a written reason
  it does not.** The reasons are in section 4; most are "the vanilla-era
  game has no such thing" (gems, catalysts, upgrade tracks, hero talents).

## 1. The gap, tool by tool

| Raidbots | Here today | After this spec | Section |
| --- | --- | --- | --- |
| Quick Sim | `/sim`: run, breakdown, uptimes, timelines, distribution, saved, history, compare-to-log | Adds fight styles, full buff/consumable panel, Smart Sim precision, margin-of-error line, sample-iteration log, report title | 4, 5 |
| Top Gear | nothing | `/sim/gear`: candidates per slot, enchants, consumables, talents, named sets; staged precision; ranked results | 2, 3 |
| Droptimizer | nothing | `/sim/drops`: raids per boss, dungeons per boss, world bosses, crafted, reputation, PvP; phase-aware | 2, 6 |
| Talent loadouts in Top Gear | planner builds exist; not simmable in bulk | loadouts as candidates in Top Gear and a `/sim/talents` tab | 2, 3.4 |
| Stat Weights | engine has it | `/sim/weights` with the warning | 7 |
| Advanced | nothing | request JSON editor on every tool | 8 |
| Armory / addon / history input | all three | addon export gains bags, bank, loadouts | 9 |
| Premium queue skip, larger sims | premium flag, server lane for single runs | caps by lane; server runs every kind | 10 |
| Share URL, run again, report options | share, run again | report title, browser notification, open in new tab | 5.4 |

## 2. One request kind for every combination tool

### 2.1 Envelope amendment

`SimRequest` gains an optional `bulk` block. A request with no `bulk` is
exactly today's single run; nothing existing changes shape.

```
bulk:
  candidates:      [ { slot, item_id, enchant?, suffix?, origin } ]   # origin: equipped | bag | bank | search | drop:<source-id> | set:<name>
  talents:         [ { name, talents } ]                               # positional strings; empty means the character's own
  sets:            [ { name, gear: [GearSlot] } ]                      # whole-set alternatives (Gear Compare's use case)
  combinations:    bool        # false: one substitution at a time (Droptimizer); true: every valid combination (Top Gear)
  precision:       fast | normal | high
  cap:             int         # combinations the lane allows; the planner refuses more, it never silently trims
```

`SimResult` gains `combos`: a ranked array of `{ substitutions:
[GearSlot|talents|set], dps: Estimate, delta: Estimate, stage: int }`,
plus `equipped: Estimate` and `stages: [ { iterations, combos_run } ]`.
`delta` is against `equipped`, with its own error, so the page can say
"within error" honestly.

The spec amendment to `2026-09-14-simulator-interfaces.md` is the two
blocks above and the `kind` column on `sims` (section 10.3).

### 2.2 The planner lives in the `sim` module, once

`sim/bulk` (new package, Go, compiled into both the wasm and the native
binary) owns everything that decides *which* sims run:

- **Expansion.** From the base character and the candidates, the valid
  combinations: slot fit from the item's inventory type and the build's
  item table (the same `simdb` the engine loads); class and level
  requirements; faction; unique-equipped and unique-category limits from the
  fork's item limit categories; rings and trinkets tried in both slots;
  two-hand versus main-hand-plus-off-hand as competing shapes; dual-wield
  specs try weapon pairs both ways; enchants only where the enchant's slot
  and item class allow. Enchant defaulting: a candidate without an enchant
  inherits the equipped item's enchant for that slot when it fits, the way
  Raidbots and the fork's `auto_enchant` do.
- **Staging.** `fast` runs every combination at 100 iterations, keeps the
  top quarter plus anything within two standard errors of the cut, reruns
  at 1,000, keeps the top ten plus ties, and runs the finalists and the
  equipped set at the precision's final count (3,000; `high` is 10,000).
  The equipped set is in every stage so every delta is paired. `normal`
  skips the 100-iteration stage; `high` also doubles the finalist count.
  This is the fork's `fast_mode` ladder made explicit so both lanes agree.
- **Ranking and grouping.** Sorted by mean; "within error" groups are
  runs whose delta intervals overlap the leader's. The result carries the
  group boundaries so the page does not recompute statistics.

Exports: `simPlan(request) → stage 1 requests`, `simRank(stage results,
request) → next stage requests or final SimResult`. The browser pool runs
each stage's requests through the existing `simRun` sharding; the native
job calls the same two functions in a loop. Nothing about statistics
exists in TypeScript.

### 2.3 Caps

| Lane | Combinations | Precision | Where the number comes from |
| --- | --- | --- | --- |
| Browser | 400 | fast or normal | 400 × 100 iterations first stage is about a minute on a mid laptop with 8 workers; a phone gets half the cap by `hardwareConcurrency` |
| Premium | 20,000 | any | the fork's own `maxIterations` guard, Cloud Run at 4 CPUs finishes within the 15-minute job timeout |

The page shows the live combination count as candidates are ticked, the
way Raidbots does ("1 valid combination"), and the run button says what
would exceed the cap and by how much.

## 3. Top Gear, `/sim/gear`

### 3.1 Inputs, top to bottom

1. **Character strip** as `/sim`, same source switcher.
2. **Gear** — the slot grid. Each slot lists: the equipped item; bag and
   bank items that fit the slot (from the export, section 9); items added
   by search; items pinned from Droptimizer. A row is a checkbox, an icon, a
   name, an item level, and a "copy and modify" menu to add the same item
   with a different enchant or suffix ("of the Bear" etc. from the build's
   suffix table). Locks: a slot can be locked to the equipped item so it is
   never substituted. Rings and trinkets show one list for both slots.
3. **Item search** — name, minimum item level, slot, source filter, "only
   usable by this character" on by default. Backed by the build's
   `items.json` client-side; no API call.
4. **Enchants** — per slot, the enchants the build's enchant table allows
   for that slot, each a checkbox; "keep current" and "none" rows; a
   selection cap per slot shown as Raidbots shows it.
5. **Consumables and buffs** — the full panel from section 4.3, plus
   "try each of these" checkboxes on flasks, weapon oils and food so a
   consumable can be a candidate rather than a setting.
6. **Talents** — the character's own; every planner build saved by the
   signed-in player for this class; in-game loadouts from the export;
   "add custom" opens the planner inline and returns a string.
7. **Sets** — named whole-gear alternatives ("my AQ set", "my PvP set"),
   built from the export or pasted as a second export string. A set is one
   candidate that replaces every slot at once.
8. **Run** — combination count, precision, cap notice, the lane switch for
   premium, one button.

### 3.2 Rules the page states in words

The Raidbots "notes and limitations" list, ours: equipped enchants carry
over where they fit; rings and trinkets are tried in both slots; two-hand
and one-hand-plus-off-hand are competing shapes; dual-wield tries both
orders; unique-equipped is respected; nothing the character cannot equip
is ever simmed; item search can find items the character cannot obtain.

### 3.3 Results

- The equipped set's DPS as the baseline line, then ranked rows: DPS,
  delta with its error, percent, and the substitutions as icon chips with
  tooltips. Rows in the leader's within-error group carry the same rank.
- A per-slot summary: for each slot, the item the winner uses and the
  gain the slot contributed (from the single-substitution stage where
  `combinations` is on, else from the run itself).
- "Open in planner" carries the winning set into the planner; "Copy to
  addon" produces the export string for the winning set so the companion
  can show the swaps in game.
- Saved like every sim at `/sim/<id>` with the ranked table, and in
  history as "Top Gear · 38 combinations".

### 3.4 Talent compare, `/sim/talents`

The same page with gear locked and the talent list as the only candidates:
the character's build against every saved planner build and the export's
loadouts, ranked. It exists as its own tab because the question "which
build" is asked far more often alone than with gear.

## 4. Settings parity

### 4.1 Fight styles

A named style sets the encounter fields the engine already has plus two
new ones. Styles, with Raidbots' name where ours differs:

| Style | Targets | Execute | Movement | Notes |
| --- | --- | --- | --- | --- |
| Patchwerk | 1 | on | none | the default |
| Execute heavy ("Execute Patchwerk") | 1 | 35% of the fight under 20% | none | |
| Light movement | 1 | on | 5 s out of melee or casting every 45 s | new engine field |
| Heavy movement | 1 | on | 5 s every 20 s | new engine field |
| Cleave 2 / Cleave 3 / Cleave 5 ("Hectic Add Cleave", "Cleave Add") | 2, 3, 5 | on | none | adds spawn with the boss |
| Dungeon pull ("Dungeon Slice") | 1 boss then packs of 3, 5, 3 | off | none | a target-count timeline, new engine field |
| Target dummy | 1 | off | none | no debuffs, no target armor reduction, infinite duration capped at the fight length |

Movement and the target-count timeline are the two engine additions: an
`Encounter.movement` block (interval, duration, kind: away or casting-only)
honoured by the APL's movement conditions, and `Encounter.targets_over_time`
as `[ { at_sec, count } ]`. Both are Forever fork commits behind the pin,
tested like every core change.

Also exposed beside the style: number of targets (1 to 10), fight length
(20 seconds to 10 minutes, default 3 minutes), duration variation (0 to
30%, default 20%), target level (60 to 63), target armor (preset per
level with an override), target type (humanoid, undead, beast and the
rest, since it changes Hunter and Warlock abilities).

### 4.2 Precision and Smart Sim

Three precisions: fast (500), normal (3,000), high (10,000). Beside them,
**"until ±0.5%"**: the run continues in 1,000-iteration steps until the
DPS error is inside half a percent or the lane's iteration ceiling
(browser 30,000, premium 100,000) is hit, and the results line says which.
This is Raidbots' Smart Sim made visible rather than silent. The
margin-of-error line ("±41 DPS, 0.4%") is always shown, as is the
iteration count and processing time.

### 4.3 Buffs, debuffs, consumables: the full panel

The preset stays the default. "Custom" opens the full list from
`sim/request/IDS.md`, grouped the way the engine groups them: raid buffs,
party buffs, this player's blessings and auras, debuffs on the target,
world buffs, consumables by kind (flask, battle and guardian elixirs, food,
weapon oil and stone, potions and runes, engineering explosives). Each row
is the id's display name and icon from the build; graded buffs (talented
versions) become a three-way choice, which is the vocabulary change IDS.md
already reserves. Cooldown timing for major cooldowns and potions (use on
pull, at a time, at execute) rides along in the same panel; the engine's
`Cooldowns` message already carries it.

### 4.4 What Raidbots has that we deliberately do not

Gems and sockets, catalyst charges, upgrade currencies and tracks, item
sets with tier-bonus minimums, hero-talent trees, crafted-stat variants,
bonus rolls, SimC version channel. The first six do not exist in the
game; crafted variants are suffixes here and handled by copy-and-modify;
bonus rolls have no equivalent; the engine version is pinned and shown,
not chosen. Tier set bonuses are counted by the engine from the gear
itself, so a "minimum set bonus" filter is a results filter ("only combos
keeping 4-piece"), added to the Top Gear results view rather than the
inputs.

## 5. Report parity

### 5.1 What a finished sim shows

Everything today's `/sim` results show, plus:

- **Margin of error, iterations, processing time, engine version and
  lane** as a details card beside the results, the way Raidbots' sidebar
  reads. The engine version links to `/sim/specs`.
- **Buff uptime with counts** (the aura table gains a count column where
  it lacks one).
- **Sample iteration log**: one iteration's casts in order with the
  pre-pull section separated, resources at each cast, and the note that it
  is one iteration and not a guide. The engine's cast log for the median
  iteration is already produced for the cast timeline; this is a table view
  of it.
- **Rotation card**: the rotation the run used with a link to its page, and
  for a non-validated spec the fidelity note.

### 5.2 Top Gear and Droptimizer reports

Section 3.3 and section 6.3. Both are saved sims with a kind and render
from the same `/sim/<id>` route.

### 5.3 Stat weights report

Section 7.

### 5.4 Report options

Report title (defaults to "Top Gear · Fury · 38 combinations"), browser
notification when a server run completes, open the result in a new tab.
Share URL exists; it stays.

## 6. Droptimizer, `/sim/drops`

### 6.1 Sources

A source is a set of items with a place they come from. The picker shows,
grouped, with icons from the zone pages:

- **Raids** — each raid, then each boss with its loot; a raid card runs
  every boss. Only raids the phase calendar has released are selectable;
  unreleased ones show with their date and a "show upcoming" toggle.
- **Dungeons** — each dungeon, by boss, plus "trash and chests".
- **World bosses.**
- **Crafted** — by profession, split into "my professions" and all.
- **Reputation** — by faction and standing, with the character's standing
  from the Armory where known.
- **PvP** — rank sets and the honor vendors, by rank.
- **Quests** — off by default; on request, since a quest reward is a
  one-time source.

### 6.2 Data

The fork's item database already carries sources for every Era item from
AtlasLoot and Wowhead: drops by zone and NPC, crafted by profession and
spell, sold by, reputation, quest. The data lane imports that into
`data/builds/<build>/loot.json` keyed by item id, joined to the build's
zones, and adds a curated overlay `data/curated/loot/*.json` for Forever's
own additions (items the fork's database has never seen, changed drop
locations, the phase each raid opens), with sources and notes like every
curated fact. Drop chances are not in either database; the page never
shows a probability it does not have. "Chance of an upgrade" is stated as
"3 of the 11 drops on this boss are upgrades", which is what the data
supports.

### 6.3 Run and results

Droptimizer is a bulk request with `combinations: false`, every item of
the chosen sources as a candidate, and the equipped set in every stage. The
results view is by source: each boss with its upgrades ranked, the best
gain per boss, and a flat "every upgrade" list. Rings and trinkets are
tried in both slots; a two-hand drop competes against the equipped
main-plus-off-hand. The same rules card as Top Gear.

A drop can be pinned into Top Gear with one click, which is how "this boss
drops a weapon and a ring, what if I got both" is answered.

## 7. Stat weights, `/sim/weights`

The engine's `StatWeights` request over the loaded character and settings.
The page opens with the caution, in our words: weights are a linear
guess at a non-linear thing; sim the actual items instead; here they are
anyway because addons want them. Results: each stat's weight with its
error, the reference stat normalised to one, and a "copy for Pawn" string.
Runs in the browser at normal precision, on the server for high.

## 8. Advanced

Every tool has a "Request" drawer: the exact JSON the run will send,
editable, validated by the same `Validate` the engine applies, with the
errors shown inline. A pasted request from another player's share loads
the whole page state. This is the escape hatch: any field the panel does
not expose is reachable here, and a share URL of an edited request is a
full reproduction.

## 9. Input parity

### 9.1 The addon export, version 2

The companion's export string gains sections after the gear list:

```
FS1:<build>:<class>:<race>:<talents>:<gear>|bags=<slot list>|bank=<slot list>|sets=<name>=<gear>;…|loadouts=<name>=<talents>;…
```

Bag and bank items are `item_id[:enchant[:suffix]]` and only equippable
items are exported. The parser in `web/src/lib/sim` accepts version 1
strings unchanged. The companion keeps its size in the chat-paste limit by
omitting bank when the string would exceed it and saying so in game.

### 9.2 Armory

Unchanged: equipped gear and talents only, with the Raidbots-style notice
that bags need the addon.

### 9.3 History

Lists every kind with its headline ("Top Gear · +41 DPS from Vis'kag"),
filterable by kind.

## 10. Compute and storage

### 10.1 Browser

The worker pool runs each stage's requests as ordinary sharded runs with a
progress line "stage 2 of 3 · 31 of 96 combinations". Abort cancels the
stage and returns what finished, marked partial, the way single runs do.

### 10.2 Server

`POST /v1/sims` accepts the same request; the job row records the kind;
the `sim-run` Cloud Run job runs the native planner loop and streams stage
progress. Job resources: bulk requests get the 4-CPU job at a 15-minute
timeout; a request the planner estimates past that is refused at submit
with the estimate.

### 10.3 Storage

`sims` gains `kind` (`run`, `gear`, `drops`, `talents`, `weights`) and the
result blob carries `combos` or `weights`. Migration 0014. Saved sims of
every kind are public at `/sim/<id>` with the unfurl card naming the kind
and headline.

## 11. Validation and tests

- `sim/bulk` is table-tested for expansion rules against the fixture
  characters (unique-equipped, both-slot rings, weapon shapes, enchant
  fit) and for staging arithmetic on fabricated results.
- The wasm smoke in `sim.yml` runs one Top Gear plan of four candidates
  end to end at 100 iterations.
- The API test suite covers the kind column, submit-time cap refusal and
  job progress for a bulk request with the fixture engine.
- Web e2e per page with the fake engine, plus one gated real-engine run of
  Top Gear with two candidates.
- The nightly validation job is unchanged; fidelity applies to the
  rotation, not to the optimiser.

## 12. Lanes and order

Five lanes, parallel where they do not touch the same files:

1. **Engine** (fork): movement block, targets-over-time, target-dummy
   mode; cast-log resources for the sample log. Pin bump.
2. **Sim module**: envelope amendment, `sim/bulk`, wasm exports
   `simPlan`/`simRank`, native loop in `forever-sim`.
3. **Data**: `loot.json` import from the fork database, curated overlay,
   enchant table, suffix table, phase dates for raids.
4. **API**: migration 0014, kind on submit and history, cap refusal,
   bulk in the job.
5. **Web**: settings panel, fight styles, precision and Smart Sim, report
   additions, then `/sim/gear`, `/sim/talents`, `/sim/drops`,
   `/sim/weights`, the request drawer.
6. **Companion**: export version 2.

Before launch on 4 November: lanes 1, 2, 4, 6, the settings and report
work, Top Gear and talent compare, dungeon and crafted sources for
Droptimizer. Before raids on 9 December: raid sources with the phase gate,
stat weights, world bosses and PvP.

## 13. Risks

- **Forever-only items with no source.** The fork's database covers
  Era's items; Forever's additions need curation. The overlay exists for
  this, and a source-less item is still simmable from search.
- **Browser cap feels small.** Four hundred combinations is a modest Top
  Gear. The cap is one number, tunable after measuring real devices, and
  the run button always says what premium would allow.
- **Movement modelling is spec-dependent.** A movement window is only as
  honest as each rotation's handling of it. Styles that use movement carry
  a note until the validation job has parses for them.
- **Export string length.** Bank contents can push past the chat paste
  limit; the companion drops bank first and says so.
