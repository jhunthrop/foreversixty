# Sim-accuracy round 22: paladin

Lane paladin. Site wt/r22-paladin from int/r22 (e3b8ce0d3: platform, items, druid, hunter and warrior merged, with the druid lane's crit-aura change), fork lane/r22-paladin from int/r22 (3c44b0559, no paladin commit). Client 1.60.1.70790; no new client data since round 21 (no commit touches the paladin talents, trainables, spellconst or hotfix-push.json). Every sim ran through `$R/sim.sh`, frozen BiS (`$R/bis-frozen`), guide raid builds (`$R/rt`) at 60, the band's published build below 60, Holy with `-heal`, Protection with `-tank`. Baseline is this branch before any change (`$R/paladin/sb-base`, `$R/paladin/base.tsv`), not `$R/before.tsv`. Lane files: `$R/paladin/` (run.sh, runb.sh, variant rotations, *.tsv, notes/, cap/).

**Result: 1 CONFIRMED-BUG, fixed (`old`).** Holy's party Flash of Light should be rank 5 at raid 60. Two mana-pacing findings (Retribution's twist and Consecration pace, Protection's half-mana Consecration gate) are measured and OPEN: the best threshold differs between the raid and bare presets and between leveling levels, and no APL value can tell those apart. They need a platform value (see "For other lanes").

## Seeds and leads

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Guide-accuracy round 22 paladin: no new sims requested; its two carried requests (the Seal of Command opener gate and Consecration ranks 1 to 4 at 50 to 53) were settled in round 21 (O1, C1). Its four prose findings (R22-1 to R22-4) are guide-lane text, already reworded on main. | design/reviews/guide-accuracy/round22/paladin.md (origin/main) | Aligned (nothing to sim) | - | - |
| L2 | Round 21 still open: Seal of Righteousness rank 8 procs on a level-60 log, Oath of Light on a low rank. Still no live Forever log. Last looked 08:10 UTC 5 Nov: api.foreversixty.gg/v1/reports/recent lists only the 20 October report; Warcraft Logs v1 zones on classic, fresh, www and sod list no Forever zone; foreverlogs.gg's API answers "API access not allowed". | live checks | OPEN (log-gated) | - | - |
| L3 | The 3% crit change (Leader of the Pack and Moonkin Aura exclusive) moved raid 60 by Holy -3.6%, Protection -2.3%, Retribution -2.4% on this branch's baseline. Near ties re-checked at raid 60, 3,000 iterations to explore, 10,000 at seeds 1 and 50001 where close: Retribution Consecration rank 4 (-1.9%) and 3, Seal of Righteousness rank 7 (-0.5%), Seal of Command rank 4 (-0.3%), the 0.8 pace against 0.6/0.7/0.9 (within error at raid); Protection Consecration rank 4 (-1.9%) and 1, Holy Strike rank 7 (-1.1%); Holy tank Flash of Light rank 4 (-0.3%), Holy Shock rank 3 (-0.1%), Holy Light rank 7 (-1.0%) and the party Flash of Light (H1). Nothing flipped; H1 predates the change. Race picks are all within 1% and stay. | `$R/paladin/base.tsv`, explore runs in the lane log | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The party Flash of Light (everyone else below 80%) should be rank 5 at raid 60.** What was wrong: the card cast rank 4 on the party at every preset. At raid 60, rank 5 heals more: 700.6/701.4 to 701.6/702.0 effective HPS alliance and 683.0/683.1 to 684.6/683.9 horde (seeds 1/50001, 10,000 iterations). Three more alliance seeds give +0.1, +1.0 and +1.7 against a seed-to-seed spread of about 0.5, so the gain is real and small (+0.1 to +0.2%). The round 21 binary (`$R/sb-before`) shows the same gain (727.1 to 728.1, 727.7 to 728.8), so the crit change did not cause it; round 20 judged only the tank row at raid 60. Bare 60 never casts the party lines (identical numbers), so the authored rank 4 stays there and below 60. **Fix:** a raid-only row (rank-choices.json, 19941 at 60, rank 5) and its embedded copy; holy.md priority step 5 and the card note say rank 5 in a raid, rank 4 without raid buffs and from 50. Test first: `TestHolyPartyFlashOfLightIsRankFiveInARaid` (red: lines 12 to 16 cast 19941; green). | `$R/paladin/holy-party5.tsv`, `holy-party5-bare.tsv`, `holy-after.tsv` | CONFIRMED-BUG (rank row, gain beyond error at both seeds and factions) | old (round 20 H1 re-judged the Holy rows at 43 to 60 and kept the party row; raid 60 was checked on the tank row only) | S c05f93f85 |
| H2 | Fresh look at the raid-60 heal fight: Holy Light 33.5 casts, tank Flash of Light 125.4, Holy Shock 4.1 (51.5% crit with Boundless Light on the party), Divine Favor 2.9 (the tank below 40%), Oath of Light once; Illumination returns half the base cost on 40 crits a fight; mana lasts 178 of 300 s on the profile. Oath, Divine Favor, Illumination, Holy Power, Boundless Light, Healing Light and Light's Vigil (never cast by the card) as round 20 H4 and round 21 N2 have them; no paladin row changed in 70790. | `-heal` breakdown | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Fresh look at the raid-60 tank fight: Holy Strike 28 casts a fight (10 s less Shieldsworn's 4), Hammer of the Righteous 27.5, Holy Shield 17.7 casts with its damage 61.9 times (blocks plus Shieldsworn's shield strikes, round 10 T1), Seal of Fury held 99.9%, Judgement of Fury 21.4, Swift Judgement 3.5, Consecration 16.7, Redoubt 30.7% uptime, Iron Creed 70.5%; boss swings 43.7% blocked, 11.2% crushing. Amounts and threat as round 20 P2 and round 21 N2. | `-tank` breakdown | Aligned | - | - |
| P2 | **The half-mana Consecration gate (row 12) is not the best rule, but no single rule beats it everywhere.** A pace line (mana percent at least the fight's remaining percent) gains at raid 60 (1125.6/1125.8 to 1128.5/1127.8 threat per second, +0.19/+0.18%) and at bare 60 (484.3/483.9 to 498.2/498.0, +2.9%), and at most leveling levels (+3 to +7% at the band floors 20, 30, 40 and 50), but loses at 26 to 29 (-2.5 to -3.6%) and 36 to 39 (-0.8 to -1.9%), 3,000 iterations seed 1, both factions. Pace 0.8 and 0.9 also lose at 26 to 29 and 39. The best threshold moves with the mana income, which an APL cannot read. | `$R/paladin/prot-lvl.tsv`, `prot-lvl2.tsv`, `pace10k.tsv`; variants `prot-p*.json`, `prot-g*.json` | OPEN (needs a platform APL value; see below) | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Round 21's changes re-read: the Consecration rows at 50 to 53 (by race) hold on this branch (bare is untouched by the crit change); Sanctified Judgement's mana makes no threat (`TestSanctifiedJudgementManaAddsNoThreat` path unchanged); the paced twist runs at raid (Seal of Command 14.5 casts, Unsealed Wrath 27.4 releases a fight, 0.1 s out of mana). | raid-60 breakdown | Aligned | - | - |
| R2 | **The 0.8 pace (twist lines 6 to 8 and Consecration line 11) overspends without raid mana, and the best pace differs by preset and level.** Bare 60, raid build: a stricter pace of 1.3 on all four lines gains 297.9/298.0 to 309.6/309.6 (+3.9%, seeds 1/50001, 10,000 iterations), but loses at raid 60 (774.8 to 737.9/737.2, -4.8%). Neither line alone helps (1.3 on the twist only 293.5, on Consecration only 302.0 at 3,000). Below 60 the Consecration line decides: dropping Consecration on one target gains at 47 (+5.1%), 53 (+2.9%) and 57 (+5.8%) but loses at 43 (-7.1%) and 33 (-3.5%), all bare, 3,000 iterations. Offset lines (mana percent at least a x the fight left minus b) do not separate the presets either. What does separate them is the mana income: adding Judgement of Wisdom alone to the bare preset makes 0.8 the best again (355.0 against 344.8), Blessing of Wisdom alone does not (317.1 against 324.0). A Judgement of Wisdom gate gives exactly the raid number and +3.9% at bare 60, but the rank rewriter turns 20355 into the rank the paladin has learned below 58 and the condition drops out, so it is not a rule a card can carry. Round 10 recorded the same limit for Consecration ("an exact per-preset optimum would need a preset value in the APL"); round 21's twist pace inherits it. | `$R/paladin/bare-grid.tsv`, `raid-grid.tsv`, `bare-buff-grid.tsv`, `lvl-cgrid.tsv`, `aff.tsv`, `pace10k.tsv`; variants `ret-t*-c*.json`, `ret-jow*.json`, `ret-aff*.json` | OPEN (needs a platform APL value) | - | - |
| R3 | Seal of Command, Judgement of Command, Seal of Righteousness, Holy Strike, Hammer of Wrath, Exorcism, Consecration, Unsealed Wrath, Vengeance, Conviction, Two-Handed Weapon Specialization, Champion of the Light, Sanctified Judgement: as round 20 R2 and round 21 N2 checked them on the 70790 rows; the raid-60 breakdown shows no unexpected cast, miss or crit rate (Seal procs and Judgements on the melee table, Consecration ticks critting on the spell table, Hammer of Wrath on the ranged table). | breakdown; engine sim/paladin | Aligned | - | - |

## Cross-cutting: Blizzard's Forever notes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | The thread (https://us.forums.blizzard.com/en/wow/t/2360696, fetched 08:00 UTC 5 Nov) is unchanged since round 21's read (same 10 posts; post 9 last edited 4 Nov 19:35, before round 21 fetched it; a line-by-line diff of every post finds no change). Every paladin line (122 lines, `$R/paladin/notes/pal.txt`) re-read against the engine: Righteous Fury 33% plus 90% Holy, Holy Strike 25 to 50% and 10 s, Holy Power's crit-healing change, Divine Precision 4/8%, Oath of Light, Boundless Light 15/30%, Redoubt (the later client's 30% wins, round 21 N3), Shield Specialization 4% no threat, Sacred Duty, Iron Creed, Shieldsworn, Conviction 2/4/6% on all spells and attacks, Two-Handed Weapon Specialization 2/4/6% Physical and Holy and -7/-14/-20% threat, Champion of the Light 12 to 60% and 0.2 s a rank, Vengeance on non-periodic crits to 5 stacks, Unsealed Wrath, Sanctified Judgement no threat, Consecration 0.042, the Judgement coefficients. All as the engine has them; the unmodelled lines (Reckoning, Judgement of Fury's taunt, Retribution Aura, Salvation, Turn Evil) have no published effect. | notes; engine sim/paladin | Aligned | - | - |

## Before/after

Raid 60, frozen BiS, guide raid builds, 10,000 iterations, seeds 1 / 50001, against this branch's own baseline (`$R/paladin/base.tsv`):
- Holy: 700.6/701.4 to 701.6/702.0 alliance (+0.14/+0.09%), 683.0/683.1 to 684.6/683.9 horde (+0.23/+0.12%) (H1).
- Protection: 1125.6/1125.8 and 1123.8/1123.3, unchanged.
- Retribution: 772.9/772.5 and 779.8/779.8, unchanged.
Bare presets and every band below 60: unchanged.

## Files to regenerate

- addon-data, then Data.lua (addonrotation): Holy's raid-60 party Flash of Light rows show rank 5, and their card note's new last clause.
- sim/request ladder golden testdata/ladder/paladin-holy.golden.md only if it carries a raid-60 row (TestRotationLadder).
- Fork: nothing.

## Measure at refresh

- Holy raid 60: the party Flash of Light row on the re-ranked gear (the gain is small, +0.1 to +0.2%).
- Carried from round 21: the Retribution gear search under the paced twist; Retribution band 50 at 53, Undead, rank 3 against rank 4.

## For other lanes

- **Platform (sim/core, the APL):** an APL value for the player's mana income, for example mana regained per second over the last N seconds, or a "time until out of mana at the current spend" value. With it Retribution's twist and Consecration pace (R2) and Protection's Consecration gate (P2) can be one buff-independent rule ("spend on filler only while the income covers the core buttons for the rest of the fight"). The measured stake: Retribution +3.9% at bare 60 and up to +5.8% on leveling bands, Protection +2.9% at bare 60 and up to +7% at band floors, raid 60 unchanged or slightly up. A preset flag would work for 60 but not for the leveling levels, where the best rule moves level by level.
- Platform, small: the rank rewriter (`sim/request/rotation_ranks.go`) rewrites an `auraIsActive` id on the target (Judgement of Wisdom 20355) to the paladin's own learned rank, so a condition on another player's debuff vanishes below its trainer level. Harmless today (no shipped card reads one); worth knowing before any card does.
- No shared code touched. No commit moves DPS outside paladin.

## For the guide lane

- Shipped: holy.md priority step 5 (Flash of Light rank 5 on the party in a raid, rank 4 without raid buffs and from 50) and the party row's card note. `_prose.test.ts` passes. The rendered card shows rank 4 until addon-data is regenerated.
- wow-player on the rendered page: SHIP (see "Review").

## Review

wow-player on the rendered Holy page (dev build of this branch, 1440 and 390, `$R/paladin/cap`): **SHIP**, no MUSTs. SHOULDs applied: step 5 names the buffs (Blessing of Wisdom, Mana Spring, the raid's Intellect buffs) and says rank 4 stays best from 50 "even once you have rank 5" (rank 4 from 50 is the measured leveling rank: rows give ranks 1 and 2 below 50). Not applied, for the guide lane: the overview's "cheap Flash of Lights" (holy.md:29) could read plain "Flash of Light"; steps 4 and 5 repeat the without-raid-buffs caveat. The card's party row shows rank 4 until addon-data is regenerated; check it flips to rank 5.

## Still open

- R2 and P2 (mana pacing; platform value needed).
- Seal of Righteousness rank 8 procs and Oath of Light on a low rank: no live log (last looked 08:10 UTC 5 Nov).
- Round 20 P3 (Protection BiS use effects, platform) unchanged.

## Sources

- Blizzard, WoW Forever Beta Development Notes, https://us.forums.blizzard.com/en/wow/t/2360696 (JSON, fetched 2026-11-05 ~08:00 UTC; identical to round 21's copy).
- Client 70790: data/builds/1.60.1.70790 (no paladin change since round 21).
- Round 10 ranks report (design/reviews/2026-11-03-sim-audit-round10/ranks.md R4 and the pace note) for the history of the Consecration pace.
