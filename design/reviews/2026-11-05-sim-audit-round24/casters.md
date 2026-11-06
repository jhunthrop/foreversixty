# Sim-accuracy round 24: Casters (mage and warlock)

Site branch wt/r24-casters (on int/r24 5f25fdc04). Fork lane/r24-casters has no commits (int/r24 215970648): no engine change was needed. Build mode: no wow-player gate; `_prose.test.ts` passes.

Method for every rank row: `spec-breakdown -band <floor> -level N -preset bare -race <race> -rank <authored id>=<rank> -bis-dir $R/bis-frozen -json`, the player's talents at N (platform's tool fix), 10,000 iterations, seeds 1 and 50001, every race of the class (mage: Human, Gnome, Undead, Troll, Orc; warlock: Human, Gnome, Orc, Undead, Troll). A row is taken on a gain above 0.15% at both seeds; a race key only where the races' best ranks differ beyond error. Each new row was checked to apply with no `-rank` and that the rank was cast (per-spell ids in `-json`). Tools and raw tables: `$R/casters/tools/`, `p1.tsv` (suspect rows), `s2.tsv` (3,000-iteration sweep screen), `p3.tsv`, `p4.tsv`, `p5.tsv` (10,000-iteration fills), `combo.tsv` (Fire Blast x Pyroblast grid).

