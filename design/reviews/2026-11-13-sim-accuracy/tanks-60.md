# Sim accuracy 2026-11-13: the three tanks at 60, and the tank tier list

Lane `tank-tiers`, brief `design/reviews/sim-accuracy/brief.md`. Scope: why /tiers Tanks (Alliance, raid 60, ranking c34c2647c on engine 4c595ee6b) reads Protection Warrior 228 taken a second, Feral Bear 351 and Protection Paladin 352, with "Health" at 47,812 / 38,019 / 30,330. Site branch from origin/main d730bbcac; engine pin 5515ae0ef (unchanged); client 1.60.1.70790.

**Totals: 0 engine bugs. 1 page bug fixed (the Health column was effective health, so damage taken was counted twice), 1 ranker change (the published metrics gain `health`), the tank bars and the vs #1 column fixed. Engine pin unchanged.**

## What each column measures

| column | source | what it is |
|---|---|---|
| Taken / s | `metrics.dtps` | the boss's damage on the tank a second, 3,000 iterations of the 180 s raid fight (`sim/score/tank.go`, `TankRunOf`) |
| Health (before) | `metrics.effective_health` | `health x raw / dtps`, raw = the stand-in boss's unmitigated 1,456 a second (`TankBoss.RawDPS`: 2,500 x 1.165 / 2.0). Not a hit-point pool: it already contains Taken / s |
| Health (now) | `metrics.health` | the character's maximum health with gear, talents, raid buffs and consumes (`score.MaxHealth`, the engine's own stat computation) |
| Threat / s | `metrics.tps` | threat a second |
| vs #1 | tier-list.ts | taken over the least taken, less one |
| bar | tier-row-view.ts | least taken over the row's own; before, stretched from 90% of the lowest row |

Derived from the published file (health = EH x dtps / 1,456) and confirmed by `spec-breakdown -tank` on the box: maximum health is warrior 7,484, bear 9,171, paladin 7,329. Those sit in the Classic raid-buffed range. The page showed 47,444 / 38,034 / 30,330 because effective health divides by the damage taken: the warrior's 35% less damage became 25% more "health" on top, though the bear's real pool is 23% larger. The rail note's line "effective health is only the hit-point pool going in" was wrong for the same reason.

## Method

