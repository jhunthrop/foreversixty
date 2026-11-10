# Addon first-run tour (2.0.38), wow-player, round 1

Verdict: SHIP (with fixes). The chrome matches the panel. Advanced Combat Logging under System, Network and /combatlog are right for the 1.15-era client. Barrage, Sniper Shot, Lacerating Strikes and melee Survival match Forever's tree.

## Must-fix as raised, and how each was settled

- **M1.** At level 24 the board read "12 of 15 points" beside "Next point at level 25". That figure came from the mock. The Lua prints Progress (matched of allowed), and the board now reads 15 of 15. Lethal Attacks is Forever's Beast Mastery tier-1 talent 105011 (see the spec's accuracy note and data/builds/1.60.1.70790/talents/hunter.json), not the era Survival talent.
- **M2.** The job line "in chat when you zone in" and the raid-day line are built:
  - RaidAssign prints the jobs once on entering the raid.
  - The chip's Tonight line shipped in 2.0.37.
  - The panel shows the Job line.

  Only answering in game and the phone app are not live, and both carry "coming".
- **M3.** The Forever Score on page 1 is live: player tooltips have shown it since 2.0.36.

## Shoulds raised

- S1: show the tooltip example as real tooltip lines.
- S2: fix the loop's reading order.
- S3: drop the loop caption.
- S4: say "sign in with Battle.net" once.
- S5: shorten the page-2 source labels.
- S6: say the companion detects logging off, only if that is true.
- S7: confirm when hiding the score takes effect.
- S8: add "(melee)" to the Survival button.
- S9: fill the empty space on pages 1 to 4.
- S10: add a hover to Tooltips only.
