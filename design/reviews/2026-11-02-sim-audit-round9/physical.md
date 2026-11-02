# Sim audit, round 9: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), plus hunter pets, on client 1.60.1.70707.

- **Baseline:** fork pin 40021374a with site origin/main 55a1212f7, each with its own `make simdb`.
- **Branches:** fork fixes are on `lane/r9-physical` (E), site changes on `wt/r9-physical` (S).
- **Measurements:** `sim/cmd/spec-breakdown`, 10,000 iterations, seeds 1 and 50001, both factions, on the published sets.
- **Client rows:** the layered 70707 raw tables with the hotfix push (SpellEffect, SpellCategories, SpellClassOptions, SpellLevels, SpellAuraOptions, SpellCooldowns, SpellPower, Spell, SpellItemEnchantment, ItemEffect, ItemXItemEffect), plus wago's SpellEquippedItems for 70707, read 2026-11-02. Forever tooltips are `https://nether.wowhead.com/forever/tooltip/spell/<id>` (or `/item/<id>`), read 2026-11-02; a tooltip and its client row count as one source.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Marksmanship's held build stays held.** The re-run search on the round-8 engine finds the same winner: two points of Improved Arcane Shot, Deflection and Surefooted, paid for by Endurance Training, Unleashed Fury and a point of Ferocity. At 10,000 iterations it is +0.8% alliance and +0.9% horde, both seeds (649.2 to 654.6; 653.1 to 659.1). That is under the 1% bar. | `scratchpad/r9/physical/ts-mm-raid/hunter-marksmanship-raid.md`; `mm-hold.txt` | HOLD (under 1%) | - |
| A2 | **Assassination's leveling build adopts one change.** The search was re-run with Murder kept, since it is an owner raid keep. The best variant that keeps every unmodeled guide talent moves the single Venom point to Seal Fate. Venom is a finisher the rotation never casts, so that point did nothing.<br>Measured, bare, 10,000 iterations, both seeds:<br>- band 60: +1.6% (216.2 to 219.7)<br>- band 50: +1.1 to +1.3%<br>- band 40: +0.5%<br>The raid build has no Venom and does not move.<br>The rule holds: the gain is beyond error and over 1% at the search's level 60. The Assassination tree keeps its 32 points. Nothing the rotation casts is dropped. The capstone it gives up is one the guide already says the rotation should not cast.<br>The guide's two talent sentences now say this. `addon-builds.json` is regenerated from the guide. Round 7's +3.6% variant also dropped Murder, so it stays rejected. | `ts-assn-bare/rogue-assassination.md`; `assn-venom.txt` | adopted (rule passes) | S 39bcdf481 |
| A3 | **Curse of Recklessness is re-verified.** Ranks 704, 7658, 7659 and 11717 are aura 674 at -175, -285, -395 and -505 armor, learned at 14, 28, 42 and 56. Each sits beside aura 92 (no flee), with no attack power effect. The tooltip reads "reducing the target's armor by 505". The engine's `CurseOfRecklessnessRanks` match, and `TestCurseOfRecklessnessTakesTheRanksArmorOnly` holds them. | SpellEffect rows above; `sim/core/raid_debuff_ranks.go` | aligned | - |
| A4 | **The 26 client weapon procs (30 rows now) are re-verified.** A script follows each item to its trigger-2 proc spell (ItemXItemEffect, then ItemEffect), and that spell to its damage effect. All 30 rows in `clientWeaponProcs` match the layered SpellEffect base points and variance. Spell and item link match for every row. | `scratchpad/r9/physical/procs.py` output | aligned | - |
| A5 | **Expose Armor used Era's armor values.** Forever's Expose Armor takes Sunder Armor's amounts per combo point. Ranks 1 to 5 (8647, 8649, 8650, 11197, 11198, learned at 14, 26, 36, 46 and 56) remove 90, 180, 270, 360 and 450 per point. A five-point rank 5 is 2250. The engine had Era's 80 to 340 per point (1700 at five).<br>Forever's Improved Expose Armor (14168) now reads "Reduces the Energy cost ... by 10, and refunds 2 Combo Points when cast with 5 Combo Points". It no longer adds armor. Even so, the engine's raid toggle added 50% for its "improved" setting.<br>The raid toggle is now a five-point cast of the raid's rank, either way (`ExposeArmorRanks`). The rogue's own cast takes the same table.<br>No published number moves: no curated rotation casts Expose Armor, and the site exposes no Expose Armor toggle. | Wowhead Forever 8647 to 11198 ("1 point: 450 Armor ... 5 points: 2250 Armor"), 14168. The client rows carry aura 674 at 0 (server-computed), so the tooltip is the single source for the amounts. Era 11198: 340 a point. Tests `TestExposeArmorIsForevers` and `TestExposeArmorPerComboPointIsForevers` failed first. | CONFIRMED-BUG (single source for the amount; 0 published DPS) | E 9c28bcb51 |
| A6 | **Armor-debuff stacking is still open.** The engine applies each debuff once: Curse of Recklessness is one shared aura; Faerie Fire and Faerie Fire (Feral) share the "minor" slot; Sunder Armor and Expose Armor share the exclusive "major" slot. That is Era's rule. Forever moved all four to aura 674, and the client cannot say whether the server groups them differently. The aggregate Forever Logs API has no event data, so a log cannot settle it from here. | debuffs.go 536-640; SpellEffect 11597, 9907, 11717, 11198 | OPEN (needs a combat log with damage per hit under each debuff mix) | - |
| A7 | **The crit sources for every physical spec are mapped.** The full table went to the crit lane early (`scratchpad/r9/physical/crit-notes.md`). In short:<br>- Forever turned many "crit" auras into aura 290 (all crit: melee, ranged and spell). These are Leader of the Pack, Moonkin Aura, Elixir of the Mongoose, Grilled Squid, Malice, Lethal Attacks, Berserker Stance, Recklessness, Ruthless and Elune's Light. The engine's single crit stat is right for all of them.<br>- The ones that must NOT reach spells are Elemental Sharpening Stone (aura 52 through its enchant), Cruelty (aura 52 limited to melee weapons), gear "Increased Critical" (aura 52) and Agility.<br>- Spell-only gear crit (aura 57) must not reach melee.<br>- Rogue poisons and every weapon proc are DefenseType 1 (magic), so they should roll spell crit. Today they crit at 31 to 40%, the melee rate.<br>- Arcane Shot and Serpent Sting are DefenseType 3 (ranged) and stay on the physical table.<br>- Forever's world buffs (Rallying Cry, Songflower, Slip'kik's Savvy) are server dummies. Their tooltips say all crit, so the engine is right on crit.<br>- Hunter pets copy the owner's single crit stat.<br>The lane changed no crit plumbing. | SpellEffect aura types; SpellEquippedItems (12320: class 2, subclass mask 173555); SpellCategories DefenseType; tooltips 22888, 15366, 22820 | sent to the crit lane | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Pendulum of Doom, Blackhand Doomsaw and Vis'kag the Bloodletter ignored their execute bonus.** The client's own text for Fatal Wound 10373, 16549 and 1305394 reads "Deals $s2% increased damage to targets below $s3% health". Wowhead Forever renders it as 50%, 50% and 75% below 25%. Round 8 modelled the same clause on Gut Ripper (18107, 100%). The four now share one table. None of the three is a published pick; Pendulum of Doom is a ranked Enhancement alternative. | Spell.csv Description_lang for the four spells; tooltips 10373, 16549, 1305394, 18107; test `TestFatalWoundsHitHarderBelowAQuarter` failed first | CONFIRMED-BUG (0 published DPS) | E b57ebe9f0 |
| X2 | Other clauses in the weapon-proc texts were checked. Searing Needle's Fire vulnerability is mage-only. Verigan's Fist's elemental multiplier and the frost bolts' slow never reach the type-neutral target. | Spell.csv for the 30 proc spells | aligned | - |
| X3 | Side effects of round 8 were re-checked on the nine specs:<br>- Berserker Stance at 30 and Whirlwind at 36 hold. W4 below revisits Fury's stance.<br>- Mutilate rank 1 is at 30.<br>- Strength of Earth is 42.<br>- Smoked Desert Dumplings are the strength food and Grilled Squid is 1% crit (aura 290).<br>- Leader of the Pack and Moonkin Aura are both aura 290 at 3, in reach at their talent levels.<br>- Raid debuffs are by the raid's level.<br>- Manual Crowd Pummeler and Gut Ripper's execute bonus are unchanged. | breakdowns `scratchpad/r9/physical/bd/`; client rows above | aligned | - |
| X4 | Strength of the Champion (Arcanite Champion) shows about 10 procs a fight for Arms. That is one a minute from every main-hand hit, including Hamstring weaving, which Classic allows. The rate itself is the engine's assumption, as in round 8. | Arms band-60 raid breakdown | OPEN (carried: no rate in the client) | - |

