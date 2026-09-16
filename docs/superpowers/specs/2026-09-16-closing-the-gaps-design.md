# Closing the gaps

Status: draft for review. Written 2026-09-16 after thirty-six persona review rounds on the
report page, when the remaining asks had settled into six features rather than fixes.

## What it is for

Against Warcraft Logs the report page is at parity or ahead on a single pull's tables,
deaths, mechanics, the whole-night view and mobile, and behind on threat, Compare, phases
and per-ability graphs. Those four are closable with the data we already parse. Two
smaller asks that every healer and DPS round repeated, pet casts and aura refreshes, ride
along. The features, in the order the reviewers asked:

1. Threat over time, and threat measured inside a window.
2. Compare at ability level, player against player, with a CSV.
3. Resources: cap line, time at maximum, wasted at cap, CSV.
4. Pet rows on Casts, and aura refresh lines in Events.
5. Per-ability graphs on the main chart.
6. Phases, from curated per-boss tables.

Out of scope, and why: tank stance and taunt threat multipliers wait for Forever's ability
data; rankings depth is a population problem; replay needs positions we do not draw; spell
and item tooltips wait for Forever's data files.

## Global constraints

- The engine version goes to 0.4.0 with the first summary-shape change and stays there for
  the rest; the fixture is regenerated and prettier-formatted on every engine change, and
  the fixture test reads the version from the engine source.
- Every summary field added is optional on the web (`?:` in `types.ts`): a report parsed
  before 0.4.0 renders as it does today, with the new panels absent rather than broken.
- Every figure that is prorated rather than measured carries the tilde and the existing
  approximate title; a figure that is measured never does.
- The night (`?fight=all`) folds what folds by name and hides what cannot fold (per-second
  series across pulls); nothing on the night prints a figure the fold cannot stand behind.
- Phone first: every new panel is checked at 390px with no horizontal scroll, a 44px tap
  target, and its words beside its figures where the desktop has column headers.
- No third-party requests; the DuckDB measures use the shared query layer and the
  versioned fight-file urls.
- Each feature is a task group in the plan, reviewed and merged to main on its own, so a
  later feature never holds an earlier one back.

## 1. Threat over time, and threat measured inside a window

### The question

"Did I pull the boss off the tank at 3:54" needs threat as a line, per enemy, with every
player on one axis, and a window that reads threat from the window's own events. Today the
Threat tab has whole-fight totals per enemy and prorates them under a brush, and says so.

### Engine

`ThreatPair` gains `series: []int64`: the threat the player built on that enemy in each
whole second of the fight, on the same one-second buckets the damage series use, rounded
to an integer. `creditThreat` writes it wherever it adds to the pair's total, so the sum of
a pair's series is its total to the unit, and healing threat spread over engaged enemies
lands in the same buckets. Nothing else in the model changes; the base model's version
string stays `base-1` and `complete` stays false until the modifiers land.

Size: fight 20 has 27 pairs over 724 seconds, about 20,000 small integers, roughly 60 KB
on a 353 KB file. Acceptable; a pair with no threat at all is not emitted, as today.

### Web

- A threat chart on the Threat tab, above the table, for the picked enemy. With "Every
  enemy" picked it shows the enemy with the most threat and says which. One cumulative
  line per player, in class colours, on one axis; taunt marks as gold ticks on the time
  axis, named on hover; the same brush the page's window uses, so a window set here is
  the report's window. The tank is not privileged: the line that is highest is on top.
- Under a window the table stops prorating. Two figures per player: **standing** threat
  at the window's end (the cumulative line's value there, the number that decides aggro)
  and **built** inside the window (the series summed over the window). The table sorts by
  standing, the bar draws standing, the share is of standing. Both are measured, so no
  tilde and no "no order to read" note; the note becomes one line saying which figure is
  which.
- The whole fight is unchanged: standing at the end equals the total.
- The taunt list's "Around it" sets the window as today, and now lands on a measured table.
- The night keeps totals only. The chart is not shown on the night, and the note says
  a night has no clock to draw threat on.
- Glossary: "Standing threat", "Threat built in a window".

### Model honesty

The line is the base model's accumulation. A taunt in the game sets the taunter's threat
to the top; the model does not do that, and the note under the chart says so in the same
words the table's note uses today, until the modifiers land.

## 2. Compare at ability level, player against player, with a CSV

### The question

"What did the player above me do differently" and "what did I do differently on the kill"
are ability lists side by side with a signed difference, not two cards and subtraction.

### Web (no engine change)

