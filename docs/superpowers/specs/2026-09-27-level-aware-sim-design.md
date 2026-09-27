# Level-aware simulation

**Date:** 2026-09-27. **Status:** approved ("fix them"). Owner asked why the sim only runs at 60 and whether it can know what spells a character has.

## Today

Nothing in the game pins the sim to 60; four pins of ours do.

| Pin | Where |
|---|---|
| The addon export carries no level | `addon/ForeverSixty/Codec.lua` (FS1 has class, race, talents, gear, optional sections) |
| The site sends `character.level = 60` always | `web/src/lib/sim/character.ts` `SIM_LEVEL`; `sim/api/envelope.go` refuses any other value |
| The engine builds every character at `CharacterMaxLevel` | `.engine`/fork `sim/core/character.go`; `proto.Player` has no level field |
| Rotations name max-rank spell ids | `data/curated/apl/*.json`, embedded in `sim/request/apl/` |

What already works: every class spell registers a rank only when `config.RequiredLevel <= character.Level` (Season of Discovery heritage, 61 such gates), the spell constants we ship carry every rank's learn level (`spell_level`), items carry `required_level`, and the game tables we fetch (`basemp`, `hppersta`) are per level. What does not: 18 class files pick rank numbers from maps keyed on the SoD brackets 25/40/50/60 (empty for any other level); base strength/agility/intellect/spirit/stamina exist only as the level-60 table.

## Sources for what is missing

Wowhead's Forever gear planner payload (`builds/<build>/raw/wowhead-gear-planner.js`, already fetched per build) carries, per class, per level 1..60: health, mana, agility, strength, intellect, spirit, stamina (`baseStats.stats[class][statId][level]`, statId 0 mana, 1 health, 3 agi, 4 str, 5 int, 6 spi, 7 sta; level-60 values match the engine's tables to the point), race offsets (`baseStats.raceOffsets`), and spell crit per intellect per level (`critSpell[class][level]`). Physical crit per agility per level is not in it: the engine keeps its level-60 `CritPerAgiAtLevel` and marks it unconfirmed below 60. Spell learn levels come from our own `spellconst`.

## Design

1. **Data** (`data/`): `normalize` also emits `builds/<build>/levels.json` from the wowhead payload (base stats per class per level, race offsets, spell crit per int per level) and `builds/<build>/spellranks.json` from spellconst: per class, spell name -> ranks `[{id, rank, level}]` sorted by rank, only spells with more than one id or a learn level. Both fixture-tested. The engine's generator reads `levels.json`; `sim/request` embeds `spellranks.json`.
2. **Engine core** (fork `jhunthrop/wowsims-forever`, branch `forever`): `proto.Player.level` (0 means max); `NewCharacter` builds at that level; base stats, health and mana per level from a generated table (tool reads `levels.json`), attack power by level per class from the Classic formulas the level-60 constants encode (warrior/paladin `3*level-20`, hunter/rogue `2*level-20` melee, etc.: derive from the existing 60 constants, do not invent), spell crit per int per level from the table; hunter pet level follows the owner. `HighestRankAtLevel(learnLevels, level)` is in `sim/core/level_ranks.go` (landed ahead of the lanes).
3. **Engine classes**: each of the 18 bracket-keyed files takes its rank from `HighestRankAtLevel` over that spell's real learn levels (from `spellconst`), and the spell's per-rank data is indexed by that rank. A level with no rank learned registers nothing for that spell.
4. **Rotation ranks** (`sim/request`, Go, used by the API and the browser wasm alike): before a rotation is handed to the engine, every `spellId` in it whose id is a ranked spell is rewritten to the highest rank the character's level has learned, by spell name within the class, using the embedded `spellranks.json`. A spell with no rank learned at that level has its action dropped (the APL already drops nil operands). Level 60 rewrites to itself.
5. **Request and validation** (`sim/api/envelope.go`, `sim/request/request.go`): `character.level` accepts 1..60; the player proto carries it; the default target level is the character's level plus 3 when the request names none, with the target armour table extended by level (`web/src/lib/sim/settings.ts` `TARGET_ARMOR_BY_LEVEL` and its Go twin).
6. **Transport** (addon + web): the export gains `level=<n>` (an optional FS1 section, `UnitLevel("player")`); the site's decoder reads it; `toCharacterSpec` sends the export's level (a hand-built planner build still sends 60); the strip shows the real level; the sim scope copy stops saying "at level 60" and says the level it ran at.

## Order

Data, engine core, engine classes, transport and request lanes run in parallel; only step 5's one line that copies the level into the player proto waits for the engine core to land and the pin (`sim/enginever/version.go`, `make engine-pin`) to move. The pin moves once, after core and classes merge on the fork's `forever` branch and its `go test ./...` is green.

## Out of scope

Stat weights per level (Top Gear runs weights live), healing and tanking specs, and per-level physical crit constants beyond the marked approximation.

## Global constraints

- Fork: `gofmt`, `go vet ./...`, `go test ./...` green. Repo: data `ruff` + `pytest` (80%), web `npm test`, Go `go test ./...` in `sim/` and `api/`.
- Nothing branches on a build string; every new table is generated from data with its source named.
- No lane merges or pushes; each reports and stops.
