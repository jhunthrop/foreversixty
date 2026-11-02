# Sim audit round 11: healers

Lane: wt/r11-healers (site, from origin/main 8f047a62f) and lane/r11-healers (fork, from cfcb76e80). Specs: Restoration druid, Restoration shaman, Holy priest, Discipline priest, Holy paladin, at 60 (raid and bare) and at the bare bands 20 to 50. The sub-60 heal profile, its pull-size multiplier and Restoration druid's band-20 Furor belong to the healprofile lane; Conviction on paladin heals belongs to the crit lane. Neither is touched here.

Client truth: 1.60.1.70707 with the hotfix push layered in (the raw CSVs fetched 2026-11-01 21:42 UTC). `push_totals` on 2026-11-02 found no new push for any spell, item or talent table since then (only a Faction count the snapshot did not track).

Method. Round 10's four class sweeps (design/reviews/2026-11-03-sim-audit-round10/healers/aud-*.md) re-read every cast rank, talent curve and mask, the 60 gear, the racials and the consumables against the client. This round I did not repeat those static reads row by row. I tested them against new evidence instead, mostly the beta logs:
- **Event-level beta logs**: 2,256 healer encounter timelines (round 10's fetch, levels about 20 to 35), plus 168 new damage timelines for twelve casters (2026-11-02).
- **The sub-60 bands**: every cast list at bands 20 to 50, every published item at every band, the band builds.
- **Mana**: the regen formulas, the five-second rule, base mana by level.
- **The heal profile's tank split and overheal**, against the logs.

The scripts are in [healers/logtools/](healers/logtools/).

The short version:
- **Heals learned below level 20 kept only part of their coefficient. Forever's logs say they keep all of it.** Round 8 added the server cores' sub-20 penalty on core evidence alone. Seven characters' two-rank or three-rank fits (four classes' heals: Rejuvenation, Healing Touch, Healing Wave, Holy Light) agree on one healing power and one talent multiplier only with the whole coefficient. With the penalty the fitted multiplier falls to 0.50 to 0.81, below any talent's. Fixed for heals; damage spells keep the rule for now (one weak data point, for the casters lanes).
- **Forever's healer potion was missing.** Mender's Potion and Major Mender's Potion (+31 and +88 healing for 30 s) are the third family of round 10's 30-second potions. Major Mana Potion still measures 3.5 to 10% ahead for every raid healer, so no preset changes.
- **Round 10's 30-second potions never fired for a healer.** They waited for an enemy target, and a healer targets a friend. Latent: no healer carried one.
- **Aligned, now with logs behind the client rows:**
  - Heal-over-time durations and tick counts (Rejuvenation 12 s and 4 ticks, Regrowth 21 s and 7, Renew 15 s and 5).
  - Penance's three pulses, one a second, the first on the cast.
  - Holy Shock's heal scaling with Healing Light.
  - The heal profile's tank share and overheal.

