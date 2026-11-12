# Sim accuracy 2026-11-12: Shadow priest

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-shadow`. Site from origin/main e9eddbd7f, fork from master 4f8bb748d (the pin). Client on the site: 1.60.1.70790 with its hotfix push; live client 1.60.1.70888. Scope: Shadow, end to end: engine against the client and Blizzard's notes to 12 Nov, engine against real fights (foreverlogs.gg and our API), builds, rotations and rank rows under the adoption rule. Discipline and Holy only where a shared spell moved them (none did). Last Shadow audits: 2026-11-11 priests (`design/reviews/2026-11-11-sim-accuracy/priests.md`, whose engine-against-client table stands), rounds 24, 25, 29 and 30.

**Totals: 1 confirmed engine bug, old, shared core, fixed (W1: the wand's first shot of every run came a full wand speed late; Shadow bare 20 +5.6 to +11.3%, bare 30 +4.3 to +5.2%, bare 40 +2.4 to +3.3%, bare 50 +1.3 to +3.1%, bare 60 +0.3%, raid 60 unchanged; the mages' bare 20 +1.2 to +3.4%; warlocks never shoot). 1 data finding, new, no engine change (D1: Earthen Silk Slippers lost 18 spell damage in Blizzard's 11 Nov hotfix; the 10 Nov rank wore the old row at Shadow bare 40 and raid 60 and in six other caster BiS files; the 11 Nov data row is right and the running re-rank reads it). 1 open (T1). 14 new log checks aligned. Engine pin 4f8bb748d to fea547248.**

## Method

1. **Blizzard since 10 Nov.** "WoW: Forever Hotfixes" (t/2403040) read to post 8 (11 Nov, posted 12 Nov 00:48 UTC): the priest lines are Confounding Flash (Gnome) usable in Shadowform and Twilight Focus working in Shadowform (both 11 Nov). Neither is modelled on the DPS dummy (a racial flash heal; a pushback talent, and the Shadow fight has no pushback), so no number moves. The items line "Earthen Silk Slippers no longer give more spell damage than intended" is D1.
2. **Client and push.** The 2026-11-11 priests lane's read of 70790 against 70888 (zero priest differences, two float rows in the push) stands; the site's `hotfix-push.json` names no priest row. Re-read line by line for Shadow: `sim/priest/{mind_flay,mind_blast,shadow_word_pain,shadow_word_death,devouring_plague,vampiric_embrace,talents,shoot,starshards,mana_abilities,racials,_items,item_sets_tier1}.go`, `sim/priest/shadow/shadow_priest.go`, and the core they lean on (`core/dot.go`'s Dynamic re-snapshot under Mind Flay's ramp, `core/spell_outcome.go`'s crit multiplier under Shadowform's +100% crit damage bonus, `core/cost.go`'s school cost multiplier under Shadowform's -50%, `core/racial_touch_of_the_grave.go`, `core/wand.go`), against `data/builds/1.60.1.70790/spellconst/priest.json`, `trainables/priest.json` and `sets.json`.
3. **Real fights.** Our API's two public reports are the 11 Nov lane's (no priest of use; the list endpoint needs a sign-in, so no newer upload could be read). foreverlogs.gg now has 25 public reports (2815 to 2853): one new (2853, Shadowfang Keep, 11 Nov) with two priests at 18 to 25 (427 and 181 events, healers; nothing for Shadow). The Shadow log is still Itso (2839, Undead, 28 to 33; 4,178 events, every one read through `meta_events` with cursor paging). This lane read what the 11 Nov lane did not: the wand's timing event by event (W1, with three warlock logs beside it: Life 2820, Taza 2839, Valpuri 2847), Touch of the Grave's rate, amount, trigger and cooldown, Mind Flay's channel lengths, Shadow Word: Pain's tick interval, Shadow Vulnerability's stacks, and whether any of the 31 priests in the 25 reports ever cast Shadowfiend.
4. **Sims** on the Windows box through `tools/remote-test.sh` (`--dir sim`, `--engine-sha` for the lane's fork branch, one sim at a time, `nice -n 15`, GOMAXPROCS=5): the fork's full suite in a writable clone (`shadow/results.out` is the run; one `.results` golden moved, the mage's), and spec-breakdown before and after at the Shadow band floors and the wand levels below Mind Flay, raid 60, and one bare cell per mage and warlock spec (`shadow/measure.sh`).

## Engine against the client (1.60.1.70790 with its push)

The 11 Nov table stands (every Shadow spell, talent, racial and heal). Added this lane:

| area | engine | client or source | verdict |
|---|---|---|---|
| Shadowform's crit damage bonus | `SpellMod_CritDamageBonus_Flat` 1.0 on the Shadow school; the crit multiplier is `1 + (1.5 - 1) x CritDamageBonus` with CritDamageBonus 2, so a Shadow crit is 2.0x, ticks included | 15473 effect 4, aura 108 misc 15 amount 100 ("increasing the critical strike damage bonus of your Shadow spells by 100%") | aligned |
| Shadowform's cost | `SchoolCostMultiplier[Shadow] += -50`, applied after the spell's own cost mods (`ApplyCostModifiers`: flat, then school, then the spell's multiplier) | 15473 effect 3, aura 72 misc 32 amount -50; Mental Agility and Devouring Contagion are aura 108 spell mods | aligned (multiplicative, not additive) |
| Mind Flay's ramp over a dynamic snapshot | `Dynamic` re-takes the snapshot (base plus spell power) before every tick, then `OnTick` multiplies the whole snapshot by `ImprovedMindFlayRamp` | the log's 1.20 and 1.44 of the first tick include the spell power part (11 Nov L1) | aligned |
| Starshards (Night Elf) | 6 ticks a second apart, 0.167 a tick, 30 s category cooldown, ranks at 10/18/26/34/42/50/58 | 10797 to 19305: duration 6000, period 1000, sp 0.167, category 30000 | aligned |
| Shadowfiend 401977 | not registered as a cast (round 3: no learn level, never cast in beta) | `trainables/priest.json` still lists it as an active skill-line row at level 1; 0 casts by 31 priests in 25 launch reports (L8) | aligned, stays out |
| Touch of the Grave (Undead) | 30% a landed spell hit, 15 s cooldown, 5% of max health, never a tick | 1260201/1260198 (11 Nov read); the log's rate, cooldown and amount (L1 to L3) | aligned |
| **The wand's shots** | **held the caster a full wand speed before every shot, so a run's first shot came a wand speed after the global cooldown** | **the logs: a 0.50 s windup, the first fire at the GCD's end (W1)** | **fixed, old** |
| Tier 1 Raiments of Conviction, 5 pieces | the client row's spell mod on Devouring Plague's cooldown, registered through `clientsetbonus` | `sets.json` 2105: "Reduces the cooldown of your Devouring Plague spell by 60 sec", on a spell whose cooldown has been 24 s since 70425 | open (T1) |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| W1 | **A wand's first shot fired a full wand speed late.** `core/wand.go` modelled Shoot as a hardcast of the wand's speed, gated on the global cooldown, so after an instant the first shot landed a wand speed after the GCD ended, and after a hardcast or a channel a wand speed after it ended. The logs show a different shape: every shot is logged as a start and a fire 0.50 s apart (the windup; 95 pairs for Itso, 336 for Life, 145 for Taza, median 0.50 to 0.53 s); after an instant the first start sits 1.0 to 1.3 s after the cast, the fire at the GCD's end (27 of 65 first starts for Life and 14 of 19 for Taza are inside the 1.5 s, none earlier than 0.86 s); after a Mind Flay channel the first fire sits 0.2 to 1.1 s after it ends (reaction plus the windup); and a run's shots are the wand's speed apart (1.5 s, 162 of 280 gaps for Life, 50 of 69 for Itso). Shoot's cast time is now the windup left since the caster's last spell ended (an instant or hardcast completing, a channel's last tick), or the swing left since the last shot, whichever is longer (`shootCastTime`); the GCD gate and the hold during a run are unchanged. Red first: `TestShootCastTimeIsTheWindupLeftOrTheSwingLeft` (40d0a1d0c), green at 5ce7567af; the mage results golden moved (fea547248). Shared core: it reaches every caster that shoots (the mages when dry; warlocks Life Tap instead and never shoot in a card). | foreverlogs 2820, 2839, 2847 (`shadow/wand3.py`) | old (the Shoot cast rule of round 6) | E 5ce7567af, 8deb05faf, 458f5e2e2, fea547248 |

## Data, no engine change

| # | finding | evidence | where it lands |
|---|---|---|---|
| D1 | **Earthen Silk Slippers (254013) carried +24 spell damage and 18 healing until Blizzard's 11 Nov hotfix; the row is now 6 spell damage, 17 healing, 7 Spirit.** The site's `items/priest.json` moved with the 11 Nov nightly (4fb79c6d9 on 10 Nov still had 24/18/7; 1c0ff79f9 on 11 Nov has 6/17/7), but the published BiS was ranked on 10 Nov 08:17 from the old row, so a level-36 green is the Shadow feet pick at bare 40 and raid 60 (score 24, Omnicast Boots 0.4 DPS behind), and sits in the Balance, Elemental, three mage, Holy, Discipline and Retribution files too. The engine's `simdb.bin` row and the ranker both read the data, so nothing in the engine or the ranker is wrong; the re-rank running on CI (bis.yml) reads the 11 Nov row. | Blizzard 11 Nov; `git log` of `items/priest.json`; `bis/priest-shadow.json` bands 4, 5, 10, 11 | MR74 |

## Open

| # | finding | why open |
|---|---|---|
| T1 | **Raiments of Conviction's 5-piece reads "Reduces the cooldown of your Devouring Plague spell by 60 sec" on a 24 s cooldown.** The text is from the 3 min era (70009); if the live mod is still -60 s the set makes Devouring Plague free of cooldown, which would put five tier pieces on the raid card. Nothing published wears the set (raids-1 opens 9 Dec, no raid preset carries it), so nothing moves now. | needs the raid lane: the 70888 row for the set's spell and how `clientsetbonus` clamps a negative cooldown, then a card decision when the set is obtainable |

## Logs against the engine

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Touch of the Grave's cooldown is 15 s.** 79 procs (52 landed, 27 immune on Gnomeregan's mechanicals) across 63 encounters; the shortest gap between two procs is 15.3 s, the median 19.8 s, none under 15 s. | Itso | aligned |
| L2 | **It drains 5% of the priest's maximum health, flat.** Medians 61, 61, 60, 61, 60.5 in the first encounters and 57, 55, 56.5, 57, 56, 50 later: a fixed share of a maximum health that moves with Fortitude and food, not a level-scaled amount (it falls while the priest levels). 0 of 52 crit. | Itso | aligned (the exact max health is not in the log) |
| L3 | **It triggers on a landed cast, never on a tick.** Within 300 ms before each proc: a Mind Flay start (31), a Shadow Word: Pain cast (11), a Mind Blast hit (14), a wand hit (11), a Devouring Plague cast (1); never a lone tick. The engine's `OnSpellHitDealt` on `ProcMaskMeleeOrRanged | ProcMaskSpellDamage` is exactly that set. | Itso | aligned |
| L4 | **The wand (W1).** See the fix: windup 0.50 s, first fire at the GCD's end after an instant, 0.2 to 1.1 s after a channel, 1.5 s between the shots of a run. | Itso, Life, Taza, Valpuri | fixed |
| L5 | **Mind Flay channels run their 3 s unless the player breaks them.** 201 of 270 channels tick three times over 3.0 s; the 61 shorter ones (1 to 2 ticks) are followed within 0.4 s by nothing (58), a new Mind Flay (8), a Pain, a Mind Blast or a Fade: the player's own clips, not pushback. The card never clips. | Itso | aligned |
| L6 | **Shadow Word: Pain ticks every 3.0 s** (165 of 196 intervals at 3.0; the 1.5 s and 1.2 s ones are two Pains on two targets interleaved) and the first tick lands 3 s after the cast. | Itso | aligned |
| L7 | **Shadow Vulnerability stacks 1 to 5 and refreshes at 5** (202 applications at 1, 74 at 2, 42 at 3, 20 at 4, 10 at 5, 80 refreshes). | Itso | aligned (11 Nov L6) |
| L8 | **No priest casts Shadowfiend.** 31 priests across 25 reports, 0 Shadowfiend casts or buffs; the round 3 decision (not learnable) stands at launch. | all reports' actor lists and the three Shadow-relevant event sets | aligned |
| L9 | **Mind Blast's cast event is its completion** (damage 0.0 s after it) and its gaps bind nothing (the priest casts it when convenient); Cannibalize returns 115 to 124 mana a tick at 28 to 33; Spirit Tap only on kills. | Itso | aligned |
| L10 | **Vampiric Embrace heals 20% off every landed Shadow hit and tick, never crits** (178 heals, 310 overheals). | Itso | aligned (11 Nov L5) |

## Before and after

Band gear at the band floor (and the wand levels inside band 20), the guide's leveling talents at each level, 5,000 iterations (3,000 for the mage and warlock cells), seed 1, engine 4f8bb748d against fea547248 (the only change is W1). Relative to the old engine, with the sum of both runs' errors. The last column is the seconds the sim spent out of mana before and after.

| cell | before | after | gain | error | OOM |
|---|---|---|---|---|---|
| priest-shadow bare 20 at 10, Human | 13.39 | 14.63 | +9.3% | 0.15% | 26 to 16 s |
| priest-shadow bare 20 at 10, Undead | 14.42 | 15.72 | +9.0% | 0.14% | 24 to 14 s |
| priest-shadow bare 20 at 15, Human | 24.47 | 27.24 | +11.3% | 0.13% | 1 to 0 s |
| priest-shadow bare 20 at 15, Undead | 24.16 | 25.96 | +7.5% | 0.12% | 0 to 0 s |
| priest-shadow bare 20 at 20, Human | 38.38 | 40.83 | +6.4% | 0.14% | 90 to 91 s |
| priest-shadow bare 20 at 20, Undead | 41.09 | 43.91 | +6.9% | 0.14% | 69 to 70 s |
| priest-shadow bare 20 at 25, Human | 40.55 | 43.28 | +6.7% | 0.14% | 99 to 96 s |
| priest-shadow bare 20 at 25, Undead | 43.94 | 46.78 | +6.5% | 0.14% | 81 to 79 s |
| priest-shadow bare 20 at 27, Human | 42.73 | 45.20 | +5.8% | 0.14% | 90 to 87 s |
| priest-shadow bare 20 at 27, Undead | 46.35 | 48.95 | +5.6% | 0.14% | 70 to 66 s |
| priest-shadow bare 30 at 30, Human | 62.72 | 65.95 | +5.2% | 0.13% | 100 to 96 s |
| priest-shadow bare 30 at 30, Undead | 68.77 | 71.82 | +4.4% | 0.13% | 86 to 84 s |
| priest-shadow bare 30 at 35, Human | 69.43 | 72.85 | +4.9% | 0.14% | 102 to 99 s |
| priest-shadow bare 30 at 35, Undead | 77.24 | 80.53 | +4.3% | 0.15% | 85 to 85 s |
| priest-shadow bare 40 at 40, Human | 90.09 | 93.05 | +3.3% | 0.17% | 86 to 90 s |
| priest-shadow bare 40 at 40, Undead | 102.37 | 104.83 | +2.4% | 0.18% | 77 to 80 s |
| priest-shadow bare 50 at 50, Human | 218.65 | 225.41 | +3.1% | 0.20% | 50 to 49 s |
| priest-shadow bare 50 at 50, Undead | 251.25 | 254.59 | +1.3% | 0.21% | 33 to 33 s |
| priest-shadow bare 60 at 60, Human | 418.46 | 419.80 | +0.3% | 0.17% | 12 to 12 s |
| priest-shadow bare 60 at 60, Undead | 445.01 | 446.40 | +0.3% | 0.17% | 12 to 12 s |
| priest-shadow raid 60 at 60, Human | 782.06 | 782.06 | +0.0% | 0.10% | 0 to 0 s |
| priest-shadow raid 60 at 60, Undead | 769.79 | 769.79 | +0.0% | 0.10% | 0 to 0 s |
| mage-frost bare 20 at 20, Alliance | 50.57 | 51.19 | +1.2% | 0.24% | 44 to 44 s |
| mage-arcane bare 20 at 20, Alliance | 41.80 | 42.66 | +2.0% | 0.26% | 44 to 43 s |
| mage-fire bare 20 at 20, Alliance | 38.56 | 39.86 | +3.4% | 0.30% | 73 to 73 s |
| warlock-affliction bare 20 at 20, Alliance | 61.73 | 61.73 | +0.0% | 0.13% | 0 to 0 s |
| warlock-destruction bare 20 at 20, Alliance | 62.89 | 62.89 | +0.0% | 0.14% | 0 to 0 s |
| warlock-demonology bare 20 at 20, Alliance | 56.90 | 56.90 | +0.0% | 0.14% | 0 to 0 s |
| mage-frost bare 40 at 40, Alliance | 171.69 | 172.03 | +0.2% | 0.27% | 8 to 8 s |
| warlock-affliction bare 40 at 40, Alliance | 177.80 | 177.80 | +0.0% | 0.11% | 0 to 0 s |

Shadow's gain is largest where the wand is the filler (10 to 27, no Mind Flay yet) and where the bare sets run dry for most of the fight (30 to 50, where Shoot fills the dry half); at raid 60 the card never shoots. The mages shoot only when dry (bare 20 +1.2 to +3.4%, bare 40 +0.2%); the warlocks Life Tap and never shoot, so every warlock cell is identical. The fork's results goldens moved for the mage suite only (the Frost sets that run dry, +0.1 to +1.6%); priest, warlock and every other suite are green unchanged.

## Builds, rotations and rank rows

Nothing in the client moved, so the builds, the card and the three leveling orders stand. The rank rows were re-judged on the new engine, because a cheaper dry spell is exactly what a cheaper rank buys: every Shadow rank row's band edges (`shadow/rankcells.txt`, 109 cells), the level's bare band gear, the row's rank against the next rank and the top rank the level can train, seed 1 at 5,000 iterations, then seed 50001 for every cell whose best rank was not the row's (`shadow/rankcells-seed2.txt`, 100 cells). A row moves on 0.2 DPS or more at both seeds, the round 23 rule. **73 rows hold, 8 tie within error, 25 move** (`shadow/rankjob*.sh`, the runs are the lane's `rankjob*.out`):

| spell, level | was | now | gain, seed 1 and 50001 |
|---|---|---|---|
| Mind Blast 35, Gnome | rank 4 | rank 5 | +0.4%, +0.4% |
| Mind Blast 40, Dwarf, Night Elf, Troll | rank 4 | rank 5 (every race now) | +0.4% and +0.3%; +0.4% and +0.4%; +0.6% and +0.5% |
| Mind Blast 50, Dwarf, Night Elf | rank 5 | rank 6 (ahead of rank 7 by 0.2%) | +0.7%, +0.7% |
| Mind Blast 50, Troll | rank 5 | the top rank (7) | +0.6%, +0.5% |
| Mind Blast 50, Undead | rank 6 | the top rank (7) | +1.0%, +0.9% |
| Mind Blast 51, Dwarf, Night Elf, Troll | rank 5 | the top rank (7) | +1.6%, +1.6%, +1.1% at both seeds |
| Mind Blast 52, Dwarf; 53, Human | rank 5; rank 6 | the top rank (8) | +1.9%; +0.3% at both seeds |
| Devouring Plague 28, Troll | rank 1 | rank 2 | +0.6%, +0.5% |
| Devouring Plague 50, Human; 51, Human | rank 1 | rank 2 | +0.3%; +0.5% at both seeds |
| Devouring Plague 51, Dwarf, Night Elf, Troll | rank 2 | rank 3 | +0.4%, +0.4%, +0.2% at both seeds |
| Devouring Plague 59, Human, Undead | rank 2 | rank 3 | +0.5%, +0.9% at both seeds |
| Devouring Plague bare 60, Dwarf, Night Elf | rank 2 | the top rank (6) | +3.7%, +3.6% |
| Devouring Plague bare 60, Troll | rank 2 | rank 3 (ahead of the top rank by 0.7%) | +1.2%, +1.2% |
| Shadow Word: Death 40, Human | rank 1 | rank 2 | +0.3%, +0.3% |

Held within error (the row stays): Mind Blast 35 Troll, 36 Undead, 40 Human and Undead (rank 5 against 6); Devouring Plague 36 Dwarf and Troll, 37 Dwarf, 38 Troll, 43 Undead, 49 Dwarf, 52 Human; Shadow Word: Death 40 Night Elf. Every Shadow Word: Pain row holds (26 to 39, every race). The rows between a band's edges keep the band's rank and are re-judged with MR75. `data/curated/rank-choices.json` carries the moves with their notes; `rank_choices_priest_r23_test.go` and `rank_choices_remeasured_test.go` pin them; the guide's three rank lines say the same.

## Guide and card

The Shadow card's Shoot note said a shot "holds you for the wand's full speed, so it never fits between spells"; it now says what the shot does (cannot start during the global cooldown; once you are shooting, the first shot fires half a second after you press it and the rest come at the wand's speed until you cast again; below 28 that is most of your damage). The three mage Shoot notes carried the same sentence and say the same now. The warlock notes ("cannot start during a global cooldown") were already right. The guide's prose names no wand timing and stands. The wow-player verdict is `design/reviews/2026-11-12-sim-accuracy/shadow-wow-player.md`.

## Gates

All on the box through `tools/remote-test.sh`, engine fea547248:

- Fork: the full suite in a writable clone (`go test --tags=with_db ./sim/...`): every package green; the one results golden that moved is the mage's (regenerated in fea547248); `go test ./sim/core/ -run "Shoot|Wand"` green on the new test.
- Site: `npm run lint` (0 errors), `npm run check` (0 errors, 18 hints), `npm test` on `src/lib/sim`, `src/content/guides` (the prose guard reads the new notes and the guide) and `src/fixtures` (104 files, 2,054 tests green); `tools/apl_check.py` (every fork and site copy matches its curated rotation; the notes are the curated file's alone); `gofmt`; `go test ./rankchoices/ ./request/` whole (green, including the round 23, round 24, remeasured, level-gap and lookup pins rewritten to the moved rows, the rotations and APL checks, and `TestRotationLadder` on the seven regenerated goldens: priest-shadow, the three mages, the three warlocks); `go test ./adapter/ ./cmd/spec-breakdown/` green.
- No Playwright spec renders the Shoot note or a rank line by text, so no e2e spec was touched or run.

## Measure at refresh

Three rows added to the ledger's table:

- **MR74**, the casters' feet on the 11 Nov Earthen Silk Slippers hotfix (D1): every band whose feet pick is 254013 (Shadow bare 40 and raid 60 both factions; Balance, Elemental, the three mages, Holy, Discipline, Retribution).
- **MR75**, Shadow on the wand's first-shot windup (W1): bare bands 20 to 60, BiS (the wand slot and the mana stats), weights (Intellect, Spirit and mp5 against spell power, now that a dry spell costs less), races, the leveling walks at 10 to 27; the rank rows are re-judged below.
- **MR76**, the mages on the wand's first-shot windup (W1): bare bands 20 to 50 where the sets run dry (the wand slot, weights, walks, rank rows); raid 60 never shoots. Warlocks Life Tap and never shoot: every warlock cell is identical.

## For the owner

Nothing needs the owner. T1 (the tier 1 Devouring Plague set bonus) is for the raid lane before 9 Dec.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 6 to 8, 10 and 11 Nov, fetched 12 Nov).
- `data/builds/1.60.1.70790/spellconst/priest.json`, `trainables/priest.json`, `sets.json`, `items/priest.json` (and its history), `bis/priest-shadow.json`, `hotfix-push.json`.
- foreverlogs.gg reports 2815 to 2853 (`/api/reports/public`, `/api/reports/<id>/meta_actors`, `/api/reports/<id>/meta_events?sourceCharacterIds=<id>` with cursor paging), read 12 Nov; the scripts are in `shadow/`.
- The 2026-11-11 priests report for the full engine-against-client table.
