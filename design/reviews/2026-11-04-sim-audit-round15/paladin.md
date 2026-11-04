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

## Wave 2: three guide-accuracy claims checked against 70755

| # | claim | evidence and source | verdict | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **"Lay on Hands costs no mana in Forever"** (guide-accuracy O2; prose changed on land/guides-r1) | Two readings of the same client data. **Flag:** Era SpellMisc Attributes_1 for 633/2800/10310 is 0x20002 (0x2 = uses all power); 70755 has 0x20000, so the flag is gone. The guide lane read that as "free". **Text:** the 70755 client's own Spell.Description_lang for 633 and 10310 was rewritten for Forever. Era reads "Drains all of the Paladin's remaining mana when used." 70755 reads "Drains all of the Paladin's remaining Mana when used, **but does not interrupt Mana regeneration**." That is new text, written for Forever. It matches the flag change: the client-side "use all power" cost (which counts as a mana spend and starts the five-second rule) is replaced by a drain that leaves regeneration running. Wowhead Forever shows the same text (nether.wowhead.com/forever/tooltip/spell/10310, read 2026-11-03). The engine already does exactly this: `SpendMana(CurrentMana())` (lay_on_hands.go:58), and `SpendMana` (core/mana.go:96-110) never touches the five-second rule. | **Guide lane's conclusion is wrong; engine right.** Lay on Hands drains all mana and does not stop regeneration. No engine change. The guide prose must revert (see "For the guide lane, wave 2"). | - | - |
| W2 | **"Exorcism never crits"**: make the rule rest on a source | Searched for one: (1) SpellMisc Attributes_2 has no can't-crit bit (879 to 10314, 70755 and Era alike). (2) SpellCategories DefenseType is 1 (magic, crit allowed) in both builds; the table did change for Judgement of Righteousness (20281: 1 to 2, which the engine follows), so Blizzard did edit that table for paladins, but not for Exorcism. (3) The Wowhead Forever tooltip (879, read 2026-11-03) and the client description, "Causes $s1 Holy damage to an Undead or Demon target.", say nothing. (4) No web source: a search found nothing Forever-specific on Exorcism crits. (5) No new logs: api.foreversixty.gg/v1/reports/recent has one report (2026-10-20). The beta logs remain the only evidence: 0 crits in 641 hits from 39 paladins, against 18 of 367 Judgement of Righteousness crits by the same players, with about 35 Exorcism crits expected. | **Single source (server behaviour in logs); no client or Blizzard source exists.** The engine rule stays, cited to the logs (exorcism.go:86-92). It is OPEN for the owner: re-check on a launch log at 60. | - | - |
| W3 | **"Kings doesn't stack with Heart of the Lion"** | No client exclusivity row: the tooltips of 20217, 25898, 409580 and 409583 are silent, and Forever's 409583 equals SoD's apart from the attack power. Two SoD sources, and Forever ships the SoD spell unchanged: the Blizzard forums thread "Blessing of Kings doesn't stack with Aspect of Lion BUT" (us.forums.blizzard.com/en/wow/t/1738967, 2023-12-11: Kings replaces one of the two Lion buffs; the hunter keeps his extra 10%), and Game Rant's SoD rune ranking (gamerant.com/wow-season-of-discovery-best-sod-runes-ranked-each-class/: "which doesn't stack"). The engine multiplied them: a priest with both read 1.21x every stat, a hunter with Kings 1.331x. | **CONFIRMED-BUG (two SoD sources; no Forever source reports a change).** Kings and the Lion area aura's stats now share one exclusive category. The Lion's attack power (a separate effect) and the hunter's own 409580 still apply. `TestKingsAndHeartOfTheLionDoNotStack` (sim/kings_heart_of_the_lion_test.go) failed first (1.21 and 1.331), then passed; `TestHeartOfTheLion*` still pass. **Shared code (sim/core/buffs.go): every spec.** | old | E fe4b177ac |

**W3 reach (flagged).** The change only matters where Kings and the Lion meet. The raid preset carries the Lion and not Kings (presets.json), and no kit buff or APL casts Kings, so no published number moves. Raid 60, all 28 specs, both factions, seeds 1 and 50001, 10,000 iterations, against before.tsv: see "Wave 2 measure". It moves the BiS tool only when a user adds Kings to a raid preset with the Lion. Not changed: the paladin's own cast-applied Kings (sim/paladin/blessings.go `kingsEffect`) still multiplies with the Lion. No APL casts it, so nothing moves today, but the next blessing change should route it through the same category.