## Confirmed and fixed

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **A heal learned below level 20 lost part of its spell power coefficient.** The engine scaled every cast rank learned at level L < 20 by 1 - (20 - L) x 3.75%. Examples: Rejuvenation rank 3 (level 16) 0.85; Holy Light rank 3 (14) 0.775; Healing Wave rank 1 (6) 0.475; Heal rank 1 (16) 0.85. That rule is cmangos' and vmangos', and round 8 added it on that evidence alone (its report flagged it unverified).<br><br>**The test.** Two ranks of one heal from one character share every talent multiplier. So their mean raw non-crit amounts fix both the healing power H and the multiplier M. A true model gives an M of at least 1, and an H above the gear's.<br><br>**Results, each pair or set from one character:**<br>- Rejuvenation ranks 3 and 5 (16 and 28; no per-level growth, so the level guess cannot matter). Aye ticks 59.6 / 81.7, Flody 60.8 / 83.7. Without the penalty: M 1.16 and 1.20, H 142 and 138. With it: M 0.77 and 0.81.<br>- Healing Touch ranks 2 and 5 (Im, 222 / 857): M 1.24 without, 0.67 with.<br>- Healing Wave ranks 1 and 4 (Xum, 140 / 510): 1.24 without, 0.58 with.<br>- Holy Light ranks 3 and 4 (Cessair): 0.94 without, 0.50 with. Ranks 3 to 5 (Reneth): residuals ±5 without, ±11 with.<br>- Dorion, three spells with one multiplier (Healing Light reaches all three). Flash of Light rank 3 (183.7, 168 heals) and the Holy Shock heal (233.2) give M 1.19 and H 190. Then Holy Light rank 3 (level 14) should be 356.8 without the penalty and 320.4 with it. Observed: 355.1 over 54 heals.<br><br>**Damage.** Of twelve casters, only one cast two ranks of one damage spell (a mage's Arcane Explosion ranks 2 and 3: M 0.96 without, 0.68 with). That is too little to move damage spells, so they keep the rule.<br><br>**The fix.** `levelPenalty` returns 1 for a spell with the healing proc mask. The heal rank tests and the conformance report follow. | `sim/core/spell.go` levelPenalty; client SpellEffect/SpellLevels for 1430, 2091, 5186, 5189, 332, 939, 647, 1026, 1042, 19939, 1311605. foreverlogs.gg reports 2747, 2748, 2752, 2753, 2701, 2710, 2811 (timelines fetched 2026-11-02); `logtools/pen.py`, `paladin3.py`, `pendmg.py`. Test `TestAHealLearnedBelowTwentyKeepsItsCoefficient` failed first (0.68 and 0.17 at level 16, 0.23 and 0.0575 at level 1). | CONFIRMED-BUG (logs, seven characters, four classes) | old (round 8, E 87743552b) | E 074edfadf |
| C2 | **Forever's healer potion was not modelled.** Round 10 added Frenzy and Spellblasting. The third family of Forever's 30-second potions is the healer's:<br>- Mender's Potion 250946 (spell 1251945): aura 135, +31 Spell Healing.<br>- Major Mender's Potion 250949 (1251948): +88.<br>Both last 30 s (DurationIndex 9) in potion category 4 (120,000 ms). The recipes are vendor-sold like Spellblasting's (Wowhead Forever lists both Major recipes from a vendor, at skill 280 and 285). Added as Potions 27 and 28.<br>Measured against Major Mana Potion at raid 60 (horde, 3,000 iterations, seeds 1 and 50001; three Mender's Potions a fight):<br>- Restoration druid -3.4%<br>- Restoration shaman -7.8%<br>- Holy priest -9.9%<br>- Discipline -8.2%<br>- Holy paladin -5.3%<br>The healers run their mana down, so the mana potion keeps every slot and no preset changes. | Client rows above; https://nether.wowhead.com/forever/tooltip/item/250949 "Increases Spell Healing by 88 for 30 sec." (read 2026-11-02). `TestBurstPotionsMatchTheClient` extended first (did not build). | CONFIRMED (missing consumable; no published change) | old (round 10's potion pass looked at the family and missed it) | E 89da891cc, S 6c9a09b4a (IDS.md) |
| C3 | **A 30-second potion never fired for a healer.** The burst potion was a DPS cooldown that is neither helpful nor a self buff. Such a cooldown waits for an enemy target, and a healer's target is a friend, so a Mender's Potion was never drunk (0 casts, identical HPS with or without it). It is now a self buff. Damage specs and tanks are unchanged: their target is the enemy. | `sim/core/consumes.go` makeBurstPotionMCD; `major_cooldown.go` landsOnTarget. `TestBurstPotionIsADPSCooldownOnThePotionTimer` gained the self-buff check and failed first. | CONFIRMED-BUG (latent: no healer carried one) | R10 E 61734a1cd | E 89da891cc |

