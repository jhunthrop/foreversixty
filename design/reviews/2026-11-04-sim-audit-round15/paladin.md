# Sim-accuracy round 15: paladin

Lane paladin. Site wt/r15-paladin from int/r15 (fd320b60e), fork lane/r15-paladin from int/r15 (d1d9f2775). Client 1.60.1.70755. Every sim went through the lock wrapper (`sim.sh`). I made no engine change, so the fork branch is still at d1d9f2775.

## Seed items (guide-accuracy round 1)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Holy band 30 skips Oath of Light (S1).** The ladder already takes Oath at 30. The round-11 row `paladin-holy: Oath of Light 1` in data/curated/ladder-first.json (d450cf8ec, 2026-11-02) gives band 30 `50320215011100000`, Oath included. The guide auditor read `bis/paladin-holy.json` band 30 (`50320215012000000`, no Oath). That file was generated 2026-11-01T20:41, before the row, and is the plain walk. The web strip (`leveling-builds.ts` `ladderFirstFor`), addon-builds, talent-search and the BiS ranker all apply the row. In the engine Oath clearly works at 30: healing per fight 30.5k to 34.9k with Oath, band 30 bare, 3,000 iterations. | probe of `LadderTalentString` with `ForSpec`; spec-breakdown `-talents` | Aligned (stale BiS file; see Measure at refresh) | guides-handoff | - |
| G2 | **Dwarf vs Human for Retribution (S2 / PALADIN-12).** Alliance, guide raidBuild and gear, 10,000 iterations. Raid 60: Human 761.98 / 761.98, Dwarf 761.57 / 761.93 (-0.05% / -0.01%, a tie). Bare 60: Human 264.22 / 264.11, Dwarf 260.99 / 261.08 (Human +1.2%). Artisan's Precision is in the engine as 1% hit on the merged hit stat (racials.go:296-300, client 1259719). Under the race rule, Human stays. Web check, since raid ties: no Forever-specific source ranks the two. Classic-era guides favor Human for PvE and Dwarf for PvP (rankedboost.com/world-of-warcraft/classic-retribution-paladin-build, read 2026-11-03). The Forever racial overview (timesaver.gg/blog/wow-forever-racial-changes-best-race-every-class, read 2026-11-03) lists both races with no ranking. | spec-breakdown `-race` | Aligned (Human stays) | guides-handoff | - |
| G3 | **Lay on Hands mana (owner O2).** 70755 SpellMisc Attributes_1 for 633/2800/10310 is 0x20000. Era 1.15.9.70576 has 0x20002, so the client dropped the old drain-all-power bit (0x2). The Forever tooltip still says "Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration" (nether.wowhead.com/forever/tooltip/spell/633 and /10310, read 2026-11-03). That points to a server-side drain with no five-second-rule restart. The engine matches the tooltip: `SpendMana(CurrentMana())` (lay_on_hands.go:58), and `SpendMana` (core/mana.go:96-110) never touches the five-second rule. There is no SpellPower row; the cooldown is 20 min. Tooltip plus client table count as one source. | wago SpellMisc 70755 and Era; Wowhead Forever | Aligned; owner question answered as "yes, it drains" (single source) | - | - |
| G4 | **Exorcism never crits (owner O1).** The client has no can't-crit flag: SpellMisc Attributes_2 is 0 on all six ranks (879 to 10314). The engine's no-crit rule (exorcism.go:86-92) rests only on beta logs, 0 crits in 641 casts. At even 5% crit, 0 in 641 has odds of about e^-32, so the logs are strong. But this is server behaviour, and no 70755 log was checked. | wago SpellMisc 70755 | OPEN (owner; a 70755-era log of Exorcism on an undead or demon settles it) | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Holy leveling bands spent points a healer cannot use.** Band 20's 11th point went to Purifying Power (cheaper Cleanse, shorter Exorcism and Holy Wrath cooldowns; the healer casts none of them) while Illumination was open, since Reverence 2 is held and Illumination is on the 10-point row. Band 30's 19th point went to Divine Precision (Holy spell hit; heals cannot miss) while Infusion of Light's second rank was open. The plain top-down walk takes the left column first. The curated reservation is now Illumination 1, Divine Favor 1, Illumination 4, Oath of Light 1, Illumination 5, Infusion of Light 2, in that order, so each level holds every point of the level before. Band 20 becomes `50320001…` and band 30 `50320215002100000`. Oath still comes at 25 and Holy Shock at 32. Bands 35 to 60 are unchanged. Band 20 effective HPS: +3.8% alliance, +3.3% horde, both seeds. Band 30 level (Infusion of Light's second rank is worth about nothing; a free point replaced by a live one). | `TestHolyPaladinBandsSpendNoDeadPoints` failed first, then passed; `TestCuratedLadderFirstIsSound` passes; talents/paladin.json 70755 rank texts | CONFIRMED-BUG | old (round 10/11's healprofile ladders tested only Oath; missed) | S eacf50ecf |
| H2 | Divine Precision 1, Purifying Power 2 and Precision 3 at band 35 to 60 are gating filler. Light's Vigil needs 30 Holy points, and every Holy talent left open (Divine Strength, Improved Seals, Holy Conduit, Consecrated Ground) does nothing for a healer either. The Prot and Ret points beyond Divine Intellect 5, Benediction 5 and Conviction 3 are Guardian's Favor and Sacred Duty utility plus filler. No live replacement exists. | tier sums on talents/paladin.json | Aligned | - | - |
| H3 | 70755 changes nothing for Holy. A diff of every paladin spell and talent row from 70753 to 70755 (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellLevels, SpellClassOptions, SpellEquippedItems) finds one change in total: 697263 (Two-Handed Weapon Specialization, see R1). | c755 fetch raw tables | Aligned | - | - |
| H4 | Checked against the 70755 talent texts and engine code: Healing Light 12%; Reverence 15% a rank (talents.go:225); Illumination 100% chance for 50% of base cost; Divine Favor; Infusion of Light 0.5 s a rank on the next Holy Light within 15 s (healing_talents.go:84-124); Oath of Light per-level amounts (oath_of_light.go); Holy Power; Boundless Light. Rank rows: every paladin row in rank-choices.json is trainable at its band (Holy Shock 4 is spell level 56 at bare 60; Flash of Light 3 at 40). | code read; spellconst/paladin.json | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Divine Protection and Divine Shield checked against the 70755 SpellEffect rows. 498/5573: aura 25 (pacify) and immunity; the engine cancels auto attack. 642/1020: immunity plus aura 79 at -50% damage; the engine halves damage dealt. They share one 300 s category cooldown. Templar's Bulwark: 110 mana, 8 s, absorb equal to 100% of health. | wago SpellEffect 70755; defensive_cooldowns.go | Aligned | - | - |
| P2 | Band builds: the round-14 band-20 Bulwark ruling holds (`15301100…`). Band 30 takes Holy Shield and band 40 Shieldsworn. No dead point has a live alternative. Improved Hammer of Justice, Stoicism and Reckoning are not taken. | ladder probe | Aligned | - | - |
| P3 | No 70755 row changes for Protection (H3). Holy Shield crit is still OPEN from round 11. | - | Aligned / OPEN | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Two-Handed Weapon Specialization in 70755.** SpellEffect 697263 EffectMiscValue_0 changes from 1 to 3. It is the school mask of effect 0 (Physical to Physical and Holy), not an item-subclass mask as round 14's build.md X3 guessed. The engine already applies both schools (talents.go:85-111). | c755 70753 to 70755 diff | Aligned (closes build.md X3) | - | - |
| R2 | Vengeance: 3% a stack, 5 stacks, 30 s, from non-periodic crits (talents.go:113-162; client 20050 CumulativeAura 5). Closes build.md X2. | code; talents/paladin.json | Aligned | - | - |
| R3 | Raid build fillers: Vindication 3 (lowers the boss's attack power) and Repentance 1 (crowd control). The tier sums for Unsealed Wrath need four points beyond the 27 live Retribution points, and every alternative (Deflection, Unyielding Faith, Pursuit of Justice, Eye for an Eye, Turn Evil) is just as dead for damage. | tier sums | Aligned | - | - |
| R4 | Band 20 Seal of Command in place of Sanctified Judgement's first rank: -0.8% (40.46 to 40.13, alliance, 3,000 iterations). The rotation keeps Seal of Righteousness at 20, so SoC adds nothing there. | spec-breakdown | Aligned | - | - |

## Before/after

Raid 60 is identical for all three specs by construction: my only change is the Holy reservation, which leaves bands 35 to 60 unchanged. Retribution raid 60 on my branch reproduces before.tsv (alliance 762.0 at both seeds). Holy, bare, own heal profile, 10,000 iterations, effective HPS:

| band | faction | seed 1 | seed 50001 | % (1 / 50001) |
|---|---|---|---|---|
| 20 | alliance | 39.1 to 40.6 | 39.1 to 40.6 | +3.8 / +3.8 |
| 20 | horde | 38.9 to 40.2 | 38.9 to 40.2 | +3.3 / +3.3 |
| 30 | alliance | 116.3 to 116.3 | 116.3 to 116.3 | 0.0 / 0.0 |
| 30 | horde | 113.6 to 113.7 | 113.6 to 113.6 | +0.1 / 0.0 |

## Files to regenerate

- `sim/request/testdata/ladder/paladin-holy.golden.md` (TestRotationLadder; bands 20 to 30 move). web `leveling-builds.test.ts` compares against it and fails until the regen.
- `data/builds/1.60.1.70755/addon-builds.json` (Holy point order).
- Holy BiS bands 20 and 30 at the refresh (see below), plus dps-per-point and coach starters, if they carry Holy band talents.

## Measure at refresh

- Holy bands 20 and 30 on the new reservation. The published `bis/paladin-holy.json` (2026-11-01) still has the plain walk at 20 and 30. It even predates round 11's Oath row, which is why the guide audit saw no Oath at 30.

## For other lanes

- None. No shared code was touched.

## For the guide lane

- holy.md:112: "If you do level as Holy, this build picks up Oath of Light in your 30s and Holy Shock in your 40s, but questing as Holy is slow." Replace with: "If you do level as Holy, this build takes Illumination at 20, Oath of Light at 25 and Holy Shock at 32, but questing as Holy is slow."
- protection.md:109 and holy.md:92 (Lay on Hands): the drain is right. Optional addition from the Forever tooltip: "...but your mana keeps regenerating as normal afterwards."

## Still open

- G4: Exorcism crit (owner).
- Holy Shield crit (round 11 O3).
- Retribution leveling point order: the walk takes Vindication 3 before Seal of Command around levels 22 to 25. It does not show at published bands, and SoC is not used at 20, so I left it.
- Warcraft Logs is still blocked (gcloud), so no level-60 log check was possible.
