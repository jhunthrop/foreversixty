# Sim audit round 13: hybrids (2026-11-03)

This report re-audits, with fresh eyes, Feral cat, Feral bear (threat), Enhancement, Retribution and Protection paladin (threat), at every band.

**Setup:**
- Client 1.60.1.70753. Its hotfix-push file (`data/builds/1.60.1.70753/hotfix-push.json`) has zero rows in every table, so the shipped tables are the client. Raw tables come from round 11's copies (`scratchpad/r11/c70753/builds/*/raw`, with `shipped/` and `wago-push/` for 70707).
- Engine: `lane/r13-hybrids` on fork int/r13 (076a5a517). Site: `wt/r13-hybrids` on int/r13 (93c5d2510).
- 10,000 iterations at seeds 1 and 50001 unless a row says otherwise. "Before" is `$R/before.tsv`, or the same binary with the published talents for the band rows (only the talents differ).
- Nothing nightly-owned was regenerated or committed. The druid APL copies from `make apl-sync` stay uncommitted.

Fix-sha column: "E" is a fork commit on `lane/r13-hybrids`, "S" a site commit on `wt/r13-hybrids`.

**In short: 2 CONFIRMED-BUGs, both fixed.**
- The band 30 and 40 bear builds left out Lacerate and Primal Bite (old).
- Templar's Bulwark was gated at level 30, which a 70707 hotfix removed (new-data).

## Method: what was checked, and how

