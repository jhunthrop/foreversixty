# Sim audit round 12: healers

Lane: wt/r12-healers (site, from c6bdc855c) and lane/r12-healers (fork, from 2606bc05e). Specs: Holy priest, Discipline priest, Holy paladin, Restoration druid, Restoration shaman, at raid 60 and the bare bands.

Client truth: 1.60.1.70753, the raw CSVs fetched for round 11 (scratchpad r11/c70753). They already include 70707's push. On 2026-11-03 wago lists no hotfix push for 70753 (`push_totals` returned 0 for Spell, SkillLineAbility, Talent, TraitDefinitionEffectPoints, ItemSparse and ItemEffect). Between 70707 plus its push and 70753, the only class-spell changes are feral ones (Poised, a Ferocious Bite variance digit). No healer row and no healer talent curve moved.

Method:
- **Full client-row sweep.** Every heal, heal over time, shield, mana tool and healing talent the five lists cast or the builds take was compared field by field against the client rows. The tables are in [healers/sweep-priest-paladin.md](healers/sweep-priest-paladin.md) and [healers/sweep-druid-shaman.md](healers/sweep-druid-shaman.md); each row is a spell or talent with every field checked. Two read-only helpers did the first pass. I re-read every mismatch they reported against the client row and the code before confirming it. The row printer and the build decoder are [healers/rows.py](healers/rows.py) and [healers/decode.py](healers/decode.py).
- **Fresh re-audit.** I re-checked each list's casts, mana and crit tables at raid 60 in spec-breakdown, the landing's Binding Heal reorder, the guide builds against the published builds, and the published races.

The short version:
- **Vampiric Embrace healed the group for 5% of Shadow damage. The client says 20%.** Its aura also missed Shadow ticks and counted other players' Shadow damage. Both are fixed. No published rotation casts it, so no number moves, Shadow's included.
- **Idol of Swiftness was never registered.** It takes 3 s off Swiftmend, and it is an Enchanting relic on Phase 1 loot. Every Restoration druid set left its idol slot empty. Equipped, it adds +1.2 to +1.3% at raid 60. The pick is for the refresh.
- **Gift of Nature leaves Tranquility out, and that follows the client.** Its periodic effect does not name Tranquility, and the beta logs agree. Aligned.
- **Two small, latent fixes:** Reflection's ranks 1 and 2 (17% and 33%, not thirds of 50%), and Mental Agility reaching Power Infusion.
- **Prayer of Mending travel stays open.** The cast has a missile speed of 40 and the engine has no travel. It moves Holy priest by 3 to 6% depending on the distance between allies, which no client row states.

