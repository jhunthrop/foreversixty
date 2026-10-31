# Review: wt/logs-engine-counts (logs engine 0.6.0 kill counts)

Reviewed 2026-10-30 against `origin/main...HEAD`. Read only; nothing edited or committed.

## Verdict: FIX FIRST

There is one HIGH: neutral mobs the group kills are dropped from both counts. The fix is
small and stays inside `logs/engine/fight/counts.go` plus its tests. Everything else (the
migration, the rollout, the readers, the web fallbacks) is safe to ship before the Nov 3
freeze.

## Checks run

- `go vet ./...` in `logs/` and `api/`: clean.
- `go test ./...` in `logs/`: all packages pass.
- `go test -p 1` on api `internal/engine`, `internal/reports`, `internal/parse`, `internal/db`,
  `internal/coach`, `cmd/samplecut` and `cmd/seedguild`, against a private database
  (`forever_gc_cr`, since dropped): all pass.
- The 0059 down and then up migration, run by hand on a migrated database: both apply cleanly.
- A one-off scan of the sample report's `events.parquet` files: 675 distinct hostile NPC deaths
  and 86 neutral ones. 3 of the neutral deaths have a PARTY_KILL (`Galestrider`), and the
  engine counts none of the 86.

## CRITICAL

None.

## HIGH

### H1. Neutral mobs the group kills count as neither a mob death nor a group kill
`logs/engine/fight/counts.go:112` (`isMob` requires `units.Hostile`), used at `:87` (the
PartyKill case) and `:102` (observeDeath). The tests lock this in at
`logs/engine/fight/counts_test.go:18,104-105` (0xa28 is treated as "a critter") and in the
oracle at `logs/engine/session/counts_test.go:34`.

The combat log flags a neutral (yellow) mob as REACTION_NEUTRAL (0x20) for its whole life,
including while it is fighting you and on its UNIT_DIED and PARTY_KILL lines. The reaction
bit cannot tell a critter from a neutral beast. Many real mobs are neutral: plainstriders,
kodos, boars and zhevras in the starting and leveling zones. The Forever sample alone has 3
`Galestrider` PARTY_KILLs that drop out. The retail v16 fixture has PARTY_KILLs on real
enemies (`Tormented Soul` x24, `Lingering Manifestation` x4) that also drop out.

Engine 0.5 counted every PARTY_KILL, so this is a regression for `group_kills`. Launch week
is mostly leveling and dungeon logs. A pull of neutral beasts would show "No mobs died", or
"N mobs died · your group killed 0", which is the kind of wrong number this lane exists to
remove.

Fix: treat a non-player NPC that is hostile, **or** neutral and touched by the group, as a
mob:
- A PARTY_KILL on a neutral NPC counts toward both `mob_deaths` and `group_kills`.
- Track group damage on neutral NPCs as well. At `:80-82`, record `lastHitByGroup` when the
  dest is hostile or neutral. In observeDeath, count a neutral death only if the group
  damaged that unit, so critters that outsiders or environmental damage killed stay out.
- To drop critters the group kills by accident with AoE: either skip units whose advanced
  max HP is below a small threshold, or use a critter NPC-ID set built from wago (creature
  type 8).
- Change the test's `flagNeutral` case to two cases: a neutral beast the group killed (counts)
  and a critter (does not). Update the session oracle to match.

## MEDIUM

### M1. Player deaths and group_deaths are not deduplicated, and may count Feign Death
`logs/engine/fight/counts.go:96-99`. Every UNIT_DIED for a player adds one. That is correct
for a real second death (Soulstone, Reincarnation, a battle res). If the client logs a
hunter's Feign Death as UNIT_DIED, though, each feign also adds one to `deaths` and the new
`group_deaths`. `group_deaths` now drives the fight row's death mark and the subtitle
(`web/src/components/report/ReportView.svelte:297`). I could not confirm this either way:
there is no Forever hunter log in the fixtures. `deaths` behaved the same way before this
branch, so only `group_deaths` is new exposure.
Fix: check against a Forever log with a hunter who feigns. If feigns do log UNIT_DIED, skip a
player death when that player has an active Feign Death aura, or when the same GUID acts
again within a few seconds without a Resurrect event.

### M2. A unit a group member has charmed or mind-controlled is not counted when it dies charmed
`counts.go:112`. While a priest's Mind Control is up, the creature carries friendly flags
with MINE, PARTY or RAID affiliation. If it dies in that state, `isMob` rejects it. Credit in
the other direction is correct: damage dealt by the charmed unit has group affiliation, so
its killing blows count as the group's. This is rare on trash, so it is MEDIUM.
Fix: remember the creature GUIDs seen with hostile flags earlier in the fight, and accept
them at death even when their flags are now friendly.

## LOW

