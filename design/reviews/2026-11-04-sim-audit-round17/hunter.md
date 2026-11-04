# Round 17, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755 with the late re-layer. Site wt/r17-hunter from int/r17 d7fc51fe7; fork lane/r17-hunter from int/r17 f637ec32c.
Method: spec-breakdown through `$R/sim.sh`, guide builds (`$R/rt`), `-bis-dir $R/bis-frozen`. 3000 iterations to explore; 10000 at seeds 1 and 50001, both factions, for every reported number.

**Result: 2 CONFIRMED-BUGs, both fixed** (1 engine, 1 site with guide text). Summon Hawk stays OPEN: no level-60 log exists yet.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Summon Hawk at 60: still not sizable.** foreverlogs.gg has no report with encounters after launch: 2814 ("Test", 2026-11-04 12:24 UTC) has 0 encounters, 2815 and up do not exist. api.foreversixty.gg/v1/reports/recent lists one report, from 2026-10-20. The web has no measured level-60 swing either. Icy Veins' BM guide (updated 2026-11-03 "for launch") says only that the hawk "scales with your pet's damage", with no numbers (single source). Two extrapolations now both land above the owner's 20-damage threshold: round 16's straight-line fit (about 23 a swing before armor) and round 16's measured hawk/pet ratio (0.24 to 0.27 of the pet's non-crit swing) applied to the level-60 cat. Both are extrapolations, so under the adoption rule nothing changes. | foreverlogs `reports/2813..2819`; `/v1/reports/recent` (2026-11-04); https://www.icy-veins.com/wow-forever/beast-mastery-hunter-ranged-dps-pve-guide (read 2026-11-04) | OPEN. Owner's two-minute check stands (round 16 hunter.md H2): ten hawk swings on a level 60+ target. It now looks likely to land over 20. | - | - |
| H2 | **The pet buff block was too wide: the pet never took its own hunter's Heart of the Lion.** Forever keeps group buffs off pets by flagging each one "can only target players" (SpellMisc Attributes_3 0x100). Round 16 listed them, and I added Windfury Totem (8515/10609/10612, flagged). Heart of the Lion's area buff 409583 is **not** flagged (Attributes_3 0 in 70755 and Era). It is SPELL_AURA 137/99/124 on ImplicitTarget 31 (every nearby ally), and the tooltip says "for all nearby allies". Unflagged area auras do reach pets in Forever: the beta timelines show Trueshot Aura (1299346, unflagged) applied to the caster's own pet (Kaiah to her Cat, plus six other hunters' pets). They also show Retribution Aura and Blood Pact landing on pets. In Season of Discovery, the same spell id buffed the pet (EU forum t/489837, 2024-01-13: "Heart of the Lion give ... and of course the Pet"). The engine's `RefusePlayerBuffs` returned before any buff, so the pet missed +10% stats and 40 + 4 a level attack power (280 at 60). Fix: `core.Pet.TakeHeartOfTheLion()`, which only the hunter pet calls. The pet now takes Heart of the Lion and nothing else. Trueshot Aura stays off it in effect, because Forever's rows are ranged attack power only (aura 124; Era also had aura 99). Test first: `TestHunterPetReceivesOnlyHeartOfTheLionFromPlayerBuffs` (red: 136/100 Str/Agi, 252 AP; green: 149.6/110, 559.2 AP plus the inherited share). | wago SpellMisc and SpellEffect 1.60.1.70755 and 1.15.9 (r16 `c755/`, `era/`); spelltips 409580; foreverlogs 2695 to 2764 `buff_applied` with target_type pet (`full/`, `fl/`); https://eu.forums.blizzard.com/en/wow/t/pet-less-hunterlone-wolf-rune-thoughts/489837 | CONFIRMED-BUG. **Single caveat:** no Forever log shows 409583 itself on a pet, because no beta hunter in the 48 encounters had Heart of the Lion up. In-game check: with Heart of the Lion on, hover the pet's buff bar. | old dde2af473 (round 1 theorycraft fix); round 16 H1 looked and missed it | E c6f2e362e |
| H3 | **Aspect of the Hawk: holds.** Rank by level 10/18/28/38/48/58 (trainables), 20/35/50/70/90/110 ranged AP (push), rank 7 25296 a pending book. Strongest-learned rank logic (`aspects.go getMaxHawkRank`) is right. The Deadly Aspects proc is on ranged autos only (round 16 H3). Physical school, so round 16's MaxLevel penalty (18815822c) cannot touch it. Raid 60 BM: 1 cast, 99.9% uptime, 9.45 Deadly Aspects procs. | trainables/hunter.json; `hotfix_push_70753_test.go`; spec-breakdown | Aligned | - | - |
| H4 | Round 16's shared changes against hunters: the MaxLevel-keyed penalty (E 18815822c) reaches only non-physical casts with a coefficient. Every hunter magic spell has BonusCoefficient 0 (Arcane Shot, Serpent Sting, Volley, Immolation and Explosive Trap), so no hunter number moves. A cast Kings no longer stacks with Heart of the Lion (E 913221285). Neither reaches the hunter pet. | `sim/core/spell.go levelPenalty`; hunter spell files | Aligned | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Aimed Shot rank 1 from 28, level by level.** The 40-49 row held at every level (+2.8 to +4.2%), but its neighbours were missing. Since round 13 put Aimed Shot on Multi-Shot's cooldown it is a rotation shot, and the bare bands run dry for 105 to 122 s. Rank 1 (75 mana, +40) beats the highest rank learned at every level from 28: 28 +2.0 to +2.3%, 33 +1.7 to +1.8%, 36/39 +3.0%, 50 +0.2 to +0.4%, 53 +0.1 to +0.3%, 56 +0.1 to +0.3%, 59 +0.6%. Every level was beyond error at seeds 1 and 50001, Horde and Alliance. The row is now 28-59 rank 1, and bare 60 is unchanged. Downranking Serpent Sting or Arcane Shot (rank 3) loses at every band. | spec-breakdown `-level` sweep (`ranks.sh`, `floors.sh` in `$R/hunter/`), 10000 iterations; rank 3 Serpent Sting/Arcane Shot probes at 3000 | CONFIRMED-BUG | old (rows last judged in round 13 e7a45f12e, before the -level sweep) | S 7e395c864 |
| B2 | Raid 60 after H2: cat white 182 DPS (was 131), 23% of damage. Everything else is as in round 16 B1 (Multi-Shot 20.6, Arcane 27.0, Aimed 7.6, Serpent 12.7 casts; 0 s out of mana). Builds re-decoded: every point casts or gates. | spec-breakdown | Aligned | - | - |
| B3 | Sniper Shot's cast scaling with ranged haste (guide-accuracy rounds 11 to 13 lead): 1310786/1310687 carry the exact SpellMisc attribute words of Aimed Shot (Attributes_0 4259858, _2 131072, _4 402653184), so they follow the same ranged-haste rule. The engine treats both alike. | wago SpellMisc 70755 | Aligned (closes the carried LIKELY) | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Aimed Shot rank 1 at 28-49 (row was missing).** Bare bands run dry for 33 to 97 s. Rank 1 beats the default at 28 (+1.7 to +2.1%), 30, 33 (+1.2 to +1.5%), 36 (+3.6 to +3.9%), 39, 40 (+2.2%), 43, 46 (+3.7 to +3.8%) and 49 (+3.5 to +3.7%). At 50 it ties (0.0 to +0.1, inside error at Horde seed 50001), so the row ends at 49. Bare 60 is unchanged. | as B1 | CONFIRMED-BUG | old | S 7e395c864 |
| M2 | **Lone Wolf still beats the pet in a raid after H2.** On the raid build with the Lone Wolf point moved to Rapid Killing (pet out), DPS is 787.7, against 801.1 with Lone Wolf (-1.7%). The leveling build at raid is 765.8. The guide's raid build holds. | spec-breakdown 3000 it, seed 1, Horde | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Lone Wolf against the pet after H2: raid 60 861.0 (Lone Wolf) against 839.9 with the pet (-2.5%), so the guide holds. **Bare 60 flips:** the pet build reads 375.4 against 368.3 with Lone Wolf (+1.9%, seed 1). Under the owner's one-build-at-60 ruling nothing changes. Recorded for the owner. | spec-breakdown 3000 it | Aligned (owner ruling) | - | - |
| S2 | Rank probes at 30-56: Raptor Strike rank 3, Immolation Trap rank 1 and Lacerate rank 1 lose at every level (-3 to -11%), and the top ranks hold. Mongoose Bite: no casts at bare 30, because the boss never swings at the hunter there and the build has no Expose Prey yet. The window aura (5302 by dodge, 1310726 by Expose Prey) matches the client. Counterattack only fires on the hunter's parry, which is right. | `ranks.sh` with sb-rs/-it/-lc; `mongoose_bite.go` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The "pre-existing fork test failures" (round 16 X2, integration open item 8) are a missing build tag, not missing items.** `TestHunterPetDoesNotInheritTheSharpeningStone`, `TestForeverRangedHitItemsReachRangedAttacksOnly` and `TestForeverScopesReachRangedAttacksOnly` say "Needs -tags=with_db". With `go test -tags=with_db ./sim/hunter/` the whole package passes on the lane head. TestP1Mage is probably the same. | `pet_stone_crit_test.go:13`; test run | Aligned (for other lanes) | - | - |
| X2 | APL top-level notes (not rendered) still carry the 10-31 sentence "Aimed Shot is only the pre-pull opener" and the 70707 "rank 6 (55)". Later dated lines in the same notes supersede both. This is history, not a bug. | data/curated/apl/hunter-*.json | Aligned | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations (before = $R/before.tsv):