## Warrior: Arms and Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **Fury from 30 to 35 opened in Berserker Stance with nothing to use there.** Berserker Stance is learned at 30, but Whirlwind comes at 36 and Bloodthirst is the 31-point talent. So the stance gave a level-30 Fury warrior 3% crit and took away Overpower, its one special besides Heroic Strike.<br>Fury now opens in Battle Stance until it knows Whirlwind. Measured at band 30 bare, horde, both seeds: +3.9% (61.8 to 64.2). Forcing Battle Stance at bands 40 to 60 loses 4 to 12%, so they keep Berserker Stance.<br>This settles round 8's open W4. The APL's Overpower note says so, and the fork's copy is synced. TestP1DPSWarrior's Fury rows (below 36) move with it. | SpellLevels 2458 (30), 1680 (36); measured with a prepull Battle Stance (`fury-battle.json`); test `TestFuryOpensInBattleStanceUntilWhirlwind` failed first | CONFIRMED-BUG (rotation) | E 21dca29f9; S 105a8463b, d77242465 |
| W2 | Abilities checked against the client and the Forever tooltip, all aligned:<br>- **Bloodthirst** 23894: 45% AP + 48, rank 1 to 4 bases 30 to 48, 30 rage, 6 s.<br>- **Whirlwind** 1680: both weapons, 25 rage, 10 s, level 36.<br>- **Mortal Strike**: +85, +110, +135 and +160 by rank, 30 rage, 6 s.<br>- **Slam** 11605: +87, 1.5 s cast, 18 s, 15 rage. Improved Slam is -0.25 s a point and -3 s cooldown (12862).<br>- **Overpower** 11585: +35, 5 rage, 5 s.<br>- **Execute** 20662: 600 + 15 a rage, 15 rage.<br>- **Heroic Strike** 11567: +138, 15 rage.<br>- **Hamstring** 7373: 45, 10 rage.<br>- **Rend** 11574: 21 a tick.<br>- **Battle Shout** 11551: 111 + 0.6 a level, 116 at 60; pinned against the CSV.<br>- **Charge**: 9, 12 and 15 rage; Improved Charge +3 a point; 6 rage at 2/2 shows in the Fury breakdown.<br>- **Death Wish** 12328: +20%, 5% damage taken, 3 min.<br>- **Recklessness** 1719: aura 290 at 100.<br>- **Berserker Stance passive** 7381: aura 290 at 3.<br>- **Molten Fury** 1249113: 55 AP and 55 RAP. | SpellEffect, SpellCooldowns, SpellPower; tooltips 21553, 23894, 1680, 11605 | aligned | - |
| W3 | Talents the published builds take, against the talent search's own probe (each probe shares the base run's seed):<br>- **Moving damage:** Improved Tactical Mastery, Improved Overpower, Anger Management, Deep Wounds, Two-Handed Weapon Specialization, Impale, Bloodthrill, Weaponmaster, Improved Slam, Mortal Strike, Cruelty, Unbridled Wrath, Improved Bloodrage (Arms); and Improved Heroic Strike, Improved Charge, Furious Precision, Dual Wield Specialization, Improved Execute, Death Wish, Flurry and Bloodthirst (Fury).<br>- **No damage, correctly:** Deflection, Improved Rend (Fury casts no Rend), Sweeping Strikes (single target), Blood Craze and Enrage (the boss hits the tank).<br>- **Unmodeled and utility only:** Lingering Rage and Improved Intercept.<br>- **Search result:** Arms' guide is the best found. Fury's best is +0.9%, under the bar. | `ts-warrior-arms/`, `ts-warrior-fury/` | aligned | - |
| W4 | Unbridled Wrath also procs on Heroic Strike. It credits about 18 rage a fight more than white swings alone, because the engine treats Heroic Strike as a swing. The client's proc type mask is 4 (melee auto attack). This is round 8's carried "Heroic Strike as a swing" row. | Fury breakdown: 103.5 rage against about 85.6 from white hits | OPEN (carried; a log of UW procs) | - |
| W5 | Normalized rage: the off-hand factor (0.5) and crit multiplier (2.0) are still unverified. The Forever Logs API is aggregate-only, so it cannot settle them. | `sim/core/rage.go` | OPEN (carried) | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Checked against the client and the tooltip:<br>- **Shield Slam** 23925: 640 to 670 plus block value, 20 rage, 6 s.<br>- **Revenge** 11601: 109 to 133 + 25% AP, 5 rage, 5 s.<br>- **Thunder Clap** 11581: 103 damage, 20% slower attacks, 20 rage, 6 s.<br>- **Shield Block** 2565: +75% block for 2 blocks, 5 s.<br>- **Sunder Armor**: 450 a stack.<br>- Block value is 1 per 20 Strength.<br>- Concussion Blow is taken but never cast; it is a stun.<br>Threat values stay round 7's carried row. | SpellEffect, tooltips 23925, 11601, 11581, 2565 | aligned | - |

