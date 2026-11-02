# Sim audit round 11: regression lane

The owner's ask: re-audit every change rounds 9 and 10 made, for side effects, and name the commit behind each bug, so the bug count can be tracked from round to round.

The fork range checked is `40021374a..cfcb76e80`: 88 commits, with round 9 pinned at 4d3129cae. The site ranges are 55a1212f7..bb619a188 (round 9) and bb619a188..8f047a62f (round 10). The engine branch is `lane/r11-regression`, from cfcb76e80; the site branch is `wt/r11-regression`.

**Result: 5 confirmed bugs, 3 from round 9 and 2 from round 10, none older.** None of the five moves a published DPS or HPS number. Two are tooling or coach bugs, one is a shaman edge case on a set no published band wears, and two are copy.

## How it was checked

1. **Crit trace across three pins.** I added a temporary hook to the three crit functions, reusing round 10's trace patch; it was never committed. It logs each spell's first crit chance and its parts: every-table, melee and spell crit, the spell's own bonus, school crit, and Agility and Intellect. I ran it on 40021374a (round 8), 4d3129cae (round 9) and cfcb76e80 (round 10). Each run covered every published spec at raid 60 and bare 20 to 50, horde, 300 iterations. Each pin ran against its own site checkout, because origin/main's raid preset names potions that 4d3129cae does not know. The round-9 to round-10 diff found 43 changed rows; round 8 to round 9 found 755.
2. For round 8 to round 9 I compared total crit: old Crit against new Crit + MeleeCrit + SpellCrit, per spec and unit. Every residual was explained (below).
3. Every round-9 and round-10 class, item, aura, rank and heal-profile change was read as a diff, checked against what else it touches, and given at least one targeted Go test or 3000-iteration measurement.