- Every row of the Compare table expands. Expanded, it lists the player's abilities from
  the metric's table (damage done for damage and DPS, healing for healing and HPS, damage
  taken for taken and DTPS; nothing for threat, which has no ability split) in both pulls:
  ability, this pull, the other pull, the signed difference, sorted by the size of the
  difference. Abilities present in one pull only show a dash on the other side.
- A second picker beside "Compare with": **a player**. With a player picked, the table is
  the two players inside the current pull (and window), one row per ability from the
  metric's table, same three columns. The per-player table above collapses to the two
  players' totals.
- Both views honour the window the way the table does today, and mark what is prorated.
- A CSV of whatever is on screen: the player table, or the expanded ability diff, in the
  same shape as every other table's CSV.
- The url carries the picked player as `vs=<guid>`, alongside `with=` and `cmetric=`.

## 3. Resources: cap, time at maximum, wasted at cap, CSV

### Engine

`ResourceTrack` gains `max: int64` (the largest `MaxPower` read for the track), `at_max_ms`
(whole seconds the reading was at max, times 1000, on the same buckets the series uses)
and `wasted: int64` (the sum of `OverEnergize` on energize events for the track, which is
what the client says was gained past the cap). Version 0.4.0.

### Web

- The graph draws the cap as a line and shades the seconds at maximum.
- The figure line reads `peak N · at cap M% of the fight · wasted W` beside the existing
  low and never-empty readouts; each figure carries a title saying what it is.
- A CSV of the series: second, reading, at cap.
- Over the night, wasted sums and at-cap time sums; the cap line uses the largest max.
- Glossary: "At cap", "Wasted".

## 4. Pet rows on Casts, and aura refresh lines in Events

### Engine

`CastRow` gains `owner_guid` (the pet's owner from the registry, or the caster's own guid).
Version 0.4.0.

### Web

- The Casts source scope keeps a pet's rows under its owner, the way the damage tables
  fold a pet's damage: picking the healer shows the statue's and Yu'lon's casts too, each
  row saying `via Jade Serpent Statue` after the spell, as ability rows do.
- Events: the full stream includes `aura_refresh` lines, worded "Deadclasslol refreshed
  Renewing Mist on Hobolol", under the "Auras applied" toggle. Web only; the parquet has
  the lines.

## 5. Per-ability graphs on the main chart

### Web (no engine change)

- Every ability row in an opened actor row has an "On the chart" control. It measures that
  ability's effective amount per second from the fight's events (a measure on the shared
  query layer, scoped to the actor and spell over the whole fight) and draws it as an extra
  line on the main chart, in the ability's school colour, labelled in the legend with the
  actor and ability. One ability at a time; picking another replaces it; the control reads
  "Off the chart" while it is up.
- The line respects the brush like the main series. The night has no main chart of this
  kind and the control is absent there.
- The url does not carry it; it is a look, not a view.

## 6. Phases, from curated per-boss tables

### The data

The mechanics table for an encounter gains `phases`: an ordered list of `{ "name", "starts"
}` where `starts` is one of `{ "spell_id", "on": "cast_success" | "aura_applied" |
"aura_removed" }` on any enemy, or `{ "health_pct": N }` on the boss's own health, read from
the advanced block on its lines. Phase 1 starts at the pull and needs no entry. Curated
the way mechanics rows are: the drafting tool lists, per encounter, the enemy casts and
aura changes that happen exactly once per pull at a consistent boss health, as candidates.

### Engine

The summary gains `phases: [{ "name", "start_ms", "end_ms" }]` per fight, each phase from
its trigger to the next trigger or the fight's end; an encounter with no table or no
phases emits an empty list. Version 0.4.0.

### Web

- Phase bands on the main chart, named at their left edge; a phase band is a window
  preset ("Phase 2 · 1:12 to 2:40") in the presets strip, so every measured table can be
  read per phase with one click.
- The fight header's outcome adds the phase reached on a wipe ("Wipe 59% · in Phase 3").
- Compare aligns by phase when both pulls have phases: a "Phase" picker beside the window,
  which sets each side's window to its own phase's span.
- The night folds phase reached into the fight list and nothing else.
- Fights without phases render exactly as today.

## Testing

- Engine: table tests per feature on the fixture accumulator (series sums to the pair's
  total; at-max and wasted; owner guid; phase spans from each trigger kind), goldens
  refreshed once at 0.4.0.
- Web: unit tests for the window measure over threat series, the Compare diff builders,
  the resource readouts and the phase presets; e2e on the fixture for each panel at desktop
  and phone; the measured-ability chart under the DuckDB route helper.
- Harness: one smoke per feature against the sample log, reconciling the new figures with
  the parquet in DuckDB, the way the swing and absorb bugs were found.
- Review: the five personas run once after all six have merged, not after each.
