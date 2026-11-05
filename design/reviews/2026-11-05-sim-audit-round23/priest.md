# Sim-accuracy round 23: priest

Lane priest. Site wt/r23-priest from int/r23 c741fc86c. Fork lane/r23-priest from int/r23 4da7c5599 (no engine change this round). Client 1.60.1.70790.

All sims ran through `$R/sim.sh` on the frozen BiS (`$R/bis-frozen`), 10,000 iterations, seeds 1 and 50001. Bands were judged on the ladder leveling build at each level (`$R/priest/tools/ladder.py`), raid 60 on the guide build (`$R/rt`). Screens ran at 3,000 iterations, seed 1.

**Error.** At the bands, both Shadow damage and healing come out deterministic to 0.1 between the two seeds (the printed error is 0.0). As in rounds 21 and 22, a rank row is taken on a gain of 0.2 or more at both seeds. Above 50, healing differs by up to 0.4 HPS between seeds, so rows there need more than that.

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Round 22's still-open items** (P2 heals scaling with the target's level, P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew's crit, S2 Dark Inspiration's Spirit base). No new evidence. All four need a live log from a level-40+ priest. | logs checked 13:30 UTC 5 Nov (below) | OPEN | - | - |
| P2 | **Guides r23 sim 1: Prayer of Mending "the cheapest healing you can cast".** Holy, shipped list, healed per mana by spell:<br>- raid 60: Prayer of Mending 10.5 against Renew 7.0, Greater Heal 5.1, Flash Heal 3.7. It lands 3.3 heals a cast (79.1 heals from 24.3 casts).<br>- bare 40: 7.0 against Renew 4.9. Bare 44: 7.1 against 4.9. Bare 50: 8.8 against 5.8. Bare 60: 10.2 against 5.7.<br>The claim holds at every band, and the heals a cast stay at 3.3 to 3.6. | `$R/priest/holy-raid.txt`; band runs in this lane's log | Aligned | - | - |
| P3 | **Guides r23 sim 2: Shadow Word: Pain at 6 seconds left instead of 10.** At raid 60 the two are within error (+0.1 DPS). Bare 30, 40, 50 and 60 give the same number to the decimal. shadow.md:89's reason holds per global: two ticks (2 x 127 before modifiers) beat the 195 a Mind Flay does in 1.5 s. The card prints no number. | `$R/priest/g23.txt`, `swp6-bands.txt` | Aligned (no change) | - | - |
| P4 | **Guides r23 sim 3: the Discipline shield with no health gate.** Raid 60 drops from 530.4 / 526.0 to 525.9 / 521.1 (-0.9%, both seeds). The 90% gate stays. | `$R/priest/g23.txt` | Aligned | - | - |
| P5 | **Guides r23 sim 4: Holy Greater Heal rank 1 behind the pace** (the four under-35% party lines and the under-50% tank line). Raid 60 drops from 771.2 / 788.0 to 767.7 / 782.9 (-0.5% / -0.6%). The top rank stays. | `$R/priest/g23.txt` | Aligned | - | - |
| P6 | **Blizzard's Forever notes since round 22.** Thread 2360696 has ten posts. Post 10 (4 Nov 19:35 UTC) is UI only. Post 9 (3 Nov, last edited 4 Nov 19:35) is the one round 22 already read. I re-read every priest line against the engine:<br>- Shadow Word: Death backlash is 30% and comes only on a landed cast (`shadow_word_death.go:111-118`).<br>- Vampiric Embrace healing cannot crit: `GainHealth`, no crit roll.<br>- Holy Binding is modelled.<br>- Shadowform is no longer dropped by Cure Disease. The engine drops it only on a Holy cast.<br>The one line outside the priest section, the Legacy perk "The Quick and the Dead should now affect more Priest spells" (16 Oct), is an account perk. The sim has no Legacy perks. | https://us.forums.blizzard.com/en/wow/t/2360696 (JSON read 13:00 UTC 5 Nov) | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Shadow Word: Pain cast the newest rank at 26 to 29.** The rows started at 30, so from 26 the list cast rank 4, a dear rank that the mana-bound low bands cannot keep up.<br>Measured every race, level by level:<br>- 26 and 27: rank 3 beats rank 4 for every race (+3.0% to +6.4% at 26, +2.7% to +5.4% at 27).<br>- 28: Mind Flay arrives and band-20 gear runs dry for over half the fight (a Dwarf is out of mana for 106 of 180 s). Rank 2 costs a third of rank 4, stays up (7.9 casts against rank 4's 2.2), and beats rank 4 by +3.5% to +5.2% for every race. It also beats rank 3 for every race but an Undead, who ties.<br>- 29: rank 2 wins for a Dwarf, Gnome, Night Elf and Troll (+0.5% to +3.0% over rank 3). A Human (+0.4%) and an Undead (+2.5%) do better with rank 3.<br>- 30 to 39: the old rank 3 row holds against ranks 2, 4 and 5 for every race. It had no note before; it has one now.<br>**Fix:** new rows (rank 3 at 26-27, rank 2 at 28, and at 29 race by race), plus shadow.md's leveling line. | `$R/priest/swp-all.txt`; `TestRound23ShadowPainAndPlagueRankRows` failed, then passed | CONFIRMED (rank rows) | old (no round walked 26-29; the 30-39 row predates the round-16 notes) | S 91d3464e3 |
| S2 | **Devouring Plague rows re-walked** with S1's cheaper Pain, and the unnoted 30-39 rank 1 row checked level by level:<br>- 28-29: rank 1 now also wins for a Night Elf (+0.5% / +0.7%), and for a Dwarf at 29 (+0.4%). The Gnome and Troll rows stand (a Troll ties).<br>- 30-35: rank 1 holds for every race.<br>- 36: a Human gains +0.3% to +0.4% with rank 2. Rank 3, learned at 36, is behind both.<br>- 37: a Troll gains +0.6% with rank 2.<br>- 38 and 39: rank 1 holds for every race.<br>**Fix:** the rows above, and shadow.md's Plague clause. | `$R/priest/dp-rows.txt`, `dp2829.txt` | CONFIRMED (rank rows) | old | S 91d3464e3 |
| S3 | **Mind Blast rank 5 cast at 34, the level it arrives.** Rank 4 beats it at 34 for every race but a Gnome, who ties: Human +0.6%, Dwarf +0.3%, Night Elf +0.5%, Troll +0.5% to +0.7%, Undead +0.5% to +0.6%. At 35, a Gnome, Troll or Undead still gains (+0.4% to +0.9%). At 36, a Night Elf or Undead still gains (+0.3% to +0.4%). From 37, rank 5 wins. | `$R/priest/mb34.txt`; `TestRound23ShadowMindBlastAndDeathRankRows` failed, then passed | CONFIRMED (rank rows) | old | S 5e791ceba |
| S4 | **Shadow Word: Death rank 3 cast at 48 and 49.** Round 18's rank 2 row began at 50, but rank 2 already wins at 48 and 49 for every race, from +0.2% (Troll) to +1.4% (Dwarf). At 40, where rank 2 arrives, rank 1 is better for a Human (+1.0% to +1.2%), a Dwarf and a Night Elf. At 41 a Night Elf still gains with rank 1 (+0.5%). At 43 a Gnome gains with it (+0.2%).<br>**Fix:** the rank 2 row now starts at 48, plus the race rows at 40, 41 and 43, and shadow.md's line. | `$R/priest/swd.txt` | CONFIRMED (rank rows) | old (round 18 walked 50-59 only) | S 5e791ceba |
| S5 | **Mind Flay ranks** at 28 to 58, sampled at the floor, +3 and +6 after each new rank: the newest rank wins everywhere. I also checked the channel itself in a debug log: three ticks a second apart, and a resisted cast costs only the global. The breakdown's Mind Flay "Hits" count the landed cast too, so it reads 3.6 a cast. | `$R/priest/screen.txt`; throwaway engine log (not committed) | Aligned | - | - |
| S6 | **Cost stacking on Devouring Plague.** The engine adds Devouring Contagion's -50% and Mental Agility's -10%, then multiplies by Shadowform's aura-72 school multiplier: 985 x 0.4 x 0.5 = 197 a cast. A full multiply would give 222. This is the open "cost modifiers add" item (mage, rounds 14 and 15). Shadow is never short of mana at raid 60 (0.0 s out of mana), so it moves no published number. Mental Agility's client mask (SpellEffect 687728: 0x82519AEB / 0x30 / 0 / 0x40000000) leaves out Shadow Word: Death (class mask 0 / 0x2), as the engine does (Death costs 170 = 340 x 0.5). | breakdown mana table; client SpellEffect, SpellClassOptions 84290 | OPEN (shared; no number moves) | - | - |
| S7 | Vampiric Embrace cannot miss (round 22's change): 6.02 casts, 6.02 landed at raid 60. Its cost of 18 = 40 x 0.5 x 0.9. | raid breakdown | Aligned | - | - |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Power Word: Shield cast the newest rank the level it arrives.** Round 14's top-rank rule was set at 60 (rank 4 held against rank 10). Level by level on band gear, an older rank heals a little more for some races:<br>- rank 5 at 36 to 39 for a Night Elf, Undead or Troll (+0.1% to +0.4%; a Dwarf from 37);<br>- rank 5 at 40 for a Human, Undead or Troll (+0.2% to +0.4%);<br>- rank 4 at 30 for an Undead (+0.3%);<br>- rank 8 at 54 and 55 for a Human (+0.15% to +0.2%, 0.3 to 0.4 HPS).<br>Every other race and level wins or ties with the newest rank.<br>**Fix:** the rows above. The shield note drops "at the top rank" and says the top rank is for 60. discipline.md's leveling paragraph names the 36-40 rows. `TestDisciplineShieldsAtTheTopRank` takes the Human rank 5 at 40. | `$R/priest/heal-walk.txt`, `heal-walk2.txt`, `a2.txt`; `TestRound23PriestShieldAndFlashHealRankRows` failed, then passed | CONFIRMED (rank rows) | old | S 4ff713edc |
| D2 | Penance (r1 against r2 at 40 and 42, r2 against r3 at 50 and 52) and Flash Heal at each new rank: the newest rank wins. | `$R/priest/heal-screen.txt` | Aligned | - | - |
| D3 | **Round 22's rows read back.** Heal rank 1 at 22-35 and rank 3 at 36-39. The tank's Lesser Heal below 16 renders in the band-10 card (addon-data). The note's "keeps your whole healing bonus until 33" matches the engine's level-gap rule (16 + 17, or MaxLevel 21 + 12). The note reads right at every band from 16 to 39. | `addon-data.json` band 10; rank-choices | Aligned | - | - |
| D4 | Prayer of Healing rank 1 at 40-49 (no note): rank 1 keeps its whole bonus to 51 (MaxLevel 39 + 12). Not re-measured. | client SpellLevels 97502 | Aligned (by rule) | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Flash Heal rank 3 at 32 for an Undead.** Rank 2 heals +0.4% there. Every other race gains 2% to 6% with rank 3, and so does an Undead from 33.<br>**Fix:** a one-level row. Both Holy Flash Heal notes drop "at the top rank". | `$R/priest/hw-h.txt` | CONFIRMED (rank row) | old | S 4ff713edc |
| H2 | **Prayer of Healing rank 1 at 50-59** (no note; rank 1 loses 5% of its bonus a level from 52). It holds against ranks 2 and 3 at 50 to 56. At 58 and 59 rank 2 is within about 0.2% (Dwarf 59 +0.2%), inside the seed spread at that level. The row stands. | `$R/priest/poh-holy.txt` | Aligned | - | - |
| H3 | **Round 22's Heal note read at every band** against the rows (22-26 rank 1; 27 rank 1 or 2 by race; 28 rank 2, 1 or 3; 29 rank 1; 30-31 rank 1 Alliance and rank 3 Horde; 32 rank 2, 1 or 3; 33 rank 1; 34 rank 1 for a Troll, otherwise 3; 35-37 rank 3; 38-39 rank 2, 3 or 4). Every clause matches. | rank-choices, `priest-holy.json` line 30 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The "newest rank the level it arrives" gap (S1, S3, S4, D1, H1). Earlier rounds walked a rank where a row already existed, or from the band floor, and missed the levels right after a new rank arrives and before the next band. Healers' Renew and Greater Heal were walked that way already (rounds 18 and 19). | this report | - | - | - |
| X2 | **Logs:** last checked 13:30 UTC 5 Nov. `api.foreversixty.gg/v1/reports/recent` lists only the 20 October report. foreverlogs.gg/reports/public returns its shell with no reports. I did not try WCL (the key answered 401 at the round's start). | | OPEN | - | - |

## Before/after

- **Raid 60:** unchanged to the decimal at both seeds and both factions. Discipline 530.4 / 526.0, Holy 771.2 / 788.0, Shadow 724.3 / 717.5 (seed 1, Alliance / Horde). No row touches 60 (`$R/priest/after.tsv` = `$R/before.tsv`).
- **Bands, Shadow (before = `$R/sb-before`, after = lane, seed 1):**
  - 26 Human 39.3 → 41.8 (+6.4%)
  - 27 Troll 39.6 → 41.6 (+5.1%)
  - 28 Dwarf 41.4 → 43.5 (+5.1%)
  - 28 Night Elf 41.4 → 43.6 (+5.3%)
  - 29 Human 43.9 → 45.3 (+3.2%)
  - 29 Night Elf 43.7 → 45.3 (+3.7%)
  - 36 Human 67.5 → 67.8 (+0.4%)
  - 37 Troll 67.2 → 67.6 (+0.6%)
  - S3 and S4 rows: +0.2% to +1.4% where they apply.
- **Bands, Discipline and Holy:** +0.1% to +0.4% at the D1 and H1 rows' levels.
- **Other classes:** none. Every change is a priest rank row, note or guide line. No engine commit.

## Files to regenerate

- `data/builds/1.60.1.70790/addon-data.json`, then `addon/ForeverSixty/Data.lua`: the Shadow Word: Pain, Devouring Plague, Mind Blast and Shadow Word: Death rows (race_spell_ids); the Discipline shield rows and note; the Holy Flash Heal row and the two Flash Heal notes.
- TestRotationLadder golden `sim/request/testdata/ladder/priest-discipline.golden.md`: it passes on int/r23 and fails on this branch, so the D1 shield rows move it. The Shadow and Holy goldens still pass. The mage-fire and warlock goldens fail here as well; that comes from the lanes merged into int/r23, not from this lane.
- Coach starters `data/curated/coach/priest-*.{md,json}`, if they carry the shield or Flash Heal note text.
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json`, `sim/request/apl/priest-{discipline,holy}.apl.json` (notes only). The fork's `ui/healing_priest/apls/forever_{discipline,holy}.apl.json` notes were not synced: the rotation is unchanged, and apl-check ignores notes.

## Measure at refresh

- The S1 to S4, D1 and H1 rows on re-ranked band gear. Most are under 1% and ride on band-floor gear. S1's 28-29 rank 2 depends on band-20 gear running out of mana.
- Holy Prayer of Healing rank 1 against rank 2 at 58-59.
- Carried: every Heal row (round 22), the Shadow Word: Death order (round 21), Discipline Gnome against Human, the Devouring Plague 28-29 rows.

## For other lanes

- **Every lane:** X1's method. Walk the levels right after each new rank arrives, not just the band floors and the existing rows. A rank's first levels on band gear often lose to the rank below.
- **Mage / platform:** S6, cost modifiers adding (Devouring Contagion and Mental Agility add, and Shadowform multiplies after them). No priest number moves.

## For the guide lane

- Written in this lane: shadow.md:137 (Shadow Word: Pain 26-29, Devouring Plague 36-37, Mind Blast at 34, Shadow Word: Death from 48 and the rank 1 at 40); discipline.md's leveling paragraph (shield rank 5 at 36-40); the Discipline shield note ("Cast the top rank at 60; while leveling, some races keep the older rank for a few levels after a new one arrives"); and the two Holy Flash Heal notes, which drop "at the top rank". `_prose.test.ts` passes. Build mode: no wow-player gate.
- Not written: the Undead shield rank 4 at 30, the Human rank 8 at 54-55 and the Holy Undead Flash Heal rank 2 at 32 are single-race, single-level rows that only the card shows.

## Still open

- Round 20's P2 (heals scaling with the target's level), P3 (Eureka! on a live Gnome Shadow log), P4 (Lightwell Renew's crit), S2 (Dark Inspiration's Spirit base). Each needs a live level-40+ priest log.
- S6 (cost modifiers add or multiply), shared with mage.

## Sources

- Blizzard, WoW Forever Beta Development Notes, thread 2360696 (posts 1 to 10, JSON read 13:00 UTC 5 Nov): https://us.forums.blizzard.com/en/wow/t/2360696
- Client 70790: `scratchpad/b70790/70790/{SpellEffect,SpellClassOptions,SpellLevels}.csv`; `data/builds/1.60.1.70790/{spellranks,talents/priest}.json`
- Guides r23 and r24 priest reports on origin/main: `design/reviews/guide-accuracy/round23/priest.md` ("Sims requested"), `round24/priest.md`
