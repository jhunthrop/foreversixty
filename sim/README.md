# sim

The simulator's Go module: the request and result envelopes, the engine
pin, the two ends of the protobuf boundary, the two artifacts, and the
log-measurement harness.

| Package | What |
|---|---|
| `api` | `SimRequest` / `SimResult`, JSON, mirrored in `web/src/lib/sim/types.ts` |
| `enginever` | the pinned engine sha, written only by `make engine-pin` |
| `request` | our envelope to the engine's `RaidSimRequest` |
| `request/apl` | the launch specs' rotations, written only by `make apl-sync` |
| `adapter` | the engine's `RaidSimResult` to a `logs` summary |
| `combine` | split a run across workers, put it back together |
| `measure` | recover combat constants from a real combat log |
| `cmd/wasm` | `sim.wasm` + `sim.js`, four JSON exports |
| `cmd/forever-sim` | the native binary for the server lane |
| `cmd/forever-measure` | the measurement tool |
| `internal/simdb` | the active build's item database, embedded in both artifacts |

Build both artifacts: `make artifacts` from the repository root.

**After a fresh clone, run `make simdb` once.** `internal/simdb` embeds
`internal/simdb/simdb.bin`, which is a git-ignored copy of
`data/builds/<build>/simdb.bin` for the build named by
`web/src/data/active-build.json`; `//go:embed` resolves at compile time, so
without it every `go build`, `go test` and `go vet` in this module fails
with `pattern simdb.bin: no matching files found`. Neither artifact is
built `--tags=with_db`: that tag carries the engine's vanilla item table,
and Forever re-itemises.

## The two targets that cross into the engine fork

They do different things, and one of them was believed to do both, which
is how two copies of a rotation drifted apart.

`make engine-pin` writes `enginever/version.go` and nothing else. It
reads the fork's HEAD - refusing a dirty checkout, because a sha that
names a dirty tree names nothing - and writes the short sha into that one
generated file. It never writes into the fork, and it does not carry
rotations, presets or tables across.

`make apl-sync` carries the rotations. For every
`data/curated/apl/<spec>.json` marked `"state": "written"`, it writes the
file's `rotation` block into two places: the fork's
`ui/<class>/apls/forever_<spec>.apl.json`, which the fork's own spec
tests run, and `request/apl/<spec>.apl.json`, which both artifacts embed.
The curated file is the single place a rotation is edited; both of those
are copies, and editing either one directly means the sim measures a
rotation nobody wrote down.

The copy is the curated bytes, dedented one level, rather than a
re-print of the parsed JSON: the curated layout keeps a short object on
one line and expands a long one, and re-printing would churn the fork's
tree on formatting alone. Syncing an unchanged rotation writes nothing.

`make apl-check` proves both copies against the curated source - a
missing copy, a drifted copy, and a fork copy with no curated source all
fail - so forgetting `apl-sync` is caught in CI rather than in a golden.

## Measuring combat constants

The client tables do not carry whether periodic damage crits, the
multiplier a critical tick uses, how unified Hit interacts with the
weapon-skill miss table, proc chances, or internal cooldowns. All of
them are visible in a combat log. `forever-measure` reads one and prints
them.

```
forever-measure -log Logs/WoWCombatLog.txt -actor "Yourname-Beta"
```

Add `-json` to feed it into something. Add `-min-samples N` to see what
the tool is withholding and why; anything reading `insufficient data`
needs a longer run, not a smaller floor. `-spell-power` and
`-attack-power` **override** what the log reports and are almost never
needed: Forever's log writes the acting unit's attack power, spell power,
armour and level into the advanced block of every damage line, so the
character sheet comes out of the log.

**Turn advanced logging on before anything else.** `/combatlog` starts
and stops the log; advanced logging is the setting that fills that
nineteen-field block. Without it the tool has no character sheet, no
target level, and no attack table keyed by level, and it will say so.

Do **not** check this by reading the header: the client writes one
header when it loads, with `ADVANCED_LOG_ENABLED,0`, and another when
`/combatlog` starts, with `1`, and the first one is not about your file.
Check a damage line instead — it has nineteen fields between the unit
block and the damage numbers when advanced logging is on:

    grep -m1 SPELL_DAMAGE Logs/WoWCombatLog.txt | tr ',' '\n' | wc -l

A spell-damage line is about 33 fields with advanced logging and about
14 without. `forever-measure` says the same thing in words in its
footer, from `Advanced.OK` per event rather than from the header.

