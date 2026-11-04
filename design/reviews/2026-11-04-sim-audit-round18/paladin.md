# Sim-accuracy round 18: paladin

Lane paladin. Site wt/r18-paladin from int/r18 (a6dc2019c), fork lane/r18-paladin from int/r18 (5a063241f). Client 1.60.1.70755 with the late re-layer. Every sim and sim-running test went through the lock wrapper. Measurements on frozen BiS ($R/bis-frozen) and guide builds ($R/rt).

**Result:** 2 CONFIRMED-BUGs, both in Seal of Righteousness's spell damage coefficient (S1, S2), and one rotation change that follows from them (S3: Retribution keeps rank 7 from 58). Reckoning was measured against Conviction and Deflection and loses on threat and survival (P1). The Oath of Light question is still open: there is no owner result and no log yet (H1).

## Seed 1: Seal of Righteousness coefficients, every rank

Client rows: wago SpellEffect 1.60.1.70755 (fetched 2026-11-04), EffectBonusCoefficient on effect 0. The same values appear in every Forever build's spellconst/paladin.json (69893 to 70755). There is no region-91 hotfix row for any of these effects (hotfix-push.json).

| rank | Seal aura | aura coefficient: Era 1.15.9.70576 / Forever 69893-70374 / Forever 70425+ | swing proc | proc coefficient: Era / Forever 69893-70374 / Forever 70425+ | engine before | engine after |
|---|---|---|---|---|---|---|
| 1 | 20154, 21084 | 0.029 / 0.1 / 0.1 | 25742, 25741 | 0.029 / 0.1 / 0 | 0.029 | 0.1 |
| 2 | 20287 | 0.063 / 0.1 / 0.1 | 25740 | 0.063 / 0.1 / 0 | 0.063 | 0.1 |
| 3 | 20288 | 0.093 / 0.1 / 0.1 | 25739 | 0.093 / 0.1 / 0 | 0.093 | 0.1 |
| 4-7 | 20289-20292 | 0.1 / 0.1 / 0.1 | 25738-25735 | 0.1 / 0.1 / 0 | 0.1 | 0.1 |
| 8 | 20293 | **0 / 0 / 0** | 25713 | 0.1 / 0.1 / 0 | 0.1 | **0** |

