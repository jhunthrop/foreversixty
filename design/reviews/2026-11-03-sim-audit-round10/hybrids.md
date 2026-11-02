# Sim audit round 10: hybrids (2026-11-02)

This report covers Retribution, Protection paladin (threat), Enhancement, Feral cat and Feral bear (threat).

**Setup:**
- Client 1.60.1.70707, layered tables (hotfix push included), the 2026-11-01 21:42Z fetch (`scratchpad/r9/hybrids/raw`).
- Published sets and builds: the committed `data/builds/1.60.1.70707/bis/`. Nothing nightly-owned was regenerated or committed.
- Engine: pin 4d3129cae plus `lane/r10-hybrids`. Site: `wt/r10-hybrids` on origin/main bb619a188.
- Method:
  - every band (bare 20 to 60, raid 60) of every spec was run and its cast and aura tables read;
  - every talent in the published builds was decoded and read against its client row (SpellEffect, SpellAuraOptions proc masks, SpellMisc attributes, SpellClassOptions masks, trait curves);
  - for the crit split, every effect that reads crit was checked for its proc mask, its table and whether it may fire off a triggered spell (aura Attributes_3 0x4000000 "can proc from procs", spell Attributes_3 0x200 "not a proc");
  - Retribution's first raid iteration was read in the engine's debug log (a scratch copy of spec-breakdown with `Debug`, not committed).

Fix-sha column: "E" is a fork commit on `lane/r10-hybrids`, "S" a site commit on `wt/r10-hybrids`.

**In short: 3 CONFIRMED-BUGs, all fixed.** One moves a published number: Retribution's raid rotation, +9.3 to +9.4%. Exorcism's crit (P4, an integrator follow-up) moves only Undead and Demon targets. The third is a display id.