## Restoration druid

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Sub-20 coefficient (C1): Rejuvenation rank 3 and Regrowth rank 2 at band 20 now keep their whole coefficient. | C1 | CONFIRMED-BUG | old | E 074edfadf |
| D2 | **Heal-over-time durations and ticks, from the logs.** Per application with no recast between:<br>- Rejuvenation ranks 2 to 5: 12.0 s, 4 ticks (852 applications).<br>- Regrowth ranks 2 to 4: 21.0 s, 7 ticks (167).<br>- Tranquility rank 1: 5 ticks at most; it is a channel, often cut short.<br>These match the client durations (12000, 21000, 10000) and the engine's tick counts. No logged druid has Nature's Splendor. | `logtools/hotdur.py` | Aligned (logs + client) | - | - |
| D3 | **Bloom's heal (1328622) has no SpellCategories row** (DefenseType 0), like Lightwell Renew (round 10 P3). The engine lets it crit. It has no class mask, so Gift of Nature does not reach it (the engine's mask mods agree). The 70707 text matches the engine: 100% of Spirit on a natural expiry. No log has Bloom: the talent arrived with 70707. | client rows; https://nether.wowhead.com/forever/tooltip/spell/1328621 (read 2026-11-02) | OPEN (one logged Bloom crit, or its absence) | - | - |
| D4 | **Bands 20 to 50, re-read:**<br>- The casts: Rejuvenation 1430 / 2090 / 2090 / 9839, Regrowth 8938 / 8940 / 8941 / 9856, Tranquility 740 / 8918 / 9862, Swiftmend and Innervate from 50.<br>- Each band's learned ranks; Innervate learned at 40. At band 40 Innervate is never cast because mana outlasts the fight (2,933 s).<br>- Base mana by level against the client BaseMP (354 at 20, matching).<br>- The published gear's effects: the Staff of Westfall and Gnarled Necromancer's Staff equips are aura 129 (run speed); Ankh of Life from 40, as at 60; Evonice's is feather fall.<br>- The night elf and tauren racials. | `logtools/itemfx.py`; spec-breakdown bands 20 to 50 | Aligned | - | - |
| D5 | Rejuvenation's same-op stacking (round 10 D3) stays open. Rosters carry tree totals only, not individual talents, so the logs cannot split Gift of Nature from Improved Rejuvenation. | roster snapshots | OPEN | - | - |

## Restoration shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Sub-20 coefficient (C1): Healing Wave ranks 1 to 4 (the band-20 cast is rank 4, level 18). | C1 (Xum's ranks 1 and 4) | CONFIRMED-BUG | old | E 074edfadf |
| S2 | **Mana Tide Totem is a tier-3 talent in Forever, not tier 7.** A 16-point Restoration shaman (Xum, report 2710) casts rank 1 (16190). The client tree puts it at tier 3, and SkillLineAbility trains ranks 2 to 4 at 38, 48 and 58. The band-40 build takes it, at rank 2 there. It is never cast at band 40: the list drops it under 25% mana, and mana outlasts the fight (1,471 s). Band 50 casts rank 3 0.2 times a fight. | talents/shaman.json; SkillLineAbility 9858, 9859, 59647; logs | Aligned | - | - |
| S3 | **Bands 20 to 50, re-read:**<br>- Healing Wave 913 / 939 / 8005, Chain Heal 10622 and Riptide 1239242 at 50, Healing Stream ranks 1 to 4 (5 + 0.022 x healing a tick for rank 1, matching 5672).<br>- Water Shield (a tier-2 talent, so from band 30).<br>- The gear's effects; dwarf and tauren racials. | client rows; spec-breakdown | Aligned | - | - |
| S4 | Chain Heal's jumps can't be checked in the logs: no logged shaman is level 40. Purification and Healing Way stacking stay open, for the same roster reason as D5. | logs | OPEN | - | - |

