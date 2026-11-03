# Sim audit round 12: hybrids (2026-11-03)

This report covers Retribution, Protection paladin (threat), Enhancement, Feral cat and Feral bear (threat), at every band.

**Setup:**
- Client 1.60.1.70753. `python -m pipeline push-check --build 1.60.1.70753` on 2026-11-03 returns `moved=false`: wago lists no 70753 push, so the shipped tables are the client. Raw tables come from round 11's 70707 and 70753 copies (`scratchpad/r11/c70753/builds/*/raw`, with `shipped/` and `wago-push/`).
- Wowhead Forever tooltips (nether.wowhead.com/forever/tooltip/spell/<id>) were read 2026-11-03.
- Engine: pin 2606bc05e plus `lane/r12-hybrids`. Site: `wt/r12-hybrids` on c6bdc855c.
- "Before" is `$R/sb-before` from `$R/site-before/sim` and `$R/before.tsv`. All numbers use 10,000 iterations at seeds 1 and 50001 unless a row says otherwise.
- Nothing nightly-owned was regenerated or committed.

Fix-sha column: "E" is a fork commit on `lane/r12-hybrids`, "S" a site commit on `wt/r12-hybrids`.

**In short: 6 CONFIRMED-BUGs, all fixed.**
- Poised for the cat (new-data).
- The band 30 Protection build left out Holy Shield (old).
- Enrage's armor penalty in Dire Bear Form (old).
- Primal Bite's level 60 rank (new-data).
- Hammer of Wrath's travel time (old; found by the warlock lane).
- The cat's Claw note (landing).

Also adopted under the talent rule: a better cat leveling build. It goes to the guides lane, because the build line is in their file.