## Mage: Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Fire Blast 30 to 35 is rank 2, not rank 3; 36 to 39 stays rank 3.** Fire Blast is cast for its Heating Up stack, and on leveling mana the cheaper rank buys more of them (19 to 22 casts a fight against 14 to 19). Judged jointly with Pyroblast (F2): with the best Pyroblast rank at each level, rank 2 beats rank 3 by 0.4% to 1.5% at 30 to 35 for every race (and the new rank 4 at 30 by 2.2% to 2.8%); from 36, when Pyroblast's cheap rank frees the mana, rank 3 beats rank 2 by 0.4% to 2.4%. Alone (with Pyroblast on its default) rank 2 also won at 30 to 36. | p1.tsv, p3.tsv, combo.tsv | CONFIRMED-BUG | old (round 13 row, band-floor talents) | S c73026b21 |
| F2 | **Pyroblast had no row; the cheap ranks win from 30 to 42.** It goes out on 3 Heating Up stacks as a 1.5 s cast, so its cost matters more than its size. With the F1 Fire Blast ranks: rank 2 at 30 to 33 (+0.1% to +0.7% over rank 3), rank 3 at 34 and 35 (rank 2 -0.2% to -0.8%), rank 2 at 36 and 37 (+0.9% to +1.9% over the new rank 4, +0.1% to +0.4% over rank 3), rank 3 at 38 to 40 (+0.3% to +0.9%), rank 3 at 41 for a Human, Orc, Troll or Undead (a Gnome keeps rank 4), rank 4 at 42 when rank 5 arrives (Orc, Troll +0.2% to +0.4%, the rest tie). The two spells interact: Fire Blast 2 with Pyroblast 2 at 36 was 96.9 (Orc) against 98.2 for Fire Blast 3 with Pyroblast 2, so they were set from the joint grid. | s2.tsv, p4.tsv, p5.tsv, combo.tsv | CONFIRMED-BUG | old (the rank sweep at the levels after a new rank; no round swept Pyroblast) | S c73026b21 |
| F3 | **Scorch rank 1 from 28, not 30.** At 28 and 29, when rank 2 arrives, rank 1 wins by 0.4% to 0.7% (28, every race) and 0.2% to 0.3% (29; Gnome and Human tie). The 30 to 39 rank 1 row holds on the level's talents (rank 2 -0.3% to -2.8%); 40 to 45 rank 2 and 46 to 49 rank 3 hold. | p1.tsv, p4.tsv | CONFIRMED-BUG | old (rank sweep) | S c73026b21 |
| F4 | **Fireball rank 4 at 24 for every race but a Gnome** (+0.2% to +0.6%; a Gnome loses 0.1%), and **rank 10 at bare 60** (+0.2% to +0.3% for a Gnome, Human or Troll; Orc and Undead tie). The raid set keeps rank 11 (row is `preset: bare`). | p4.tsv | CONFIRMED-BUG | old (rank sweep) | S c73026b21 |
| F5 | Guide-r26 sim 1, Pyroblast on 2 or 1 Heating Up stacks at raid 60 (guide raid build, frozen BiS): 2 stacks -5.0% (Alliance 687.9 / 688.4 against 724.4 / 724.9; Horde 676.1 / 676.3 against 711.2 / 712.3), 1 stack -18%. fire.md:75 and card row 4 stand. | raid.tsv | Aligned | - | - |
| F6 | Guide-r26 sim 2, potion: Major Spellblasting Potion against Major Mana Potion at raid 60: -2.9% (703.5 / 703.8 against 724.4 / 724.9 Alliance; 690.2 / 691.0 against 711.2 / 712.3 Horde). The mana potion stays; fire.md :91, :97, :105 and the starter stand. | raid.tsv | Aligned | - | - |
| F7 | fire.md:121 still put Wand Specialization at 44 and 45 after the platform lane moved it to 41 and 42 (S 502efcb28). Rewritten: Wand Specialization at 41 and 42, Accelerant's last point at 43, Burning Soul at 44 and 45. | data/curated/ladder-order.json | CONFIRMED (guide text) | R24 502efcb28 (text not followed) | S c73026b21 |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Frostbolt rank 4 from 26, not 30.** When rank 5 arrives at 26, rank 4 beats it by 1.2% to 2.0% for every race at 26 to 29; the 30 to 33 row holds on the level's talents (rank 5 -0.7% to -1.7%, rank 3 -1.0% to -1.4%). 22 to 25 keep the default. | p1.tsv, p4.tsv | CONFIRMED-BUG | old (round 16 row started at the band floor) | S c73026b21 |
| A2 | Guide-r25/r26 optional sim, the leveling Blast past 3 stacks (arcane.md:74). Shipped card (past 3 while mana% is above the fight's remaining share and above 20%) against stopping at 3, leveling build, bare: on a 60 s fight stopping loses 9% at 30 to 45 and 2% to 4% at 50 to 56; on a 180 s fight it is within 0.8% at 30 and 40 but stopping gains 1.3% to 4.5% at 45 to 60; at 120 s within 0.6%. Raid 60 (raid build): stopping loses 0.6% (705.8 against 710.8 Alliance seed 1). Thresholds 0.4 and 0.6 in place of 0.2 trade one fight length for the other (3,000-iteration exploration, arc3.tsv). No single card rule wins at every fight length, so the card stays; the guide's leveling half now says what the card does on a short pull and to stop at 3 on a long dungeon boss. | arc.tsv, arc2.tsv, arc3.tsv, raid.tsv | Aligned (guide text added) | - | S c73026b21 |

## Mage: Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Frostbolt 29 to 31 by race.** The 26 to 29 rank 4 row holds at 26 to 28 (rank 5 -1.7% to -5%). At 29, when the late-20s Frost talents arrive, a Gnome (+1.2%), Undead (+0.5%) or Orc (+0.2%) casts rank 5; a Human or Troll keeps rank 4 (rank 5 -0.3% to -0.7%). At 30 and 31 a Human keeps rank 4 (+0.4% to +1.0%); every other race loses 0.3% to 1.8% on it. The 32 to 33 rank 5 row holds (rank 4 -0.7% to -2.7%). | p1.tsv, p3.tsv | CONFIRMED-BUG | old (round 16 rows) | S c73026b21 |
| R2 | Ice Barrier rank 1 beats ranks 2 to 4 at every level from 46 to 59 on the bare fight (+0.2% to +1.1%), because the engine models no damage taken (`sim/mage/ice_barrier.go:26`), so only the mana cost differs. At raid 60 rank 1 loses 0.1% to 0.2% (619.2 against 619.8). Not adopted: while leveling, mobs hit you, and the shield's size is what keeps its extra 5% Frost damage and the pushback protection up; the sim cannot price that. Web search found no Forever source that downranks it (wow-forever.gg/spotlight/ice-barrier, foreverdb.net/spell/11426, read 6 Nov 2026). | p4.tsv, raid.tsv | Aligned (model limit, no row) | - | - |

## Warlock: Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Shadow Bolt 20 to 23 is rank 2, not rank 3.** Rank 2 beats rank 3 by 0.6% to 1.1% for every race (and rank 4 more): on level 20 gear the cheap bolt keeps you casting. At 24, when rank 2 starts losing its damage to your level, a Gnome or Human casts rank 3 (+0.2% to +0.3%), a Troll loses 0.3% on it, the rest tie; rank 4 from 25. At 28, when rank 5 arrives, a Troll keeps rank 4 (+0.2% to +0.4%). The 30 to 31 rank 4 row holds (rank 5 -0.1% to -0.3%). destruction.md:128 says so. | p1.tsv, p3.tsv | CONFIRMED-BUG | old (round 19/20 row on band-floor talents) | S c73026b21 |
| D2 | Corruption rank 3 at 41 holds (rank 4 -0.1% to -0.4%). Shadow Bolt at 36 to 43 holds (36 to 39 within error or a loss for rank 5; from 40 Incinerate replaces it). | p1.tsv, p4.tsv, p5.tsv | Aligned | - | - |

## Warlock: Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Shadow Bolt rank 5 at 36 and 37** when rank 6 arrives: +0.2% to +0.5% for a Gnome, Orc, Troll or Undead (a Human ties); a Gnome keeps it at 38 (+0.2%; a Human loses 0.1% to 0.2%). Rank 6 from 38 elsewhere; 40 to 43 hold. demonology.md says so. | s2.tsv, p4.tsv, p5.tsv | CONFIRMED-BUG | old (rank sweep) | S c73026b21 |
| M2 | Shadow Bolt rows 20 to 21 (rank 3) and 28 to 30 (rank 4) hold on the level's talents (challengers -0.1% to -0.7% or within error); 22 to 27 keep the default. | p1.tsv, p3.tsv | Aligned | - | - |

## Warlock: Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Rank sweep over every ranked spell the card casts (Shadow Bolt, Immolate, Corruption, Bane of Agony, Siphon Life, Curse of the Elements, Life Tap): no challenger above 0.1% at the two levels after any new rank. | s2.tsv | Aligned | - | - |
| L2 | Guide-r26 optional sim (R26-4): raid 60, guide raid build, frozen BiS, Horde, 10,000 iterations: 7.55 Shadow Trance procs a fight, 6.43 Shadow Bolts cast, 5.88 hits at 20.2% crit, so about 1.2 Improved Shadow Bolt debuffs a fight: at most 14 of 180 seconds (8%). spec-breakdown prints no target aura uptime; the figure is casts x hits x crit x 12 s. The R26-4 wording ("when a Nightfall bolt crits ... for 12 seconds") is right. | affl-raid.txt | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Every confirmed row here was set on band-floor talents or never swept at the levels right after a new rank; all are `old`. They cluster where a cheap rank's mana carries a leveling fight (Fire's Heating Up spells, 20s Shadow Bolt). | platform-suspect-rank-rows.tsv | - | - | - |
| X2 | Blizzard notes: thread t/2360696 unchanged since 4 Nov 19:35 UTC (posts 1 to 10, as round 23 read them). New thread "WoW: Forever Hotfixes - November 5, 2026" (t/2403040, 6 Nov 02:34 UTC): no mage or warlock line. Player threads on warlock pet health (t/2402430, t/2401991, t/2402196, t/2401780) carry no blue post. | forum JSON, read 6 Nov 2026 ~09:40 UTC | Aligned | - | - |
| X3 | Logs, last looked 2026-11-06 10:10 UTC: api.foreversixty.gg/v1/reports/recent lists only z2opt2x2gs5u (20 Oct); foreverlogs.gg answers 200 with no report API; Warcraft Logs not tried (key 401, owner). | curl | OPEN | - | - |