How the server reads these values:
- Before 70425, ranks 1 to 7 carried 0.1 on both the aura and the proc.
- Blizzard's beta note of 16 October says: "Fixed an issue causing Seal of Righteousness to gain almost twice as much benefit from Spell Damage than intended." The 70425 client is the build that zeroed every proc. It is the same build that carries the notes' Judgement of Righteousness 0.429 (sor.go:30).
- So the server adds the aura's coefficient and the proc's. If it read only the proc, zeroing the procs would have removed all scaling. If it read only the aura, there would have been no double to fix.
- Rank 8's aura was 0 in Era and is still 0. Once the proc was zeroed, rank 8 kept no spell damage at all.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Seal of Righteousness ranks 1 to 3 used Era's downranked coefficients** (0.029, 0.063, 0.093). Forever gives every rank from 1 to 7 a coefficient of 0.1. | Client rows in the table above. Forever's own new tooltip agrees: it gives every rank the same spell power term (SpellDescriptionVariables 1023, `$1h=0.0765*$SPH`, `$2h=0.12*$SPH`; Era's description has no spell power term). Test first: `TestSealOfRighteousnessCoefficientFollowsTheClientPerRank` (rank 3 at 25) measured 0.093, then 0.100. | CONFIRMED-BUG | old (the Classic port; round 4 P1 fixed the double benefit but kept the per-rank values; no later round compared ranks) | E af69f750a |
| S2 | **Seal of Righteousness rank 8 got 0.1 of spell damage; the client gives it 0.** A level-58+ paladin in rank 8 gets nothing from spell damage on the swing proc (the Judgement keeps its 0.429). | Aura 20293 effect 0 has coefficient 0 in every Forever build. Proc 25713 has 0 since 70425. See the mechanism above. **Against it:** the Forever tooltip's spell power term (variable set 1023) is shared by all eight ranks, so the in-game rank 8 tooltip still shows 0.0765/0.12 x spell power. That is a hand-written formula, not the coefficient the server reads. Single client source plus Blizzard's note; no level-60 log exists yet. **Settles it:** a level-60 Retribution log, comparing Seal of Righteousness rank 8 proc damage at two spell damage values (or rank 7 against rank 8 on the same weapon). Test first: same test, rank 8 at 60, measured 0.100, then 0.000. | CONFIRMED-BUG (on the client row; flagged for a level-60 log) | old in the engine. The data has been this way since 70425; the engine's Era proc values hid it, and round 4 (bb9912dc0) kept rank 8's 0.1. | E af69f750a |
| S3 | **Retribution should keep rank 7 from 58.** With S2, rank 7's 0.1 coefficient beats rank 8's bigger base. The rotation's two `castPaladinPrimarySeal` lines (the Crusader swap and the refresh) always cast the top rank learned, and no rank row could reach them. They are now ranked `castSpell` lines. The whole rotation authors Seal of Righteousness at rank 7 (20292). A rank row at 60 for both presets has to live in the rotation itself (`TestCuratedRankChoicesAreCastableDownranks`), and below 58 the rewrite already casts the highest rank learned. | 10,000 iterations, seeds 1 and 50001, frozen BiS, guide builds. Rank 7 against rank 8 (after S2): raid 60 alliance 785.0/784.5 against 775.1/774.5 (+1.28%/+1.29%); horde 792.4/791.9 against 782.6/782.2 (+1.25%/+1.24%); bare 60 alliance 298.0/298.0 against 288.1/288.1 (+3.44%); horde 296.5/296.5 against 284.5/284.5 (+4.2%); band 50 bare at 58: 201.6/201.5 against 192.7/192.8 (+4.6%); at 59: 201.5/201.5 against 194.7/194.8 (+3.5%). No race key is needed (both factions agree). Tests: `TestRetributionKeepsSealOfRighteousnessRank7From58` (57 to 60, bare and raid, human, dwarf and undead) and `TestRetributionCastsNoPrimarySealAction`. Both failed on the old rotation and pass now. | CONFIRMED (follows S2) | old (follows S2) | S b193ac8d8, E ad9dbe2eb (fork APL copy) |

Guide text shipped with S3 (retribution.md):
- :89 "At 58, stay on rank 7. In Forever rank 8 gets nothing from your spell damage, even though its tooltip still lists it, and the spell damage Champion of the Light gives you alone makes rank 7 hit harder on every swing. Training rank 8 puts it on your bar in place of rank 7, so drag rank 7 back from your spellbook."
  - Check: at 60, rank 8's base is about 8 damage a swing above rank 7 on a 3.5 s two-hander (1.2 x 0.01 x 191 x 3.5). At 0.11 a point, rank 7 passes it at about 75 spell damage, and Champion of the Light 5 alone gives more than that.
- :91 "Switch to Seal of Righteousness (rank 7)".
- :103 "it scales with your weapon's speed and your spell damage (rank 8 does not, so you keep rank 7), not attack power".

Card notes changed:
- line 2: "Rank 7 from 58, straight after the opening Crusader Judgement: rank 8 gets nothing from your spell damage. ...". This line is now a Seal of Righteousness cast on the card.
- line 5: "150 to 190 mana" (Seal of Righteousness rank 7 costs 170 and Seal of Command 210, less Benediction).
- line 8: "From 58 use rank 7: rank 8 gets nothing from your spell damage."

`_prose.test.ts` passes. The wow-player verdict is under "For the guide lane".

Not changed, and open: the one-hand/two-hand split on the spell damage term. The engine uses 0.1 with a one-hander and 0.11 with a two-hander ("from testing in SoD", sor.go). Forever's new tooltip says 0.0765 and 0.12 for every rank. That is a description formula, not the coefficient, so it is OPEN; the same level-60 log settles it.

