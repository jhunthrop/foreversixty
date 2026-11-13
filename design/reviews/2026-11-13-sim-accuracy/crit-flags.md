# Sim accuracy 2026-11-13: every class's periodic crit flags

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `crit-flags`. Site from origin/main 25f3696db, fork from master 79358b3c2 (the pin). Client on the site: 1.60.1.70790; the reference here is the live 1.60.1.70888 with every hotfix push (`useHotfixes=1`), checked again on 70938 (no Can't Crit or PERIODIC_CAN_CRIT bit moved on any row between the two builds). The 70938 data run reached main during the lane (52cd11091) but the site's active build is still 70790.

Why: the owner's 31 October rule ("bleeds and non-channelled DoTs always crit; channels follow the client PERIODIC_CAN_CRIT flag") went into the engine as a blanket rule (`core.periodicDamageCanCrit`, fork 09ca529f3). Deep Wound (412613) turned out to be flagged Can't Crit with zero crits in the logs (dps-warriors.md, 12 Nov) and the rogues were all aligned (rogues-60.md). This lane does the same check for every other class.

**Totals: 2 engine bugs, both old, fixed. (1) Ignite's ticks crit in the engine under the true-dot rule, where Forever's Ignite (talent 11119, aura 412538) burns through the server tick 412545, which carries SPELL_ATTR2_CANT_CRIT exactly as Deep Wound's 412613 does; the logs are silent (3 ticks, none critical). Fire raid 60 -0.9 to -1.1%, Fire bare 20 to 60 -0.05 to -0.3%, Arcane raid 60 -0.4 to -0.5%. (2) Blizzard, Hurricane, Rain of Fire and Volley never crit in the engine (unflagged channels), where each pulse is a separate direct-damage tick row with no Can't Crit, and the logs show those ticks crit: Blizzard 902 of 2,767, Rain of Fire 200 of 3,315, Volley 52 of 478 on the hunter's ranged table. No published rotation casts any of the four, so no published number moves. Every other periodic in every published rotation, card, proc and BiS item rolls crit exactly as its client row says, and every log check agrees. Engine pin 79358b3c2 to 8f61b9e6a.**

## Method

1. **Client.** `crit-flags/client.py` walks every class spell (SpellClassOptions), every talent rank and three hops of trigger spells in the hotfixed 70888 SpellEffect and SpellMisc, keeps every row with a periodic damage, leech or heal aura (3, 8, 53, 89), a periodic trigger (23, 226) or a direct-damage row reached through one, and prints SPELL_ATTR2_CANT_CRIT (Attributes_2 0x20000000), SPELL_ATTR8_PERIODIC_CAN_CRIT (Attributes_8 0x200) and the channel bits (`client_70888.tsv`, 586 rows). A second pass took every number in the fork's class and item code and listed every one the client marks Can't Crit with a damage or heal effect. Server tick rows (the targets of a 226 dummy aura, which the client does not link) were found by name in SpellName and confirmed in the logs: Ignite 412545, Deep Wound 412613, Blizzard 1279976 to 1279980 and 1279949, Flamestrike 1279983 to 1279990, Consecration 1280345 to 1280349, Rain of Fire 1282380 to 1282385, Hurricane 1278759, 1278965, 1278968, Volley 1279715, 1279719, 1279721.
2. **Engine.** Every file that registers a dot or heal over time (57 files across the eight classes, `common/item_effects`, `common/item_sets` and the hunter pets), the rule in `core/dot.go` and `Dot.OutcomeTick`, and the spell id each dot carries against the client row.
3. **Logs.** Our reports, read only: all 7 reports, 1,178 fights, every fight's events from R2 (`crit-flags/r2get.py`, `ticks.py`, `join.py`; `joined_ours.tsv` is per class and spell, every damage and heal event, critical over all, players). foreverlogs.gg: all 26 public reports, the server tick rows above queried by spell id (`fl_ticks.py`, `fl_ticks.tsv`), plus the 2026-11-12 lanes' cached report events (`ticks_fl.tsv`). Warcraft Logs still has no Forever partition (rogues-60.md). Every log is below level 50 (the advanced log's level field is the zone's scaled level), so the logs answer whether a row can crit, not how often at 60. Private reports are counted only in aggregate.
4. **Sims** on the Windows box only, one at a time, `nice -n 15`, GOMAXPROCS=5: the fork tests red at the pin and green at the fix in a throwaway worktree of the box's clone (`fork_tests.sh`), `spec-breakdown` at 10,000 iterations, seed 1, on every published entry Ignite reaches at both engines (`bands.sh`), and the site gates (`gates.sh`).

## The table

Client flags are the hotfixed 70888 rows (70938 identical). "Logs" is critical ticks over all ticks, ours plus foreverlogs.gg, summed over the trained ranks. "Rule" is the owner's 31 October rule as the engine applies it.

| effect | client flag | engine before | logs | verdict |
|---|---|---|---|---|
| **Warrior** Rend 772 to 11574 | PERIODIC_CAN_CRIT | crits, bleed rule, melee table | 54 / 771 (7.0%) | aligned |
| Deep Wounds, tick 412613 | CANT_CRIT | never (fixed 12 Nov, `NeverCrits`) | 0 / 206 | aligned |
| **Druid** Rake 1822 to 9904 | PERIODIC_CAN_CRIT | crits, bleed rule | 556 / 3,635 (15.3%) | aligned |
| Rip 1079 to 9896 | PERIODIC_CAN_CRIT | crits, bleed rule | 40 / 287 | aligned |
| Pounce Bleed 9007 to 9827 | PERIODIC_CAN_CRIT | crits, bleed rule | 3 / 9 | aligned |
| Lacerate 1235826 (bear) | PERIODIC_CAN_CRIT | crits, bleed rule | 6 / 36 | aligned |
| Moonfire 8921 to 9835 | PERIODIC_CAN_CRIT | crits, spell table | 107 / 955 | aligned |
| Insect Swarm 5570 to 24977 | PERIODIC_CAN_CRIT | crits, spell table | 22 / 326 | aligned |
| Hurricane 16914 to 17402, ticks 1278759, 1278965, 1278968 | channel unflagged (aura 226); tick rows direct damage, no CANT_CRIT | never (unflagged channel) | no log | **fixed B2**: crits (logs silent, client allows) |
| Rejuvenation 774 to 25299 | PERIODIC_CAN_CRIT | crits, healing crit | 65 / 874 (rank 1 0 / 192 at levels 4 to 10) | aligned |
| Regrowth's heal over time 8936 to 9858 | PERIODIC_CAN_CRIT | crits | 188 / 1,633 | aligned |
| Tranquility 740 to 9863 | PERIODIC_CAN_CRIT | crits | 2 / 36 | aligned |
| Wild Growth 1238214, 1238215 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Lifebloom | not a Forever spell (only the Season of Discovery 408124) | not registered | | n/a |
| **Hunter** Serpent Sting 1978 to 25295 | PERIODIC_CAN_CRIT | crits, spell table | 29 / 662 | aligned |
| Immolation Trap 13797 to 27023 | PERIODIC_CAN_CRIT | crits | 4 / 31 | aligned |
| Explosive Trap 13812 to 27025 | PERIODIC_CAN_CRIT | crits | 2 / 24 | aligned |
| Lacerate 1299332, Lacerating Strikes 1310536 | PERIODIC_CAN_CRIT | crits (bleed rule; Lacerating Strikes `CanCrit`) | no log | aligned |
| Volley 1510 to 14295, ticks 1279715, 1279719, 1279721 | channel unflagged (aura 226); tick rows direct damage, DefenseType 3 (ranged), no CANT_CRIT | never; arcane spell table | 52 / 478 (10.9%, one hunter's ranged rate) | **fixed B2**: crits on the ranged table |
| **Mage** Ignite, talent 11119, aura 412538, tick 412545 | tick CANT_CRIT (Attributes_2 0x20000004) | crits, true-dot rule, spell table (17% of ticks at raid 60) | 0 / 3 | **fixed B1**: never crits (logs silent) |
| Fireball's burn 133 to 25306 | PERIODIC_CAN_CRIT | crits | 21 / 270 | aligned |
| Pyroblast's burn 11366 to 18809 | PERIODIC_CAN_CRIT | crits | 0 / 2 | aligned |
| Frostfire Bolt's burn 1237312 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Blizzard 10 to 10187, ticks 1279976 to 1279980, 1279949 | channel unflagged (aura 226); tick rows direct damage, no CANT_CRIT | never (unflagged channel) | 902 / 2,767 (821 / 1,405 from one Shatter mage; 81 / 1,362 from four others) | **fixed B2**: crits |
| Flamestrike's burn, ticks 1279983 to 1279990 | aura 226; tick rows no CANT_CRIT | crits, true-dot rule | 55 / 939 | aligned |
| Arcane Missiles 5143 to 25345, missiles 7268 to 25346, 1303498 | channel unflagged; each missile a direct hit, no CANT_CRIT | each missile rolls hit and crit | 58 / 507 | aligned |
| **Priest** Shadow Word: Pain 589 to 10894 | PERIODIC_CAN_CRIT | crits | 40 / 704 | aligned |
| Devouring Plague 2944 to 19280 | PERIODIC_CAN_CRIT | crits (`CanCrit`, per-tick roll) | 23 / 344 | aligned |
| Mind Flay 15407 to 18807 | PERIODIC_CAN_CRIT (a flagged channel) | crits | 11 / 161 | aligned |
| Holy Fire's burn 14914 to 15261 | PERIODIC_CAN_CRIT | crits | 4 / 36 | aligned |
| Starshards 10797 to 19305 | PERIODIC_CAN_CRIT (channel) | crits | 3 / 41 | aligned |
| Renew 139 to 10929 | PERIODIC_CAN_CRIT | crits | 30 / 488 | aligned |
| Penance pulses 1240723 to 1240730, 1316991, 1316993 | direct heal and damage rows, no CANT_CRIT | each pulse rolls crit | 3 / 49 heals, 2 / 44 damage | aligned |
| Prayer of Mending | not periodic (direct heals on a charge) | direct heal crit | | n/a |
| Vampiric Touch | not a Forever spell (Season of Discovery 402668) | not registered | | n/a |
| **Shaman** Flame Shock 8050 to 29228 | PERIODIC_CAN_CRIT | crits | 35 / 291 | aligned |
| Searing Totem's Attack 3606 to 10436, Magma Totem 8187 to 10581 | direct damage, no CANT_CRIT | each attack or pulse rolls crit | no log | aligned (not periodic) |
| Healing Stream 5672 to 10461 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Riptide's heal over time 1239242 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Earth Shield 974 | a charge aura, no periodic effect | not a dot | | n/a |
| Lightning Shield 26363 to 26370 | CANT_CRIT (direct) | always hit, never crits | 0 / 441 | aligned |
| **Warlock** Corruption 172 to 25311 | PERIODIC_CAN_CRIT | crits | 11 / 187 | aligned |
| Curse of Agony / Bane of Agony 980 to 11713 | PERIODIC_CAN_CRIT | crits | 12 / 307 | aligned |
| Immolate's burn 348 to 25309 | PERIODIC_CAN_CRIT | crits | 1 / 53 | aligned |
| Siphon Life 18265 to 18881 | PERIODIC_CAN_CRIT | crits | 7 / 85 | aligned |
| Drain Life 689 to 11700 | PERIODIC_CAN_CRIT (channel) | crits | 3 / 54 | aligned |
| Drain Soul 1120 to 11675 | PERIODIC_CAN_CRIT (channel) | crits | 0 / 7 | aligned |
| Wrack 1316697 | PERIODIC_CAN_CRIT (channel) | crits | 1 / 13 | aligned |
| Bane of Doom / Curse of Doom 603 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Unstable Affliction 1242971, 1242972 | PERIODIC_CAN_CRIT | crits | no log | aligned |
| Rain of Fire 5740 to 11678, ticks 1282380 to 1282385 | channel unflagged (aura 226); tick rows direct damage, no CANT_CRIT | never (unflagged channel) | 200 / 3,315 (6.0%) | **fixed B2**: crits |
| Hellfire 1949 to 11684 | PERIODIC_CAN_CRIT | not registered in the engine | 0 / 7 | n/a |
| Health Funnel 755 to 11695 | CANT_CRIT | not a damage dot here | 0 / 19 | n/a |
| **Paladin** Consecration 26573 to 20924, ticks 1280345 to 1280349 | aura 226; tick rows direct damage, no CANT_CRIT | crits, true-dot rule, spell table | 3,945 / 35,665 (11.1%) | aligned |
| Holy Strike, the Seals and Judgements | direct, no periodic effect | direct crit | | n/a |
| **Items on published BiS** Blaze 16898, Bloodrazor Rend 17504, Eskhandar's Rake 22639, Skullforge Brand 17484, Hand of Ragnaros 21162, Fireball 18796, 18082, 21159, Corrosive Poison 13526, Poison 18203, Blight 9796, Siphon Health 18652, Poison Cloud 11790 | PERIODIC_CAN_CRIT | crit | Corrosive Poison 1 / 25, Poison Cloud 0 / 6 | aligned |
| Firebreather's burn 16413, Spire of Hakkar's Blood Siphon 1299943 | unflagged true dots, no CANT_CRIT | crit, true-dot rule | no log | owner's rule stands (the client does not forbid) |
| Thunderbrew's Boot Flask 12257 (pulses 12278) | channel trigger; direct pulses, no CANT_CRIT | each pulse rolls crit | no log | aligned |

Three unflagged heals over time the logs show never crit (Thrill of Adventure 1261483, 0 / 2,028; First Aid, 0 / 283; Cannibalize, 0 / 102) are not player spell rows any card uses, and the engine's heal over time rule (the flag alone) already keeps them flat.

## B1: Ignite's ticks crit (CONFIRMED-BUG, old, fixed)

**Client.** Forever's Ignite is one talent rank, 11119 ("burn for an additional 40% of your spell's damage over 412538's duration"). 412538 is a 4 s aura whose only effect is a dummy periodic (aura 226) every 2 s; the server casts 412545 on each pulse. 412545 is a direct Fire row with Attributes_2 0x20000004: SPELL_ATTR2_CANT_CRIT, the flag that took Deep Wound's crits away. The classic 12654 the engine names is not in the Forever client at all. The engine's own comments already read 412538 for the burn's shape; it applied the owner's true-dot rule to the ticks.

**Logs.** Three Ignite ticks in every readable Forever log (two in our reports, one on foreverlogs.gg), none critical: silent. The brief's rule applies (client forbids, logs silent: follow the client).

**Fix.** `mage/ignite.go`: the burn's `DotConfig` sets `NeverCrits` (`core.DotConfig.NeverCrits`, added for Deep Wounds). Red first: `TestIgniteTicksNeverCrit` (replacing `TestIgniteTicksCritWithoutFeedingIgnite`) fails at 79358b3c2 ("Ignite ticked 0 plain and 2 critical at a forced crit, want 2 plain and none") and passes at 8f61b9e6a.

**Before and after** (`spec-breakdown`, 10,000 iterations, seed 1; error under 0.06%):

| entry | change |
|---|---|
| Fire raid 60 Horde | -0.87% |
| Fire raid 60 Alliance | -1.06% |
| Fire bare 60 Horde / Alliance | -0.33% / -0.25% |
| Fire bare 50 | -0.33% |
| Fire bare 40 | -0.25% |
| Fire bare 30 | -0.18% |
| Fire bare 20 | -0.05% |
| Arcane raid 60 Horde / Alliance (Ignite 5/5 in the raid build) | -0.39% / -0.54% |

Frost and the bare Arcane builds take no Ignite. Ignite was 10.6% of Fire raid 60 damage at the pin with 17% of its ticks critical; crit is now worth a little less to Fire through Ignite's ticks, and the weights follow at the refresh (MR101). Nothing else changes: no rotation, build or order decision rests on Ignite's tick crits (the Combustion and Heating Up lines already count only direct crits).

## B2: the ground channels never crit (CONFIRMED-BUG, old, fixed, nothing published moves)

**Client.** Blizzard, Hurricane, Rain of Fire and Volley are channels whose only periodic effect is a dummy (aura 226); each pulse is a separate server row with plain school damage (Effect 2) and no Can't Crit. The channel's missing PERIODIC_CAN_CRIT has no periodic damage effect to apply to, so the rule's "channels follow the flag" never reaches these pulses: they crit as direct spells. Volley's tick rows are DefenseType 3, so they roll on the ranged table.

**Logs.** Blizzard 902 critical of 2,767 ticks (one mage at 58% on frozen targets; four others 6%), Rain of Fire 200 of 3,315 (14 warlocks), Volley 52 of 478 (one hunter, 10.9%, his ranged rate). Hurricane: no log. The logs win, and agree with the tick rows.

**Fix.** The four area dots set `CanCrit` (the field's doc now names this case), and Volley carries `DefenseType: DefenseTypeRanged`, so its pulses crit on the ranged table with the ranged multiplier and Mortal Shots. Red first: `conformance/TestGroundChannelTicksCrit` fails at the pin for all four (and Volley's defense type) and passes at the fix.

**Reach.** No published rotation casts Blizzard, Hurricane, Rain of Fire or Volley (every `sim/request/apl/*.apl.json` checked), and no card text says they tick flat. The engine's pack fights (`spec-breakdown -targets`) are the only place these numbers show; listed for the refresh as MR102 so a future pack page reads them right.

## Held and open

| # | item | why |
|---|---|---|
| H1 | Firebreather's burn (16413) and Spire of Hakkar's Blood Siphon (1299943) | unflagged non-channelled dots; the client does not forbid them (no Can't Crit), so the owner's rule stands; no log |
| H2 | Rejuvenation rank 1 0 / 192 crits at levels 4 to 10 | every higher rank crits at 5 to 11%; at those levels spell crit is 2 to 3%, so 0 / 192 is a one in a hundred draw; the client flags it; nothing at 60 rests on it. Watch it in later logs |
| O1 | every rate at 60 | every log is below 50; these checks settle whether a row crits, not the rate at 60 |
| O2 | `TestConformance/warlock` red at the pin | the warlock conformance golden is stale on master before this lane (red at 79358b3c2 and at the fix, same message); owned by the warlock lane, not touched here |

## Gates

- **Fork, box, throwaway worktree** (`crit-flags/fork_tests.sh`): both tests red at 79358b3c2, green at 8f61b9e6a; `go test -tags=with_db -p 1` over `sim/core`, `sim/mage/...`, `sim/druid/...`, `sim/warlock/...`, `sim/hunter/...`, `sim/common/...`: ok; `sim/conformance` ok but O2. No suite golden moved (the class suite fixtures cast none of the four channels and carry no Ignite build).
- **Site, pinned to 8f61b9e6a** (`crit-flags/gates.sh`, `gates2.sh`): `npm run lint` clean; after `npm run sync`, `astro check` 0 errors (2,310 files) and `npx vitest run src/lib/sim src/content/guides src/lib/guides` 131 files, 2,481 tests passed; `apl_check` against the pinned fork clean; `go vet ./...`; `go test ./request/` with the ladder goldens rewritten (only `mage-fire.golden.md` moved), then `./request/`, `./leveling/`, `./rankchoices/`, `./score/`, `./adapter/`, `./cmd/spec-breakdown/`, `./cmd/leveling-bis/`, `./cmd/addon-builds/`, `./cmd/talent-search/`: ok.
- **wow-player, mage:** SHIP on the one Fire guide line (the burn never crits); fire.md, arcane.md, the Fire card notes and the Fire coach read correctly with Ignite's ticks flat.

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| MR101 Fire and Arcane on Ignite without tick crits | mage-fire every bare band 20 to 60 and raid 60; mage-arcane raid 60: picks, weights (crit a little less), the race sweep, the published set_dps | B1 |
| MR102 the ground channels crit | nothing published; any pack-fight figure for Frost, Balance, Destruction or the hunters | B2 |

## Shas

- **Fork `lane/crit-flags` on master 79358b3c2:** ecab41abe (the tests, red at the pin), 8f61b9e6a (the fix). Master fast-forwarded to 8f61b9e6a.
- **Site pin:** 8f61b9e6a (`make engine-pin`).
- **Site:** the pin, the Fire guide's Ignite line, the ladder goldens the gates rewrote, this report and `crit-flags/` (scripts and tallies), the ledger row, MR101 and MR102.

## Sources

- wago.tools 1.60.1.70888 and 1.60.1.70938 with `useHotfixes=1` (SpellEffect, SpellMisc, SpellCategories, SpellClassOptions, Talent, SpellName, Spell).
- Our reports: `reports`, `fights` (read only) and every fight's `events.parquet` in R2 for the 7 reports.
- foreverlogs.gg `/api/reports/public` and `/api/reports/<id>/meta_events?spellIds=...` for the 26 public reports (13 Nov).