## Holy and Discipline priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Sub-20 coefficient (C1).** It reaches:<br>- Heal rank 1 (level 16): band 20's cast for both specs, and Discipline's authored downrank in its list, cast at raid 60 (rank-choices.json lifts it to rank 4 at bare 60).<br>- Renew ranks 1 and 2.<br>- Lesser Heal; Divine Grace's level-10 rank.<br>Discipline at raid 60 gains from it (see Before/after). Power Word: Shield never took the penalty: its absorb reads the rank's coefficient directly. | C1; conformance priest.golden.md | CONFIRMED-BUG | old | E 074edfadf |
| P2 | **Penance is three pulses one second apart, the first on the cast.** The channel spells (402277, 1240732, 1240733, 1316992) last 2 s with a 1 s period and carry Attributes_5 0x200 (a tick on application). In the logs, 34 of 40 volleys from a level-30 Discipline priest (Altu, report 2747) are 3 pulses 1.0 s apart; the rest are cut short. This matches `penance.go` (an instant pulse plus 2 ticks). | SpellMisc rows; `logtools/penance.py` | Aligned (logs + client) | - | - |
| P3 | **Renew lasts 15 s for 5 ticks** (rank 2, 23 clean applications). | `logtools/hotdur.py` | Aligned (logs) | - | - |
| P4 | **Spirit of Redemption has no passive in Forever.** 20711 is a single dummy effect, and the server text is the death form only. The engine gives it nothing. | client row; https://nether.wowhead.com/forever/tooltip/spell/20711 (read 2026-11-02) | Aligned (FOREVER-CHANGE from Era's +5% Spirit) | - | - |
| P5 | **Lightwell.** The 70707 server text matches the engine: "restore 1600 health over 10 sec. Being attacked cancels the effect. Lightwell lasts for 3 min or 5 charges"; Lightwell Renew restores 320 every 2 s. Its crit stays round 10's open P3. | https://nether.wowhead.com/forever/tooltip/spell/27871, /27874 (read 2026-11-02) | Aligned | - | - |
| P6 | **Holy Binding** (1328569): proc 100 on helpful magic spells (ProcTypeMask 0x4000). The server text is "healed for 50% as much", Flash Heal only, not on the priest. This matches `applyHolyBinding`. | client row and tooltip | Aligned | - | - |
| P7 | **Dark Sacrifice, ranks 1 to 4 at bands 20 to 50** (1277324 to 1277327): 80 / 136 / 196 / 258 health and mana every 3 s, cast at each band at its own rank. Bands 20 to 50 cast nothing a band cannot have learned: Prayer of Mending from 40, Penance from 40, Binding Heal at 50. | client rows; spec-breakdown | Aligned | - | - |
| P8 | Round 10's open P3 to P6 (Lightwell Renew's crit, Renewed Hope per pulse, Sudden Insight per target, Litany across ranks) stay open. The logs have no Renewed Hope (a 31-point talent), no Lightwell and no Devout set. | - | OPEN | - | - |