## Assigned items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Holy Shield rank 4 (1239570)** matches the client on every number.<br>- Client rows: level 60 (SpellLevels 148343); taught by SkillLineAbility 53829 (skill 267, AcquireMethod 0); 240 mana (SpellPower 310878); 10 s category cooldown and 1.5 s global (SpellCooldowns 94637); 10 s duration; 4 charges (SpellAuraOptions 242217); +30% block (effect 0, aura 51); 221 damage at 0.08 of spell power (effect 1, aura 43); the text's 20% extra threat; family mask [0,64,0,0], the same as ranks 1 to 3, so Benediction reaches it.<br>- Engine (`sim/paladin/holy_shield.go:41`): 240, 221, 0.08, 30, 4, 10 s, x1.2 threat, Benediction multiplier.<br>- Ranks by band: 40 casts 20927, 50 casts 20928, 60 casts 1239570 (spellranks maps the authored rank 4 down).<br>- The whole landing move is rank 4. The same APL with rank 3 gives 1021.9 threat/s against 1100.7 (raid 60, horde, seed 1, 3,000 iterations), which is the landing's +7.7%. The proc's threat goes from 204.0 to 283.9 a second.<br>- Wowhead's tooltip for 1239570 shows no mana cost. That is Wowhead's data, not the client's; client tables win.<br>- Crit table: still open (round 11 O3). | client rows above; `bd/prot-*.txt`; `rot/prot-r3.json` | Aligned | - | - |
| A2 | **Retribution's Seal opener** (landing 7dd072e12). The raid build has Unsealed Wrath, so it opens with Seal of Command rank 5 (20920, 189 mana after Benediction). The band builds lack Unsealed Wrath and open with Seal of Righteousness at the band's rank (20288 at 20, 20290 at 40, 20293 at 60). | `bd/ret-*.txt` | Aligned | - | - |
| A3 | **Poised for the cat.** 70753 teaches every druid Poised at 32 (SkillLineAbility 59889, passive 1328829): "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10." Its buff 1328827 (SpellEffect 1368040) is aura 107, SpellMod 14 (cost), -10, with 1 charge and no duration (DurationIndex 21). Its family mask [8622080, 536870912] covers Claw, Rake, Shred, Ravage, Pounce, Rip, Ferocious Bite and Cower; it does not reach Maul, Swipe, Lacerate or Shifting Power. The engine did not have it.<br>**What changed:** a fight opens out of combat at full Energy, so the opener (Ravage) costs 50 instead of 60. Only the first ability gets the discount.<br>**Shared code:** the physical lane's `sim/core/poised.go` (`core.RegisterPoised`, written for the rogue's 1328826) is copied byte for byte (sha1 9342ca9e); `sim/druid/poised.go` adds the druid's rows. | Test first: `TestPoisedTakesTenEnergyOffTheCatsFirstAbility`, `TestCatPoisedIsNotKnownBelowThirtyTwo`, `TestCatPoisedReachesEveryEnergyAbility` (sim/druid/feral). They did not build before the fix and pass after. | CONFIRMED-BUG | new-data (70753) | E 3a8651388 |
| A4 | **The cat leveling build (adoption rule).**<br>- The gap to the raid build is +4.1% at bare 60 (265.5 to 276.4, horde, seed 1) and +5.3% at raid 60, not 9%.<br>- The raid build cannot replace the leveling build: it drops Brutal Impact and Feral Charge, which the engine does not model.<br>- A bare-60 talent search (`-keep` Shifting Power, Leader of the Pack, Nature's Focus, Subtlety) found a winner that keeps every unmodeled guide talent: `FS1:1.60.1.70753:druid:night-elf:01/50212032123012123001/535033:`. Against the guide it is Heart of the Wild 3 instead of 5, Genesis 1 and Natural Shapeshifter 3.<br>- Bare 60: +1.09% / +1.17% horde, +1.27% / +1.12% alliance (seeds 1 / 50001).<br>- Band 50 (the walk gives Genesis 1, Nature's Focus 5, Heart of the Wild 3, Naturalist 1): +1.6% / +1.6% horde, +1.5% / +1.6% alliance.<br>- Bands 20 to 40 do not change (same Feral tree).<br>- It keeps every unmodeled talent, keeps the Feral tree, and drops nothing the rotation casts. **Adopted**; the build line belongs to the guides lane (For other lanes).<br>- Whether bare 60 uses the right build: by design, bare 60 spends the leveling build (`leveling-bis/raid_build.go`), so the better leveling build is the fix. | `ts-out/druid-feral.md`; `-talents` runs | Adopted (rule), not a sim bug | old (e43958ce5, 2026-10-31) | guides lane |
| A5 | **Claw against Shred.**<br>- Shred is learnable at 20 (5221, SkillLineAbility; `shred.go:13`). The sim assumes you stand behind the target, the ranker's position for every melee spec, so Shred from behind is the sim's right default.<br>- Per energy, Claw leads until Shredding Attacks is full: at 20, Claw costs 40 (Ferocity 5) and Shred 54 (1 point).<br>- The walk fills Shredding Attacks at 22, so the APL's switch at 24 (gated on Ferocious Bite) is two levels late. No band falls between 21 and 23, so no band number moves.<br>- Shred spending Clearcasting at band 20 measured +0.7% (45.0 to 45.3, both seeds). That is under the 1% rotation bar, so it is held.<br>- The APL's note said rank 1 Shred "hits for less than Claw". It hits harder: 65 against 50 on average at band 20. The note is now right.<br>- The guide's own lines are right: Claw while solo, because a mob faces you; Claw in a raid when you cannot get behind. | `rot/cat-shred20.json`, `rot/cat-clawonly.json`; breakdowns | CONFIRMED-BUG (note only) | landing c0f5432b2 | S 7ea0ba759 |
| A6 | **Omen of Clarity is baseline.** 16864 is taught by SkillLineAbility 9402 (skill 574, AcquireMethod 0) at level 20 (SpellLevels 102825). Wowhead says "Requires level 20". It is in no talent tree. The engine gives it to every druid from 20 (`talents.go:369`), in cat, bear and Balance, with Moonkin Form doubling its chance; no build can take it. This agrees with the trainables lane if its ruling is "baseline". | client rows; `omen_test.go` | Aligned | - | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Hammer of Wrath had no travel time.**<br>- What the client says: every rank (24275, 24274, 24239) has SpellMisc Speed 35. It is a thrown hammer, like Fireball (24) and Wrath (20).<br>- What was wrong: the engine dealt the damage at the end of the cast.<br>- What changed: the damage rolls at cast end and lands after the distance divided by 35.<br>- At melee range the flight is about 0.14 s, so this is noise: raid 60 +0.01 / -0.01%. Relayed from the warlock lane. | Test first: `TestHammerOfWrathFliesAtTheClientsSpeed` (failed at speed 0, passes at 35) | CONFIRMED-BUG | old | E 16e9403b2 |
| P2 | **Costs at raid 60 and bare 40, per cast, against the client:**<br>- Judgement: 6% of base mana, less Benediction (81.7).<br>- Seal of Righteousness: 200 to 180. Seal of Command: 210 to 189.<br>- Consecration: 565 to 282.4 (Holy Conduit plus Benediction, added).<br>- Holy Strike: 20 to 18.<br>- Hammer of Wrath: 425 to 255 (Holy Conduit; Benediction's mask leaves it out).<br>- At bare 40: Seal of Righteousness 120 to 108, Consecration 26573 135 to 121.5. | `bd/*-after.txt`, cost sweep | Aligned | - | - |
| P3 | **Learnability at bands 20 to 60:** every id the rotation casts is a SkillLineAbility row, a talent, or a Seal or Judgement effect spell. | sweep | Aligned | - | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **The band 30 build left out Holy Shield.**<br>- What the client says: Holy Shield (20925, learned at 30) sits on the 21-point row. A level 30 paladin's 21 points are the 20 above that row plus one.<br>- What was wrong: the top-row-down walk spent the 21st point on Improved Righteous Fury, so the band 30 tank never cast its main ability, which the rotation casts.<br>- What changed: `ladder-first.json` now reserves Holy Shield, then Shieldsworn (so band 40 keeps the capstone the spine already gave it), each as soon as it fits. Only band 30 changes.<br>- Bare 30 tank fight: threat 65.1 to 90.1 (horde), 60.5 to 84.8 (alliance), +38 to +40% at both seeds. Damage taken rises 1.9% without the Improved Righteous Fury point.<br>- The published number moves at the refresh. | Test first: `TestProtectionTakesHolyShieldAtThirty` (sim/leveling) failed on `...-15302102511000000-...` and passes; `TestCuratedLadderFirstIsSound` covers the new row | CONFIRMED-BUG | old (the walk, 78efc0d3e; no Protection row since round 10's mechanism) | S 9db09c660 |
| T2 | **Seal of Fury and Judgement of Fury ranks by band** (1311654/1311655 at 20 through 20418/20414 at 60) are the client's effect and rank spells. Holy Shield ranks by band: see A1. | sweep | Aligned | - | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Fire Nova on Clearcasting** (landing b2e4a3bc5).<br>- Fire Nova 408341 to 408345 is taught at 12/22/32/42/52 (SkillLineAbility 8734 to 8738): 375 mana at rank 5, 10 s cooldown, a dummy cast.<br>- Its damage spell 408428 is 239 ± 11% + 1.3 a level, capped at level 57, with a 0.143 coefficient.<br>- The engine matches (`fire_nova.go:25`).<br>- Clearcasting 16246 (-100% cost) covers Lightning Bolt, Chain Lightning, the three shocks, Fire Nova and Lava Burst (mask -1743781885 / 4096). The engine's set (shaman damage spells with a spell proc mask) is the same set; Searing Totem's attacks are in neither. | client rows; `talents.go:114` | Aligned | - | - |
| N2 | **The landing's build swap.** The raid build trades Earth's Grasp 2 for Improved Ghost Wolf 2 (both unmodeled utility, no DPS change). The leveling build drops Earth's Grasp and Guardian Totems for Improved Ghost Wolf and Shamanistic Focus (+3.9% bare 60 per the commit). This is a guide-lane decision like the Shadow ruling, not a sim bug. The published entries still carry the old builds until the refresh. | b2e4a3bc5 | Aligned (note) | - | - |
| N3 | **Costs:**<br>- Stormstrike: 8% of base mana (121.6 at 60).<br>- Earth Shock: 450 x 0.55 (Shamanistic Focus). It averages 242.4 because Clearcasting sometimes lands on it.<br>- Totems at their client costs.<br>- Bare 40: Earth Shock 240; Lightning Bolt cast free on Maelstrom.<br>Learnability: only totem attacks and effect spells fall outside SkillLineAbility. | cost sweep | Aligned | - | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Poised (A3), the leveling build (A4), Claw (A5), Omen (A6). | above | see above | | |
| F2 | **Costs:** Shred 42 (Shredding Attacks 3), averaging 37.8 with Clearcasting; Rip 30; Ferocious Bite 35; Ravage 60, now 50 under Poised; Shifting Power 55% of base mana (684). | cost sweep | Aligned | - | - |
| F3 | **Natural Shapeshifter.** It does reach Shifting Power (`shifting_power.go:45`). The published raid entry still carries the pre-c0f5432b2 build (Restoration `505003`, no Natural Shapeshifter), so the raid number moves at the refresh. | bis/druid-feral.json | Aligned (Measure at refresh) | - | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Enrage took 27% of the armor in Dire Bear Form.**<br>- What the client says: Enrage (5229) "reduces base armor by 27% in Bear Form and 16% in Dire Bear Form. The druid is considered in combat for the duration." Wowhead Forever's tooltip reads the same.<br>- What was wrong: the armor effect itself is a dummy with no amount, and the engine applied Bear Form's 27% at every level. Every bear from 40, in Dire Bear Form, lost 11 points of armor too many while Enrage was up. The landing opens every fight with Enrage.<br>- What changed: the penalty is now the form's (`bearFormTier.EnrageArmorReductionPercent`).<br>- Raid 60: damage taken 593.3 to 592.4 (seed 1) and 593.6 to 592.4 (seed 50001); threat -0.02 to -0.08%.<br>- The "in combat" line confirms the 1 s pre-pull: rage does not decay. | Test first: `TestEnrageArmorPenaltyFollowsTheForm` (sim/druid) did not build before the fix and passes; TestFeralTank golden regenerated (DTPS 609.8 to 609.3, 678.4 to 677.5) | CONFIRMED-BUG | old (86dc10877, 2026-10-07) | E 21445df6f |
| G2 | **Primal Bite's level 60 rank was one no druid learns.**<br>- What the client says: SkillLineAbility teaches 407995 at 40 (row 53146), 1238069 at 50 (53141) and 1238070 at 60 (53142, rank 3, superseding 1238069). Rank 4 (1238073) states the same 77 at 60, but push 113283 removed its learn row (53143, status 2) in 70707, and 70753 ships without it.<br>- What was wrong: the engine and the rotation named 1238073, so the addon and coach pointed at a spell missing from the player's spellbook.<br>- What changed: the engine registers 1238070 at 60, and the APL casts rank 3. Damage, cost and cooldown are identical, so no number moves (curated rotation: 1010.7 threat/s against 1011.5 before). | Test first: `TestPrimalBiteLiveRanks` and the level 60 kit test now want 1238070; they failed before the fix | CONFIRMED-BUG | new-data (70707 push 113283) | E 57068d0fd, S be6afacce |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Learnability sweep** of every cast id at bands 20/30/40/50/60 and raid 60 for the five specs, against the 70753 layered SkillLineAbility (AcquireMethod 0) and talent trees. Every miss is a proc, aura, buff, item, totem-attack or effect spell, except G2. | sweep script output | Aligned (except G2) | - | - |
| X2 | `raid-keeps.json` still keeps Nature's Focus and Subtlety for druid-feral, but the guide's raid build dropped them on the commit's "do nothing for a cat" (c0f5432b2, +1.6%). The record and the build disagree; for the coordinator. | raid-keeps.json; c0f5432b2 | OPEN (record) | - | - |