**For the guide lane, wave 2** (both lines are on origin/land/guides-r1):
- protection.md:109: replace "Lay on Hands heals you to full on a 20 minute cooldown, and in Forever it costs no mana, so you keep your Holy Shield and your threat after it. Save it for the swing that would kill you." with "Lay on Hands heals you to full on a 20 minute cooldown, but it takes every point of your mana with it. Your mana keeps regenerating afterwards, but until it comes back you have no Holy Shield and little threat, so save it for the swing that would kill you and have a mana potion ready."
- holy.md:92: replace "Lay on Hands is a full heal on a 20 minute cooldown, and in Forever it no longer empties your mana bar, so there is no reason to hold it for the end of a fight. Use it on anyone about to die whom Divine Favor and Holy Light could not have saved." with "Lay on Hands is a full heal on a 20 minute cooldown, and it empties your mana bar, though your mana keeps regenerating as normal afterwards. Use it on someone about to die whom Divine Favor and Holy Light could not have saved, ideally late in the fight."
- The guide-accuracy report's O2 answer and its "needs sim round" note (drop the SpendMana line) should be withdrawn.

**Wave 2 measure.** `measure.sh` with the lane binary (fork fe4b177ac + site wt/r15-paladin), raid 60: 28 specs × 2 factions × 2 seeds = 112 rows, all identical to before.tsv (0 differ, 0 errors). Output: scratchpad r15/paladin/after-w2.tsv.

**Wave 2 counts.** One CONFIRMED-BUG (W3, old, E fe4b177ac, shared code). W1 reverses the guide lane's prose change, so the engine is unchanged. W2 stays a single-source owner item.

## Wave 3: rank choices on the level-gap rule

