# Sim-accuracy round 22: Druid

Lane `druid`:
- Site: wt/r22-druid from int/r22 d5cc4d6b4 (later merged with int/r22 3337ca488). Fork: lane/r22-druid from int/r22 170a59c39 (later merged with int/r22, which carries the platform lane's Rend and Tear option 69c6325fb).
- Client: 1.60.1.70790 (raw tables `scratchpad/b70790/70790`, plus wago SpellEffect re-fetched 5 Nov ~07:40 UTC); Blizzard's Forever notes thread 2360696 (re-checked 07:30 UTC: 8 posts, last 4 Nov 19:35, nothing new); Wowhead Forever tooltips.
- Measured with spec-breakdown: 10,000 iterations, seeds 1 and 50001, frozen BiS (`$R/bis-frozen`), guide builds (`$R/rt`). For levels 10 to 19, a band-10 entry is made from band 20 with its gear stripped (`$R/druid/bis10`).

**Summary: 5 confirmed, all fixed.**
- **One fix moves every spec.** Leader of the Pack and Moonkin Aura gave 6% crit together. They are exclusive since Blizzard's 3 November notes, so together they now give 3%. Raid 60 loses 1.2% to 3.7% for every spec. The fix is in core, flagged to the integrator during the round.
- Feral: Rip and Ferocious Bite now take 25% more of attack power (3 Nov notes); +1.9% at raid 60 and +1.4 to +1.7% on bare bands.
- Feral below 20: no Cat Form until the druid learns it.
- Caster Faerie Fire: no longer resets the swing.
- Bloom: aligned and documented.
- Balance +2.35% lead: not adopted (adoption rule).

**Logs (last checked 07:27 UTC 5 Nov).** None:
- api.foreversixty.gg/v1/reports/recent lists only the 20 October report.
- foreverlogs.gg's API answers "API access not allowed".
- No Warcraft Logs v1 host (classic, fresh, sod, vanilla, www) lists a Forever zone.

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Bloom and BloomSpellCoeff.** Bloom's heal is 100% of Spirit (50% when dispelled), with no healing power.<br>- The heal 1328622 is SpellEffect effect 10, base 1, EffectBonusCoefficient 0. The talent 1328621 is two aura-4 dummies, 100 and 50: the Spirit shares the server reads.<br>- Its SpellClassMask is zero in every word, so family-masked talents (Gift of Nature, Improved Rejuvenation) do not reach it.<br>- SpellMisc Attributes_3 0x40010000 lacks IGNORE_CASTER_MODIFIERS 0x20000000, so unmasked healing multipliers apply. Attributes_2 lacks CANT_CRIT, so it crits.<br>- The engine already did all of this: `bloom.go` sets no BonusCoefficient and no class mask. `BloomSpellCoeff` 0.2143 (1.5/3.5/2, the hybrid-instant convention) is the generator filling a zero column, and nothing reads it.<br>- The seed calls it "Lifebloom's bloom". Lifebloom 409824 is a SoD rune, not in trainables/druid.json, and its `LifebloomSpellCoeff` is also unread.<br>- The new test `TestBloomTakesNoHealingPowerOrGiftOfNature` fails if Bloom reads the 0.2143 or gets a Rejuvenation mask (721.9 against 292 Spirit). The comment in `bloom.go` now cites the rows and closes the old "unconfirmed: healing modifiers". | wago SpellEffect / SpellMisc 70790; nether.wowhead.com/forever/tooltip/spell/1328621 and /1328622 | Aligned (documented, pinned) | - | E 29c9fa4f1 |
| S2 | **Balance +2.35% lead (MythicSim).** The lead swaps Genesis 2 to 5 and adds Heart of the Wild 5, paid by dropping Nature's Focus 5 and Subtlety 3 (`5532220115501351-0-055000`).<br>- Verified: horde 646.9 to 660.9 at seed 1 and 647.2 to 661.2 at seed 50001 (+2.16%).<br>- **Not adopted, by the adoption rule.** Both dropped talents do something real for a moonkin in a raid. Nature's Focus is 70% pushback protection on Arcane and Nature spells. Subtlety is -30% threat on Nature and Arcane spells. Both are in the 70790 client text and on raid-keeps.json. The engine models neither, and the rule keeps every unmodeled talent.<br>- A small keep-respecting talent search found nothing better than the guide: raid, horde, keeps applied, 3000-iteration finalists. Genesis credits +3.4 a point. Every Balance point it could come from credits +6.7 or more, and Heart of the Wild +0.5.<br>- This closes round 14 B3 (owner hold) as a lane decision: the build stays. | spec-breakdown runs above; `$R/druid/ts-out/druid-balance-raid.md`; talents/druid.json | Decided: keep the guide build | - | - |

## Cross-cutting (core, moves every spec)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Leader of the Pack and Moonkin Aura stacked to 6% crit. Forever makes them exclusive.**<br>**What was wrong:** the raid preset carries both. `core/buffs.go:247-255` added the single all-table crit stat (`stats.Crit`) once for each, so every spec in the raid fight had 6% crit from the two druid auras. A moonkin's own form aura also stacked with a raid Leader of the Pack.<br>**What the client and Blizzard say:**<br>- Moonkin Aura 24907 (SpellEffect 701335) and Leader of the Pack 24932 (702727) are both effect 65 (raid area aura), aura 290 (crit chance of every kind), 3. The same type, so they can only be exclusive with each other. No hotfix row touches either.<br>- Blizzard's notes, 3 November section: "Moonkin Form: The critical strike bonus from this talent is again an increase to all types of critical strikes and is again exclusive with Leader of the Pack." Also: "Leader of the Pack: ... is again exclusive with Moonkin Aura."<br>- Wowhead Forever tooltips for 24907 and 24932 both read "Increases Critical Strike chance by 3%." With both at 3%, whichever one wins, the raid gets 3%.<br>- The guide pages already said this (guide accuracy rounds 1 and 2).<br>**Measured before the fix** (3000 iterations): Fury's Bloodthirst crit went from 29.2% to 26.3%, and Fire's Scorch from 33.8% to 31.1%.<br>**Fix:** `core.RaidHasDruidCritAura` gives 3% once when either aura reaches the level. The druid's form reads the same check, so a moonkin beside a raid Leader of the Pack adds nothing. The preset keeps both druids: a 40-man raid has a feral and a moonkin.<br>Tests: `TestLeaderOfThePackAndMoonkinAuraAreExclusive` (red 0.06, green 0.03), `TestMoonkinFormAddsNoCritBesideLeaderOfThePack` (red 0.03, green 0), and core's `TestLeaderOfThePackAndMoonkinAuraReachEveryTable` (now 3 for either or both). | wago SpellEffect 70790; Blizzard thread 2360696; nether.wowhead.com/forever/tooltip/spell/24907, /24932 | CONFIRMED-BUG | new-data (3 Nov server change). The raid-buff code is from 2338b7f5a (round 8); rounds 5 and 9 to 10 checked the aura type, not exclusivity. | E 5d23a8747 |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Caster Faerie Fire reset the swing timer.** The 1 October notes say "Faerie Fire no longer resets your swing timer when used." The caster spell carried `SpellFlagResetAttackSwing`; the feral-form spell did not. No shipped card casts caster Faerie Fire, so no number moves. Test `TestFaerieFireDoesNotResetTheSwingTimer` (red, then green). | Blizzard notes, 1 Oct | CONFIRMED-BUG (no number moves) | old (rounds 14 and 21 read the feral-form spell) | E 655bee16c |
| L2 | Re-audit of Balance:<br>- **Spells:** Wrath, Starfire (+40% base), Moonfire, Insect Swarm, Eclipse, Nature's Grace, Moonkin Form (Omen x2), Innervate. As round 20 L3 and round 21 L1/L2; no druid client row changed in 70790.<br>- **Raid build, point by point:** every point is modeled, except the two keeps (S2).<br>- **BiS:** Frozen Heart of the Mountain's use is Frost and Shadow only (`item_effects.go:3741`). Balance takes the trinket for its 0.9% hit, so the use rightly adds nothing. Rune of the Guard Captain is 0.7% hit and 14 attack power. Idol of the Moon, Ironfeather 2-piece and Sageclaw are as round 14 B2. | engine; nether.wowhead.com item tooltips 249469, 19120, 20070 | Aligned | - | - |

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Rip and Ferocious Bite kept Era's attack power shares.**<br>**What the notes say:** Blizzard, 3 Nov: "Rip: Base points and attack power scaling increased by about 25%." The same line exists for Ferocious Bite, Rake and Pounce.<br>**What the client shows:** the base half is in the client push. Rip rank 6 went 15 to 19 a tick and 25.5 to 31.9 a point; Bite rank 6 went 145 to 181 a point; Rake rank 4 went 61 to 76. Rounds 18 to 21 checked those.<br>**What was wrong:** the attack power half is server-side. The engine kept Era's 1% of attack power a point a tick on Rip (6% a point over 12 s, four points at most) and 3% a point on Bite.<br>**Fix:** 1.25% and 3.75%. Rip now takes 30% of attack power at 4 to 5 points, and Bite 18.75% at 5.<br>**Rake and Pounce stay with no attack power term.** Their live tooltips say nothing of attack power, while Rip's and Bite's do ("and by your Attack Power"), and Era gave them none. Neither is cast in the published fights.<br>Test `TestRipAndBiteAttackPowerSharesAreRaisedAQuarter` (red at Era's values, green). | Blizzard thread 2360696 (3 Nov); wago SpellEffect 1079 to 9896, 22829; tooltips 9896, 22829, 9904, 9827 | CONFIRMED-BUG (Blizzard note; the base half of the same line is in the client, as the Serpent Sting precedent) | old (rounds 18 to 21 checked only the base half of the 3 Nov line; round 21 C3) | E c7a5fa846 |
| C2 | **A feral druid below 20 started every fight in Cat Form.**<br>**What was wrong:** at the pull it had cat stats and the paw weapon, and at 18 and 19 it cast the feral-form Faerie Fire in a form it has not learned. The first Moonfire then cancelled the form. Cat Form 768 is trained at 20 (trainables/druid.json).<br>**Fix:** the spell registers from 20 (round 21 law: gate by registering from the trainer level). Below 20 the druid starts in caster form and swings its own weapon.<br>Tests: `TestCatFormIsLearnedAtTwenty` (red at 19) and `TestFeralBelowTwentySwingsItsOwnWeapon` (red: 65.8-damage paw on an unarmed druid).<br>The feral APL's engineering note called this a "deliberately unfixed" inaccuracy; the note is updated. **Caster Faerie Fire** (platform's second question) also registers at every level, but no card casts it (tag 0), so it is left as is. | trainables/druid.json; `$R/druid/bis10` runs | CONFIRMED-BUG (sub-20 only) | old (rotation-accuracy program 2026-09-28 saw it and left it; platform r22 found the missing gate) | E f749164f6, S 47aca2b53 |
| C3 | **Rake upkeep, re-measured.** Rake is now stronger (+25% base, Genesis) and raid fights assume a bleed. A Rake row before Shred is still -4.9% at raid 60 (horde: 813.6 against 855.7 at 3000 iterations). MythicSim's Rake hits 244 a tick against our 42 + Genesis. That gap is their attack power model, which the live tooltip does not support (C1). Rake stays out. | `$R/druid/feral-rake.json` | Aligned | - | - |
| C4 | **BiS:** Idol of the Dream (tested), Molten Heart of the Mountain (55 attack power, `item_effects.go:3705`), The Unstoppable Force (2% crit and a stun proc), Manual Crowd Pummeler and the Rune of the Guard Captain: as their tooltips. | tooltips 249470, 19323 | Aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Thorns (8 Oct):** Vanilla base values and a 6% spell power share. ThornsRanks 3/6/9/12/15/18 match the client (467 to 9910, aura 15). The spell power share is not modeled, and a bear's spell power is about 0, so it is negligible. | wago SpellEffect; `core/buff_ranks.go:353` | Aligned (negligible) | - | - |
| B2 | **Ranks:** Primal Bite has 3 live ranks (40/50/60; the client's rank 4 1238073 is not trainable, `primal_bite.go:45`). Lacerate is 25/34/42/50/58. Maul, Swipe and Demoralizing Roar match trainables/druid.json. | trainables/druid.json | Aligned | - | - |
| B3 | **Threat constants:** Maul 1.75, Swipe 1.75, Lacerate 2.0. Primal Bite 2.0 is backed by "roughly doubled". No new source, and no log. | as round 21 B4 | OPEN (a threat capture) | - | - |
| B4 | **BiS:** Arena Grand Master (absorb use), Talisman of Arathor (absorb), the PvP insignias: survival uses with no threat or damage. | tooltips 19024, 20071, 17904 | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Bloom: S1. Idol of Swiftness takes 3 s off Swiftmend (1306488, aura 107 -3000 ms, tooltip). Dormant Heart of the Mountain's Eternal Power triggers on Healing Touch (`client_procs.go`). Wild Growth's 3 s cast and front-loaded ticks, Swiftmend not consuming, Rejuvenation and Tranquility crits: as round 21 R1. | tooltips 279250, 249473 | Aligned | - | - |

## Blizzard's druid lines (24 Sep to 4 Nov), read one by one

- **Fixed this round:**
  - Faerie Fire swing reset (1 Oct): L1.
  - Rip and Bite attack power (3 Nov): C1.
  - Moonkin Form and Leader of the Pack exclusive (3 Nov): X1.
- **Aligned, already in the engine:**
  - 24 Sep: Rejuvenation, Tranquility and Wild Growth crit; Wrath +50%.
  - Thorns: B1.
  - Primal Bite rename and Blood Frenzy name.
  - 1 Oct: bear crit rage (+75%, then +100% on 8 Oct); Swipe 3% attack power (round 21 B2); Tiger's Fury and King of the Jungle removed (in no APL or tree); Primal Bite threat x2; Shifting Power.
  - 8 Oct: mana potions in forms; Cat Form health; Demoralizing Roar threat; Wrath ranks; Natural Instinct; Shifting Power not at full energy; Thick Hide.
  - 16 Oct: the high-rank Rejuvenation fix; Omen of Clarity not from a free Faerie Fire; Faerie Fire threat in Bear Form; Shred at 20; Insect Swarm hit stacking; Furor energy.
  - 3 Nov: Lacerate a talent; Ferocious Bite rank 1 at 24; Bear Form 115% (round 21 B1); Poised; Shred, Claw and Ravage percentages; Starfire +40%; Genesis; Heart of the Wild, Furor, Naturalist, Reflection and Subtlety moves; Bloom; Swiftmend; Wild Growth 3 s; Rend and Tear 3%; Primal Bite 3 ranks.
- **Not modeled, no effect here:**
  - All forms immune to Disarm.
  - Feral Charge on allies.
  - Feral Charge (Cat) range and the shapeshift-override visuals.
  - Quest fixes.

## Before/after

Raid 60, guide builds, frozen BiS, 10,000 iterations. "Before" is `$R/before.tsv` (r21 landed). "After" is this branch merged with int/r22, which includes the platform lane's Rend and Tear option (feral about +0.85%, bear about +0.5%, by the platform lane's figure). Seed 1 is absolute; both seeds give the relative change.

| Spec | alliance s1 | horde s1 | relative, both seeds and factions |
|---|---|---|---|
| Balance | 652.8 to 635.7 | 646.9 to 628.7 | -2.6 to -2.8% (X1) |
| Feral cat | 886.1 to 888.0 | 870.6 to 871.2 | +0.07 to +0.21%. Split: C1 is +1.85/+1.87% (871.9 to 888.0, 855.2 to 871.2); X1 with platform's Rend and Tear is -1.6 to -1.8% |
| Feral bear (threat/s) | 1756.3 to 1715.1 | 1750.6 to 1710.5 | -2.3 to -2.5% (X1) |
| Restoration (effective HPS) | 939.5 to 927.3 | 935.3 to 923.9 | -1.2 to -1.3% (X1) |

**Every other spec moves through X1 alone** (no other lane change reaches raid 60 at this int/r22), in % at both seeds and factions:

| Class | Spec and change |
|---|---|
| Hunter | BM -2.80 to -2.88; MM -2.94 to -3.01; Survival -2.92 to -2.95 |
| Mage | Arcane -2.43 to -2.51; Fire -3.14 to -3.27; Frost -2.07 to -2.12 |
| Paladin | Holy -3.61 to -3.73; Protection -2.28 to -2.29; Retribution -2.43 to -2.45 |
| Priest | Discipline -1.76 to -1.90; Holy -1.20 to -1.32; Shadow -2.42 to -2.44 |
| Rogue | Assassination -2.45 to -2.63; Combat -2.23 to -2.43; Subtlety -2.38 to -2.57 |
| Shaman | Elemental -2.30 to -2.34; Enhancement -2.74 to -2.89; Restoration -2.18 to -2.23 |
| Warlock | Affliction -2.01 to -2.02; Demonology -2.52 to -2.60; Destruction -2.14 to -2.21 |
| Warrior | Arms -2.76 to -2.93; Fury -3.35 to -3.51; Protection -2.34 to -2.37 |

The full table is in `$R/druid/after.tsv`.

**Below 60, feral, horde** (both seeds within 0.1):
- Bare 30: 94.7 to 96.0. Bare 40: 141.8 to 144.0. Bare 50: 201.2 to 204.5. Bare 60: 342.5 to 348.3 (+1.4 to +1.7%, C1).
- Levels 14, 18 and 19 on the stripped band-10 entry: 6.0 to 5.8, 6.7 to 6.4, 6.9 to 6.6 (-3 to -4%, C2: no paw at the pull).
- Band 20 is unchanged (51.3).
- Bare bands carry no raid auras, so X1 does not reach them.

## Files to regenerate

- **Fork goldens outside druid, moved by X1 (FullBuffs carries both auras):** hunter TestP1Hunter; mage TestP1Mage; paladin TestProtection and TestRetribution; priest TestP1Shadow; rogue TestCombatDaggers and TestCombatSinisterStrike; shaman TestElemental, TestEnhancement and TestWardenShaman; warlock TestWarlockDSRuin and TestWarlockSMRuin; warrior TestP1DPSWarrior and TestP1TankWarrior. Also any `sim/core/testdata/conformance` golden that sims with both auras. Regenerated here: `sim/druid/tank/TestFeralTank.results` (crit 32 to 29).
- **Ladder goldens:** `sim/request/testdata/ladder/druid-feral.golden.md` (rung 10: C2; every rung: C1). The bear, Balance and Restoration ladders are bare and unchanged, unless a rung uses the raid preset.
- **After the pin:** stat weights and dps-per-point for every spec at raid 60 (3% less crit shifts the crit weights); BiS for every spec (the queued CI re-rank); feral weights at every band (more attack power value through Rip and Bite). Coach starters, if they quote raid numbers.
- No card, APL action or rank row changed, so addon-data and Data.lua need nothing from this lane.

## Measure at refresh

- Every spec's raid-60 weights and picks on 3% crit (X1). The crit-heavy picks (Fury, Fire, the hunters) first.
- Feral at the refresh: Rip at 5 points against Bite with Rip still up (APL rows 6 and 7). This lane did not re-tune them for the stronger finishers.

## For other lanes

- **Integrator:** X1 is a core change that moves all 28 specs (table above). I sent the evidence to you during the round. The druid APLs are unchanged, so no APL copies need syncing.
- **Platform:** caster Faerie Fire (9907, tag 0) still registers at every level as rank 4. No card casts it. Whether to register it by rank from 18 is a platform tidy-up. `BloomSpellCoeff` and `LifebloomSpellCoeff` are unread convention fills. The generator could emit the client's 0 for triggered heals, so no later ability file reads 0.2143 by mistake (`sim/core/spellconst/spellconst_test.go:106` pins today's fill, so this is a core decision).

