# Sim-accuracy round 6: pipeline lane

Owner priority 3: the spell generators read the server's hotfix push by themselves. Round 5 read the 1.60.1.70707 push by hand (`scratchpad/r5/hotfix70707/`), because every spell generator on the site and in the fork read only the shipped DB2 CSVs.

Branches: site `wt/r6-pipeline` (S), fork `lane/r6-pipeline` from 8baffe5fc (E). Nothing pushed, pinned or dispatched.

## What changed

- **One layering rule, applied once, in place** (`data/pipeline/hotfix_push.py`, S f9155ec20). `pipeline fetch` downloads the shipped tables. It then reads wago.tools' push log for the build and layers it into `raw/<Table>.csv` for every fetched table except ItemSparse/Item. `hotfix_merge` keeps those two, by its own rule.
- **The rule** (`layer_rows`):
  - Per record id, the newest push wins: highest push id, then wago's entry id.
  - Status 1 replaces the shipped row, or adds a row the client lacks.
  - Status 2 drops the row.
  - Statuses 3 (invalidated) and 4 (not public) carry no values. The shipped row stands, and the record is counted as withheld.
  - A cell equal by value keeps the shipped print.
  - A push row of the wrong width, or whose ID column disagrees with its record id, fails the run.
- **On disk.** The file as fetched moves to `raw/shipped/`. The push is cached per table in `raw/wago-push/`, with `LAYERED.json` holding the counts and the sha of each layered file. Re-layering is idempotent. A fresh fetch over a layered build becomes the new shipped copy.
- **Offline.** `pipeline layer-hotfix-push --build <b> --offline` re-layers from the cache, and fails when there is no cache.
- **Every reader is covered.** All of them read `raw/<Table>.csv`, so none needed code: normalize, simconst, trainables, levels/spellranks, simdb, `tools/gen_set_rows.py`, and the fork's relicmods, periodiccrit and eureka. data.yml, engine.yml and bis.yml all run `fetch`. Comments were added to data.yml, engine.yml, `data/README.md` and `docs/engine-follow.md`.
- **Fails loudly.** An HTTP error, a page with no data, or a search whose counts do not add up fails `fetch`. Nothing falls back silently to shipped rows.
- **The wago search pager drops rows, and the fetch works around it.**
  - The pager returns 25 rows a page in push order, with ties in no stable order. Paging past page one repeats some rows and never shows others, the same ones every time. CurvePoint at 70707 came back as 182 of 212.
  - The fetch splits each search by the last digits of the record id until every answer fits on one page. It then checks that the children add up to the parent's total. A record whose id equals the suffix is found by the last digit of its push id.
  - A full 70707 fetch takes 1,337 requests, about 3 minutes.
- **Removed spells.** simconst now drops a spell whose SpellName row the push removed (S 2be8245c1). Without this, spellconst wrote 8 nameless constants and spellranks grouped them under "".

Tests, written first and seen failing: `data/tests/test_hotfix_push.py` (25: changed, added and removed rows, newest push wins, withheld, print width, schema errors, determinism, the pager split against a fake wago that drops rows, HTTP and parse failures, idempotence, refetch, offline) and `data/tests/test_hotfix_push_generators.py` (7). The generator tests run on the real 70707 rows and pushes:
- simconst: Ice Barrier's pair 3/3 becomes 5/5, and Frigid Power's 3/3 becomes 15/15.
- spell text: Touch of the Grave's `$proccooldown` goes from 25 sec to 15 sec.
- trainables: follow a skill-line row the push removes or adds.
- simconst: drops a spell whose name the push removed.
- CLI: `fetch` always layers, and `--offline` is an explicit mode.

Touched site tests pass: 134 across hotfix_push, wago, wago_hotfixes, simconst, trainables, spelltext and levels. ruff is clean.