## Before/after

Raid 60 (frozen BiS, guide raid builds, both factions, seeds 1 and 50001): unchanged for all six specs (e.g. Fire 724.4 / 724.9 Alliance, Destruction 776.3 / 776.8 Horde, identical to before.tsv). Every change is a leveling row; bare 60 moves only for Fire (Fireball rank 10, +0.2% to +0.3%). Per changed level, relative gain at both seeds (absolute at seed 1 in the tables above and the tsv files): Fire +0.2% to +2.8% (30 to 42), Arcane +1.2% to +2.0% (26 to 29), Frost +0.2% to +1.2% (29 to 31), Destruction +0.2% to +1.1% (20 to 24, 28), Demonology +0.2% to +0.5% (36 to 38). Checked to apply with no `-rank`: Fire 36 Orc (Scorch 2948, Pyroblast 12505, Fire Blast 2138 after the joint fix), Arcane 27 Troll 57.07 (7322), Frost 30 Human 81.82 (7322), Destruction 21 Orc 61.20 (695), Demonology 36 Undead 117.71 (1106), Fire 24 Undead 44.74 (3140), Fire bare 60 Human 364.45 (10150).

## Files to regenerate

- addon-data (rotations/mage-* and warlock-* rank lines and race_spell_ids follow rank-choices.json), Data.lua.
- coach starters (`persona_starter.py --all`) if any starter repeats a rank (none found by grep for these spells).
- `sim/rankchoices/rank-choices.json` was copied by hand (`make rank-choices-sync`).

