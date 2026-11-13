# Sim accuracy 2026-11-13: Arcane mage and Shadow priest on the corrected 70938 data

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-arcshadow`. Site from origin/main a544d507a, fork from master 5603877e4 (the pin), rebased onto f17711960 for the landing. Client: 1.60.1.70938 with the corrected hotfix layer (273623f0f, `item-hotfix.md`). Scope: Arcane and Shadow at bare 50, bare 60 and raid 60. Both last had a full pass on 12 Nov (`2026-11-12-sim-accuracy/arcane.md`, `shadow.md`), before the client move and before the push layer read anything on 70938.

**Totals: 2 fixed, both old, both shared engine. M1 (mana model): a mage's four gems and the Demonic Rune share one two-minute category (ItemEffect SpellCategoryID 1153), so a 180 s fight has room for two of them, and the engine popped whichever gem first fit the deficit: in practice the Ruby and the 375 to 425 mana Agate, with the Demonic Rune the raid preset carries never used. The gems now go largest first, and with a rune in the conjured slot the gems below the biggest stay in the bag (Ruby, then rune). Arcane raid 60 +0.46 to +0.59%, bare 60 +0.60 to +0.80%, bare 50 +0.15 to +0.55%; Fire raid 60 +0.25 to +0.33%, bare 50 and 60 +1.0 to +1.5%; Frost bare 50 +0.5%, raid 60 and bare 60 -0.0 to -0.2% (it drains slowly, so a rune or Citrine two minutes after the Ruby is sometimes not needed by then). S1: Mental Agility's cost mask names Shadowform (15473) and the engine's mask left it out; the form's prepull 40% of base mana now costs 10% less with 3/3. No Shadow number moves (the bar refills before the pull). Every other spell, talent, proc, aura and buff the two raid cards and builds touch matches the 70938 rows and their push. Engine pin 5603877e4 (f17711960 on main by the landing) to 037798db1.**

## Method

1. **70938 against 70790 for both classes.** `spellconst`, `talents`, `trainables` for mage and priest, `sets.json`, `simbuffs.json`, `simconsumes.json`. The corrected push layer (`hotfix-push.json`, `raw/wago-push/`): the rows it changes, and every withheld (invalidated or not public) record of SpellEffect, SpellCooldowns, SpellPower, SpellMisc, SpellAuraOptions, SpellCategories, SpellLevels and SpellClassOptions mapped back to its spell through the hotfixed 70938 CSVs from wago.tools (`logs/rows.py`).
2. **Engine against the rows.** Re-read at the pin: `mage/{arcane_blast,arcane_missiles,arcane_wrath,missile_barrage,mana_gem,evocation,talents,ignite}.go`, `priest/{mind_blast,mind_flay,shadow_word_pain,shadow_word_death,devouring_plague,vampiric_embrace,talents,priest}.go`, the class masks of every talent modifier the two raid builds take, decoded against each spell's 70938 family mask, and `core/consumes.go` (runes and potions), `core/debuffs.go` (Judgement of Wisdom).
3. **The mana model on the 180 s raid fight.** `spec-breakdown` on every published entry at raid 60 and bare 20 to 60, both factions, 10,000 iterations (5,000 in the sweeps), seed 1, reading cast counts, conjured items used, out-of-mana time and the mana ledger (`baseline.sh`, `bands.sh`, `results/`).
4. **Logs.** Our reports through the Cloud SQL proxy and R2 (read only, every fight's events of all 8 reports to 13 Nov 18:38, including the new ju2t73py6pib). foreverlogs.gg's API answered every request on 13 Nov 23:00 UTC with 403 "API access not allowed" (browser user agent and referer too), so no public report was read this lane. Private reports are counted only in aggregate; no player is named.
5. **Compute.** Every sim, Go test and golden on the Windows box (`tools/remote-test.sh --engine-sha`, `fork_tests.sh` over ssh), one at a time, `nice -n 15`, GOMAXPROCS=5. The box ran active build 1.60.1.70938 with simdb sha256 ea39a6a17bfa, the same file as origin/main's.

## Client: 70790 against 70938, and the push

- **Mage:** every spell row the Arcane card or either build uses is unchanged but for float formatting (the fractional-values lane's floats). Two new NPC rows (Frost Armor 1283419, Blast Wave 1291176), neither a player spell.
- **Priest:** float formatting, and Divine Spirit and Prayer of Spirit's family mask gaining word 1 0x400000 (a raid buff the engine applies as stats; no modifier reads that bit).
- **Talents and trainables:** the build stamp only.
- **The push on 70938:** 20 rows change (the Stormstrike, Maelstrom Weapon and Mental Dexterity curves and effects, Gloves of Holy Might's proc mask, Wendel Mathers' dummy). None is a mage or priest row, and no withheld record maps to a mage or priest spell or talent.
- **Consumables:** eleven new "Endless" elixirs and scrolls (289081 to 289334, quality 1, item level 30). No loot, vendor or recipe source in the data, so they stay out of the preset.

## Engine against 70938

| area | engine | client 70938 | verdict |
|---|---|---|---|
| Arcane Blast | 15% of base mana, +100% a stack, 4 stacks, 12% of Intellect as spell damage a stack, 18 s | 400574 to 1239700 cost_pct 15; 400573 aura 174 12, aura 108 100 misc 14, CumulativeAura 4 | aligned (logs L1) |
| Arcane Missiles | each missile a direct hit that can crit; 0.286 a missile | crit-flags.md row; missiles' SpellEffect | aligned (L3) |
| Missile Barrage | 40% on Blast, 20% on Wrath, Fireball, Frostbolt; -100% cost, 0.5 s ticks | 400588 aura 42 40; 400589 -100 misc 14, -50 misc 1 | aligned |
| Arcane Concentration and Clearcasting | 10% at 5/5, 1 s proc cooldown, next damage spell free | 11213 ProcCategoryRecovery 1000; 12536 -1000 misc 14, 1 charge | aligned (L4) |
| Arcane Wrath | 451 at 60, 6 s, 395 mana, coefficient 1.0, +5% of max mana as spell damage | 1326082; effect 1 dummy 5; SpellCooldowns 102218 to 102221 6000 | aligned (L2) |
| Arcane Power, Presence of Mind | +30% damage and cost, 15 s, 3 min; next cast instant, 3 min | 12042, 12043 | aligned |
| Arcane Mind, Instability, Impact, Meditation, Incineration, Accelerant, Ignite, Heating Up, Improved Fire Blast, Precision | per rank as the 12 Nov table; Ignite's ticks flat (crit-flags B1) | 70938 trait values: 10% Int and +100% crit bonus, 3%, 6%, 50%, 10%, -0.5 s, 40%, -25% a stack, -2 s, 6% | aligned |
| Evocation | 8 s, +1500%, all regen while channelling, 8 min | 12051 | aligned |
| Mana gems and runes | one category, two minutes | ItemEffect 5514, 5513, 8007, 8008, 12662, 20520: SpellCategoryID 1153, 120000 ms | category aligned; **M1: the order** |
| Shadowform | +10% Shadow, -50% Shadow cost after the spell mods, +100% crit bonus, 40% of base mana | 15473 | aligned |
| Mental Agility | -10% cost at 3/3 on SW:P, DP, VE, Smite, Holy Fire | 14520 mask also names Shadowform | **S1** |
| Twin Disciplines | +5% on SW:D direct and SW:P, DP, Mind Flay periodic | 1225132 e0 misc 0 reaches SW:D, e1 misc 22 reaches SW:P, DP, MF, Starshards | aligned |
| Devouring Contagion, Improved SW:P, Improved Mind Blast, Darkness, Improved Mind Flay, Early Demise, Inner Focus | -50% DP cost; +6 s; -2.5 s; +10%; 1.2 a tick, compounding; +30% crit at 20%; free MB with +25% crit, no periodic crit | 1309950, 15275, 15273, 15259, 1225139, 1310076, 14751 | aligned |
| Shadow Weaving | 100% at 3/3, +2% a stack to the priest's Shadow damage, 5 stacks, 15 s | 15257 proc 0x10000; 15258 aura 270 2 misc 32, CumulativeAura 5 ("the Shadow damage the target takes from you") | aligned for one priest on one target |
| Vampiric Embrace | 20% of landed Shadow damage to each party member, 30 s, 30 s cooldown, 40 mana, never misses | 15286 | aligned. Group value: the raid 60 card's VE heals 150 a second raw to each party member (27,000 a fight) for one global every 30 s; the raid search's +3.7 to +4.0% for dropping it is that global (raid-searches.md), held |
| Devouring Plague, SW:P, Mind Blast, SW:D, Mind Flay | rank 6 106 a tick for 24 s on a 24 s cooldown; 18 s; 8 s; 18 s with 30% backlash; 3 ticks, 0.167 | 19280, 10894, 10947, 1309636, 18807 | aligned |
| Judgement of Wisdom | 50% on any landed direct hit, spells and each missile included | 20355 ProcTypeMask 0x222A8 (magic spells taken included), trigger 1826 a server dummy | aligned with the healer-mana lane's rule; no JoW energize in our logs to measure |
| Raid buffs and consumables | the curated preset | presets.json; the 70938 consumables | aligned (Arcane casts no Spellblasting Potion; its mana potion is worth more) |

## Logs

The highest-level casters in our reports are an Arcane mage at a scaled level 33 (132 casts; Missiles rank 4, so at least 32), a second at 27, and a Shadow priest at 32 (no Shadowform yet). No caster at 50 or 60 has been uploaded, and foreverlogs.gg is closed to the API. The advanced log carries current mana before the cast, the cost and the spell power, so this lane read costs and damage directly for the first time.

| # | check | evidence | verdict |
|---|---|---|---|
| L1 | Arcane Blast's cost rises 100% a stack | the first three Blasts of a chain take 104, 206 and 311 mana off the bar (691 base mana x 15% = 104, then x2, x3); the log's cost field shows 104 every time, the bar shows the stacks | aligned |
| L2 | Arcane Wrath adds 5% of max mana as spell damage | 140 mana; non-crit hits 344 to 358 at 43 spell power and 3,576 max mana, against the engine's 341 to 365 (rank 1 at 33, 119 to 143, plus 43 + 178.8); the one 395 came on Blast stacks | aligned |
| L3 | Missiles take 0.286 of spell power and the stacks | non-crit missiles at 79 with no stack (66.6 + 0.286 x 43 = 78.9) and steps of 7 a stack | aligned |
| L4 | Clearcasting frees the next damage spell | cost 0 and an unchanged bar on every Blast cast under it | aligned |
| L5 | Shadow ticks crit | SW:P 16 of 279, DP 12 of 154, Mind Flay 8 of 72 | aligned (crit-flags) |
| L6 | Gems and runes at 60 | no 60 in any log | Open |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| M1 | **The mana gems wasted the shared category.** Every gem was its own major cooldown that popped once the deficit passed its own maximum, all on the conjured timer the Demonic Rune also uses. The Agate's threshold (425) is met first, so it took a two-minute slot from the Ruby, the Citrine or the rune. Published raid 60 Arcane: Ruby 1.00 and Agate 0.95 a fight, the rune 0; bare 60: Ruby and Agate 0.75. A gem now waits while a bigger gem is unused, and with a Demonic Rune in the conjured slot (900 to 1500, more than every gem below the Ruby, and there for every use) the smaller gems stay unused: raid 60 is Ruby then rune (1.00 and 0.98), bare 60 Ruby then Citrine, bare 50 Citrine then Jade. The guides already told players to do this. | ItemEffect 1153 rows; `results/baseline-5603877e4.out` | old | E 8418c94fc (red: `TestManaGemsGoLargestFirst`, `TestManaGemsBelowTheBiggestHoldForADemonicRune`), ab166b582, 286ad19e6 (TestP1Mage results) |
| S1 | **Mental Agility missed Shadowform.** 14520's cost modifier mask names 15473 (word 0 bit 31); the engine's mask did not, so the form cost 550 instead of 495 with 3/3. Shadowform is cast 20 s before the pull and the bar refills, so no Shadow cell moves (identical to the hundredth, `results/shadow-bands-*`). The one figure it moves is the Shadow ladder golden's level-60 rung (a priest in a weapon and a wand with no Intellect gear, whose bar does not refill before the pull): 310.7 to 312.1. | 14520, 15473 | old | E d98bec9b4 (red: `TestMentalAgilityCutsShadowformsCost`), 037798db1 |

## Before and after (M1)

Published entries, 5,000 iterations, seed 1, 5603877e4 against 9f34117f9 (the fix alone); error is the sum of both runs' (`results/mage-bands-*.out`).

| entry | Alliance | Horde | error |
|---|---|---|---|
| Arcane raid 60 | +0.46% | +0.59% | 0.16% |
| Arcane bare 60 | +0.80% | +0.60% | 0.21% |
| Arcane bare 50 | +0.15% | +0.55% | 0.23% |
| Arcane bare 40 | -0.16% | -0.10% | 0.24% |
| Fire raid 60 | +0.25% | +0.33% | 0.16% |
| Fire bare 60 | +1.21% | +1.53% | 0.27% |
| Fire bare 50 | +1.03% | +1.21% | 0.26% |
| Fire bare 40 | -0.04% | -0.03% | 0.26% |
| Frost raid 60 | -0.09% | -0.13% | 0.14% |
| Frost bare 60 | -0.02% | -0.24% | 0.18% |
| Frost bare 50 | +0.55% | +0.54% | 0.22% |
| Frost bare 40 | +0.15% | +0.05% | 0.22% |
| every mage at bare 20 and 30 | 0 | 0 | |

Out of mana at Arcane raid 60 falls from 1.2 to 1.0 s, at bare 60 from 12 to 10 s. Frost at 60 loses within its error at raid and just past it at bare 60 Horde: it drains slowly, so the Ruby comes late and the rune or Citrine after it is used in 41 to 83% of fights, where the old order's early Agate always fit. The Frost guide now says to spend the second one only once you are well down. The fork suite's 300 s no-buff Frost and Arcane fixtures fall 1.5 to 3.1% for the same reason (a late Ruby leaves no third window); no published fight is 300 s.

## Held and open

- **Frost's second conjured slot (Open).** An earlier pop (a smaller gem when waiting would cost a window) would win back Frost's 0.1 to 0.2% at 60; it needs a rule on fight length and spend time, and no Frost log at 60 exists to set it.
- **Vampiric Embrace (Held, unchanged).** The raid searches' +3.7 to +4.0% for dropping it is the global it costs; the group heal is the reason the card keeps it.
- **Raiments of Conviction's 5-piece (Open, 12 Nov T1).** Still for the raid lane before 9 Dec.
- **Logs at 50 and 60 (Open).** None in our reports; foreverlogs.gg's API refused every request this lane.

## Measure at refresh

- **MR109**, the mages on the largest-first gems: mage-arcane, mage-fire, mage-frost raid 60 and bare 40 to 60, both factions: picks, weights (Intellect and mana regeneration against spell power), set_dps; the leveling walks from 38.

## Guide and card

The three mage guides' "Mana in a 3 minute fight" paragraphs: Ruby first, then a Demonic Rune two minutes later or the Citrine without one, and a smaller gem spent first wastes one of the two slots; Frost spends its second only once it is well down. The card notes (Evocation, Shoot) already say gem, potion and rune and stand. Shadow: no text names Mental Agility on the form, so none changes. wow-player: Arcane and Fire SHIP, Frost HOLD then SHIP, Shadow SHIP (`arcane-shadow-70938-wow-player.md`).

## Gates

On the box at pin 037798db1 (`gates.sh`): fork mage, priest (all three packages) and core suites green, the TestP1Mage results regenerated; site lint, astro check, the sim and guide vitest, apl-check, go vet, the request suite with the ladder goldens and the adapter goldens rewritten on the new engine, and the leveling, score, rank-choice and tool tests. Results under `results/`.

## Shas

- Fork `lane/sim-arcshadow` on master f17711960: 8418c94fc, ab166b582, 286ad19e6, d98bec9b4, 037798db1; master fast-forwarded to 037798db1, on top of the Enhancement lane's f17711960 (pinned on the site while this lane ran).
- Site: the pin (f17711960 to 037798db1), the three mage guides, the ladder goldens (mage-arcane, mage-fire, mage-frost and priest-shadow: one rung each at 60 or the gem levels), this report, `arcane-shadow-70938/` and the ledger row.
