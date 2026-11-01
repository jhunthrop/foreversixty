# Sim audit round 3: hybrids (2026-11-01)

Druid (Balance, Feral cat, Feral bear, Restoration), Shaman (Elemental, Enhancement, Restoration) and Paladin (Retribution, Protection, Holy).

**Setup:**
- Client 1.60.1.70707.
- Engine: pin c8d29ea33, plus fork branch `lane/r3-hybrids`.
- Site: branch `wt/r3-hybrids`.
- Round 2 is `design/reviews/2026-11-01-sim-audit-round2/hybrids.md`.

**How claims were checked:**
- Client tables: the raw 70707 wago CSVs (`SpellEffect`, `SpellClassOptions`, `SpellAuraOptions`, `SpellLabel`, `SpellAuraRestrictions`, `SpellProcsPerMinute`, `SpellCooldowns`), fetched 2026-11-01.
- Wowhead Forever tooltips: `nether.wowhead.com/forever/tooltip/spell/<id>`, fetched 2026-11-01.
- Beta logs: foreverlogs.gg's report pages (the JSON the site itself reads: `buff_uptimes`, `casts`, `character_abilities` with `scope=boss`), 156 druid and 9 Retribution characters from the round-2 report list, levels about 13 to 30.
- Sim runs: `sim/cmd/spec-breakdown` and `sim/cmd/talent-search`, raid preset, level-60 horde band unless noted, 10000 iterations at seeds 1 and 50001.

