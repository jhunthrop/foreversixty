# Sim accuracy 2026-11-11: Fire mage, Frost mage and Balance druid

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane: sim-casters. Site from origin/main cfe2a6f67, fork from master e586233d7 (the pin); rebased on origin/main 46e750bbe and fork master 3bdfab5b6 (the paladin lane) before the push. Client on the site: 1.60.1.70790; live client 1.60.1.70888 (Blizzard's 6, 9 and 10 Nov hotfixes are hand rows in the engine until the client move, ledger MR41 and MR46). The three specs were last audited spell by spell in rounds 23 and 24; round 28 walked Frost's order and round 29 every order; the hotfixes lane landed Arcane Concentration's split today.

**Totals: 3 confirmed bugs, all fixed. One engine (Hurricane's tick table sat 2 above the client at every rank; no card casts it, so no published number moves), one tool (spec-breakdown simmed every raid entry without its published enchants, 5 to 10% under the site's own set_dps), one rank row (Balance's Insect Swarm rank 3 row at 54 to 56 lost to rank 4 on round 29's order, +0.5 to +1.0% at 53 to 58). Engine against the client: every amount, coefficient, cost, cast time, cooldown, tick count and talent value the three rotations touch matches 1.60.1.70790 and the hotfixed talent text. Engine against real fights: the first public Frost log with Fingers of Frost (Gnomeregan, foreverlogs.gg 2839) agrees with the engine on the proc rate, the charge spend, Ice Lance's frozen multiplier and coefficient, Ice Shards' crit size and every Frost amount. No Fire log with Ignite or Heating Up and no Balance log at 40 or above exists yet. Engine pin e586233d7 (3bdfab5b6 after the paladin lane) to 8dacbe217.**

## Method

1. **Blizzard since 4 Nov.** "WoW: Forever Hotfixes" (t/2403040) read to its newest post (post 6, 11 Nov 03:27 UTC: the druid Omen of Clarity and Nature's Swiftness lines the hotfixes lane already carries); "Class Tuning Incoming" (t/2404092) posts 1 to 20; the development notes thread (t/2360696). No Fire, Frost or Balance line since 10 Nov.
2. **Client.** `data/builds/1.60.1.70790/spellconst/{mage,druid}.json` for every spell the three cards cast or the builds talent (Fireball, Scorch, Pyroblast, Fire Blast, Ignite, Combustion, Heating Up, Fire Power, Master of Elements, Frostbolt, Ice Lance, Cone of Cold, Ice Barrier, Frostfire Bolt, Fingers of Frost, Arcane Missiles, Arcane Blast, Missile Barrage, Clearcasting, Arcane Power, Evocation, Blizzard, Flamestrike; Wrath, Starfire, Moonfire, Insect Swarm, Innervate, Moonkin Form and Aura, Eclipse, Nature's Grace, Omen of Clarity, Faerie Fire, Hurricane), `talents/{mage,druid}.json` for every rank's text, and `hotfix-push.json` (the 70790 push carries no mage or druid class row).
3. **Engine.** Every mage and druid file the three specs touch in the fork at e586233d7 (`sim/mage/*.go`, `sim/druid/*.go`, `sim/druid/balance`, `core/debuffs.go`, `core/spell_result.go`, `core/spell_outcome.go`, `core/buffs.go` Innervate), read line by line against the rows above.
4. **Logs.**
   - Our API's reports: gxrgvz4yzsi2 (8 Nov, advanced logging, a city realm) holds 649 mages and druids at levels 9 to 28 with a handful of casts each (one Frostbolt rank 5 sample at 87 spell power matches the client roll); z2opt2x2gs5u adds nothing.
   - foreverlogs.gg: all 24 public reports (2815 to 2850, 5 to 11 Nov) through the report API with browser headers (`/api/reports/public`, `/casts`, `/character_spell_damage`, `/encounters/<id>/timeline`), paced at 4 to 5 s a call. Frost mages in 2815, 2818, 2823, 2827, 2828, 2839, 2848; a Fire mage in 2836 (Deadmines, about level 20) and 2834; Balance druids in 2818 (level 14 to 20, Wrath only), 2820 (a bear) and 2827/2836 (healing). The one caster with a Frost build deep enough to matter is Anaking in 2839 (Gnomeregan, 9 Nov, Ice Lance rank 2, Fingers of Frost, Frostbite, Ice Shards 5/5, no Winter's Chill): 18 encounters read event by event (714 of the mage's events).
5. **Sims** on the Windows box (`tools/remote-test.sh`, one `spec-breakdown` at a time, `nice -n 15`, GOMAXPROCS=5, `make simdb` first because the module embeds the build's enchants.json and spellranks.json): raid 60 both factions at seeds 1 and 50001, 10,000 iterations; the Insect Swarm rank sweep at 53 to 58 for a night elf and a tauren, seeds 1 and 50001, 10,000 iterations; the published raid entries re-simmed at the BiS commit (2de65052d) and the BiS engine (0e04b1284) to chase the set_dps gap.

## Confirmed and fixed

| # | where | finding | source | fix |
|---|---|---|---|---|
| E1 | engine, druid | Hurricane's tick table read 70, 100 and 134 a tick with the client's growth and caps; the client's tick spells (1278965, 1278968, 1278759, the rows the generated `Hurricane*` arrays already carry) are 68, 98 and 132. AoE only; no card casts it and the sim is single target, so no published number moves. | spellconst 1278965/1278968/1278759 | fork c6b29db31 (red, a build failure on the box), 8dacbe217: `hurricane.go` reads the generated rows through `clientdamage.FromTable`; `TestHurricaneTickDamageMatchesClient` |
| T1 | tool, sim/cmd/spec-breakdown | A raid entry's breakdown dropped the entry's enchants: `bisSlot` had no `enchant_id`, so the gear went to the engine bare while the ranker measures `set_dps` wearing them (leveling-bis `withRaidEnchants`). Fire raid 60 Alliance read 751 against the page's 820 (-8.4%), Frost 628 against 685 (-8.3%), Balance 655 against 731 (-10.4%), Combat rogue 852 against 902 (-5.5%), all on the BiS commit and the BiS engine, so neither the engine nor today's data run moved anything. Every raid-60 before/after in the 2026-11-11 lanes ran on bare gear; their deltas are relative and stand. | published `bis/*.json` slot rows carry `enchant_id`; `spec-breakdown` header promises equality with `set_dps` | site 38dfb5068 (red, a build failure on the box), 10b49b36a: `bisSlot.EnchantID`, gear carries `Enchant`; `TestRaidEntryWearsItsPublishedEnchants`, `TestBareEntryWearsNoEnchant` |
| R1 | rank row, Balance | The round 19 Insect Swarm rank 3 row at 54 to 56 (re-judged round 24 on the old order) loses on round 29's order (Heart of the Wild at 42 to 46, Naturalist at 47 to 51): rank 4 beats rank 3 at 53 by +0.5/+0.7% (Alliance seeds 1/50001) and +0.5/+0.7% (Horde), 54 by +0.8/+0.7% and +0.7/+0.7%, 55 by +0.7/+0.7% and +0.7/+0.7%, 56 by +0.6/+0.8% and +0.7/+0.7%, 57 by +0.7/+0.8% and +0.7/+0.7%, 58 by +1.0/+1.0% and +0.9/+1.0% (errors 0.04 to 0.07%). The round 30 lane re-judged Wrath and Starfire on the moved levels but not this row. | `spec-breakdown -level -rank 24977=3` against `=4` | site f927fc40d (red: `TestBalanceRankRowsFortyToFiftyNine` at 54 to 56 wants rank 4, six cases), 4c94cc933: the row is gone from `rank-choices.json` (synced copy), balance.md's rank list stops at 38 |

## Checked, nothing to change

**Fire.** Fireball ranks 1 to 12 (amounts, growth, caps, coefficient 1.0 from rank 5, costs, 3.5 s, the flat DoT ticks 1 to 15 at 2 s over 4/6/6/8 s, no spell power on the DoT, which the core snapshot honours since `Dot.BonusCoefficient` is the dot's own); Scorch 1 to 7 (0.429, 1.5 s); Pyroblast 1 to 8 (1.0; DoT 11 to 53 a tick at 0.15, four 3 s ticks; 6 s cast); Fire Blast 1 to 7 (0.429, category 19's 8 s shared with Ice Lance); Flamestrike direct 0.157 and ground ticks 0.032 (1279983 rows); Blast Wave; Ignite 8% a rank over two 2 s ticks, rolling, no target double dip (412538); Heating Up 25% a stack to 3, 20 s, a Fire Blast hit a stack, a Presence of Mind Pyroblast spending none (400625); Fire Power 1/2/3% an effect to 4/8/12% counting the mage's own Fire dots and Improved Scorch; Improved Scorch 33/67/100% for 3% a stack to 5 for 30 s (core 12873); Master of Elements 10% a rank of the base cost on Fire and Frost crits; Critical Mass 2% x3, Incineration 2% x5 on Fire Blast, Ice Lance, Arcane Blast and Scorch, Accelerant 0.1 s x5 on Fireball, Frostfire Bolt and Arcane Wrath, Improved Fire Blast 1 s x2, Burning Soul 10% threat x3, Combustion 10% crit a hit until 3 non-periodic crits (28682 cap 10, 3 min); Arcane Blast 12% Intellect a stack to 4 for 18 s, +100% cost a stack, 15% base mana, 1.5 s (400573, 400574 to 1239700); Missile Barrage 10/20% on Fireball, Frostbolt, Frostfire Bolt and Arcane Wrath, 20/40% on Arcane Blast, the free half-time channel (400589); Arcane Concentration's split over a Missiles channel (hotfixes lane M1); Arcane Power 30/30 for 15 s; Evocation 8 s at +1500% on 8 min; Mage Armor 50% regen while casting (aura 134 amt 50 on every rank).

**Frost.** Frostbolt 1 to 11 (0.814 from rank 5, 3 s, binary); Ice Lance 1 to 6 (every rank now carries the 8 s category cooldown in the 70790 tables, 300% to Frozen, coefficient 0.107 from the beta logs: see the log section); Cone of Cold 1 to 5 (0.129, 10 s, chill bit, the reach rule); Ice Barrier 15% and 5% from the hotfixed text (the shipped rows still read 5 and 5), 30 s, a minute; Frostfire Bolt 1 to 3 (0.814, 3 s, the 9/13/19 flat ticks over 9 s); Fingers of Frost 15% a chill, 1/2 charges, 15 s, spent as the cast resolves; Frostbite 5% a rank, a root a boss ignores; Shatter 17/33/50; Ice Shards 20% a rank; Frost Channeling 5% cost and 10% threat a rank; Improved Frostbolt 0.1 s a rank; Winter's Chill 20% a rank to apply, 20% Cone of Cold damage and 7% spell damage a stack to 5 (hotfixed text; 12579); Improved Cone of Cold 12/23/35; Arctic Reach 10% a rank; Precision 2% a rank on every school; Blizzard ticks 0.042 with the Improved Blizzard chill rolling Frostbite at the rank's level-gap share.

**Balance.** Wrath 1 to 8 (0.429/0.486/0.571, 91 at 54 with 1.0 a level to 60, costs 10 to 120, 1.5 to 2.0 s); Starfire 1 to 7 (1.0, 533 plus 3.4 a level at 60 to 66, 3.5 s, costs to 340); Moonfire 1 to 10 (0.15 direct, 0.13 a tick, 3 s ticks over 9/12 s); Insect Swarm 1 to 5 (8 to 31 a tick at 0.158, 2 s ticks over 12 s, the -2% hit); Innervate 6 min, 20 s, +400% regen and full regen while casting, 5% base mana; Moonkin Form 35% base mana, the 3% aura once and exclusive with Leader of the Pack, Omen's chance doubled and its cooldown halved (24858 effects 4 and 5); Eclipse 0.17/0.33/0.5 s off the next two Starfires a Wrath, four charges, 15 s; Nature's Grace 10% haste and 10% off the global for 3 s on every non-periodic crit, from the hit for a Wrath; Improved Moonfire 5% damage and crit a rank; Moonfury 2% a rank on both schools; Vengeance 20% a rank including Insect Swarm; Moonglow 8/17/25; Improved Wrath 0.1 s and 10% a rank; Genesis 2% a rank (hotfixed text; the row reads 1); Nature's Majesty 2% a rank; Nature's Reach 2% hit a rank; Nature's Splendor one tick on Moonfire and Insect Swarm; Improved Starfire 0.1 s a rank (the stun is not modelled: a boss ignores it); Naturalist 2% a rank; Living Spirit 5% a rank; Omen of Clarity at the beta logs' 4% a spell with a 10 s cooldown; Faerie Fire 505 armor at rank 4, no swing reset; the melee swing stays off at 30 yards (`core/attack.go` gates it at `MaxMeleeAttackDistance`).

**Known rulings re-read and standing:** Ice Lance waits Fire Blast's 8 s category cooldown (every Ice Lance rank now carries `cd0/8000` in the 70790 tables); Leader of the Pack and Moonkin Aura are exclusive (`core.RaidHasDruidCritAura`, the raid preset lists both and the form adds no second 3%); Ignite does not double-dip (`SpellFlagIgnoreTargetModifiers`, snapshot multiplier 1).

## Logs against the engine

**Frost, Anaking (foreverlogs.gg 2839, Gnomeregan, 9 Nov; Frostbolt rank 5, Ice Lance rank 2, Cone of Cold rank 1, Frost Nova rank 2, Blizzard rank 2; Ice Shards 5/5, Fingers of Frost, Frostbite; level about 30 to 33, spell power about 88 read off the Frostbolt hits).** 18 encounters, 714 of the mage's events.

| what | log | engine | verdict |
|---|---|---|---|
| Frostbolt rank 5 non-crit | 171 to 195 (57 hits), median 181 | 104 plus 1.3 a level to 30, 9.1% wide, plus 0.814 x spell power: 99 to 119 plus about 72 | aligned at 88 spell power |
| Frostbolt crit | 342 to 386, median 366, 2.02 x the non-crit | 1.5 x (1 + 5 x 20% Ice Shards): 2.0 x | aligned |
| Ice Lance rank 2 unfrozen | 44 to 52 (10 hits) | 36 plus 0.4 a level to 32, 16.7% wide, plus 0.107 x 88: 41 to 53 | aligned; a 0.429 coefficient would read 70 to 82 |
| Ice Lance frozen (a Fingers of Frost spend 0.6 to 1.5 s before it) | 193, 200, 210; crits 375 to 421 | 4 x the unfrozen roll; crits 2 x that | aligned (one 95 is a frozen hit half resisted) |
| Fingers of Frost per chill | 16 fresh grants over 95 Frostbolt and Cone of Cold landings (16.8%); 7 after 66 Frostbolts (10.6%) | 15% a chill application | aligned within the sample's error |
| the charge spend | the first charge leaves at the Frostbolt's cast completion (18 ms after `SPELL_CAST_SUCCESS`), the second at Ice Lance's; the bolt in flight still crits as Frozen | spent as the cast resolves; the in-flight bolt keeps Shatter | aligned |
| Frostbite | 16 freezes or boss immunes over 95 chills (16.8%) | 5% a rank, 3/3 | aligned |
| Cone of Cold rank 1 | 108 to 121 non-crit | 97 plus 0.8 a level to 31, 9.7% wide, plus 0.129 x 88 | aligned |
| Frost Nova rank 2, Blizzard rank 2 ticks | 36 to 41; 46 to 55 (median 46) | 34 plus 0.5 a level, 5.7% wide, plus 0.029 x 88; 42 plus 0.2 a level plus 0.042 x 88 | aligned |

- **An observation, not a bug:** 11 of the 16 Frostbite procs land in the same instant as a Fingers of Frost grant. Independent 15% rolls would pair about 2 of 16. The server seems to roll the two Chill procs together; on a boss, which Frostbite cannot freeze, it changes nothing the sim measures. Recorded for a later Frost lane with a 60 log.
- Winter's Chill: no stack events (the mage has no points), as expected at that level.
- **Open, carried:** Ice Lance's coefficient at 60 (a rank 6 log) and Arcane Power's cost.

**Fire, Ximeya (2836, Deadmines, 9 Nov, about level 20; Fireball rank 4, Fire Blast rank 2, Flamestrike).** 172 Fireball damage events: direct hits 70 to 117 and the flat 3-a-tick burn (2 on a partial, 4 on a tick crit), both the client's rows; no Ignite, Heating Up or Scorch in the build, so nothing of the Fire kit above Fireball is in a public log yet. Our 8 Nov report's mages are the same picture at 12 to 28. **Open:** a Fire log with Ignite and Heating Up at 30 or above.

**Balance.** No public or own log has a Balance druid at 40 or above (2818's Never at 14 to 20 casts Wrath rank 3 at the client's 30 to 33 plus spell power; 2827/2836's Yingarna heals). **Open, carried:** Eclipse on a missed Wrath, Nature's Grace from the hit, Omen's rate in Moonkin Form.

## Before and after

**Engine.** Hurricane's tick table (E1) reaches no card and no single-target fight: every published DPS, weight and rank is unchanged at the new pin. Raid 60 on the new pin against the old (10,000 iterations, Alliance seed 1, the published enchants worn): Fire 817.2, Frost 674.6, Balance 729.8, within error of the same entries on e586233d7 (749.3, 618.0 and 654.2 bare; the gap is the enchants, below).

**The tool (T1) against the site.** spec-breakdown with the published enchants, seed 7, 3,000 iterations, Alliance raid 60, against the page's `set_dps`:

| spec | page (BiS engine 0e04b1284, 10 Nov) | breakdown, BiS engine, today's data | breakdown, pin 8dacbe217 |
|---|---|---|---|
| Fire | 820.3 | 820.3 | 815.4 (-0.6%: Arcane Concentration's split, hotfixes lane) |
| Frost | 685.3 | 675.7 (-1.4%) | 675.2 |
| Balance | 730.8 | not run | 730.8 |

- Fire and Balance reproduce the page to the decimal: the breakdown and the ranker now measure the same character.
- **Frost's 1.4% is data drift, not the engine:** today's data run (1c0ff79f9, 09:07 UTC) restated Earthen Silk Slippers (254013, the Frost raid feet) from 24 spell power, 18 healing and 7 Spirit to 6 spell power, 17 healing and 7 Spirit (items/mage.json; Robe of the Archmage's use text moved from "Restores 500 mana" to "375 to 625" with no stat change). The published Frost entry predates it; the next BiS run re-ranks the slot (ledger MR54).

**The Insect Swarm row (R1).** Bare band 50 at the levels the row covered, rank 4 over the published rank 3 (10,000 iterations): 54 +0.8/+0.7% Alliance, +0.7/+0.7% Horde (seeds 1/50001); 55 +0.7/+0.7%, +0.7/+0.7%; 56 +0.6/+0.8%, +0.7/+0.7%. 53 and 57 to 58 already cast rank 4 and gain nothing. Raid 60 and bare 60 are untouched (rank 5 from 60).

**Raid 60 at the new pin, both factions, the published enchants worn (10,000 iterations, seeds 1 / 50001):** Fire Alliance 817.2 / 818.0, Horde 803.7 / 803.6; Frost Alliance 674.6 / 675.5, Horde 669.4 / 669.4; Balance Alliance 729.8 / 730.0, Horde 708.3 / 707.6 (errors 0.4 to 0.5). These are the first raid-60 breakdowns on the site's own gear; the 2026-11-11 lanes' raid-60 figures before this landing were measured bare.

## Gates

Every Go test and sim ran on the Windows box through `tools/remote-test.sh`.

- **Fork (a writable copy of the pinned module, `go test -tags=with_db`):** `sim/druid` fails to build at c6b29db31 (`HurricaneTickDamage` undefined: the red), and `sim/druid` and `sim/druid/balance` pass at 8dacbe217 (and at the rebased sha after the paladin lane's master).
- **Site, red first:** `cmd/spec-breakdown` fails to build at 38dfb5068 (`bisSlot` has no `EnchantID`); `TestBalanceRankRowsFortyToFiftyNine` fails at f927fc40d on the six 54 to 56 cases.
- **Site, green at the landing:** `go vet ./...`; `go test` for `sim/request` (the ladder goldens of every spec, `TestRound30DamageRankRows`, `TestEmbeddedRankChoicesMatchCurated`, `TestLadderIsDeterministic`), `sim/rankchoices`, `sim/leveling`, `sim/specs` and the whole `sim/cmd/spec-breakdown` package (333 s).
- **Web:** `vitest run src/content/guides` (534 tests: prose guards, sections, builds).
- No APL, build, preset or golden moved, so `apl-check` and the goldens hold as origin/main left them; no regen was run.

## Guide and card

- **Balance (`balance.md`):** the rank list's Insect Swarm line stops at 38 (R1). wow-player: SHIP ("a clean bullet, the same shape as its Moonfire and Wrath neighbours"; the closing "Everywhere else, cast the newest rank" covers 39 on). The reviewer's optional tidy of the paragraph's "20s and 30s" lead-in against the Wrath rank 6 line at 46 and 47 predates this lane and is left.
- **Fire and Frost:** no prose or card changes; nothing the engine models moved.

## Measure at refresh (ledger MR54)

- **MR54, Balance:** bands 50 and 60, every published number at 53 to 58 (the Insect Swarm row), and the addon's rank rows (Data.lua) at the next regen.
- **Frost raid 60:** the feet slot after Earthen Silk Slippers lost 18 spell power in today's data (254013); the published set_dps is 1.4% above what the site's own data now sims.

## Sim changes in this landing

- Fork `lane/sim-casters`, rebased on master 3bdfab5b6 and fast-forwarded into master: c6b29db31 (test, red on the box), 8dacbe217 (fix). Site pin 8dacbe217 (4af50ac1b).
- Site: 38dfb5068 (red) and 10b49b36a (spec-breakdown enchants); f927fc40d (red) and 4c94cc933 (the Insect Swarm row and the guide line).
- No APL, build or preset moved; `apl-check` holds.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 1 to 6, read 11 Nov).
- Blizzard, Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092.
- data/builds/1.60.1.70790: spellconst, talents, hotfix-push.json.
- foreverlogs.gg reports 2815 to 2850; encounter timelines of 2839 (18) and 2836 (7).
- api.foreversixty.gg report gxrgvz4yzsi2 (fights 1, 4, 8, 11, 54, 56, 57, 59, 78).
