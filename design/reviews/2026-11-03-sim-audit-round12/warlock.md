# Sim-accuracy round 12: warlock (2026-11-03)

Specs: Affliction, Demonology, Destruction and their demons, bands 20 to 60, bare and raid presets. Client 1.60.1.70753. wago.tools has no hotfix push for 70753 (`hotfix_push.push_totals` for Spell, Talent, TraitDefinition, ItemEffect, ItemSparse: 0 entries, read 2026-11-03 08:00 local), so the shipped tables are the layered ones (round 11's copies in scratch `r11/c70753/builds/1.60.1.70753/raw`). Pin 2606bc05e, site c6bdc855c. Lane engine `lane/r12-warlock`, site branch `wt/r12-warlock`.

Method: the round-11 report first, then the assigned items, then a fresh pass. Every spell the three APLs cast and every talent in the published builds (data/builds/1.60.1.70753/bis/warlock-*.json, decoded against talents/warlock.json) was read against SpellEffect, SpellMisc (attributes, duration, cast time, missile speed), SpellCooldowns, SpellPower, SpellLevels, SpellAuraOptions, and the class masks resolved bit by bit through SpellClassOptions set 5. I grepped rounds 1 to 11 for each spell before calling it new. Measurements used spec-breakdown at 10,000 iterations, seeds 1 and 50001.

Totals: 1 CONFIRMED-BUG, fixed. Origin: 1 old. It moves every warlock band by 0.0 to 0.4% (down). No published number outside the warlock group moves.

## Assigned: the landing's APL lines

The only warlock APL change in 8d0f78ea0..c6bdc855c is round 11's own ac4937071: 11 lines per spec, Curse of the Elements 1311680 "rank 4" before the bane. No guide lane touched the warlocks (fork 47df4a89d..2606bc05e has no warlock commit).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Curse of the Elements line, all three specs. Ranks by band match the trainer: 440892 at 20, 1311676 at 30, 1311677 at 40, 1311680 at 50 and 60 (the line's "rank 4" resolves to the band's top learned rank; spec-breakdown shows each id cast). The client: 5 min (SpellDuration 300000), GCD 1.5 s, instant, 50/100/150/200 mana, effect 0 aura 22 -30/-45/-60/-75 and effect 1 aura 87 +4/6/8/10% on misc 126 (every magic school), "Only one Curse per Warlock". The engine (`curses.go:211-276`, constants_auto_gen.go:824-831) matches every number. It is cast 1.1 times a fight (a miss is recast) and never in the raid preset, where the raid's curse is up (`CurseOfElementsOn`). | SpellEffect, SpellMisc, SpellPower, SpellLevels for 440892/1311676/1311677/1311680; trainables/warlock.json; spec-breakdown bare 20 to 60 | Aligned | - | - |
| L2 | The notes beside it: "sits beside your bane for five minutes; two copies do not stack" matches the client text and duration. The generated copies (sim/request/apl/warlock-*.apl.json) also carry the guide lanes' note rewrites; each was read. "Corruption ... a 2 second cast without Improved Corruption" (Demonology) matches 11672's cast time 2000; "Soul Fire every 6 seconds while Decimation is up ... 2.4 second Soul Fires that cost no shard" matches 440870 (-90% of 60 s) and 440873 (-40% after Bane's -2 s, aura 256 no reagent). | SpellMisc 11672; 440870, 440873 | Aligned | - | - |

## Assigned: Orc Command on demons (70753)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | 70753 adds a second effect to the warlock's Command 20575 (SpellEffect 697943): aura 108 (percent modifier), misc 0 (damage), +5, on class mask 12288, which is Firebolt (bit 12) and Lash of Pain (bit 13) only. Effect 0 was already there: aura 107 misc 8, +5 on mask bit 26, the demons' "Tamed Pet Passive (DND)" 18727 to 18730 (aura 79, all schools, base 0). Unholy Power 18769 raises the same passive (+10 at rank 5), which confirms the passive is the demons' damage aura. So in the client an orc's demon deals +5% on everything through the passive, added to Unholy Power on the same aura (1 + 0.10 + 0.05 = 1.15), and Firebolt and Lash of Pain get a further +5% as a spell modifier. The engine (`core/racials.go:245-260`) multiplies every demon's damage by 1.05 on top of Unholy Power's own multiplier (1.10 x 1.05 = 1.155) and has no Firebolt/Lash term. For an orc with Unholy Power 5: Succubus melee should be 1.15 (engine 1.155, 0.4% high), Lash of Pain and Firebolt 1.15 x 1.05 = 1.2075 (engine 1.155, 4.5% low). On the raid sims' pet shares that nets about +0.04% (Demonology: melee 62.7 DPS -0.27, Lash 11.2 DPS +0.51) to +0.06% (Affliction); the Imp is never out in a published rotation. No published warlock is an orc (troll and gnome), so nothing published moves. The auras lane owns racials.go: see "For other lanes". | SpellEffect 697943 (new in 70753, diffed against 70707), 20575 effect 0, 18727-18730, 18769; SpellClassOptions set 5 | CONFIRMED in the client; fix belongs to auras | new-data | - |

## Assigned: held under the adoption rule

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Affliction raid: 1 Improved Drains to Improved Bane of Agony (25520010135001050-0005200000000000000-0550004000000000), horde, 10,000 iterations: 572.1 to 576.2 (seed 1), 572.2 to 576.3 (seed 50001), +0.72% both, error ±0.3. Under 1%. Held. (The coordinator's Shadow-priest ruling, moving a point out of a talent with no effect, could apply: no rotation casts a drain. That call is the coordinator's, not this lane's.) | spec-breakdown -talents | Aligned (held) | - | - |
| H2 | Destruction raid: Immolate line above Conflagrate (lines 6 and 7 swapped), horde: 638.0 to 639.6 (seed 1), 638.2 to 639.4 (seed 50001), +0.25% / +0.19%. Beyond error, under 1%. Held. | spec-breakdown -rotation | Aligned (held) | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Shadow Bolt landed on the cast; in the client it travels.** Every Shadow Bolt rank (686 to 25307) has SpellMisc.Speed 20 yards a second. The engine gave it no missile speed, while Soul Fire and Death Coil (24) and other classes' bolts (Fireball 24, Frostbolt 28, Wrath 20, Lightning Bolt 20) already travel. From the ranker's 30 yards a bolt now lands 1.5 s after the cast, so a bolt still flying when the fight ends deals nothing, and Improved Shadow Bolt and Decimation go up when it lands (both trigger on the hit: 17793 and 440870 proc on landing spell damage). Test `TestShadowBoltAndIncinerateTravelAtTheClientSpeed` failed first (0 s, want 1.5 s) and passes. The Incinerate damage test now steps the sim until the bolt lands. | SpellMisc.Speed for 686, 695, 705, 1088, 1106, 7641, 11659, 11660, 11661, 25307 (all 20); `shadowbolt.go`, `core/spell.go:671` (TravelTime = distance / MissileSpeed), `sim/leveling/distance.go:50` (casters at 30 yd) | CONFIRMED-BUG (client column; the engine models the same column for Soul Fire, Death Coil and other classes) | old: wowsims classic's Shadow Bolt never had it; no round checked missile speeds | E 1b8bd1d7c |
| A2 | Spells against the client: Shadow Bolt r9 11661 (251 + 1.8 a level, 0.857, 3.0 s, 370), Corruption r6 11672 (57 every 3 s for 18 s, 0.20, 2.0 s, 290), Bane of Agony r6 11713 (46 every 2 s for 24 s, 0.133, 215), Siphon Life r4 18881 (41 every 3 s for 30 s, 0.05, 365), Immolate r7 11668 (146 + 2.3 a level at 0.20, then 52 every 3 s for 15 s at 0.13, 2.0 s, 370), Life Tap r6 11689 (420 + 1 a level from 56, + Spirit, x1.2: `lifetap.go:27-44`). Crit: Corruption, Agony, Siphon Life and Doom carry PERIODIC_CAN_CRIT (Attributes_8 0x200) and are non-channelled DoTs anyway. Ranks per band are the band's top trainable ranks (bare 20 to 50 casts read off spec-breakdown). | client rows; engine tables | Aligned | - | - |
| A3 | Talents, masks resolved bit by bit: Suppression (aura 54/55, mask 0), Improved Corruption (-2 s; +10% on misc 22, periodic), Improved Bane of Agony (misc 22, Agony), Pandemic (misc 15 crit damage, 33/67/100% on Corruption, Agony, Doom, the drains, Siphon Life, Wrack), Malevolence (misc 7 crit chance on Shadow Bolt, Corruption, Drain Life, Shadowburn, Agony, Drain Soul, Death Coil, Lash of Pain, Siphon Life, Doom, Wrack, Havoc), Nightfall (2% a rank, 17941 10 s), Shadow Mastery (misc 0 direct on Shadow Bolt, Corruption, Shadowburn, Death Coil; misc 22 periodic on the banes, Corruption, the drains, Siphon Life, Wrack), Amplify Curse (round 9 F2: the masks name Weakness and Exhaustion, the text and Blizzard's replies Agony). All match `talents.go`. | 18174, 17810, 18827, 427712, 1310949, 18094, 18271, 18288 | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | A1's Shadow Bolt fix moves Demonology the most (-0.25 to -0.44% at raid 60): Shadow Bolt is 41% of its damage, and Decimation's Soul Fire window now opens when the bolt lands. | spec-breakdown | follow-on of A1 | - | E 1b8bd1d7c |
| D2 | Demonic Sacrifice: Imp 18789 "Burning Shadow" +15% Shadow (aura 79 misc 32), Succubus 18791 "Touch of Fire" +15% Fire; Soul Link 25228 aura 79 +3% for both; Master Demonologist 2% a rank (bp 10 at 5); Demonic Knowledge 100% of level; Unholy Power +10 on the demons' passive; Improved Sayaad misc 8 on Lash of Pain; Fel Vitality 15%; Master Summoner -4 s and -40%; Decimation 440870/440873 (-90% cooldown, -40% cast, 6% at 35%). Summon costs 80% (Imp) and 100% (Succubus) base mana, Soul Link 20%: moving the prepull summons 4 minutes earlier changes nothing (639.3 both), so the prepull's mana does not bind. | client rows; `talents.go:426-776`; spec-breakdown with the prepull at -300 s | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Incinerate landed on the cast.** All three ranks (412758, 1293812, 1293813) have SpellMisc.Speed 20, the same as Shadow Bolt; the engine dealt the damage at once. Fixed with A1, same test. Incinerate's +25% reads Immolate at the cast, as the engine's other bolts compute their hit at the cast. | SpellMisc.Speed for the three ranks; `incinerate.go` | CONFIRMED-BUG (with A1) | old (wowsims' SoD Incinerate had 24; 620976647, 2024-11-22, removed it with the SoD runes and the Forever file was rebuilt without it) | E 1b8bd1d7c |
| X2 | Talent masks: Agonizing Flames (misc 7 Searing Pain crit; misc 0 on Shadow Bolt, Immolate, Rain of Fire, Shadowburn, Searing Pain, Incinerate, Soul Fire, Conflagrate; misc 22 on Immolate and Rain of Fire, so Immolate's ticks get it too, as the engine's whole-spell multiplier gives), Aftermath (misc 0 on Immolate: initial hit only, as `immolate.go:110` does), Ruin and Cataclysm (masks equal the engine's Destruction flag for every cast spell), Bane (-0.5 s Shadow Bolt, Immolate, Incinerate; -2 s Soul Fire), Fire and Brimstone (8/17/25%), Shadow and Flame (2% a rank for 20 s, 20% a rank Immolate kept). Conflagrate 18932 (282 + 1.3 a level, 0.429, 10 s), Shadowburn 18871 (266 + 1.8 a level, 0.429, 15 s), Bane of Doom 603 (1742, coefficient 4.0, 60 s, 60 s cooldown, PERIODIC_CAN_CRIT). | client rows; `talents.go:843-1326` | Aligned | - | - |
| X3 | Eureka! (Gnome, the alliance race) for warlocks, 1259821: its damage mask names Shadow Bolt, Immolate, Rain of Fire, Shadowburn, Searing Pain, Death Coil, Soul Fire and Conflagrate, not Incinerate (word 1 bit 6). `core/eureka.go` builds the list from the client masks. Troll Berserking 20554: 10% cast and attack speed, 15 s, 3 min. | 1259821, 20554 | Aligned | - | - |

## Demons

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | The warlock's spell modifiers reach the demon's spells where their masks name them (the server applies an owner's spell mods to its pet's spells): Malevolence and Improved Sayaad on Lash of Pain (Lash crits 10.5% for Affliction, 5.6% without Malevolence). Lash of Pain 11780: 50 flat, 0.429, 12 s, 160 mana. The Imp's Firebolt has client speed 16, but the Imp is never out in a published rotation; not changed. | 11780, 1310949, 18754; spec-breakdown | Aligned | - | - |
| P2 | Round 10's open demon rows (crit at 60, the sheets, the warlock's crit reaching demons): no level-60 log yet; api.foreversixty.gg holds no level-60 report. | | OPEN | - | - |

## Gear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | 84 published items across the warlock bands. Those with an effect beyond stats: Kindling Stave (Ignition), Frozen Heart of the Mountain (Chthonic Power), Rune of the Guard Captain (Authority, attack power), Robe of the Void (Heal Pet), Mark of the Chosen and Uther's Strength (when struck), Cookie's Stirring Rod (cooking), Archmagister's Faceted Pendant (experience). Unchanged from round 11; no published item is in a set. | ItemXItemEffect, ItemEffect, Spell; bis/*.json | Aligned | - | - |

## Before/after

10,000 iterations, seeds 1 and 50001. Before: `$R/sb-before` from `$R/site-before/sim` (c6bdc855c on 2606bc05e). After: wt/r12-warlock on lane/r12-warlock 1b8bd1d7c. Seed 1 absolute, then % at seed 1 / 50001.

| Spec, band, preset | Horde s1 | Horde % | Alliance s1 | Alliance % |
|---|---|---|---|---|
| Affliction 20 bare | 54.7 to 54.6 | -0.18 / -0.37 | 55.0 to 54.9 | -0.18 / -0.18 |
| Affliction 30 bare | 90.1 to 89.9 | -0.22 / -0.22 | 90.5 to 90.2 | -0.33 / -0.33 |
| Affliction 40 bare | 140.1 to 139.7 | -0.29 / -0.21 | 140.1 to 139.8 | -0.21 / -0.21 |
| Affliction 50 bare | 207.5 to 207.1 | -0.19 / -0.24 | 209.3 to 208.9 | -0.19 / -0.29 |
| Affliction 60 bare | 307.9 to 307.3 | -0.19 / -0.32 | 306.6 to 305.9 | -0.23 / -0.33 |
| Affliction 60 raid | 572.1 to 571.9 | -0.03 / -0.09 | 570.2 to 570.0 | -0.04 / -0.09 |
| Demonology 20 bare | 48.4 to 48.3 | -0.21 / -0.21 | 49.8 to 49.6 | -0.40 / -0.40 |
| Demonology 30 bare | 82.2 to 82.0 | -0.24 / -0.24 | 82.1 to 81.9 | -0.24 / -0.24 |
| Demonology 40 bare | 148.1 to 147.7 | -0.27 / -0.34 | 148.1 to 147.7 | -0.27 / -0.34 |
| Demonology 50 bare | 222.3 to 221.9 | -0.18 / -0.22 | 222.0 to 221.5 | -0.23 / -0.14 |
| Demonology 60 bare | 347.1 to 347.0 | -0.03 / -0.03 | 347.1 to 346.8 | -0.09 / -0.14 |
| Demonology 60 raid | 641.4 to 639.3 | -0.33 / -0.25 | 636.5 to 635.5 | -0.16 / -0.44 |
| Destruction 20 bare | 53.2 to 53.0 | -0.38 / -0.19 | 53.4 to 53.2 | -0.37 / -0.37 |
| Destruction 30 bare | 91.1 to 91.0 | -0.11 / -0.22 | 91.7 to 91.5 | -0.22 / -0.22 |
| Destruction 40 bare | 147.2 to 147.0 | -0.14 / -0.14 | 148.0 to 147.8 | -0.14 / -0.14 |
| Destruction 50 bare | 251.5 to 250.9 | -0.24 / -0.20 | 253.4 to 252.9 | -0.20 / -0.12 |
| Destruction 60 bare | 378.3 to 377.6 | -0.19 / -0.21 | 382.0 to 381.1 | -0.24 / -0.24 |
| Destruction 60 raid | 638.0 to 636.2 | -0.28 / -0.30 | 635.0 to 634.0 | -0.16 / -0.27 |

The fork's `GOWORK=off go test -tags=with_db ./sim/warlock/...` passes; no fork golden moved (the fork's warlock DPS tests stand at 0 yd).

## Files to regenerate

- sim/request/testdata/ladder/warlock-affliction.golden.md, warlock-demonology.golden.md, warlock-destruction.golden.md (TestRotationLadder): every rung moves slightly.

## Measure at refresh

- The warlock bands' weights and picks after the bolt travel time (A1, X1): spell power and hit lose a little value at the end of each fight; picks should not move.

## For other lanes

- **Auras (racials.go, Orc Command, O1):** in 70753 the warlock's Command 20575 is two effects. Effect 0 adds +5 to the demons' damage passive (18727-18730, aura 79), the same aura Unholy Power adds +10 to, so the two add: 1.15 with Unholy Power 5, not the engine's 1.10 x 1.05. Effect 1 (new, SpellEffect 697943) is +5% damage on Firebolt and Lash of Pain only (mask 12288). Wanted for an orc warlock with Unholy Power 5: demon melee x1.15, Lash of Pain and Firebolt x1.2075; the engine gives x1.155 to both. Net about +0.04 to +0.06% for an orc warlock at raid 60; no published warlock is an orc. The hunter's 20576 effect 0 is the same pattern on class set 9, for the physical lane to check against the hunter pet passive.
- **Other casters and hunters (missile speeds, not confirmed by this lane):** APL-cast spells whose engine missile speed differs from SpellMisc.Speed in 70753: hunter Aimed Shot 20904, Arcane Shot 14287, Serpent Sting 13555 (client 40, engine 24) and Multi-Shot 2643 (client 30, engine 24); paladin Hammer of Wrath 24239 (client 35, engine none); the wand Shoot 5019 (client 20, `core/wand.go:61` 24); Holy priest's Prayer of Mending 1240827 (client 40, a heal). Mage, druid, shaman and warlock bolts now match.
- **Integrator:** E 1b8bd1d7c moves only the warlock specs (it touches sim/warlock/shadowbolt.go and incinerate.go). Regenerate the three warlock ladder goldens.

## Still open

- From earlier rounds, unchanged: demon crit and sheets at 60 and the warlock's crit reaching demons (P2); additive versus multiplied percent modifiers (Aftermath, Improved Corruption); Ignition's rate; Siphon Life tick crit (held for the owner); shadow wands (#419).
- H1 and H2 stay held.