## Assigned items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Vampiric Embrace healed the group for 5% of the priest's Shadow damage. Forever's is 20%.**<br>- 15286's one effect (SpellEffect 694156, aura 4 dummy) states 20.<br>- The server text: "causes all party members to be healed for 20% of any Shadow spell damage you deal for 30 sec."<br>- The 5% was Season of Discovery's, kept from upstream with Improved Vampiric Embrace commented out.<br>The Shadow guide already says "20% of every Shadow spell you land", so the page was right and the engine was wrong. | client row; https://nether.wowhead.com/forever/tooltip/spell/15286 (read 2026-11-03). Test `TestVampiricEmbraceHealsThePartyTheClientsShare` failed first (0.05 against 0.2). | CONFIRMED-BUG (latent: no rotation casts it) | old (upstream SoD, kept at 36c1def4e) | E 8a128d69d |
| A1b | **Vampiric Embrace healed only off direct hits, and off anyone's Shadow damage.**<br>- 15286's proc row (SpellAuraOptions 128620, ProcTypeMask 0x800A0000) fires on harmful magic spells taken and on periodic damage taken.<br>- The text says "any Shadow spell damage you deal".<br>- The engine's aura had only OnSpellHitTaken, so Shadow Word: Pain, Devouring Plague and Mind Flay ticks healed nobody, and a second player's Shadow hit on the target healed this priest's group.<br>It now heals off the priest's own Shadow hits and ticks only. | Test `TestVampiricEmbraceHealsOffThePriestsOwnShadowHitsAndTicks` failed first (another player's damage healed 200; no periodic callback). | CONFIRMED-BUG (latent) | old (upstream) | E 1c113274b |
| A1c | **Shadow does not move.** priest-shadow's list does not cast Vampiric Embrace (spell 15286 is absent from data/curated/apl/priest-shadow.json). Raid 60 is 712.7 / 712.9 / 720.0 / 720.0 before and after (alliance s1 / s50001, horde s1 / s50001). The guide tells a Shadow priest to keep it up on the boss; whether the list should cast it is for the casters lane (below). | spec-breakdown, 10,000 iterations | Aligned (no move) | - | - |
| A2 | **Gift of Nature does not reach Tranquility, and that follows the client.** 17104 has two effects:<br>- Effect 0 (SpellEffect 695243) is aura 108, misc 0, the direct-heal op. Its mask is word 0 0xE0 (Healing Touch 0x20, Regrowth 0x40, Tranquility 0x80) and word 1 0x2 (Swiftmend).<br>- Effect 1 (1291839, Forever's) is aura 108, misc 22, the periodic op. Its mask is 0x50 (Rejuvenation 0x10, Regrowth 0x40) and word 1 0x4000000 (Wild Growth).<br>Tranquility (740, 8918, 9862, 9863; SpellClassOptions word 0 0x80) has only a periodic heal: effect 35, aura 8, every 2 s. The direct op names it, but it has no direct part to raise. The periodic op that would raise its ticks leaves it out.<br>Round 10's logs agree: two level-35 druids' Tranquility ticks fit their Healing Touch healing power only without Gift of Nature (reports 2748 and 2770).<br>The engine's set (every heal but Tranquility) equals the union of the two effects. The tooltip's "all your healing spells" is looser than the masks. | client rows; round 10 healers A11 | Aligned (client masks + logs) | - | - |
| A3 | **Healer haste goes negative for the priests (for the casters lane).** The published haste scale factors are Discipline -8.0 at raid 60 and -12.3 at bare 60, and Holy -5.4 and -29.3. Druid band 20 is -3.1. The priests run out of mana mid-fight (Discipline's bar lasts 118 s of 300), and the lists are paced on mana against the fight clock, not on cast speed. So faster casts spend the bar sooner, and effective HPS over the fight falls. That looks like the mechanism. I did not change anything. | bis/<spec>.json haste_scale_factor; spec-breakdown -heal | for the casters lane | - | - |
| A4 | **Blood Fury on healing (for the auras lane).** No published healer is an orc: human and undead priests and paladins, night elf and tauren druids, dwarf and tauren shamans. The stale Holy priest entry is troll, and it casts Berserking (26297) twice a fight. So Blood Fury reaches no published healer number. | bis races; spec-breakdown | for the auras lane (no healer affected) | - | - |

## Holy and Discipline priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Every priest heal rank matches 70753:** level, cost, cast time, coefficient, base, variance, per-level growth and cap.<br>- Lesser Heal 1-3, Heal 1-4, Flash Heal 1-7, Greater Heal 1-5, Prayer of Healing 1-5, Renew 1-10 (15 s, 5 ticks, no tick on application, ticks can crit).<br>- Power Word: Shield 1-10 (4 s category), Penance 1-4 (3 pulses), Prayer of Mending 1-3 (5 charges, 10 s), Binding Heal 1-6.<br>- Lightwell 1-3 and its Renew, Dark Sacrifice 1-5, Divine Grace 1-7, Weakened Soul 15 s, Inner Focus. | sweep-priest-paladin.md, priest ranks table | Aligned | - | - |
| P2 | **Every taken priest talent's curve and mask matches:** Twin Disciplines, Mental Strength, Mental Agility, Inner Focus, Improved Power Word: Shield, Meditation, Soul Warding, Renewed Hope, Divine Aegis, Improved Renew, Holy Specialization, Divine Fury, Inspiration, Improved Healing, Holy Binding, Litany of Light, Spiritual Guidance, Spiritual Healing. | sweep-priest-paladin.md, priest talents table | Aligned | - | - |
| P3 | **Heals learned below 20 keep their whole coefficient on every path** (round 11's rule). Every priest heal registers with the healing proc mask. Power Word: Shield reads its coefficient directly. Lightwell Renew, Dark Sacrifice and the Holy Binding heal are not cast ranks. | `core/spell.go:803-811`; sweep | Aligned | - | - |
| P4 | **Mental Agility did not reach Power Infusion.**<br>- Its cost modifier (14520, SpellEffect 687728, aura 108 op 14, -10%) names word 0 0x8251c1eb.<br>- That mask carries Power Infusion's 0x80000000 (SpellClassOptions of 10060).<br>- Power Infusion had no engine class mask, so the modifier could not name it. It now has one. | client masks. `TestMentalAgilityFollowsTheClientMask` extended first (failed). | CONFIRMED-BUG (latent: Power Infusion is out of every default) | old | E 3ae9bba12 |
| P5 | **The landing's grouped Binding Heal lines (f726a437b) cast exactly as the interleaved order did.** On the current list, re-interleaving the four party Binding Heal and Flash Heal pairs gives identical effective HPS: 653.6 horde and 649.8 alliance at raid 60, 216.1 at bare 50 and 317.8 at bare 60 (3,000 iterations, seed 1). Each pair's conditions differ only by "you are under 90%", so no state can pick a different heal. The tank's Binding Heal line stays ahead of its Flash Heal. | spec-breakdown -rotation, both orders | Aligned | - | - |
| P6 | **Holy's tank Flash Heal at rank 2 is right at raid 60.** The list never runs dry with it (the bar lasts 480 s), so I tested whether a bigger rank heals more. Rank 2 696.0, rank 4 647.1, rank 5 637.8, rank 6 666.3, rank 7 683.3 (horde, seed 1, 3,000 iterations). The bigger ranks run the bar out at 140 to 220 s. | spec-breakdown -rotation | Aligned | - | - |
| P7 | **Prayer of Mending has no travel time in the engine.**<br>- The cast 1240827 (and 401859, 1240826) has SpellMisc Speed 40, and the engine places the charges at once.<br>- The holder auras 1240848 and 1240849 have Speed 0, and no Forever jump spell with a speed can be identified.<br>Delaying every placement and hop (scratch hack, not committed) moves Holy raid 60 from 696.7 to 673.4 at 0.25 s a hop (-3.3%) and to 655.6 at 0.75 s (-5.9%). The cause is that fewer hops land before the 10 s recast replaces the charges. The size depends on the distance between allies, which no client row gives, and Prayer of Mending is a tier-6 talent, so no beta log has it. | SpellMisc rows; spec-breakdown with the hack | OPEN (a level-60 log with Prayer of Mending hops settles the delay) | - | - |
| P8 | **Dark Sacrifice's ranks each have their own cooldown timer.** The client shares category 1144 (600 s) across ranks. A fight casts one rank, so no number can move. | SpellCategories 1277324-1277328 | Aligned in effect (unreachable) | - | - |
| P9 | **Prayer of Mending's holder auras carry Variance 0.109.** The engine heals the flat centre. The mean is the same, and the variance sits on a dummy aura that a script may not read. | SpellEffect 1047993, 1239448 | OPEN (low; jump amounts from a log) | - | - |
| P10 | Round 10 and 11's open priest rows have no new evidence: Lightwell Renew's crit, Renewed Hope per Penance pulse, Sudden Insight per target, Litany of Light across ranks, Holy Specialization and Spiritual Healing on the Prayer of Mending heal, Inner Focus on Penance pulses, and Divine Aegis's cap at 40 and 50. | sweep | OPEN (carried) | - | - |
| P11 | **The guide builds out-heal the published builds:** Holy 696.0 to 711.0 (+2.2%), Discipline 653.6 to 665.3 (+1.8%), at raid 60 (horde, seed 1, 3,000 iterations). The published bis entries are from 2026-11-01. They still carry Wand Specialization (Discipline) and the gnome and troll races (Holy). This is the nightly's refresh, not a sim bug. | spec-breakdown -talents with the guide strings | Measure at refresh | - | - |

## Holy paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Every paladin heal and mana tool matches 70753:**<br>- Holy Light 1-9 and Flash of Light 1-6.<br>- Holy Shock, cast and heal. 1311606 is Forever's only learnable rank: ranks 2 to 4 are AcquireMethod 3 runes. It has a 10 s category cooldown and costs 160.<br>- Divine Favor (mask 0x80206000, 120 s), Lay on Hands 1-3, Oath of Light, Light's Vigil, and the Illumination and Infusion of Light buffs. | sweep-priest-paladin.md, paladin tables; SkillLineAbility 13935-13940, 59253 | Aligned | - | - |
| H2 | **Every taken paladin talent matches:** Spiritual Focus, Healing Light, Reverence, Illumination, Divine Favor, Infusion of Light, Boundless Light, Holy Power, Divine Intellect, Benediction, Conviction. | sweep | Aligned | - | - |
| H3 | **Illumination's fourth effect** (20210 e3, aura 4, base 10) is not modelled. The server text names only the chance and the 50% refund. | SpellEffect rows; tooltip read 2026-11-03 | OPEN (low; an Illumination energize log) | - | - |
| H4 | Holy Power's crit size, Oath of Light's scaling and Illumination on Holy Shock rank 1 are still open (no level-40+ paladin log). | - | OPEN (carried) | - | - |

## Restoration druid

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Idol of Swiftness (279250) was never registered.**<br>- **What it is:** an Enchanting relic, item level 65, requires level 60, epic.<br>- **The client:** its equip spell 1306488 "Swiftmend Cooldown Reduction" is aura 107 op 11, -3000 ms, on word 1 0x2, Swiftmend's mask. The server text: "Reduces the cooldown of your Swiftmend spell by 3 sec."<br>- **The bug:** the druid relic table mapped only Rip and Shifting Power, so the idol could not register. The gear search scored it at zero, and every Restoration set left its idol slot empty. The shaman's Totem of Urgency had the same gap in round 10.<br>- **The fix:** the table names Swiftmend and the idol is registered. The fork's item database gains its row the way the other Forever relics did (overrides.go, and one item inserted into db.bin and db.json; the rest of db.bin round-trips byte for byte).<br>- **Measured equipped at raid 60** (10,000 iterations): horde 833.7 to 844.2 / 833.5 to 843.9, alliance 834.6 to 844.8 / 833.8 to 845.0 (+1.2 to +1.3%). Swiftmend goes from 17.6 to 21.9 casts a fight. | ItemXItemEffect -> 1306488; https://nether.wowhead.com/forever/tooltip/item/279250 (read 2026-11-03); loot.json crafted:enchanting. `TestIdolOfSwiftnessCutsSwiftmendsCooldownOnly` failed first (the idol could not register). | CONFIRMED-BUG | old (c88324d55 modelled the DPS idols only) | E 3c6942947 |
| D2 | **Reflection's ranks 1 and 2 gave 16.67% and 33.33% of regeneration while casting.** The client's trait curve (TraitDefinitionEffectPoints 134381, curve 85365) states 17, 33 and 50. Rank 3 was already exact. It reaches only the stale band-20 entry with Reflection 1, and there it measures 39.4 to 39.4 (bare 20, horde, both seeds). | client curve; `TestReflectionKeepsAShareOfRegenerationWhileCasting` changed to the client values first (failed) | CONFIRMED-BUG (negligible) | old | E 12ebf562f |
| D3 | **Every druid heal rank matches:** Healing Touch 1-11, Regrowth 1-9 (21 s, 7 ticks), Rejuvenation 1-11 (12 s, 4 ticks), Tranquility 1-4 (5 ticks over 10 s, 5 min), Wild Growth 1-3, Swiftmend (20% base mana, 15 s), Bloom (100% and 50% of Spirit), Innervate (5%, 6 min, 20 s), Nature's Swiftness and Omen of Clarity. | sweep-druid-shaman.md | Aligned | - | - |
| D4 | **Every druid talent's curve and mask matches:** Gift of Nature, Improved Rejuvenation, Tranquil Spirit, Improved Tranquility, Naturalist, Nature's Focus, Living Spirit, Heart of the Wild, Genesis (leaves Tranquility out), Nature's Majesty, and Nature's Splendor (+3 s Rejuvenation, +6 s Regrowth). | sweep | Aligned | - | - |
| D5 | **Swiftmend with Nature's Splendor pays out 5 Rejuvenation or 9 Regrowth ticks.** The client row and the text ("the full duration") do not say whether the extended duration counts. If only the base duration counts, Swiftmend off Regrowth is overpaid by 9/7. No logged druid has Splendor. | `druid/swiftmend.go:68` | OPEN (one level-60 log with Splendor) | - | - |
| D6 | **The guide build heals more than the stale published one:** 834.5 to 840.6 (+0.7%, raid 60, horde, seed 1). The published 60 entry still takes Improved Wrath 4, which heals nothing. | spec-breakdown -talents | Measure at refresh | - | - |
| D7 | Bloom's crit (round 11 D3) and Rejuvenation's same-op stacking (round 10 D3) are still open. | - | OPEN (carried) | - | - |

## Restoration shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Every shaman heal and totem matches 70753:**<br>- Healing Wave 1-10 (rank 2 for the 40-59 rank rows), Lesser Healing Wave 1-6 and Chain Heal 1-3 (3 targets, ×0.5 a jump).<br>- Riptide 1-3: the direct heal, 5 ticks over 15 s, the 6 s cooldown and the 25% Chain Heal bonus.<br>- Healing Stream 1-5 (every 2 s, the party), Mana Tide 1-4 (4 ticks, 5 min) and Mana Spring 1-4.<br>- Water Shield (3 charges, 3.5 s internal cooldown, 2% mana, no periodic proc) and Nature's Swiftness. | sweep-druid-shaman.md | Aligned | - | - |
| S2 | **Every taken shaman talent matches:** Improved Healing Wave, Tidal Focus, Mindfulness, Healing Focus, Tidal Mastery, Restorative Totems, Healing Way, Purification, Ancestral Knowledge, Thundering Strikes; Ancestral Healing (armor, unmodelled). | sweep | Aligned | - | - |
| S3 | Two wrong node ids in code comments (Tidal Focus is 104729, Healing Focus 104733). No number depends on them. | `shaman/talents_restoration.go:16`, `talents.go:588` | Aligned (comment only, not changed) | - | - |
| S4 | Chain Heal's jump order, and how Purification and Healing Way stack, are still open. | - | OPEN (carried) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **No hotfix push for 70753** on wago (2026-11-03). Every healer row above is the shipped 70753 row. | `pipeline.hotfix_push.push_totals` | Aligned | - | - |
| X2 | **The Devout set's mana proc** (450576, 5%, ProcTypeMask 0x14000, 200 mana) fires 5 to 9 times a fight. That is 5% of each healer's casts. | `core/client_sets_gen.go:465-468`; spec-breakdown | Aligned | - | - |

## Before/after

Before: $R/before.tsv (raid 60) and $R/sb-before from $R/site-before/sim. After: lane/r12-healers 3ae9bba12 with wt/r12-healers. 10,000 iterations, effective HPS (DPS for Shadow).

| spec | case | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|---|
| Discipline priest | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 653.5 |
| Holy priest | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 696.7 |
| Holy paladin | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 645.1 |
| Restoration druid | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 833.7 |
| Restoration druid | bare 20 | - | - | 0.00% | 0.00% | 39.4 |
| Restoration shaman | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 840.2 |
| Shadow priest | raid 60 | 0.00% | 0.00% | 0.00% | 0.00% | 720.0 |
| Restoration druid, Idol of Swiftness equipped | raid 60 | +1.22% | +1.34% | +1.26% | +1.25% | 833.7 to 844.2 |

No published number moves. Vampiric Embrace and Power Infusion are not cast. The idol is on no published set until the refresh. The only build that holds Reflection 1 is the stale band-20 entry, and there the change is under a tenth of a point.

## Files to regenerate

None of mine. The fork's item database gained one row in place (Idol of Swiftness: db.bin, db.json and overrides.go). That is not a regen.

## Measure at refresh

- Restoration druid at 60, raid and bare: the idol slot should fill with Idol of Swiftness (D1).
- Holy and Discipline priest, and Restoration druid: the guide builds measure ahead of the stale published ones (P11, D6). Holy's races are already on the list (MR6).

## For other lanes

- **Casters (Shadow).** Vampiric Embrace now heals the group for 20% of the priest's Shadow damage, off ticks as well as hits (E 8a128d69d, E 1c113274b). Shadow's DPS does not move, because its list does not cast Vampiric Embrace. The guide says "On the boss, keep it up". Decide whether the list casts it: it costs a global every 30 s, and it adds group healing the DPS metric does not count.
- **Casters (haste).** See A3: the priests' negative haste factors look like mana pacing (faster casts empty a mana-bound bar sooner). This lane changed nothing.
- **Auras.** No published healer is an orc, so Blood Fury reaches no healer number (A4).
- **Warlock lane and coordinator.** The Prayer of Mending travel question is answered in P7: OPEN, with its size measured.
- **Integrator.** The fork's item database gains Idol of Swiftness (279250), in the same three files round 10 used for Totem of Urgency. Fork tests touched: priest, priest/healing, priest/shadow, druid/..., conformance, common/item_sets, all passing with `-tags=with_db`. No golden moved. No site Go code changed, so the ladder goldens do not move from this lane.

## Still open

- Prayer of Mending travel and hop delay (P7) and its holder variance (P9).
- Swiftmend's payout with Nature's Splendor (D5).
- Illumination's fourth effect (H3). Holy Power, Oath of Light, and Illumination on Holy Shock rank 1 (H4).
- Inner Focus not spent by a Power Infusion cast (14751 e0 names 0x80000000; the engine casts Power Infusion with no on-cast hook). Latent, because Power Infusion is out of every default. It is left for a round that puts Power Infusion in a preset.
- The carried rows: P10, D7, S4.
