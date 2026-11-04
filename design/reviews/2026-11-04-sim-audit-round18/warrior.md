# Round 18, lane warrior (Arms, Fury, Protection)

Site `wt/r18-warrior` on int/r18 f2733b9f7. Engine `lane/r18-warrior` on fork int/r18, merged up to af1507a41 (as the coordinator asked), head 4742b0075. Client 1.60.1.70755 with the late hotfix re-layer. Client rows come from the layered 70755 raw tables (`scratchpad/c755/fetch/data/builds/1.60.1.70755/raw/`), `talents/warrior.json` and Wowhead Forever tooltips (nether.wowhead.com/forever).

**Result: 0 confirmed warrior sim bugs.** One shared-core fix came from the coordinator's racials audit: the Berserking id plus three wrong racial comments (see "Racials"). It moves no number. No guide-visible change, so there is no guide text and no rendered review.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Heroic Strike as a white swing for procs** (Unbridled Wrath, Flurry, Reckless Swings). Still no level-60 Forever log. The newest foreverlogs.gg report is still 2814 ("Test", 2026-11-04 12:24 UTC); 2815 and up return "Report not found". `api.foreversixty.gg/v1/reports/recent` lists one 2026-10-20 report with no boss. A web search found no Forever source that settles it. wow.gg's Forever Fury overview (updated 2026-09-18, https://wow.gg/guides/warrior-fury-forever-overview) says nothing about it, and it contradicts the client on Unbridled Wrath ("raises your maximum rage pool"), so it is not a source. The engine stays as round 17 decided. | foreverlogs.gg API, 2026-11-04 ~17:00 UTC; web | OPEN (settles with a level-60 log after launch) | - | - |
| S2a | **Round-18 rogue Thistle Tea change (E 1aa12229b)** touches only energy bars (`thistleTeaFits` reads `MaxEnergy`). A warrior has no energy bar and never registers the tea. **Round-18 Hemorrhage change (E 05ccb0d1d)** is in `sim/rogue` only. | fork diff b2c4a2d05..int/r18 | Aligned | - | - |
| S2b | **Round-17 hit weight (E 7c8f710d0)** widens the hit sweep past the weapon cap only when `computeSpellHitProfile` finds a damage spell. `castsDamageAtTarget` (`core/hitprofile_spell.go:57`) needs `SpellFlagAPL`, `ProcMaskSpellDamage` and a non-physical school. Thunder Clap is the warrior's only `ProcMaskSpellDamage` spell, and it is Physical. Explosives are `ProcMaskEmpty` with no APL flag (`core/consumes.go` `newBasicExplosiveSpellConfig`). So no warrior profile changes. `dps-per-point.json` warrior rows are byte-identical between d7fc51fe7 and the round-17 regen. | code; JSON diff | Aligned | - | - |
| S2c | **No drift, measured.** Raid 60 at int/r18, both factions, seeds 1 and 50001, 10000 iterations, frozen BiS, guide builds: all 12 warrior values equal `$R/before.tsv` to the decimal, before and after this lane's commit. Fork `go test -tags=with_db ./sim/warrior/...` is green (all three packages). Site `TestRotationLadder/warrior-*` compares clean against the committed goldens. | measure.sh; go test | Aligned | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Every rank text of every talent the leveling and raid builds take was re-read against the engine constants. Improved Heroic Strike 1/2/3; Improved Rend 12/23/35; Improved Charge 3/6; Improved Tactical Mastery 5/10/15; Improved Overpower 25/50; Deep Wounds 20/40/60; Two-Handed Weapon Specialization 2/4/6 (`twoHandedWeaponSpecializationPct`); Impale 10/20; Bloodthrill 10/20; Weaponmaster 1-5 crit / 3-15 armour; Improved Slam; Mortal Strike; Cruelty 1-5; Unbridled Wrath 12-60; Enrage 2-10 at 30%; Reckless Swings; Improved Execute 3/5; Improved Bloodrage 25/50; Iron Will. All match. | talents/warrior.json; `talents.go`, `talents_protection.go` | Aligned | - | - |
| A2 | Reckless Swings' health cost is a **power cost** (SpellPower PowerType -2, PowerCostPct 25), not a damage event. So it cannot trigger Enrage, whose ProcTypeMask 139944 covers taken melee, ranged and spell hits only (12317 SpellAuraOptions). The engine pays it only when tanking and fires no hit-taken event. Enrage 5 stays a tier-3 gate filler, as round 17 A4 found. In a real raid it does fire on boss and raid damage. The only alternative filler with any effect is Booming Voice, worth about 2 rage a fight. | SpellPower 315180; SpellAuraOptions 124859; `reckless_swings.go` | Aligned | - | - |
| A3 | Arcanite Champion (12790): ItemEffect 99501 is chance on hit, 16916 = Strength 120 plus a heal of 360 with 0.5 variance, DurationIndex 9 (30 s). The engine matches (270-450, 120, 30 s). The proc rate is the engine's assumed 1 PPM, because the 70755 tables carry no SpellProcsPerMinute row. Shared item code, noted only. | ItemEffect, SpellEffect 694634/5 | Aligned (rate unconfirmed, shared) | - | - |
| A4 | BiS set and use effects on raid 60 (frozen BiS): only one piece of any set (Highlander's/Defiler's Resolution) is worn. Thunderbrew's Boot Flask (Alliance) is modelled (round 16 X4). Frozen and Molten Heart's biome doubling is off on the type-neutral fight, by design (`common/item_effects.go:3696`). | ItemSet, ItemSetSpell | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Raid breakdown (horde Troll, seed 1): Bloodthirst 22.1, Whirlwind 16.4, Execute 24.1, Heroic Strike 34.1, Reckless Swings 10.6, Berserker Rage 5.8, Death Wish 1.09 (10 rage, client SpellPower 178801 ManaCost 100), Recklessness 1, Mighty Rage Potion 2 (30 s potions). Unbridled Wrath gives 97.2 rage. It matches round 17. | spec-breakdown | Aligned | - | - |
| F2 | Death Wish 12328 on 70755: +20% physical (aura 79), fear immunity, **+5% damage taken** (aura 87, 5), 3 min, 1.5 s global. Recklessness 1719: 100% crit, +20% damage taken, 30 min. Both match the talent text and the engine. | SpellEffect 690203-5, 687178-80 | Aligned | - | - |
| F3 | **Windfury Hamstring line for Fury.** Arms gains 1.6% from Hamstring under Windfury (round 17 A5), so I tested the same line on Fury, horde, 10000 iterations. Ahead of Heroic Strike: 852.6 at seed 1 (-2.4%). As the last filler: 852.3 / 852.6 at seeds 1 / 50001, against 873.3 / 873.2 without it (-2.4%). Fury needs its globals and rage for Bloodthirst, Whirlwind and Heroic Strike, so the APL is right to leave Hamstring out. | spec-breakdown `-rotation` | Aligned | - | - |
| F4 | The Heroic Strike guard `not spellIsKnown(Execute 20662 rank 5)` resolves correctly below 56. The rank rewrite (`sim/request/rotation_ranks.go` `rewriteRankedSpellIDs`) maps the condition's id to the learned Execute rank, so from 24 to 55 the guard reads "Execute is known" and Heroic Strike holds from 25%. | code | Aligned | - | - |
| F5 | Troll Berserking fires once every 3 minutes with the client's flat 10% variant only. The five tagged custom-percentage cooldowns that share its timer never cast (0 uptime in the JSON breakdown). | spec-breakdown `-json` | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Tank fight (alliance, seed 1, 3000 iterations): 1097.4 threat/s, which is round 17's figure. Revenge 30.0, Shield Slam 30.2, Sunder 47.8, Heroic Strike on 55.7 of 85.5 swings, Shield Block 78% up, chance of death 0%. | spec-breakdown `-tank` | Aligned | - | - |
| P2 | Protection talents rank by rank against the engine: Shield Specialization 1-5% block and 20-100% for 5 rage on a block; Master of Defense 50/100% for 5 rage on a dodge or parry with a shield; Defiance 5/10/15%; Anticipation 4 a point; Improved Revenge 20/40/60; Improved Thunder Clap 2 a point; Improved Sunder Armor 1 a point; Focused Rage 1 a point; Bastion 2 a point; Last Stand 30%/20 s. All match. Improved Revenge's 15% proc chance (SpellAuraOptions 125470) is a leftover with no proc effect (12797 has only effect 0, aura 108 at 60), so nothing to model. | talents/warrior.json; SpellEffect 690356 | Aligned | - | - |
| P3 | Revenge rank 5: 121 with 0.197 variance (109.07-132.93) + 0.25 AP in Defensive Stance. That matches the Wowhead Forever tooltip, `[109.07 + ((sp71 0.25 * sp12797) * AP)]`. Improved Revenge (aura 108, misc 0) scales the whole hit, as the engine's flat-percent mod does. | tooltip 11601; SpellEffect 689703 | Aligned | - | - |
| P4 | Demoralizing Shout rank 5: -196 at level 54 plus 1.4 a level, so 204 at 60, and threat 43 plus 0.8 a level. The engine and the card note ("204") agree. Round 17's "-196" was the base value. | SpellEffect 690063, 1361987 | Aligned | - | - |
| P5 | **Revenge and Shield Slam threat constants** (Revenge 2.25 x damage plus 2.25 x 2 x level; Shield Slam flat 508) are SoD-era upstream values (E d30d87d9c, 2024-05-06). The client carries no threat effect for either. This is carried OPEN since round 8 (physical P2). Revenge is 37% of Protection's threat. A Forever tank log with threat, or a blue post, settles it. | `revenge.go:19`, `shield_slam.go:22` | OPEN (carried, not new) | - | - |

## Racials (coordinator item, shared core)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Berserking used a spell id the client does not ship.** `makeBerserkingCooldown` registered 26297, Classic's energy variant. 70755 has no 26296 or 26297: SpellName and SkillLineAbility hold only 20554 (skill line 733, Racial - Troll). It is now 20554 (`berserkingSpellID`). The test pins 20554 and checks that 26297 is gone. It failed first ("not registered under the client id 20554") and passes now. Results are unchanged to the decimal at 10000 iterations: Troll Fury 873.3 / 873.2 and Troll Destruction 782.7 / 783.4 at seeds 1 / 50001, before and after. The priest healing test constant moved with it. | SpellName, SkillLineAbility; `racials_test.go` TestElunesLightAndBerserkingMatchTheClientRows | CONFIRMED (id only, no number) | old | E 4742b0075 |
| R2 | **Shatter Curse comment** (`racials.go` ~230) said "-15% magic damage taken" and "immunes". Client 1299026: aura 87 at **-5** on school mask 126 for 8 s (DurationIndex 31), two curse dispels (effect 38), 3 min cooldown. Wowhead Forever says "removes all Curses and Banes, and reduces all Magical damage taken by 5% for 8 sec". Corrected. | SpellEffect 1322979-81; tooltip 1299026 | CONFIRMED (comment) | old | E 4742b0075 |
| R3 | **War Stomp comment** (~398) said a 2 min cooldown. Client 20549: SpellCooldowns 90000 (90 s), a 0.5 s cast (CastingTimeIndex 3), 2 s stun (DurationIndex 39), 8 yd (radius 14). The tooltip says 1.5 min. Corrected. | SpellCooldowns 55607; tooltip 20549 | CONFIRMED (comment) | old | E 4742b0075 |
| R4 | **Regeneration comment** (~498) said 10% of regeneration in combat. Client 20555: aura 88 at 20 and aura 116 at 20. The tooltip says "increased by 20% ... 20% of total Health regeneration will continue during combat". Corrected. Rapid Regeneration's "channels back 50%" is right (1260270: 10% every 1.2 s for 6 s, channelled, 3 min). | SpellEffect 697794/5; tooltip 20555 | CONFIRMED (comment) | old | E 4742b0075 |
| R5 | Stale duration notes: the racials test said the shipped row "still gives 10 s", and racials.go cited 70753. The 70755 SpellMisc 329971 is DurationIndex 8 = 15000 ms. Both texts are corrected. | SpellDuration 8 | note | - | E 4742b0075 |

Old-id references found (not changed, for the integrator):
- **Site:** `data/curated/engine-spell-names.json:208` names 26297 "Berserking". After the pin it is unused, because 20554 is named from the client's spells.json. Remove it after the pin moves. Removing it before would leave 26297 unnamed at the old pin.
- **Fork, upstream files:**
  - `ui/rogue/apls/combat_backstab_sweaty.apl.json:8` and `combat_sinister_strike_sweaty.apl.json:8` cast `{spellId: 26297, tag: 6}`. These are non-Forever APLs that no Forever test loads.
  - `ui/raid/components/importers/raid_wcl_importer.tsx:187` maps Classic log ids.
- No site rotation, addon data, coach or guide references either id.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Round-15/16/17 build strings decoded against 70755 (Arms build and raidBuild, Fury build, Protection build): every point is legal. Every point either acts in the rotation or is a forced tier or prerequisite filler (round 17 A4/F5). | decode of talents/warrior.json | Aligned | - | - |
| X2 | Alliance raid preset with Windfury Totem: legal, since Dwarf Shaman is one of Forever's new combos (Deep Dive, news.blizzard.com/en-us/article/24303313). | memory; preset | Aligned | - | - |

## Before/after (raid 60, 10000 iterations, guide builds via `rt`, frozen BiS)

No commit moves a warrior number. The after values (`$R/warrior/after.tsv`) equal `$R/before.tsv`:

| Spec | alliance s1 / s50001 | horde s1 / s50001 |
|---|---|---|
| warrior-arms | 820.4 / 820.4 (0.00%) | 826.4 / 826.5 (0.00%) |
| warrior-fury | 879.4 / 878.9 (0.00%) | 873.3 / 873.2 (0.00%) |
| warrior-protection (threat/s) | 1097.0 / 1096.3 (0.00%) | 1088.0 / 1088.0 (0.00%) |

Specs outside the class: E 4742b0075 changes only the Berserking action id, for every Troll. Troll Fury and Troll Destruction are identical to the decimal at both seeds.

## Files to regenerate

- None. At or after the pin, the integrator removes `data/curated/engine-spell-names.json` "26297" (see Racials). That is a curated edit, not a regen.

## Measure at refresh

- Carried: Arms and Fury race order on the round-15 builds (guide-accuracy round 14 asks, two seeds, Skyborne included); Protection races on the tank fight; Arms Alliance Boot Flask and Frozen Heart picks.

## For other lanes

- Core: E 4742b0075 touches `sim/core/racials.go` (every Troll). No number moves.
- Items: Arcanite Champion's proc rate is still an assumed 1 PPM (A3). Arena Grand Master's absorb is still unmodelled (round 17 P6).
- Log tooling: no launch logs yet. After 23:00 UTC launch, the first level-60 warrior logs settle S1 and P5.

## For the guide lane

- Nothing. No build, APL or card change.

## Still open

- S1: Heroic Strike as an auto attack for procs (level-60 log).
- P5: Revenge and Shield Slam threat constants (log or blue post), carried since round 8.
- WARRIOR-13 (Improved Shield Wall) stays with the owner.
- Warcraft Logs is blocked (gcloud login). I did not work around it.
