# Sim-accuracy round 24: melee (rogue and warrior)

Site `wt/r24-melee` from int/r24 26c258c20 (platform tool fix in: `-level N` uses the player's talents at N, a forced `-rank` wins). Fork `lane/r24-melee` from int/r24 215970648, unchanged. Client 1.60.1.70790. BUILD MODE.

**Result: 0 confirmed rogue or warrior sim bugs. No engine, APL, rank-row, build or guide change.** Rogue and warrior have no rows in `data/curated/rank-choices.json` and none in `platform-suspect-rank-rows.tsv`, so nothing needed a re-judge. The rank sweep finds no rank that a lower one could beat. The engine casts the newest trainer rank at every level after a new rank, in all six specs. Round 23's Assassination change (Rupture removed) holds at 10,000 iterations on both seeds and both factions. Warrior N2 (Rend and Thunder Clap attack power) stays OPEN: no log and no new source.

Tools and outputs: `$R/melee/tools/{dominance,sweep,decode}.py`, `$R/melee/out/`, `$R/melee/rot/` (R = the round-24 scratchpad). Measurements use frozen BiS (`$R/bis-frozen`) and guide builds (`$R/rt`).

## Rank sweep (both classes)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **No lower rank can win a rogue row.** I compared every trainer rank with the rank below it, at its learn level and the level after. The spells are Sinister Strike, Backstab, Eviscerate, Slice and Dice, Rupture, Garrote, Ambush, Shank (1329283 to 1329549 at 44), Mutilate (399956, 1241582 at 50, 1241584 at 60), Expose Armor, Kidney Shot, Gouge, Kick and Feint. Each step keeps the same energy cost, global, cooldown and cast time, and every effect is equal or larger: Shank's energy rises 15 to 25 for the same 5 energy, and Mutilate's flat bonus rises 39.6 to 56.7 to 79.2 a weapon at 60 energy. No rank costs more than the one below it, so no level is short enough on energy for a lower rank to win. The low-rank penalty only widens the gap. | `out/dominance-rogue.txt` on `spellconst/rogue.json`, `trainables/rogue.json`; spelltips 1329283, 1329549, 399956, 1241582, 1241584 | Aligned | - | - |
| R2 | **Warrior re-check on the current tree.** I re-ran round 23's R1 on today's 70790 rows. The one flag is Demoralizing Shout's threat effect (63), which is equal at each rank's level cap (19 against 19 at 24): a float artifact, not a weaker rank. The attack-power cut is larger at every new rank (-63 to -77 at 24, -161 to -196 at 54). | `out/dominance-warrior.txt` | Aligned | - | - |
| R3 | **The engine casts the newest trainer rank at every new-rank level and the level after.** I ran 297 bare runs on the fixed tool (`-level N`, the player's talents at N), levels 10 to 60, all six specs, Horde, 1,000 iterations. I checked each cast id against trainables (book ranks excluded) and found 0 mismatches. | `out/sweep-*.txt`, `out/low-*.txt` | Aligned | - | - |
| R4 | **The tool runs a rogue unarmed below 19** (Blackfang and Assassin's Blade need 19). At 12 a white hit averages 3.1 damage, and Slice and Dice loses to no Slice and Dice by 25 to 28%. That number says nothing about a rogue with a weapon, so I made no row or card change below 19. See "For other lanes". | `-json` at 12; `bis-frozen/rogue-assassination.json` band 20 | Tool limit (not a sim bug) | - | - |

## Rogue

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Fresh re-audit of round 23's change (Assassination drops Rupture).** I put the old Blood Splatter Rupture line back into today's card. Raid 60, raid build, 10,000 iterations, seeds 1 / 50001: Alliance card 792.8 / 792.8, with Rupture 780.6 / 780.6. Horde card 790.2 / 790.6, with Rupture 778.8 / 778.4. The card wins by 1.6% at both seeds and on both factions. The engine's Rupture matches the client and the round 23 reading: (35 + 4.73 x cp) a tick at 60, Blood Splatter adds 3 ticks, bleed ticks crit. Coup de Grace and Aggression name only Eviscerate (talents 14162, 18427), so the finisher comparison stands. Blood Splatter keeps a real effect (the opener Garrote +9 s; client 1329256), so the point is not dead. | `rot/assa-plusrupture.json`; rupture.go:94-120; talents/rogue.json | Aligned | - | - |
| K2 | **Guide-accuracy round 26 optional: Assassination's three-point Slice and Dice clause on the leveling build.** Bare, Horde, 3,000 iterations, seeds 1 / 50001, card against the four-point-only gate: 20 43.1/43.0 vs 43.1/43.2 (tie); 22 43.7 vs 43.7; 30 75.7/75.5 vs 75.4/75.2; 40 129.3/129.4 vs 129.1/129.2; 50 206.1/206.0 vs 205.4/205.1; 60 319.5/319.6 vs 318.7/318.5. The clause ties or gains 0.2 to 0.4% at every level on the leveling build, so R26-5's rule is right for both builds. The wording fix is the guide lane's. | `rot/assa-snd4only.json` | Aligned | - | - |
| K3 | Assassination casts no Eviscerate at 20 to 23 (the `expected_idle` note). Slice and Dice at three takes every set before a fourth point. Eviscerate at five placed ahead of Slice and Dice ties the card (43.1 / 44.3 at 20 / 23), and no Slice and Dice loses (42.3 / 42.8). | `rot/assa-evis5first.json`, `rot/assa-nosnd.json` | Aligned | - | - |
| K4 | Raid builds against tolls (Combat Lightning Reflexes 4 and Composure 2; Subtlety Master of Deception, Camouflage, Elusiveness): unchanged since round 23 C1/S2. Decoded with `tools/decode.py`. | talents/rogue.json | Aligned | - | - |

## Warrior

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **N2, Rend and Thunder Clap attack-power scaling (Blizzard deep dive, 30 Sep).** Re-read 6 Nov 10:30 UTC. Wowhead's Forever tooltips still show only flat numbers (Rend 11574 "147 damage over 21 sec", 11573 "126", Thunder Clap 11581 "103 damage"). No new player reading turned up, and no log exists (below). The 0-coefficient rule still needs two independent log readings. | https://nether.wowhead.com/forever/tooltip/spell/11574, /11573, /11581 | OPEN (carried: a Forever log with Rend ticks and Thunder Clap hits at two known attack powers) | old | - |
| W2 | **Guide-accuracy round 26: Protection mitigation order on one harness.** Raid 60, tank fight, guide build, Horde, 10,000 iterations, seeds 1 / 50001. Damage taken a second: card 252.8 / 253.1 (TMI 23.2 / 23.3). Without Demoralizing Shout 329.4 / 329.4 (+30.3%, TMI 31.8). Without Shield Block 311.7 / 311.9 (+23.3%, TMI 30.6). Without Thunder Clap 298.9 / 299.1 (+18.2%, TMI 27.6). So Demoralizing Shout is the largest cut, then Shield Block, then Thunder Clap. That matches the card's figures and R26-3's arithmetic. | `out/prot-mitigation.txt`, `rot/prot-*.json` | Aligned (guide-lane wording, below) | guides-handoff | - |
| W3 | **Guide-accuracy round 26: Bloodrage ceiling (R26-11).** Raid 60, 10,000 iterations, seeds 1 / 50001. Protection casts Bloodrage 3.51 times, Fury 3.53 / 3.52, out of about 4 possible in 180 s (pre-pull, then every 60 s). The ceiling (60 for Protection, 80 for Fury) costs at most half a cast a fight. spec-breakdown has no rage timeline, so the share of the fight above the ceiling and the longest stretch need the coach's F1 on a log. | sims | Aligned (F1 on a log is the coach lane's) | - | - |
| W4 | **Guide-accuracy round 26: Fury execute hold (R26-10).** Raid 60: Bloodthirst 22.1 casts against 30 possible, Whirlwind 16.4, Execute 24.0. The 36-second execute phase holds about 6 Bloodthirsts, which is most of the gap. F1 on a sim log is coach code and is not run here. | sims | Aligned (coach lane) | - | - |
| W5 | Raid builds for dead points. Fury's Arms points (Improved Rend 3, Improved Charge 2, Improved Tactical Mastery 3) are the 8-point toll to Deep Wounds and Impale; every other point in Arms rows 1 to 3 is just as idle for Fury. Arms and Protection are as in round 15 P2. | `tools/decode.py`; talents/warrior.json | Aligned | - | - |
| W6 | The layered 70790 push (80f4a3569): the only warrior row is Defensive Stance Passive 7376 (threat 80%, already law from round 21; the data lane re-checks it). No rogue row. | gr26 `hotfix-diff-70790.txt` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Blizzard's Forever notes: the thread JSON (6 Nov 10:20 UTC) still has 10 posts, and the last one is 4 Nov 19:35 UTC. No new rogue or warrior line. A web search for newer Forever hotfixes found nothing. | https://us.forums.blizzard.com/en/wow/t/2360696.json | Aligned | - | - |
| X2 | Logs, last looked 6 Nov 10:26 UTC. api.foreversixty.gg `/v1/reports/recent` still lists only the 20 Oct report. foreverlogs.gg `/api/reports/public` answers 403. Warcraft Logs v1 `/v1/zones` answers 401 "Invalid key specified" on www, classic, fresh, vanilla and sod. | curl | OPEN | - | - |
| X3 | Round 24 shared commits that touch melee: Troll Berserking autocasts only the client copy (E 248e1bdad). Rogue and warrior cards leave Berserking to autocast, so nothing moves. | fork log 3b4d6b054..215970648 | Aligned | - | - |

## Before/after

No change. Raid 60 equals `$R/before.tsv`: Assassination 792.8 Alliance / 790.2 Horde, Fury 878.2, Protection 1,453.2 threat/s (Horde, seed 1).

## Files to regenerate

None.

## Measure at refresh

- Carried: Assassination raid 60 on the re-ranked BiS without Rupture; rogue picks with one-hand axes; the Subtlety Hemorrhage gate and the 50-energy Shank on re-ranked gear; the Protection BiS on the 80% stance.

## For other lanes

- **Platform (tool):** `spec-breakdown -band 20 -level N` below the band-20 weapons' required level runs the character unarmed. For a rogue that means below 19 (Blackfang, Assassin's Blade). Any melee rank row or card judgment at 10 to 18 made this way is meaningless. Give those levels a level-legal weapon (the band-10 pick, or the best weapon in the entry's alternatives that the level allows), or have the tool refuse with a message.
- Coach: R26-10 and R26-11 need F1 run on a log (W3, W4 give the sim cast counts).

## For the guide lane

- `coach/warrior-protection.md:21` (R26-3): the harness order is Demoralizing Shout (cut 23%), then Shield Block (19%), then Thunder Clap (15%). The round 26 hand edit stands.
- `protection.md:69` "Your biggest button" is still on main (round 23 P1).
- `data/curated/apl/warrior-fury.json:41` and `:280`, the Berserker Rage note: "hold it for the fears on Magmadar, Onyxia and Nefarian". Forever's raids are Barrow Deeps, Hyjal Summit and Onyxia. Suggest leading with Onyxia (for example "hold it for Onyxia's fear") and leaving out the later bosses until their raids open. Card-note edit plus addon-data regen.
- assassination.md (R26-5): K2 shows the three-point rule holds on both builds; only the "over 20 seconds" reason is scoped by build.

## Guide text not done

- None (no sim-side change).

## Still open

- Warrior: S1 Defiance 7/13/20, S2 Revenge, Shield Slam and Thunder Clap threat, S3 Deep Wounds' double multiplier, N2 Rend and Thunder Clap attack power. Rogue: Death Mark storage, the Mutilate and Shank off-hand penalty, poison coefficients and Backstab's flat bonus. All of these need a Forever log. Owner: level-60 book ranks, WARRIOR-13, Combat Eviscerate at five.

## Sources

- Client 1.60.1.70790: `spellconst/{rogue,warrior}.json`, `spelltips/rogue.json`, `trainables/{rogue,warrior}.json`, `talents/{rogue,warrior}.json`, `hotfix-push.json` (layered, 80f4a3569), gr26 `hotfix-diff-70790.txt`.
- Wowhead Forever tooltips 11574, 11573, 11581 (6 Nov 10:30 UTC).
- Blizzard thread 2360696 (JSON, 6 Nov 10:20 UTC). Web search "WoW Forever hotfixes November 2026 rogue warrior" (6 Nov): nothing Forever-specific.
- Round 23 rogue and warrior reports; guide-accuracy round 26 rogue and warrior.
