# Sim-accuracy round 27: integration

Scoped round in ship mode, working round 26's open list. Bases: site int/r27 from main af9721524, fork int/r27 from 227d95bd1. Engine pin after: **978e4c8aa**.

## Confirmed bugs: 9

| lane | fixed | by origin | report |
|---|---|---|---|
| orders | 6 | 4 old, 2 R26 (the round 26 Skyborne druid rows) | orders.md |
| tools | 3 | 3 old | tools.md |
| **total** | **9** | **7 old, 2 R26** | |

The orders lane also fixed one stale test, `TestRestorationRankRowsBelowTwenty`, which still checked round 24's rows. It is not counted as a sim bug.

What changed, in a player's words:
- **Leveling orders** (each spec spent its first points on talents that do nothing in a fight):
  - **Subtlety:** Opportunity at 10 to 14, before Master of Deception.
  - **Arms:** Improved Rend at 10 to 12, before Deflection: +0.9 to 4.8% at 10 to 17.
  - **Assassination:** Malice at 10 to 14. Improved Gouge and Remorseless Attacks wait until 22 to 26. +6.7% at 14, +3.0% at 20; two Forever leveling guides agree.
  - **Combat:** one point of Lightning Reflexes at 14, then Precision and Puncturing Wounds by 20. +5.5% at 18 and 20 with a dagger, +2.9% at 20 with a sword.
- **Skyborne druids below 20:** a full grid on the Skyborne races (564 runs) found 10 of 40 level-race cells wrong in round 26's rows. Balance at 14 and 18 and Restoration at 10 to 16 now name their own races.
- **Combat Slice and Dice at bare 60:** no change. It loses 0.4 to 0.6% bare and breaks even at about one attack power buff, and no in-game state separates bare from raid. Recorded.
- **Tools:**
  - `spec-breakdown -race` refuses a race the class cannot be in that faction, using the race sweep's legal-race source.
  - The fork's class suites crashed only because they were run without `--tags=with_db`. Items 12620 and 12640 are in the item database, and the panic now names the missing tag. The three suites pass with the tag.
  - All 54 engine-only spell ids have an icon source in `curated/extra-spell-icons.json`, and a test keeps it that way.

## Before/after

Only the specs a fix touches were run (ship mode, 10,000 iterations, seeds 1 and 50001, both factions); the per-level tables are in orders.md. Raid 60 does not move: every order change is spent by 20, and the rank rows sit below 20. The tools changes run no sims.

## Regenerated (S 69d6e5a1a, regen790.sh at pin 978e4c8aa)

Regenerated: the Subtlety, Arms, Assassination and Combat ladder goldens, addon-builds, dps-per-point, and the build and rotation rows in Data.lua. No BiS pick moves. The fork's Restoration druid rotation notes are synced (E 978e4c8aa).

**Not regenerated: spellicons.json.** `pipeline icons` needs the fresh 70790 client tables, and wago.tools answered 503 through the round. The coordinator chose to ship without it. The icon entries for the 54 ids land as their own small change when wago is back.

## Re-rank needed (one-spec bis.yml run after landing)

- **rogue-combat band 20** and **rogue-assassination band 20**: their leveling orders change the band 20 talents. Both are in `knownStaleBISFiles` (S a57f9ce25) until the re-rank, which ships as addon 2.0.32.

## Open, for round 28

- Other specs open their leveling walks with points that do nothing in a fight: warlock-destruction, warlock-demonology, mage-frost, druid-feral, hunter-survival and paladin-protection. Each needs the same level-by-level check.
- The icon regen for the 54 ids, once wago.tools answers.
- Arms A5 (an Overpower clears its own Bloodthrill proc) needs a log.
- Logs: still only the 20 Oct test report at 23:48 UTC 6 Nov.
