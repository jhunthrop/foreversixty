# Sim accuracy 2026-11-11: the three priest specs

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-priests`. Site from origin/main cfe2a6f67, fork from master e586233d7 (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70888. Scope: Shadow, Holy and Discipline, end to end: engine against the client (both builds, hotfixed), engine against real fights (foreverlogs.gg and our API), builds, rotations and rank rows under the adoption rule. Last priest audits: rounds 24 and 25 (2026-11-06), round 29's leveling walks, round 30's rank rows (`round30-open.md`).

**Totals: 1 confirmed engine bug, old, fixed (Dark Sacrifice's tick ignored the client's per-level growth; Undead only, between rank levels only, under 1% anywhere). 0 new data: the 70888 client and its push change no priest row. 19 checks against the first priest logs aligned, 3 held for lack of a level-40+ log. No build, rotation or rank row moves. Engine pin e586233d7 to fe282fdfb.**

## Method

1. **Blizzard since 6 Nov.** "WoW: Forever Hotfixes" (t/2403040, posts 4 to 6: 6, 9 and 10 Nov), "Class Tuning Incoming" (t/2404092) and the development notes thread (t/2360696, post 9). Priest lines: the Revelation enchant no longer procs off Vampiric Embrace healing (10 Nov), Holy Nova unlearns with its talent (10 Nov), Shadowform survives Fear Ward and Power Infusion persists on allies (9 Nov). None is modelled or on a card; no number moved.
2. **Client and push.** wago.tools CSVs for 1.60.1.70790 and 1.60.1.70888, both with the hotfix push layered (`useHotfixes=1`): SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellInterrupts, SpellProcsPerMinute, Spell, SpellName, plus the SpellDuration and SpellCastTimes indices priest rows reference. Priest rows are every spell with SpellClassSet 6 plus the racials, Shadowfiend, the talent auras and the set bonuses (645 spells, 899 effect rows). 70790 against 70888: zero differences. The 70888 push over its shipped tables touches two Holy Fire effect rows for float precision only (1.29999995232 to 1.2999999523163). So the site's 70790 tables are the live priest.
3. **Engine against the 70790 tables**, every spell, talent, aura and racial the three cards and builds touch: `sim/priest/*.go` read line by line against `data/builds/1.60.1.70790/spellconst/priest.json` and `talents/priest.json` (the full table is below).
4. **Real fights.** Our API holds two public reports (z2opt2x2gs5u, 20 Oct; gxrgvz4yzsi2, 8 Nov open world) with no priest of use. foreverlogs.gg has 24 public reports (2815 to 2850, 5 to 11 Nov): their actor lists name 25 priests, 19 of them with 20 or more events; every event those 19 did was read through the report API (`meta_events?sourceCharacterIds=<id>`, 500 a page, paced). The useful ones: a Shadow priest at 28 to 33 (Itso, report 2839, Gnomeregan and Dalaran: 275 Mind Flay channels, 283 Shadow Word: Pain ticks, 48 Mind Blasts, 4 Vampiric Embraces), a Smite priest at 22 to 27 (Emanuele, 2848: 107 Holy Fires, 214 Smites, 49 Mind Blasts) and six healers at 14 to 27 (Net, Shadow, Luminare, Cogara and others: 977 Renew ticks, 201 Heals, 215 Lesser Heals, 48 Flash Heals). No priest in any public log is 40 or above, so Penance, Prayer of Mending, Lightwell, Shadowform, Shadow Word: Death and the tier 5+ talents still have no log. Heal events in these logs carry the effective amount (overheal is a separate event), so heals settle rates and timings, not formulas; damage settles both.
5. **Sims** on the Windows box through `tools/remote-test.sh` (`--dir sim`, `--engine-sha` for the lane's fork branch, `nice -n 15`, GOMAXPROCS=5, one sim at a time): the fork's priest suites (`-tags with_db`), and spec-breakdown before and after the fix at the levels it touches.

## Engine against the client (1.60.1.70790, hotfixed)

Every number below matched; the one row that did not is D1.

| area | engine | client | verdict |
|---|---|---|---|
| Mind Blast | 1.5 s cast, 8 s category cooldown less 0.5 s a rank of Improved Mind Blast, coefficient 0.429, rolls and growth from the generated ladder | 8092 to 10947: cast 1500, category 8000, sp 0.429; talent 15273 0.5 s a rank | aligned |
| Mind Flay | 3 ticks a second apart, 0.167 a tick, costs 45 to 205, Improved Mind Flay compounds 10% a rank per tick, Dark Inspiration on the third tick | 15407 to 18807: duration 3000, period 1000, sp 0.167; 1225139 "10%/20% more than its previous instance" | aligned |
| Shadow Word: Pain | 6 ticks of 3 s, +1 tick a rank of Improved Shadow Word: Pain, 0.2 a tick, rank 8 keeps 10894's 127 | 589 to 10894: 18000/3000, sp 0.2, amounts 5 to 127; 15275 +3 s a rank | aligned |
| Shadow Word: Death | instant, 18 s category cooldown, 0.429, 30% of max health backlash on a landed cast, none on a miss | 1309595 to 1309636: category 18000, sp 0.429, effect 2 amount 30; Blizzard 3 Nov "no longer damage the Priest when the spell misses" | aligned |
| Devouring Plague | 8 ticks of 3 s, 24 s category cooldown, 0.1 a tick, ticks crit, Devouring Contagion -25% cost a rank | 2944 to 19280: 24000/3000, category 24000, sp 0.1; 1309950 25%/50% | aligned |
| Vampiric Embrace | 20% of the priest's landed Shadow damage heals the party, 30 s, 30 s cooldown, 40 mana, cannot be resisted, never crits | 15286: aura 4 amount 20, duration 30000, cooldown 30000, cost 40, Attributes_3 0x40000; Blizzard 3 Nov "can no longer roll a critical effect" | aligned |
| Shadowform | +10% Shadow damage, -50% Shadow cost (school multiplier after the spell mods), +100% Shadow crit damage bonus, -15% physical taken, 40% base mana, 1.5 s GCD and cooldown, Holy casts drop it | 15473: aura 79 10 (misc 32), aura 72 -50 (misc 32), aura 108 100 (misc 15), aura 87 -15 (misc 1), PowerCostPct 40, category 1500 | aligned |
| Shadow Weaving | 2% a stack, 5 stacks, 15 s, 33% a rank to add a stack | 15257 proc 100% a rank third, 15258 aura 270 amount 2 for 15000, text "33/67/100%"; the log's debuff is named Shadow Vulnerability (15258) and stacks to 5 | aligned (modelled as a self buff: the text's "the Shadow damage the target takes from you" is the priest's own damage either way) |
| Darkness, Shadow Focus, Twin Disciplines | 2% Shadow damage a rank; 1% Shadow hit a rank; 1% instant damage and healing a rank | 15259, 15260, 1225132 texts | aligned |
| Inner Focus | next priest spell free, +25% crit on non-periodic, 3 min | 14751: aura 108 -100 misc 14, aura 107 25 misc 7, cooldown 180000 | aligned |
| Dark Inspiration | 3% Spirit a rank for 10 s off Mind Blast and a full Mind Flay | 1326168: aura 137 15 at rank 5, duration 10000 | aligned |
| Power Infusion | +20% spell damage and healing for 15 s, 3 min, 20% base mana | 10060: aura 136 20, aura 79 20 (misc 126), duration 15000, cooldown 180000, PowerCostPct 20 | aligned |
| Smite, Holy Fire | Smite cast 1.5/2.0/2.5 s, 0.429/0.571/0.714; Holy Fire 3.5 s, direct 0.75, 5 ticks of 2 s at 4/5/6/7/9/10/13/15 and 0.05 | 585 to 10934; 14914 to 15261 (effect 1 amounts 4 to 15, period 2000, duration 10000, sp 0.05) | aligned |
| Heals (Lesser Heal, Heal, Flash Heal, Greater Heal, Prayer of Healing) | every rank's cost, cast, level, amount, variance, growth and coefficient in `heal_ranks.go` | 2050 to 25316 | aligned (TestHealingLaddersMatchTheClient pins them) |
| Renew | 5 ticks of 3 s, 0.2 a tick, 9 to 166, ticks crit | 139 to 25315: 15000/3000, aura 8, sp 0.2; every rank carries PERIODIC_CAN_CRIT | aligned |
| Power Word: Shield, Weakened Soul | absorb + 0.1 healing, 30 s, 4 s category cooldown, Weakened Soul 15 s, Improved 7/14/20%, Soul Warding -4 s and -15% | 17 to 10901: aura 69, sp 0.1, 30000, category 4000; 6788 15000; 14748; 402000 | aligned |
| Penance | 12 s cooldown, three pulses (instant, then every 1 s for 2 s), 0.19, 126/200/268/425, costs 150 to 385 | 402174, 1240720, 1240721, 1316995 and their heal helpers 402289, 1240723, 1240724, 1316991 | aligned |
| Prayer of Mending | 172/298/413, 0.429, 5 jumps, 30 s, 10 s cooldown, costs 210/305/390 | 401859, 1240826, 1240827 (effect 1 amount 5), auras 401877 and 12408xx duration 30000, category 10000 | aligned |
| Lightwell | 225/295/365 mana, 1.5 s cast, 10 min, 180 s well, 5 charges, renew 160/233/320 every 2 s for 10 s with no coefficient | 724, 27870, 27871; 7001, 27873, 27874 | aligned |
| Holy Nova | damage 27 to 187 and heal 51 to 311, both 0.107, costs 185 to 750 | 15237 to 27801 and 23455 to 27805 | aligned |
| Inner Fire | 315/495/720/945/1170/1395 armor, 10 min, Improved +15% a rank | 588 to 10952 (aura 674), 14747 | aligned |
| Racial heals | Divine Grace at Desperate Prayer's amounts below 50% health; Contingency Plan ward 30 s, trips below 35%, shield 155 to 926, heal over time 25 to 134 every 3 s for 15 s at 0.2 | 1277370 to 1277378 (effect 1 amount 50); 1277462 to 1277640, 1277463 to 1277466, 1277456 to 1277460 | aligned |
| Dark Sacrifice | 5 ticks of 3 s, the rank's amount taken as health and given as mana plus Spirit once, 10 min | 1277324 to 1277328: period 3000, duration 15000, category 600000, amounts 80/136/196/258/320, **EffectRealPointsPerLevel 1, MaxLevel 28/38/48/58/68** | **D1: the growth was missing** |
| Healing talents | Improved Renew 5%, Holy Specialization 1%, Divine Fury 0.1 s, Inspiration 8/17/25, Improved Healing 5%, Searing Light 2/5, Litany of Light 5/10, Spiritual Healing 3/7/10, Spiritual Guidance 5% healing and 1/3/5/6/8% damage of Spirit, Renewed Hope 2% and 1 s a rank, Divine Aegis 5% a rank to 20 a level for 12 s, Holy Binding 50%, Twilight Focus 23/47/70, Mental Strength 2%, Meditation 17/33/50, Mental Agility 3/7/10, Silent Resolve 10% | the 70790 rank texts, line by line | aligned |
| Not modelled, nothing on a card depends on them | Power in Light (its 70790 text now also counts Shadow Word: Pain), Holy Precision, Martyrdom, Blessed Recovery, Spirit of Redemption, Blackout, Silence, Spirit Tap (no kill in a boss fight), Devouring Contagion's spread, Improved Inner Fire's charges, the Searing Light free Holy Nova | | named, as before |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| D1 | **Dark Sacrifice took and gave the rank's own-level amount at every level.** The client grows each rank 1 a level above its level to its max level (SpellEffect EffectRealPointsPerLevel 1 on both effects, SpellLevels MaxLevel 28/38/48/58/68; the generated `DarkSacrificePointsPerLevel` and `DarkSacrificeMaxLevel` already carried it), so an Undead between rank levels was short up to 8 a tick, 40 mana a cast: rank 1 at 28 is 88 a tick, not 80. At the rank levels themselves (20, 30, 40, 50, 60: every published band floor and raid 60) the row is the engine's, so nothing published moves; the ladder goldens move at rung 38. `mana_abilities.go` now reads the growth through `clientdamage.Roll`, as every damage spell in the package does. Red first: `TestDarkSacrificeTickGrowsOneALevelAboveItsRank` (13b9fcc59, red at compile since the helper did not exist), green at fe282fdfb. | 70790 SpellEffect rows of 1277324 to 1277328 | old (the racial's first model, round 11) | E fe282fdfb |

## Logs against the engine

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Improved Mind Flay compounds.** In 180 full, crit-free channels the second tick is 1.20 of the first (median; 1.18 to 1.23) and the third 1.44 (1.42 to 1.47). That is 1.2 and 1.2 squared, the engine's `ImprovedMindFlayRamp` at 2/2; an additive reading would give 1.40. Ticks land 1.00 s apart, the first 1.0 s after the cast. | Itso, 2839 | aligned |
| L2 | **Mind Flay ticks crit, at the priest's spell crit, for 1.5x.** 22 of 721 ticks crit (3.1%); a crit first tick is 1.25 of the non-crit second tick that follows it, which is 1.5 over the 1.2 ramp. The client marks every Mind Flay rank PERIODIC_CAN_CRIT and the engine's channel rule reads that flag. | Itso | aligned |
| L3 | **Renew ticks crit.** 53 of 977 ticks across six healers (5.4%), the same rate as their direct heals (Heal 5.5%, Lesser Heal 5.1%). Every trained Renew rank carries the client flag and the engine's heal-over-time rule reads it. | Net, Shadow, Luminare, Cogara, Hallow | aligned |
| L4 | **Shadow Word: Pain, Devouring Plague and Holy Fire ticks crit.** 34 of 686 Pain ticks (5.0%), 3 of 68 Plague ticks, 9 of 145 Holy Fire ticks (6.2%), at 1.48x (Pain, 10 crits). The owner's rule (non-channelled DoTs crit) and the client flag agree. | all priests | aligned |
| L5 | **Vampiric Embrace heals exactly 20% and never crits.** 80 of 80 Shadow damage events inside a Vampiric Embrace window produce one heal per party member within 60 ms, 0.20 of the damage (0.19 to 0.22 after rounding); 0 of 488 heal or overheal events crit. The engine's share is 0.20 and its heals roll no crit. | Itso, 15290 | aligned |
| L6 | **Shadow Weaving's debuff (Shadow Vulnerability 15258) stacks to 5 and is added by a landed Shadow Word: Pain, Devouring Plague, Mind Blast and Mind Flay cast**, not by each tick: 348 applications and refreshes against about 400 such casts and hits, stacks 1 to 5 seen, 15 s. The engine adds a stack on exactly those events. | Itso | aligned (rank unknown, so the per-rank chance is not read) |
| L7 | **One spell power explains every Shadow spell.** At Itso's level (28 to 33 by his ranks) a Pain tick of 48, a first Mind Flay tick of 49 and a Mind Blast of 192 (medians, non-crit) all imply about 65 to 70 spell power under five Shadow Weaving stacks through the engine's bases (30, 34 and 151 to 158) and coefficients (0.2, 0.167, 0.429). | Itso | aligned (spell power not in the log) |
| L8 | **One spell power explains every Holy spell.** Emanuele (22 to 27): Holy Fire ticks of 6 (4 + 0.05 x 40), Holy Fire hits of 87, Smites of 90, Mind Blasts of 118, Pain ticks of 27 and Plague ticks of 21 all sit on about 40 spell power through the engine's rows. Holy Fire's periodic coefficient 0.05 and base 4 are confirmed by the tick alone. | Emanuele, 2848 | aligned |
| L9 | **Direct crits are 1.5x.** Smite 1.49 (10 crits), Shadow Word: Pain 1.48, Mind Blast 1.43 (2 crits). | Itso, Emanuele | aligned |
| L10 | **Mind Blast's cooldown.** The shortest gap between two Mind Blast hits is 11.0 s in a log with a 1.5 s cast and travel; no gap under 8 s. | Itso | aligned (not binding) |
| L11 | **Holy Fire's duration.** 35 of the casts that ticked at all gave 3 ticks and 19 gave 5; targets die under it. Ticks arrive every 2 s. | Emanuele | aligned |
| L12 | Spirit Tap procs on kills (25 gains for Itso), never inside a boss fight; the engine never activates it. Touch of the Grave never crits. | Itso, Net | aligned |
| L13 | **Heals scale with the target's level (round 20 P2)**, Eureka! on a Gnome Shadow log (P3), Lightwell Renew's crit source (P4), Dark Inspiration's Spirit base (round 20 S2), Penance's three pulses, Prayer of Mending's jumps, Renewed Hope, Divine Aegis, Inspiration, Shadow Word: Death's backlash and Shadowform: no public log has a priest at 40 or above, and no Gnome priest logged. | foreverlogs actor lists | held (no log) |

## Builds, rotations and rank rows

Nothing in the client moved, so the round 24, 25, 29 and 30 judgements stand: the Shadow card (Pain with 10 s left, Vampiric Embrace, Devouring Plague on its cooldown, Inner Focus into Mind Blast, Mind Blast, Shadow Word: Death on cooldown, Mind Flay, Shoot), the Holy and Discipline lists, the three leveling orders and the 33 priest rank rows. D1 changes an Undead's mana only between rank levels; the rank rows at those levels were judged with 10,000-iteration pairs whose margins (0.1 to 1%) sit well above the mana D1 adds (see below), and the raid and band-floor numbers are untouched, so no row was re-judged.

## Before and after

Undead, bare band gear, the guide's leveling talents at each level, 3,000 iterations, seed 1, engine e586233d7 against fe282fdfb (the only change is D1). Every level below sits between two rank levels, where D1 adds 5 to 8 mana a tick to the one Dark Sacrifice the card casts; the figures are relative to the old engine, with the sim's own error.

| spec, level | before | after | gain | error |
|---|---|---|---|---|
| Shadow bare 20 at 25 | 43.86 | 43.93 | +0.18% | 0.04 (beyond) |
| Shadow bare 30 at 35 | 77.21 | 77.28 | +0.09% | 0.07 (at the edge) |
| Shadow bare 40 at 45 | 187.62 | 188.01 | +0.21% | 0.25 (within) |
| Shadow bare 50 at 55 | 275.95 | 276.56 | +0.22% | 0.40 (within) |
| Shadow raid 60 | 718.86 | 718.86 | 0 (Dark Sacrifice is never cast) | 0.47 |
| Holy bare 30 at 35 (effective HPS) | 89.40 | 89.59 | +0.21% | 0.05 (beyond); mana lasts 148.6 to 149.8 s, the cast gives 808 to 832 mana |
| Holy bare 50 at 55 | 331.81 | 332.52 | +0.21% | 0.26 (within); 1,520 to 1,545 mana |
| Discipline bare 30 at 35 | 100.69 | 100.84 | +0.15% | 0.03 (beyond) |
| Discipline bare 50 at 55 | 224.67 | 224.91 | +0.11% | 0.12 (within) |

At the band floors (20, 30, 40, 50, 60) the rank's own-level row is what both engines read, so every published BiS, weight, race and rank-row number is unchanged. The three priest ladder goldens (an Undead, rungs 10 to 60) move at rung 38 only, where the rank 2 cast gives 144 a tick instead of 136: Discipline 59.7 to 59.9, Holy 61.1 to 61.6, Shadow 64.1 to 64.2 (300 iterations); they are regenerated on the box in the pin commit.

## Guide and card

No page text quotes a Dark Sacrifice amount: the guides say "well over a thousand mana once a fight" and the card notes say "it turns your own health into mana over 15 seconds, plus your Spirit". Both stay true. No build, card or rank row moved, so there is no guide prose to ship and no wow-player review this lane.

## Gates

- Fork, on the box (`--engine-sha fe282fdfb`): `go test -tags with_db` for sim/priest, sim/priest/healing and sim/priest/shadow (ok), and sim/core (ok). Without `-tags with_db` sim/priest/healing's `TestHealerFiresASpellPowerTrinket` panics on item 18820 at the pin too; it needs the tag, as its message says.
- Site, on the box: `npm run lint` (0 errors), `astro check` (0 errors), `vitest run src/lib/sim` (the pin's version and fixture assertions), `go test ./request/ -run TestRotationLadder` green on the old pin, red for the three priest goldens on the new one until the rung 38 regeneration, then green again at the pin commit. No APL, guide, rank row or Playwright surface changed, so apl-check, the guide guards and the e2e specs were not run.

## Measure at refresh

Nothing. No published band or raid number moves (D1 lands between rank levels only).

## For the owner

- The first level-60 priest logs (a Shadow priest with Shadowform and Shadow Word: Death, a Holy or Discipline priest with Penance, Prayer of Mending and Lightwell) would close L13's seven held checks. Our own report pipeline with advanced logging is the best source.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 4 to 6); Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092; development notes, https://us.forums.blizzard.com/en/wow/t/2360696 (post 9).
- wago.tools DB2 CSVs for 1.60.1.70790 and 1.60.1.70888, hotfixed and shipped.
- `data/builds/1.60.1.70790/spellconst/priest.json`, `talents/priest.json`.
- foreverlogs.gg reports 2815 to 2850 (`/api/reports/public`, `/api/reports/<id>/meta_actors`, `/api/reports/<id>/meta_events?sourceCharacterIds=<id>`).
- Wowhead Forever tooltip 1309636 (Shadow Word: Death's backlash text).