## Measure at refresh

- Fire, Arcane and Frost leveling bands on refreshed band gear: the cheap-rank margins are mana-bound and move with gear.

## For other lanes

- **Moves DPS outside mage and warlock: nothing.** Only mage and warlock rows changed.
- Platform: `TestCuratedRankChoicesAreCastableDownranks` still fails on int/r24 for paladin-protection Consecration rows 49 and 58 (not casters).
- Platform: spec-breakdown prints no target aura uptime (Improved Shadow Bolt, Improved Scorch); a debuff table would settle R26-4 style questions directly.

## For the guide lane

- arcane.md:74 leveling half now states the past-3 rule and the long-fight exception (A2); fire.md:119 names the leveling ranks; frost.md, destruction.md:128 and demonology.md name theirs.
- MAGE-R26 coach Fire Blast / Pyroblast lines quote no rank; nothing to change there.

## Still open

- X3 (logs). Carried: Ice Lance's coefficient at 60 (level 56+ Frost log), Arcane Power's cost (mana log), Arcane Subtlety's resistance cut (SIM-R7-1), Demonic Brand damage, demon crit from gear (#331).
- Pending owner, not reopened: level-60 trainer book ranks, Siphon Life tick crit, warlock Orc, the Demonology Horde race.
- Ice Barrier's absorb is not modeled (R2): a leveling damage-taken model would decide its rank while leveling.

## Sources

- Client 1.60.1.70790 spellranks.json; data/curated/rank-choices.json; data/curated/apl/mage-*.json, warlock-*.json.
- https://us.forums.blizzard.com/en/wow/t/2360696 and t/2403040 (read 6 Nov 2026).
- https://wow-forever.gg/spotlight/ice-barrier/ and https://foreverdb.net/spell/11426 (read 6 Nov 2026).
- Round 26 guide-accuracy mage.md and warlock.md ("Sims requested").