Branches merged int/r15: site 1e96d8df2, fork 9c511a2cd. The rule: a rank keeps its whole spell-power or healing bonus until the caster is 17 levels past the rank's learn level, then loses 5% a level (`core.LevelPenalty`; priest lane E 8d9ca6aae, 974614b8b). Method:
- Sweep: one variant binary per alternative rank, with sim/rankchoices/rank-choices.json edited for the build and then restored. Alliance, 3,000 iterations.
- Confirmation: 10,000 iterations at seeds 1 and 50001.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W3-1 | **Consecration never took the level-gap rule.** Its first-targets tick adds 0.042 × spell power by hand (consecration.go `consecrationTickDamage`), outside the spell's BonusCoefficient that `levelPenalty` scales. So rank 1 (learned at 20) kept its whole coefficient at 40, 50 and 60, and Retribution's numbers did not move at all under the new rule. The coefficient is now `0.042 × core.LevelPenalty(rank level, caster level)`. | `TestConsecrationCoefficientFollowsTheLevelGap` failed first (undefined helper), then passed; sim/paladin, retribution and protection packages pass, and no fork golden moved | CONFIRMED-BUG | R15 wave 2 (8d9ca6aae's rule did not reach a hand-rolled coefficient) | E 86fe4bf34 |
| W3-2 | **Holy Shock rank 1 at raid 60.** It is learned at 30, so at 60 it keeps 35% of its healing bonus. Ranks 2/3/4 at raid 60 (3,000 iterations): 658.2 / - / 659.1 against 656.1. At 10,000 iterations rank 4 gains +0.50 / +0.58% alliance and +0.52 / +0.42% horde (seeds 1 / 50001). Mana still runs dry before the end (154 to 145 s against 158 to 147 s), so the row rule's mana clause holds. The bare-60 row widens to both presets. | `TestHolyPaladinHolyShockAtBareSixty` now asserts rank 4 at raid too; it failed first, then passed | CONFIRMED-BUG | R15 wave 2 (the rule makes rank 1 lose) | S e2d3455ea |
| W3-3 | **Protection rows hold.** Consecration rank 1 at 40/50/bare 60 beats ranks 2 to 5: threat 191.3 / 278.5 / 444.1 against 186.5 to 188.0 / 271.6 to 273.5 / 427.8 to 438.2 (3,000 iterations), because Protection has almost no spell power at these bands. Holy Strike rank 1 at 40 beats rank 5 by +0.58 / +0.68% (10,000 iterations, both seeds) and ranks 2 to 4 at 3,000. | variant sweeps | Aligned (rows kept) | - | - |
| W3-4 | **Retribution rows hold after W3-1.** Consecration rank 1 at 40: 112.8 / 112.8 against rank 3 (max) 111.3 / 111.2 (+1.3%). At 50: 166.5 / 166.5 against max 161.6 / 161.7 (+3.0%); ranks 2 and 3 lose too. The 30-39 row (gap under 17) and the bare-60 rank 4 row (learned at 50) are untouched by the rule. | variant sweeps, 10,000 iterations | Aligned (rows kept) | - | - |
| W3-5 | **Holy rows re-picked by the priest lane do not lose.** Against no Holy rows at all: bare 20 +14.8%, 30 +29.4%, 40 +16.2 to +16.3%, 50 +6.4 to +6.5%, bare 60 +0.6% (both seeds). Fixed-rank casts in the Protection and Retribution APLs are authored at their top ranks, so they cast the top rank learned at each band, and a top rank is never penalised. | 10,000 iterations | Aligned | - | - |

**Wave 3 before/after** (alliance, 10,000 iterations, seed 1 / 50001). Before = $R/sb-before (round start). After = this branch.

| spec | band | before | after | % |
|---|---|---|---|---|
| Retribution | bare 30 | 72.5 / 72.6 | 72.5 / 72.6 | 0 / 0 |
| Retribution | bare 40 | 113.7 / 113.7 | 112.8 / 112.8 | -0.8 / -0.8 |
| Retribution | bare 50 | 174.1 / 174.2 | 166.5 / 166.5 | -4.4 / -4.4 |
| Retribution | bare 60 | 264.2 / 264.1 | 264.2 / 264.1 | 0 / 0 |
| Retribution | raid 60 | 762.0 / 762.0 | 762.0 / 762.0 | 0 / 0 |
| Protection (threat) | bare 40 | 191.2 / 191.5 | 191.2 / 191.5 | 0 / 0 |
| Protection | bare 50 | 278.0 / 278.2 | 277.9 / 278.2 | 0 / 0 |
| Protection | bare 60 | 443.9 / 443.3 | 443.8 / 443.2 | 0 / 0 |
| Protection | raid 60 | 1113.8 / 1112.5 | 1113.8 / 1112.5 | 0 / 0 |
| Holy (eff. HPS) | bare 20 | 39.1 / 39.1 | 38.0 / 38.0 | -2.8 / -2.8 |
| Holy | bare 30 | 101.7 / 101.7 | 98.6 / 98.5 | -3.1 / -3.2 |
| Holy | bare 40 | 190.7 / 190.5 | 145.5 / 145.4 | -23.7 / -23.7 |
| Holy | bare 50 | 218.2 / 218.1 | 158.4 / 158.3 | -27.4 / -27.4 |
| Holy | bare 60 | 350.8 / 350.6 | 246.4 / 245.7 | -29.8 / -29.9 |
| Holy | raid 60 | 676.5 / 677.0 | 659.6 / 660.3 | -2.5 / -2.5 |

What each row means:
- **Holy bands 20 and 30** still use the published (stale) band talents. On the wave-1 walk, band 20 is +3.3 to +3.8% on top of these.
- **Holy's large bare losses are the rule itself**, not a rank choice. Bare 40-60 keeps low-rank heals for mana, and those heals now lose most of their bonus. The priest lane's re-pick already recovers +6 to +29% over no rows (W3-5).
- **Holy raid 60**: the rule costs -3.0%, and Holy Shock rank 4 wins back +0.5%.
- **Retribution's -0.8% and -4.4%** come from Consecration rank 1's spell-power share fading at 40 and 50. Rank 1 still wins there on mana.

**Wave 3 files to regenerate:** TestRotationLadder goldens for paladin-retribution (bands 40, 50) and paladin-holy (raid 60), and the Holy and Retribution BiS bands at the refresh.

**For other lanes:** any class code that adds a spell-power share by hand, outside BonusCoefficient or a Dot or Hot BonusCoefficient, misses the rule the same way. Paladin's only such site was Consecration (grep for `GetBonusDamage(` and `GetBonusHealing(` inside OnTick or ApplyEffects).

**Wave 3 follow-up (S 416b2280f, E 932a0357b).** `TestCuratedRankChoicesAreCastableDownranks` refuses a row at 60 for both presets, so W3-2 moved into the rotation:
- Holy's six Holy Shock lines (curated, sim/request and fork copies) now name rank 4 (20930), and the Holy Shock row is gone. The APL notes are untouched.
- Below 60 the top rank learned ties rank 1 (3,000 iterations: band 30 98.7 / 98.7, 40 145.4 / 145.5, 50 158.3 / 158.5), so no row keeps rank 1.
- Raid 60 659.6 and bare 60 246.4 are unchanged from the row.
- The smoke test's talent excuse now names 20930.
- Tests that pass: the rank-choice test, `TestHolyPaladinHolyShockAtBareSixty`, `TestHealerListsCastTheTopTrainerRankUnlessChosen`, the paladin rotation smoke tests, and `make apl-check`.
- Not mine: `TestEveryWrittenRotationRunsInTheEngine/hunter-marksmanship` fails on int/r15 with an unresolved 1310786.