### L1. During the rollout, an old revision rewrites report.json without the new fields
`api/internal/reports/ingest.go:669-683` (WriteReportJSON). While the old revision still
takes traffic, any fight upload it handles rewrites `reports/<id>/report.json` from the old
`FightEntry`, which drops `mob_deaths`, `group_kills`, `group_deaths` and `group` for every
fight in that report. This lasts only until the next upload on the new revision. The web then
shows no count, never a wrong one, so this is cosmetic.
Fix: none needed. Optionally drain the old revision quickly.

### L2. An old revision's upsert leaves stale new columns behind
`api/internal/reports/store.go:276-283`. The old code's `on conflict do update` does not set
the four new columns. If a fight the new revision wrote is re-sent with a different SHA and
lands on an old revision, the old counts survive next to the new `players` and `deaths`. This
needs a re-send during the rollout window. If you want to close it, have the follow-up
migration that drops `npc_kills` also null the counts on rows where `uploaded_at` falls in
the deploy window.

### L3. Fights indexed before 0.6.0 never get counts
Old rows stay null for good. A re-send with the same SHA returns early at `ingest.go:302-304`
and is never recounted. This is the intended behaviour (no count, never a wrong one), but the
production path has no backfill like `samplecut -recount`. Consider a follow-up command that
recounts from each fight's stored `events.parquet`, which is what recount.go already does
for the sample.

### L4. No migration test for 0059
`api/internal/db/migrations/0059_fight_counts.*.sql`. The repo has migration tests such as
`adherence_migration_test.go` (0037) and the 0052 guild attendance test. The migration itself
is safe: it adds nullable columns with no default, which Postgres does as a catalog-only
change, and `npc_kills` keeps `not null default 0` (0005), so new-code inserts that omit it
succeed. The down migration reverses it.
Fix: add a short test that up adds the four nullable columns and down removes them.

### L5. The scoped subtitle count and the scoped death list can disagree
`web/src/components/report/ReportView.svelte:297-298`. Under group scope the subtitle reads
`group_deaths` (UNIT_DIED lines whose victim was flagged group), but the death list below it
filters the summary's deaths by `groupScopeNow.group`. If the summary deduplicates deaths
differently, or the feign case in M1 is real, the two numbers differ on the same screen.
Fix: once M1 is settled, read both from one source.

## Areas checked with no findings

- **GUID deduplication.** `died` and `groupKilled` are keyed by GUID, so a PARTY_KILL plus a
  UNIT_DIED for one mob is one death, and so is a doubled UNIT_DIED. A mob that dies twice
  (fake death and reset) counts once, which is right for a kill count.
- **Pets and guardians.** A pet carries its owner's affiliation, so the `lastHitByGroup`
  fallback credits pet killing blows. Pets are kept out of `group` because only player GUIDs
  are recorded. Hostile summons die through UNIT_DESTROYED and UNIT_DISSIPATES, which are
  excluded.
- **Rollout and readers.** Every read is nil-safe. The store scans NULL into `*int` and nil
  slices, and the JSON uses `omitempty`. On the web, `killCounts` and `sumKillCounts` return
  null when any pull lacks counts, and `fightGroupScope` falls back to `groupScope`. The
  Swift and Kotlin models had `npcKills` optional and the new fields are optional too.
  Rolling back to the old API is safe: it reads `npc_kills = 0` on new rows, and the old web
  never displayed that field. The API does not gate uploads on engine version.
- **Performance.** `Counter.Observe` makes the same two `units.Parse` calls per event that the
  old `assign` loop made. What is new is one map write per Damage event plus small per-fight
  sets, which is negligible next to the parse even on a 40-player log. `Apply` sorts players
  and group, as `snapshotPlayers` already did.
- **Session restore.** `session.State` nils the open fight and replays it, so the
  `restoreCounter` double-count described in its comment cannot happen through the session.

## Re-review: fix round 1c26af802 (2026-10-30)

This is a scoped re-review of commit 1c26af802. It is read only. The only change is this
appended section.

### Verdict: FIX FIRST

H1 and M2 are fixed. The M1 Feign Death undo adds a new HIGH: on a real log it takes back
real deaths.

### Checks run

- `go test -race -count=1 ./...` in `logs/`: every package passes.
- I ran the `fights -json` output of the parent commit (4e221dcc6) and of 1c26af802 on all
  four fixtures (`v16`, `v22`, `v22-shapes`, `forever-1.60`). The fight lists are identical, so
  the "goldens unchanged" claim holds. The fixtures are too small to exercise the change,
  though: none has a neutral pull and none has a wipe.
- I ran the same comparison on the cached real retail log
  `logs/testdata/cache/wowp-retail-v16.txt` (Sanguine Depths, Necrotic Wake). It goes from 63
  fights to 64, and the counts differ in 17 fights. The details are below.

### HIGH

