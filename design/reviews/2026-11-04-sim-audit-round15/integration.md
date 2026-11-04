# Sim-accuracy round 15: integration (2026-11-04)

Round 15 started from the landed round 14 (site origin/main fd320b60e, fork pin d1d9f2775) on client 1.60.1.70755. One auditor per class, plus a platform lane for the guide page components. Ten waves folded in the seeds that guide-accuracy rounds 1 to 9 handed over, and the coordinator's later items. The lane reports sit beside this file: warrior, rogue, shaman, priest, druid, warlock, mage, hunter, paladin and platform.

**Not a zero round.** About 70 confirmed bugs were fixed (see "Confirmed by lane"). Round 16 is needed.

**Branches.**
- Site `int/r15`: every lane branch plus origin/main up to the 70755 hotfix re-layer (43dc0f074) and guide-accuracy rounds 2 to 7, with the pin and one regen.
- Fork `int/r15` at **9f7f1f6e4**, pushed to origin. Fork master is not fast-forwarded.
- Pin: `chore(engine): pin 9f7f1f6e4`.

## Regenerated once, at the pin

- Ladder goldens: 26 specs.
- addon-builds.json, addon-data.json and Data.lua.
- Coach starters (persona_starter --all) and coach rules.
- dps-per-point.json for every spec.
- genraidconsumes. Unchanged after the shaman lane's own regen.

apl-check passes on all 56 copies, and apl-sync wrote nothing. The fork keeps the curated notes out by design. ladder-first, ladder-order and rank-choices copies match.

Not regenerated, by rule: BiS files, simdb.bin and weights. The CI BiS refresh owns them and waits on this pin.

## Before/after, raid 60

**Method.** Before is the round-start engine d1d9f2775 on the site at fd320b60e. After is int/r15 at the pin.
- **Gear:** both sides run on the same frozen copy of main's newest published BiS gear (de044125c). The nightly BiS refresh raised most specs 3 to 15% on its own, and that gear move is left out of this table.
- **Talents:** each side runs on its own guides' raid build. The after side's builds include guide-accuracy's build changes from main as well as round 15's. Destruction's +2% is guide round 2's Pandemic/Malevolence build, not a round 15 change.
- **Runs:** 10,000 iterations at seeds 1 and 50001, both factions. Healers are in effective HPS and tanks in threat per second.

