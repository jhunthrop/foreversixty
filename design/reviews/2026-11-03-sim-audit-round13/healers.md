# Sim audit round 13: healers

Lane: site wt/r13-healers and fork lane/r13-healers. Both start at int/r13 (site 93c5d2510, fork 076a5a517). The fork branch is fast-forwarded to lane/r13-core e88a33c97 (random regen-tick phase), as the coordinator asked, so the Discipline haste check runs on core's change. Specs: Holy and Discipline priest, Shadow priest, Holy paladin, Restoration shaman, Restoration druid, and the heal profile. Client 1.60.1.70753, read through round 12's row printer (design/reviews/2026-11-03-sim-audit-round12/healers/rows.py). Every measurement ran one at a time under `nice -n 15` with GOMAXPROCS=9.

The short version:
- **The Restoration druid's party Rejuvenation waited on the mana pace, and its tank Rejuvenation was rank 4.** It is the druid's cheapest healing. It is now rank 1 on the party, cast whenever someone is hurt, and rank 10 on the tank in a raid. Raid 60 +12.0%, bare 60 +40%. Every band below gains (+2 to +14%).
- **Holy's Prayer of Mending waited on the mana pace too.** Its own note and the guide say "every time it is ready". Raid 60 +1.4 to +1.6%, bare 60 +22 to +23%, bare 40 and 50 +10 to +11%.
- **Shadow now keeps Vampiric Embrace up**, as the guide says: raid 60 -3.9 to -4.0% DPS (a global every 30 s; the group healing is not counted).
- **The level-60 heal split from Warcraft Logs is not finished.** The fetcher and the rule are ready, but gcloud's token expired after 12 responses, so 60 keeps 0.659 (the coordinator's fallback).
- **Discipline's peak at exactly 0% haste is a core tie-break artifact, not a Discipline bug.** The explanation is in X3, sent to core.

## Coordinator items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **The level-60 tank and group split from Classic Era logs: not finished.**<br>- **The rule.** For each Molten Core, Onyxia and Blackwing Lair boss, take 10 logged kills at raid size 30 or more. They come from the S0 dps rankings on vanilla.warcraftlogs.com (pages 2, 5 and 10), at most one kill per report, ordered by a hash of report and fight.<br>- **A kill's tank share** is the damage taken by the players Warcraft Logs' summary lists as tanks for that fight, over the damage taken by every player. Pets are left out.<br>- **The band value** is the median over every sampled kill. Every boss has the same number of kills, so every boss weighs the same.<br>- **Trash is out:** the level-60 profile is a boss fight, unlike the sub-60 bands' trash pull.<br>- **What happened.** gcloud's token expired after 12 cached responses ("Reauthentication failed. cannot prompt during non-interactive execution"). I did not work around the reauth.<br>- **The cache holds one usable kill:** Onyxia, report ZCgbD2ydYN8rtkVw fight 2 (40 players, 82 s, Naxxramas-era gear). Its tank share is 0.312. One kill is not a sample.<br>Level 60 keeps 0.659. | Tools: [healers/wcl.py](healers/wcl.py) (cached, rate-limited; the key is read from an environment variable and never written) and [healers/raidsplit.py](healers/raidsplit.py). The cache sits in the lane's scratch folder, fetched 2026-11-03, and has no key in it. | OPEN (rerun `raidsplit.py` once gcloud is re-authenticated) | - | - |
| A2 | **Discipline's 1% peak at exactly 0% haste** (core lane's question). Measured on core e88a33c97, horde raid 60, 10,000 iterations: 0% 655.7 / 655.9, +0.01% 652.7 / 653.0, +0.5% 653.4 / 653.7 (seeds 1 / 50001).<br>No Discipline line tests a time with exact equality. Two same-instant ties do it; see X3. | scratch spec-breakdown with bonus spell haste; one-iteration combat logs | Not a Discipline bug (sent to core) | - | - |
| A3 | **The guides lane's proposed wording, "every rank from level 20 up carries the same healing power bonus", is wrong.** In Forever, heals learned below 20 keep their whole coefficient. Round 11 C1 showed it from beta logs, and `core/spell.go` levelPenalty returns 1 for healing. The three priest notes ("keeps its full healing power bonus") are right as written. | round 11 healers C1; engine | Aligned (no change) | - | - |
| A4 | **Shadow never cast Vampiric Embrace,** though the guide says "On the boss, keep it up: it costs a global every 30 seconds".<br>- **The client:** 15286 costs 40 mana, has a 30 s cooldown (SpellCooldowns 55127) and lasts 30 s, and heals the group for 20% of Shadow damage.<br>- **The fix:** the list now casts it whenever the target lacks it, with at least 10 s of the fight left, once the talent is taken (bands 40 and up).<br>- **The cost:** raid 60 DPS -4.00 / -3.93% alliance, -3.92 / -3.94% horde. Bare 40 / 50 / 60 -3.9 / -3.5 / -3.1%. Bare 30 has no talent and is unchanged.<br>The DPS metric does not count the group healing. | `TestShadowKeepsVampiricEmbraceUp` failed first, then passed. web/src/content/guides/priest/shadow.md:81 | CONFIRMED-BUG (rotation against the guide) | old (no Shadow list has cast it; guide since 7de4be717) | S 175d4ff96, E ec5a105ee |
| A5 | **Walk H2: Swiftmend reserved at druid bands 30 and 40 loses about 4%.** That is real. Swiftmend does not consume the HoT in Forever (round 12 D5). The reserved point displaces Gift of Nature or Improved Rejuvenation, which raise every HoT tick, and Swiftmend costs 20% of base mana on mana-bound bands. D2 makes HoT points worth more again. | walk.md H2; druid/swiftmend.go | Aligned (keep the walk) | - | - |
| A6 | **Walk X2:** between published bands, Discipline at 35 drops Divine Aegis, and Restoration shaman 46 to 49 drops Riptide. This touches the addon's point order and the strips only, not a published band. Riptide is cast, so the shaman gap should be closed by a reservation. | walk.md X2 | OPEN (next round) | R10/R11/R12 rows | - |