## Rogue: Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Subtlety never pooled energy for Ghostly Strike.** The Ghostly Strike line says "on cooldown". But Hemorrhage (35 energy) went off the moment energy reached 35, so energy never reached Ghostly Strike's 40 (14278: 40 energy, 20 s). It was cast 1.4 times a fight instead of about 9.<br>Hemorrhage now holds back while Ghostly Strike is under a second from ready, unless energy is at 75 or more. Ghostly Strike rises to about 8.7 a fight.<br>Measured: +0.1 to +0.2% on the band-60 raid preset (both factions, both seeds), and level to +0.2% bare. The fork's copy is synced. | breakdowns; `rogue-audit/sub-pool.json`; test `TestSubtletyHemorrhageWaitsForGhostlyStrike` failed first | CONFIRMED-BUG (rotation; about +0.15%) | S 431c6006e, d77242465; E effe86939 |
| R2 | Checked against client rows, rank curves and the Forever tooltips, all aligned:<br>- **Poisons:** proc chances 20% (Instant) and 30% (Deadly, Wound); Improved Poisons +2 a point; Vile Poisons +4% a point. Instant Poison VI is 88 ±27.7%. Deadly Poison IV ticks 18 every 3 s, 5 stacks. Venom is +30% damage and +10 points of proc chance.<br>- **Builders:** Backstab, Sinister Strike and Ambush use normalized weapon damage, +140 at 150%, +68 and +116 at 250%. Mutilate is (weapon + 67) × 75% per hand, 60 energy, 2 combo points, +20% against a poisoned target. Hemorrhage is 100% (145% with a dagger), with the 15% Rupture debuff. Ghostly Strike is 125% (180%).<br>- **Finishers:** Eviscerate rank 8 is 48 to 144 + 151 a point. Rupture rank 6 is 35 + 4.73 a point a tick. Slice and Dice is 6 + 3 s a point, 20% or 30% haste.<br>- **Talents:** Ruthlessness, Relentless Strikes, Seal Fate (500 ms internal cooldown), Initiative, Thousand Cuts, Lethality, Puncturing Wounds, Improved Ambush, Cold Blood, Malice, Hack and Slash, Improved Eviscerate, Aggression, Opportunity, Serrated Blades, Quietus, Murder, Dual Wield Specialization, Flawless Execution, Improved Sinister Strike, Precision, Weapon Expertise.<br>- **Cooldowns:** Adrenaline Rush, Blade Flurry, Thistle Tea, Premeditation.<br>- **Ranks per band** match what each band's sim casts.<br>- **Sets:** Stormshroud 2-piece and 3-piece, and Devilsaur.<br>- **Racial:** troll Berserking, 10% for 15 s. | rogue fork of this lane; `rogue-audit/` | aligned | - |
| R3 | Rogue poisons crit at the melee rate, 31 to 40%. They are magic (DefenseType 1). The crit lane owns the split (A7). | breakdowns | for the crit lane | - |
| R4 | Carried: the attack power share of Eviscerate and Rupture (the client's BonusCoefficientFromAP is 0, and the tooltips say only "increased by Attack Power"). Also carried: Combat's Slice and Dice gate (round 2 C2), and Venom adding with Vile Poisons rather than multiplying (the guide already says the model flatters Venom). | | OPEN (carried) | - |

## Hunter: Beast Mastery, Marksmanship, Survival, pets

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Savage Strikes reached only Raptor Strike and Mongoose Bite.** Savage Strikes (19159) is aura 107 misc 7 (crit chance, 2% a point), with class mask 66 / 528384 / 2097152. Each ability's own SpellClassOptions row puts these inside that mask:<br>- Raptor Strike and Mongoose Bite (mask 0 bit 2)<br>- Wing Clip (bit 64)<br>- Counterattack (mask 1 bit 524288)<br>- Lacerate (bit 4096)<br>- Strider Kick (mask 2 bit 2097152)<br>The tooltip reads "all your melee abilities". Survival's Strider Kick, Wing Clip, Lacerate and Counterattack now get the 4% too. | SpellClassOptions 14266, 14271, 14268, 20910, 1299332, 1317257; test `TestSavageStrikesReachesEveryMeleeAbility` failed first | CONFIRMED-BUG | E 63686fa7e |
| H2 | **No hunter's cat ever bit.** The cat cast Bite only while its focus regen beat Claw's cost / 1.6, which is 15.6 focus a second. A pet regenerates 26.25 focus every 5.25 s: 5 a second, or 6 with Bestial Discipline. So no breakdown had a Bite.<br>Bite (17261: 35 focus, 10 s, 81 to 99) deals 2.6 damage a focus, against Claw's 2.0 (3009: 25 focus, 43 to 59). The cat now bites whenever Bite is ready, pooling focus for it, and claws otherwise.<br>Bite's rank table also takes the client's rolls. Ranks 1 to 3 and 5 had used rank 4's 31 to 37, and rank 8 stopped at 91 instead of 99. | `sim/hunter/pet.go:267`, `core/focus.go:12-13`; SpellEffect 17253 to 17261; tests `TestCatBitesWheneverBiteIsReady` and `TestPetBiteRanksRollTheClientRange` failed first | CONFIRMED-BUG | E 2cd085dce |
| H3 | **Every hunter below 60 carried level-52 arrows and a level-55 quiver.** The site gave every band Thorium Headed Arrows (17.5 DPS, Requires Level 52) and a 15% quiver (Gnoll Skin Bandolier, level 55).<br>A hunter now carries the best its level can use:<br>- Ammo: Sharp Arrow 3.5 DPS from 10, Razor Arrow 7.5 from 25, Jagged Arrow 13 from 40, Mithril Gyro-Shot 15 from 44, Ice Threaded Arrow 16.5 at 51, Thorium from 52.<br>- Quiver: Light Quiver 10%, Heavy Quiver 12% from 30, Quickdraw Quiver 13% from 40, the Bandolier from 55.<br>The engine's ammo list had nothing below Razor Arrow, so it gains SharpArrow and HeavyShot (a proto enum change). The engine does not match ammo to bow or gun, so from 44 to 50 the Gyro-Shot stands in for a bow's Jagged Arrow. The published band-50 sets wear guns. Band 60 does not move. | Wowhead Forever items 2515, 3030, 11285, 10513, 18042, 2101, 7371, 8217, 19320; tests `TestAmmoDamagePerSecondIsTheClients` (fork) and `TestAHuntersAmmoAndQuiverFollowItsLevel` (site) failed first | CONFIRMED-BUG | E 2a8b74fff; S 9c7de3c53 |
| H4 | **Beast Mastery's Serpent Sting clipped a tick on most casts.** The rotation refreshed the sting with under 5 s left, but a refresh drops the remaining ticks. So 16 casts gave 55 ticks in a 180 s fight.<br>Refreshing under 1.5 s left keeps every tick but at most the last. Measured: +0.2 to +0.3% raid (both factions, both seeds) and +0.1 to +1.0% bare. The fork's copy is synced. | `bm-ss15.json` measurements; test `TestBeastMasteryRefreshesSerpentStingWithoutClippingATick` failed first | CONFIRMED-BUG (rotation) | S e164c7c0a, d77242465; E 629e104be |
| H5 | **Mortal Shots' crit damage skipped Serpent Sting.** Mortal Shots (19485, aura 108 misc 15: crit damage bonus) has mask 0 = 424449, which includes bit 16384, Serpent Sting's own class mask. Forever's DoTs crit (owner ruling), so the sting's tick crits take the bonus, as Arcane Shot's do. The effect is small. | SpellEffect 19485, SpellClassOptions 13555; test `TestMortalShotsReachesSerpentSting` failed first | CONFIRMED-BUG | E 44d0f1707 |
| H6 | **A pet's base stats and weapon damage were its level-60 values at every band** (Str 136, Agi 100, 18.17 to 27.66 a second of swing). A hunter's pet is the hunter's level, so a band-20 cat hit like a level-60 one. The pet is about 30% of Beast Mastery's damage at the low bands.<br>**Source.** The 70707 client has no pet stat table: raw/ carries no CreatureFamily stats, and Hunter Pet Scaling 415429's channels are server-computed (bp 0). Forever is taken to keep Era's pet system. Era's per-level sheet is pet_levelstats entry 1, identical row for row in mangoszero's and tortoise-wow's world databases (level 30: Str 61, Agi 52, Sta 115, Int 32, Spi 44; level 60: 136, 100, 256, 50, 80, which is the engine's level-60 sheet except Stamina 274). The two share a lineage, so treat them as one Era source. Era's server code is the second source for damage: vmangos' Pet::InitStatsForLevel makes it proportional to the pet's level (petlevel × 1.2075 to 1.5225 a 2 s swing, "from vanilla pet tab screenshots"), and cmangos' is linear in petlevel.<br>**Fix.** The pet takes the table's stats at its level; Stamina keeps the engine's 274 at 60 and follows the table's curve below. Weapon damage scales by level / 60. Level 60 does not move. TrinityCore has no 1.x database to check. | mangoszero `World/Setup/FullDB/pet_levelstats.sql`; tortoise-wow `sql/base/tw_world_pet_levelstats.sql`; vmangos `src/game/Objects/Pet.cpp:1411-1414`; cmangos `src/game/Entities/Pet.cpp` (all read 2026-11-02); test `TestAHunterPetsStatsAndDamageFollowItsLevel` failed first | CONFIRMED-BUG (Era source; no Forever table) | E ac0ae9ae6 |
| H6b | **Warlock demons do not have the level-60 problem.** Each demon carries sheets at 25, 40, 50 and 60, interpolated between them, and extrapolated below 25 (`sim/warlock/pet.go:287-320`). Their values differ from Era's pet_levelstats (entries 416, 1860, 1863, 417) by up to about 30% on some stats. Imp at 60: Str 122 and Int 264 match, but Agi 35 against 27 and Spi 260 against 197. Era's table is itself uneven: the Imp's Intellect falls from 98 at 25 to 63 at 30. So the sheets stay as they are. | mangoszero pet_levelstats; `sim/warlock/imp.go`, `succubus.go` | aligned (scales by level); values OPEN against Era's table | - || H7 | Lacerating Strikes: a new Mongoose Bite replaces the bleed with no rollover. It delivers about 18 DPS where a full 40% would give about 34. The client row is a server dummy (1310533 aura 4 at 40), and no beta log reaches the 31-point talent. | `lacerating_strikes.go:75` | OPEN | - |
| H8 | Checked against client rows and tooltips, all aligned:<br>- **Shots:** Multi-Shot 2643 (normalized, no flat amount), Aimed Shot 20904 (+166, 2 s, 310 mana), Arcane Shot 14287 (217 Arcane, 190 mana, 6 s minus 0.3 s a point), Serpent Sting 13555 (83 a tick, 15 s).<br>- **Cooldowns and aspects:** Rapid Fire 3045 (40%, 15 s, 5 min), Aspect of the Hawk 14322 (55 RAP; Quick Shots 30%, 12 s), Bestial Wrath 19574 (+50%, 18 s, 2 min).<br>- **Survival melee:** Raptor Strike 14266 (+70, ranks 3, 4, 6, 7, 8 by band), Mongoose Bite, Counterattack 20910 (50% weapon + 110), Lacerate (17, 28, 40, 58 a tick), Strider Kick 1317257 (100% normalized, 8 s), Wing Clip 14268 (50), Aspect of the Beast 1299447 (110 AP; Quick Strikes 30%), Immolation Trap by band.<br>- **Talents (trait curves):** Lethal Attacks, Mortal Shots, Barrage, Efficiency, Careful Aim, Improved Stings, Focused Fire, Ferocity, Unleashed Fury, Frenzy, Predator's Edge, Resourcefulness, Expose Prey, Survivalist, Lightning Reflexes, Ranged Weapon Specialization, Surefooted, Rapid Killing.<br>- **Pets:** Claw ranks by band, Happiness ×1.25, the cat's ×1.10 damage.<br>- **Racials:** Troll Berserking and Dwarf Artisan's Precision. | hunter fork of this lane; `hunter-audit/rows1.txt`, `tooltips.txt` | aligned | - |
| H9 | Cosmetic: troll Berserking's engine `Note` still says "for 10 s". The aura is 15 s (27 October notes). | `sim/core/racials.go:486` | note only | - |
| H10 | Carried: pets inherit 10% of attack power (guides only), Aspect of the Hawk's proc mask, and the spell power coefficients of Serpent Sting and Arcane Shot. | | OPEN (carried) | - |

