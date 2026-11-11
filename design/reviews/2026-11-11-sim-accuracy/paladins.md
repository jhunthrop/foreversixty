# Sim accuracy 2026-11-11: Retribution and Holy paladin

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-paladins`. Site from origin/main cfe2a6f67, fork from master e586233d7 (the pin at the start; rebased on fe282fdfb, the priest lane's pin, before the push). Client on the site: 1.60.1.70790; live client 1.60.1.70888. Scope: Retribution and Holy, end to end: engine against the client and the hotfix push, engine against real fights, builds, rotations and rank rows under the adoption rule. Protection was audited today by the tanks lane and is not re-audited here; one shared spell fix moves its numbers, and those are reported.

**Totals: 1 confirmed engine bug, old, fixed. Holy Strike's weapon percent (25% to 50% by rank) was multiplying the flat Holy bonus and the spell damage term along with the weapon damage; the live logs show it scales the weapon damage only. Retribution +5.0 to +6.0% bare at every band and +2.0% at raid 60; Protection threat +1.2 to +6.6% bare and +3.2% at raid 60; Holy unchanged. 0 new data: Blizzard's 6, 9 and 10 Nov hotfixes and the 70888 client change nothing a paladin card casts. 24 checks against the first paladin logs aligned (Oath of Light's amount among them, open since round 20), 1 open (Sanctified Judgement's refund is 60% of the seal's cost for six logged paladins and 66% for five). Engine pin e586233d7 to 3bdfab5b6 (on fe282fdfb).**

## Method

1. **Blizzard since 6 Nov.** "WoW: Forever Hotfixes" (t/2403040, posts 4 to 6: 6, 9 and 10 Nov) and "Class Tuning Incoming" (t/2404092). Paladin lines: Cooldown Manager fixes for Seals and Redoubt, and "Consecration could stop dealing increased damage to enemies" (9 Nov; the first-four-targets bonus on a pack, which the single-target fights here never lose). Nothing a card casts moved.
2. **Client and push.** The 70888 tables against 70790 were diffed today by the hotfixes lane (`hotfixes-other-classes.md`): the two paladin rows are Consecration's despawn fix (1280353) and an equipped-weapon row on Unsealed Wrath. Not repeated. The engine was read against `data/builds/1.60.1.70790/spellconst/paladin.json`, `talents/paladin.json` and `spelltips/paladin.json` (the push layered), every spell, talent and aura the Retribution and Holy cards and builds touch.
3. **Real fights.**
   - Our API's 8 Nov advanced-logging report (gxrgvz4yzsi2, 104 fights, open world on a PvP realm): every event with the advanced fields (`base_amount`, the hit before armor and before the crit multiplier; the source's attack power and spell damage on swings; `adv_power_cost` on casts). 26 paladins appear; one, a level 30 to 35 Retribution paladin with Verigan's Fist (3.2 s two-hander, COMBATANT_INFO), fought for two hours: 313 white swings, 80 Seal of Command procs, 78 Holy Strikes, 93 Judgements of Command, 128 Judgements, 1,307 Consecration ticks.
   - foreverlogs.gg: all 24 public reports (2815 to 2850, 5 to 11 Nov). Every encounter whose participants hold a paladin was read in full (timeline, participants, combatant info where the uploader logged it): 17 reports, 153 encounters, 111 paladin rows, 29,908 paladin events. These logs carry no armor-free base amount, no power and no level; a paladin's level is read from the ranks he casts, his weapon speed from his swing timer outside Seal of the Crusader, and his weapon from combatant info where present.
4. **Sims** on the Windows box through `tools/remote-test.sh` (`--pin-engine` on the old pin, `--engine-sha` on the lane's fork branch; `nice -n 15`, GOMAXPROCS=5, one sim at a time): spec-breakdown before and after at raid 60 and the bare bands 20 to 60, both factions, seeds 1 and 50001, 10,000 iterations; the Holy Strike rank rows re-judged level by level; the fork's paladin and core suites (`-tags with_db`) in a writable clone; the three paladin ladder goldens.

## Confirmed and fixed

| # | finding | evidence | origin | fix |
|---|---|---|---|---|
| H1 | **Holy Strike's weapon percent cut the flat Holy bonus and the spell damage.** `holy_strike.go` applied effect 31's percent (25, 29, 32, 36, 39, 43, 46, 50 by rank) as the spell's DamageMultiplier over the whole base, flat plus normalized weapon damage, and so over the 0.429 spell damage term too. The client's own text reads "causes 50% weapon damage plus an additional 81 to 105 as Holy damage", and the logs agree: the 8 Nov paladin's 78 non-critical rank 4 hits (36%, 24 to 25 flat at his level) average 106 against his own pre-armor white swings of 198 and Seal of Command procs of 145 (70% of weapon damage, so 207 of weapon damage; the two agree on his weapon). 0.36 x 202 + 25 is 98, and 107 once his own Judgement of the Crusader (+58 Holy, through the coefficient) is counted on the quarter of hits that had it; the old reading gives 82. On foreverlogs.gg, two Seal of Command paladins land Holy Strike at 0.73 and 0.75 of a Seal of Command proc (78 and 14 hits): the fixed formula predicts 0.72 and 0.74, the old one 0.60. A third, rank 3 at 20 to 27 with a 3.2 s two-hander, lands 93 against white hits of 95 to 125: 0.32 x 125 + 19 + 25 (his Judgement of the Crusader) is 84 before his Holy multipliers, against 54 for the old reading. Now `baseDamage = flat roll + percent x normalized weapon damage`, DamageMultiplier 1 (1.2 with Sacred Arbiter), BonusCoefficient 0.429 whole. Red first: `TestHolyStrikeWeaponPercentLeavesTheFlatAndSpellDamageWhole` (rank 8, Arcanite Reaper, 500 spell damage, a same-level target: 287.9 against a wanted 460.7; green at 460). | 8 Nov report fights 9 to 49; foreverlogs 2842 (Bool), 2818 (Dawgie) | old (the file's first version, 2026-10) | E 187a29e7a (fix), b6ee909dc (test), 3bdfab5b6 (test on a same-level target, Protection regression results) |

Before and after (spec-breakdown, 10,000 iterations, seeds 1 / 50001, both factions, the old pin e586233d7 against 87ba79a52, the same fix before the rebase):

| spec | cell | before (alliance s1) | after | change (alliance s1 / s50001, horde s1 / s50001) |
|---|---|---|---|---|
| Retribution | raid 60 | 792.6 | 808.8 | +2.04 / +2.04, +2.07 / +2.05% |
| Retribution | bare 20 | 46.7 | 49.4 | +5.8 / +5.8, +6.0 / +5.7% |
| Retribution | bare 30 | 85.1 | 89.7 | +5.4 / +5.5, +5.4 / +5.4% |
| Retribution | bare 40 | 133.5 | 141.3 | +5.8 / +5.9, +5.3 / +5.3% |
| Retribution | bare 50 | 195.4 | 207.1 | +6.0 / +6.0, +5.5 / +5.5% |
| Retribution | bare 60 | 344.3 | 361.7 | +5.1 / +5.0, +5.0 / +5.1% |
| Protection (threat /s) | raid 60 | 1125.8 | 1162.4 | +3.25 / +3.24, +3.22 / +3.22% |
| Protection | bare 20 | 48.5 | 51.7 | +6.6 / +6.6, +5.6 / +5.6% |
| Protection | bare 30 | 116.9 | 120.2 | +2.8 / +2.8, +2.7 / +2.6% |
| Protection | bare 40 | 261.1 | 264.3 | +1.3 / +1.3, +1.2 / +1.2% |
| Protection | bare 50 | 385.3 | 400.2 | +3.9 / +3.9, +3.7 / +3.7% |
| Protection | bare 60 | 496.2 | 523.1 | +5.4 / +5.4, +5.3 / +5.3% |

Holy casts no Holy Strike; its raid 60 heal fight is unchanged (739.68 effective HPS at seed 1 on both pins). The fork's Protection regression suite moves with it (Phase 1 average DPS 591 to 607, threat 1370 to 1422; the spell power weight 0.304 to 0.346), recorded in `TestProtection.results`; the Retribution suite's rotation casts no Holy Strike and is unchanged.

## Logs against the engine

Aligned unless marked. "Our" is the 8 Nov advanced report; "fl" is foreverlogs.gg.

| # | check | evidence | verdict |
|---|---|---|---|
| L1 | **Seal of Command procs 7 a minute.** 99 procs on 261 landed white swings under the Seal with a 3.2 s weapon, 37.9%; 7 x 3.2 / 60 is 37.3%. (fl: 36 procs on 138 swings over four paladins, 26%, a thin and window-dependent sample.) | our | aligned |
| L2 | **Seal of Command proc is 70% of weapon damage.** Proc base 145.0 against white base 198.2, 0.732; 0.7 plus the 0.2 x 7 spell damage and a share of hits under Judgement of the Crusader. | our | aligned |
| L3 | **Judgement of Command is half its roll unless the target is stunned.** 93 non-critical hits in two clusters, 85 to 108 and 171 to 205, the second under the paladin's own Hammer of Justice. | our | aligned |
| L4 | **Seals last 30 s.** 30.0 to 30.1 s from Seal of Command's application to its fall when not recast. | our | aligned |
| L5 | **Judgement of the Crusader lasts 40 s and the paladin's own white swings refresh it.** 99 refreshes, every one at a landed swing (plain, or with a Seal of Command proc on the same swing); Holy Strike alone refreshed none (no Sacred Arbiter on that paladin). | our | aligned |
| L6 | **Judgement of the Crusader's bonus reaches Holy hits through the coefficient, not flat.** Consecration rank 2 ticks are 10 to 12 on the first targets (3 + 7, plus 0.042 x 58 is 2.4, plus his Holy multipliers) and 2 to 3 beyond the fourth; a flat +58 a tick would read 68. Seal of Righteousness procs on foreverlogs carry 4 to 7 more than the formula while the paladin's own Crusader Judgement is up (0.0765 or 0.12 x 58 or 92). | our, fl 2825, 2845, 2848 | aligned (core `SchoolBonusDamageTaken x BonusCoefficient`) |
| L7 | **Consecration: 8 ticks a second apart, 3 + 7 a tick at rank 2 to the first four targets and 3 to the rest.** 1,307 ticks: 3 (506), 10 to 12 (975), 2 (15), 8 to 9 (40); the 11s and 12s are the 10 under Two-Handed Weapon Specialization and Vengeance. | our | aligned |
| L8 | **Consecration ticks crit.** 229 critical ticks of 1,536. | our | aligned (owner's DoT rule) |
| L9 | **Exorcism never crits.** 0 of 22 hits (1 ours, 21 fl), on top of round 2's 0 of 641. | our, fl | aligned |
| L10 | **Vengeance stacks to 5.** Applied doses at 2, 3, 4 and 5 stacks (21, 15, 10, 7 ours; 17, 13, 11, 8, 12 at 1 to 5 fl). | our, fl | aligned (70755 push) |
| L11 | **Seal of Righteousness proc damage** follows 0.01 x (amount + per level) x 0.85 (one-hander) or 1.2 (two-hander) x speed plus 0.0765 or 0.12 of spell damage: twelve paladins at ranks 2 to 6 land inside the rank's own-level to max-level range once their spell damage and their Crusader Judgement are counted (2.0 s, 2.3 s, 2.4 s one-handers; 3.1 to 3.7 s two-handers; a hasted 2.31 s reading is a 3.23 s two-hander under Seal of the Crusader). | fl 2815 to 2848 | aligned |
| L12 | **Blessing of Wisdom rank 2 returns 18 mana every 5 s** (305 of 312 ticks). | our | aligned |
| L13 | **Sanctified Judgement refunds on every Judgement at rank 3** (128 of 128 for the 8 Nov paladin, 19 of 19, 6 of 6, 3 of 3 for three others; a rank 2 paladin on foreverlogs refunds 40%). | our, fl | aligned (chance) |
| L14 | **Sanctified Judgement's amount.** 60% of the judged Seal's base cost for six paladins (39 of 65, 66 of 110, 54 of 90, 84 of 140, 75 of 125, 39 of 65) and 66% for five (72.6 of 110 and 42.9 of 65 for the 8 Nov paladin, over 128 refunds; 92.6 of 140, 40 of 60 and 26.4 of 40, 43 of 65, 28.6 of 65 for a rank 2 paladin, 44% of 65). The talent's rank 3 row says 60%, the engine says 60%. The extra tenth is on both factions, on one-handers and two-handers, at level 20 and at level 40, and on rank 2 as well as rank 3; it is not Benediction (the costs paid are full), not Vengeance (constant through a fight), not the race. No buff or talent in the logs separates the two groups. | our, fl 2818, 2823, 2834, 2842, 2845 | OPEN (see Still open) |
| L15 | **Oath of Light adds 1.9167 x level to Flash of Light and 6.6667 x level to Holy Light on the sworn target, whole.** A Holy paladin at 30 to 33 (Divine Favor, Seal of Righteousness rank 4): Flash of Light rank 2 on the sworn tank 183.3 (24 hits, sd 4), on others 123.5 (42 hits, sd 3), the difference 59.8 against 59.4 at 31; Holy Light rank 4 on the tank at most 640 against 432 on others, 208 against 206.7. Round 20's open item: the amount is whole and takes no cast-time coefficient on a rank near the caster's level. The low-rank penalty itself still waits for a paladin casting a rank 17 levels below his own. | fl 2834 (Onyooms, 7 encounters) | aligned (round 20 closed for near-level ranks) |
| L16 | **Illumination.** 12 mana returns on 22 critical heals for that paladin; his rank is not in the log (3 of 5 would be 60%). | fl 2834 | not comparable |
| L17 | **Judgement of Wisdom and Light, Seal of Light and Wisdom** (the 50% and 15 a minute assumptions): no logged paladin runs them long enough to count (7 Seal of Light heals in all). | fl | no data |
| L18 | **Holy Shock, Light's Vigil, Unsealed Wrath's twist, Hammer of Wrath:** no paladin in the logs has reached them (the one Unsealed Wrath paladin at 39 to 43 shows "Echo of Command" and "Echo of Righteousness" stacks, the live client's name for the stored Seals, which the engine's aura does not carry; a name, not a number). | fl 2842 | no data |

Two reading notes for later lanes. The advanced report's attack power field on a paladin's swings runs well above what his level and gear allow (644 to 905 at level 30 to 35), so the weapon part of a formula is best read from his own Seal of Command procs or his swings, never from that field. The foreverlogs timeline gives no base amount, so a white hit there is after armor while a Holy hit is not; the Seal of Command proc, at a fixed 70% of weapon damage, is the armor-free yardstick.

## Engine against the client (1.60.1.70790, hotfixed)

Every number below matched the client row; H1 above is the one that did not.

| area | engine | client | verdict |
|---|---|---|---|
| Judgement | 6% of base mana, off the global cooldown, 10 s less 1 s a rank of Improved Judgement, never consumes the Seal | 20271 cost_pct 6, cooldown 10000, gcd 0; talent 105705 "1 sec" a rank; Deep Dive "no longer consumes Seals" | aligned |
| Seal of Righteousness | 8 ranks at 1 to 58, 20 to 200 mana, 30 s; proc 0.01 x (amount + per level to max level) x 0.85 or 1.2 x speed, 0.0765 or 0.12 of spell damage; Judgement 15 to 170 centre, 0.429 | 20154 to 20293 (amounts 108 to 1786, per level 18 to 47), the tooltip's `$1h=0.0765*$SPH`, `$2h=0.12*$SPH`; 20187 to 20286 | aligned (L11) |
| Seal of Command | talent-granted, 5 ranks at 20 to 60, 65 to 210 mana; 7 a minute with a 1 s cooldown, 70% weapon damage, 0.29 (0.2 effective); Judgement 97 to 356 centre halved unless stunned, 0.429, melee table without active defence | 105696; 20375 to 20920; 20424 effect 31 amount 70 sp 0.29; 20467 to 20966; tooltip "138 to 146 if the target is stunned or incapacitated" | aligned (L1 to L3) |
| Seal of the Crusader | 6 ranks, 31 to 306 attack power plus per level, +40% speed and swings at 1/1.4; Judgement 23/35/58/92/127/161 Holy taken for 40 s, refreshed by the paladin's melee | 21082 to 20308 (aura 99, aura 9 at 40), 21183 to 20303 (aura 14, 40000) | aligned (L5, L6) |
| Holy Strike | 8 ranks at 6 to 60, 5 to 20 mana, 10 s cooldown (6 with Shieldsworn); flat 12 to 93 with 25% width and per-level growth; 25% to 50% of normalized weapon damage; 0.429 | 679 to 10333: effect 121 (amount, variance 0.25, points per level), effect 31 (25 to 50), sp 0.429, category 10000 | fixed (H1) |
| Consecration | 5 ranks at 20 to 60, 135 to 565 mana, 8 s cooldown, 8 ticks of 1 s, 2 to 12 a tick plus 4 to 27 to the first four, 0.042 of spell damage on that share | 26573 to 20924 and 1280345 to 1280349 | aligned (L7) |
| Hammer of Wrath | 3 ranks at 44, 52, 60, 295 to 425 mana, 1 s cast less 0.2 s a rank of Champion of the Light, 6 s cooldown, 288 to 498 centre, 0.429, below 20% health, thrown at speed 35 | 24275, 24274, 24239 | aligned |
| Exorcism | 6 ranks at 20 to 60, 85 to 345 mana, 15 s cooldown, 79 to 502 centre, 0.429, Undead and Demons, never crits | 879 to 10314 | aligned (L9) |
| Swift Judgement | a second Judgement of the active Seal, no cost, no global, 60 s | 1310994 | aligned |
| Benediction, Improved Judgement, Conviction, Sanctified Judgement | 2% a rank on instants; 1 s a rank; 2% crit a rank on every table (push 113139); 33/66/100% to return 20/40/60% of the judged Seal's cost | 105706, 105705, 105703, 105701 | aligned (L14 open on the amount) |
| Two-Handed Weapon Specialization, Vengeance, Champion of the Light, Sacred Arbiter, Unsealed Wrath | 2/4/6% Physical and Holy with a two-hander, 7/14/20% less threat; 1/2/3% a stack to 5 for 30 s; 12% of Intellect a rank as spell damage and 0.2 s off Hammer of Wrath; Holy Strike +20% and Judgement refresh; 3 stored Seals for 15 s | 105697, 105693 (20050 stacks 5), 110882, 105700, 105692 | aligned (L10) |
| Improved Seals, Reverence, Holy Conduit, Divine Strength, Divine Intellect, Precision | 3% a rank on Seals and Judgements; 15% a rank of regeneration while casting; 20/40% off Consecration, Holy Wrath, Exorcism, Hammer of Wrath; 2% a rank Strength; 2% a rank Intellect; 1% a rank hit | 105334, 110871, 105704, 113630, 105332, 105638 | aligned |
| Holy Light, Flash of Light | 9 ranks at 1 to 60, 35 to 660 mana, 2.5 s, 0.714; 6 ranks at 20 to 58, 35 to 140 mana, 1.5 s, 0.429 | 635 to 25292, 19750 to 19943 | aligned |
| Holy Shock | 4 ranks at 30, 40, 48, 56, 160 to 325 mana, 10 s cooldown, heal 114 to 320 centre and damage 134 to 348, 0.429 | 1311606, 20473, 20929, 20930 and their heal and damage rows | aligned |
| Light's Vigil | 3 ranks at 40, 50, 60, 730 to 1340 mana, 1.5 s cast, 6 s cooldown, 30 s; the next Holy Shock on the target heals its party 324 to 704 centre at 0.143 and takes no cooldown | 1310911, 1311590, 1311595 and the heal rows | aligned |
| Oath of Light | 8% of base mana, 5 min, one target; 6.6667 / 1.9167 / 3.8333 a level on Holy Light, Flash of Light, Holy Shock, whole on a near-level rank | 1328548 (points per level 6.66667, 1.91667, 3.83333) | aligned (L15) |
| Spiritual Focus, Healing Light, Illumination, Divine Favor, Infusion of Light, Boundless Light, Holy Power | 14% a rank; 4% a rank; 20% a rank for 50% of base cost and 50% regeneration for 15 s; 100% crit, 2 min, 4% of base mana; 0.5 s a rank for 15 s; 15% a rank Holy Shock crit off the Oath target; 10% a rank crit bonus | 105335, 105333, 105329 (20272 aura 134 at 50), 105325 (20216 cost_pct 4), 110873, 113631, 105321 | aligned |
| Blessings of Might and Wisdom, Greater | 14 to 133 attack power; 12 to 40 mana every 5 s; the Greater ranks at 52/60 and 54/60 | 19740 to 25291, 19742 to 25290, 25782, 25916, 25894, 25918 | aligned (L12) |

## Builds, rotations and rank rows

- **Retribution raid build and twist.** Unchanged. The twist's four-fifths pace and the Consecration budget were measured in rounds 22 to 24 and no number this lane touched moves them: Holy Strike's gain is the same on every line. The raid build's three dead points (Vindication) remain the only legal way to Unsealed Wrath's 30 points (round 29).
- **Holy raid build and rotation.** Unchanged (round 29 H5: no live replacement for the seven Protection points).
- **Holy Strike rank rows.** The fix raises the flat bonus's share of every Holy Strike, so a lower rank loses more than before. Re-judged on the box after the fix, bare, `-level`, both published races, seeds 1 and 50001, 10,000 iterations (`$S/ranks-after.txt`):
  - Protection's one Holy Strike row, rank 1 at 40, loses on the fixed formula: the top rank (5) makes 4.65% more threat for a Dwarf (276.6 against 264.3 at both seeds) and 4.2% more for an Undead (281.7 / 281.9 against 270.4 / 270.3); rank 1 also loses 4.2 to 4.9% at 12 to 16 (both races) and 5.0 to 5.2% at 39 and 41. The row is removed (`data/curated/rank-choices.json` and the `sim/rankchoices` copy), the card note no longer sends a 40 tank to rank 1, and `TestProtectionHolyStrikeRankRows` and `TestRankChoicesRemeasured`'s Holy Strike case now want the top rank at 40. Retribution keeps no Holy Strike row and needs none: the top rank beats the rank below by 0.6% at 21, 2.2% at 29, 2.2% at 45 and 0.8% at bare 60 (Human, both seeds).
- **Leveling orders.** Not re-walked: Holy Strike is baseline, so no talent's value moves against another except through Sacred Arbiter (+20% of a now larger hit), which the Retribution order already takes at 25 as the first live point after Seal of Command.

## Gates

- Fork, on the box, in a writable clone at 3bdfab5b6: `go test -tags=with_db -p 1 ./sim/paladin/... ./sim/core/`: ok (paladin, holy, protection, retribution, core). The red test at b6ee909dc read 287.9 against 460.7.
- Site, on the box with `--engine-sha 3bdfab5b6`: the paladin ladder goldens regenerated (`FOREVER_UPDATE_GOLDEN=1 go test -run 'TestRotationLadder/paladin-(retribution|protection|holy)$' ./request/`; Retribution and Protection move at every rung, Protection's rung 40 once more when its rank 1 row went, 47.9 to 53.1; Holy is unchanged); `go test -skip TestRotationLadder ./request/ ./rankchoices/ ./leveling/` ok; `TestRotationLadder/paladin` ok; `npm test -- src/content/guides src/lib/bis` 974 passed (the prose, sections, builds and card tests).
- Python: `data/tests/test_apl.py`, `test_addonrotation.py`, `test_card_note_levels.py`: 112 passed; `test_every_allowlist_line_still_matches_a_note` is red on origin/main (druid, hunter and mage lines waiting on the addon-data regen, no paladin line) and stays so.
- Guide text: wow-player NO-SHIP on the first wording ("takes the whole of it" read as 100% of spell damage where the share is 0.429; one sentence carried three ideas), both MUSTs taken verbatim, then SHIP.

## Files

- Fork `lane/sim-paladins`, on fe282fdfb: b6ee909dc (test, red), 187a29e7a (fix), 3bdfab5b6 (same-level target, `TestProtection.results`). Master fast-forwarded to 3bdfab5b6.
- Site: `sim/paladin` untouched (the fix is the fork's); `sim/enginever/version.go` and `web/src/lib/sim/version.ts` pinned to 3bdfab5b6; `sim/request/testdata/ladder/paladin-{retribution,protection,holy}.golden.md`; `data/curated/apl/paladin-retribution.json` and `sim/request/apl/paladin-retribution.apl.json` (the Holy Strike note); `web/src/content/guides/paladin/retribution.md` (the stat paragraph); `data/curated/rank-choices.json` and `sim/rankchoices/rank-choices.json` (the Protection Holy Strike row); `data/curated/apl/paladin-protection.json` and `sim/request/apl/paladin-protection.apl.json` (its note); `sim/request/protection_holy_strike_rank_test.go` and `sim/request/rank_choices_remeasured_test.go`.
- No regens: no BiS rank, data run or simdb rebuild. Lane scripts `tools/lane-pal-batch.sh` and `tools/lane-pal-ranks.sh` were scratch and are not pushed.

## Measure at refresh

- MR52: Retribution and Protection BiS, weights and stat order on the fixed Holy Strike (spell power and Intellect are worth more: the fork's Protection stat weights move spell power 0.304 to 0.346), every band and raid 60; the Retribution guide's statPriority keeps Spell power sixth until the weights say otherwise.
- MR53: Protection band 40 on the top Holy Strike rank (the published entry's casts carry rank 1), and the addon data's Protection Holy Strike note and row at the nightly regen.

## Still open

- **L14, Sanctified Judgement's 66%.** Five of eleven logged paladins get 66% of the Seal's cost back, six get 60%; both at rank 3 (100% refunds) and the one rank 2 paladin's 44% against 40%. Modelling 60% keeps the client's row and half the logs. Whatever adds the tenth (a +10% to mana returned the logs do not show, or a server rule on the "mana cost of the judged seal") is worth asking the owner to watch for on his own paladin: one Judgement under Seal of Command rank 5 with no buffs shows 126 or 139.
- **Oath of Light's low-rank penalty** (round 20): still no paladin casting a rank 17 levels below his own.
- **Holy Shock, Light's Vigil, the twist, Hammer of Wrath** in logs: no paladin above 43 yet.
- **The ranker's raid 60 Retribution score** (850.7 in the published 10 Nov entry) against spec-breakdown's 792.6 on the same gear and preset at e586233d7: a 7% gap between the two tools on one spec, not chased here; for the platform lane to read (the ranker's set measure runs 3,000 iterations at seed 7 with a 0.2 duration variation and a 0.25 execute ratio).