## Holy and Discipline priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Holy's Prayer of Mending waited on the mana pace.**<br>- **The bug:** the line required mana at 0.9 of the fight's remaining share. Prayer of Mending is Holy's cheapest healing (8.6 healed a mana, against Heal's 5.6). Its note and the guide say "every time it is ready".<br>- **The fix:** the gate is removed.<br>- **Measured** (10,000 iterations, seeds 1 / 50001): raid 60 alliance +1.44 / +1.52%, horde +1.42 / +1.53%. Bare 60 alliance +22.0 / +21.9%, horde +23.2 / +23.1%. Bare 40 +11.0 / +10.9% and bare 50 +10.5 / +10.5% (horde). | `TestHolyCastsPrayerOfMendingWheneverReady` failed first, then passed; guide holy.md:70 | CONFIRMED-BUG | old (the pacing pass 54986deb6, 2026-10-07) | S 0dcd484d5, E 1614f6c49 |
| P2 | **Every other mana-paced line, released one spell at a time** (raid and bare 60, horde, 3,000 iterations): Holy's Heal and Prayer of Healing; Discipline's Shield, Binding Heal, Flash Heal, Prayer of Healing, Greater Heal and Lesser Heal. None gains beyond error at both presets. The best is Discipline's Greater Heal: raid +0.6%, bare +0.2%. Those gates hold back costly heals, which is their job. | [healers/ungate.py](healers/ungate.py) | Aligned | - | - |
| P3 | **Holy's Renew rank on the tank at raid 60:** rank 3 698.8, rank 6 700.2, rank 7 699.5, rank 8 695.8, rank 9 687.0 (3,000 iterations). Within error, so rank 3 stays. | spec-breakdown -rotation | Aligned | - | - |
| P4 | **Fresh read of the talents the builds take, against 70753's rank text:** Twin Disciplines' masks, Mental Agility, Inner Focus (+25% non-periodic), Holy Specialization (1% a rank), Spiritual Guidance, Spiritual Healing, Improved Healing, Litany of Light (5% a rank of base cost, when the previous heal differs), Holy Binding (50% of Flash Heal, no crit or modifiers), Divine Aegis (5% a rank, capped at 20 a level), Inspiration, Meditation (17/33/50).<br>Spirit of Redemption has no Spirit bonus in Forever's text, and the engine gives none. | talents/priest.json; priest/healing_talents.go | Aligned | - | - |
| P5 | **Downranking.** The engine charges heals no level penalty, below 20 or above (round 11 C1; Era has no per-rank penalty above 20). No level-60 log can test a low rank yet. | core/spell.go levelPenalty | Aligned (beta logs; Era rule) | - | - |
| P6 | Round 12's open priest rows (Prayer of Mending travel P7 and variance P9; the Lightwell Renew crit; Inner Focus with Power Infusion) have no new evidence. | - | OPEN (carried) | - | - |