## Before/after

**Before:** fork 40021374a with site 55a1212f7. **After:** `lane/r9-physical` (44d0f1707) with `wt/r9-physical`. Each cell gives the relative change at seed 1 / seed 50001, then the absolute values at seed 1.

Band 60, raid preset (Protection in threat per second):

| Spec | alliance % | horde % | seed 1 alliance / horde |
|---|---|---|---|
| Arms | +0.00 / +0.00 | +0.00 / +0.00 | 708.9 → 708.9 / 720.8 → 720.8 |
| Fury | +0.00 / +0.00 | +0.00 / +0.00 | 754.4 → 754.4 / 757.8 → 757.8 |
| Protection (threat) | +0.00 / +0.00 | +0.00 / +0.00 | 1049.0 → 1049.0 / 1036.9 → 1036.9 |
| Assassination | +0.00 / +0.00 | +0.00 / +0.00 | 611.7 → 611.7 / 609.3 → 609.3 |
| Combat | +0.00 / +0.00 | +0.00 / +0.00 | 605.8 → 605.8 / 606.0 → 606.0 |
| Subtlety | +0.15 / +0.17 | +0.11 / +0.15 | 532.6 → 533.4 / 532.6 → 533.2 |
| Beast Mastery | +0.90 / +0.82 | +0.94 / +0.86 | 658.1 → 664.0 / 661.8 → 668.0 |
| Marksmanship | +0.32 / +0.35 | +0.40 / +0.38 | 649.2 → 651.3 / 653.1 → 655.7 |
| Survival | +0.61 / +0.60 | +0.61 / +0.59 | 756.3 → 760.9 / 759.5 → 764.1 |

