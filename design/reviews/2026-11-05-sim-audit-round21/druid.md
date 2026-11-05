# Sim-accuracy round 21: Druid

Lane `druid`:
- Site: wt/r21-druid, from int/r21 93c66ebc8. Fork: lane/r21-druid, from int/r21 d5c708834.
- Client: 1.60.1.70790 (raw tables `scratchpad/b70790/70790`, wago CSV re-fetched 5 Nov), Blizzard's Forever patch notes (24 Sep to 4 Nov sections, read 5 Nov ~03:20 UTC), Wowhead Forever tooltips.
- Method: seeds first, then the round-20 guide leads the integrator added (guide accuracy round 20), then a fresh read of every druid line in Blizzard's notes against the engine. Measured with spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, frozen BiS (`$R/bis-frozen`), guide builds (`$R/rt`). Levels 10 to 19 have no published band, so a band-10 entry was made from the band-20 one with the gear stripped (`$R/druid/bis10`; the test builds the same thing in a temp dir).

**Summary: 6 confirmed, all fixed (B7 added after the merge of int/r21). Bear threat moves a lot (+45% raid 60: Bear Form's 115% from Blizzard's 3 November notes, Swipe's attack power share from the 1 October notes, and no single-target Swipe). Balance below 20 and the cat at 21 to 23 get a better card. Raid 60 is unchanged for Balance, cat and Restoration.**

**Logs.** None usable. api.foreversixty.gg/v1/reports/recent still lists only the 20 October report (checked 03:46 UTC 5 Nov); foreverlogs.gg public API answers "Forbidden" to a plain fetch, and the platform lane's 03:02 UTC check (browser headers) found the list empty and every id "Report not found"; WCL has no Forever zone (platform, 03:02). Every log-gated druid item stays OPEN below.

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Balance leveling, Eclipse Wrath on and off at 36, 40 and 50** (guide r19). With the leveling build's single Eclipse point the Wrath line wins: +0.9% at 40, +0.7% at 45, +1.0% at 50, +1.0% at 60 bare, both seeds, both factions (e.g. horde seed 1 at 50: 217.2 against 215.1; at 60 348.1 against 344.6). At 36 and 39 the two are identical: below 40 the mana line already casts Wrath nearly every global (Starfire 0.8 to 1.6 casts a fight), so the Eclipse line changes nothing. **Decision: keep the `auraIsKnown(408248)` gate; the one-point loop is worth casting**, so the note no longer limits it to the raid build, and balance.md's leveling section says when to start it (40, with Moonkin Form and Innervate). | `w/eclipse.tsv` (lane dir) | Aligned (note and prose updated) | - | S b2ddafb03, S bf29c62a7 |
| S2 | **Balance card had nothing to cast below 20 while mana was ahead.** Starfire trains at 20; below it the only Wrath outside the fight's end was the mana-pacing line (cast while mana is behind the fight). A level 14 moonkin with mana to spare and Moonfire ticking stood idle: geared for mana, it cast 2.9 Wraths in a 60 s fight.<br>**What a druid does:** Wrath is the filler until Starfire.<br>**Fix:** the mana line also fires while Starfire is unknown. Bare, no gear: +1.0% at 12, +0.2 to +0.5% at 11 and 14 to 19, both seeds and factions; level 10 -0.7% horde, 0.0% alliance (Moonfire rank 1 is the better mana there; one level, accepted). Test `TestBalanceBelowTwentyCastsWrathWhileManaIsAhead` (red 2.9 Wraths, green 35). | `w/low10.tsv` | CONFIRMED-BUG | old (guides-handoff, guide r19 sim finding; the mana line came in 0143f30ab, no round looked below 20) | S b2ddafb03, E 512030156 |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Starfire against Blizzard's 3 Nov "+40% base": rank 6 471 + 3.2 a level from spell level 58 (SpellLevels 9876, MaxLevel 64), rank 7 533 + 3.4 (the 70755 push). Wowhead's Forever tooltips show 451-529 and 510-597: they render each rank at its MaxLevel (rank 6 at 64: 471 + 6 x 3.2 = 490; Wrath 9912 97, 8905 73.4 the same way), not at 60. The engine's min(level, MaxLevel) growth is right. | SpellEffect, SpellLevels 70790; nether.wowhead.com tooltips 9876, 25298, 9875, 9912, 8905 | Aligned | - | - |
| L2 | Moonkin Form crit on every table (3 Nov "again an increase to all types"), Wrath ranks (8 Oct), Eclipse, Improved Wrath: as round 20 L3/L4, unchanged in 70790. | Blizzard notes; round 20 | Aligned | - | - |

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Builder switched at Ferocious Bite (24), not at Shredding Attacks.** Behind the target, bare band 20, leveling order: Shred beats Claw by +1.0% at 21 (48 energy with two points), +7.5% at 22 and +7.8% at 23; Claw wins by 4% at 20 (Shred 54). Both seeds, both factions (horde seed 1: 21 51.85 against 51.34; 22 55.34 against 51.48; 20 49.25 against 51.32).<br>**What was wrong:** a cat at 21 to 23 was told to Claw though Shred already does more for its energy; the card's `expected_idle` reason dated from the 155% Shred.<br>**Fix:** Shred `currentLevel >= 21`, Claw `< 21`; notes, the expected_idle reason and feral.md's leveling line say 21. Test `TestFeralBuilderIsShredFromTwentyOne` (red at 21 and 23, green). | `w/fer.tsv` | CONFIRMED-BUG | old (guides-handoff, guide r20 DRUID-R20-2 sim finding; the gate predates the 200% Shred of 3 Nov) | S bf29c62a7, E 92094e766 |
| C2 | **Rend and Tear on any bleeding target** (guide r20 question). The talent reads "on Bleeding targets" (1223246: one aura 4 dummy, no mask, the server decides); the engine counts the druid's own Lacerate, Rip, Rake and Pounce unless `AssumeBleedActive`. Decision: any bleed counts (the client text, and the WotLK talent of the same name counted other players' Deep Wounds and Rupture). Effect on the published fights is the opener only (the cat's first Rip, the bear's first Lacerate, a few seconds a fight), and only in a raid that has another bleeder; the site's spec options do not know the preset, so the flag cannot be set for raid only without plumbing. | talents/druid.json; SpellEffect 1211040 | Decided; engine gap OPEN-small (platform: preset-aware `assume_bleed_active` for feral; the bear options have no such field) | - | - |
| C3 | Shred 200%, Claw 140%, Ravage 435%, Rip/Rake/Bite/Pounce +25% (3 Nov), Ferocious Bite rank 1 at 24, Predatory Discipline's new text, Shred at 20 (16 Oct): all in the engine as round 18 to 20 found them. | Blizzard notes; round 20 C4 | Aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Bear Form threat bonus was 50%; Forever's is 115%.** Blizzard's notes of 3 November: "Bear Form and Dire Bear Form: Threat bonus increased from 50% to 115%." The cached 70790 SpellEffect row 700189 (21178 aura 10) still reads 50, and no hotfix row touches it, but the same note's "Defensive Stance: Threat bonus increased from 30% to 80%" is also absent from the cached row (686475 reads 30) while Wowhead's Forever tooltips for 71 and 7376 render 80% (and 107% with Defiance): the note's threat bonuses are server data the cached table predates. Righteous Fury's 33% from the same note is in the table (1367626).<br>**What was wrong:** the bear generated 1.5x threat instead of 2.15x.<br>**Fix:** `bearFormThreatMultiplier` 2.15. Test `TestBearFormThreatIsOneHundredFifteenPercentOver`. | Blizzard thread 2360696 (3 Nov section, read 5 Nov); wago SpellEffect 70790; nether.wowhead.com/forever/tooltip/spell/71 and /7376 | CONFIRMED-BUG | new-data (3 Nov server change; no round compared the notes' threat lines with the table) | E 77674506e |
| B2 | **Swipe had no attack power term.** 1 October notes: "Fixed a bug causing Swipe to not scale with Attack Power. It will now correctly gain 3% of the Druid's attack power." The client rows (779 to 9908) carry BonusCoefficientFromAP 0; the engine added nothing. Each target's hit is now the roll plus 3% of attack power (`SwipeHitDamage`). Test `TestSwipeGainsThreePercentOfAttackPower`. | Blizzard notes (1 Oct); SpellEffect 70790 | CONFIRMED-BUG (Blizzard note, as the round-20 Serpent Sting precedent) | old (in the notes since 1 Oct; rounds 6 to 20 read Swipe's flat amount only) | E 77674506e |
| B3 | **Single-target Swipe at 40 rage cost threat.** No Swipe on one target: +0.8% threat at raid 60 (1742.0 to 1756.3 alliance seed 1, 1741.8 to 1757.3 seed 50001; horde +0.7%), +1.6% on the round-start engine, +1% at bare 30 to 50, a tie at 20. Rage thresholds 25/30/50/60/70/80 sit between (higher is better). The rage does more in Maul, Lacerate and Primal Bite.<br>**Fix:** the Swipe line needs `numberTargets >= 2` as well as 40 rage; card note and feral-bear.md priority item 6. Test `TestBearSwipesOnlyOnAPack` (red 6.1 Swipes, green 0). | `w/` tank runs (lane dir) | CONFIRMED-BUG | old (the line predates round 6; no round tested Swipe out of the single-target card) | S 96b829d1f, E 221241f5a |
| B4 | **Threat constants.** Maul and Swipe 1.75 (Classic), Lacerate 2.0 ("a high amount", client text), Primal Bite 2.0: the 1 Oct note "Primal Bite: Threat generated by this ability has been roughly doubled" supports 2.0 for Primal Bite (one source). Maul, Swipe and Lacerate have no Forever statement. For the platform lane: this evidence. | Blizzard notes; threat.go | Primal Bite Aligned (one source); Maul, Swipe, Lacerate OPEN (a threat capture) | - | - |
| B7 | **Primal Bite against Maul** (guide accuracy round 21). Engine: Primal Bite 2.0x threat, 20 rage less Ferocity (15), 6 s cooldown, on the global; Maul 1.75x, 15 less Ferocity (10), on the next swing, and the swing it replaces earns no rage (`core/rage.go`: only a pure auto attack pays). Card order: Primal Bite first (row 3), Maul at 15 (row 7). Measured, 10,000 iterations, seeds 1 and 50001, both factions: Maul first ties Primal Bite first everywhere (within 0.16%); pooling for Primal Bite (Maul at 35 rage or while it is 1.5 s+ off cooldown) is +0.1 to +0.3% at bare 40 to 60 but -0.7% at 20 and -0.2% at raid 60; **Maul at 30 wins +0.2 to +0.8% at every level from 40 to 53** (40 302.3 to 304.2, 43 350.6 to 353.3, 50 498.1 to 499.5, 53 506.0 to 507.8; horde the same), ties at 54 to 55 and loses 0.3 to 0.5% below 40, from 56 and at raid 60.<br>**What was wrong:** from 40, when Primal Bite arrives, to 53 a bear is short of rage, and the 15-rage Maul left Primal Bite waiting on rage (28.4 of a possible 30 Bites a fight at 43).<br>**Fix:** the Maul line needs 30 rage from 40 to 53 (`currentLevel`), 15 otherwise; card note and feral-bear.md priority item 5. Test `TestBearMaulKeepsRageForPrimalBiteFromFortyToFiftyThree` (red 28.35 Bites, green 29.45). Primal Bite's 2.0 stays (B4). | `w/pb10k.tsv`, `w/maul30-levels.tsv`, `w/pb-after.tsv` | CONFIRMED-BUG (small, leveling only) | old (guides-handoff; the 15 predates Primal Bite's rank ladder; no round swept the Maul threshold by level) | S fbde7096e, E 01aba8c70 |
| B5 | Feral Instinct 3/3 in the bear build now acts only on packs (Swipe) in the published single-target fight. Kept: it is the pack talent the page's Adds paragraph relies on, and nothing else in reach does more for a tank on trash; not a point that does nothing for the spec. | build `0/5123033202310212551/055` | Decided (keep) | - | - |
| B6 | Bear crit rage (+75% 1 Oct, core's 2.0x normalized crit since), Faerie Fire threat in form (16 Oct), Demoralizing Roar threat (8 Oct, test-pinned): in the engine. | Blizzard notes; core/rage.go; demoralizing_roar_threat_test.go | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Blizzard's druid healing lines against the engine: Rejuvenation crits (24 Sep; every HoT tick rolls crit, `healing.go:103`), Wild Growth 3 s cast and Naturalist's 0.1 s a rank on it (3 Nov), Swiftmend needs your own HoT and consumes nothing (3 Nov), Bloom on a natural expiry (3 Nov), the high-rank Rejuvenation overwrite fix (16 Oct, nothing to model). Round 20's rank rows were judged level by level then; nothing in 70790 moved them (druid files identical to 70755, guide r19 D4). | Blizzard notes; engine healing.go, swiftmend.go, bloom.go, wild_growth.go | Aligned | - | - |

## Before/after

Raid 60 (`$R/before.tsv` / `$R/sb-before` against this branch, guide builds, frozen BiS, 10,000 iterations):

| Spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| Balance | 652.8 (0.0%) | 653.5 (0.0%) | 646.9 (0.0%) | 647.2 (0.0%) |
| Feral cat | 886.1 (0.0%) | 885.6 (0.0%) | 870.6 (0.0%) | 870.5 (0.0%) |
| Feral bear (threat/s) | 1209.1 to 1756.3 (+45.3%) | 1208.8 to 1757.3 (+45.4%) | 1206.5 to 1750.6 (+45.1%) | 1205.1 to 1750.7 (+45.3%) |
| Restoration | 939.5 (0.0%) | 939.5 (0.0%) | 935.3 (0.0%) | 935.6 (0.0%) |

Bear split at seed 1 alliance: B1 + B2 1209.1 to 1742.0 (+44.1%, of which the threat bonus is x1.433), B3 +0.8% on top. Bear own damage 453 a second, unchanged to 0.1%.

Below 60 (seed 1, horde; both seeds and factions agree): Balance 11 +0.2%, 12 5.04 to 5.09 (+1.0%), 14 +0.3%, 16 +0.4%, 18 +0.5%, 19 +0.6%, 10 -0.7% (alliance 0.0%); cat 21 +1.0%, 22 51.5 to 55.3 (+7.5%), 23 +7.8%, 20 and 24+ 0.0%; bear bare 30 to 50 about +1% from B3 (3,000 iterations) on top of the x1.43 threat; B7 bare 40 to 53 +0.2 to +0.8% (40 302.3 to 304.2 alliance seed 1), 20, 60 and raid 60 unchanged.

No commit on this lane moves any spec outside the druid class (fork changes are sim/druid and ui druid APLs only).

## Files to regenerate

- `data/builds/1.60.1.70790/addon-data.json` and `addon/ForeverSixty/Data.lua`: card rows and first-sentence notes for Balance (mana Wrath, Eclipse Wrath, Starfire), Feral (Shred/Claw gates and notes, Rip note) and bear (Swipe gate and note). The rendered cards show the old text until then.
- Ladder goldens `sim/request/testdata/ladder/druid-balance.golden.md` (level 10 rung now casts more Wrath), `druid-feral-bear.golden.md` (Swipe AP, no single-target Swipe), and `druid-feral.golden.md` if a rung sits at 21 to 23 (rungs are 10, 20, 30: likely unchanged).
- After the pin: bear tank weights, stat-weights/dps-per-point and BiS for druid-feral-bear (threat x1.43 and Swipe's AP share shift weights toward attack power a little); coach starters if they quote bear threat.
- Fork goldens regenerated here: `sim/druid/tank/TestFeralTank.results` (twice: B1/B2, then B3). Fork `sim/conformance` goldens for mage, warlock, paladin, priest and shaman were already stale at d5c708834 (not this lane).
- Copied by hand: `sim/request/apl/druid-{balance,feral,feral-bear}.apl.json` and the fork's `forever_{balance,feral,feral_bear}.apl.json` (apl_check green).

## Measure at refresh

- Bear weights and BiS on the new threat model (B1 to B3).
- Feral bands 20 on re-ranked gear: the 21 switch (C1) holds by 1%, the first to re-check.

## For other lanes

- **Warrior (and integrator), from B1's evidence:** Blizzard's 3 Nov note "Defensive Stance: Threat bonus increased from 30% to 80%" is live: Wowhead's Forever tooltips for Defensive Stance (71) read "Increases all threat generated by 80%" and Defensive Stance Passive (7376) "[Defiance: 107% / 80%]", while the cached 70790 row 686475 and the engine (`sim/warrior/stances.go:164`, 0.3) still have 30. The same note says Defiance 7/13/20%; the tooltip's 107% is 1.8 x 1.15, so Defiance's rank 3 still reads 15% there. Not changed by this lane.
- **Platform:** bear threat constants (B4): Primal Bite's "roughly doubled" (1 Oct) supports 2.0; Maul, Swipe and Lacerate have no Forever statement. C2: a preset-aware `assume_bleed_active` if the raid is taken to carry another bleeder.
- **Integrator, merge:** `data/curated/apl/druid-balance.json` and its request copy conflict with main (guides r19) on the Eclipse and mana Wrath notes; this branch already carries guides r19's and r20's wording for every other druid note it touches, so take this branch's side. feral.md line 119's Shred sentence and the Feral Shred/Claw notes supersede guides r20's "from 22".

## For the guide lane

Done here, prose test green (18), wow-player SHIP on the rendered pages (1440 crops, `$R/druid/cap/`): Balance and bear SHIP with SHOULDs applied (From 4 Moonfire, "0.17 seconds", Eclipse 35 against 40, "Until 40, Starfire is for the opener", Swipe "on a pack of three"); Feral SHIP, one SHOULD applied ("Ravage and Pounce both need you behind"). Bear B7 (Maul at 30 rage at levels 40 to 53): HOLD on level-versus-rage wording, fixed ("At levels 40 to 53 ... only at 30 rage", card note with units), then SHIP on the recapture. Open SHOULDs for the guide lane: feral.md Levels 20 to 60 is one nine-sentence paragraph (split builder / openers / finishers); feral-bear.md Weaknesses "a bear the boss is not hitting starves" could read "runs out of rage"; the Claw card names 54 where the page gives 48.

## Still open

- Eclipse on a missed Wrath; Ferocious Bite's attack power share; the unarmed paw: live logs (Balance 40+, cat 40+ with Bites, a bare-handed cat).
- Maul, Swipe and Lacerate threat multipliers (B4); Rend and Tear on others' bleeds in the raid fight (C2).
- Pending owner, not reopened: as the brief.
