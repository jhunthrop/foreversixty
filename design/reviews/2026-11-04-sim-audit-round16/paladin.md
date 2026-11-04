# Sim-accuracy round 16: paladin

Lane paladin. Site wt/r16-paladin from int/r16 (c62d77ab8), fork lane/r16-paladin from int/r16 (9f7f1f6e4). Client 1.60.1.70755 with the late hotfix re-layer (43dc0f074). Every sim and sim-running test went through the lock wrapper. Client rows re-fetched from wago.tools for 1.60.1.70755 on 2026-11-04 (SpellEffect, SpellAuraOptions, SpellMisc, SpellCooldowns, SpellPower, SpellCategories, SpellLevels, SpellClassOptions).

One CONFIRMED-BUG, fixed (seed 3). No published number moves. No APL, build string or rank row changed, so no guide text ships with this lane.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Exorcism and Holy Shield crit on a level-60 log: no log exists yet.** api.foreversixty.gg/v1/reports/recent returns one report (2026-10-20). foreverlogs.gg report ids run to 2813 (2026-10-29, title "shield"); ids 2814 to 2849 and 2850, 2900, 3000 to 5000 return "Report not found". The list endpoint needs an API key we do not have. Warcraft Logs is still blocked (gcloud). Client rows unchanged: Exorcism 10314 SpellMisc has no can't-crit bit; Holy Shield's damage is its own aura-43 effect (coefficient 0.08) on all four ranks, which the engine rolls on the spell table. | foreverlogs API probe 2026-11-04; wago SpellMisc/SpellEffect 70755 | OPEN (both). Round 15's call stands: Exorcism never crits, Holy Shield crits. A launch log settles both. | - | - |
| S2a | **Crusader opener and both Crusader judgement fixes hold.** Bare 60 alliance, guide raid build, frozen gear: 290.1 DPS (round 15 reported 288.5 on its gear); exactly one Seal of the Crusader and one Judgement of the Crusader per fight; the debuff is one aura with the top learned rank's bonus (20303, 161), refreshed by the paladin's own melee and skipped when the raid's debuff is up (prepull condition on 20303). 70755 text for 20308 confirms 40 s, refreshed by melee, Judgement does not consume the Seal. `TestOwnCrusaderJudgementUsesTheTopRankBonus`, `TestCrusaderReachesConsecrationTicks` pass. | spec-breakdown; spelltips 20308/20271; fork tests | Aligned | - | - |
| S2b | **Consecration under the low-rank penalty** holds: the first targets' 0.042 share is the dot's BonusCoefficient, so core's level penalty and Judgement of the Crusader both reach it. `TestConsecrationCoefficientFollowsTheLevelGap` passes. Client 20924: 96 + 216 over 8 s, 565 mana, 8 s cooldown. | fork tests; spelltips | Aligned | - | - |
| S2c | **Holy Shock rank 4** in all six Holy lines (20930: 325 mana, 348 damage / 320 heal, 10 s). Libram of Infusion (279248) +6% crit reaches the heal half (client mask 0x200000/0x10000; test passes). | spelltips; relic_mods_auto_gen.go | Aligned | - | - |
| S2d | **Stoneform on the GCD and the Dwarf Protection line** hold; `TestStoneformMatchesTheClientRow` passes and the Protection list keeps the below-70% line. | fork tests; curated APL | Aligned | - | - |
| S2e | **Raid Kings exclusive with Heart of the Lion** holds (`TestKingsAndHeartOfTheLionDoNotStack` passes). | fork test | Aligned | - | - |
| S2f | **Holy leveling reservation** (Illumination 1, Divine Favor 1, Illumination 4, Oath of Light 1, Illumination 5, Infusion of Light 2) unchanged on int/r16; bands 35 to 60 untouched. | data/curated ladder rows | Aligned | - | - |
| S3 | **A paladin's own cast Blessing of Kings still stacked with Heart of the Lion.** Round 15 put the raid Kings and the Lion's area stats in one exclusive category, but the cast-applied Kings (sim/paladin/blessings.go `kingsEffect`) multiplied its own +10% on top: a paladin standing in a hunter's Lion aura read 1.21x every stat after casting Kings on himself. In game Kings and the Lion's shared 10% do not stack (round 15 W3's two SoD sources; Forever ships 409583 unchanged). The cast aura now joins the same category through a new exported core helper `core.KingsStatsEffect`. | `TestCastKingsDoesNotStackWithHeartOfTheLion` (sim/paladin/retribution/kings_cast_lion_test.go) failed first (1.10x with the Lion), then passed; without the Lion the cast still gives 1.10x and fades cleanly. Kings/Lion core tests and all four paladin packages pass (`-tags=with_db`). | CONFIRMED-BUG | old (round 15 W3 saw it and left it, no rotation casts Kings) | E 913221285 |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Oath of Light's flat bonus skips the low-rank penalty.** healing.go:92 and holy_shock.go:158 add `coefficient × oathOfLightBonus` by hand, outside the spell's BonusCoefficient, so core's level-gap penalty never reaches it (the same pattern as round 15's Consecration, which the grep advice there missed because this path calls no GetBonusHealing). Consequence: Holy Light rank 1 on the Oath target at 40 gets 267 × 0.714 = about 190 on a 43 base for 35 mana, which is why rank-choices.json keeps Holy Light rank 1 at bands 20-29 and 40-49 for the tank lines, including the Divine Favor emergency heal below 40%. No source settles how the Oath's amount is treated: Blessing of Light in vanilla added its full amount with no coefficient, and patch 2.3.0 first penalised low ranks ("Lower ranks of Flash of Light and Holy Light are now properly penalized when used with this Blessing", warcraft.wiki.gg/wiki/Blessing_of_Light, read 2026-11-04). The engine's coefficient-only reading matches neither. ForeverChanges (foreverchanges.pro/downranking, updated 2026-10-16) does not mention Oath or flat bonuses; a web search found no Forever source (2026-11-04). | code read; web sources above | OPEN. A Forever log of a low-rank Holy Light on the Oath target at 40+ settles it. If the Oath follows the penalty, apply `core.LevelPenalty` to the Oath term and re-pick the Holy Light rows at 20-29 and 40-49. | - | - |
| H2 | Checked against 70755 rows and texts: Holy Light 1-9 and Flash of Light 1-6 amounts and costs; Holy Shock 1-4; Divine Favor; Illumination (100% for 50% of base cost at rank 5); Infusion of Light (1.0 s); Boundless Light (30%); Holy Power (50%); Healing Light (12%); Reverence (30%); Light's Vigil (not cast; round 9 P3's reasoning holds); Dormant Heart of the Mountain is modelled in common/item_effects/client_procs.go (the "NOT implemented" comment at common/item_effects.go:3756 is stale, cosmetic); Hand of Edward the Odd is a melee proc and idles for a healer; Uther's Strength is a melee-struck absorb and idles. ForeverChanges' "a heal over time gets its full bonus on every rank" line is contradicted on the same page ("at level 30, Rejuvenation rank 1 keeps 55% of its bonus"), so core's penalty on HoTs stays. | spelltips, talents/paladin.json, wago SpellEffect; bis-frozen holy 60 | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Reckoning 5 in place of Anticipation 5** (`15302102513001531`, legal: tier gates 23/29/32). Raid 60 alliance, 10,000 iterations: threat 1151.9 / 1152.2 to 1180.7 / 1180.9 (+2.5% both seeds), chance of death 0% both, damage taken +5.4%. Bare 60: threat +0.6 / +0.7%, chance of death 26.7 / 26.8% to 48.4 / 47.1%. | spec-breakdown `-tank -talents` | Aligned: not adopted (the bear Nature's Reach precedent: a threat gain that nearly doubles chance of death is rejected) | - | - |
| P2 | Benediction 5 in place of Deflection 5: raid threat 1152.0 to 1150.5, damage taken 367 to 396 (3,000 iterations). Protection is not mana-bound at raid. | spec-breakdown | Aligned | - | - |
| P3 | Righteous Fury 25780: two aura-10 effects, Holy 90 and all schools 33, which the engine multiplies (Holy 1.9 × 1.33). WoW applies each MOD_THREAT aura as a separate multiplier, so this is right. Holy Shield 1-4: 4 charges on every rank (SpellAuraOptions ProcCharges 4 on 20925/20927/20928/1239570, though the tooltips of ranks 2 and 3 omit the line), 30% block, 75/110/153/221, coefficient 0.08. Seal of Fury 7: proc 20418 35 + 0.1, Judgement 20414 153 + 3.69/level, 0.429. Redoubt, Reckoning, Iron Creed, Shield Specialization, Toughness, Anticipation, Sacred Duty, Improved Righteous Fury: engine per-rank constants match the rank-1, rank-2 and rank-5 texts. | wago SpellEffect/SpellAuraOptions 70755; tank_talents.go | Aligned | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Checked against 70755: Judgement of Righteousness 20286 (170 + 4.1/level, 0.429, melee table); Seal of Righteousness proc (1786 + 47/level dummy, 0.1 after Blizzard's 16 October fix); Seal of Command 20424 (70%, 0.29, 1 s ProcCategoryRecovery); Exorcism 10314 (502 + 3.2/level, 0.429, 15 s); Hammer of Wrath 24239 (498, 6 s, instant with Champion of the Light 5); Holy Strike 10333 (50% weapon + 93, 0.429); Consecration 5; Unsealed Wrath 1310735; Holy Conduit and Improved Seals masks; Sanctified Judgement; Vengeance; Two-Handed Weapon Specialization. Raid BiS: Highlander's 2-piece is +5 Stamina only; Thunderbrew's Boot Flask and Cloak of Flames are modelled. | wago rows; code read; bis-frozen | Aligned | - | - |
| R2 | 70755 late re-layer (43dc0f074): no paladin spellconst, spelltips or talent change; paladin item and set text changes are whitespace only. | git show 43dc0f074 | Aligned | - | - |

## Wave 2: Holy rank rows on the MaxLevel + 12 rule

Merged int/r16 (site 88ebcb13e, fork ddc356138; the priest lane's E 18815822c keys the low-rank penalty on the rank's client MaxLevel + 12). Variant binaries with sim/rankchoices/rank-choices.json edited per variant and restored; Holy bare, effective HPS, own heal profile.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W2-1 | **Holy Light at 40 to 49: rank 2 now beats the rank-1 row.** 3,000 iterations, alliance, seed 1: rank 1 145.8, rank 2 149.3, ranks 3 to 6 146.9. 10,000 iterations: alliance 145.9 / 145.9 to 149.5 / 149.3 (+2.5 / +2.3%), horde 141.4 / 141.4 to 142.6 / 142.6 (+0.85 / +0.85%), seeds 1 / 50001. Both races gain beyond error, so the row is not race-keyed. | variant runs; `TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus` now expects rank 2 at 40 and 49 (it failed on the new row first, as written), rankchoices, rank-choice and preset-rank tests pass | CONFIRMED-BUG (row stale under the new rule) | R16 18815822c | S (this commit) |
| W2-2 | Holy Light at 20 to 29 keeps rank 1: 48.8 against 41.9 for ranks 2 and 3 (seed 1, 3,000 iterations; +16%, far beyond error). | variant runs | Aligned (kept) | - | - |
| W2-3 | Flash of Light (19941 line) at 40 to 49 keeps rank 1: 145.8 against 144.0 (rank 2) and 140.5 (rank 3). The 19942 rows at 26 to 39 and 42 to 49 do not apply at band 40 (it casts the top rank learned, 3), unchanged. | variant runs | Aligned (kept) | - | - |

Guide text: none needed. No guide prose or card note names a leveling Holy Light rank (holy.md:76 and :81 and the card note speak to rank 8 at 60), and the web does not render rank rows, so the rendered page is unchanged and there was nothing for wow-player to review.

Wave 2 files to regenerate: sim/request/testdata/ladder/paladin-holy.golden.md (band 40), and addon outputs that read rank rows (data/pipeline/addonrotation.py: addon-data / Data.lua).

## Before/after

Raid 60, 10,000 iterations, frozen BiS, guide raid builds (`measure.sh`, ALLRT=1). Identical to before.tsv on all 12 rows (0.00%):

| spec | alliance seed 1 / 50001 | horde seed 1 / 50001 |
|---|---|---|
| Holy (effective HPS) | 671.7 / 671.7 | 655.6 / 655.3 |
| Protection (threat/s) | 1151.9 / 1152.2 | 1150.0 / 1149.5 |
| Retribution | 788.5 / 788.0 | 795.8 / 795.4 |

## Files to regenerate

- None. No fork golden moved.

## Measure at refresh

- Nothing new.

## For other lanes

- E 913221285 adds `core.KingsStatsEffect` to sim/core/buffs.go (shared file) and splits `makeExclusiveBuff` into a build-phase wrapper around `newExclusiveStatsEffect`; behaviour of every existing caller is unchanged. Only a cast Blessing of Kings uses the new helper; no rotation casts it, so no spec's number moves.
- Cosmetic: the comment at sim/common/item_effects.go:3756 says Dormant Heart of the Mountain is not implemented; it is (client_procs.go:310).

## For the guide lane

- Nothing. No APL, build string, rank row or card note changed.

## Still open

- S1: Exorcism and Holy Shield crit (no level-60 log yet; foreverlogs has nothing after 2026-10-29).
- H1: whether Oath of Light's flat bonus takes the low-rank penalty (and with it the Holy Light rank 1 rows at 20-29 and 40-49).
- Guide-accuracy carried sims (Retribution race on a Beast target; Protection threat by ability; the Dwarf raid edge) were not re-run: race picks are not engine bugs.
- Warcraft Logs still blocked (gcloud).
