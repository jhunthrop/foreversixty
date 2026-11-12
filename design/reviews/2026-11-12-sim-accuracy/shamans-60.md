# Sim accuracy 2026-11-12: Enhancement and Elemental at 50 and 60

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main aafe382b5, fork from master 0fb4dcebb (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push (wago.tools, `useHotfixes=1`). Scope: Enhancement and Elemental at 50 and 60, bare and the raid 60 preset, against real fights recorded after Blizzard's 10 November hotfix; failing that, the engine against the client for every row the hotfix touched, the raid talent and rotation searches on the fixed `rotation-search`, Enhancement's imbue rows and the kit at 50 and 60.

**Totals: no post-hotfix log exists anywhere we can read, so this is the engine-against-client pass. 1 engine bug, old, fixed: Chain Lightning restored its bounce multiplier after the hits landed, so a Stormstrike charge the first hit had spent came back on the spell for good and compounded (the rotation search valued a Chain Lightning line at 1e40 DPS for Enhancement; no published number moves, since no single-target card casts it). Every row the 10 November hotfix touched, every level 50 and 60 spell, totem, imbue and talent the two cards use, and the three core changes since the last shaman pin match the hotfixed 70888 client or Blizzard's words; Elemental's five Lightning Bolt rows at 50 to 60 and Enhancement's imbue at 50 and 60 hold. The raid builds and all four rotations hold; the two level 60 leveling builds move (Elemental +4.7%, Enhancement +5.0% at bare 60, both clean under the adoption rule; MR87, MR88). Engine pin 0fb4dcebb to 4c595ee6b (the fix rebuilt on the fork master the warriors lane moved to 837cca30a).**

## Logs: nothing recorded after the hotfix carries the hotfixed numbers

1. **foreverlogs.gg.** The public list is still the 25 reports of the last shaman pass (2815 to 2853, the newest uploaded 12 November 01:37 UTC and started 11 November 23:51 UTC). The deploy sits between 2847 (started 10 November 19:45 UTC) and 2849 (11 November 06:45 UTC): every Stormstrike user in 2847 and 2848 (Josh, Torben, Voss, Tapi, 94 to 259 Stormstrikes each) still has a 6 s cooldown (the tenth percentile of the gaps between casts is 5.99 to 6.00 s). The three reports since (2849, 2850, 2853: a Dalaran city log and two dungeon nights, 54 shamans between them) hold no Stormstrike, no Maelstrom Weapon, no Lightning Overload and no Windfury Weapon in a fight: the one busy shaman is Rhambo in 2853, a level-23 Elemental (14 talent points, no Lightning Overload yet) with 854 Lightning Bolt cast events; the rest are travel buffs. Rhambo's bolts cannot show the overload rows, which start at 31.
2. **Our API.** The public recent feed (`GET /v1/reports/recent`) is empty; our 8 November advanced report (gxrgvz4yzsi2) predates the hotfix and was read by the last pass. The reports table itself (private uploads since 10 November) is reachable only through the Cloud SQL proxy, and the laptop's gcloud login has expired (`gcloud auth login` is interactive), so no private report could be listed. Needs the owner.
3. **Warcraft Logs** has no Forever zone (checked again: the vanilla and retail hosts refuse automated reads, and the two earlier lanes' finding stands).

So every 10 November number (Stormstrike 100% on 8 s, Mental Dexterity 33/66/100, Maelstrom Weapon 20% a stack and 6 to 10 a minute, Windfury Weapon under Windfury Totem) still rests on Blizzard's post alone. Stormstrike's 17364 in the hotfixed 70888 push still reads effect 2 = 150 and RecoveryTime 6000, Mental Dexterity's 415140 still 100 at 3/3 by curve, and 408498 still carries a 20/50/5 triple with no rate, so none can move to a client row yet (MR46 stands).

## Engine against the client at 50 and 60

Every row below was read from the hotfixed 1.60.1.70888 tables (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellLevels, SpellAuraOptions, one spell a request) and from the site's vendored 70790 `spellconst/shaman.json`; the fork at the pin on the right.

| row | client | engine | verdict |
|---|---|---|---|
| Stormstrike 17364 | effect 121 (normalized weapon damage), effect 2 still 150 and RecoveryTime 6000 in the push; 8% base mana; debuff aura 271 +20 on mask 0x100003, ProcCharges 1, ProcTypeMask 0x20000 (a harmful spell taken from the caster spends it), DurationIndex 29 (12 s) | 100% and 8 s by hand (Blizzard's 10 November line), 8% base mana, +20% on Lightning Bolt, Chain Lightning and Earth Shock for one charge, 12 s | aligned (hand values until a client row states them) |
| Mental Dexterity 415140 | aura 598 100 at 3/3 by curve 85272 (50/100/150 unchanged in the push) | 33/66/100 by hand | aligned with Blizzard's line |
| Maelstrom Weapon 408498 and its stack 408505 | the talent's proc aura has ProcTypeMask 0x10014 (melee auto attacks, melee abilities and harmful magic spells), ProcChance 100 and no rate; the stack 408505 is aura 108 -20 on cast time (misc 10) and -20 on cost (misc 14) for class mask 1 (Lightning Bolt), CumulativeAura 5, DurationIndex 9 (30 s) | 20% a stack at every rank, five stacks, 30 s; 6 to 10 a minute by rank as a chance on the main hand's speed, rolled on auto attacks and melee abilities, and on Lightning Bolt only through Totem of the Storm | aligned where the client speaks (stack size, duration, cost and cast cut, the ability bit); the rate stays server side, see Open |
| Windfury Weapon rank 3 and 4 (10486, 16362; buffs 10484, 16361) | 140 and 165 mana; buff aura 99 249 + 8.3 a level to 58 and 333 + 12.5 to 68, effect 19 two attacks, DurationIndex 65 (1.5 s), ProcTypeMask 4 | ranks learned at 50 and 60, 249 + 8.3 and 333 + 12.5 by level, two strikes with the bonus attack power, 20% a main-hand hit with 1.5 s between procs | aligned |
| Flametongue Weapon rank 5 and 6 procs (16343, 16344) | 1876 + 62 a level (46 to 54), 2498 + 78 (56 to 64), effect 3 | the last pass's by-level rows | aligned |
| Lightning Bolt rank 9 and 10 (15207, 15208) | 295 + 1.6 (50 to 55) and 353 + 2.1 (56 to 61), variance 0.114 and 0.108, 0.714, 190 and 220 mana, CastingTimeIndex 19 (2.5 s) | the same rows | aligned |
| Lightning Overload rank 9 and 10 (408475, 408477) and Chain Lightning's rank 4 (408484) | 147.5 + 0.8, 176.5 + 1.05 and 255 + 1.65 at 0.357 in 70888 | the same by hand (the vendored 70790 file still holds 89 + 0.5 and 98 + 0.6) | aligned (MR46) |
| Chain Lightning rank 4 (10605) | 510 + 3.3 at 0.714, 495 mana, category cooldown 6 s, 2.5 s cast | the same | aligned |
| Earth Shock rank 6 and 7 (10413, 10414) | 383 + 1.6 and 531 + 3.4 at 0.386, 345 and 450 mana, category 6 s | the same | aligned |
| Flame Shock rank 6 (29228) | 292 + 3.8 at 0.214, a 80 tick every 3 s at 0.1, 410 mana, 12 s (21 with Call of Flame) | the same | aligned |
| Searing Totem rank 5 and 6 and its Attack (10437, 10438, 10436) | 145 and 170 mana, 50 and 55 s, the bolt 47 at 0.017 every 2.2 s | the same | aligned |
| Strength of Earth rank 4 and 5 (10441, 25362), Grace of Air rank 2 and 3 (10626, 25360), Mana Spring rank 4 (10494), Windfury Totem rank 2 and 3 (10608, 10610) | 42 and 53 Strength, 77 and 89 Agility, 10 mana every 2 s, 179 and 246 attack power with one extra attack | `core/buff_ranks.go` and `core/buffs.go` carry the same amounts and levels (52, 60, 56, 60, 56, 42, 52) | aligned |
| Flurry (16256, 16257) | the buff has ProcCharges 3, ProcTypeMask 4 (auto attacks) and ProcCategoryRecovery 500, DurationIndex 8 | three stacks, a white hit spends one with 500 ms between spends, 15 s | aligned (the 500 ms is in the client) |
| Improved Stormstrike 1238931 | aura 134 50 for DurationIndex 8 | 50% regeneration while casting for 15 s | aligned |
| the three core changes since the last shaman pin | `IsExtraAttack` (paladins), the relic's ranged skill (paladins), `NewClientEquipModItemEffect` and Water Shield's charges (Restoration relics) | none reaches Enhancement or Elemental: the raid 60 and bare 60 and 50 breakdowns at the pin reproduce the published `set_dps` to the decimal (857.5 against 856.8 Horde raid, 873.4 against 874.4 Alliance, 313.9 bare 60, 212.3 bare 50; Elemental 749.6 and 748.0 raid, 287.9 bare 60, 166.5 bare 50) | nothing to change |

The kit at 50 and 60 (`sim/leveling/kit.go`): one imbue on the main hand only, Windfury Weapon from 30, never a second weapon (mail from 40 is the ranker's armor rule, not the engine's). The raid entry keeps Windfury Weapon over the preset's sharpening stone (`sim/request/preset.go` Layer) and receives the raid's Windfury Totem on top, as the breakdown shows (31.6 Windfury Weapon strikes and 19.4 totem extra attacks an iteration at raid 60). The one stale sentence was in `data/curated/presets.json`'s raid note ("Only a shaman whose main hand carries Windfury Weapon does not receive the totem"); it now states the hotfix rule.

## Elemental's Lightning Bolt rows at 50 to 60 (MR43)

`spec-breakdown -level N -rank 15208=R`, bare band 50 (band 60 for the last row), both factions' guide races (dwarf, orc), seeds 1 and 50001, 10,000 iterations, errors 0.2 to 0.3; the overload rows are the hotfixed half-bolt ones. Every row holds on the new overload: the newest rank costs far more than the fight can pay (the dry seconds are in brackets).

| level | row | row's rank | newest rank | row over newest |
|---|---|---|---|---|
| 50 | rank 6 (45 to 50) | 166.4 to 166.5 (62 to 64 s dry) | rank 9: 153.9 to 154.5 (103 to 104 s dry) | +7.8 to +8.1% |
| 51 | rank 7 (51 to 56) | 174.0 to 174.5 (71 to 73 s) | rank 9: 163.2 to 163.7 (98 to 99 s) | +6.6% (rank 6 ties rank 7 at 51: -0.1 to +0.2%) |
| 56 | rank 7 | 206.0 to 206.2 (52 to 53 s) | rank 10: 188.3 to 188.5 (95 to 96 s); rank 9: 195.1 to 195.6 | +9.4%; +5.4% over rank 9 |
| 57 | rank 8 (57 to 59) | 204.7 to 205.0 (67 to 68 s) | rank 10: 191.0 to 191.4 (94 s); rank 9: 197.4 to 198.0 | +7.1%; +3.6% over rank 9 |
| 60 bare | rank 8 | 285.7 to 287.1 (40 to 42 s) | rank 10: 272.5 to 273.2 (70 to 71 s); rank 9: 277.4 to 278.2 | +4.8 to +5.2%; +3.0 to +3.2% over rank 9 |

The guide's rank list and the card's filler note stand.

## Enhancement's imbue at 50 and 60

`spec-breakdown -consumes` with the one imbue swapped and the raid's other consumables kept (`jobs.sh` part B, `jobs2.sh`), 10,000 iterations, seed 1 (raid 60: seeds 1 and 50001, both factions). Windfury Weapon is the imbue at every band the kit carries it:

| band | Windfury Weapon | Rockbiter Weapon | Flametongue Weapon | Windfury over the next |
|---|---|---|---|---|
| bare 50 (Horde) | 212.3 | 199.9 | 203.1 | +4.5% over Flametongue, +6.2% over Rockbiter |
| bare 60 (Horde) | 313.9 | 301.4 | 301.4 | +4.1% over both |
| raid 60 Alliance | 873.4 and 874.0 | 773.5 and 772.8 | 719.3 and 718.2 | +12.9% over Rockbiter |
| raid 60 Horde | 857.5 and 856.3 | 760.6 and 760.6 | 705.7 and 704.6 | +12.7% over Rockbiter |

Under the raid's Windfury Totem the weapon's own procs add 31.6 strikes an iteration on top of the totem's 19.4 extra attacks (the 10 November rule), so the gap over Rockbiter is wider than the +10.8% the hotfix lane measured on the old pin's gear. The kit's rule (Windfury from 30) and the guide's imbue lines stand.

## Chain Lightning compounded a spent Stormstrike charge (fixed)

The two Enhancement rotation searches (`jobs2.sh`, `rotation/shaman-enhancement-raid.md` and `rotation/shaman-enhancement.md`) reported a Chain Lightning line worth 1e40 DPS. Reproduced at the pin with the card plus one `castSpell 10605` line before Earth Shock (`enhancement-cl.json`, bare 60 Horde, 200 iterations): Chain Lightning's average hit is 1.27e9 and grows every cast. The cause is in `chain_lightning.go`'s bounce loop: it scales `spell.DamageMultiplier` down a bounce at a time, deals the hits, then puts the entering multiplier back. Dealing the first hit spends the Stormstrike charge, whose dynamic mod (`stormstrike.go`, +20% on Lightning Bolt, Chain Lightning and Earth Shock, one charge) divides that same multiplier by 1.2 as it leaves, so the restore reinstated the spent charge on the spell for good and every Stormstrike then Chain Lightning pair compounded 1.2. Lightning Bolt and Earth Shock never save and restore around their hit, so they were never affected.

- **Fix (fork `lane/sim-shamans60`):** af3ba15fb the test, red (`TestChainLightningKeepsNoSpentStormstrikeCharge`: three Stormstrike then Chain Lightning rounds leave the multiplier at its entering value and the next Chain Lightning at plain damage; Stormstrike is landed until the charge is up, since its own roll can miss); 4c595ee6b the fix: the multiplier is restored once every bounce is rolled, before the hits are dealt, so a mod that leaves on the hit stays gone. On the box with the fork's item database (`-tags with_db`, three runs each) the test is red at af3ba15fb and green at 4c595ee6b; `TestEnhancement.results` and `TestElemental.results` are byte-identical on the fixed engine (the fork's suites cast no Chain Lightning on a Stormstrike target), so no fork golden moves. The fork master had moved to 837cca30a (the warriors lane) while this lane ran; the two commits sit on it and master fast-forwards to 4c595ee6b.
- **Published numbers:** none move. No Enhancement card casts Chain Lightning on one target, and the Elemental card (which casts it on two or more) has no Stormstrike; the raid 60, bare 60 and bare 50 breakdowns are identical on both engines. The guide's pack advice for Enhancement (Chain Lightning while your mana lasts) was never simmed, so no prose changes.

## Talent and rotation searches at 60 (on the fixed `rotation-search`)

`talent-search` and `rotation-search` on the pin, raid 60 on the raid build and bare 60 on the leveling build, Alliance BiS gear (dwarf), seed 7, the finalists at the ranker's weights level (`jobs2.sh`; the reports are in `jobs2.txt`). The keep lists hold what the cards cast or lean on (Stormstrike, Maelstrom Weapon, Rage of the Farseer, Elemental Focus, Improved Ghost Wolf, Tidal Focus, Shamanistic Focus; Lightning Mastery, Call of Thunder, Elemental Focus, Trial by Fire, Elemental Warding, Call of Flame, Tidal Focus, Improved Reincarnation).

| search | guide | best | verdict |
|---|---|---|---|
| Enhancement raid talents | 821.5 ± 2.2 | deep Elemental (Mindfulness 3 to Reverberation 2 and Call of Flame 1), 827.7 ± 2.2, +0.8% | under the 1% bar: the raid build holds |
| Enhancement raid rotation | the card | the card (the Chain Lightning blow-up aside; Flame Shock inserted +3.1 ± 3.1, within error) | holds |
| Enhancement bare 60 talents (leveling build) | 315.4 ± 0.8 | deep Enhancement: Convection 5, Elemental Devastation 3, Improved Stormstrike 2, Mindfulness 3 for Concussion 2 and Call of Flame 3; 331.3 ± 0.8, **+5.0%**, keeps every unmodeled guide talent (Improved Ghost Wolf) | **adopted** (below) |
| Enhancement bare 60 rotation | the card | the card, Chain Lightning aside; Flame Shock inserted at the top +27.8 ± 1.2 | held: Flame Shock at the top of a bare rotation competes with Earth Shock on the shared shock cooldown and the search put it above Stormstrike; a Flame Shock line for the leveling card is a rotation question for a later lane, with the new build worn (the raid search calls it no effect) |
| Elemental raid talents | 674.7 ± 1.0 | Improved Reincarnation 2 to Shamanistic Focus and Totemic Focus, +0.9% | drops an unmodeled guide talent; counting only clean variants the guide is best: holds |
| Elemental raid rotation | the card | the card (+0.0 ± 1.5) | holds |
| Elemental bare 60 talents (leveling build) | 287.8 ± 0.9 | Ancestral Knowledge 5 and Shamanistic Focus 1 for Totemic Focus 5 and the third point of Mindfulness; 301.5 ± 1.0, **+4.7%**, keeps every unmodeled guide talent (Elemental Warding, Improved Reincarnation) | **adopted** (below) |
| Elemental bare 60 rotation | the card | the card (+0.0 ± 1.3) | holds |

### The two level 60 leveling builds

Both winners keep every unmodeled talent, hold the spec tree, gain beyond error by more than 1% and drop nothing the card casts (Totemic Focus and Call of Flame are passives), so the talent adoption rule takes them:

- **Elemental:** `5500331220001351/05500001/052002` (was `.../05/053052`). After Tidal Focus and Thundering Strikes: Ancestral Knowledge 5, Shamanistic Focus 1, Mindfulness 2, Improved Reincarnation 2; Totemic Focus is gone. The band 50 entry moves too: Mindfulness 2 at 46 and 47 beats Thundering Strikes' last two points (the order table below), so bands 50 and 60 bare re-rank at the refresh (MR87).
- **Enhancement:** `503/055030310051202251/053` (was `323003/.../052`). Convection 5, Elemental Devastation 3, Improved Stormstrike 2, Mindfulness 3; Concussion and Call of Flame are gone. The band 50 entry (Convection 3, Tidal Focus 5 on the 33 Enhancement points) moves too, since Improved Stormstrike at 49 and 50 beats Convection's second and third points by 5.0% (the order table below); bands 50 and 60 bare re-rank at the refresh (MR88).

### The order of the last points (`orders.sh`, on the fixed engine)

Bare band-50 gear at the level, both factions' guide races, seeds 1 and 50001, 10,000 iterations (errors 0.2); `data/curated/ladder-first.json` carries the orders and `make ladder-first-sync` the embedded copy.

| level | candidates (points past the band-50 entry) | DPS | pick |
|---|---|---|---|
| Enhancement 50 | Convection 3 (the published entry) against Convection 1 + Improved Stormstrike 2 | 211.3 to 212.3 (2.5 to 6.8 s dry) against 222.1 to 222.5 (never dry) | **Improved Stormstrike at 49 and 50, +5.0%**: the band 50 entry moves (MR88) |
| Enhancement 53 | Convection 4 + Improved Stormstrike 2 against Convection 5 + Elemental Devastation 1 (no Improved Stormstrike) | 233.4 to 234.1 against 220.3 to 221.4 | Improved Stormstrike first, Convection next |
| Enhancement 57 | Convection 5, Elemental Devastation 1, Mindfulness 2 against Convection 5, Elemental Devastation 3 (both with Improved Stormstrike 2) | 258.5 to 259.3 against 249.2 to 249.9 | Mindfulness before Elemental Devastation's second and third points, +3.7% |
| Elemental 47 | Tidal Focus 5 + Thundering Strikes 2 against Tidal Focus 5 + Mindfulness 2 (band-40 gear) | 118.1 to 118.6 (83 s dry) against 130.3 to 130.7 (70 s) | **Mindfulness at 46 and 47, +10%** |
| Elemental 50 | Thundering Strikes 5 (the published entry) against Thundering Strikes 3 + Mindfulness 2 | 166.4 to 166.5 (62 to 64 s dry) against 185.6 to 185.9 (44 to 46 s) | +11.6%: the band 50 entry moves (MR87) |
| Elemental 55 | Ancestral Knowledge 5 against Ancestral Knowledge 3 + Mindfulness 2 | 175.7 to 176.2 (75 to 77 s dry) against 194.0 to 194.7 (62 to 64 s) | Mindfulness before Ancestral Knowledge, +10% |
| Elemental 58 | Ancestral Knowledge 5 + Mindfulness 2, then Shamanistic Focus against Improved Reincarnation 1 against Mindfulness 3 | 217.6 to 218.3 against 198.4 to 198.9 against 210.5 to 210.8 | Shamanistic Focus at 58, Improved Reincarnation last |
| Enhancement 55 | Convection 5 + Mindfulness 1 against Convection 3 + Mindfulness 3 (both with Improved Stormstrike 2) | 243.7 to 244.5 against 250.0 to 251.7 | Mindfulness before Convection's last points, +2.9% |
| Enhancement 57 | Mindfulness 3 + Convection 4 against Convection 5 + Elemental Devastation 1 + Mindfulness 2 | 258.8 to 260.6 against 258.5 to 259.3 | a tie within noise; Mindfulness first, Convection to 5, then Elemental Devastation (which needs five Elemental points) |

Orders landed: Enhancement 49 and 50 Improved Stormstrike, 51 to 53 Mindfulness, 54 to 57 Convection, 58 to 60 Elemental Devastation; Elemental 46 and 47 Mindfulness, 48 to 52 Thundering Strikes, 53 to 57 Ancestral Knowledge, 58 Shamanistic Focus, 59 and 60 Improved Reincarnation. `sim/leveling/ladder_first_walks_dps_test.go` carries the new rungs and the two ladder goldens are regenerated on the box. The guides' leveling sections and talent paragraphs say the same; the Enhancement guide gains an Improved Stormstrike bullet.

## Held and open

- **Maelstrom Weapon's realized rate.** The engine rolls the rank's 6 to 10 a minute as a chance on the main hand's speed (10 x 3.8 / 60 = 63% at raid 60) against every melee hit, so with Stormstrike, the two Windfury Weapon strikes and the totem's extra attacks the raid shaman sees 76 stacks an iteration, about 25 a minute, not 10. The client's proc mask (0x10014) says abilities do trigger it, so the per-hit reading is the natural one, and Blizzard's "was 20 times per minute regardless of rank" matches what the old 50%-a-hit engine gave at raid 60 (about 21 a minute); but whether "6 to 10 times per minute" is the per-swing chance or a cap only a post-hotfix Enhancement log with a known weapon speed can tell. Open; the first level 60 Enhancement log settles it, and the glancing question on Windfury Weapon's strikes with it.
- **Flametongue Totem over Windfury Totem** when both stand, and the Rockbiter bleed: held from the last pass, nothing new.
- **Private reports since 10 November** in our own API: unreadable from this laptop until the owner runs `gcloud auth login` (the proxy's token has expired). The first one with a level 60 Enhancement or Elemental shaman and advanced logging on closes every open row above.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 5, 6 and 8; nothing after post 8 as of 12 November 19:00 UTC).
- Blizzard, Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092.
- wago.tools DB2 CSVs for 1.60.1.70888 with its hotfix push (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellLevels, SpellAuraOptions), filtered one spell a request.
- foreverlogs.gg reports 2845 to 2853 (rosters, every shaman's damage, cast, buff, avoidance and resource events; scripts under `shamans-60/`).
- foreversixty API `GET /v1/reports/recent`.