## Shadow priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| SH1 | Vampiric Embrace: see A4. | - | CONFIRMED-BUG | old | S 175d4ff96 |
| SH2 | **Shadowform's client effects:**<br>- +10% Shadow damage and -15% Physical damage taken.<br>- Aura 72, -50% Shadow cost.<br>- Aura 108 misc 15, +100% crit bonus, on mask 0x280a010 / 0x2 (Mind Blast, Shadow Word: Pain, Mind Flay, Devouring Plague; Shadow Word: Death).<br>The engine's crit-bonus mod covers the priest's Shadow class spells, the same set. | 15473 SpellEffect 692710-1239839; priest/talents.go:486 | Aligned | - | - |
| SH3 | **Improved Mind Flay** "each instance deals 20% more than its previous instance" (2/2): the engine compounds it (1.0, 1.2, 1.44). The effect is a dummy aura (1225139 e0, 20), so the text decides. One source. | priest/talents.go:563 | Aligned (single source) | - | - |
| SH4 | **Mind Flay's 12.5% crit** in the breakdown is not a lower crit rate. The hit count includes the channel's own landing (about 35 a fight), which cannot crit. The ticks crit at about 17%, like Shadow Word: Pain's. | spec-breakdown, raid 60 | Aligned | - | - |
| SH5 | **Round 12's open item, Mind Flay at band 20 ahead of Shadow Focus and Blackout.** Every build that reaches Mind Flay by 20 loses to the walk's (bare 20, horde, 3,000 iterations): walk 36.4; SF5 IMB3 ISWP2 MF 36.1; SF5 IMB5 MF 34.9; SF5 IMB4 ISWP1 MF 35.4. Band 20 is mana-bound and on the wand. | spec-breakdown -talents | Aligned | - | - |

