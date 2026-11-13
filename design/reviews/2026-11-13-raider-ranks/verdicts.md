# Raider ranks on guild settings: persona verdicts (2026-11-13)

Surface: the "Raider ranks" row in Members and access on the guild settings page
(web/src/components/guild/GuildRaiderRanks.svelte). Captures in this folder come from
web/tests/captures/raider-ranks/capture.spec.ts. All fixture names are made up.

| Persona | Round 1 | Final |
|---|---|---|
| guild-officer-reviewer | FIX | SHIP (round 2) |
| ux-designer | FIX | SHIP (round 2) |
| wow-player | FIX | SHIP (round 4) |
| ai-slop-reviewer | FAIL | PASS (round 4) |
| info-density-reviewer | FIX | SHIP (round 2) |

## Fixed from the MUSTs
- The summary follows the boxes as they change, not the saved value; no box ticked reads the addon default.
- Failed and loading states show no made-up count.
- Each rank shows its level 60 count: the setup rank lines gained `at_sixty` (characters.level >= 60),
  and the sample names list the 60s first.
- Holder lines carry their unit ("13 of 14 at 60"); one short line says what ticking does.

## Declined, with reason
- Rank names: nothing syncs them (the Blizzard roster API gives indexes only; the addon export drops
  GetGuildInfo's rank name). Needs an addon export change in its own lane.
- Class mix per rank, sync recency, item level: the level 60 count answers raid rank versus alt rank;
  the rest lives on Roster and Readiness (one job per page).
- "Tick everything below officers" shortcut; checkbox size and disabled accent (site-wide global.css).

## Open shoulds
- The failed row is taller than the loading skeleton at 340 px (ux-designer).
- "+N" names cannot be opened; long sample lines wrap at 390 (wow-player, ai-slop-reviewer).
