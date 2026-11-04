# Sim audit round 14: regression lane, casters and healers (regcast)

Scope: every change rounds 12 and 13 and the land2 landing made to mages, warlocks, Balance, Elemental, Shadow and the five healers, and to the shared paths they use (heal profile, rank rows by race, trainables, regen phase). Each was re-read against what its round claimed, then shown with a test or a run.

Branches: site `wt/r14-regcast` from int/r14 (72471fcb3), engine `lane/r14-regcast` from int/r14 (5195a8889, no commits). Client 1.60.1.70753 with the 2026-11-03 hotfix push layered. Runs: spec-breakdown, 10,000 iterations, seeds 1 and 50001, one at a time under `nice -n 15`, `GOMAXPROCS=9`. Rank rows were tested by building spec-breakdown with one row taken out of the embedded rank-choices.json (`$R/regcast/mkbin.sh`, restored after each build).

**Result: 2 confirmed bugs.**
- One R13: Holy's Heal rank 1 row at 50 to 59, overtaken by round 13's Prayer of Mending fix. Fixed here.
- One old: the Shadow coach starter's priority contradicts the guide and the rotation. Patch routed to platform.

Every other round 12, round 13 and landing change in this group does what its round said.

## The landing (land2)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **The pin.** data/proto/ENGINE_SHA and sim/enginever/version.go both read 5195a8889. | files | Aligned | - | - |
| L2 | **apl-sync (fc90c1fcc).** `make apl-check ENGINE_DIR=<lane engine>`: all 56 copies match their curated rotations (28 specs, a fork copy and a site copy each). That includes the three healer copies the landing did not touch (priest-holy, priest-shadow, druid-restoration), which round 13 synced by hand. | apl-check output | Aligned | - | - |
| L3 | **addon-builds and the ladder (f9b264a4f).** `TestCommittedAddonBuildsFollowTheGuides` passes. The addon's point order matches the ladder's band builds wherever I decoded it:<br>- Elemental: Tidal Focus 5 by 50, and `-0533520000000000` at 60 in the ladder golden.<br>- Balance leveling: 30 is `5232220113…`, with Improved Entangling Roots at 0 and Naturalist and Nature's Grace in.<br>- Affliction raid: Improved Bane of Agony is taken. | `go test ./cmd/addon-builds`; testdata/ladder/*.golden.md; addon-builds.json decoded by tree and tier | Aligned | - | - |
| L4 | **addon-data rank rows (32fb15c5f).** Each band's ids in addon-data.json equal what the rows pick. Checked: Elemental 30 to 50 (Lightning Bolt rank 4 and 7, Flame Shock rank 1 and 2, Lava Burst rank 2); Balance 30; Fire 30 and 40; Holy priest 40 and 50; Holy paladin's Flash of Light pair at 40; the Restoration druid's tank row. Only Shadow carries `raceSpellId`, and its ids match the race rows at 40 and 50 (e.g. Mind Blast 8103 Dwarf and Night Elf, 8104 for the others). The addon maps the client race token to the same slugs (`Compat.RACE_SLUGS`, NightElf→night-elf, Scourge→undead). | addon-data.json; Data.lua:1816-1828; addon/ForeverSixty/Compat.lua:575 | Aligned | - | - |
| L5 | **Coach starters (32fb15c5f).** Balance's starter now has the pre-pull Wrath, and Shadow's has Vampiric Embrace. But the same Shadow starter tells the player the wrong priority. See X1. | data/curated/coach/*.md | see X1 | - | - |
| L6 | **dps-per-point for the mana specs (32fb15c5f).** The casters' values are smooth by band. Three healers read almost zero healing per point of healing power at low bands:<br>- Restoration druid 20/30/40: 0.029 / 0.035 / 0.015, against 0.14 to 0.16 before.<br>- Restoration shaman 20 and 40: 0.034 and 0.072.<br>- Holy paladin 30: 0.058.<br>The druid's band-40 fight shows why: 201.0 effective HPS with the row and without it, at mana lasting 443 s or 315 s. More healing does not raise its effective healing there, because it already covers about 84% of the band's incoming damage (tank 482 every 3.3 s, pulse 116 every 5 s). This is the coverage question round 13 raised for the raid (X1 there), seen at the bands. | dps-per-point.json; spec-breakdown -heal -preset bare | OPEN (healers: the band coverage rule; Measure at refresh) | - | - |
| L7 | **Vendored trainables and races (E 54ea3bfb1).** All nine trainables files and races.json are byte-identical to data/builds/1.60.1.70753. The conformance goldens move only by learn-row counts and costs from the trainables lane's fixes. Haunt, Unstable Affliction and Dispel Magic stay `class_spell` with no learn row, as round 13 trainables X5 settled. | cmp; golden diffs | Aligned | - | - |
| L8 | **Searing Totem in conformance (E 5195a8889).** Engine Dot lengths are 12/14/16/18/20/22 attacks × 2.437 s inside the client's 30/35/40/45/50/55 s, so they always end less than one attack early. The new reading accepts only that. The test's three cases (29,244 accepted; 30,000 − 2,437 and 30,500 compared) pass. | `go test ./sim/conformance` | Aligned | - | - |

## Casters: mage, Balance, Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Frostfire Bolt 24 yd/s (E 99b329919).** The client's SpellMisc.Speed is 24 on all three ranks, and `TestFrostfireBoltTravelsAtTheClientSpeed` passes. No published rotation casts it. | `go test -tags=with_db ./sim/mage` | Aligned | - | - |
| M2 | **Round 12's mage changes** (Evocation's regen through the last tick, mana gems sharing category 1153 with Demonic Rune): tests pass. The regen-phase change (E e88a33c97) leaves Evocation at exactly four ticks. An 8 s channel always spans four 2 s ticks, whatever the phase. | fork mage tests; `core/mana.go:263` | Aligned | - | - |
| B1 | **Nature's Grace on landing (E a86de2495).** Only Wrath has a missile speed among Balance's crits, and only Wrath takes the landing path (`wrath.go`). Starfire, Moonfire and the rest still proc in `OnSpellHitDealt` (`talents.go:157`, guarded on `MissileSpeed == 0`). Nothing procs twice. | code; druid tests pass | Aligned | - | - |
| B2 | **The pre-pull Wrath (E ed156274a, S e25400173).** The coach starter and addon-data both carry it. At 1.5 s (Improved Wrath 5/5, the walk's first five points) it ends on the pull. | coach/druid-balance.md; addon-data | Aligned | - | - |
| B3 | **Leveling build (S b7bc3b144) and the band-30 rows.** On the new band-30 ladder build (`5232220113000000-…`), each row still wins beyond error, horde, seeds 1 / 50001:<br>- Moonfire rank 3: 74.0 against 70.8 without it (+4.5%).<br>- Wrath rank 4: 74.0 against 69.2 (+6.9%).<br>- Insect Swarm rank 1: 74.0 against 72.3 (+2.4%). | mkbin variants ba-mf30/ba-wr30/ba-is30 | Aligned | - | - |
| E1 | **Lava Burst wait (E 25fca77c0, S 457a602df).** The `wait` line is skipped at zero length (`APLActionWait.IsReady` needs a duration above 0). So with Lava Burst ready but unaffordable, the shaman falls through to the cheap bolt and does not stall. Below the talent, `spellTimeToReady` on an unregistered spell drops the line. | `core/apl_actions_timing.go:34` | Aligned | - | - |
| E2 | **Searing Totem 2.437 s (E 5ca93cfa3).** Magma Totem keeps its own 2 s pulse (`fire_totems.go:174`). New check in round 13's 60 beta timelines (`r13/casters/tls`): a totem's first Attack lands a median 2.21 s after the drop (142 drops; 88 within 2.2 to 2.3 s), against the engine's 2.437 s. Either way a totem fits the same number of attacks in its lifetime (22 in 55 s), so no count changes. The engine re-drops the totem when its Dot ends, 1.4 s before the totem would expire, right after the last attack it could make. That is what a careful player does. | `py/searfirst.py` over r13 timelines | Aligned (observation) | - | - |
| E3 | **Tidal Focus at 50/60 (S 11f258619) and the Elemental rows.** On the new ladder builds, every row still wins beyond error (±0.1), horde, seeds 1 / 50001:<br>- Bare 50, 164.8 / 165.0 with all rows. Without each: Lightning Bolt rank 4 120.4 (+37%); 15208 at rank 7 164.1 (+0.4%); Flame Shock rank 2 158.2 (+4.2%); Lava Burst rank 2 158.0 (+4.3%).<br>- Bare 60, 203.4 / 203.5. Without each: Lightning Bolt rank 4 156.8 (+30%); rank 7 202.5 (+0.4%); Lava Burst rank 2 198.8 (+2.3%). | mkbin el-* variants, `-talents` the ladder strings | Aligned | - | - |
| E4 | **Published bare 50/60 still run the old build.** spec-breakdown's default reads bis/shaman-elemental.json, whose bare 50/60 talents predate Tidal Focus: 155.2 against 164.8 on the new build at bare 50. That is the BiS refresh's job (round 13 MR24), not a bug. | bis/shaman-elemental.json | Aligned (MR24) | - | - |

## Warlock

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Improved Bane of Agony (S 66fa0d042)** is in the raid order addon-builds writes (the test passes). Its +5% is unchanged in `talents.go`, and the warlock fork tests pass. | tests | Aligned | - | - |
| W2 | **Round 12's warlock changes** (Shadow Bolt and Incinerate at 20 yd/s, Command's +5% demon bolts): tests pass. | `go test -tags=with_db ./sim/warlock/...` | Aligned | - | - |
| W3 | **The rows after the regen phase.**<br>- Destruction Shadow Bolt rank 4 at 30: 91.0 / 90.9 against 90.4 (+0.6%). Holds.<br>- Destruction Corruption rank 3 at 40 (round 12 adopted it at +0.16%, error 0.13) is now neutral: horde 147.2 against 147.1, alliance 147.7 against 147.8. Taking it out moves nothing, so it does no harm. | mkbin wd-* | Aligned (Corruption row neutral) | - | - |
| W4 | **Soul Link in conformance (E 5195a8889).** The disposition says "no rotation casts it", but Demonology's opener casts it at −2 s (round 13 regcast F9). It is text in a report only, and no number reads it. | sim/conformance/dispositions.go:92 | Aligned (wording, warlock lane) | - | - |

## Healers and Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Holy's Heal rank 1 row at 50 to 59 lost 12% once Prayer of Mending ran freely.** <ul><li>The row (round 11, 7f4387eeb) replaced the list's Heal rank 4 (6064) with rank 1 (2054) at 50 to 59.</li><li>Round 13 removed Prayer of Mending's mana gate (S 0dcd484d5) and did not re-measure the Holy rows.</li><li>Bare 50 effective HPS, row against no row, seeds 1 / 50001: horde 210.4 / 210.3 against 239.4 / 239.4 (+13.8%); alliance 216.0 / 215.9 against 243.1 / 243.1 (+12.6%). On the ladder build: 213.3 against 238.4 (+11.8%).</li><li>With the gate put back the row ties: 193.7 against 194.2 on the ladder build. So the round-13 change is what turned it.</li><li>The fight still runs dry without the row (about 100 s of 300).</li></ul>For a player: at 50 to 59 a Holy priest's filler Heal is the rank-4 downrank the note already names, not rank 1. Prayer of Mending now saves enough mana to afford it. | mkbin holy-heal50; `-rotation` with the gate restored; `TestHolyHealAtFiftyCastsTheListsRankFour` failed (rank 1 at 50, 55 and 59 for all six races), then passed | CONFIRMED-BUG | R13 0dcd484d5 (the row it overtook: 7f4387eeb) | S 9dba07e78 |
| H2 | **The other Holy rows after Prayer of Mending,** horde, seeds 1 / 50001, row against no row:<br>- Bare 60: Flash Heal rank 7 +16%; Prayer of Healing rank 1 +0.6 / +0.5%.<br>- 50: Flash Heal rank 6 +2.8%; Prayer of Healing rank 1 +1.4%; Renew rank 1 +4.7%.<br>- 40: Flash Heal rank 4 +1.6%; Prayer of Healing rank 1 +0.6%; Renew rank 1 +3.9%.<br>Inside the new band-50 set (Heal row gone): Flash Heal rank 6 +24%, Renew rank 1 +2.1%. Prayer of Healing rank 1 is only +0.17 / +0.29% (239.4 against 239.0 / 238.7), at the edge of the healer bar. | mkbin holy-* | Aligned; band-50 Prayer of Healing at the edge (healers) | - | - |
| H3 | **Restoration druid tank Rejuvenation rows after D2** (the party on rank 1, unpaced), horde, seeds 1 / 50001:<br>- 30: rank 4 110.6 against 107.2 without (+3.2%).<br>- 40: 201.0 either way (tie).<br>- 50: rank 1 365.2 against 342.5 (+6.6%).<br>- Bare 60: rank 4 582.0 against 573.2 (+1.5%).<br>All hold or tie. But at 30 the row now ends the fight with mana unspent (lasts 318 s; without it, 234 s). The row file's own rule says a healer row never does that. A rank between 4 and 10 may gain more there. | mkbin rd-* | Aligned; band 30 mana note (healers) | - | - |
| H4 | **Shadow's Vampiric Embrace (S 175d4ff96, E ec5a105ee):** 6.0 casts and 5.4 landings in 180 s at raid 60. That is once a cooldown, with a resisted cast waiting out its cooldown as in game. The race rows still win with it in the rotation. Measured on Undead, the published horde race (round 13's figures in brackets, measured before Vampiric Embrace), seeds 1 / 50001:<br>- Mind Blast 50 to 59 rank 3: 230.6 against 229.8 without it, +0.35% (+0.4%).<br>- Mind Blast 40 to 49 rank 4: 149.7 against 148.0, +1.1% (+1.4%).<br>- Devouring Plague 50 to 59 rank 1: 230.6 against 222.6, +3.6% (+2.9 to +4.2%). | spec-breakdown; mkbin sh-* | Aligned | - | - |
| H5 | **Round 12's healer changes** (Vampiric Embrace 20% off the priest's own Shadow damage, Blood Fury on healing, Mental Agility, Idol of Swiftness, Reflection, Inner Fire before the pull): the fork priest, druid, paladin and shaman tests pass. No published healer is an orc. | `go test -tags=with_db ./sim/{priest,druid,paladin,shaman}/...` | Aligned | - | - |
| H6 | **Heal profile by band:** the site tests pass (`HealBand`, `HealProfile`, `CuratedHeal`). Every band reads its own row. 60 keeps 0.659, and Warcraft Logs is blocked (gcloud). | `go test ./request -run …` | Aligned | - | - |
| H7 | **Restoration druid's party Rejuvenation rank 1 (S 383d137f1):** 55 casts a fight at raid 60, 30 healed a mana, 4.8 ticks a cast (Nature's Splendor makes it 15 s). The coefficient on rank 1 rests on round 11 C1. The rank-1 tooltip check is pending the owner (ruling 5) and is not reopened here. | spec-breakdown -heal | Aligned (pending owner) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The Shadow coach starter gives the wrong priority.** The coach persona's "priority, in words" says:<br>- "Inner Focus into Devouring Plague outside the execute window";<br>- "Mind Blast only when Mind Flay cannot be cast";<br>- "Shadow Word: Death first below twenty percent".<br>The guide (shadow.md, Rotation and priority) and the rotation say otherwise: Inner Focus on Mind Blast every time it is up, Mind Blast on cooldown, Shadow Word: Death on cooldown, Mind Flay with everything else, and "nothing changes below 20%". The lines are hand overrides from October's rotation search (tools/coach/starter_overrides.py:140-151), and the launch-day guide replaced that order. The landing regenerated the starter, which kept them. The generator's own test (`test_guide_order_statements_win`) pins the old order.<br>For a player: the coach would tell a Shadow priest to hold Mind Blast behind Mind Flay and save Shadow Word: Death for the execute, the opposite of the guide. | starter_overrides.py; shadow.md; data/curated/apl/priest-shadow.json; data/curated/coach/priest-shadow.md:24-27 | CONFIRMED-BUG (coach copy) | old (dcd34cf96, 2026-10-27; stale since the guide's adcaad714, 2026-11-02) | - (patch routed to platform) |
| X2 | `GUIDE_LINES` in starter_overrides.py has `"priest-discipline"` twice. Python keeps the second, so the first one's Binding Heal line is dead code. The starter still gets the guide's own Binding Heal line, so nothing visible changes. | starter_overrides.py | Aligned (lint, in the X1 patch note) | - | - |

## Fresh re-audit (new eyes)

- **Searing Totem's first attack** (E2): checked against logs for the first time. It changes no count.
- **Every caster and healer rank row on its current rotation and band build**, not the one it was adopted on (E3, B3, H1 to H4, W3, and Fire Blast and Arcane Frostbolt below). This found H1.
- **Fire Blast rank 3 and Arcane Frostbolt rank 4 at 30** (horde, seeds 1 / 50001): Fire 72.0 against 71.5 / 71.6 without; Arcane 86.4 / 86.6 against 85.7. Both hold.
- **Coach starter lines against the guides** for the group's specs:
  - Shadow is X1.
  - Elemental's order (Lava Burst first while Flame Shock is up) matches the rotation.
  - Fire, Affliction, Holy paladin, Discipline and Restoration druid match their guides.

## Before/after

Only H1 moves a number: Holy priest, bare band 50 (levels 50 to 59). Effective HPS, 10,000 iterations, before = `$R/sb-before`, after = this branch.

| spec, band | faction | seed 1 | % (seed 1 / 50001) |
|---|---|---|---|
| priest-holy bare 50 | horde | 210.4 to 239.4 | +13.8 / +13.8 |
| priest-holy bare 50 | alliance | 216.0 to 243.1 | +12.5 / +12.6 |

Raid 60 does not move. The spot checks equal $R/before.tsv: priest-holy horde s1 708.7, shaman-elemental 480.3. No spec outside the group moves.

## Files to regenerate

- sim/request `TestRotationLadder` golden for **priest-holy** (band 50 casts Heal 6064, not 2054). It fails until regenerated.
- **addon-data.json and Data.lua**: Holy's band-50 Heal lines become 6064.
- **data/curated/coach/priest-shadow.md** (`python tools/coach/persona_starter.py --spec priest-shadow`) once the X1 patch lands. `test_generator_is_deterministic_and_committed_files_match` fails between the two.
- **dps-per-point** for priest-holy band 50, at the next measure.

## Measure at refresh

- Holy priest bare band 50 picks and weights without the Heal rank 1 row (H1).
- Elemental bare 50/60 and Balance bare 30 to 60 publish on their new builds (E4; round 13 MR24).
- Healer healing-power weights at low bands (L6), once healers settle the band coverage.

## For other lanes

- **Platform (tools/coach):** apply `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r14/regcast/patches/shadow-coach-starter.patch` (X1). It does three things:
  - rewrites the four Shadow override lines from shadow.md;
  - puts Mind Blast and Shadow Word: Death ahead of Mind Flay in `GUIDE_ORDER`;
  - flips the order assertion in `test_guide_order_statements_win` and adds `test_shadow_priority_follows_the_launch_guide`.

  With it, `pytest tools/coach/test_persona_starter.py` passes except the committed-file check, until priest-shadow.md is regenerated. Also drop the first, dead `"priest-discipline"` entry (X2).
- **Healers:**
  - I removed Holy's Heal 50 to 59 row (S 9dba07e78; you own healer rows). Please re-measure Holy's band-50 set on your round-14 profile. Prayer of Healing rank 1 at 50 is at the edge inside the new set (H2).
  - Restoration druid band 30: the tank row now ends the fight with mana unspent (H3).
  - Low-band healing-power weights are near zero for the druid, shaman and Holy paladin (L6). This is the same coverage rule as ruling 3, below 60.
- **Warlock:** the conformance disposition for Soul Link says no rotation casts it; Demonology's opener does (W4).
- **Integrator:** the ladder golden, addon-data and coach regenerations above.

## Still open

- L6 (band coverage and healer weights), H2 (Holy band-50 Prayer of Healing at the edge), H3 (druid band-30 mana).
- Carried, with no new evidence: Imp Firebolt speed (pet distance), Demonology pre-pull mana, Elemental Mastery's charge, Eclipse timing, Prayer of Mending hop travel, Inner Focus on Power Infusion. Warcraft Logs work is blocked (gcloud login expired).