## Crit split and the round-10 crit sweep

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Healing Stream's ticks still crit from a school's spell crit. Round 10 took the SpellCrit stat back off each tick so the totem crits on every-table crit only. In the same round, the crit lane sent school crit into every heal. So The Ten Storms 5-piece ("Nature spells", aura 57, which is spell crit) still reached the ticks. | sim/shaman/water_totems.go ApplyEffects; sim/core/spell_result.go healingCritRating. `TestHealingStreamIgnoresANatureSchoolSpellCrit` failed first: 50 of 50 ticks crit at +100% Nature crit, want 0. | CONFIRMED-BUG (no published number; no band wears the set) | R10 E 30003fb0d with E b4546cdb6 | E aa1f866fb |
| C2 | Every round-9 to round-10 crit change on a published spec and band traces to an intended fix, and none is unexplained. Conviction moved Crit to MeleeCrit: Retribution's and Protection's Consecration and Holy Shield lose 6%, seals and Hammer of Wrath keep it. Hack and Slash moved off the poisons and Stormshroud (-5%). Healing Stream lost its Intellect crit. The Succubus moved to its measured sheet. Combustion no longer reaches Ignite. Savage Strikes and Predator's Edge reach Lacerate (+4). The stone now works in forms (cat and bear +2) and on the tanks (Protection warrior +2, paladin +1 spell from the oil). Gloves of Holy Might and Thornspike roll the ranged table. | crit trace e9 to e10, 43 rows | Aligned | - | - |
| C3 | The per-spell stone (E 56a5e5134) gives the player exactly the rolls the old MeleeCrit stat did. Warrior, rogue, Retribution and Survival melee rows are identical between the two pins. It never reaches a ranged attack (the CastType and DefenseType filters), a pet, or a DefenseType-3 proc. Two stones still add +2 each to both hands, which is what cmangos does: the weapon-dependent aura applies to every attack type whose equipped weapon fits the mask. | trace; sim/core/consumes.go addMeleeWeaponCrit | Aligned | - | - |
| C4 | The round-9 split changed only the table each roll reads. Melee rows for warriors and rogues are unchanged, and poisons lose their Agility crit as intended. The total-crit residuals between rounds 8 and 9 are explained by three things. (a) Era's base melee column is now kept, an unread table for casters: mage 3.2, priest 3.0, warlock 2.0, shaman 1.7, druid 0.9, paladin 0.7. (b) Round 9's preset changes at 60: Blasted Lands per spec adds about 27 Intellect for casters and about 30 Agility for the Agility specs. (c) Enhancement now drops its own Grace of Air, which adds 49 Agility at bare 50 and 85 at raid 60. That is legitimate: the raid's Windfury Totem belongs to another shaman, and its raid buff list carries no Grace of Air. | resid.py with Agility and Intellect columns | Aligned | - | - |
| C5 | Fury at bare 30 loses Berserker Stance's 3% crit. That is round 9's Battle Stance until Whirlwind at 36, as intended. | trace | Aligned | - | - |
| C6 | Elemental Devastation is fed by Earth Shock and Lightning Bolt crits only. The Searing Totem bolt is ProcMaskEmpty, so a totem crit never procs it (about 10 procs a fight at 60, against about 10 spell crits). | sim/shaman/talents.go, fire_totems.go; Enhancement 60 breakdown | Aligned | - | - |
| C8 | The same trace was run for the variants the first pass skipped, at the same three pins (each against its own site checkout), 300 iterations, seed 1: Alliance at raid 60 and bare 20 to 50, bare 60 for both factions, and the raid preset at bands 20 to 50 for both factions. No published BiS entry exists for raid below 60, so the bare band's gear was relabelled raid in a scratch bis-dir. Round 9 to round 10 changed 118 rows on Alliance and bare 60, and 281 on raid below 60. Every change is either an intended fix (Conviction, Hack and Slash, Healing Stream, Savage Strikes and Lacerate, Combustion off Ignite, the stone in forms and on tanks, the oil on Protection paladin, the stone leaving the Survival pet and Thornspike's ranged roll) or a first-call artifact: Night Elf Elune's Light (+10 for 15 s), Elemental Devastation, Divine Favor, Inner Focus and Combustion stacks. Round 8 to round 9 total-crit residuals are explained as before: the base melee column, the round-9 per-spec Blasted Lands and food changes (they scale with the per-level crit rates, so they are larger at band 20), Enhancement's own Grace of Air, Fury's Battle Stance at 30, and the pets. The racial heal masks give Discipline's Forever racial heals (1277377, 1277378) Holy Specialization's +5, as round 9 intended. | crit/t2-*, cmp2.py, resid2.py | Aligned | - | - |
| C9 | Side note from the raid-below-60 trace: spec-breakdown (and any engine-side resolve of the raid preset) does not cut consumables to the character's level. At band 20 the raid set carries Elemental Sharpening Stone (needs 50) and Brilliant Wizard Oil (45). The /sim page's own set, sim/request/raid-consumes.json, is cut by the client's RequiredLevel (`ResolvedPreset.UsableAt`), and the published bands below 60 run bare. So no published or page number is affected. This predates round 9; round 10 only added the tanks' rows. | sim/request/raid_consumes.go, preset.go:286; spec-breakdown band 20 raid | OPEN (old, tooling only) | old | - |
| C7 | The crit-sweep fixes read for side effects: Conviction on heals (ruling 2 is the crit lane's and was left alone), Libram of Infusion, Exorcism, Inner Focus off Renew, Totem of Urgency, demon crit, hunter pet spell crit and Web, Weaponmaster and Hack and Slash with the off-hand correction. | class review below | Aligned | - | - |

## Rank choices, heal profiles, Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | talent-search and rotation-search ignored the bare-60 rank-choice rows. Neither sent a preset, so a bare-60 search fought the raid preset, which has no raid-60 rows. rotation-search also ran the curated rotation with no rank rewrite. 14 specs have bare-60 rows, and any bare-60 adoption measured with these tools rested on the wrong baseline. | Bare-60 BiS set, horde, seed 1, 3000 iterations, no preset against bare: Elemental 140.3 vs 190.8, Shadow 342.0 vs 353.7, Retribution 242.3 vs 246.3. The bare figure is spec-breakdown's bare 60. `TestSearchRequestNamesItsPresetForTheRankChoices` and `TestCandidatesCastThePresetsRankChoices` failed first. | CONFIRMED-BUG (tooling) | R9 S 9de1cf834 (the bare-60 rows); S 65212f9be fixed only the ranker | S 531bd7f36 |
| K2 | The coach calls the site's own authored downrank a low rank. Elemental's raid-60 Lightning Bolt is rank 9 by design, and rank 10 trains at 56. F5 exempted only rank-choices.json rows, so a shaman following our rotation was told "N casts of Rank 9". | api/internal/coach/f5_ranks.go; data/curated/apl/shaman-elemental.json (15207); trainables rank 10 = 15208 at 56. The new f5 case failed first. Live after the addon-data.json refresh. | CONFIRMED-BUG (coach copy) | R9 S dcd25dee3; the round-10 coach fix S 2b859fadf covered rows only | S a2d8e3723 |
| K3 | Every rank-choice row is checked against the client. Each row names a trainables family, the authored id is in that family, and the chosen rank exists and is learned by the row's from_level. No two rows overlap. Every chosen rank's cost, amount and coefficient match the fork's conformance goldens, including the below-20 penalty 1 - (20 - L) x 0.0375 (Holy Light rank 3 x0.775). | script join; sim/core/spell.go levelPenalty | Aligned | - | - |
| K4 | The rank rewrite covers conditions as well as casts, so Retribution's mana reserve prices the chosen rank. /sim, Top Gear, weights, the ranker and spec-breakdown send the preset. The addon rank table follows the rows; addon-data.json is held output. | rotation_ranks.go, run.ts, bulk-store-request.ts, addonrotation.py | Aligned | - | - |
| K5 | Band heal profiles reach healers only. Attach and AttachWeights are called from the healer fight, spec-breakdown heal, the in-process healing run, the ranker's heal engine and dps-per-point's healer path; no DPS or tank path reads them. The band follows the healer's own level. ladder-first.json is read by the ranker, talent-search, rotation-search, dps-per-point and addon-builds. | caller audit; Holy priest band 30 vs 60 | Aligned | - | - |
| K6 | Retribution's paced Consecration. At raid 60 over 60, 180 and 360 s it spends 0 to 0.2 s out of mana with about 88% uptime. Bare 30 to 59 casts rank 1 and bare 60 rank 4, per the rows. Bare 20, 50 and 60 run dry for 6 to 9 s at the very end, which is the line spending its pool. Bare 20's loss is ruling 3 and was not reopened. Protection paladin's APL is untouched by rounds 9 and 10 and keeps its own gate. | 7 runs, 1000 iterations | Aligned | - | - |
| K7 | The twist wait reads Unsealed Wrath's 15 s refresh correctly. At raid 60 Seal of Command is cast 14.8 times and Righteousness 15.1 (it was about 28 each, one per global); the bare bands never twist. | unsealed_wrath.go | Aligned | - | - |

## Forms, potions, auras and items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Round 10 moved the bear's raid potion to Mighty Rage Potion but left three strings naming Greater Stoneshield: the bear rotation's autocast note, the presets note, and a row reason calling it "the bear's potion below level 46". Both potions need level 46 in the client's ItemSparse, so Stoneshield is never the bear's potion. The preset itself is right. | data/curated/apl/druid-feral-bear.json, presets.json; ItemSparse 13455 and 13442 RequiredLevel 46. `tank_potion_notes_test.go` failed first. | CONFIRMED-BUG (copy) | R10 S fd109430b | S cf89e7c2c; E 098a439dc (fork APL copy) |
| F2 | Only druids shapeshift, so the form rules move nothing else. Explosives still cancel the form. Balance at 60 drinks Spellblasting twice and a Demonic Rune in Moonkin Form (99.8% uptime). Thistle Tea's move into core is line for line the deleted rogue file. | diff; Balance and cat breakdowns | Aligned | - | - |
| F3 | The stone is applied once in forms: white swing +2.0 crit, Shred +1.8. Weightstones reach the paw through BonusMHDps (Shred 693 to 699). Windfury Totem procs about 45 times a fight on the cat, and the Windfury Weapon exclusion holds. | cat with and without | Aligned | - | - |
| F4 | The four 30-second potions match the client: 16, 47, 28 and 80; 30 s duration; category 4 with a 120 s recovery shared with the mana and rage potions; levels 25, 55, 25 and 55. No APL prepots. Mighty Rage gives 45 to 75 rage and 60 Strength for 20 s; the bear drinks two a fight. | SpellEffect, SpellDuration, SpellCategories, ItemSparse | Aligned | - | - |
| F6 | The three procs from E d454d6b42 that the first pass did not check, row by row against the 70707 tables (wago.tools, fetched 2026-11-02). Barman Shanker (12791): ItemEffect trigger 2 (chance on hit) to 1300787 Rend, which is DefenseType 2, physical, mechanic 15 (bleed), duration index 165 (7 s); effect 0 is aura 3 at 14 every 1000 ms (7 ticks) and effect 1 is aura 4 with base points 2, the "2 times from behind" multiplier. The engine has 14 a tick for 7 s, doubled when not in front, a melee special hit to land, and ticks that crit as a bleed (`Dot.OutcomeTick` with CanCrit, the dots-crit ruling). The Lobotomizer (19324): trigger 2 to 1290950 Brain Damage, school 32 (Shadow), DefenseType 1, effect 0 school damage 250 with variance 0.4 (200 to 300), bonus coefficient 0; effect 1 is aura 65 at -25 (cast slow) for 30 s, which does nothing to a boss. The engine has Shadow, magic hit and crit, Roll(200, 300), no coefficient. Draconic Infused Emblem (22268): trigger 1 (equip) to 1318931, aura 42 proc on ProcTypeMask 0x10000 (harmful spell), ProcChance 100, no PPM row, triggering 1318930 for 10 s (duration index 1): aura 13 +35 spell damage (misc 126, all magic) and aura 135 +35 healing. Its third effect (aura 180, +35 against creature type 2, Dragonkin) meets no published target, which is type-neutral. The engine has +35 spell damage and healing for 10 s on a harmful spell. The one-a-minute rate stands in for a client ProcChance of 100 with no rate row, the same open item as every unstated proc. | ItemXItemEffect, ItemEffect, SpellEffect, SpellAuraOptions, SpellCategories, SpellMisc, SpellDuration (wago 70707) | Aligned (the proc rate stays OPEN with the other unstated rates) | - | - |
| F5 | All match the client: Thorns, Retribution Aura and Stoneskin ranks; Strength of Earth and Grace of Air amounts; party Mana Tide by level; Windfury and Rockbiter AP by level; Blasted Lands buffs, foods and salmon; Orc's Command; the raid's Winter's Chill and Curse of Shadow flags; tanks' oil and stone levels; base mana; Aegis's 5 s cooldown; Dormant Heart; the execute weapons; the fourteen retuned procs; Blazing Rapier; the use effects of Mindtap, Burst of Knowledge, Nat Pagle's, Sanctified Orb, Diamond Flask and Smolderweb's Eye; Gutwrencher and Razor Gauntlets. | client rows; engine tests | Aligned | - | - |

## Class changes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | The fork UI's basic Protection paladin preset gated Holy Shield on Redoubt's old buff id 20134. Round 9 moved Redoubt to the 70707 id 20128, so the line could never fire. Neither the site's APL nor the fork tests use this preset. | sim/paladin/tank_talents.go:78; ui/protection_paladin/apls/basic_prot.apl.json | CONFIRMED-BUG (fork UI only) | R9 E 261a43a91 | E fe7f027f5 |
| G2 | The other 33 class changes are aligned; see the closing table. The full `-tags=with_db` suites for all nine classes and sim/core pass on the lane. | per-change checks | Aligned | - | - |

## Before/after

No fix moves a published DPS or HPS number, so there is no 10,000-iteration table.
- C1: Restoration shaman effective HPS is unchanged, horde, seed 1, 3000 iterations: 834.1 at raid 60, 245.4 at bare 50, 65.1 at bare 40.
- K1 changes only the search tools' bare-60 baseline, which now equals spec-breakdown's bare 60. Seed 1, 3000 iterations: Elemental 140.3 to 190.8, Shadow 342.0 to 353.7, Retribution 242.3 to 246.3.
- K2, F1 and G1 are copy, coach or fork-UI changes.

## For other lanes

- **Ranks lane:** `parseRankChoices` checks overlap only between rows that are adjacent after sorting. A non-adjacent overlap in the same preset would pass, and rankAt would take the first match. No such row exists today.
- **Auras (presets):** raid rows at levels 35 to 49 give Fury, Arms, Survival, Retribution, Protection warrior, cat and bear no main-hand weapon consumable. The off hand gets a Dense Sharpening Stone from 35, but the main hand waits for the Elemental Sharpening Stone at 50, though the Dense stone and weightstone need level 35. This is old for the warriors and R9 S be3314ab5 for Survival. It affects only /sim's raid option at those levels and is not measured. The fix is to list the Dense stone ahead of the Elemental one in each main-hand list.
- **Crit lane:** C1 is fixed in the shaman file. If a core "every-table only" heal helper is preferred, it can replace the two subtractions.
- **Platform (coach):** K2 touches api/internal/coach/f5_ranks.go and api/internal/sims/rotation.go. `go vet ./...` in api is clean.
- **Integrator:** the BiS json files that embed preset notes pick up F1's wording at the refresh (held output).

## Still open

- Healing Stream's crit rule rests on low-level logs, where shamans without Thundering Strikes show 0 crits in 21,134 ticks. That is equally consistent with the totem not critting at all without Thundering Strikes. At raid 60 the sim gives it 13% (raid every-table crit). A level-60 log would settle it (healers or crit lane).
- C9: the engine-side raid preset is not cut to level below 60 (old, tooling only). Platform or ranker lane: confirm the API never resolves the raid preset server-side for a sub-60 /sim request, since the page sends its level-cut set.
- Draconic Infused Emblem's rate: the client says ProcChance 100 with no rate row; it is held at one a minute with the other unstated procs until a log shows the uptime.

## Closing table

| round-9/10 change | checked how | verdict | bugs caused |
|---|---|---|---|
| Crit split (R9 E c14bf02a5; S 67f5e9d9b, 8235fc1b9) | three-pin crit trace on both factions, the published presets plus raid 20 to 50 and bare 60; residuals with Agility and Intellect | aligned | 0 |
| Conviction (R10 E fc8c92a4e, 03ec2ecd8) | trace; heals keep it | aligned | 0 |
| Gloves of Holy Might / Thornspike table (R10 E 495779597, c2eed2ba6) | trace (DefenseType 3, no stone) | aligned | 0 |
| School crit to heals (R10 E b4546cdb6) | test with Ten Storms 5 on Healing Stream | bug | 1 (C1, with E 30003fb0d) |
| Emerald Dragon Whelp (R10 E 49db95da4) | diff | aligned | 0 |
| Per-spell stone, pets, forms (R10 E 56a5e5134, 193ecffc2, 2eb068e75, merge) | trace; cat with and without | aligned | 0 |
| Weaponmaster / Hack and Slash (R10 E b71cc6975, 5849efeca) | trace (poisons -5, melee unchanged) | aligned | 0 |
| Demon crit (R10 E de87d2393) | trace; Lash of Pain 5 to 6% at 20 to 50 | aligned | 0 |
| Hunter pet spell crit, Web (R10 E db7e15a84) | inheritance map | aligned | 0 |
| Healing Stream crit (R10 E 30003fb0d) | new test | bug | 1 (C1) |
| Combustion off Ignite (R10 E 2af230c0e) | trace; stack logic | aligned | 0 |
| Inner Focus off Renew, Mental Agility (R10 E 8a241f75a) | masks | aligned | 0 |
| Libram of Infusion, Totem of Urgency, Exorcism (R10 E ac11976ed, 999f1085b, f1c7fa888) | masks, suites | aligned | 0 |
| Downrank rule (R9) | rotation_ranks.go, resolver | aligned | 0 |
| Per-band rank-choice rows (R9 S 9de1cf834 and re-measures) | client join; preset plumbing in every tool | bug | 1 (K1, S 9de1cf834) |
| Raid-60 authored downranks (R9 S dcd25dee3, E 0a9ef0d0c) | coach F5 test | bug | 1 (K2, S dcd25dee3) |
| Rank costs, amounts, coefficients, below-20 penalty | conformance goldens | aligned | 0 |
| Addon rank table | addonrotation.py | aligned | 0 |
| Coach reads rank choices (R10 S 2b859fadf, 8f047a62f) | test | incomplete, counted under K2 | 0 |
| Band heal profiles (R10 healprofile_bands.go) | caller audit; band 30 vs 60 | aligned | 0 |
| ladder-first.json (R10) | consumer audit | aligned | 0 |
| Retribution's paced Consecration (R10) | 7 runs over bands, presets and lengths | aligned | 0 |
| Retribution's twist wait (R10) | Unsealed Wrath read; cast counts | aligned | 0 |
| Form consumables, Thistle Tea (R10 E d5c560d31) | diff; Balance and cat runs | aligned | 0 |
| 30-second potions (R10 E 61734a1cd) | client rows | aligned | 0 |
| Bear's Mighty Rage (R10 S fd109430b) | client rows; preset notes test | bug (copy) | 1 (F1, S fd109430b) |
| Aura ranks: Thorns, Retribution Aura, Stoneskin, Strength of Earth, Grace of Air, Mana Tide (R9 E b7c81850b, 09d7ffddb, 23aee9172) | every rank against the client | aligned | 0 |
| Windfury/Rockbiter AP by level, Windfury Totem under Windfury Weapon (R9 E 60205a769, S cefa42f10) | client rows; raid buff list | aligned | 0 |
| Blasted Lands per spec, foods, Orc's Command, Winter's Chill / Curse of Shadow (R9) | client rows | aligned | 0 |
| Tanks' imbue (R10 S f8ad6d401) | required levels | aligned | 0 |
| Base mana (R9 E 918047272) | every-level test | aligned | 0 |
| Aegis, Dormant Heart, terrain riders, execute weapons (R9) | client rows, callers | aligned | 0 |
| The round-10 item fixes (E 71a463e62, bc2ef10fa, d454d6b42, caa405a7d, 3bf3a1d1f) | client rows, every item one by one (F6 covers Barman Shanker, The Lobotomizer and Draconic Infused Emblem) | aligned | 0 |
| Evocation (R9 E 388fbc7c8) | Frost bare 20 | aligned | 0 |
| Clearcasting / Omen / Maelstrom at the proc (R9 E d3b5b796b, ff4bae666, d985fbb84) | proc paths; suites | aligned | 0 |
| Hunter pets and ammo by level (R9 E ac0ae9ae6, 2a8b74fff) | sheets; site tables | aligned | 0 |
| Savage Strikes, Lacerate, Predator's Edge, Mortal Shots, Serpent Sting refresh, Bite, Summon Hawk (R9/R10) | masks, trace, breakdowns | aligned | 0 |
| Expose Armor 450, Fury stance, Ghostly Strike (R9) | runs at 30 and 40; diffs | aligned | 0 |
| Twin Disciplines, racial heal masks, Nature's Swiftness, Swiftmend, Nature's Grace, Gift of Nature (R9/R10) | masks; Restoration druid run | aligned | 0 |
| Fingers of Frost, Improved Flamestrike, Precision (R10) | flags, mage suite | aligned | 0 |
| Touch of the Grave, Spellstone, Demonology prepull (R9/R10) | runs; all 28 APLs' prepulls | aligned | 0 |
| Talent buff ids, Elemental Devastation id (R9 E 261a43a91, R10 E af3901df3) | grep for the old ids | bug (fork UI) | 1 (G1, E 261a43a91) |

**Totals: 5 confirmed bugs. R9: 3 (K1 S 9de1cf834, K2 S dcd25dee3, G1 E 261a43a91). R10: 2 (C1 E 30003fb0d with E b4546cdb6, F1 S fd109430b). Old: 0. New-data: 0.** K2 is counted under round 9, which introduced the authored rank; round 10's coach fix (S 2b859fadf) missed it.