## Assigned items: crit-split side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Vengeance** (20049) triggers on any non-periodic crit, melee or spell, including seal procs and Judgements. Client: proc mask 69972 (0x11154: melee swings, melee- and ranged-class spells, and harmful magic and no-class spells, no periodic bit) and Attributes_3 0x4000000, so triggered spells (seal procs, Judgement damage, Thrash's extra swings) proc it too. Engine: `OnSpellHitDealt` on `DidCrit()` (`paladin/talents.go:150`); ticks go to the periodic handler. Consecration's tick 1280349 is a direct hit, but it carries Attributes_3 0x2000000 (treated as periodic), so it does not proc Vengeance, and the engine's dot does not either. Its buff 20050 is aura 79 for Physical and Holy (misc 3), 3% a stack at 3/3, three stacks, 30 s: matches `vengeanceMultiplier`. | client rows; engine | Aligned | - |
| C2 | **Conviction** (20117) is aura 290 (every table), and it needs a melee weapon equipped (SpellEquippedItems class 2, mask 173555). Engine `stats.Crit`. | client | Aligned | - |
| C3 | **Sanctified Seals** is not in Forever's tree. The Holy tree's **Improved Seals** (20224, +15%) masks Judgement of Righteousness, Command and Fury and the Righteousness, Command and Fury procs (family 10 masks 0x2000400, 0x20000A00, 0x40). This matches `PaladinSpellMaskSealsAndJudgementsDamage`. | SpellClassOptions scan | Aligned | - |
| C4 | **Seal and Judgement crit tables.** Seal of Righteousness 25713, Seal of Command 20424, Judgement of Righteousness 20286, Judgement of Command 20966, Seal and Judgement of Fury, and Holy Strike 10333 are all DefenseType 2, so they roll melee crit at 200%. Consecration 20924 and 1280349, Exorcism 10314, Holy Wrath and Holy Shield 20925-20928 are DefenseType 1, so they roll spell crit at 150%. Hammer of Wrath 24239 is DefenseType 3 (ranged). The engine sets each the same way (`sor.go:96,132`, `soc.go:82,118`, `consecration.go:88`, `exorcism.go:52`, `holy_shield.go:76`, `hammer_of_wrath.go:50`) and `critMultiplierAs` gives magic 150% and the rest 200%. Raid 60 measures melee 23.0%, Consecration 25.6% (Intellect plus the 3.5 spell base). | SpellCategories; breakdowns | Aligned | - |
| C5 | **Hammer of Wrath and the sharpening stone.** The engine takes the stone's 2% back off ranged rolls, so Hammer of Wrath crits 21.0% against 23.0% melee. The stone's spell 22755 has no SpellEquippedItems row. In TrinityCore, an aura-52 spell with no equipped-item class raises melee, off-hand and ranged crit alike, which would put Hammer of Wrath at 23%. That is an emulator, one source. About 0.1% of Retribution's damage. | `scratchpad/r10/hybrids/bd/paladin-retribution-raid60.txt` | OPEN (crit lane's round 9 "which hands an aura-52 enchant reaches") | - |
| C6 | **Holy Shield crit** (carried, round 9 T1). The client row is a damage shield (aura 43). The engine rolls spell crit (19.0% at raid 60). I searched 960 cached beta character breakdowns (`scratchpad/r9/crit/fl`, `fl2`) and found no Holy Shield. It is a level-40 talent and the beta stops near 30. | cached logs | OPEN (needs a 40+ log) | - |
| C7 | **Enhancement: Flurry.** The talent 16256 procs on mask 20 (white swings and melee-class spells) from a crit, so Stormstrike and Windfury Weapon's hits count and spells do not. The buff 16257 has 3 charges, spent by white swings only (mask 4), with a 500 ms ProcCategoryRecovery. Windfury's hits 439440 and 439441 are triggered, but they carry Attributes_3 0x200 ("not a proc"), so they may proc Flurry although the talent lacks "can proc from procs". The engine matches (`shaman/talents.go:458`, `:514`). | client rows | Aligned | - |
| C8 | **Enhancement: Elemental Devastation.** The talent 30160 procs on mask 0x10000 (harmful magic spells, no periodic) and cannot proc from procs. Its buff 30165 is aura 52, melee crit, +9% for 10 s. The engine's `ProcMaskSpellDamage` trigger leaves out item procs and Searing Totem's attack (that one is `ProcMaskEmpty`; in game it is the totem's spell, not the shaman's). The buff's display id was still Era's 29178: see X1. | client; breakdown | Aligned (display fixed) | E af3901df3 |
| C9 | **Enhancement: Windfury Weapon.** Every rank's enchant (283, 284, 525, 1669) is an equip spell for 439431: 20% chance, 1.5 s ProcCategoryRecovery, proc mask 20. That triggers 439440 and 439441, melee-class (DefenseType 2) weapon-damage hits. The engine's two yellow hits (two-roll, can be dodged, no glancing) are this model. The fork's "vanilla uses two extra attacks" TODO describes the Era row 16361, which no enchant uses. | SpellItemEnchantment, SpellEffect | Aligned | - |
| C10 | **Enhancement:**<br>- Stormstrike 17364 (melee table; its 20% charge spent by the next Lightning Bolt, Chain Lightning or Earth Shock to land; 8% base mana, 8 s).<br>- Earth Shock and Flame Shock (DefenseType 1, spell table).<br>- Lightning Shield (not cast).<br>- Maelstrom Weapon: 50% a melee hit (effect 1, not curved); trait curve 85306 is −4 to −20% cast time and cost a stack; the proc mask's magic bit is the Totem of the Storm case the engine models.<br>- Thundering Strikes: aura 290, needs a melee weapon.<br>- Weapon Mastery is not in Forever's tree. | client rows, TraitDefinitionEffectPoints 23728 | Aligned | - |
| C11 | **Druids:**<br>- Primal Fury is Forever's **Blood Frenzy** (16958): proc mask 20 (no periodic, no spells), can proc from procs. A bear's melee crit gives 5 Rage; a cat builder's crit gives a combo point. Engine `druid/talents.go:247` matches.<br>- **Natural Instinct** (1223242): +20% crit damage bonus, masked on Claw, Ferocious Bite, Lacerate, Maul, Pounce, Primal Bite, Rake, Ravage, Rip, Shred and Swipe. That equals `DruidSpellMaskMeleeAbilities \| Rip`, and it reaches Rip's and Lacerate's ticks through the dot crit multiplier.<br>- **Sharpened Claws**, **Nature's Majesty** and **Leader of the Pack** (24932) are aura 290, every table. The talent's own Leader of the Pack and the raid's are one aura.<br>- **Predatory Strikes** is attack power by level, no crit.<br>- Rip, Rake and Lacerate ticks roll the melee table (bleed rule). Rip measures 34.5%, the same as white. | client rows; breakdowns | Aligned | - |
| C12 | **Item effects on the published sets.** I listed every use and equip spell on the 383 distinct items across the five specs' bands. None reads crit or procs on crit. | `scratchpad/r10/hybrids/itemprocs.txt` | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **The Unsealed Wrath twist traded the two Seals every global cooldown.** It should have waited for one swing.<br>- What the client says: Unsealed Wrath (1310735) stores a copy of the active Seal on each swing, three at most. A swing that lands under a different Seal releases them.<br>- What the rotation meant to do (its own note): store three Seals of Command, switch to Seal of Righteousness, and "after one swing under Seal of Righteousness, go straight back".<br>- What was wrong: the way back fired on "Unsealed Wrath has at least one stack", and the three stored Command copies already met that. So 1.5 s after the switch, Seal of Command went back up, usually before any swing. The two Seals then traded every global until a swing happened to land under Righteousness. The debug log shows eight Seal casts in a row from 11.5 s to 23.5 s.<br>- The cost: about 28 casts of each Seal a fight, each a global and about 200 mana. Holy Strike fell to 15 casts and Consecration to 7.<br>- What changed: the way back waits until Unsealed Wrath restarted after the Seal was cast (its time left plus 15 s exceeds the Seal's time left). That moment is the Righteousness swing landing.<br>- Result: Seals 15 casts each, Holy Strike 18, Consecration 16. No time out of mana.<br>- Bare bands are unchanged: no bare build takes Unsealed Wrath.<br>- For the wow-player review: this is the twist the guide already describes, now done the way it says. | Debug log of iteration 1 (`scratchpad/r10/hybrids/ret.log`); `data/curated/apl/paladin-retribution.json` line 6. Test first: `TestRetributionTwistWaitsForTheRighteousnessSwing` (sim/cmd/spec-breakdown) failed at 0.59 Righteousness casts a landed swing, and passes now. | CONFIRMED-BUG (rotation, engine log) | S a30f50a3e, E 32ccbcc64 (fork copy) |
| P2 | Checked fresh:<br>- Hammer of Wrath ranks 24275/24274/24239 at 44/52/60: 1 s cast, 1 s global, 6 s cooldown, 295/360/425 mana. With Champion of the Light 5/5 it is instant.<br>- Holy Strike ranks by band (1866 at 20, 2495 at 40, 5569 at 50, 10333 at 60).<br>- Seal of Righteousness and Judgement ranks by band.<br>- Seal of Command's proc, 20424, can miss and be dodged (no always-hit or no-active-defense attribute), as the engine rolls it. Judgement of Command (Attributes_3 0x40000) always hits. Judgement of Righteousness can miss but not be dodged.<br>- Precision is melee and spell hit (aura 54 and 55).<br>- Two-Handed Weapon Specialization: physical only.<br>- Sacred Arbiter masks Holy Strike.<br>- Unsealed Wrath's stored Righteousness copy released by a Windfury Totem swing gives one Righteousness hit (log, 26.51 s), as the text says.<br>- Frozen Heart's Chthonic Power is Frost and Shadow only, so it gives Retribution nothing. | client rows; log | Aligned | - |
| P3 | The rotation's top-level note is stale: it says "the guide builds do not take Unsealed Wrath" and "Consecration is deliberately left out". The raid build takes Unsealed Wrath and the list casts Consecration. It is prose only. I left it, because the ranks lane rewrites the Consecration line this round. | APL notes | Note (for the ranks lane) | - |
| P4 | **Exorcism crit in the engine; in Forever it never crits.** (Integrator follow-up, from the crit lane's log tally.)<br>- What the logs say: public foreverlogs.gg reports (39 paladins at levels 20 to 30, the crit lane's cached breakdowns `scratchpad/r9/crit/fl`, `fl2`, read 2026-11-02) show 0 crits in 641 Exorcism hits (879: 320 hits, 5614: 321). The same paladins' Judgement of Righteousness crit 18 of 367 (4.9%). At their roughly 5 to 6% spell crit, about 35 Exorcism crits were due; the chance of none is about e^-35.<br>- What the client says: Exorcism 879 to 10314 is DefenseType 1 and carries no cannot-crit attribute (Attributes_2 0x20000000 is clear), so the rule lives on the server. The Season of Discovery copies 415068 to 415073 sit on another skill line and are not trained. The hotfix push does not touch these rows.<br>- What was wrong: the engine rolled spell crit (24.5% at raid 60 against Undead).<br>- What changed: Exorcism rolls hit only. The fork's TestProtection golden moves with it.<br>- Published numbers do not move: Exorcism only lands on Undead and Demons, and the ranker's target type is unknown. Against Undead: Retribution raid 60 -1.0% (721.4 to 714.0), Protection raid 60 -0.8 / -0.6%, Retribution bare 40 -1.0%. | logs above. Test first: `TestExorcismNeverCrits` (sim/cmd/spec-breakdown, Undead target) failed at 24.5% and passes now. | CONFIRMED-BUG (logs, 641 hits) | E f1c7fa888, S (test) |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Checked fresh:<br>- Holy Shield 20928: 4 charges, 10 s, 10 s category cooldown, 195 mana, damage shield 153 + 0.08 coefficient. Its crit is open (C6).<br>- Hammer of the Righteous and Holy Strike roll melee crit; the boss parries and blocks them from the front.<br>- Judgement of Fury and Swift Judgement.<br>- Shieldsworn's shield strike fires Holy Shield's damage.<br>- Consecration's tick is a direct spell; spell crit 19.2%.<br>- The rotation's mana reserve for Holy Shield.<br>- No published item reads crit (C12). | client rows; breakdowns | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | Checked fresh, besides C7 to C10:<br>- Elemental Weapons: rows 0x400000 Rockbiter +20%, 0x800000 Windfury +40%, 0x1200000 Flametongue and Frostbrand +15%.<br>- Mental Dexterity: aura 598, attack power 100% of Intellect.<br>- Mental Quickness: aura 174 and 175, 30% of Intellect.<br>- Shamanistic Focus: −45% on Shocks and Lightning Shield.<br>- Rage of the Farseer: aura 342, 30%, 25 s, 3 min.<br>- The rotation: Stormstrike, a free Lightning Bolt at five Maelstrom stacks, Searing Totem, then Earth Shock above 50% mana. | client rows; breakdowns | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Checked fresh:<br>- Shred 155% + 80, Ravage 350% + 98, Rake, Rip and Ferocious Bite rows, against `shred.go`, `ravage.go` and `rip.go`.<br>- Blood Frenzy 1/2: a 50% combo point on a builder crit.<br>- Natural Instinct, Rend and Tear and Savage Fury masks.<br>- Shifting Power.<br>- Rip's bleed crit.<br>- Ferocious Bite rank 6 (31018) is a book.<br>- No item reads crit. | client rows; breakdowns | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | **Primal Bite 1238073** ("rank 4" in spellranks.json and in the rotation) has no SkillLineAbility row. The learnable top rank is 1238070, which supersedes 1238069. The two rows are identical (100% weapon + 77, 6 s, 20 Rage), so damage and threat do not move. Display only: the breakdown names a spell no druid trains. | SkillLineAbility 53141, 53142, 53146; SpellEffect | Note (cosmetic; the spellranks pipeline is the ranker's) | - |
| G2 | Checked fresh:<br>- Blood Frenzy 2/2 gives a bear 5 Rage on any melee crit, 21 procs a fight.<br>- Natural Reaction.<br>- Lacerate ticks crit on the melee table.<br>- Maul, Swipe and Lacerate rows.<br>- Threat items carried (round 9 A2, A3). | client rows; breakdowns | Aligned | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Elemental Devastation's buff showed Era's id.** The breakdown read "spell (29178)"; the client has one buff, 30165, at every rank (29177 and 29178 have no SpellName row). Round 9 routed this to the crit lane, and it was still open. Display only. | SpellName. Test first: `TestElementalDevastationBuffCarriesTheClientsId` (build failure, then passes). | CONFIRMED-BUG (display) | E af3901df3 |
| X2 | **Hit is one stat in the engine** (melee and spell). For my specs that is right: Precision states aura 54 and 55, and only 3 item equip spells in the whole client state melee or ranged hit alone (1294769, 1301127, 1306681); item hit is rating. Across every tree, the only talents with one hit table are Warrior's Furious Precision 1323963 (aura 54 only) and Mage's Precision 29438 (aura 55 only). | SpellEffect scan | Aligned (my group); see "For other lanes" | - |

## Before/after

Before: `$R/sb-before` (pin 4d3129cae, site origin/main bb619a188). After: `lane/r10-hybrids` 32ccbcc64 with `wt/r10-hybrids`. 10,000 iterations, raid 60.

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-retribution | +9.4% | +9.4% | +9.3% | +9.4% | 621.3 -> 679.6 |
| shaman-enhancement | 0.0% | - | - | - | 575.8 -> 575.8 |
| paladin-protection (threat) | 0.0% | - | - | - | 432.4 -> 432.4 |
| druid-feral | 0.0% | - | - | - | 593.0 -> 593.0 |
| druid-feral-bear (threat) | 0.0% | - | - | - | 306.8 -> 306.8 |

Retribution bare 50 and 60 (3,000 iterations, seed 1) are unchanged (165.9 and 241.3).

Exorcism (P4) alone, from engine 32ccbcc64 to f1c7fa888, horde, 10,000 iterations: the published unknown-target numbers do not move (Retribution raid 679.6, Protection raid 432.4, both bare 60). Against Undead, seed 1 / 50001: Retribution raid 60 -1.0 / -1.0%, Protection raid 60 -0.8 / -0.6%, Retribution bare 40 -1.0 / -1.1%, Protection bare 40 0.0 / -0.1%.

## For other lanes

- **Crit lane:**
  - C5, the stone and ranged rolls: TrinityCore gives an aura-52 spell with no equipped-item class to ranged crit too, so it would reach Hammer of Wrath.
  - X2: Furious Precision is melee-hit-only and today also raises the warrior's spell-table weapon procs (negligible).
- **Ranks lane:** P3. Retribution's top-level note still says Consecration is left out and Unsealed Wrath is not taken; rewrite it with the Consecration gate change. My commit only touches line 6 (the way back to Seal of Command) and its note.
- **Ranker / spellranks:** G1. Primal Bite 1238073 is not learnable (no SkillLineAbility row); the bear rotation and spellranks.json name it as rank 4. No number moves.
- **Integrator:**
  - Retribution raid 60 moves +9.4%, so the tier list, BiS picks and weights move for that row at the refresh.
  - `TestRotationLadder` goldens were not regenerated.
  - Nothing outside my group moves.
  - The fork commit 32ccbcc64 is the APL copy.

## Still open

1. C6, Holy Shield crit: needs a 40+ Forever log.
2. C5, the stone on ranged rolls: crit lane.
3. Carried from round 9: Faerie Fire's form threat; Lacerate and Primal Bite's 2.0 threat; Righteous Fury's multiplication.
4. Seal of Command and Stormstrike have no SpellEquippedItems rows. In TrinityCore's rule only weapon-class spells fire weapon-enchant procs, which would keep both away from Windfury Totem and weapon procs; the engine lets both trigger them. Emulator-only evidence, not changed.