Specs moving more than 1% (seed 1 values; the last column is % at seed 1 / 50001):

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| druid-balance | alliance | 609.2 to 652.8 | +7.16 / +7.17 |
| druid-balance | horde | 603.8 to 646.9 | +7.14 / +7.12 |
| druid-feral | alliance | 724.7 to 871.6 | +20.27 / +20.25 |
| druid-feral | horde | 715.7 to 856.8 | +19.71 / +19.83 |
| druid-feral-bear | alliance | 1073.3 to 1200.3 | +11.83 / +11.83 |
| druid-feral-bear | horde | 1067.8 to 1194.2 | +11.84 / +11.85 |
| druid-restoration | alliance | 1064.1 to 939.5 | -11.71 / -11.71 |
| druid-restoration | horde | 1060.0 to 935.3 | -11.76 / -11.77 |
| hunter-marksmanship | alliance | 786.7 to 801.5 | +1.88 / +1.94 |
| hunter-marksmanship | horde | 785.3 to 801.2 | +2.02 / +2.04 |
| mage-arcane | alliance | 674.8 to 717.6 | +6.34 / +6.36 |
| mage-arcane | horde | 664.7 to 704.3 | +5.96 / +6.11 |
| paladin-holy | alliance | 700.7 to 671.7 | -4.14 / -4.18 |
| paladin-holy | horde | 687.1 to 655.6 | -4.58 / -4.73 |
| paladin-protection | alliance | 1136.6 to 1151.9 | +1.35 / +1.27 |
| paladin-protection | horde | 1135.1 to 1150.0 | +1.31 / +1.30 |
| priest-discipline | alliance | 611.5 to 504.6 | -17.48 / -17.51 |
| priest-discipline | horde | 598.4 to 494.5 | -17.36 / -17.39 |
| priest-holy | alliance | 846.8 to 773.8 | -8.62 / -8.62 |
| priest-holy | horde | 857.8 to 793.5 | -7.50 / -7.60 |
| rogue-assassination | alliance | 765.3 to 722.9 | -5.54 / -5.47 |
| rogue-assassination | horde | 767.5 to 726.2 | -5.38 / -5.50 |
| rogue-subtlety | alliance | 678.4 to 722.3 | +6.47 / +6.39 |
| rogue-subtlety | horde | 673.9 to 716.8 | +6.37 / +6.46 |
| shaman-elemental | alliance | 661.9 to 669.4 | +1.13 / +1.16 |
| shaman-elemental | horde | 659.1 to 666.1 | +1.06 / +1.08 |
| shaman-enhancement | alliance | 823.5 to 855.3 | +3.86 / +3.85 |
| shaman-enhancement | horde | 828.1 to 855.3 | +3.28 / +3.41 |
| shaman-restoration | alliance | 793.0 to 759.4 | -4.24 / -4.26 |
| shaman-restoration | horde | 792.0 to 757.1 | -4.41 / -4.42 |
| warlock-affliction | alliance | 638.7 to 736.7 | +15.34 / +15.30 |
| warlock-affliction | horde | 634.4 to 733.1 | +15.56 / +15.48 |
| warlock-demonology | alliance | 707.1 to 716.6 | +1.34 / +1.43 |
| warlock-demonology | horde | 697.3 to 706.7 | +1.35 / +1.32 |
| warlock-destruction | alliance | 765.0 to 780.2 | +1.99 / +1.98 |
| warlock-destruction | horde | 761.6 to 776.4 | +1.94 / +1.96 |
| warrior-arms | alliance | 777.9 to 820.4 | +5.46 / +5.26 |
| warrior-arms | horde | 784.3 to 826.4 | +5.37 / +5.46 |

**Within 1%:** Beast Mastery, Survival, Fire, Frost (+0.45), Retribution (+0.98), Shadow (+0.85), Combat, Fury (+0.25 to 0.42) and Protection warrior.

**What moves them:**
- **Cat:** the 70755 server push 113403 followed in the engine (Claw 140%, Shred 200%, Ravage 435% weapon damage, plus new Rake, Rip and Ferocious Bite values on every rank), Rend and Tear at 3% a point, and Nature's Reach in the raid build.
- **Bear:** Rend and Tear, plus Furor 1 / Naturalist 5 (threat +7.7%, chance of death unchanged at 60).
- **Balance:** the pushed Starfire values on all seven ranks.
- **Affliction:** Wrack (the 70755 push raised it to 86 a tick) in the rotation and raid build, plus Bane of Doom.
- **Arcane:** the raid build takes Fire second, with Arcane Wrath.
- **Subtlety:** Shank ranks resolved below 44, Eviscerate held for Death Mark, Preparation fixed.
- **Arms:** dead Deflection and Blood Craze points moved to Spearing Strike, Reckless Swings and Improved Execute.
- **Marksmanship:** Sniper Shot.
- **Enhancement:** Rockbiter under the raid's Windfury Totem.
- **Assassination (down):** Mutilate's off hand now takes the off-hand penalty, as the client marks it. The Rupture line gives about 1% back.
- **Healers (down):** the low-rank penalty below. Rank re-picks recovered most of the engine-only loss (-10 to -20%). Discipline's -17% is larger than the lane's -11 to -12% on the older gear, because a bigger healing bonus on penalised low ranks loses more. Discipline also gains Power Word: Shield rank 10 (+1.0 to 1.4%).

## The low-rank penalty (the round's biggest shared change)

**The rule.** A spell rank keeps its full Spell Power / Spell Healing bonus until the caster is 17 levels past the rank's learn level, then loses 5% a level, down to 0. It applies to heals and damage alike. It replaces Era's sub-20 rule.
- Fork E 8d9ca6aae and 974614b8b.
- Spells that added spell power by hand were made to follow it: Consecration (E 86fe4bf34), Lightning Overload and Power Word: Shield (E 27c8c6367).
- Wrack opts out: it has one rank that scales to the caster's level.