Causes:
- Beast Mastery: the cat's Bite (H2) and Serpent Sting's refresh (H4), plus Mortal Shots on the sting (H5).
- Marksmanship: the cat's Bite (H2).
- Survival: Savage Strikes (H1) and the cat's Bite (H2).
- Subtlety: Ghostly Strike pooling (R1).
- The rest: nothing in this lane touches them at band 60.

Bare bands (only the bands that move):

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Fury 30 | +3.91 / +3.91 | +3.88 / +3.89 | 61.4 → 63.8 / 61.8 → 64.2 | W1, Battle Stance until 36 |
| Subtlety 40 | +0.42 / +0.42 | +0.21 / +0.32 | 95.9 → 96.3 / 95.0 → 95.2 | R1 |
| Subtlety 50 | +0.07 / +0.15 | +0.00 / +0.00 | 134.7 → 134.8 / 135.1 → 135.1 | R1 |
| Subtlety 60 | +0.14 / +0.09 | +0.09 / +0.09 | 216.6 → 216.9 / 216.0 → 216.2 | R1 |
| Beast Mastery 20 | -15.69 / -15.69 | -15.72 / -15.62 | 103.9 → 87.6 / 105.6 → 89.0 | H3 ammo and quiver by level; H2, H4, H5 |
| Beast Mastery 30 | -8.33 / -8.34 | -8.22 / -8.28 | 134.4 → 123.2 / 136.3 → 125.1 | H3 ammo and quiver by level; H2, H4, H5 |
| Beast Mastery 40 | -2.51 / -2.40 | -2.42 / -2.42 | 191.4 → 186.6 / 190.4 → 185.8 | H3 ammo and quiver by level; H2, H4, H5 |
| Beast Mastery 50 | -1.29 / -1.33 | -1.20 / -1.20 | 240.9 → 237.8 / 242.5 → 239.6 | H3 ammo and quiver by level; H2, H4, H5 |
| Beast Mastery 60 | +1.68 / +1.72 | +1.65 / +1.65 | 308.9 → 314.1 / 309.3 → 314.4 | H3 ammo and quiver by level; H2, H4, H5 |
| Marksmanship 20 | -14.90 / -14.82 | -14.60 / -14.70 | 102.0 → 86.8 / 103.4 → 88.3 | H3; H2 |
| Marksmanship 30 | -9.20 / -9.27 | -9.03 / -9.03 | 136.9 → 124.3 / 138.4 → 125.9 | H3; H2 |
| Marksmanship 40 | -3.09 / -3.14 | -3.48 / -3.37 | 184.6 → 178.9 / 184.0 → 177.6 | H3; H2 |
| Marksmanship 50 | -2.15 / -2.07 | -2.14 / -2.10 | 241.7 → 236.5 / 242.8 → 237.6 | H3; H2 |
| Marksmanship 60 | +0.50 / +0.56 | +0.46 / +0.66 | 302.9 → 304.4 / 305.2 → 306.6 | H3; H2 |
| Survival 20 | +0.58 / +0.73 | +0.72 / +0.57 | 69.0 → 69.4 / 69.9 → 70.4 | H1; H2 |
| Survival 30 | +0.68 / +0.87 | +0.67 / +0.77 | 103.4 → 104.1 / 103.9 → 104.6 | H1; H2 |
| Survival 40 | +0.91 / +0.97 | +1.10 / +1.16 | 175.2 → 176.8 / 173.3 → 175.2 | H1; H2 |
| Survival 50 | +1.11 / +1.07 | +0.94 / +1.03 | 233.7 → 236.3 / 222.8 → 224.9 | H1; H2 |
| Survival 60 | +0.75 / +0.75 | +0.76 / +0.76 | 320.7 → 323.1 / 316.8 → 319.2 | H1; H2 |

