# Round 16, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755 (late re-layer 43dc0f074: hunter rows are text-only). Site wt/r16-hunter from int/r16 c62d77ab8; fork lane/r16-hunter from int/r16 9f7f1f6e4.
Sims: spec-breakdown through `$R/sim.sh`, guide builds, `-bis-dir $R/bis-frozen`, 3000 iterations to explore. No fix moved DPS, so no 10000-iteration before/after was needed.
Log tools: `hunter-tools/` next to this file (wago diffs, wago hotfix search, foreverlogs fetch and the three analyses).

**Result: 0 CONFIRMED-BUGs.** Seed 1 (pets and buffs) is settled by the client and the engine is right. Seed 3 (Hawk's proc) is settled by the beta logs and the engine is right. Seed 2 (the hawk's assault) is now measured in the logs, but its level-60 value is an extrapolation, so it stays OPEN with a two-minute check.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Pets and player buffs: settled. Group buffs never reach a Forever pet.** The 70755 client sets SpellMisc `Attributes_3` bit 0x100 on every party and raid buff. That bit is `SPELL_ATTR3_ONLY_TARGET_PLAYERS`, "can only target players" (TrinityCore SharedDefines.h; WowPacketParser SpellEnums.cs). The flagged spells are Battle Shout (all ranks), Leader of the Pack (24932), Gift of the Wild, Prayer of Fortitude, Prayer of Spirit, Prayer of Shadow Protection, Arcane Brilliance, the Greater Blessings (Might, Kings, Wisdom), Strength of Earth, Grace of Air, Moonkin Aura and Devotion Aura. Current Era (1.15.9.70576) has 0 on every one of them. The control spells (Fireball, Sunder Armor, Retribution Aura, Sanctity Aura, the world buffs) did not get the bit. No wago hotfix row touches any of these SpellMisc records. This is Blizzard's sentence in data: "Pets can no longer receive player buffs that increase their stats." The engine refuses every player buff on the pet (`sim/hunter/pet.go` `RefusePlayerBuffs`), which matches. The single-target versions (Blessing of Might, Mark of the Wild, Power Word: Fortitude, Blessing of Kings, Arcane Intellect) are not flagged, so a player could hand-cast one on a pet. No raid does that, and Blizzard's sentence says it gives no stats anyway. The pet still gets the raid's buffs indirectly, through its 10% of the hunter's attack power and all of the hunter's crit (Hunter Pet Scaling 415429), which the engine models. The +26% upper bound from round 15 was never real. | wago SpellMisc 1.60.1.70755 vs 1.15.9.70576 (`hunter-tools/attrdiff.py`, `byname.py`); wago hotfix search for each SpellMisc record (`hf.py`: 0 rows); Blizzard deep dive (news.blizzard.com/en-us/article/24301515, read 2026-11-04); foreverchanges.pro/hunter-pets (updated 2026-09-24: "Your pet no longer takes buffs from other players, even those it took in Classic"); zockify.com/forever/hunter (updated 2026-11-01). Fork comment E 86fe8c147 cites the flag, so later rounds do not reopen this. | Aligned (settled; closes round 15 W7-1 and open item 1) | - | E 86fe8c147 (comment only) |
| H2 | **Summon Hawk's 18 s assault: measured in the beta logs, not yet priced at 60.** 8 BM hunters (levels 25 to 35) cast Summon Hawk in 48 foreverlogs encounters (`fullscan.py`, `analyze.py`, `ratio.py`). The kit: (a) The hawk is a real pet unit ("Hawk", source_type pet, owner = the hunter) that auto attacks with melee swings (spell -1). (b) It swings about every 2.26 s (median 2247 to 2273 ms at every level), and Rapid Fire hastes it (1.62 s inside Rapid Fire). (c) It lands 4.4 to 7.6 swings per summon. (d) Its average non-crit swing rises with level: 8.4 at 25, 9.1 at 26, 10.1 at 27, 9.4 at 28, 10.4 at 30 and 12.8 at 35. That is about 0.27 of the same hunter's pet swing at 26 and 0.24 at 35, so it grows more slowly than the pet. (e) It crits at 0 to 2% for hunters below Ferocity and 10.6 to 11.6% for those with it, so it seems to take Ferocity but not the hunter's crit. (f) The dive is 35 to 59. A straight-line fit puts a level-60 swing near 23 before boss armor. Two hawks would then add roughly 20 DPS, about 2.7% of BM raid 60. Round 15's dive-only run (Summon Hawk in Arcane Shot's place) was -1.6%, so the net is about +1%: inside the extrapolation's error. Under the adoption rule, nothing changes yet. The web is still split after launch (search 2026-11-04: mobalytics, wowforever.be and wowforeverbuilds take it; overgear and a Mobalytics note prefer Arcane Shot after its attack power raise). | foreverlogs.gg reports 2695 to 2764 (beta), timelines with owner fields; engine `sim/hunter/summon_hawk.go` models the dive only | OPEN. **Two-minute in-game check for the owner:** at 60 with a Beast Mastery hunter, Summon Hawk on any level 60+ target and read ten of the hawk's swings in the combat log. If the average non-crit swing on a raid boss is about 20 or more, model the hawk (a 2.25 s guardian, two at most, hasted by Rapid Fire, Ferocity and Unleashed Fury) and re-run the BM build with Summon Hawk. Under about 15, the guide stays as it is. | - | - |
| H3 | **Aspect of the Hawk's proc: Auto Shot only (the text wins).** In the same 48 encounters, 4 hunters gained Quick Shots (6150) 32 times. Every gain falls on the same millisecond as an Auto Shot cast. They made 432 Auto Shot casts while Quick Shots was down (7.4% procced, Deadly Aspects 4 to 5 points). In the same windows they also cast Summon Hawk 95 times, Multi-Shot 94, Serpent Sting 68, Aimed Shot 51 and Arcane Shot 10, 318 shots in all, with 0 procs. At the Auto Shot rate, about 23 would be expected (p about 1e-10). The client's ProcTypeMask 320 (ranged auto plus ranged-class spell), which Forever narrowed from Era's 87376, is therefore filtered server side to Auto Shot, as the talent text says. The engine rolls only on ranged autos (`talents.go` applyDeadlyAspects). Side note: the log shows the buff at the Auto Shot's cast, not at its impact. The engine rolls at impact, a timing gap of the shot's flight. That is not worth a change. | `hunter-tools/qsattr.py` on foreverlogs 2695 to 2764; wago SpellAuraOptions 125507/126992/135442 (70755: 320; Era 1.15.9: 87376) | Aligned (settled; closes round 15 open item 6) | - | - |
| H4 | **Round 15 regressions held.** Sniper Shot at raid 60 (MM, guide raid build): 9.0 casts at 1846 average, 11.6% of damage, after Multi-Shot under the 1.2 s auto gate. The MM leveling and raid strings in marksmanship.md match round 15 (`532000050/005055200152105/5`, `5/0052550011503051/50005003`). Falcon's own 5% stays in `aspect_of_the_falcon.go`, and no rotation casts Falcon. The BM leveling order in addon-builds.json has Focused Fire at 15 and 16, ahead of Endurance Training (data/curated/ladder-order.json reason intact). | spec-breakdown MM raid 60 horde 801.1 (3000 it); addon-builds.json | Aligned | - | - |
| H5 | Round 10 guide audit's LIKELY ("Falcon's proc may need the talent"): the client gives Falcon's aura its own ProcChance 5 (SpellAuraOptions 239200). The talent's two chance mods name Hawk's and Beast's class masks, and Falcon's mask is 0 (round 15 W2). The engine follows the client. No published number uses Falcon. | wago SpellAuraOptions, SpellClassOptions 70755 | Aligned (no DPS at stake) | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Raid 60 breakdown (guide raid build, frozen BiS, horde): Auto Shot 47%, cat 22% (white 194 swings at 1.2 s; Bite 16.2 and Claw 23.1 a fight), Multi-Shot 20.6 casts, Arcane Shot 27.0, Aimed Shot 7.6, Serpent Sting 12.7 casts. Bestial Wrath 2 casts, Rapid Fire 1, 0 s out of mana. Pet crit 33.8% (the hunter's crit plus Ferocity 10%), miss 2.1% (the hunter's hit, the same as Auto Shot), dodge 6.4% (from behind, no parry). Deadly Aspects 9.45 procs (10% of about 95 landed autos). Everything matches the client and the round 14/15 checks. | spec-breakdown 728.3 (3000 it) | Aligned | - | - |
| B2 | Builds decoded against talents/hunter.json: every point either does something in the rotation or is a gate (Endurance Training 4 and Bestial Swiftness 1 are the tier-0/2 path to Intimidation and Bestial Wrath; Intimidation is the prerequisite). | talents/hunter.json prerequisites | Aligned | - | - |
| B3 | Multi-Shot is the single Forever spell 2643 (6 s, 13.9% base mana, normalized weapon damage, instant). The engine's cancel and re-enable of the auto keeps the swing time, so the instant shot never delays an auto. | multi_shot.go; core/attack.go EnableAutoSwing | Aligned | - | - |
| B4 | Aspect of the Hawk 14322 at 110 (the push 113534 value; the shipped file says 55). Rank 7 (25296) is a level-60 book, pending owner. | wago hotfix 692138; spellconst/hunter.json | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Raid 60 breakdown: Auto Shot 53%, Multi-Shot 21.8 casts, Sniper Shot 9.0, Aimed Shot 5.8, Arcane Shot 17.6, Serpent Sting 11.6. Hit capped (0% miss) with Surefooted 3. No pet (Lone Wolf). | spec-breakdown 801.1 | Aligned | - | - |
| M2 | Survival dip filler: Surefooted has no prerequisite. Tier 2 needs 10 Survival points, so Survivalist 5 is filler on top of Improved Tracking 5. Every tier-0/1 Survival option is equally dead for a ranged hunter. Improved Tracking is modelled (`improvedTrackingApplies`: every Phase 1 boss is a trackable type). | talents/hunter.json; talents.go:632-660 | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Raid 60 breakdown: white 47% (main and off hand plus Windfury), Raptor Strike 24 hits, Mongoose Bite 20.3, Strider Kick 19.1, Immolation Trap 9.7, Lacerate 9.5 plus Lacerating Strikes, Wing Clip under Windfury Totem. Rapid Fire is cast, and that is right: Forever's 3045 carries aura 319 (melee haste 40) next to aura 140 ("ranged and melee attack speed by 40%"). | spec-breakdown 861.0; wago SpellEffect 3045 | Aligned | - | - |
| S2 | Build: Survivalist 3 is tier-1 filler (Savage Strikes is capped at 2; Entrapment and Improved Wing Clip do nothing in a fight). Deterrence 1 is the prerequisite for Counterattack, and Expose Prey 2 for Lacerating Strikes. | talents/hunter.json | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | BiS (frozen): no set bonus is active on any hunter raid-60 or bare-60 set (single pieces only). Effects in use: Molten Heart of the Mountain (MM, modelled in common/item_effects.go), Teebu's Firebolt (round 15 X2), Gloves of Holy Might, Electrified Dagger and Glacial Blade (bare 60 Survival). | bis-frozen/hunter-*.json; sets.json | Aligned | - | - |
| X2 | Pre-existing fork test failures on int/r16 9f7f1f6e4, before any change of mine: `TestHunterPetDoesNotInheritTheSharpeningStone` (panic: no item 1728) and `TestForeverRangedHitItemsReachRangedAttacksOnly` (no item 2098). Both items are in assets/database/db.json but not in the test item set the package loads. These look like environment issues, not hunter logic. | `go test ./sim/hunter/` through the lock | For other lanes (core/test harness) | - | - |

## Before/after

No change moves DPS (the one fork commit is a comment). Raid 60 on the lane head equals int/r16, at 3000 iterations, seed 1, horde, guide builds and frozen BiS: Beast Mastery 728.3, Marksmanship 801.1, Survival 861.0.

## Files to regenerate

None.

## Measure at refresh

Nothing new. Round 15's list stands: MM raid and bare bands on the new builds, and Survival's bare-60 Intellect weight.

## For other lanes

- Withdrawn: the "build 70783" hotfix rows are retail 12.1.5.70783 region pushes, not Forever (still 1.60.1.70755, no new push since the re-layer). Rule: check a wago hotfix row's region before reading it as Forever (Forever is region 91).
- Core/test harness: X2.
- The class rule from H1 applies to warlock demons too (`sim/warlock/pet.go` RefusePlayerBuffs). The same client flag supports it.

## For the guide lane

- marksmanship.md on int/r16 still has the stale round 15 prose (line 85: "so neither build takes it"; line 69: the 30/16 split; line 72 Barrage; line 80 Improved Stings 3), and the Sniper Shot card note is empty. **land/guides-r10 already carries all six fixes (HUNTER-R10-1 to R10-6) and the card note.** This lane did not touch them, so it must not conflict. Please land guides-r10's hunter changes with round 16.
- No prose change follows from this lane: no build, APL or card rank moved. The guide line "take no stat buffs from other players" (marksmanship.md:85) is now backed by client data.

## Guide text not done

None (no guide-affecting change).

## Still open

- H2: the hawk's level-60 swing damage (the two-minute check in H2).
- Carried: cat vs raptor (owner question 4); Ranged Weapon Specialization on Serpent Sting; Lacerating Strikes on refresh (H7); the queued Raptor Strike lifting the off-hand miss penalty (H10). Pending owner (not reopened): Deadly Aspects under Falcon text, the level-60 trainer book ranks, Arcane Shot's tooltip with spell power.
- Warcraft Logs stays blocked. foreverlogs.gg has no reports after 2813 (2026-10-29), so there are no launch-day logs yet.