## For the guide lane

- **X1:**
  - balance.md:49 and :63 ("Moonkin Form's 3% crit aura") could add "it does not stack with Leader of the Pack".
  - Likewise feral.md:51 and feral-bear.md:48 ("Brings: Leader of the Pack, 3% ...").
  - Guide accuracy rounds 1 and 2 say the index already notes exclusivity; I could not find it in index.md.
  - No wording is wrong today, so this is a Should.
- No guide build, APL action or card rank changed in this lane, so no prose was due and no wow-player pass was needed.

## Still open

- Bear threat constants: Maul, Swipe and Lacerate (B3). Needs a threat capture from a live tank log.
- Eclipse on a missed Wrath; the unarmed paw. Needs a live log of Balance at 40+ and a bare-handed cat.
- The exact Rip and Bite attack power shares. "About 25%" is taken as exactly 1.25x. A live log of a 60 cat's Rip ticks and Bites with a known attack power settles it.
- Pending owner, not reopened: as in the brief.

## Sources

- Blizzard, "WoW Forever Beta Development Notes", https://us.forums.blizzard.com/en/wow/t/2360696 (re-read 5 Nov 07:30 UTC; 24 Sep, 1 Oct, 8 Oct, 16 Oct and 3 Nov druid sections).
- Wowhead Forever tooltips (5 Nov):
  - Spells: 1328621, 1328622, 24907, 24932, 24858, 17007, 9904, 9896, 22829, 9827, 1823.
  - Items: 279250, 249469, 249470, 249473, 19120, 19323, 19024, 20071, 11302.
- wago.tools SpellEffect, build 1.60.1.70790 (re-fetched 5 Nov); cached SpellMisc 70790.