## Holy paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Illumination's refunds add up:** 48.7 a fight, 112.8 mana each. That is exactly 50% of the base cost of each crit spell (Holy Light 8, both Flash of Light ranks, Holy Shock), counted by crits. | -heal breakdown | Aligned | - | - |
| H2 | **Holy Shock crits 51.5%** on party targets: Boundless Light's +30% off the Oath target, plus the base 25%. | talents/paladin.json | Aligned | - | - |
| H3 | **Releasing the mana pace** on Holy Light or Holy Shock loses at raid 60 (-1.0%, -0.7%) and ties or loses at bare 60. | ungate.py | Aligned | - | - |
| H4 | Round 12's H3 (Illumination's fourth effect) and H4 (Holy Power's crit size, Oath of Light) have no new evidence. | - | OPEN (carried) | - | - |

## Restoration shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Releasing Chain Heal's mana pace,** the shaman's one paced line: raid 60 -5.8%, bare 60 +1.9%. Mixed, so the pace stays. | ungate.py | Aligned | - | - |
| S2 | **Healing Stream crits about 13%:** item crit only, from round 11's beta logs (0 crits in 21,134 ticks without Thundering Strikes). | shaman/water_totems.go comment | Aligned | - | - |
| S3 | Chain Heal's jump order and the Purification / Healing Way stacking are still open. | - | OPEN (carried) | - | - |

## Restoration druid

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **At raid 60 the druid's mana lasted 526 s of a 300 s fight,** with Rejuvenation rank 4 on every target. Round 9 measured that pick as +0.9%, before later rounds raised the druid's mana income.<br>The tank's Rejuvenation is now rank 10 (9841): +2.3% on its own, with the party at rank 7 (S a7070f05d). The tank's rank-choice row keeps rank 4 from 28 to 49 and at bare 60, and rank 1 from 50 to 59. | `TestRestorationDruidRejuvenationRanksByBand` failed first. 10,000 iterations: +2.36 / +2.30% alliance, +2.38 / +2.33% horde | CONFIRMED-BUG | R9 dcd25dee3 (right then, overtaken) | S a7070f05d, E 268ec0362 |
| D2 | **The party's Rejuvenation waited on the mana pace,** though it is the druid's cheapest healing.<br>- **The bug:** ungated, rank 1 (774) beats every rank at every band. Raid 60 with the party ungated, by rank: rank 1 934.0, rank 2 932.1, rank 3 925.7, rank 4 918.9, rank 7 898.2, rank 10 865.4.<br>- **Why rank 1 wins:** heals learned below 20 keep their whole coefficient (round 11 C1), so rank 1 heals about 30 a mana.<br>- **The fix:** the party lines are now rank 1, with no mana condition. The rank-7 rows from D1 are gone.<br>- **Measured** against D1 (10,000 iterations, seeds 1 / 50001): raid 60 alliance +9.43 / +9.42%, horde +9.44 / +9.38%. Bare 60 alliance +39.7 / +39.7%, horde +39.8 / +39.8%. Bare 20 / 30 / 40 / 50 horde +3.1 / +11.9 / +2.0 / +13.6% (both seeds the same). | `TestRestorationDruidPartyRejuvenationIsNotPaced` failed first, then passed | CONFIRMED-BUG | old (the written rotation 35d37c4c0, 2026-10-07) | S 383d137f1, E a21e535ce |
| D3 | **Releasing Swiftmend's or Healing Touch's mana pace as well,** measured with the party ungated at rank 4: Swiftmend raid +0.4%, bare -0.2%. Healing Touch raid +1.7%, bare -8.9%. Both stay paced. After D2 the druid ends a raid fight with mana to spare (lasts 817 s), so Healing Touch's raid pace is worth a second look. | ungate.py, 3,000 iterations | Aligned (see Measure at refresh) | - | - |
| D4 | Bloom's crit, Rejuvenation's same-op stacking and Swiftmend with Nature's Splendor (round 12 D5, D7) are still open. Rank 1 now rests on round 11 C1's sub-20 evidence, which reaches level-6 ranks (Healing Wave rank 1) but not a level-4 one. | - | OPEN (carried) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The retune rule's ceiling.** After D2 the druid's raid set covers about 95% of incoming damage (933.7 of 978 a second). The rule holds the strongest raid set at about 80 to 85%. Holy priest went from 72% to 73%. The 1.85 scale was tuned before this. | heal-profile.json tuning.reason | OPEN (owner: retune at refresh) | - | - |
| X2 | **Where the mana-pace lines go wrong.** All three bugs (P1, D2, and D1's mana surplus) come from one pattern: a mana pace on the cheapest heal. The pace should hold back costly heals and let the cheapest one run. Every paced line of every healer was released in turn (P2, H3, S1, D3), and only these two gained. | ungate.py | - | - | - |
| X3 | **The 0% haste tie** (A2; sent to core).<br>- **First tie:** Penance's 12 s cooldown ends on the instant a global or a 2.5 s or 3 s cast ends. At 0% the list casts it there. With any haste the cast ends about 0.25 ms early, and a filler takes the slot. Penance's cooldown at 12.001 s gives 21.0 casts, the hasted count. That is about a quarter of the gap.<br>- **Second tie:** raid pulses push casts back by 1.0 or 0.8 s, which puts cast starts on the pulse's phase. A pulse then lands on the exact instant of a Penance tick, and the tick wins the tie. Hasted, the pulse falls just before about half the time and cuts the last tick (core/cast.go pushBackChannel). Penance pulses per cast: 2.80 at 0%, 2.77 hasted.<br>Weakened Soul at 15.001 s changes nothing. | combat logs; scratch cooldown tests | For core (tie-breaking) | - | - |

## Before/after

Raid 60, 10,000 iterations. Before is $R/before.tsv (int/r12 93c5d2510 / 076a5a517). After is wt/r13-healers 383d137f1 with lane/r13-healers a21e535ce, which includes core e88a33c97. Effective HPS, DPS for Shadow. Absolute values are horde, seed 1.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| Holy priest | +1.45% | +1.62% | +1.40% | +1.43% | 698.9 to 708.7 |
| Discipline priest | -0.05% | -0.11% | -0.03% | -0.03% | 655.9 to 655.7 |
| Shadow priest | -4.00% | -3.93% | -3.92% | -3.94% | 720.0 to 691.8 |
| Holy paladin | +0.05% | +0.03% | -0.08% | -0.06% | 645.1 to 644.6 |
| Restoration shaman | -0.06% | +0.04% | -0.02% | +0.01% | 840.2 to 840.0 |
| Restoration druid | +11.93% | +12.05% | +11.99% | +11.97% | 833.7 to 933.7 |

Discipline, Holy paladin and Restoration shaman move only by core's regen-phase change.

Below 60 (bare, horde, seeds 1 / 50001):
- Holy: 40 +11.0 / +10.9%, 50 +10.5 / +10.5%, 60 +23.2 / +23.1% (P1).
- Restoration druid against int/r13: 20 +3.1%, 30 +11.9%, 40 +2.0%, 50 +13.6%, 60 +39.8% (D1 moves no band).
- Shadow: 30 0.0%, 40 -3.9%, 50 -3.5%, 60 -3.1% (A4).

## Files to regenerate

- sim/request `TestRotationLadder` goldens for priest-holy, priest-shadow and druid-restoration (every band).
- The fork's druid-restoration and priest-holy ui copies are synced by hand. The integrator's `make apl-sync` rewrites them with the site's notes. The fork's druid copy still carried older notes.
- The addon's Data.lua and `addon-data`, which carry the rank-choice rows and rotations (nightly-owned).
- Every BiS file for the three specs, and dps-per-point for them, at refresh.

## Measure at refresh

- Every Restoration druid pick, raid and bands: the fight is no longer mana-bound at raid 60, so Spirit and mp5 lose weight.
- Holy priest's bare-band picks (P1).
- Shadow's picks with Vampiric Embrace cast (one fewer global every 30 s).
- The heal profile's tuning scale under the retune rule (X1).

## For other lanes

- **Core:** X3, the tie-breaking that favours exact 0% haste.
- **Guides:** restoration.md:83 ("Cast Rejuvenation rank 4 on everyone") now contradicts the rotation. Suggested: "Cast Rejuvenation rank 1 on a hurt party member: your healing power adds the same to every rank and Forever keeps it whole on low ranks, so it is your cheapest healing by far. On the tank use the top rank in a raid and rank 4 while leveling." The Shadow guide's Vampiric Embrace paragraph already matches the rotation.
- **Ranks:** the healer row change is druid-restoration only (authored 9841: 28-49 rank 4, 50-59 rank 1, bare 60 rank 4), with no race key.
- **wow-player review:** D2 (rank 1 Rejuvenation on the party) and A4.

## Still open

- **Finish the level-60 split once gcloud auth is restored:** `WCL_KEY=... python3 -I raidsplit.py <dir> 10`, then update heal-profile.json's band 60, with tests first. Today the reference band refuses a logged split (`validateSplit`), so that refusal has to change with it.
- A6 (walk X2), D4, H4, S3, P6, X1, X3.