**Evidence: PARTLY CORROBORATED.**
- Blizzard's beta development notes (us.forums.blizzard.com/en/wow/t/2360696), 2026-10-01: "spells cast using spell ranks vastly below your current level receive less benefit from your Spell Damage and Spell Healing". This gives the direction only, and it is the only official sign that damage spells are included.
- ForeverChanges (foreverchanges.pro/downranking, 2026-10-16): measured heals at levels 18 to 30. It gives the 17-level threshold and the 5% slope, and states "Above level 30, and for damage spells, I assume the same rule."
- Ozgar's Forever downrank tool (wowdownrank.forever.ozgar.se, credited Defcamp & Melderon, undated): agrees on the 5% slope, the floor at 0 and the removal of the sub-20 rule. It differs in two ways:
  - It puts the threshold at 16 levels. ForeverChanges itself saw Healing Touch 1 and Holy Light 1 run one level early.
  - It exempts a spell's top trainable rank.
- The client does not encode the rule: no curve table, no tooltip text, Era-shaped SpellLevels.

**Kept as implemented** (17 levels, every rank). It is the only measured version, and Blizzard's wording keys on the level gap, not on top-rank status. The top-rank exemption touches only two cast spells (Heal rank 4, Lesser Heal rank 3), and Holy's filler no longer uses Heal 4.

**Least certain, worst first:** damage spells at any level; levels above 30; 16 versus 17; the top-rank exemption.

**Guide wording.** Keep the per-level numbers out of the guides. Suggested sentence: "In Forever, a low rank keeps its full bonus from Spell Power and Spell Healing for many levels after you learn it. Once you outlevel it by a wide margin it loses that bonus a little each level, until it heals or hits for its base amount only. At 60, the ranks you learned in your teens and early twenties get almost nothing from your gear."

## Weapon DPS in forms (checked before the pin, CONFIRMED)

- **Blizzard's Forever Class Deep Dives: Hunter and Druid** (us.forums.blizzard.com/en/wow/t/2367361, 2026-09-30):
  - "While Shapeshifted, the Druid's melee auto attack DPS ... in Bear Form, Cat Form, or Dire Bear Form is now the same as the DPS of the Druid's equipped weapon."
  - "the speed of attack is changed to 1.0 or 2.5 seconds, depending on the form."
  - "Abilities that deal weapon damage likewise use these damage values."
- **Second sources:** forum thread t/2367553 (2026-09-30, 1.0 s normalization), Zockify's Forever druid page, and Wowhead Forever's tooltips (Claw "140% normal damage", Shred "200% damage").
- **Client:** SpellShapeshiftForm CombatRoundTime is 1000 / 2500.
- **Engine:** `formPawWeapon` (sim/druid/forms.go, round 14 E 4cb26aa16) keeps the weapon's DPS at the form's speed, so a slow two-hander is not inflated. Claw, Shred, Ravage and Maul use it; Rake, Rip, Bite and Swipe are flat.
- **Open:** the form table's ±20% spread versus the weapon's own spread, and the 54.8 DPS unarmed fallback.
- **The earlier cat +61.7% / bear -25% was a coordinator error:** the cat and bear rows were swapped in a table. The real gear moves are +8.4% and +11.7%. The bear's guide raid build matches the BiS file's talents point for point.

## Calls decided this round (nothing went to the owner)

