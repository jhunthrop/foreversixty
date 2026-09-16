# Mechanics mode

Status: draft for review. Written 2026-09-16 during the report review loop, after the
raid leader persona asked for avoidable-damage analysis in every round.

## What it is for

The question a raid leader asks after a wipe is not "who did the most damage" but "who
stood in what, and what did it cost". Warcraft Logs answers it with a per-fight
**Problems** list ("Hobolol took avoidable damage from Wicked Gash", "Reglitch died to
Gloom Squall") driven by a per-encounter table of which abilities are avoidable. The
table is the whole feature: the log only says an ability hit someone; a person decided
that standing in it was a mistake.

Mechanics mode is that list for Forever Sixty, plus the two things the review loop
showed people want beside it: a per-player card ("what to tell Mishvamp before next
week") and a whole-night roll-up ("Gloom Squall hit someone on 6 of 18 pulls").

## Where the tables come from

Forever's client derives from Classic Era (build 1.15.9), whose Dungeon Journal table is
empty, so Blizzard's own per-boss ability list is **not** available to the pipeline. The
tables are curated files, and the work is making curation cheap:

1. **A drafting tool reads a report's events** and, per encounter, lists every enemy
   ability that hit a player with the evidence a curator needs: who it hit (tanks, or
   others), how many players per cast, how often the same player was hit twice, its
   school, its share of all damage taken, and whether it killed anyone. An ability that
   hits several non-tanks at once, or the same non-tank repeatedly, is drafted as
   avoidable; an ability that only ever hits the tank is drafted as unavoidable; the rest
   are drafted as unclassified.
2. **Boss-mod modules on GitHub** (BigWigs, DBM; open source) are the second reference
   for the curator: they encode, per spell id, whether a mechanic is "move away", "soak",
   "interrupt" or "dispel". They are code, not data, so they are read by a person, not
   the pipeline.
3. **The curated file** is what ships. One file per encounter under
   `data/curated/mechanics/<encounter_id>.json`, checked in, reviewed like code. It is
   the draft with the curator's corrections; the draft never ships unreviewed.

Until Forever's raids open (2026-12-09) and someone logs them, the tables cover the
encounters in the sample log (Sanguine Depths), so the whole feature is built and reviewed
against real pulls now. Forever's bosses slot in as files when the logs exist.

### Table format

```json
{
  "encounter_id": 2363,
  "name": "General Kaal",
  "mechanics": [
    { "spell_id": 322936, "name": "Wicked Gash", "kind": "avoidable", "note": "Frontal cleave; only the tank should be in it." },
    { "spell_id": 323845, "name": "Gloom Squall", "kind": "avoidable", "note": "Room-wide unless behind a pillar." },
    { "spell_id": 322935, "name": "Piercing Blur", "kind": "unavoidable", "note": "Tank damage." },
    { "spell_id": 322938, "name": "Wicked Rush", "kind": "interrupt", "note": "Kickable." },
    { "spell_id": 320788, "name": "Frozen Binds", "kind": "dispel", "note": "Magic; Detox or Purify Spirit." }
  ]
}
```

`kind` is one of `avoidable` (damage a player should not have taken), `unavoidable`
(damage the fight deals regardless; listed so it is not mistaken for a gap in the table),
`interrupt` (a cast that should have been stopped), `dispel` (a debuff that should have
been removed). Soaks and positioning checks that need per-mechanic logic are out of
scope for the first cut; the format leaves room (`kind` is a string, and a mechanic can
carry extra fields later).

## What the engine computes

Per fight, per player, per listed mechanic, from the events the engine already reads:

- `avoidable`: hits taken and damage taken, with the first and last hit's instant, and
  whether one of them was the killing blow.
- `interrupt`: how many casts of the spell started, how many were stopped, and by whom
  (the engine already keeps this for the Interrupts tab; Mechanics mode lists the ones
  the table names).
- `dispel`: how many applications on players, how many were dispelled, and how long the
  rest ran (already partly computed for the Dispels tab's "ran their course").

This lands in `summary.json` as a `mechanics` block keyed by spell id, so the web needs
no events to draw the mode, and the whole-night fold sums it. The engine reads the
curated table for the fight's encounter at parse time; a fight with no table gets an
empty block and the mode says so.

The engine version bumps, and the sample report is re-parsed, as with every summary
change.

## What the web shows

**Mechanics** becomes an enabled mode (it is already listed as "later" in the mode bar).

1. **Problems, this fight.** One line per failure, ranked by cost: avoidable damage
   taken by player and mechanic, with hits and damage and "and died to it" where the
   killing blow was that mechanic; interrupts that went through; debuffs that ran their
   course. Each line links to the tab that shows the detail (Damage Taken filtered to the
   ability, the Deaths card, Interrupts, Dispels).
2. **Per player.** A card per player with their avoidable damage total and its share of
   their damage taken, the mechanics they were hit by, and the ones they were never hit
   by (the raid leader's "what to tell each player").
3. **The night.** On `?fight=all`, per mechanic: how many pulls it hit someone on, who
   most often, and the pull list; per player, avoidable damage over the night.
4. **Honesty.** A fight whose encounter has no table says "No mechanics table for this
   boss yet" and offers the drafting tool's output for that fight, so a curator can add
   one. Unclassified abilities are listed under "not yet classified" rather than hidden.

Everything measured here is whole-fight: the mode ignores the Analyze window, like
Compare does, and says so.

## What is out of scope

- Soak and positioning checks (needs per-mechanic logic).
- Comparing a player's mechanic record against other reports (needs the API).
- Automatic curation without review.

## Testing

- Pipeline: the drafting tool on the sample log produces the expected draft for General
  Kaal (unit test on a fixture of events).
- Engine: a fixture fight with a table yields the expected `mechanics` block (golden).
- Web: the mode renders the problems list, the cards and the night roll-up from the
  fixture; a fight without a table shows the empty state; every link lands on the right
  tab with the right filter (Playwright).
- Review: the raid leader and tank personas review the mode on the sample log.
