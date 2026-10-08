# The client's crit curves: PlayerExpectedStat (2026-10-08)

## What was found

Publish 23 (bis17, engine 6cc1e6fa2) moved every caster that is not a mage or
a shadow priest down 2.4 to 3.1% on identical gear. The cause was the raid
preset dropping Grace of Air Totem (ruled 2026-10-08: it shares the air totem
slot with Windfury and the beta notes say the two do not stack). Measured with
`spec-breakdown -preset raid -buffs grace_of_air_totem`, 5,000 iterations:

| spec | without | with Grace of Air | gain |
|---|---|---|---|
| warlock-affliction | 895.3 | 916.2 | +2.3% |
| druid-balance | 515.5 | 529.8 | +2.8% |
| shaman-elemental | 478.3 | 492.5 | +3.0% |

77 Agility moved a balance druid's Starfire crit from 28.3% to 31.7%: the
engine feeds Agility into Forever's one Crit stat for every hybrid, so a
caster's spells crit more with Agility. Mages and priests were the only
classes without an Agility line, which is why they did not move.

Chasing where the rates came from turned up a primary source nobody had read:
the beta client's DB2 table `PlayerExpectedStat` (wago.tools serves it for
1.60.1.70009 and 1.60.1.70245, identical rows), one row per class per level
1 to 123 with `BaseMana`, health per Stamina, `CritPerAgility` and
`SpellCritPerIntellect`. research/08-stats.md had recorded these curves as
"entirely unpublished" and the engine comment on `CritPerAgiAtLevel` said no
per-level source existed.

## What the table says

- Level 60 `CritPerAgility` equals the engine's level-60 constant for all nine
  classes (warrior 0.0500% per point, paladin 0.0506, hunter 0.0189, rogue
  0.0345, priest 0.0500, shaman 0.0508, mage 0.0514, warlock 0.0500, druid
  0.0500). The engine was right at 60.
- Below 60 the client's rate is higher: a level-30 warrior gets 1% per 10.4
  Agility (the engine gave 1% per 20), a level-1 rogue 1% per 2.3. Every
  leveling band under 60 under-credited Agility crit, by about 2× at 30.
- `SpellCritPerIntellect` equals the per-level table the engine already had
  from Wowhead's planner, at every level checked. Wowhead mined this table.
- `BaseMana` equals `basemp.txt` at every level checked; health per Stamina is
  10 everywhere.
- Warriors and rogues have a zero Intellect column. Every class, mages and
  priests included, has an Agility column.

## What changed

- Data: `pipeline gametables` fetches `PlayerExpectedStat` beside the three
  game tables and commits it as `gametables/playerexpectedstat.csv`
  (1.60.1.69893 and 1.60.1.70009, identical bytes); the fetch fails loudly if a
  later build drops either crit column.
- Engine (lane expected-stat): `CritPerAgiAtLevel` becomes a per-level
  function generated from the client CSV; `SpellCritPerIntAtLevel` is
  generated from the same CSV with a check against Wowhead's copy; every
  class wires both lines (the table's zeros do the rest), so mages and priests
  now get the tiny Agility crit the client gives them; the paladin's
  Agility-to-dodge line, which read the crit constant, reads the dodge one.

## What stays open: do a hybrid's two crit terms sum?

Blizzard confirmed one Crit stat for spells, melee and ranged. The table gives
every class both an Agility rate and an Intellect rate. Whether the server adds
both into the one number (the engine's model, which is why Grace of Air lifts a
druid's Starfire), or keeps two pools and shares only item, talent and racial
crit between them (foreverdb.net's reading, and Output Lag's), is not in any
table: the roll lives in the server. It is worth 2 to 3% of a caster's raid
DPS, the whole value of Agility on caster gear and of Intellect on hunter,
retribution and enhancement gear.

The beta character sheet decides it in under a minute; the test is
design/reviews/2026-10-08-beta-evidence.md §9. Until then the engine keeps
the sum, flagged in `AddCritStatDependencies`.