- **Conformance golden** (`sim/core/testdata/conformance/{druid,paladin,shaman}.golden.md`): every registered spell's cost, cooldown, cast time, GCD, required level, duration and damage roll is compared with the 70753 spellconst at levels 10 to 60. Once Templar's Bulwark is fixed, the only rows that differ are Faerie Fire (Feral): free, 6 s cooldown and a 1 s global in forms, which is intended. No damage roll is "declared, differs".
- **Client rows read directly** for every mechanic the golden cannot see: procs, talent multipliers and masks. Each table below lists them.
- **Breakdowns** at raid 60 and at bare 20/30/40/50/60 for all five specs. **Ladder goldens** (`sim/request/testdata/ladder/*.golden.md`): their "Unresolved" column lists the spells a rotation names at a band but cannot cast.

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Shred:** 155% of weapon damage plus 22/32/44/64/80. The flat bonus is added before the percentage, in effect order (effect 58, then 31). Costs 60 less 6 per point of Shredding Attacks. Takes Savage Fury and Rend and Tear. | 5221/9830 SpellEffect; `shred.go:17-26,57,69,78`; round 3 beta log (1.55 x (31 + 22)) | Aligned | - | - |
| C2 | **Rip:** 6 ticks over 12 s at every combo point count (DurationIndex 29). Per tick: base 15 at rank 6, plus EffectPointsPerResource 25.5 per combo point, plus attack power at 1% per point (4% at 5 points). It always crits, per the owner ruling: 36.6% tick crit at raid 60, the same as white swings. Genesis (client mask holds Rip) and Natural Instinct (mask 0x800000) both reach it. | 9896 rows; `rip.go`; breakdown | Aligned | - | - |
| C3 | **Ferocious Bite:** EffectPointsPerResource 147 (rank 6) and 128 (rank 5), base 82/70 with variance 0.732/0.714, 2.7/2.5 damage per extra Energy, 3% of attack power per combo point. These are the engine's numbers. Predatory Discipline keeps the Energy. Rank 6 (31018) has a trainer row, but it is the Ahn'Qiraj book that `IncludeAQ` holds back; that is policy (2026-10-07-aq-ranks-off.md), not a sim bug. The engine comment "the per-combo-point step is stated nowhere" is out of date: the client states 147, and the engine uses 147. | 31018, 22829 rows; `ferocious_bite.go:27` | Aligned | - | - |
| C4 | **Talents against client masks:** Savage Fury (mask 38912: Claw, Maul, Rake, Shred, Swipe; plus Rake's periodic damage); Natural Instinct (mask 0x839800 / 0x540: every melee special, Rip and Ferocious Bite); Genesis (periodic damage including Rake, Rip and Pounce); Naturalist (all damage, +2% a rank); Nature's Majesty (aura 290); Blood Frenzy (combo point on builder crits); Leader of the Pack (raid buff in form). | mask resolution against SpellClassOptions; `talents.go:250-760` | Aligned | - | - |
| C5 | **Tiger's Fury, powershifting, Omen:** Tiger's Fury has only an AcquireMethod 3 (Season of Discovery rune) row, so no druid learns it. Shifting Power replaces powershifting as the Energy source (rune row, granted by its talent). Omen of Clarity is baseline from 20. | SkillLineAbility 59569, 9402 | Aligned (rounds 1 and 12) | - | - |
| C6 | **Round 12 holds re-judged:** Shred on Clearcasting at band 20 still measures +0.7% / +0.7% horde and +0.4% / +0.7% alliance (45.0 to 45.3). That is under the 1% rotation bar (round 11 F2 practice). The Claw-to-Shred switch at 22 rather than 24 moves no band. Both stay held. | `cat-shred20.json` | Held | - | - |

## Feral bear (threat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **The band 30 and 40 bear builds left out Lacerate and Primal Bite.**<br>**What the client says:**<br>- Lacerate is the Feral talent on the 15-point row. Its prerequisite is Savage Fury 2, and rank 1 (1328617) is learned at 25.<br>- Primal Bite is on the 30-point row, behind Leader of the Pack. Rank 1 (407995) is learned at 40.<br>- A level 30 druid's 21 points hold Lacerate; a level 40 druid's 31 points hold Primal Bite.<br>**What was wrong:** the top-row-down walk spent those points on Shredding Attacks (a Shred discount, plus a Lacerate discount the build cannot use) and on Blood Frenzy. So band 30 never cast Lacerate and band 40 never cast Primal Bite, though the rotation casts both. The ladder golden lists them as Unresolved at 30 and 40.<br>**What changed:** `ladder-first.json` now reserves Savage Fury 2, Lacerate, Leader of the Pack and Primal Bite, each as soon as it fits. Bands 50 and 60 already had both and do not change. Band 20 takes Savage Fury's first rank instead of Feral Swiftness's.<br>**New builds:**<br>- 20: `-55000001...`<br>- 30: `-55230202000100100000`, which adds Leader of the Pack<br>- 40: `-55230332023100100010`<br>**Tank fight, bare, threat per second (both seeds within 0.1):**<br>- Band 20: 58.6 to 61.4 (+4.8%) horde, 57.4 to 60.2 (+4.9%) alliance.<br>- Band 30: 77.3 to 109.0 (+41%) horde, 77.7 to 109.5 (+41%) alliance.<br>- Band 40: 175.8 to 209.1 (+19%) horde, 178.1 to 210.9 (+18%) alliance.<br>- Damage taken rises 1.9%, 1.9% and 0.8%, from the lost dodge points. This is the same trade round 12 accepted for Protection's Holy Shield. | Test first: `TestBearTakesLacerateAtThirtyAndPrimalBiteAtForty` (sim/leveling) failed ("level 30 bear build has Lacerate 0", "Primal Bite 0" at 40) and passes; `TestCuratedLadderFirstIsSound` covers the new row (bear added to its spec table). `hybrids/bear-ladder*.tsv` | CONFIRMED-BUG | old (the walk, 78efc0d3e; round 12 fixed the same shape for Protection only) | S 7cbe92a23 |
| B2 | **Threat model:**<br>- Bear Form +50% (21178, aura 10, amount 50).<br>- Maul and Swipe x1.75: Classic server values; no client column.<br>- Lacerate and Primal Bite x2.0: "high amount of threat", unconfirmed, as recorded.<br>- Cat Form x0.71.<br>- Feral Instinct is +30% Swipe damage only (16947) and the engine adds no threat for it. | `threat.go`, `bear_form.go:28`, `swipe.go:49` | Aligned (threat numbers stay single-source, as recorded) | - | - |
| B3 | **The tank fight:** chance of death 98% at raid 60, against Protection's 0.1%. The bear has no block or parry and takes 15% crushing blows. That is Era's bear trade (round 6 G2), not a form bug. Rage sources match their client rows: damage taken, Blood Frenzy +5 on 100%, Natural Reaction +5 on dodge, Enrage, Furor's 10 on shift, Mighty Rage Potion twice (30-second potions). | `sim/request/tank-encounter.json`; breakdown | Aligned | - | - |
| B4 | **Genesis's client mask resolves to Lacerate** (with Rake, Rip, Pounce, Moonfire and Insect Swarm), but `DruidSpellMaskPeriodicDamage` leaves it out. No bear build takes Genesis, so this is latent. | Genesis 1223081 mask resolution | OPEN (latent; no build reaches it) | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Stormstrike:** 8% of base mana, 8 s cooldown, normalized weapon damage, then +20% for 12 s on one charge. The charge is spent by Lightning Bolt, Chain Lightning or Earth Shock: mask 1048579 resolves to exactly those three spells. | 17364 rows; `stormstrike.go` | Aligned | - | - |
| N2 | **Rage of the Farseer:** +30% attack speed (aura 342) for 25 s on a 180 s cooldown, with no global cooldown. It is cast 1.49 times a fight because the fight length varies around 180 s. | 425336 rows; `talents.go:875` | Aligned | - | - |
| N3 | **Windfury Weapon:** 20%, 1.5 s, two hits with 333 attack power, +40% from Elemental Weapons 3. These are yellow weapon-damage hits (enchant 439431 triggers 439440/439441; round 10 C9). The Era row 16361 (extra attacks) is not what the enchant uses. | 16361, 16266 rows | Aligned (round 10) | - | - |
| N4 | **Flurry** (3 charges, white swings spend them with a 500 ms ProcCategoryRecovery, any melee crit refreshes it). **Elemental Devastation** (spell crits only; the totem's attacks carry ProcMaskEmpty and cannot proc it). **Maelstrom Weapon**, **Thundering Strikes**, **Mental Dexterity**. | 16257, 30160/30165 rows; `talents.go:197,443-534`; `fire_totems.go` | Aligned | - | - |
| N5 | **Shamans never dual wield:** every band and the raid build hold a two-handed or one-handed main hand with no off-hand weapon (raid: Forest Defender's Axe). | bis json | Aligned | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The rotation's spells match the client** on cost, cooldown, level and damage roll at every band (conformance golden). The Seal opener and switching (Unsealed Wrath only in the raid build), Holy Strike (effect 121 then 31, so the flat bonus is halved, in effect order), Sacred Arbiter and paced Consecration are as round 12 left them. | conformance golden; round 12 A2/P2 | Aligned | - | - |
| R2 | **The band builds never cast Seal of Command:** the rotation names it only with Unsealed Wrath known, and the leveling build has none by the guide's choice. The ladder golden's Unresolved 20375 at 20 and 1310735 at 40 to 60 are therefore expected. | ladder golden; guide `build:` | Aligned | - | - |

## Protection paladin (threat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **Templar's Bulwark was gated at level 30.**<br>**What the client says:** 70707 shipped Templar's Bulwark (1311015) at SpellLevels 30 (row 153328). Hotfix push 113139 set that row to 0, and 70753 ships 0. The talent (the 10-point row) is the only gate, so a paladin can use it from 19.<br>**What was wrong:** the engine still gated it at 30. Round 6's pipeline lane flagged this (P9) but read the direction backwards.<br>**What changed:** `templarsBulwarkLevel` is 0.<br>**Effect:** no published band takes the talent below 30, so no number moves today. The paladin conformance golden's five mismatch rows are now matches. | Test first: `TestTemplarsBulwarkMatchesTheClient` (sim/paladin) failed ("required level 30, client 0") and passes; protection and retribution goldens unchanged (`--tags=with_db`) | CONFIRMED-BUG | new-data (70707 push 113139) | E 0c7ed03ba, E 50cb4ed67 (conformance golden) |
| T2 | **Band 20 Protection, Bulwark against Sacred Duty's second point.**<br>- At 20 the walk spends 2 points in Sacred Duty. It only shortens Divine Shield (34), Divine Protection (not cast) and Bulwark (not taken), so it does nothing in the rotation.<br>- Templar's Bulwark fits there, and the rotation casts it.<br>- With T1, `-15301100...` against `-15302000...`: chance of death 77.0 to 65.8% horde and 74.6 to 61.4% alliance; damage taken -1.8%; but threat -2.0% (39.5 to 38.7, both seeds). Bulwark's 110 mana is real at 20.<br>- The tank's published number is threat, so this is not adopted. | `-tank` runs | OPEN (owner/coordinator: a survival gain against a 2% threat loss) | - | - |
| T3 | **Righteous Fury:** +90% Holy threat (effect 0) times +33% all threat (effect 4, the hotfix), multiplied; Improved Righteous Fury is damage taken, not threat. **Iron Creed:** +30% Holy Strike threat (mask 0x8000), and a damage-taken proc on a landed Holy Strike. **Shieldsworn:** Holy Strike 4 s shorter, and it fires Holy Shield and Shield Specialization. **Redoubt:** 10% on melee taken, +30% block for 10 s or 5 blocks. All match the engine. RF effect 2 (-100% to Holy Strike's effect 3, a script effect with no amount) changes nothing the sim models. | 25780, 20468, 1311034, 1328544, 20127/20128 rows; `righteous_fury.go`, `holy_strike.go` | Aligned | - | - |
| T4 | **Holy Shield crit** (round 11 O3, round 12 open): still no level 40+ Forever log. | - | OPEN | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The walk leaves rotation-cast talents out at bands where they fit**, in other specs too. A sweep of every ladder golden's Unresolved column kept the talent spells that meet all four tests: (a) the guide build takes them, (b) their client SpellLevels row is at or below the band, (c) the band's points meet the talent's tier gate, (d) they appear in the rotation. That leaves 43 band rows (`hybrids/walk-gaps.txt`). Examples: Restoration druid Swiftmend at 30 to 40, Arms Mortal Strike at 40, Fury Death Wish at 30, Affliction Siphon Life at 30, Destruction Conflagrate at 30, Elemental Lava Burst at 30, Assassination Mutilate at 30/38, the mages' Arcane Blast/Pyroblast/Ice Lance at 20. Each still needs its own measurement and the ladder-first adoption test. Some will hold because the walk's points are worth more. | `hybrids/walk-gaps.txt` | OPEN (for other lanes) | - | - |

## Before/after

**Raid 60:** unchanged by construction. The bear change touches bands 20 to 40 only, and the Bulwark gate matters only below 30. Measured at 10,000 iterations, horde:
- Protection: 1125.7 / 1124.4 threat/s, the same as `before.tsv` at seeds 1 / 50001.
- Every other spec in the group: 0.00%.

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-retribution | 0.00% | 0.00% | 0.00% | 0.00% | unchanged |
| paladin-protection (threat) | 0.00% | 0.00% | 0.00% | 0.00% | 1125.7 -> 1125.7 |
| shaman-enhancement | 0.00% | 0.00% | 0.00% | 0.00% | unchanged |
| druid-feral | 0.00% | 0.00% | 0.00% | 0.00% | unchanged |
| druid-feral-bear (threat) | 0.00% | 0.00% | 0.00% | 0.00% | unchanged |

**Bear bands, bare, tank fight, threat** (B1): 20 +4.8 / +4.9%, 30 +41 / +41%, 40 +19 / +18% (horde / alliance; seeds 1 and 50001 agree within 0.1 threat/s). Absolute at seed 1, horde: 58.6 -> 61.4, 77.3 -> 109.0, 175.8 -> 209.1.

## Files to regenerate

- Site `TestRotationLadder` golden `sim/request/testdata/ladder/druid-feral-bear.golden.md`: the bands 20, 30 and 40 builds. Lacerate at 30 and Primal Bite at 40 resolve.
- Published BiS and leveling-strip data for the bear at bands 20 to 40: `bis/druid-feral-bear.json`, `addon-builds.json`, and anything else that carries the band build. These move at the refresh.
- Fork goldens regenerated here: `sim/core/testdata/conformance/paladin.golden.md` only (E 50cb4ed67).
- The druid APL copies, as round 12 listed (apl-sync).

## Measure at refresh

- Bear bands 20, 30 and 40, both factions: new band builds, so threat, gear picks and weights change.

## For other lanes

- **Coordinator, ladder owners of the other specs (X1):** `hybrids/walk-gaps.txt` lists 43 band rows where the walk leaves out a talent the rotation casts, though it fits. Two have been fixed with ladder-first reservations: Protection 30 (round 12) and the bear 30/40 (this round). The rest need the same measurement.
- **Coordinator / owner (T2):** Protection band 20, Templar's Bulwark for Sacred Duty's second point: survival up sharply, threat -2%. Adopt or hold.
- **Physical / regmelee:** nothing here overlaps regmelee's list. Primal Bite rank 3 and Enrage were not touched.

## Still open

1. Holy Shield crit (T4): no 40+ Forever log yet.
2. Lacerate and Primal Bite threat x2.0, and Maul/Swipe x1.75: single-source, as recorded.
3. Genesis on Lacerate (B4): latent.
4. The band 20 Clearcasting Shred line (C6): held under the 1% rotation bar.