| Call | Decision | Reason |
|---|---|---|
| Bear Furor vs Naturalist | Furor 1, Naturalist 5 | Furor never fires (the bear never shifts in a fight). Threat +7.7% at 60 with chance of death unchanged; band 20 trades -0.5% threat for chance of death 68-76% to 41-49% (the Protection paladin band-20 precedent). |
| Bear Nature's Reach | Rejected | +2% threat, but chance of death doubles. |
| Enhancement imbue | Rockbiter under the raid's Windfury Totem, Windfury Weapon bare | Windfury Weapon turns the totem off for its wielder, and Forever's Rockbiter is 653 AP. +4.0 to 4.6% at raid 60; it loses 1.3 to 3.5% bare. No Forever source picks either; the Classic/SoD advice rests on the old Rockbiter. |
| Discipline Power Word: Shield rank | Rank 10 | Blizzard's deep dive (2026-09-30): absorbs no longer reduce rage, which was round 14's only reason for rank 4. Rank 4 keeps 5% of its bonus at 60. +1.0 to 1.4%, no band loses. |
| Shadow Wand Specialization | First, at 10-11 | Icy Veins Forever ("your first priority is to pick up Wand Specialization") and Kargoz. Band 20 +8 to 9.5%, band 30 +5 to 6%; band 40 -36 to -37% because Shadowform lands at 42. No legal order keeps both. This is the sharpest drop on a published band this round, and it is right for a player levelling in that order. |
| Exorcism crit | Never crits (kept) | No Forever source or client flag. Beta logs 0 crits in 641 casts against about 35 expected. Re-check on a level-60 launch log. |
| Decimation | 3% / 6% (kept) | The 0 in SpellEffect 1341977 is overridden by the trait curve (TraitDefinitionEffectPoints def 135428, curve 110390: 3 / 6). The engine applies it to Shadow Bolt and Searing Pain below 35%. |
| Eureka! (Gnome) | Client mask kept: no Incinerate, a separate +10% on Drain/Wrack ticks | The client table beats the tooltip and Warcraft Tavern; including Incinerate would be worth about 0.2%. No race changes (Destruction Human +0.5%, Discipline Human +0.1%; Assassination stays Night Elf). |
| Haste and the global cooldown | Aligned (Forever keeps Classic's rule) | The cited Unit.SpellGCD was dead code, now removed with a pin test. Casts haste the cast time only; the client uses 1500 on every GCD spell; the deep dive names Nature's Grace as an explicit GCD exception. |
| Lay on Hands | Drains all mana (kept) | The 70755 Forever description: "Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration." Guide round 1's "free" reading was reverted on main. |
| Wind Blessed (Skyborne) | 2% (already) | Production and round 15 both use 1.02. The shipped rows, live push 113276 and the tooltip all say 2. |
| Expertise on items | Items DO carry it (removal retracted) | 71 catalogue items in simdb.bin, e.g. 240081 "Reduces chance to be Dodged or Parried by 2.0%". items/*.json drops the stat (a data-lane bug). |

## Confirmed by lane (approximate, from the lane reports; each report has the tables)

| Lane | Confirmed and fixed | Main ones |
|---|---|---|
| warrior | 4 | Arms dead Deflection/Blood Craze points; Spearing Strike Battle Stance only; Fury dead points to Improved Berserker Rage; the stale Fury expected-names test |
| rogue | 10 | Shank unresolved below 44 (all specs); Subtlety Eviscerate vs Death Mark, Shank/GS overflow, Garrote, Preparation; Assassination Rupture; Combat Efficient Strikes; Mutilate off-hand penalty; Cold Blood order; rank labels |
| shaman | 13 | Stormstrike skipped at band 30; Elemental Alacrity dead points; NS Healing Wave rank 2 to 9; Riptide order and rows; Lightning Bolt rank 10 at raid 60; the Flame Shock and Lightning Bolt rows re-picked on the new penalty; the imbue |
| priest | 12 | the low-rank penalty and its re-picks; Holy Shock heal half, Lightning Overload and Power Word: Shield skipping it; Shadow leveling/raid builds; Mind Blast pre-pull; Starshards had no cooldown; PW:S rank 10; Wand Specialization |
| druid | 9 | the 70755 push (7 spells, every rank); Rend and Tear 3%; cat leveling caster points; Restoration dead Improved Moonfire; Nature's Reach in cat builds; Ravage behind check; Furor/Naturalist |
| warlock | 7 | Wrack never cast; Doom over Agony; Curse of the Elements first; Soul Fire with Demonic Pact; Wrack rank; Shadow Bolt rows at 20-29 |
| mage | 7 | Arcane Power/Presence of Mind on the pull; Frost dead points; Fire leveling Burning Soul; Improved Frostbolt first; Arcane Concentration then Blast; Scorch rank 2 at 40-49; the Arcane Fire-second raid build |
| hunter | 4 | BM Focused Fire order; MM Sniper Shot; MM leveling Rapid Killing; Falcon's own 5% Quick Strikes proc |
| paladin | 9 | Holy leveling dead points; Kings vs Heart of the Lion; Consecration skipped the penalty; Holy Shock rank 4; Stoneform off the GCD; Crusader judgement missing on Consecration ticks; Crusader judgement stuck at rank 1's 23; the Crusader opener; Dwarf Protection Stoneform line |
| platform | 2 | "raid build moves N points" hid talents; the hit/expertise cap figures read on the weights-sweep character and the wrong race |
| **total** | **about 77** | |

By origin, most are `old`. Over 30 came from guide-accuracy hand-offs. About 10 are `70755` or `new-data` (the druid push, Wrack, Shank ranks, Rogue tree, Lightning Bolt rank 10). One is `R14` (shaman cf5fe63bb).

## For the guide lane (the lane reports have exact sentences)

- Every healer APL note that says "no downrank penalty" or names the old ranks.
- Affliction (Wrack, Doom, curse order).
- Arms raid paragraph, plus notes for Spearing Strike and Reckless Swings.
- Fury (Lingering Rage, Piercing Howl, Berserker Rage).
- Arcane (the Fire-second build and its four new lines).
- Frost and Fire leveling order.
- Shadow (Wand first, Spirit Tap, the pre-pull Mind Blast).
- Cat builds (Nature's Reach) and the bear's Furor/Naturalist.
- Enhancement ("Keep Windfury Weapon up all fight").
- Retribution's Crusader opener (round 3's text can go in now).
- Marksmanship's Sniper Shot line.
- Subtlety Shank note ("on cooldown").
- Destruction :113 (Eureka covers Incinerate: contradicts the client).
- round4/shaman.md:74 ("engine bug": withdraw it).
- Tank stat priorities WARRIOR-2, DRUID-13 and PALADIN-6 (Expertise does exist on items).
- The low-rank wording above.

## Measure at refresh

- Every spec's raid and band BiS, because the engine changes under the gear.
- The cat, bear and Balance weights.
- Enhancement raid-60 weights (Rockbiter).
- Arcane (Fire spell power now counts).
- Destruction and Demonology band 20.
- The bear and Protection warrior hit and expertise cap figures (read on the published set now).

## Open, for round 16 (top first)

1. Hunter pets and player buffs: Blizzard's one sentence ("Pets can no longer receive player buffs that increase their stats") reads either way. Upper bound BM +26% if they all apply. A pet tooltip or log under Battle Shout / Leader of the Pack settles it.
2. Low-rank penalty parameters (above): threshold 16 vs 17, the top-rank exemption, and damage spells.
3. Skyborne base stats: cloned from Human/Orc. The client has none; a naked level-60 Skyborne profile import would give them.
4. Summon Hawk's 18 s assault: a client dummy with no data. A Forever combat log settles it.
5. Exorcism crit and Holy Shield crit on a level-60 log; Mutilate and Shank off-hand under live logs.
6. Aspect of the Hawk proc on ranged abilities (client) vs Auto Shot only (text).
7. Read Leyline lines for Skyborne healers and casters (modelled, cast by no rotation).
8. Death Mark storing white swings and poisons.
9. Cone of Cold's Winter's Chill adds its coefficient raw (no rotation casts a low rank today).
10. Cat form ±20% spread vs weapon spread; the unarmed 54.8 DPS fallback.
11. The data lane: items JSON drops expertise.
12. Coach: api/internal/coach TestWebFixtures_NoDrift and TestWebHistory_PagesAreTheRoutesShape were red before round 15.
13. Warcraft Logs stays blocked (gcloud login).