**The log has to be made deliberately.** A raid log mixes buffs, targets
and levels, and every one of those is a variable the measurement cannot
control for. Twenty minutes on a dummy gives a cleaner answer than a
night of raiding.

### Before you log

1. **Strip every buff.** No food, no flask, no elixirs, no scrolls, no
   raid buffs, no world buffs, no procs from another player. Right-click
   off anything that survives. A buff you forgot moves crit by a percent
   and the fitted constant is then wrong by a percent forever. The log
   records your attack power and spell power, but not which buff
   supplied them.
2. **Take off trinkets and any proc weapon** for the attack-table run.
   They are measured separately in run 3.
3. **Talents**: unspend anything that changes hit, crit, expertise or
   damage. If you cannot respec, write down what you have; an open-world
   log carries no `COMBATANT_INFO`, so the tool cannot know.
4. **Screenshot the character sheet anyway.** The log gives attack power,
   spell power and armour; it does not give hit, crit, expertise or
   weapon skill, and those are exactly what an attack-table fit needs.

### The three runs

Use `/combatlog` to start and stop. One file per run is easiest; the
tool takes one log at a time.

**Run 1 — the attack table (the most valuable one, and the one that
needs an instance).**
Target: a **level 63** dummy or boss, which is the case every melee spec
cares about. **An open-world log cannot do this**: the starting zones
have nothing at level 63, and the suppression terms differ per level, so
a level-60 run does not substitute. Until a dungeon is open, log what you
can at whatever level and **say which level**; the tool keys every row by
the target's level, read from the `SWING_DAMAGE_LANDED` advanced block.

- Auto-attack only. **No abilities at all** — a special uses a different
  table and mixing them makes both unreadable.
- **At least 1,000 swings.** At a 2.5-second weapon that is about 42
  minutes; at dual-wield 1.8s it is about 15 minutes for both hands.
  Fewer than 500 and the dodge and parry rates, which are small numbers,
  have error bars wider than the thing being measured.
- Stand **behind** the target if you can, for a run with no parry, then
  **in front** for a run with parry. Two files.

**Run 2 — periodic crits. This one works today, in the open world.**
- Apply your class's damage-over-time spells and **only** those. Let
  each run its full duration, re-apply, repeat.
- **At least 300 ticks per spell.** A dot ticking every 3 seconds for 18
  seconds gives 6 ticks per cast, so that is 50 casts.
- Include **one physical dot and one magic dot** if your class has both
  (Warrior Rend and Deep Wounds; Warlock Corruption and a bleed from a
  pet). Whether the two behave the same is exactly the open question, and
  it is the question this project can answer first.

**Run 3 — procs and coefficients. Also works today.**
- Put the trinket or weapon back on. Auto-attack for **at least 1,000
  swings**, or cast one spell repeatedly for at least 500 casts if the
  proc is cast-triggered.
- For coefficients: cast **one rank of one spell** at least 200 times.
  The log records your spell power on every line, so you do not have to
  write it down — but **do** drop a piece of spell-power gear and cast
  the same spell again. Two points fit a line, one does not, and the log
  will show the sheet changing between them.

### Reading the result

What each block answers:

- **PERIODIC DAMAGE** — `can crit` false across hundreds of ticks is
  evidence periodic damage does not crit for that spell. `mult` is the
  critical-tick multiplier, which the engine currently has no
  per-periodic value for. This is what Task 16's per-spell `CanCrit`
  flags are set from.
- **ATTACK TABLE** — these five rates, against a known target level and
  your character sheet, are what fit the miss, dodge, parry, glance and
  crit constants, and they are the only public evidence for how unified
  Hit meets weapon skill. A row against a level-60 target is a different
  measurement from one against a 63 and the table keeps them apart. A
  row reading `unknown level` is a target no line in the log gave a
  level for; it is kept separate rather than folded in with a known one.
  `special` counts only **physical** abilities, because a Frostbolt does
  not roll on the weapon table and would dilute every dodge rate in it.
- **PROCS** — `per swing` is the proc rate. `min gap` is a lower bound
  on the internal cooldown, and a tight one once the proc has fired
  thirty or forty times: a real ICD shows as a hard floor many samples
  never cross.
- **DAMAGE** — `mean` and `stddev` against the spell power the log
  reports are what a coefficient is fitted from. The base damage is not
  subtracted here; that is what the generated constants file is for.
- **The footer** lists what this log could not support. An open-world log
  always lists at least the missing `ENCOUNTER_START` and the missing
  level-63 target.

**Post the numbers with their sample counts.** A figure without its
count cannot be weighed against a later one.