## Cross-cutting findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | No generator read the server's hotfix push. Spell, talent and skill values the server changed reached the engine only by hand. | `data/pipeline/simconst.py:101-115`, `trainables.py:140-192`, the fork's `tools/{periodiccrit,eureka,relicmods}/gen.py` read only `raw/*.csv`. Round 5 found the pushed values by hand (casters R1b, hybrids X1). | CONFIRMED-BUG (pipeline) | S f9155ec20 |
| P2 | Round 5's hand dump of the push (`hfdump.py`) was incomplete. It paged wago's search, which repeats and drops rows. It missed 78 of 586 SpellEffect entries, 312 of 1,747 SpellMisc, 148 of 812 SpellName and 5 of 68 SpellAuraOptions. It saw no Trait*, CurvePoint, ChrRaces or ItemEffect rows at all. | Entry-id comparison of `raw/wago-push/*.json` with `r5/hotfix70707/*.json`. Rows that change a value are in P5-P11 and `scratchpad/r6/pipeline/missed-by-r5.tsv`. | CONFIRMED-BUG (the round-5 tool) | S f9155ec20 (the fetch splits searches) |
| P3 | Spells the push deletes became nameless constants. simconst kept their SpellClassOptions rows with name "". | Push SpellName/Spell status 2 for 1038, 1912, 19977-19979, 25890, 25895, 1311085, 1325499. | CONFIRMED-BUG (fresh, from P1) | S 2be8245c1 |
| P4 | The fork regenerated from the layered build (`tools/follow.sh`). The engine had hand-applied every pushed value it uses, so the layered client now agrees with it. Conformance mismatches go from 15 to 0 (druid), 14 to 0 (mage) and 10 to 5 (paladin). The hand code that indexed the old rank ladders now follows the generated tables: Lacerate's ladder is the table's; Ferocious Bite gains the pushed rank 1 (1328618: level 24, 9-19, 0.6 a point of energy, and 24 a combo point from Wowhead's live tooltip, read 2026-11-01); Wild Growth's healing table states the 3 s cast the spell already used; Holy Shield's client comparison is rank for rank. | Follow report: `verdict: review` only for 3 new trainables without a disposition and Templar's Bulwark (P9). `go test --tags=with_db ./sim/...` is green. | Aligned (no DPS move; see Before/after) | E 6d8bb01a8 |
| P5 | **Blessing of Light and Greater Blessing of Light no longer exist in Forever. The Holy paladin APL still precasts Greater Blessing of Light (25890) on a party member, and the engine adds its +400 Holy Light / +115 Flash of Light bonus healing.** | Push: SpellName and Spell rows of 19977, 19978, 19979 and 25890 have status 2 (removed). Wowhead Forever's tooltips for 19979 and 25890 return 404 (read 2026-11-01). The layered trainables drop both. `data/curated/apl/paladin-holy.json:616` casts 25890. The engine is `sim/paladin/blessing_of_light.go:31-41`. Oath of Light (1328548, pushed SkillLineAbility 59879, Holy line) looks like the replacement; it is unmodelled (round 5 L3). spec-breakdown shows 1 GBoL cast a fight (Holy paladin, band 60 raid). | CONFIRMED-BUG (two sources). Not fixed here; for the hybrids lane | - |
| P6 | Salvation (1038, 1912) and Mass Salvation (25895) are removed the same way. | Push SpellName/Spell status 2; Wowhead 404 for 1038. The engine registers them in `sim/paladin/blessing_ranks.go:99-104`. It is threat only: no DPS or HPS moves, and tanks carry none (`data/curated/presets.json` notes). | CONFIRMED (no sim number moves); for the hybrids lane to drop | - |
| P7 | Holy Shield (20925, 20928) and Shield Specialization (20150) proc masks add 0x20000 (taken harmful spells): 680 → 131752. | Push SpellAuraOptions 132432, 131453 and 131646. Round 5's dump missed all three. | OPEN, for the protection paladin lane (matters only if the tank encounter casts spells at the tank) | - |
| P8 | Echo of Fury (1311701) and Echo of Righteousness (1311704) aura options change: CumulativeAura 0 → 3, ProcChance 100 → 101, ProcCharges 1 → 0, ProcTypeMask 4 → 0. This is the same change round 5 saw on Echo of Command (247753). | Push SpellAuraOptions 247752 and 247754, both missed by round 5's dump. | OPEN, for the hybrids lane: check that Unsealed Wrath's stored copies stack to 3 for every seal | - |
| P9 | Templar's Bulwark (1311015) leaves the paladin trainables. Its pushed SpellLevels row is level 30, and the engine registers it talent-gated at level 0. | Follow report; the 5 paladin conformance mismatches left (`required_level 0->30`). | OPEN, for the protection paladin lane (talent-gated, so likely no number) | - |
| P10 | Binding Heal (401937, 1240770-1240774; levels 25-56) becomes a priest trainable, from pushed SkillLineAbility rows. Voice of Truth (1310897, level 20) and Turn Evil become paladin trainables. | Layered trainables; the follow report lists them without a disposition. | OPEN, for the casters lane: Binding Heal in the Holy/Discipline APL? | - |
| P11 | Fresh rows round 5 missed with no sim effect: Execute Attributes_0 drops 0x100 (logging); Light's Vigil range index 578 → 5; Feline Grace duration; Plainsrunning 1 → 2; Escape Rope; Endurance's second effect 1349372 (aura 50) removed; Boosted Critical Chance 1229519 effect removed. | Push SpellMisc/SpellEffect. The engine's Endurance is 5% Health only (`sim/core/racials.go:428`), as Wowhead's live text says. | Aligned | - |
| P12 | Talent trees: the layered trait tables build the same trees as the committed ones. They already come from Wowhead's live payload. Only `hotfix_only: true` drops on 3 druid, 5 paladin and 1 priest talents, because the client tables now carry them. | normalize on shipped vs layered raw: talents/, races.json, classes.json identical apart from that flag. | Aligned | - |
| P13 | Nothing re-fetches an active build when the server pushes again. data.yml's fetch runs only on a new build or a dispatch, so a 1 November push to 70707 waits for the next build. | `.github/workflows/data.yml` triggers; client.yml dispatches only on a new build. | OPEN (a schedule, or client.yml comparing the push count in `LAYERED.json`) | - |
| P14 | The engine-follow skip hash (`tools/engine/follow-inputs-hash.sh`) hashes the committed site inputs. A push that changes only a raw-read generator's output (relicmods, periodiccrit, eureka, set rows) does not trigger a follow by itself. | Script header. | OPEN (low; today's push moves periodic crit by 2 bleed ranks only) | - |