## Before/after

Raid 60, 10,000 iterations. After: `lane/r12-hybrids` 57068d0fd with `wt/r12-hybrids` 7ea0ba759. DPS, threat/s for the tanks. The bear's after uses the curated rotation (`-rotation data/curated/apl/druid-feral-bear.json`), as the apl-sync regen will give it; its before is `$R/before.tsv`.

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-retribution | +0.01% | -0.01% | 0.00% | 0.00% | 690.1 -> 690.2 |
| paladin-protection (threat) | 0.00% | 0.00% | 0.00% | 0.00% | 1101.0 -> 1101.0 |
| shaman-enhancement | 0.00% | 0.00% | 0.00% | 0.00% | 586.0 -> 586.0 |
| druid-feral | 0.00% | 0.00% | 0.00% | -0.02% | 660.0 -> 660.0 |
| druid-feral-bear (threat) | -0.08% | -0.02% | -0.02% | -0.02% | 1011.5 -> 1010.7 |

Bare bands, horde, seeds 1 / 50001:
- **Cat (Poised):** band 40 +0.35 / +0.44%, band 50 +0.52 / +0.52%, band 60 +0.53 / +0.56% (115.0 -> 115.4, 152.8 -> 153.6, 265.5 -> 266.9). Bands 20 and 30 are below Poised's level. At raid 60 the 10 Energy is lost by the fight's end (identical Shred counts).
- **Bear (Enrage):** bare 40 -0.06%, bare 60 +0.05%: noise.
- **Protection band 30:** threat +38 to +40%, at the refresh (T1).
- **Retribution (Hammer of Wrath):** bare 50 and 60, 0.00%.

