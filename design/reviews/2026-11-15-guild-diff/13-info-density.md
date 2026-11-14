# Information density gate: four guild specs (2026-11-15)

Judged without screenshots, against the specs and the code: api/internal/guilds (loot_council_*.go, raids_report_pulls.go,
raid_prep.go, readiness.go), logs/engine/summary (deaths.go, phases.go, mechanics.go), data/builds/*/loot.json, addon
RaidPrep.lua and Comms.lua. Privacy rule held throughout: names officer-only, raiders get their own rows, no directory gear
without D1. Nothing below asks for a name to reach a raider, or for a number the data lacks.

## Verdicts

| Spec | Verdict |
|---|---|
| loot-plan | SHIP WITH FIXES |
| live-pulls | SHIP WITH FIXES |
| prep-reminder | SHIP WITH FIXES |
| guild-later | SHIP WITH FIXES (briefs only; recruiting stays blocked on G1 and R1) |
None is THIN: each has a real answer above the fold. The gaps are fields we already have.

## 1. loot-plan

Strong: role groups, bands, fairness line, source tag with age, officer actions, redaction in the API.
Must fix:
1. Section 5.1 drop row and head: add drop chance. loot.json bosses carry `item_chances` per item (for example 1.0, 0.5, 0.02).
   Show it as `100%` / `2%` on the drop ItemRow and the raider's `Your drops tonight` row. A raider asks "will it even drop"; Onyxia
   has 22 drops and the left Plate shows none of that.
2. Section 5.2 raider row: add `replaces <item>` (what you wear in the slot, quality colour) and the item level delta. The officer row
   has `replaces`; the raider row has only a status word, which is the one place the raider needs it. Own data only, so privacy holds.
3. Section 5.1 left Plate: the per-role top pick is hidden until expanded. Show the count per role on each row (`D 11 · H 4 · T 2`,
   counts only, no names) so an officer sees which drops have no taker without opening 22 rows. On phone this is a second line in the
   drop picker (SelectSheet).
4. Section 5.1 row: add the council vote tally when the drop is already open on a night (`votes_yes/votes_no`, loot_council_nights.go)
   and `Won last by <name>, <date>` is NOT asked for (names). Cut nothing, but link each row to the existing LootCandidate detail
   (drill-down) so the plan and the council are one path.
6. Group head: add a Copy CSV button (officer only) for the group, as a log site table has. Phone: in the overflow Sheet.
7. Section 5.4 partial data: the count line shows `4 of 11 have gear on file`; add the per-source split (`2 addon · 1 Blizzard`)
   so an officer knows which gap to chase. Counts only; names stay in the officer Disclosure.

## 2. live-pulls

Strong: cost model ranked in one unit, names matrix, latency table with honest estimates, one shared PullCosts component.
Must fix:
1. Section 5.1 region B: each pull row shows only cost lines. Add the phase of every death and the phase the pull ended in
   (`ended in phase 3`), from logs/engine/summary/phases.go (name, start_ms, end_ms). A raid leader needs "where do we die".
2. Section 5.1 region B: add a one-line trend strip over the night: boss % at end per pull (`72 · 55 · 38 · 31`), best pull marked,
   from the fight summaries already folded. A wipe at 31% with no neighbour is a figure with no context. Phone: the strip is the
   first line under the head, a 4-wide wrap.
3. Section 2 cost lines: add the other mechanic rows the engine already builds (logs/engine/summary/mechanics.go: `casts`,
   `stopped`, `applied`, `dispelled`, `damage`, `healed`, per-player `Hits`). Interrupts stopped vs missed, dispels, avoidable hits
   are mechanics, not parses, and officers get names; raiders get counts only. Today the spec reads deaths, downtime and rotation only.
4. Section 5.1 death row: show last hits and overkill (`Death.last_hits`, `overkill`, `hp_after`, `max_hp`) in the row's expand
   (`Fireball 2,140, then Fireball 1,980`), and auras held or lost at death (deaths.go). It is the Details! death log this spec cites,
   and the data is already in PullDeathView.
5. Section 3 and 5.1 raider view: the raider's own lines are specified only as `own` flag. Specify them: own deaths with cause and
   phase, own active time against the fight length (metrics row), own damage taken, and the same trend strip. Own row only.
6. Section 5.1 region C rail: Night so far and Companion is a thin rail. Add `Deaths by phase` and, for officers, `Who died more than
   once tonight` as a count of raiders (not names) with names under officer Disclosure; plus the logging character's last upload
   age (already the `Stale` copy) as a permanent muted stamp.
7. Section 5.2 report: add Copy CSV of pulls (pull, result, boss %, duration, deaths, top cause), a sort on Deaths and Duration,
   and link each cause to its encounter mechanic. The Plate has death causes today and nothing sortable.
## 3. prep-reminder

Strong: sources and ages on every row, `Could not check` reasons that read as facts, no shame list, relay honesty.
Must fix:
1. Section 4.3 member rows are five (enchants, resistance, flask, food, attunement). The data has more: `talent_points_unspent`
   (readiness.go TalentPointsUnspent, with `talent_points_read`), `gear_gap` and `item_level_delta`, `level`, weapon enchant,
   protection potion counts by school (snapshot supplies), and per-boss `encounterItems` (snapshot! v2). Add Unspent talent points,
   Weapon enchant, Protection potions (count by school) and a `Gear gap` line to the member block. Gold only when known to fail.
2. Resistance row: `Fire 21 from gear` has no target. Show the raid's need per boss only where the data has it (encounters.json
   mechanics); otherwise print the total and nothing more. See needs data.
3. Section 4.3 officer block: add raid-level coverage tiles, counts only: `Flasks 22 of 31 checked`, `Food 28`, `Enchants 25`,
   `Attuned 36 of 38`. Same data as the rows, the line an officer actually reads at 7:00. Phone: a 2 by 2 wrap under the count line.
4. Section 4.3 officer rows: sort control (fixes first, name, class) and a class filter on a 38-row list; Copy as text for pasting
   into the officer channel. Phone: sort in the overflow.
5. Section 4.1 DM: add the raid's `Attuned` line (the readiness Attuned view) and a `Where` per fix (the ready line's where label,
   already in ready_check.go ReadyItem). The DM names the fix but not where to fix it; the addon panel does.
7. Section 4.4 addon raider panel: add item level and delta to your guild median (`item_level_delta`) as a header figure.

## 4. guild-later

Comp and bench: the shape is one line plus a bench list. Must fix:
1. Section 1 Shape: each bench row also lists what the player brings that nobody else provides AND what the raid gains by fielding
   them, both from raids_cover.go providers (`sure`, `likely`); today only the cost of benching is named. Add `Cover after swap`
   as a count line (`Buffs 14 of 16, debuffs 7 of 8`) before and after, same source, no DPS needed. Ships without K2.
2. Section 1: show the coverage count at the head of the Grid (`16 of 24 covered`) as a figure, not only a `Missing:` line.
3. Fix G2 first: the catalogue's `melee_group` need for Strength of Earth is stale per the spec itself; a density panel built on
   a wrong need is worse than none.
Recruiting search: must fix:
1. Section 2 row: add the item level (already listed in the answer) plus the equipped-slot count of rare-or-better pieces from the
   directory `equipment` summary (slot, quality, item level), and `Fills: Windfury Totem` computed from the guild's open cover
   gaps (raids_cover.go) for the character's class and spec. That is the "would fit us" answer and needs no new data.
2. ControlBar: sort by item level and by last seen, a result count, `Copy /w` per row (already specced) and a saved filter per guild.
Both stay blocked on G1 and R1.

## Needs data

- Resistance need per boss (Onyxia fire resistance threshold): not in encounters.json as a number; verify before any target is shown.
- Boss % progress per pull for a wipe: confirm the fight summary carries last boss health (the spec uses 31% already, so it should).
- Marginal raid DPS per bench player: engine spike K2, not built.
- Group assignments (per-group buffs): none exist.
- Consumables at pull: needs a real Forever log with COMBATANT_INFO (spec already holds this).