## Holy paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Sub-20 coefficient (C1): Holy Light ranks 1 to 3. Today's band casts are Flash of Light rank 1 (level 20) and higher, so no published paladin number moves. | C1 (Cessair, Reneth, Dorion) | CONFIRMED-BUG | old | E 074edfadf |
| H2 | **Holy Shock's heal takes Healing Light**, as the engine applies it. Dorion's Flash of Light rank 3, Holy Shock heal and Holy Light rank 3 fit one multiplier and one healing power. Reneth's Holy Shock heal (220.5) matches (114 + 0.429 x 170) x 1.173 from his Holy Light fit. | `logtools/paladin3.py` | Aligned (logs) | - | - |
| H3 | **Bands 20 to 50:**<br>- Flash of Light rank 1 carries bands 20 to 50: the fights are easy, and rank 1 is the most healing per mana.<br>- Holy Light and Holy Shock (rank 1, level 30) appear at 50.<br>- Lay on Hands takes the band's rank (633 / 2800 / 10310; 0 / 250 / 550 mana back).<br>- Verigan's Fist at band 20 carries a melee-hit damage proc (1322218) that a healer never triggers.<br>- Blessing of Light stays removed (round 6). | client rows; spec-breakdown | Aligned | - | - |
| H4 | Holy Power's crit size, Oath of Light's scaling, Illumination's refund on Holy Shock rank 1: still open. The logged Holy Shock users are 26-point paladins without Holy Power. The energize events that would show Illumination's refund cannot be fetched: foreverlogs.gg answers `event_types=energize` with a server error. | logs; foreverlogs.gg API, 2026-11-02 | OPEN | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Spirit regen formulas.** Every 2 s tick:<br>- druid and paladin: 15 + Spirit/5<br>- shaman: 17 + Spirit/5<br>- priest: 12.5 + Spirit/4<br>These are cmangos-classic's and vmangos' OCTRegenMPPerSpirit. No client table states them, and the logs carry no mana. | `core/mana.go:157`, `priest/priest.go:250`, `shaman/shaman.go:34` | Aligned (server cores only; single source) | - | - |
| X2 | **The five-second rule.** Mana is spent when a cast completes and the rule runs five seconds from then (`core/cast.go` hardcast OnComplete, `mana.go` SpendCost). A free cast (Clearcasting, Inner Focus) does not restart it. Intellect gives 15 mana a point past the first 20, and base mana matches the client's BaseMP at 20, 30, 40, 50 and 60. | code; gametables/basemp.txt | Aligned | - | - |
| X3 | **The heal profile's tank share and overheal, against the beta logs.**<br>- Tank share. In 99 logged encounters (5-player dungeons, levels 20 to 35), the most-healed target takes 50 to 60% of a healer's raw healing (median: druid 0.52, priest 0.50, shaman and paladin 0.60). The 60 profile asks about 66% on the tank (645 tank damage a second against 333 group damage a second).<br>- Overheal. The logs: druid 11%, shaman 10%, priest 8%, paladin 5%. Rejuvenation is the highest single spell at 18%. The sim at raid 60: 0.5 to 10%.<br>Both sit close. The profile is sized on purpose so healing falls short of the damage, which keeps overheal low. Neither is a bug. The bands' high overheal (40 to 50% at bare 30 and 40) is the healprofile lane's easy-band problem. | `logtools/split.py`; spec-breakdown | Aligned (judged; the logs are dungeons, not raids) | - | - |
| X4 | **Every published healer item at every band** (125 entries), through ItemXItemEffect: the only effects are Ankh of Life, Uther's Strength, Truefaith, Hand of Edward, Evonice's, the run-speed staves and Verigan's Fist. Ironfeather Shoulders and Stormcloth Vest are single pieces of their sets. All were already handled or have no healing value. | `logtools/itemfx.py` | Aligned | - | - |

## Before/after

Before: `$R/sb-before` (origin/main 8f047a62f on cfcb76e80). After: lane/r11-healers 89da891cc with wt/r11-healers. Effective HPS on each band's heal profile, 10,000 iterations.

| spec | case | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|---|
| Restoration druid | raid 60 | +0.0% | +0.0% | +0.0% | +0.0% | 830.0 → 830.0 |
| Restoration druid | bare 20 | - | - | +1.6% | +1.6% | 31.1 → 31.6 |
| Restoration druid | bare 30 / 40 / 50 / 60 | - | - | 0.0% | 0.0% | 54.1 / 63.1 / 239.1 / 412.5, unchanged |
| Restoration shaman | raid 60 | +0.0% | +0.0% | +0.0% | +0.0% | 834.0 → 834.0 |
| Restoration shaman | bare 20 / 30 / 40 / 50 | - | - | 0.0% | 0.0% | 38.3 / 54.2 / 65.1 / 245.3, unchanged |
| Restoration shaman | bare 60 | - | - | +9.2% | +9.1% | 373.1 → 407.4 |
| Holy priest | raid 60 | +0.0% | +0.0% | +0.0% | +0.0% | 693.8 → 693.8 |
| Holy priest | bare 20 | - | - | +0.3% | +0.6% | 32.8 → 32.9 |
| Holy priest | bare 30 / 40 / 50 / 60 | - | - | 0.0% | 0.0% | unchanged |
| Discipline priest | raid 60 | +1.8% | +1.8% | +1.8% | +1.8% | 625.3 → 636.6 |
| Discipline priest | bare 20 | - | - | +0.3% | +0.3% | 28.7 → 28.8 |
| Discipline priest | bare 50 | - | - | +0.1% | +0.1% | 192.8 → 192.9 |
| Discipline priest | bare 30 / 40 / 60 | - | - | 0.0% | 0.0% | unchanged |
| Holy paladin | every case | 0.0% | 0.0% | 0.0% | 0.0% | 641.4 raid, unchanged |