## What the layer changes in committed outputs (70707)

None of these outputs are committed here: data.yml regenerates the site outputs, and the engine follow pins the fork outputs.

Site outputs:
- **spellconst/** for druid, mage, paladin, priest, rogue and warlock: 112 value changes, listed in full below.
- **trainables/**:
  - druid: Ferocious Bite gains 1328618 (level 24). Lacerate gains 1328617 and 1328616 (levels 25 and 34). Primal Bite's levels go from 25/36/48 to 40/50/60. Wild Growth's cast goes from 0 to 3000 ms.
  - mage: Ice Lance's cooldown goes from 0 to 8000 ms.
  - paladin: Blessing of Light, Greater Blessing of Light, Salvation, Mass Salvation, Revoke and Templar's Bulwark are removed. Turn Evil (38/38/52) and Voice of Truth (20) are added. Holy Shield's levels go from 40/50/60 to 30/40/50, and its top cost from 240 to 195.
  - priest: Binding Heal is added (25, 32, 38, 44, 50, 56).
- **spellranks.json**:
  - druid: Bloom is added. Ferocious Bite goes from 5 to 6 ranks, and Lacerate from 3 to 5. Primal Bite's levels go from 25/36/48 to 40/50/60.
  - mage: Predatory Discipline is added.
  - paladin: Blessing of Light goes from 3 ranks to 1. Greater Blessing of Light, Mass Salvation, Revoke, Salvation and Templar's Bulwark are removed. Boundless Light, Shieldsworn, Stoicism, Stun and Turn Evil are added. Holy Shield's levels go from 40/50/60 to 30/40/50.
  - priest: Holy Binding is added.
- **talents/**: druid, paladin and priest drop the `hotfix_only` flag (P12).
- **spells.json**: 22 spells are added.
- **items/**: druid and paladin rune effect texts change. Mangle's text goes from "plus 26" to "plus 42". Soul Siphon's text changes to include Drain Mana.
- **simdb.bin**: one enchant is added (8737, Speed Increase). Items: 0 changed of 11,463.
- **levels.json**: unchanged.

Fork outputs (in E 6d8bb01a8):
- `sim/{druid,mage,paladin,priest,rogue}/constants_auto_gen.go`. The engine reads only FerociousBite*, Lacerate* and PrimalBite*; the other 212 changed identifiers have no reader.
- `sim/core/periodic_crit_auto_gen.go`: Lacerate 1328616 and 1328617 are added. Bleeds always crit by owner ruling, so nothing moves.
- `sim/core/client_sets_gen.go`: Summon Arcane Elemental 1324323 now triggers 1324322 directly. Set 2134 is unregistered.
- Conformance client tables, goldens and SUMMARY.md.
- The talentgen fixtures (the flag only).

`races.json` in the conformance tables also catches up with the site's current racial text; that change comes from the site, not from the push.

### spellconst value changes (spell, field, old, new)

The 8 paladin "SPELL present → removed" rows are P3/P5/P6.

| class | spell id | spell | field | old | new |
|---|---|---|---|---|---|
| druid | 22568 | Ferocious Bite | rank | 1 | 2 |
| druid | 22827 | Ferocious Bite | rank | 2 | 3 |
| druid | 22828 | Ferocious Bite | rank | 3 | 4 |
| druid | 22829 | Ferocious Bite | rank | 4 | 5 |
| druid | 24907 | Moonkin Aura | effect[0].aura | 57 | 290 |
| druid | 24932 | Leader of the Pack | effect[0].aura | 52 | 290 |
| druid | 31018 | Ferocious Bite | rank | 5 | 6 |
| druid | 407995 | Primal Bite | spell_level | 25 | 40 |
| druid | 407995 | Primal Bite | effect[0].amount | 26 | 42 |
| druid | 408120 | Wild Growth | cast_time_ms | 0 | 3000 |
| druid | 414644 | Lacerate | rank | 1 | 3 |
| druid | 1235826 | Lacerate | rank | 2 | 4 |
| druid | 1235827 | Lacerate | rank | 3 | 5 |
| druid | 1238069 | Primal Bite | spell_level | 36 | 50 |
| druid | 1238069 | Primal Bite | effect[0].amount | 38 | 62 |
| druid | 1238070 | Primal Bite | spell_level | 48 | 60 |
| druid | 1238070 | Primal Bite | effect[0].amount | 59 | 77 |
| druid | 1238214 | Wild Growth | cast_time_ms | 0 | 3000 |
| druid | 1238215 | Wild Growth | cast_time_ms | 0 | 3000 |
| druid | 1322605 | Shifting Power | cooldown_ms | 16000 | 8000 |
| druid | 1328616 | Lacerate | SPELL | absent | added (rank 2, lvl 34, effects [(6, 3, 6), (3, 0, 10)]) |
| druid | 1328617 | Lacerate | SPELL | absent | added (rank 1, lvl 25, effects [(6, 3, 4), (3, 0, 10)]) |
| druid | 1328618 | Ferocious Bite | SPELL | absent | added (rank 1, lvl 24, effects [(2, 0, 14), (3, 0, 60)]) |
| druid | 1328621 | Bloom | SPELL | absent | added (rank 0, lvl 1, effects [(6, 4, 100), (6, 4, 50)]) |
| druid | 1328622 | Bloom | SPELL | absent | added (rank 0, lvl 1, effects [(10, 0, 1)]) |
| mage | 11426 | Ice Barrier | effect[1].amount | 3 | 5 |
| mage | 11426 | Ice Barrier | effect[2].amount | 3 | 5 |
| mage | 12579 | Winter's Chill | effect[0].amount | 5 | 20 |
| mage | 13031 | Ice Barrier | effect[1].amount | 3 | 5 |
| mage | 13031 | Ice Barrier | effect[2].amount | 3 | 5 |
| mage | 13032 | Ice Barrier | effect[1].amount | 3 | 5 |
| mage | 13032 | Ice Barrier | effect[2].amount | 3 | 5 |
| mage | 13033 | Ice Barrier | effect[1].amount | 3 | 5 |
| mage | 13033 | Ice Barrier | effect[2].amount | 3 | 5 |
| mage | 400573 | Arcane Blast | duration_ms | 15000 | 18000 |
| mage | 400573 | Arcane Blast | effect[0].amount | 15 | 12 |
| mage | 400573 | Arcane Blast | effect[1].amount | 175 | 100 |
| mage | 400640 | Ice Lance | category_cooldown_ms | 0 | 8000 |
| mage | 400640 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1240044 | Ice Lance | category_cooldown_ms | 0 | 8000 |
| mage | 1240044 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1240045 | Ice Lance | category_cooldown_ms | 0 | 8000 |
| mage | 1240045 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1240046 | Ice Lance | category_cooldown_ms | 0 | 8000 |
| mage | 1240046 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1240047 | Ice Lance | category_cooldown_ms | 0 | 8000 |
| mage | 1240047 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1312002 | Ice Lance | effect[1].amount | 200 | 300 |
| mage | 1326891 | Frigid Power | effect[0].amount | 3 | 15 |
| mage | 1326891 | Frigid Power | effect[1].amount | 3 | 15 |
| mage | 1328613 | Predatory Discipline | SPELL | absent | added (rank 0, lvl 1, effects [(6, 4, 20)]) |
| paladin | 1038 | Salvation | SPELL | present | removed |
| paladin | 5923 | Holy Power | effect[0].aura | 108 | 50 |
| paladin | 5923 | Holy Power | effect[0].amount | 5 | 0 |
| paladin | 5923 | Holy Power | effect[0].misc_value | 7 | 2 |
| paladin | 5923 | Holy Power | effect[1] | {'index': 1, 'effect': 6, 'aura': 108, 'amount': 15, 'sp_coefficient': 1.0, 'misc_value': 7} | removed |
| paladin | 19977 | Blessing of Light | SPELL | present | removed |
| paladin | 19978 | Blessing of Light | SPELL | present | removed |
| paladin | 19979 | Blessing of Light | SPELL | present | removed |
| paladin | 20170 | Stun | SPELL | absent | added (rank 1, lvl 22, effects [(6, 12, 0)]) |
| paladin | 20232 | Judgement of Fury | duration_ms | 4000 | 3000 |
| paladin | 20925 | Holy Shield | cost | 150 | 100 |
| paladin | 20925 | Holy Shield | spell_level | 40 | 30 |
| paladin | 20925 | Holy Shield | effect[1].amount | 110 | 75 |
| paladin | 20927 | Holy Shield | cost | 195 | 150 |
| paladin | 20927 | Holy Shield | spell_level | 50 | 40 |
| paladin | 20927 | Holy Shield | effect[1].amount | 153 | 110 |
| paladin | 20928 | Holy Shield | cost | 240 | 195 |
| paladin | 20928 | Holy Shield | spell_level | 60 | 50 |
| paladin | 20928 | Holy Shield | effect[1].amount | 221 | 153 |
| paladin | 25771 | Forbearance | family_mask | [0, 0, 0, 0] | [0, 0, 128, 0] |
| paladin | 25780 | Righteous Fury | effect[2].misc_value | 12 | 23 |
| paladin | 25780 | Righteous Fury | effect[4] | absent | {'index': 4, 'effect': 6, 'aura': 10, 'amount': 33, 'sp_coefficient': 1.0, 'misc_value': 127} |
| paladin | 25890 | Greater Blessing of Light | SPELL | present | removed |
| paladin | 25895 | Mass Salvation | SPELL | present | removed |
| paladin | 456544 | Holy Shield | effect[0].amount | 1 | 0 |
| paladin | 456544 | Holy Shield | effect[0].sp_coefficient | 0.05000000075 | 0.079999998211861 |
| paladin | 1224697 | Sacred Duty | effect[0].aura | 137 | 107 |
| paladin | 1224697 | Sacred Duty | effect[0].amount | 4 | -10000 |
| paladin | 1224697 | Sacred Duty | effect[0].misc_value | 0 | 1 |
| paladin | 1310735 | Unsealed Wrath | name | Twist of Light | Unsealed Wrath |
| paladin | 1310735 | Unsealed Wrath | effect[0].amount | -20 | 0 |
| paladin | 1310904 | Divine Precision | effect[0].amount | 18 | 8 |
| paladin | 1310904 | Divine Precision | effect[1] | absent | {'index': 1, 'effect': 6, 'aura': 232, 'amount': -20, 'misc_value': 9} |
| paladin | 1310994 | Swift Judgement | school_mask | 1 | 2 |
| paladin | 1310994 | Swift Judgement | duration_ms | -1 | 0 |
| paladin | 1310994 | Swift Judgement | effect[0].effect | 6 | 77 |
| paladin | 1310994 | Swift Judgement | effect[0].aura | 108 | 0 |
| paladin | 1310994 | Swift Judgement | effect[0].amount | -100 | 0 |
| paladin | 1310994 | Swift Judgement | effect[0].misc_value | 14 | 0 |
| paladin | 1311015 | Templar's Bulwark | spell_level | 30 | 0 |
| paladin | 1311084 | Champion of the Light | effect[1] | absent | {'index': 1, 'effect': 6, 'aura': 107, 'amount': -1000, 'sp_coefficient': 1.0, 'misc_value': 10} |
| paladin | 1311085 | Instrument of Law | SPELL | present | removed |
| paladin | 1325499 | Revoke | SPELL | present | removed |
| paladin | 1328370 | Turn Evil | SPELL | absent | added (rank 0, lvl 0, effects [(6, 107, -1500), (6, 107, -1500), (6, 332, 1328374), (6, 332, 1328373), (6, 332, 1328372)]) |
| paladin | 1328372 | Turn Evil | SPELL | absent | added (rank 3, lvl 52, effects [(6, 7, 0), (6, 31, 25)]) |
| paladin | 1328373 | Turn Evil | SPELL | absent | added (rank 2, lvl 38, effects [(6, 7, 0), (6, 31, 25)]) |
| paladin | 1328374 | Turn Evil | SPELL | absent | added (rank 1, lvl 38, effects [(6, 7, 0), (6, 31, 25)]) |
| paladin | 1328544 | Shieldsworn | SPELL | absent | added (rank 0, lvl 1, effects [(6, 107, -4000)]) |
| paladin | 1328546 | Stoicism | SPELL | absent | added (rank 0, lvl 1, effects [(6, 232, -20), (6, 232, -20)]) |
| paladin | 1328548 | Oath of Light | SPELL | absent | added (rank 0, lvl 0, effects [(6, 226, 0), (6, 4, 0), (6, 4, 0)]) |
| paladin | 1328550 | Boundless Light | SPELL | absent | added (rank 0, lvl 1, effects [(6, 4, 0)]) |
| paladin | 1328551 | Shieldsworn | SPELL | absent | added (rank 0, lvl 1, effects [(6, 528, 0)]) |
| priest | 15007 | Resurrection Sickness | duration_ms | 600000 | 60000 |
| priest | 1309595 | Shadow Word: Death | effect[2].amount | 10 | 30 |
| priest | 1309633 | Shadow Word: Death | effect[2].amount | 10 | 30 |
| priest | 1309635 | Shadow Word: Death | effect[2].amount | 10 | 30 |
| priest | 1309636 | Shadow Word: Death | effect[2].amount | 10 | 30 |
| priest | 1328569 | Holy Binding | SPELL | absent | added (rank 0, lvl 0, effects [(6, 4, 50)]) |
| priest | 1328570 | Holy Binding | SPELL | absent | added (rank 0, lvl 1, effects [(10, 0, 1)]) |
| rogue | 2764 | Throw | cast_time_ms | 500 | 0 |
| warlock | 403511 | Improved Drains | effect[1] | absent | {'index': 1, 'effect': 6, 'aura': 108, 'amount': 20, 'sp_coefficient': 1.0, 'misc_value': 3} |

## Changed values the engine uses that no round-5 lane hand-applied

- **Blessing of Light / Greater Blessing of Light removed** (P5). This is the only one that moves a published number (Holy paladin HPS). For the hybrids lane.
- **Salvation / Mass Salvation removed** (P6): threat only.
- **Holy Shield and Shield Specialization proc masks** (P7), **Echo of Fury/Righteousness stacks** (P8), **Templar's Bulwark level** (P9), **Binding Heal trainable** (P10): unverified effect. Each is for its class lane.
- Every other pushed value the engine reads was already hand-applied in rounds 4 and 5. The evidence is the conformance mismatches that fall to zero.

Once S 2be8245c1 is in, the next follow drops the removed spells from the client tables. Three paladin tests then fail until the engine stops registering them: `TestBlessingsMatchClient` (1038 not in the client table), `TestBlessingOfLightMatchesClient` (19977) and `TestARankRegistersOnlyWhenTheCharacterHasLearnedIt` (19977). Those failing tests are the start of P5/P6's fix. The follow diff is saved as `scratchpad/r6/pipeline/follow2-after-simconst-drop.patch`. Apply it together with the paladin lane's fix, or let engine.yml open it as a review follow.

## Before/after

Engine 8baffe5fc against E 6d8bb01a8, site wt/r6-pipeline, 10,000 iterations, horde.

| spec (band, preset) | seed 1 before → after | seed 50001 before → after | relative |
|---|---|---|---|
| druid-feral (60 raid) | 616.455 → 616.455 DPS | 616.495 → 616.495 | 0.00% |
| druid-feral (20/30/40/60 bare) | 45.878 / 97.351 / 130.404 / 269.243, unchanged | unchanged | 0.00% |
| druid-feral-bear (60 raid, tank; 20-60 bare) | identical output (sorted) | identical | 0.00% |
| druid-balance (20/30/40/60 bare) | 32.500 / 61.971 / 120.502 / 250.963, unchanged | unchanged | 0.00% |
| druid-restoration (60 raid, heal) | 739.2 → 739.2 HPS | 739.7 → 739.7 | 0.00% |
| paladin-holy (60 raid, heal) | 544.2 → 544.2 HPS | 543.8 → 543.8 | 0.00% |
| paladin-protection (60 raid, tank) | identical output (sorted) | identical | 0.00% |
| mage-frost (60 raid) | 546.909 → 546.909 DPS | 547.638 → 547.638 | 0.00% |

## Still open

- P5/P6, the Blessing of Light and Salvation removal, needs an engine and APL change. For the hybrids lane, with the follow2 patch.
- P7-P10 are class-lane checks.
- P13: no re-fetch on a new push for an active build.
- P14: the follow skip hash ignores raw-only generator inputs.
- ItemSparse/Item still take wago's `useHotfixes=1` CSV (`fetch-wago-hotfixes`), which does not honour removals. Moving them to the push log would need `hotfix_merge`'s add-only rule carried over.
- Tables no generator reads are not fetched, so they are not layered: ChrRaceRacialAbility, SpellCategories, SpellProcsPerMinute, SpellEquippedItems and SpellLabel. The engine's racials and those rules stay hand-written; Orc Command is E e074b0415 from round 5.

Sources: wago.tools/hotfixes push log for build 70707 (read 2026-11-01, pushes up to 113557); `https://wago.tools/db2/<Table>/csv?build=1.60.1.70707`; Wowhead Forever tooltips `nether.wowhead.com/forever/tooltip/spell/{1328618,408120,19979,25890,1038,20550}` (read 2026-11-01).
