Verdict: FIX (round 1, commit 44ed408f4)

Reads as a player's reference: real spec names and DPS, class colours, uneven row heights. Two generated tells.

Must fix
1. Meta line describes its own layout ("DPS, then the stats...") and "Raid-ready" is vague. Cut to the one fact that changes how the number reads.
2. "rating" repeated 30+ times ("Hit rating 1.00 · Crit rating 0.46"); it also wraps to orphans at 390. Use "Hit 1.00 · Crit 0.46".

Shoulds
- Ragged gaps at 1440 in short class cells.
- 30 pairs of identical gold Guide / BiS links: mute them, gold on hover.
- Rogue class name is yellow like the links.
- Phone: "All stat weights ›" alone on its own line.
