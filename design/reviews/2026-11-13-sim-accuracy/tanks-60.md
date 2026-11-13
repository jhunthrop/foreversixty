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
- Final verdicts on the built page: same folder.

## Gates

Go: `go test ./cmd/leveling-bis/ ./score/` on the box (red at the test commit, green after). Web: `vitest run src/lib/tiers src/lib/app-data`, `astro check`, eslint and prettier on the touched files, `playwright test tests/e2e/tiers.spec.ts` on the box.