Fix-sha column: "E" is a fork commit on `lane/r3-hybrids`, "S" is a site commit on `wt/r3-hybrids`.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Judgement of the Crusader joins the raid preset** (owner ruling). | **Rank:** the preset is the level-60 band only (`sim/cmd/leveling-bis/raid_build.go`), and a level-60 paladin judges rank 6: Seal of the Crusader 20308 is learned at 52, its Judgement 20303 carries aura 14 at 161 (`spellranks.json`, `SpellEffect`). The engine's raid debuff applies 161 (`sim/core/debuffs.go:470`). No lower band carries the preset, so no other rank is needed.<br>**Improved form:** Forever's trees have no Improved Seal of the Crusader.<br>**Tests first:** `sim/request/preset_test.go` and `data/tests/test_presets.py` failed, then passed.<br>**Who moves:** a sweep of all 28 specs: only Retribution (+9.6%) and Protection paladin (+7.8% DPS, +9.3% threat). Shadow casts no Holy damage, and the Holy-damage healers (Holy and Discipline priest, Holy paladin) score healing, which does not move. | RULED, implemented | S bc6a1dce1 |
| C2 | **Blessing of Kings and Heart of the Lion: still no Forever evidence.** | **Client:** Kings (20217, 25898) carries `SpellLabel` 26, the label every Blessing shares (Might 19838 and 25782 too). Heart of the Lion (409580, 409583) has no label. `SpellAuraRestrictions` has no row for either. The client cannot say.<br>**A stray row:** 1213408 "Blessing of Kings" is a level-1, self-targeted aura 137 at +10 that nothing in the client triggers (`SpellEffect` 1195051). It may be a server-side helper; it settles nothing.<br>**Wowhead Forever 409583:** "increasing total stats by 10% for all nearby allies, and increasing total stats for the Hunter by an additional 10%", the Season of Discovery wording.<br>**Web:** no Forever page, Blizzard note or Reddit thread found (Exa and web search, 2026-11-01). Season of Discovery is still the only evidence: they do not stack (Blizzard forums, 2023-12-11).<br>**Engine today:** the raid preset does not carry Kings. A request with both multiplies them (1.10 x 1.10, `sim/core/buffs.go:302` and `:264`).<br>**Settle:** a Forever character sheet with both buffs up. | OPEN | - |
| C3 | **The in-browser sim's "raid-buffed" list is not the raid preset.** `web/src/lib/sim/settings.ts:139` carries Blessing of Kings, Sanctity Aura (not learnable in Forever), Winter's Chill and Shadow Weaving, and no Heart of the Lion. `data/curated/presets.json` refuses all four of those and carries Heart of the Lion. | Code read. Not touched: the brief keeps this lane off the /sim landing page. | OPEN (site, outside this lane) | - |
| C4 | **Blood Fury's cooldown stays 2 minutes.** A third-party summary of the 2026-10-15 notes says 3 minutes. | **Client:** `SpellCooldowns` 55612, RecoveryTime 120000 for 20572.<br>**Wowhead Forever 20572:** "2 min cooldown".<br>The engine's 2 minutes (`sim/core/racials.go:238`) matches both. It is still off the GCD (round 2). | Aligned | - |
| C5 | **The pre-pull landing hold reaches the hybrids correctly.** Elemental's pre-pull Lightning Bolt lands and counts. Enhancement's two totems are buffs. Feral's Prowl sits at -2 s, and Ravage, first in the list, fires once (1.0 cast, 1013 average). | spec-breakdown runs. | Aligned | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Omen of Clarity's spell chance was 2 procs a minute over the cast time** (a 3 s Starfire 10%, an instant 5%). Forever gives a flat chance of about 4% a spell. | **Logs:** 34 caster druids (no form above 10% uptime) gained Clearcasting 547 times from 12,402 casts and 1,198 melee hits. The old model predicts about 840 from those casts alone. A per-character least-squares fit gives instants 8% and cast-time spells 1%, so the chance does not grow with cast time. Taking out melee procs at 2 a minute leaves about 4% a spell.<br>**Reddit (2026-09-30):** "4% chance on spells and 10 sec internal cd" (https://www.reddit.com/r/classicwow/comments/1wu7hcu/).<br>**Client:** 16864 has no `SpellProcsPerMinuteID`, which fits a flat server chance.<br>**Test first:** `sim/druid/omen_test.go` failed, then passed. Moonkin Form still doubles the chance and halves the 10 s cooldown (24858 effects 4 and 5).<br>**Effect:** 5.8 procs a fight, down from 9.0. Balance DPS moves -0.1% at seed 1 and 0.0% at seed 50001: in the raid run it is not short of mana. | CONFIRMED-BUG (logs plus one claim; low levels only) | E 36b92c966 |
| B2 | Raid talent search with `-keep`: the guide build is best. | Round-2 B1. Nothing in this round changes Balance's talent values beyond B1's mana. | Aligned | - |
| B3 | Moonfire and Insect Swarm ticks crit (29.7% and 25.6%). Starfire and Wrath follow the Eclipse lines. | spec-breakdown. | Aligned | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Shred dealt Era's 225% weapon damage. Forever's is 155%.** Rank 1 and 2's flat bonuses were 24 and 24, where the client has 22 and 32. | **Client:** `SpellEffect` effect 31 (weapon percent) is 155 on every rank (5221, 6800, 8992, 9829, 9830). Effect 58 (flat) is 22, 32, 44, 64, 80.<br>**Wowhead Forever 9830:** "causing 155% damage plus 80". Wowhead Classic reads "225% damage plus 180".<br>**Logs** (Shred against the same character's paw hit, both physical, so armor cancels; crits taken out at x2):<br>- report 2717, character 107989: rank 1 Shred 82 against a 31 paw hit. Forever gives 1.55 x (31 + 22) = 82; Era would give 119.<br>- report 2710, character 96227: rank 2 Shred 120 against a 43.3 paw hit. Forever gives 117; Era would give 169.<br>- report 2771, character 140512: rank 2 Shred 114 against 36.2. Forever gives 106 before talents; Era would give 153.<br>The flat bonus is multiplied: 1.55 x 31 + 22 = 70 misses the 82.<br>**Test first:** `sim/druid/weapon_percent_test.go` checks Shred, Claw and Ravage against the client on every rank. It failed, then passed. | CONFIRMED-BUG | E 07d744876 |
| F2 | **Savage Fury missed Shred and wrongly raised Pounce.** | **Client:** 16998's mask 0x9800 holds Shred (0x8000), Claw and Rake (0x1000), and Maul and Swipe (0x800) (`SpellEffect` 693144/693145, `SpellClassOptions`). Pounce is 0x20000.<br>**Blizzard beta notes (2026-10-15):** "Shred's tooltip counts Savage Fury (the damage itself was right)".<br>**Test first:** `sim/druid/feral/savage_fury_test.go` failed, then passed.<br>Pounce is not in the APL, so the Pounce half moves nothing. | CONFIRMED-BUG | E 07d744876 |
| F3 | **Claw dealt 100% weapon damage. The client has 110%** on every rank (1082 to 9850). | **Wowhead Forever 1082:** "110% normal damage plus 27".<br>**Log:** report 2717's rank 1 Claw is 66.5 against a 31 paw hit; 1.1 x (31 + 27) = 64.<br>The APL casts Claw only before Shred is learned, so the raid number does not move. | CONFIRMED-BUG | E 07d744876 |
| F4 | **F1 to F3 together:** Feral moves -11.3% on the published band and -11.4% with a real two-hander. | Shred falls from 308 to 234 DPS; Savage Fury gives back 10% of what is left. | (measured) | - |
| F5 | **Rake upkeep, re-measured on the 155% Shred:** still -4.1% (565.1 and 566.0 against 589.0 and 588.9). It was -9% on the old Shred. The list is unchanged; the guide and the APL notes carry the new figure. | `-rotation` with a Rake upkeep line before Shred. | Aligned | S c685eb1bd |
| F6 | **The published band still wears Elemental Mage Staff (944).** The re-rank on weapon DPS has not landed: `bis/druid-feral.json` is the nightly at 17d31ecfa (committed 2026-11-01 02:48 UTC). A bis.yml run was in progress at 04:00 UTC. | `git log origin/main -- data/builds/1.60.1.70707/bis/druid-feral.json`; `gh run list --workflow bis.yml`. | Follows (nightly) | - |
| F7 | **Raid talent search, re-run with a real weapon** (Grand Marshal's Demolisher 234568 in a scratch copy of the band) on the lane engine: the guide build is best (682.1 ± 0.9). Savage Fury now credits +12.6 a point.<br>Ferocity's 5 points buy nothing, because the APL casts neither Claw nor Rake. Every alternative in tiers 0 and 1 is also worth nothing, and Blood Frenzy's second point is +0.2%. | `sim/cmd/talent-search -preset raid -keep "Shifting Power,Leader of the Pack,Nature's Focus,Subtlety"`, before and after F1. | Aligned | - |
| F8 | Omen of Clarity's melee rate stays at 2 procs a minute. | **Logs:** 10 cats (Cat Form above 70% uptime) gained Clearcasting 346 times from 6,131 paw hits, 3,106 specials and 455 spells. 2 a minute on the 1.0 s paw, plus spells at 5%, predicts 331 before the 10 s cooldown. | Aligned (logs) | - |
| F9 | The form attack uses the weapon's DPS (round-2 F1) and sits at 1.0 s; the Ravage opener counts. Rip ticks crit on the melee table (37.2%). | `sim/druid/forms.go:50`, spec-breakdown. | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | Savage Fury raises Maul and Swipe, in the client mask, and the engine applies it to both (`maul.go:54`, `swipe.go:91`). Threat is unchanged by this round (694.0 both seeds). | Client mask 0x800. | Aligned | - |
| G2 | The bear paw at the weapon's DPS and 2.5 s, and Lacerate's melee-table crits, are unchanged from round 2. | Round-2 G1 and G3. | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Omen of Clarity on heals now runs at B1's 4% a cast, not 2 a minute over the cast time. Effective HPS moves -1.5%: fewer free Healing Touches, and mana lasts 322 s instead of 358 s in a 300 s fight. | B1. | CONFIRMED-BUG (B1) | E 36b92c966 |
| R2 | Wild Growth skipped, Bloom taken: unchanged from round 2. | Round-2 R1 and R2. | Aligned | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | **Improved Reincarnation stays in the build** (owner ruling: the adoption rule keeps every unmodeled talent). | Round-2 E3. | Settled (ruling) | - |
| E2 | Fire Nova's Forever damage and its Clearcasting line hold. The damage is rank 5's 408428 roll at the 0.143 coefficient with Call of Flame and Trial by Fire. The cast fires 5.4 times a fight, and its average with crits is 550. | `sim/shaman/fire_nova.go:25`; spec-breakdown. | Aligned | - |
| E3 | The orc's Blood Fury is off the GCD, on its client 2-minute cooldown (two uses a fight), and adds 10% spell power (Wowhead Forever 20572). | C4. | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | Stormstrike is 100% weapon damage in the client (17364 effect 121, no percent effect); the engine's `DamageMultiplier: 1` matches. Windfury Weapon, Maelstrom Lightning Bolt and Earth Shock are unchanged from round 2. | `SpellEffect`; `sim/shaman/stormstrike.go:49`. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Heal profile: 611 effective HPS. Riptide, Chain Heal and Healing Wave as round 1 set them. Mana lasts 413 s. | Heal-profile run. | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Judgement of the Crusader in the preset: +9.6% on the published band. | C1. | RULED | S bc6a1dce1 |
| P2 | **The Unsealed Wrath twist still pays, more with the debuff:** 645.7/646.3 against 623.3/623.1 without the twist (+3.6%) on raidBuild `050050002/503/05233301001301351`, with no time out of mana. Every stored Seal of Righteousness copy also takes Judgement of the Crusader. | `-rotation` with and without the two `1310735` lines. | Aligned | - |
| P3 | **Holy Strike's flat bonus: is it multiplied by the weapon percent?** The engine applies the percent to (flat + normalized weapon), as Shred does (F1). Wowhead Forever 10333 reads "50% weapon damage plus an additional 97 to 121 as Holy damage". | **Log:** report 2675, character 38190, rank 1866 (32%, flat 19): Holy Strike averages 57.6 without crits. That falls between the two readings, depending on the mob's armor against the white hits used to infer the weapon roll.<br>**Stakes:** at 60 the unmultiplied reading is about +15% Holy Strike, about +1.6% Retribution. | OPEN: a training-dummy log of Holy Strike against a known weapon and attack power settles it | - |
| P4 | Raid talent search: the round-2 raidBuild with Unsealed Wrath is the guide's. | Round-2 P2 and P5. | Aligned | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Judgement of the Crusader in the preset: threat 659.7 to 721.2 (+9.3%), DPS +7.8%. Holy Strike's Sacred Arbiter refresh only refreshes a judgement that is already up, so the permanent raid debuff is unaffected. | C1; `sim/paladin/holy_strike.go:173`. | RULED | S bc6a1dce1 |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | Heal profile: 531 effective HPS, unchanged by the debuff. Flash of Light, Holy Light and Holy Shock (through 1311605) as in round 2. | Heal-profile run. | Aligned | - |

## Before/after (relative %)

**Method:**
- Raid preset, level-60 horde band, 10000 iterations, seeds 1 and 50001. The two seeds agreed within ±0.1%.
- **Before:** pin c8d29ea33 with origin/main.
- **After:** lane engine 36b92c966 with `wt/r3-hybrids`.
- Healers are scored on effective HPS on the heal profile, tanks on threat per second.

| Spec | Change | Before → after |
|---|---|---|
| Balance | B1 | -0.1% / 0.0% |
| Feral (published band, caster staff) | F1 to F3 | -11.3% (663.7 → 589.0) |
| Feral (High Warlord's Destroyer 234546 in the main hand) | F1 to F3 | -11.4% (776.2 → 687.4) |
| Feral bear (threat) | none | 0.0% |
| Restoration Druid | R1 | -1.5% |
| Elemental, Enhancement, Restoration Shaman | none | 0.0% |
| Retribution | C1 | +9.6% |
| Protection Paladin | C1 | +9.3% threat (+7.8% DPS) |
| Holy Paladin | none | 0.0% |
| Every spec outside the group | C1 sweep | 0.0% (all 28 specs run) |

Nothing nightly-owned was committed. TestRotationLadder fails on this branch as expected; the integrator regenerates it at the pin.

## Still open

1. **C2, Kings and Heart of the Lion stacking.** A Forever character sheet with both buffs settles it.
2. **C3, the in-browser "raid-buffed" list** (`web/src/lib/sim/settings.ts:139`) disagrees with the curated raid preset (Kings, Sanctity Aura, Winter's Chill, Shadow Weaving; no Heart of the Lion). It is the sim page's list, which this lane may not touch.
3. **P3, Holy Strike's flat bonus under the weapon percent.** A training-dummy log settles it.
4. **B1's 4% rests on levels 20 to 30.** A level-60 Balance or Restoration log with Clearcasting (16870) gains re-checks it after launch.
5. **Unsealed Wrath's unconfirmed reading** and **Wild Growth under a party-damage profile**: carried over from round 2.

## Sources

- Wowhead Forever tooltips (fetched 2026-11-01): 9830, 1082, 9867, 16998, 10333, 17364, 20217, 25898, 409580, 409583, 1213408, 20572; Wowhead Classic 9830.
- foreverlogs.gg reports 2667 to 2809 (beta, 2026-09 to 2026-10): `buff_uptimes`, `casts` and `character_abilities?scope=boss` per druid, and 2675's `combatants-info` for encounter 4898.
- Blizzard beta notes, 2026-10-15 (third-party summary: https://wowclassicforever.info/patch-notes/).
- Reddit, Omen of Clarity rate (2026-09-30): https://www.reddit.com/r/classicwow/comments/1wu7hcu/
- Season of Discovery, Kings and Heart of the Lion (2023-12-11): https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967
- Client 1.60.1.70707 wago tables: `SpellEffect` (688178–688180, 693144–693145, 1195051), `SpellClassOptions` (44164, 44196, 44236, 44322, 44327, 44328), `SpellLabel` (190072, 190155), `SpellAuraOptions` (129438), `SpellCooldowns` (55612).