1. **Data path.** `web/src/lib/tiers/tier-load.ts` reads `data/builds/1.60.1.70790/bis/<spec>.json` band 60 raid per faction; `tier-list.ts` sorts tanks on `metrics.dtps`; `tier-row-view.ts` draws the figures and the bar. The metrics block is `sim/cmd/leveling-bis/score_tank.go` `tankReportOf`, from `simscore.NewTankFigures` on a 3,000-iteration run of the published set; the ranker's score (effective health x (TPS/100)^0.25 x exp(-0.005 TMI)) picks the gear but is not what the page shows.
2. **Engine at 60.** Every sim on the Windows box through `tools/remote-test.sh` (`--pin-engine`, `nice -n 15`, `GOMAXPROCS=5`, one at a time): `spec-breakdown -tank` on the three raid 60 Alliance entries (3,000 iterations, seed 7, the ranker's), plus one logged iteration of each (a scratch probe, not committed) to read every boss swing's base, after-armor and after-modifier damage and every armor change.
3. **Code read** in the fork at 5515ae0ef: the enemy attack table (`core/target.go`, `core/spell_outcome.go` `outcomeEnemyMeleeWhite`), armor (`core/spell_resistances.go` `GetArmorDamageModifier`), parry haste (`core/attack.go`), Defensive Stance (`warrior/stances.go`), Thunder Clap (`warrior/thunder_clap.go`, `core/debuffs.go`), Dire Bear Form and Thick Hide (`druid/bear_form.go`, `druid/thick_hide.go`), Barkskin, Righteous Fury and Iron Creed, the three APLs' defensive cooldowns, the raid preset.
4. **Logs.** No level 60 tank log exists anywhere yet. foreverlogs.gg's 26 public reports (2815 to 2854) top out at Scarlet Monastery and Zul'Farrak (tanks at 43 or below). Our own reports table, read read-only through the Cloud SQL proxy: 7 reports, 1,178 fights, 11 boss kills (Ironforge-area dungeons, Scarlet Monastery), no tank spec rows in `fight_metrics` at 60. The raid boss melee that would settle the 60 numbers arrives with the 9 December raids. The sub-45 fight was set against 156 logged tank kills by the 11 Nov lane; nothing at 60 can be checked against a log today.

## The engine against the rules, raid 60 Alliance

| | Warrior | Bear | Paladin |
|---|---|---|---|
| maximum health | 7,484 | 9,171 (Dire Bear +1,240, Heart of the Wild +20% Stamina) | 7,329 |
| armor in the fight (from the logged after-armor ratio) | 7,809 (6,356 of it plate and shield) | 9,757 (items 1,117 x 4.6 plus bonus; Enrage's -16% for 10 s at the pull) | 8,253 plus Devotion Aura's 735 |
| boss swings a fight | 84.1 | 98.2 | 96.1 |
| miss / dodge / parry / block | 6.7 / 19.8 / 12.8 / 55.3% | 5.4 / 27.7 / 0 / 0% | 6.9 / 14.6 / 13.0 / 42.4% |
| crit / crush / hit | 0.4 / 1.7 / 3.4% | 4.6 / 15.1 / 47.3% | 2.4 / 11.3 / 9.4% |
| damage a swing, landed or not | 492 | 643 | 659 |
| damage taken a second | 230 | 351 | 352 |

Checked and right:

- **Armor.** `1 - A / (A + 400 + 85 x 63)`, 5,755 at a level 63 attacker: the Classic Era formula (TBC's 467.5 x level form is not Classic). No tank here reaches the 75% cap, so the cap the engine does not apply changes nothing.
- **The table.** One roll: miss, dodge, parry, block, crit, crush, hit. Crushing blows 15% for a three-level gap off the base defense, so gear defense does not reduce them; crit 5.6% less defense. A warrior under Forever's Shield Block (2 charges, 7 s, 5 s cooldown, the protection-warrior lane's client and log check) has miss + dodge + parry + block above 100% 83% of the time, so it is crushed 1.7% of swings; a paladin with Holy Shield up 92% still has 77% and is crushed 11%; a bear has no block or parry and takes the full 15%.
- **Dire Bear Form.** +360% on item armor (aura 142 and 466), +1,240 health, the probe's +4,021 armor on 1,117 item armor. The bear ends only 25% above the warrior's armor because the published bear set is pre-raid leather (1,117 base armor) and the warrior carries The Immovable Object (2,291 on the shield alone); at 5,755 that is 62.9% against 57.6% mitigation.
- **Damage reductions.** Defensive Stance 10% (client 7376), Barkskin 20% physical for 15 s a minute in bear (28% uptime), Righteous Fury with Improved Righteous Fury and Iron Creed on the paladin, Stoneform for the Dwarf. No Blessing of Salvation in the raid preset. Shield Wall and Last Stand fire 0 times (health never under 40%), so nothing warrior-only is counted as permanent.
- **Swing count.** The boss's swing is 2.0 s. Thunder Clap's slow is 20% in the client (aura 319, -20, every rank, 30 s at rank 6), cast 7.3 times a fight by the warrior: 84 swings. The bear's own attacks come from the front with no hit rating, so the boss parries 14% of them and its next swing comes early (parry haste, real in Classic): 98 swings. The paladin's 96 is the same effect with 5% hit.

Why the warrior takes 35% less than the bear despite less armor: 15% crushing blows against 1.7%, 5.6% boss crits against 0.4% (block pushes both off the table), the 10% stance, 13% parry against none (offset by the bear's 8% more dodge), and 17% fewer boss swings (Thunder Clap's 20% slow against the bear's parry-hasted boss). Per swing the warrior takes 492 against 643 (-23%), per second 230 against 351 (-35%). That is Classic's tank order with Forever's stronger Shield Block and Thunder Clap. The bear and paladin landing on the same figure is a coincidence of different routes (the paladin blocks and parries, the bear has more armor-to-health and dodge); Horde reads 348 and 337.

## Findings

| id | finding | status |
|---|---|---|
| P1 | The Health column showed effective health (health x raw / taken), so damage taken counted twice and the warrior won the same comparison twice; the rail note called it "the hit-point pool going in" | Fixed: the ranker publishes `metrics.health`; the column shows it, is hidden until every band has it, highest green |
| P2 | Tank bars stretched from 90% of the lowest row: 0.65 of the best rendered at 0.16 | Fixed: tank bar = least taken over the row's own, from zero (100 / 65 / 65) |
| P3 | vs #1 read "+52.9%" for more damage taken, the leader "best" though it is not best on health or threat | Fixed: "+53% taken" in whole percent, a title in words ("Takes 53% more damage a second than Protection Warrior"), the leader "lowest" |
| P4 | 351 against 352 was ranked 2 and 3 | Fixed: tanks within 1% share a rank ("=2", "≈ tie with ...") |
| P5 | On a phone no tank row showed vs #1 | Fixed: the tank row carries the phone gap figure |
| P6 | Subtitle "level 63" read as a character level | Fixed: "vs a level 63 boss" on Damage and Tanks |
| O1 | The raid preset gives the boss no Thunder Clap or Demoralizing Shout unless the tank casts them, so only a warrior tank gets the 20% slow; in a raid a second warrior would clap for a bear or paladin | Owner call (preset design) |
| O2 | Bear threat over the warrior (2,027 against 1,677) | Open, not checked here: Bear Form's 115% threat is Blizzard's 3 Nov line; the wow-player reviewer asked for a threat pass |
| O3 | Avoidance and crush chance columns on the tank list | Open SHOULD (wow-player); the figures are not published yet |

## Before and after

No sim figure moves: the engine is unchanged and nothing re-ranks here. The page, raid 60 Alliance:

| | before | after |
|---|---|---|
| Health | 47,444 / 38,034 / 30,330 (effective) | hidden until the next ranking, then 7,484 / 9,171 / 7,329 (bear green) |
| vs #1 | best / +52.9% / +53.2% | lowest / +53% taken / +53% taken |
| rank | 1 / 2 / 3 | 1 / =2 / =2 |
| bar | 100 / 16 / 16% | 100 / 65 / 65% |

## Measure at refresh

MR94 in the ledger: every tank band's `metrics.health` (all three tank specs, every band and raid 60) appears with the next ranking; the tier list's Health column comes back then.

## Reviews

- ux-designer (proposal): SHIP with changes, `design/reviews/2026-11-13-tank-tiers/ux-designer-proposal.md`.
- wow-player (proposal): SHIP with changes, `design/reviews/2026-11-13-tank-tiers/wow-player-proposal.md`. Not taken: "Damage taken / s" as the head (the heads hold to ten characters at 96 px; the rail note and the figure's label say damage taken). Avoidance and crush columns left open (O3).
- Final, on the built page (after-1440.png, after-390.png beside before-*.png): ux-designer SHIP (`ux-designer.md`), wow-player SHIP (`wow-player.md`). Their polish items stay open: one phone slot for vs #1, a sorted-column mark, phone units, one bar rule across the three tabs, avoidance and crush columns, the bear threat pass (O2).

## Gates

All on the box. Go: `go test ./cmd/leveling-bis/ -run TestTankReportPublishesMaximumHealth` red at the test commit (`report.Health undefined`), then `go vet` and `go test ./score/ ./cmd/leveling-bis/` green. Web: `vitest` over `src/lib/tiers`, `src/lib/app-data/tiers` and `src/lib/bis/tank` (141 passed), eslint and prettier on the touched files, `astro check` 0 errors after the e2e type gained `health`, `playwright test tests/e2e/tiers.spec.ts` (desktop and mobile): every tank test passes; the two failures are the DPS "one link per row" check that 1b3853d60 (a DPS figure opens the ranking sim, not this lane) leaves behind when the published files carry sim links. `src/lib/app-data/app-files.test.ts` fails on origin/main too (the shaman-elemental leveling build code), not touched here. The captures carried `health` written into the box's copy of the bis files for the run only; nothing under `bis/` is committed.

## Follow-ups, 2026-11-13 (lane `tank-followups`)

Site branch from origin/main 8083f0650; engine pin 5515ae0ef (unchanged); client 1.60.1.70790. Every sim on the Windows box (`tools/remote-test.sh --pin-engine`, `nice -n 15`, `GOMAXPROCS=5`, one at a time): `spec-breakdown -tank`, raid 60 Alliance, 3,000 iterations, seed 7.

### F1. Thunder Clap on every tank's boss (owner ruling: yes). Fixed

Before, the raid preset's debuffs (Sunder Armor, Faerie Fire, the curses, the Judgements) were the same for every spec and named no Thunder Clap, so only a warrior tank slowed the boss, with its own clap. The preset gains `debuffs_by_role` (`data/curated/presets.json`, `sim/request/preset.go` `Resolve`, validated like `debuffs`): the three tanks carry `thunder_clap`, the engine's `Debuffs.ThunderClap` at Regular, which is rank 6 (11581) at the client's 20% (aura 319, -20). Damage dealers and healers are never swung at and do not carry it. `sim/request/raid-consumes.json` (the /sim page's per-spec rows) regenerated: 23 tank rows gain the id, nothing else moved; `published-presets.json` unchanged.

No stacking. `core.ThunderClapAura` and the warrior's own `thunderClapAura` register one aura by one label (`ThunderClap-20`) with one exclusive `AtkSpdReduction` effect, so the warrior's clap refreshes the raid one instead of adding a second slow. The run shows it: the warrior's boss swings 84.1 then 83.5 times a fight, not the 70 a doubled slow would give. With the slow always up, the Protection APL's "keep it on the boss" line never fires (Thunder Clap 7.3 casts to 0); the globals go to Sunder Armor (47.4 to 54.7 casts).

| raid 60 Alliance | taken / s | TMI | threat / s | boss swings |
|---|---|---|---|---|
| Protection Warrior | 230.0 to 226.1 (-1.7%) | 23.5 to 23.0 | 1,675.4 to 1,666.8 (-0.5%) | 84.1 to 83.5 |
| Feral Bear | 351.2 to 297.3 (-15.3%) | 28.5 to 24.3 | 2,027.0 to 1,938.8 (-4.4%) | 98.2 to 83.3 |
| Protection Paladin | 351.9 to 298.4 (-15.2%) | 38.2 to 32.7 | 1,216.7 to 1,180.5 (-3.0%) | 96.1 to 81.1 |

The warrior's own small gain is the slow from the pull and through every refresh window. The bear and paladin lose threat because the slower boss gives them less rage (the bear) and fewer blocks for Holy Shield (the paladin). The warrior's lead in damage taken narrows from 35% to 24% against the bear and 24% against the paladin. Needs a re-rank: MR99.

### F2. Bear threat above the warrior. Checked: no bug in what the client or Blizzard states; the order rests on constants nothing published settles

Threat at raid 60 after F1: bear 1,939 a second, warrior 1,667. Where it comes from (threat a second): bear Maul 915 (71 casts), Lacerate 572 (78), Primal Bite 347 (29), white 73; warrior Revenge 634 (30), Shield Slam 356 (30), Heroic Strike 348 (50), Sunder Armor 145 (55), white 148. Rage, per fight (the tank report now prints it): both earn about the same, the bear 1,537 (918 from damage taken, Blood Frenzy 171, Natural Reaction 115, potion 116, white 190, Enrage 27) against the warrior's 1,584 (619 from damage taken, white 393, Shield Specialization 226, Master of Defense 137, potion 106, Bloodrage 103). The bear spends every point on threat; the warrior puts 286 of its 1,839 spent into Shield Block, which is why it takes 24% less. Threat per point of threat-spent rage is 188 for the bear and 193 for the warrior.

| item | engine | evidence | verdict |
|---|---|---|---|
| Dire Bear Form | x2.15 | 21178 aura 10 at 115 (push 113326; Blizzard 3 Nov "increased from 50% to 115%"); `BearFormPassive2BaseDamage` {115} | Aligned |
| Defensive Stance x Defiance | x1.80 x 1.20 = x2.16 | 7376 aura 10 at 80 and trait curve 109358 at 7/13/20 (push 113326; Blizzard 3 Nov); multiplied, as two MOD_THREAT auras are | Aligned |
| Feral Instinct | no threat | Forever text: Swipe damage +10% a rank and Prowl. The client row keeps Classic's shape (16947 effect 0, aura 107 on EFFECT1, +3), whose target mask our tables do not hold; if it still lifted Bear Form's threat the bear would gain, not lose | Aligned with the text |
| Faerie Fire (Feral) | 108 flat | not cast in the raid 60 rotation (0 casts), so it moves nothing here | n/a |
| Swipe | x1.75 | not cast on one target since round 21 (B3) | n/a |
| Sunder Armor | 206 + 5% of attack power | client effect 63 (11597) and the 24 Sep notes | Aligned |
| Maul | x1.75 damage | Classic's value; no Forever client column or note (cmangos gives a flat 207 instead) | Unconfirmed |
| Lacerate | x2.0 on the hit and the bleed | "a high amount of threat", no number; MR70 changed its damage (the hit counts its stack), not its threat | Unconfirmed |
| Primal Bite | x2.0 | 1 Oct "roughly doubled" | Single source |
| Shield Slam | +508 flat | "very high", no number | Unconfirmed |
| Revenge | x2.25 + 243 flat | the fork's long-standing figure, no client column | Unconfirmed |
| Heroic Strike | +145 | Classic server rows | Unconfirmed for Forever |

So the two tanks' global threat multipliers are equal to 0.5%, their rage incomes are within 3%, and the bear leads because none of its rage goes to mitigation. That is the model behaving as the client says. Whether the lead is 16% or reversed rests on the unconfirmed per-ability factors: on cmangos' flat 207 for Maul instead of x1.75 the bear would lose about 215 a second, and on x1.0 for Lacerate about 285; either alone closes most of the 272 gap, and both reverse it. Combat logs carry no threat, so no log settles them; a Forever threat meter reading (an in-game `UnitDetailedThreatSituation` capture on one boss) would. Held as O4 with that ask.

Guides: no guide, coach starter or tier-list copy says a bear out-threatens a warrior. The bear guide says Bear Form adds 115% and Primal Bite and Lacerate carry high threat; the Protection warrior guide states the stance's 80% and Defiance's 20%. Nothing changed, so no wow-player round.

### F3. The tiers e2e failure since 1b3853d60. Stale assertion, fixed

`tests/e2e/tiers.spec.ts` "dps: rows follow the published files in order" required exactly one link per row. Since 1b3853d60 a DPS row's figure is a second link, to its ranking sim (the design: spec name to BiS, figure to the sim; tanks and healers keep a plain figure), so every DPS row has two. Not a page bug. The test now finds the BiS link by its href, counts two links on a DPS row and one on a tank or healer row, and checks the DPS figure opens `/sim?`. Its CI run was green because the four-shard `e2e` job is skipped in ship-fast mode (`vars.SHIP_FAST`); only the box run caught it.

### Findings added

| id | finding | status |
|---|---|---|
| O1 | Thunder Clap on the boss only for a warrior tank | Fixed (F1), owner ruling 2026-11-13; MR99 |
| O2 | Bear threat above the warrior | Checked (F2): every client and Blizzard-stated factor aligned; the order rests on unconfirmed ability factors (O4) |
| O4 | Maul, Lacerate, Shield Slam, Revenge and Heroic Strike threat factors have no Forever number | Open: needs an in-game threat reading |
| T1 | Tiers e2e one-link check stale since 1b3853d60 | Fixed (F3) |

### Gates

On the box: `go test ./request/` red at the test commit (the three tanks without `thunder_clap`, the role list unvalidated), then `go vet` and `go test -p 1 ./request/ ./cmd/spec-breakdown/ ./cmd/leveling-bis/ ./score/` green. Web: prettier and eslint on the e2e spec, `vitest` over `src/lib/sim`, `src/lib/tiers` and `src/lib/bis`, and `playwright test tests/e2e/tiers.spec.ts` (desktop and mobile). Nothing under `bis/` regenerated.