## Seed 2: Reckoning against Conviction plus Deflection (Protection)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Guide build `05005/15302102513501031/5003` against Reckoning 5 in place of Conviction 3 and Deflection 2 (`05005/15302102513501531/3`). Dwarf, alliance, 10,000 iterations, seeds 1 and 50001. **Raid 60:** threat 1151.9/1152.2 against 1126.6/1127.2 (Reckoning -2.2%); damage taken 367.0/367.1 against 382.2/381.5; chance of death 0% for both. **Bare 60:** threat 484.3/483.9 against 469.0/469.3 (-3.1%); chance of death 26.7%/26.8% against 39.3%/39.1%. Reckoning loses on threat and on survival. | spec-breakdown -tank, `$R/paladin/reck.txt` | Aligned (build stays) | - | - |

Guide note: int/r18's protection.md:82 says Reckoning is "worth less than Conviction". That is now measured true. Guide round 15's replacement (on land/guides-r15) is also true. Either wording holds.

## Seed 3: Oath of Light

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | No owner result yet (it is item 1 on the owner's post-launch list). No level-25+ Forever log: api.foreversixty.gg/v1/reports/recent has only the 2026-10-20 report. Rows left as they are, per the brief. | API probe 2026-11-04 | OPEN (carried, round 17 O1/O4) | - | - |

## Fresh re-audit

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Coefficients against the 70755 rows. All match the engine:<br>- Exorcism 0.429, Holy Wrath 0.19, Hammer of Wrath 0.429 (24239/24274/24275), Holy Shield 0.08.<br>- Holy Shock 0.429 damage and heal (1311604-25913).<br>- Holy Light 0.714, Flash of Light 0.429, Light's Vigil heal 0.143.<br>- Judgements of Command, Righteousness and Fury 0.429; Seal of Command proc 0.29; Hammer of the Righteous 0.<br>- Consecration ticks: 0.042 on the second damage effect only (1280345-1280349).<br>The 0.15 on 429151 and Light's Vigil's damage ids (1310914, 1311592, 1311598) are not in the Forever kit. lights_vigil.go documents the latter. | spellconst/paladin.json 70755; engine sim/paladin/*.go | Aligned | - | - |
| A2 | **Seal of Fury follows the same aura-plus-proc pattern.** Its aura effect 0 carries 0.09 on ranks 1, 2, 3, 5 and 6, **0.9** on rank 4 (20419, almost certainly a typo) and 0 on rank 7 (20423). Every proc (1311647 to 20418) carries 0.1. If the server adds the two as it does for Seal of Righteousness, ranks 1 to 6 get 0.19 (rank 4: 1.0). Rank 7 gets 0.1, which is what the engine models. So Protection at 60 is the same either way, and only the leveling bands could move. The tooltip shows only the proc's flat amount, so nothing settles it. | spellconst/paladin.json 70755 (20163, 20419-20423, 1311649, 1311656) | OPEN (a level 34-41 Protection log with Seal of Fury settles rank 4) | - | - |
| A3 | Unsealed Wrath's stored Righteousness copy casts the stored rank's proc. With rank 7 active it stores rank 7, which the breakdown shows (25735, avg 208 at raid 60). | spec-breakdown | Aligned | - | - |
| A4 | Seal of Righteousness scaling levels: rank 7 grows to 56 (client MaxLevel 56); rank 8's 60 is equivalent at any level up to 60 (client MaxLevel 64). Rank 7 at 60 is 4 levels past its MaxLevel, short of the 12-level low-rank grace, so it keeps its whole coefficient. | SpellLevels 70755 (103342, 103839); core.SpellLevelPenalty | Aligned | - | - |
| A5 | Holy, Protection and Retribution talent builds and the round 16/17 regression targets were re-read; nothing in the paladin tree, SpellEffect or SpellPower rows moved since round 17 (no new 70755 push). | round 17 H1-H4, P1-P2, R1-R2 | Aligned | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS, 10,000 iterations. Before = $R/before.tsv (b2c4a2d05 on 4600b7d82). After = fork ad9dbe2eb on site b193ac8d8.

| spec | alliance seed 1 / 50001 | horde seed 1 / 50001 |
|---|---|---|
| Retribution | 788.5 to 785.0 (-0.44%) / 788.0 to 784.5 (-0.44%) | 795.8 to 792.4 (-0.43%) / 795.4 to 791.9 (-0.44%) |
| Protection (threat) | 1151.9 to 1151.9 (0) / 1152.2 to 1152.2 (0) | unchanged (no Seal of Righteousness) |
| Holy | unchanged (no Seal of Righteousness) | unchanged |

Split: S2 alone is -1.70% (alliance) and -1.66% (horde) at raid 60. S3 wins back +1.25 to +1.29%.

Bare bands, Retribution, alliance:
- 60: 307.6 to 298.0 (-3.1%); horde 60 (S2 alone, then S3): 284.5 to 296.5.
- 20 (S1): 46.0/45.9 to 46.2/46.1 (+0.4%).
- 58 and 59: see S3.

## Files to regenerate

- addon-data (Retribution rotation lines: Seal of Righteousness now shows at rank 7 from 58, and lines 2 and 8 now appear on the card; the published card shows rank 8 until then).
- sim/request ladder golden testdata/ladder/paladin-retribution.golden.md (TestRotationLadder).
- Retribution stat weights and dps-per-point (spell power and Intellect drop at 58-60 if S3 is not taken; with S3 they hold), tiers, coach starters if they quote the rotation.

## Measure at refresh

- Retribution bands 50 and 60, both presets and factions: BiS picks with the new Seal of Righteousness coefficients (spell damage items lose a little at 58-60 next to rank 7's 0.1).

## For other lanes

- Nothing shared touched. No commit moves DPS outside paladin.

## For the guide lane

- retribution.md :89, :91 and :103 and the Retribution card notes (lines 2, 5 and 8) shipped with S3 in S b193ac8d8.
- wow-player on the rendered page (dev build, captures in the lane scratchpad): first pass NO-SHIP with two MUSTs. "With any spell damage at all" was not true, and "keep rank 7 on your bar" ignored that training rank 8 replaces it on the bar. Both are fixed, and the tooltip is named. Second pass: **SHIP**.
  - Its open items: confirm the card reads "Rank 7 from 58, ..." after the nightly addon-data regen (the capture still shows the old card), and check in game that rank 7 stays in the spellbook after rank 8 is trained (as in Classic). If it does not, :89 needs a rewrite.
- protection.md:82: both the int/r18 wording and round 15's wording are true (P1). Take round 15's at the merge.

## Still open

- S2's mechanism: a level-60 log of Seal of Righteousness rank 8 against spell damage. If rank 8 does scale, revert S2, author rank 8 again and drop the rank 7 text.
- The Seal of Righteousness one-hand/two-hand spell damage split (0.1/0.11 in the engine, 0.0765/0.12 in Forever's tooltip).
- A2: Seal of Fury's aura coefficients below 60 (rank 4's 0.9).
- H1: Oath of Light (owner's check).
- X1 (carried): Exorcism and Holy Shield crit at 60.
- Warcraft Logs is still blocked (gcloud).

## Sources

- wago.tools SpellEffect, Spell, SpellLevels, SpellDescriptionVariables and SpellXDescriptionVariables, builds 1.60.1.70755 and 1.15.9.70576, fetched 2026-11-04. Forever spellconst for every build under data/builds/.
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/20292 and /20293 (2026-11-04). Era tooltip: https://nether.wowhead.com/classic/tooltip/spell/20293.
- Blizzard's beta notes, 16 October 2026 (quoted in sor.go and round 4 hybrids P1).
- Web search 2026-11-04: no Forever source gives a Seal of Righteousness coefficient. Classic sources disagree (warcraft.wiki.gg/wiki/Seal_of_Righteousness; us.forums.blizzard.com/en/wow/t/seal-of-righteousness/394206). SpellCoeffsForever (curseforge.com/wow/addons/spellcoeffsforever, updated 2026-10-11) derives coefficients from formulas and does not list Seals.