Every other bare band of the nine specs is level (within ±0.05%).

The hunters' bare bands 20 to 50 fall because they no longer carry level-52 arrows and a level-55 quiver. Band 60 rises from the cat's Bite and the specs' own fixes.

Assassination's new leveling build (A2) moves the published numbers only at the owner's refresh, because the bands read their talents from the nightly BiS data. Measured with `-talents` on the bare character: band 60 +1.6%, band 50 +1.1 to +1.3%, band 40 +0.5%.


Hunter pets by level (H6), measured on top of the lane's other fixes (E 44d0f1707 → ac0ae9ae6). Relative change at seed 1 / seed 50001, then the absolute values at seed 1:

| Spec, band | alliance % | horde % | seed 1 alliance / horde |
|---|---|---|---|
| Beast Mastery bare 20 | -23.40 / -23.29 | -22.92 / -23.01 | 87.6 → 67.1 / 89.0 → 68.6 |
| Beast Mastery bare 30 | -16.31 / -16.25 | -15.91 / -15.91 | 123.2 → 103.1 / 125.1 → 105.2 |
| Beast Mastery bare 40 | -10.99 / -11.03 | -10.98 / -10.93 | 186.6 → 166.1 / 185.8 → 165.4 |
| Beast Mastery bare 50 | -4.29 / -4.29 | -4.30 / -4.26 | 237.8 → 227.6 / 239.6 → 229.3 |
| Beast Mastery bare 60 | +0.00 / +0.00 | +0.00 / +0.00 | 314.1 → 314.1 / 314.4 → 314.4 |
| Beast Mastery raid 60 | +0.00 / +0.00 | +0.00 / +0.00 | 664.0 → 664.0 / 668.0 → 668.0 |
| Marksmanship bare 20 | -24.54 / -24.42 | -24.24 / -24.15 | 86.8 → 65.5 / 88.3 → 66.9 |
| Marksmanship bare 30 | -13.27 / -13.19 | -13.03 / -13.02 | 124.3 → 107.8 / 125.9 → 109.5 |
| Marksmanship bare 40 | -6.71 / -6.71 | -6.70 / -6.70 | 178.9 → 166.9 / 177.6 → 165.7 |
| Marksmanship bare 50 | -2.45 / -2.45 | -2.40 / -2.40 | 236.5 → 230.7 / 237.6 → 231.9 |
| Marksmanship bare 60 | +0.00 / +0.00 | +0.00 / +0.00 | 304.4 → 304.4 / 306.6 → 306.6 |
| Marksmanship raid 60 | +0.00 / +0.00 | +0.00 / +0.00 | 651.3 → 651.3 / 655.7 → 655.7 |
| Survival bare 20 | -28.24 / -28.24 | -27.98 / -28.02 | 69.4 → 49.8 / 70.4 → 50.7 |
| Survival bare 30 | -15.27 / -15.26 | -15.01 / -15.00 | 104.1 → 88.2 / 104.6 → 88.9 |
| Survival bare 40 | -6.45 / -6.50 | -6.45 / -6.40 | 176.8 → 165.4 / 175.2 → 163.9 |
| Survival bare 50 | -2.24 / -2.25 | -2.40 / -2.40 | 236.3 → 231.0 / 224.9 → 219.5 |
| Survival bare 60 | +0.00 / +0.00 | +0.00 / +0.00 | 323.1 → 323.1 / 319.2 → 319.2 |
| Survival raid 60 | +0.00 / +0.00 | +0.00 / +0.00 | 760.9 → 760.9 / 764.1 → 764.1 |