| spec | faction | seed 1 | seed 50001 |
|---|---|---|---|
| Beast Mastery | alliance | 730.6 to 782.8, +7.14% | +7.17% |
| Beast Mastery | horde | 728.0 to 779.6, +7.09% | +7.09% |
| Marksmanship | both | unchanged (Lone Wolf, no pet) | unchanged |
| Survival | both | unchanged (Lone Wolf, no pet) | unchanged |

Bare bands, published band talents (H2 plus the rank rows), seed 1 Horde (seed 50001 and Alliance agree within 0.5 points; full table `$R/hunter/bare.tsv`):
- BM: 20 +9.0%, 30 +12.0%, 40 +10.4%, 50 +10.9%, 60 +8.6%.
- MM: 20 +9.5%, 30 +9.1%, 40 +8.5%, 50 +6.0%, 60 +5.6%.
- SV: 20 +12.5%, 30 +8.9%, 40 +6.7%, 50 +5.5%, 60 0.0% (Lone Wolf).

No commit moves a spec outside the hunter class. The core change is opt-in, and warlock demons still refuse every buff (`TestDemonReceivesNoPlayerBuffs` passes).

## Files to regenerate

- Ladder goldens (`sim/request TestRotationLadder`) for the three hunter specs: pet stats and the Aimed Shot rank lines move.
- addon-data / Data.lua and the coach starters for hunter-beast-mastery and hunter-marksmanship (the rank rows).
- dps-per-point and stat weights for the three hunters (the pet's share grew).
- The fork golden `sim/hunter/TestP1Hunter.results` was not touched. It passes as is with `-tags=with_db`.

## Measure at refresh

- Beast Mastery every band: picks and weights with the pet on Heart of the Lion (the pet is now 23% of raid damage). Check whether a more pet-heavy raid build gains.
- Marksmanship and Survival bare bands: the leveling pet is stronger. Survival bare 60 prefers the pet by about 2% (S1, owner ruling holds).

## For other lanes

- **Warlock (and the coordinator): the same client facts put Heart of the Lion on demons.** Any raid hunter's 409583 is unflagged and reaches every nearby ally. Turning it on is one line in `sim/warlock/pet.go` (`wp.TakeHeartOfTheLion()`, after `RefusePlayerBuffs`). It would move warlock numbers and possibly the demon choice, so I left it to that lane.
- Core/test harness: X1. Run the item-dependent fork tests with `-tags=with_db`.

## For the guide lane

- Optional, accurate either way: beast-mastery.md:60 and index.md:32 say the pet "takes no stat buffs from other players". That is still true, but a player would want to know that your own Heart of the Lion does reach it. Suggested addition: "Your own Heart of the Lion still reaches it."

## Guide text

Done in this lane (S 7e395c864, revised in ec74057de and the report commit): beast-mastery.md and marksmanship.md, "Levels 20 to 60", one paragraph each on Aimed Shot rank 1. `_prose.test.ts` passes.

wow-player reviewed the rendered pages (astro dev on this branch, captured at 1440; `$R/hunter/shots/*-aimed-rank*.png`):
- First pass: SHIP WITH FIXES, no MUSTs. Applied: rank 2 stays in the spellbook; the bar button upgrades itself, so drag rank 1 back; "about a quarter of the mana"; Marksmanship 50-59 "ties with rank 1".
- Second pass: **SHIP**. Its last SHOULD (make the bar-upgrade note its own sentence) is applied and needs no re-review.

## Still open

- H1 Summon Hawk (owner check). H2's in-game confirmation (pet buff bar).
- Carried, not reopened: cat vs raptor, Ranged Weapon Specialization on Serpent Sting, Lacerating Strikes on refresh, the queued Raptor Strike and the off-hand miss penalty. Pending owner: Deadly Aspects under Falcon text, level-60 book ranks, Arcane Shot's tooltip with spell power, Heart of the Lion 280/276.
- Warcraft Logs is still blocked (gcloud login).