## Files to regenerate

- **Required with the pin:** `sim/request/apl/druid-feral-bear.apl.json` (apl-sync). Without it, the embedded rotation names 1238073, which the engine no longer registers, and the bear drops Primal Bite: 858.9 threat/s against 1010.7.
- `sim/request/apl/druid-feral.apl.json` (apl-sync; notes only).
- addon-data, for the bear's rotation id and both druid APLs' notes.
- Site `TestRotationLadder` goldens:
  - Protection band 30 (Holy Shield);
  - the cat from 40 (Poised);
  - the bear from 40 (Enrage);
  - Primal Bite's id in the druid "unused" lists (`druid-balance.golden.md:197`, `druid-restoration.golden.md:214`);
  - Retribution's and Protection's Hammer of Wrath timing (likely no visible change).
- spellranks.json and trainables output, if the trainables lane drops 1238073 (For other lanes).
- `addon-builds.json`, after the guides lane changes the cat's leveling build.
- Fork goldens regenerated here: `sim/druid/tank/TestFeralTank.results` only.

## Measure at refresh

- Protection band 30, both factions: Holy Shield in the band build, so threat and gear picks change.
- The cat at bands 50 and 60 bare on the new leveling build (A4), and raid 60 on c0f5432b2's raid build (Natural Shapeshifter 3).
- Enhancement's bands on b2e4a3bc5's builds: the published entries still carry the old ones.