#### R1. The Feign Death undo erases real deaths: wipes with a corpse run, Reincarnation, and probably Soulstone
`logs/engine/fight/counts.go`: `signOfLife`, `undoDeath`, and the `Resurrect` case in `Observe`.

The undo assumes every way back from a real death logs `SPELL_RESURRECT`. Several common ones
log nothing of the kind:

- **Releasing and running back from the graveyard logs no resurrect event.** On the retail log,
  the General Kaal wipe (02:30:12 to 02:42:16, one encounter) has 9 player `UNIT_DIED` lines and
  0 `SPELL_RESURRECT`. Each player runs back and casts again before `ENCOUNTER_END`, so the
  fight goes from **deaths 9 to 0, and group_deaths 9 to 0**. Trash fights with a corpse run
  inside them lose deaths the same way: 5 to 2, 2 to 0, 1 to 0, 3 to 0, and 5 to 4.
- **Reincarnation logs as a `SPELL_CAST_SUCCESS` by the dead shaman** (spell 21169, line
  110445). It comes 3.5 s after the shaman's `UNIT_DIED` at 02:58:06.772, and no
  `SPELL_RESURRECT` follows. That trash fight drops from 1 death to 0. This is the "known gap"
  in the round-3 notes, and the log confirms it. Classic's Reincarnation (20608) is very likely
  logged the same way.
- **Soulstone** has no case in the fixtures. Treat it as at risk too.

Feigns are still unconfirmed, but these are confirmed real deaths. The effect is that a wipe
shows "0 deaths" on the fight row, and on the unscoped subtitle, which reads `fight.deaths`.
That is the kind of wrong number this lane exists to remove. Before this commit, `deaths`
counted these deaths correctly.

Fix, in order of preference:
1. Undo a held death only when the same player cast Feign Death (`SPELL_CAST_SUCCESS` or an
   aura, spell 5384 by name or id) at or just before the `UNIT_DIED`. Until then, keep every
   other player death.
2. Or drop the undo until a Forever hunter log shows whether feigns log `UNIT_DIED` at all.

Either way, add a regression test built from the Kaal sequence (death, no resurrect, cast:
the death stays) and one built from the Reincarnation sequence.

### MEDIUM

#### R2. Fight boundaries move on real logs, though the fixture goldens do not show it
`fight.go` `hostileCombat` now uses `units.Enemy`. On the retail log:
- An old 132 s "fight" from 01:54:34 to 01:56:46 disappears. It was ambient NPCs fighting each
  other: neutral `Grubby Dirtcruncher` against hostile `Gluttonous Tick`, which the old
  hostile-vs-non-hostile test counted as combat. With that event gone, the next trash fight
  splits into three pulls. **Both changes are improvements.**
- In the other direction, ambient fights between a friendly NPC and a neutral one now keep a
  fight open. In Necrotic Wake, `Xandria` (0xa18) against `Patchwerk Soldier` (0xa28) stretches
  the trash fight that starts at 03:24:02 from 03:24:33 to 03:26:37 (+2 min). This is the same
  class of problem the old rule had with friendly NPCs against hostile ones, so it is not new,
  but it now applies to neutral units too.
- Boss encounters are not split or merged, because encounter boundaries still win.
- Critters and a single aggro'd neutral: a one-hit critter kill makes a trash segment shorter
  than `MinTrash` (3 s), which is dropped. A real neutral pull in a levelling zone now becomes a
  fight, which is intended. A bystander killing yellow beasts within log range also opens fights
  now, as bystanders already did with red mobs, because `hostileCombat` does not require a
  group member on either side.

Suggestion: not a blocker. A follow-up could require a group-affiliated unit on one side,
which would remove both kinds of ambient sustain. The levelling-log fixture case is still
missing: add a small neutral-pull fixture so that a future change to the rule shows up in a
golden.

### H1: fixed
- A neutral NPC counts when the group damaged it or PARTY_KILLed it. A neutral NPC nobody in
  the group touched does not count.
- Critters are excluded by known max HP of 10 or less.
- Health that was never shown: in the Forever dialect, `SPELL_DAMAGE` and
  `SWING_DAMAGE_LANDED` carry the **dest's** advanced block (checked in `forever-1.60.log`), and
  `observeUnits` reads advanced blocks on every event kind. So a critter the group damages
  normally reveals its health and is excluded. A unit counts with unknown health only when
  advanced logging is off, which the upload check already warns about, or on a PARTY_KILL with
  no damage line. **This is acceptable.**
- The 10 HP cutoff is not checked against Forever critter data. Noted, not a finding.

### M2: fixed
`hostileSeen` accepts an NPC seen hostile on any earlier line of the fight. The test covers both
a charmed mob that was seen hostile (counts) and one never seen hostile (does not count).

### M1: see R1
A death that is never taken back simply stays counted at fight end, which is correct. Restoring
a session drops `pendingDeaths`, which only means a feign cannot be undone after a restore.
That is harmless.
