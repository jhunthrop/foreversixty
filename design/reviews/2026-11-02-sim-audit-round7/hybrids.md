# Sim audit round 7: hybrids (2026-11-02)

This report covers Balance, Feral cat, Feral bear, Protection paladin, Retribution, Elemental and Enhancement. It also covers the damage side of the healer specs, where it touches these.

**Setup:**
- Client 1.60.1.70707. Published sets and builds are the 8baffe5fc nightly in `data/builds/1.60.1.70707/bis/`. The re-rank (bis run 41363895819) was cancelled by the owner, so item 3 is dropped. No ranker runs were made.
- Engine: pin b57765fb9 plus the fork branch `lane/r7-hybrids`.
- Site: `wt/r7-hybrids`.
- Client rows are the layered 70707 tables, the hotfix push included (copied from the casters lane's 2026-11-01 21:42Z fetch). SpellCategories comes from `scratchpad/wago/` (2026-10-31).
- Wowhead Forever tooltips were read on 2026-11-01. TrinityCore and cmangos were used for proc and hit-mask conventions.
- The fresh re-audit ran as three read-only sub-audits, one each for druid, paladin and shaman. I re-checked every row they called confirmed against the client rows before writing a test.

Fix-sha column: "E" is a fork commit on `lane/r7-hybrids`, "S" a site commit on `wt/r7-hybrids`.

**In short: 12 CONFIRMED-BUGs. 11 are fixed and 1 is routed.**
- 5 came from items 1 to 3: the Moonkin double count, Insect Swarm's crit, Judgement of Fury, the planner's equip stats, and held-in-off-hand items.
- 6 came from the fresh re-audit: Strength of Earth, Grace of Air, Mark of the Chosen, Iron Creed, the shared immunity cooldown, and Hammer of the Righteous.
- The routed one is a ranker bug that leaves the off hand empty for one-hand casters (X3). leveling-bis has no owner this round.
- The biggest cross-group move is **Strength of Earth Totem**. It gave Era's 61 Strength where Forever's rank is 42, so every melee spec on the raid preset falls about 1.2%.

## Item 1: side effects of round 6's hybrid changes

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Round 6's Moonkin crit fix (B1) was only half done, and Balance counted as a Moonkin without the talent.**<br>- Every Balance druid started in Moonkin Form, talent or not.<br>- Below 40, a druid without the talent still raised the raid's Moonkin Aura for itself (+3% crit), and rolled Omen of Clarity at the form's doubled chance and halved cooldown (24858 effects 4 and 5).<br>- With the talent, the form's own aura gave 3%, and the raid flag it raised gave 3% more. So every bare band had 6% where 24907 (aura 290, 3) gives 3% once.<br>- Now a druid without the talent starts in caster form, and the form no longer raises the raid flag on top of its own aura. | `druid.go:233` (`raidBuffs.MoonkinAura = true` when in form), `druid.New` (starting form), `balance.go:32`. Test first: `TestBalanceDruidsMoonkinAuraIsCountedOnce` (in form without the talent, and the raid flag raised; then neither). | CONFIRMED-BUG (client rows; the engine's own one-aura rule) | E 2c9c6de96 |
| S2 | **Insect Swarm could not crit outside Moonkin Form.**<br>- The swarm zeroed its crit chance outside the form at application, and since round 6's dynamic DoTs, on every tick. This is a Season of Discovery leftover.<br>- Forever's form (24858, 24907) is armor, Omen and a 3% aura; 24977 is a plain DoT (auras 3 and 54).<br>- It was latent while every Balance druid started as a Moonkin. S1 makes it live below 40. | `insect_swarm.go:106, 113`; owner ruling (non-channelled DoTs always crit). Test first: `TestInsectSwarmTicksCritOutsideMoonkinForm` (crit 0, then the caster's). | CONFIRMED-BUG (owner ruling plus client rows) | E 348385a66 |
| S3 | **Judgement of Fury still rolled the spell table.** Round 6 moved Righteousness and Command to the melee table. Fury, Protection's judgement, carries the same rows at all seven ranks (20183, 20411 to 20414, 1311650, 1311655): SpellCategories DefenseType 2 (row 76345 for 20414) and SpellMisc Attributes_0 0x250000, which includes 0x200000 (no active defense). It missed about 16% and crit for 1.5x. It now rolls the melee table without dodge, parry or block. | `seal_of_fury.go:194, 209`. Test first: `TestJudgementOfFuryRollsTheMeleeTableWithoutActiveDefenses` (magic and a 16% miss rate, then melee and under 9%). | CONFIRMED-BUG (client tables, the same rows as round 6 P1) | E 3ae4ef775 |
| S4 | Checked and aligned:<br>- Maul and Swipe at 1.75x; only Maul and Swipe carry it.<br>- Moonfire and Wrath registered once, in Balance and in Feral.<br>- Vengeance on Insect Swarm, and Nature's Reach on every spell.<br>- Natural Instinct on Rip: its mask_0 0x839800 also names Ferocious Bite, which shares Rip's bit, and the engine's mask matches.<br>- Righteous Fury mana-gain threat, and Shield Specialization's no-threat mana.<br>- Judgement of Righteousness and Command on the melee table: crits at 2x, Precision applies, Vengeance procs.<br>- Hammer of Wrath on the ranged table (DefenseType 3).<br>- Windfury Weapon takes Elemental Weapons once.<br>- Stormstrike's charge goes only to Lightning Bolt, Chain Lightning and Earth Shock (mask 0x100003).<br>- Flurry lasts 15 s.<br>- Moonfire and Flame Shock read spell power and multipliers each tick: their OnSnapshot is a plain Snapshot of a fixed base. | Sub-audits and file reads | Aligned | - |

## Item 2: ledger rows

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | **"Benediction's planner row lacks its healing line."** Benediction is the priest staff 18608, not the paladin talent.<br>- simdb fills a client item's on-equip stats from Wowhead Forever wherever the client row states zero (`pipeline/simdb/equip_fill.py`), because a hotfix item's ItemXItemEffect link never reaches the client tables.<br>- `items/<class>.json`, which the planner shows, never did. 159 client items there lack the spell damage and healing the sim scores. Examples: Benediction's 195/65, Staff of Jordan, Hand of Edward the Odd's 90 (the Balance and Elemental main hand).<br>- normalize now applies simdb's rule to the class lists: a missing stat is taken, a stated one is kept, and only client rows are touched. | 18608's client effects (23041 Call Anathema, 23236 +2% Holy crit, no healing); `raw/wowhead/items.json` `splheal` 195, `spldmg` 65. Test first: `test_a_client_item_takes_wowheads_equip_stats_where_its_row_has_none` (ImportError, then pass). The items json regenerates at the data run (not committed). | CONFIRMED-BUG (display; the sim already had the stats) | S 26b8573a3 |
| L2 | **"Raid-buffed doesn't add the class kit (Protection paladin's Devotion Aura at 60)."** The published number is not affected.<br>- spec-breakdown (`main.go:235-247`) and the ranker (`leveling-bis/character.go withKit`) both call `ResolvedPreset.Layer(KitBuffs, ...)`, so `devotion_aura` is on the raid preset too. `core/buffs.go:317` applies it as rank 10293 (735 armor) at 60.<br>- Forever has no Improved Devotion Aura talent, so the plain buff is right.<br>- The gap is only the /sim page's Raid-buffed set, which `web/src/lib/sim/settings.ts` builds. That is the page lane's file. | Code read | Not a sim bug; routed to the page lane | - |
| L3 | Blocked rows stay blocked, with no new evidence: Lacerate and Primal Bite's factor, Holy Shield's flat threat, Elemental Mastery, Windfury white or yellow, Righteous Fury's multiplication, and Arcanite Champion's proc rate. | | Blocked | - |

## Item 3: post-nightly sets

Dropped on the owner's update: bis run 41363895819 is cancelled and the workflows are disabled. The published sets are still the 8baffe5fc ones. Two findings came out of reading them (X2, X3).

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Strength of Earth and Grace of Air Totem had Era's values at every level.**<br>- Forever retuned both. Strength of Earth's auras are 7, 14, 25 and 42 Strength at levels 10, 24, 38 and 52, and 53 at 60 (8076, 8162, 8163, 10441, 25362). Grace of Air's are 49 and 77 Agility at 42 and 56, and 89 at 60 (8836, 10626, 25360). The shipped tables and the push agree, and Wowhead Forever's 10442 reads "by 42" and 10627 "by 77". The level-60 ranks are the Ahn'Qiraj books, as in Era.<br>- The engine gave 61 Strength and 67 Agility at every level.<br>- The raid preset's `strength_of_earth_totem` and Enhancement's own totem both read that one value. That is +19 Strength (before Kings) for every melee spec on the raid preset, and +54 at level 20 for a leveling Enhancement shaman.<br>- Both totems are now level-aware `core.BuffRanks` tables, with the books behind IncludeAQ. Grace of Air is in no preset. | `core/buffs.go:72, 102, 1331, 1360`. Test first: `TestEarthAndAirTotemsFollowTheClientRanks` (build failure, then pass). Site pin: `TestEarthAndAirTotemRanksMatchTheClient`, which needs the round-7 pin. | CONFIRMED-BUG (client rows plus Wowhead; fresh) | E d6c407a19, S ec47e8568 |
| X2 | **Held-in-off-hand items replaced the main-hand weapon.**<br>- Every orb, tome and rod (Item ClassID 4, SubclassID 0, InventoryType 23; 178 planner items) reached simdb as ItemTypeWeapon with no weapon or hand type.<br>- The engine builds equipment with `Equipment.EquipItem` (`core/database.go:216`). It ignores the request's slot and files an unknown-hand weapon under the main hand.<br>- Example: an Elemental shaman with Hand of Edward the Odd and Master Dragonslayer's Orb simmed at 436 DPS, against 468 with an empty off hand. With the orb typed as an off hand, it sims at 485. Balance: 528 with the orb, 567 empty, 586 with the fix.<br>- simdb now types held items WeaponTypeOffHand and HandTypeOffHand. | `pipeline/simdb/items.py _apply_class_subclass`. Test first: `test_a_held_in_off_hand_item_is_an_off_hand_not_a_main_hand` (weapon type 0, then 5). simdb.bin regenerates at the data run (not committed). | CONFIRMED-BUG (fresh; engine code plus measurement) | S 1bd3920b5 |
| X3 | **Route to the ranker owner (not fixed; leveling-bis is off limits):** Balance, Elemental and Shadow priest publish band 60 with an empty off hand (`empty_reason` "no_sourced_item", despite 259 sourced off-hand candidates for Elemental).<br>- The cause is in `verify.go`'s swap promotion. When a one-hander beats the picked two-hander in the main hand, the off hand stays empty: it was emptied for the two-hander, and the swap was measured that way.<br>- Before X2, any held item tried there would also have replaced the main hand. A shield added to the current Elemental set is +1.7% (Evergreen Shield 279262); a held item is about +4% (the orb, a later-phase item).<br>- The fix: re-pick the off hand after promoting a one-hander, then re-rank after X2's simdb. | `sim/cmd/leveling-bis/verify.go:384-408`; the bis json Coverage rows | CONFIRMED (ranker) | routed |
| X4 | **Mark of the Chosen** (17774) is 21 all stats with a 2-minute proc cooldown. The engine gave Era's 25 with no cooldown, so procs refreshed the buff (32.6% uptime on the horde bear). | SpellEffect 21970 (aura 29, misc -1, 21); SpellAuraOptions 132914 for 21969 (ProcChance 2, ProcCategoryRecovery 120000); Wowhead Forever "by 21 for 1 min (Proc chance: 2%, 2m cooldown)". `common/item_effects.go:2951`. Test first: `TestMarkOfTheChosenIsTwentyOneStatsOnATwoMinuteCooldown` (25 stats and 33 procs, then 21 and 1). | CONFIRMED-BUG (client plus Wowhead; fresh) | E 1b876c7f2 |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | Moonkin form and aura (S1), and Insect Swarm's crit (S2). | | CONFIRMED-BUG | E 2c9c6de96, E 348385a66 |
| B2 | Moonglow and Improved Wrath's cost cuts add (Wrath is 30 mana a cast). Carried from round 6 B4. | SpellPower 313696 | OPEN (a logged Wrath cost settles it; 0 s out of mana) | - |
| B3 | Starfire rank 7 (25298) is in our trainables at 60. IncludeAQ keeps it out, as an Ahn'Qiraj book rank. | SkillLineAbility 13235 | OPEN (whether Forever's trainers teach it; about +2 to 3%) | - |
| B4 | Checked and aligned:<br>- Spells: Wrath 9912, Moonfire 9835, Starfire 9876, Insect Swarm 24977.<br>- Omen of Clarity and Clearcasting (10 s proc cooldown).<br>- Talents (against the trait curves): Genesis, Nature's Majesty, Vengeance, Moonfury, Improved Moonfire, Nature's Reach, Nature's Grace, Eclipse, Naturalist.<br>- Gear: Idol of the Moon, Frozen Heart of the Mountain, Hand of Edward the Odd (its proc needs a melee hit).<br>- Race: Tauren Endurance. | Sub-audit, client rows | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Strength of Earth's raid value (X1): -1.2%. | | CONFIRMED-BUG | E d6c407a19 |
| F2 | Ferocious Bite casts rank 4 (22829). Rank 5 (31018, 2.7 an energy) is in our trainables at 60 and held back as an Ahn'Qiraj book, the same question as B3. | `ferocious_bite.go:44`; SkillLineAbility 14491 | OPEN (under 0.1%) | - |
| F3 | Checked and aligned:<br>- Spells: Shred, Ravage, Rip (12 s, 2 s ticks), Ferocious Bite's energy conversion, Shifting Power (55% base mana, 40 energy, 8 s).<br>- Talents: Natural Instinct's mask, Idol of the Dream (+1 Rip tick), Savage Fury, Sharpened Claws, Predatory Strikes, Blood Frenzy, Leader of the Pack.<br>- Tiger's Fury and Faerie Fire (Feral) are not trainable in 70707. | Sub-audit | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | Mark of the Chosen (X4, the horde trinket) and Strength of Earth (X1). | | CONFIRMED-BUG | E 1b876c7f2, E d6c407a19 |
| G2 | The engine casts Barkskin in Dire Bear Form. No SpellShapeshift row is fetched, and Wowhead shows no form line. | `barkskin.go:12` | OPEN (SpellShapeshift 22812, or a cast in bear on beta) | - |
| G3 | The APL casts Primal Bite 1238073, which has no SkillLineAbility row. 1238070 is the learnable rank, with identical numbers. | SkillLineAbility 53142 | Aligned (cosmetic id) | - |
| G4 | Checked and aligned: Maul rank 7 (128, 15 rage), Swipe rank 5 (83, 20 rage, can be dodged and parried), Lacerate 1235827 (15 a tick x 5, 5 stacks), Shredding Attacks, Thick Hide, Heart of the Wild, Natural Reaction, Feral Swiftness. | Sub-audit | Aligned | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Judgement of Fury on the melee table (S3): +1.4 to 1.5% threat. | | CONFIRMED-BUG | E 3ae4ef775 |
| T2 | **Iron Creed's -6% damage taken opened on missed, dodged and parried Holy Strikes.** 1311034's proc is aura 231 on a done melee-class spell (ProcTypeMask 16) with no hit mask, which TrinityCore (the default done hit mask) and cmangos (hit or crit only) both read as a landed hit or crit. About a third of strikes fail from the front. | `holy_strike.go:168`. Test first: `TestIronCreedOpensOnlyOnALandedHolyStrike` (opened 2000 of 2000, 1341 landed; then opened only on landed strikes). | CONFIRMED-BUG (two server cores) | E fa557aac7 |
| T3 | **Divine Protection and Divine Shield had separate cooldowns.** All four ids (498, 5573, 642, 1020) are SpellCategories Category 37 with a 300 s CategoryRecoveryTime. Templar's Bulwark is Category 0. | `defensive_cooldowns.go:81`. Test first: `TestDivineProtectionAndDivineShieldShareOneCooldown`. | CONFIRMED-BUG (client category; about 0, they fire about once in 100 fights) | E 35f090ff9 |
| T4 | **Hammer of the Righteous struck 3 targets; the client strikes 4.** 407632 effects 0 and 2 have EffectChainTargets 4; Wowhead Forever: "the current target and up to 3 additional nearby targets". | `hammer_of_the_righteous.go:31`. Test first: `TestHammerOfTheRighteousStrikesFourTargets`. | CONFIRMED-BUG (single target unchanged) | E c4d72f960 |
| T5 | Holy Shield's damage can crit in the engine. The client's damage is on the aura (aura 43): TrinityCore's aura 43 handler does not crit, while cmangos rolls a crit. | 20928 effect 1 | OPEN (a launch log; about -2.4% threat if it cannot crit) | - |
| T6 | Toughness has a second effect, aura 466 (bonus armor 2 to 10%), on top of item armor. The engine scales item armor only. | 20143 effect 1 | OPEN (a character-sheet reading; under 0.2% damage taken) | - |
| T7 | Holy Strike's flat damage under the 50% weapon multiplier (carried from round 6 P8). The tooltip says "plus an additional 97 to 121". | Wowhead 10333 | OPEN (carried; +2% threat and +0.9% Retribution if it is added on top) | - |
| T8 | Checked and aligned:<br>- Seal of Fury rank 7 and its proc 20418, and Judgement of Fury's amount (153, coefficient 0.429).<br>- Judgement and Swift Judgement.<br>- Hammer of the Righteous's cost and cooldown; Holy Strike with Shieldsworn; Holy Shield's mana and charges.<br>- Templar's Bulwark.<br>- Talents: Sacred Duty, Redoubt, Shield Specialization, Iron Creed's threat, Anticipation, Precision, Conviction, Divine Strength, Improved Seals.<br>- Gear: Deathbone 2-piece, Guardian Talisman (no combat effect).<br>- Race: Touch of the Grave. | Sub-audit | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Strength of Earth's raid value (X1): -1.2%. | | CONFIRMED-BUG | E d6c407a19 |
| P2 | Checked and aligned, nothing new:<br>- Seal of Command rank 5 through 20424 (70% weapon damage, 0.29, 1 s cooldown, can be dodged).<br>- Judgement of Command and Righteousness on the melee table; Hammer of Wrath on the ranged table (6 s, 425 mana before Holy Conduit).<br>- Talents: Improved Judgement, Holy Conduit, Benediction, Sanctified Judgement, Vengeance, Two-Handed Weapon Specialization, Champion of the Light.<br>- Gear: Molten Heart, Cloak of Flames (0 DPS).<br>- Carried: Arcanite Champion (blocked), 25713's base points, Holy Strike (T7). | Sub-audit | Aligned | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | The off hand is empty at band 60 (X3), and held items replaced the main hand (X2). | | CONFIRMED (ranker routed; simdb fixed) | S 1bd3920b5 |
| E2 | Searing Totem attacks every 2.5 s. Its Attack (10436) is a 2.2 s cast; the engine stretched it to stand in for missile flight. | `fire_totems.go:61`; SpellCastTimes 2200 | OPEN (a beta log of the totem's attack times; about +0.8%) | - |
| E3 | Fire Nova assumes the boss is within 10 yards of the totem, which a ranged shaman drops at its feet. | 408428 radius | OPEN (positioning; Fire Nova is 2.9% of DPS) | - |
| E4 | Checked and aligned:<br>- Spells: Lightning Bolt 15208, Lava Burst 1326047 (+20% on Flame Shock at 60), Flame Shock 29228, Earth Shock 10414, Searing Totem.<br>- Talents: Convection, Concussion, Elemental Devastation, Call of Flame, Elemental Focus, Lightning Overload, Elemental Fury, Thundering Strikes, Ancestral Knowledge, Tidal Focus.<br>- Gear: Totem of Thunder.<br>- Race: Blood Fury.<br>- Carried: Elemental Mastery (blocked), Orc Command on totems. | Sub-audit | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | Strength of Earth (X1): -1.3% raid. Bare: -11% at band 20, -8% at 30, -5% at 40, -4% at 50, -1.6% at 60. The APL's own totem gave 61 at every level. | | CONFIRMED-BUG | E d6c407a19 |
| N2 | The APL still drops its own Windfury Totem, which a Windfury Weapon shaman cannot use. Round 6 N5; +0.1%, within error. | | Rotation note | - |
| N3 | The band-60 main hand is a two-hander (Forest Defender's Axe), so no shield shows. The ranker's off-hand pool for this spec is shields and held items (owner rule), and is right. | bis json | Aligned | - |
| N4 | Checked and aligned:<br>- Talents: Reverberation, Shamanistic Focus, Mental Dexterity and Quickness, Elemental Weapons, Flurry, Stormstrike, Maelstrom Weapon, Rage of the Farseer.<br>- Windfury Weapon (333 attack power, 20%, 1.5 s cooldown), Windfury Totem, Ironfeather 2-piece.<br>- Carried: Windfury white or yellow (blocked). | Sub-audit | Aligned | - |

## Before/after

Pin b57765fb9 on origin/main against `lane/r7-hybrids` at d6c407a19 with `wt/r7-hybrids`. Raid preset, band 60, 10,000 iterations. The simdb is today's committed one, so X2's simdb change is not in these numbers. Tanks are measured in threat a second.

| spec | changes | alliance % (seed 1 / 50001) | horde % (seed 1 / 50001) | absolute, seed 1 (alliance; horde) |
|---|---|---|---|---|
| Balance | S1, S2 (bare bands only) | 0.0 / 0.0 | 0.0 / 0.0 | 565.7; 567.2 |
| Feral cat | X1 | -1.2 / -1.2 | -1.2 / -1.2 | 624.1 → 616.5; 619.6 → 612.0 |
| Feral bear (threat) | X1, X4 (horde) | -1.4 / -1.4 | -1.9 / -1.9 | 877.0 → 864.6; 885.3 → 868.6 |
| Protection paladin (threat) | S3 +1.4/+1.5, X1 -1.0, T2 (damage taken +1.3/+1.4) | +0.5 / +0.4 | +0.5 / +0.5 | 992.7 → 997.5; 1008.2 → 1013.3 |
| Retribution | X1 | -1.2 / -1.2 | -1.2 / -1.2 | 666.2 → 658.3; 664.5 → 656.7 |
| Elemental | - | 0.0 / 0.0 | 0.0 / 0.0 | 459.4; 467.6 |
| Enhancement | X1 | -1.3 / -1.2 | -1.3 / -1.3 | 608.0 → 600.4; 607.3 → 599.7 |

Bare bands, horde, seed 1 (seed 50001 within 0.1):
- **Balance:** 20: -5.2%; 30: -5.5%; 40: -2.5%; 50: -2.4%; 60: -2.4% (S1, S2).
- **Enhancement:** 20: -11.1%; 30: -8.0%; 40: -4.6%; 50: -3.8%; 60: -1.6% (X1).

**Outside the group, for the integrator to re-measure:**
- **E d6c407a19 (Strength of Earth).** Every spec on the raid preset that uses Strength: warriors, rogues, hunters (melee attack power only, so Survival most) and Protection warrior threat. The fork goldens that moved: TestP1DPSWarrior, TestP1TankWarrior, TestP1Mage.
- **E 1b876c7f2 (Mark of the Chosen).** Any set that wears it and is hit. It is 18 specs' candidate at some band, but it procs only on hits taken, so it matters for tanks.
- **S 1bd3920b5 (held items).** Every caster that can hold an off-hand item, after the data run and a re-rank: mages, warlocks, priests and the hybrid casters.
- **S 26b8573a3 (the planner's items).** Display only.

**Not done:**
- Nothing nightly-owned was committed.
- `TestRotationLadder` was not regenerated.
- No ranker run was made (the bis item was dropped).

**Fork goldens regenerated:** TestProtection (S3, T2, T4), and with X1: TestFeralTank, TestP1Mage, TestProtection, TestElemental, TestEnhancement, TestP1DPSWarrior and TestP1TankWarrior. The full fork suite passes with `-tags=with_db`; `sim/web` needs `binary_dist`, as before.

## Still open

1. **X3, routed to the ranker owner.** Verify's main-hand promotion leaves the off hand empty: Balance, Elemental and Shadow priest at band 60. Re-rank after X2's simdb.
2. **The Ahn'Qiraj book ranks Forever lists as trainable at 60:** Starfire 7 (B3), Ferocious Bite 5 (F2), and Strength of Earth 5 and Grace of Air 3 (X1). IncludeAQ keeps all four out. Whether Forever's trainers teach them settles it.
3. **T5** Holy Shield crit. **T6** Toughness's bonus armor. **T7** Holy Strike's flat damage.
4. **G2** Barkskin in bear. **E2** Searing Totem's 2.2 s attack. **E3** Fire Nova range. **B2** Wrath's cost stacking.
5. **L2** for the page lane: Raid-buffed on /sim adds no class kit, so Protection paladin's Devotion Aura is missing there. The published numbers have it.
6. Blocked rows unchanged: Lacerate and Primal Bite, Holy Shield's flat threat, Elemental Mastery, Windfury, Righteous Fury, and Arcanite Champion.