## For other lanes

- **Guides (web/src/content/guides/druid/feral.md):**
  - Frontmatter: `build: 'FS1:1.60.1.70753:druid:night-elf:01/50212032123012123001/535033:'`.
  - Prose for the leveling-build paragraph: "Genesis, 3 points of Heart of the Wild and Natural Shapeshifter's third point beat the last 2 points of Heart of the Wild at 50 and 60, about 1.5% at 50 and 1% at 60; Feral Charge and Brutal Impact stay for questing."
  - The Claw lines (86 and 112) are right as written. If you want the reason: "In a group from behind, Shred takes over once Shredding Attacks is full, around 22."
  - `_prose.test.ts` must pass.
- **Trainables:** `trainables/druid.json` lists Primal Bite 1238073 as rank 4 at 60 through the "no learn row" extras rule (`trainables.py:248`). Its row was removed by push 113283 (SkillLineAbility record 53143, status 2), so it is not learnable; rank 3 (1238070) is the level 60 rank. Omen of Clarity: the engine and every druid build treat it as baseline at 20 (SkillLineAbility 9402, AcquireMethod 0). If you rule otherwise, tell this lane.
- **Platform (coach):** `data/curated/coach/druid-feral-bear.json:11` names 1238073; it should be 1238070 (verified line `spellconst/druid.json#1238070`).
- **Physical:** `sim/core/poised.go` in this lane is your file copied byte for byte (sha1 9342ca9eb0dfc147012dc07bbd0005de3e56f74c). At the merge, take yours if it changed. The druid side is `sim/druid/poised.go`, using `core.RegisterPoised` and `core.ClassMaskTable`.
- **Warlock / coordinator:** Hammer of Wrath travel is done (P1).
- **Coordinator:** X2, the raid-keeps record against the cat's raid build.

## Still open

1. Holy Shield crit (round 11 O3): no 40+ Forever log yet.
2. The band 20 Clearcasting Shred (+0.7%) and the Claw-to-Shred switch at 22 rather than 24: under the 1% rotation bar, held.
3. Round 11's G4 (rage from damage taken uses the boss's level) and the round 10 log list.
4. The cat's raid 60 Poised Energy: real, but lost by the end of the fight in the sim (identical Shred counts). A launch log of the opener would show whether real cats keep it.