Where it comes from: every move is C1. The Mender's potion is carried by no preset, and C3 moves no published number.
- **Restoration shaman bare 60 +9.2%:** its curated rank choice there is Healing Wave rank 3 (547, level 12), its main heal (94 casts a fight), which kept only 0.70 of its coefficient.
- **Discipline raid 60 +1.8%:** its list's authored Heal rank 1 (level 16).
- **Bands 20 to 50 move little:** today's easy band fights cap the healers (the healprofile lane's multiplier changes that).
- **Holy paladin:** casts no rank below 20.

A follow-on, not adopted: with C1, Discipline's Heal rank 1 lines cast as Lesser Heal rank 3 measure +2.0% at raid 60 (horde, 3,000 iterations, 636.4 → 649.4 at seed 1, 636.9 → 650.3 at 50001). It needs the adoption rule's 10,000-iteration, both-faction check (Still open).

## For other lanes

- **Ranks lane and healprofile lane (rank choices).** C1 makes every heal rank learned below 20 worth more per mana. Re-measure the sub-60 healer rows of rank-choices.json on the merged ladder (healprofile), and Discipline's authored Heal rank 1 at raid 60 (ranks). Lesser Heal rank 3 (level 10) gains the most: its coefficient goes from 0.446 to 0.714, so it may now beat Heal rank 1 as Discipline's mana-bound filler at 60. That is a rotation question under the adoption rule, not measured here. Round 9's downrank candidates assumed the penalty: Holy Light 639, Healing Wave 547, Heal 2054, Renew 139.
- **Casters, warlock and hybrids lanes (damage ranks).** Damage spells keep the sub-20 penalty. The only logged pair, a mage's Arcane Explosion ranks 2 and 3 (report 2706), leans the same way as the heals: M 0.96 without the penalty, 0.68 with it. A caster casting two ranks of Frostbolt, Shadow Bolt, Smite or Shadow Word: Pain across level 20 settles it (`logtools/pendmg.py` takes damage timelines).
- **Auras lane.** C2 and C3 touch `sim/core/consumes.go` (burstPotions, makeBurstPotionMCD) and add Potions 27 and 28 in `proto/common.proto`. A potion change of yours this round will conflict there; the numbers need renumbering if you also took 27.
- **Crit lane.** Bloom's heal (1328622) has DefenseType 0, like Lightwell Renew, and the engine lets it crit (D3).
- **Platform lane.** /sim's consumable panel can offer the Mender's potions once simbuffs.json is regenerated at the refresh.
- **Integrator.**
  - data/proto must follow the pin for Potions 27 and 28 (S 6c9a09b4a regenerated sim/request/IDS.md).
  - TestRotationLadder moves for every healer: C1 at the sub-60 bands, and Discipline at 60. Not regenerated here.

## Held-output tests

None deselected. `go test ./request/` fails only TestRotationLadder (the integrator's goldens) on this branch. The fork suite (`-tags=with_db`, every package but sim/web) passes. The conformance goldens for druid, paladin, priest and shaman were regenerated in E 074edfadf. `go vet ./...` in api passes.

## Still open

- Discipline's Lesser Heal rank 3 filler (+2.0% at 3,000 iterations after C1; not yet run through the adoption rule).

- Damage spells' sub-20 penalty (casters lanes, above).
- Bloom's crit (D3); Lightwell Renew's crit, Renewed Hope per pulse, Sudden Insight per target, Litany across ranks (round 10 P3 to P6).
- Holy Power, Oath of Light, Illumination on Holy Shock rank 1 (each needs a level-40+ paladin log; energize events are not fetchable today).
- Same-op stacking on Rejuvenation and on Healing Wave (needs per-talent rosters).
- The spirit regen formulas rest on server cores alone.