## For other lanes

- **Crit lane:** `scratchpad/r9/physical/crit-notes.md` (A7) has the source-by-source table. The physical specs care most about three things:
  - Rogue poisons and weapon procs should roll spell crit; they are DefenseType 1.
  - Cruelty, Elemental Sharpening Stone, gear aura-52 crit and Agility must stay off spells.
  - A hunter pet should inherit the physical channel.
  - The engine is right to treat Forever's aura-290 buffs (Leader of the Pack, Moonkin Aura, Mongoose, Squid, Malice, Lethal Attacks, Berserker Stance, Recklessness, Ruthless, Elune's Light) as all crit.
- **Auras lane:** Forever's world buffs, from the tooltips read 2026-11-02:
  - Rallying Cry of the Dragonslayer gives +80 spell damage and healing besides 10% crit and 140 AP. The engine has no spell power on it.
  - Fengus' Ferocity adds +230 spell damage and healing to its 200 AP.
  - Warchief's Blessing is 15% haste to spells and attacks; the engine gives melee only.
  - Rallying Cry and Slip'kik's Savvy have "No effect while in a raid instance".
  - None of these is in a preset.
- **Items lane:** E b57ebe9f0 touches `sim/common/item_effects.go` (Gut Ripper, Blackhand Doomsaw, Pendulum of Doom, Vis'kag) and `weapon_proc_rolls.go`. Merge with care. These weapons are ranker alternatives across classes; Pendulum of Doom is one for Enhancement.
- **Integrator:**
  - E 2a8b74fff changes `proto/hunter.proto`: Hunter_Options.Ammo gains SharpArrow = 13 and HeavyShot = 14. data/proto needs `simproto` at the pin. Only hunters use the new values (S 9c7de3c53).
  - TestRotationLadder rows move for these specs. They were not regenerated here.
    - hunter-beast-mastery, hunter-marksmanship and hunter-survival: E 63686fa7e, 2cd085dce, 629e104be, 44d0f1707, ac0ae9ae6; S 9c7de3c53, e164c7c0a.
    - rogue-assassination: S 39bcdf481, the build.
    - warrior-fury: E 21dca29f9.
    - rogue-subtlety: S 431c6006e, if its ladder rows read the rotation.
  - Fork golden TestP1DPSWarrior is regenerated in E 21dca29f9. No other golden moved: the full fork suite (`-tags=with_db ./sim/...`) passes at ac0ae9ae6, except `sim/web`, which fails setup as it does at the pin.
  - No fork commit here moves a spec outside this group, apart from the item-proc rows in X1 for wearers of those weapons.

## Still open

- Warlock demon sheets against Era's pet_levelstats (H6b).
- Armor-debuff stacking under aura 674 (A6).
- Normalized rage's off-hand factor and crit multiplier (W5). Unbridled Wrath on Heroic Strike (W4).
- Lacerating Strikes' rollover (H7).
- Carried: Eviscerate's and Rupture's attack power, Venom's model, Combat's Slice and Dice gate, the Hawk proc mask, pet attack power inheritance, Serpent Sting's and Arcane Shot's spell power, threat values, rage from damage taken, and weapon-proc rates.
- Marksmanship's +0.9% build (held, A1).

Sources (read 2026-11-02): the layered 70707 raw tables listed above. Wago's SpellEquippedItems 70707 (`https://wago.tools/db2/SpellEquippedItems/csv?build=1.60.1.70707`). Wowhead Forever spell tooltips 11198, 8647, 8649, 8650, 11197, 14168, 11597, 11717, 9907, 22888, 15366, 22820, 16609, 22817, 24907, 24932, 19159, 21553, 23894, 1680, 11605, 11581, 23925, 11601, 2565, 18107, 10373, 16549, 1305394. Wowhead Forever item tooltips 2164, 2515, 3030, 11285, 10513, 18042, 2101, 7371, 8217, 7278, 19320, 18714.
