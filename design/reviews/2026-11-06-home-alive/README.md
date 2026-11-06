# Home, alive: three directions (2026-11-06)

Owner at 2000 px on the signed-in home: "the site still feels dead and dull ... boring and lifeless." Three mocks of the same character (Go Hogs, level 20 Fury Warrior, Horde, in a guild, Battle.net imported, no logs or sims). Generator `design/mocks/home_alive.py`, mocks `design/mocks/web/web-home-alive-{a,b,c}.html`, captures by `design/mocks/home_alive_shoot.mjs` (`home-alive-<x>-<w>.png` is the first screen at 1440x1100 and 2000x1200, `-full` the whole page). Upgrades, deltas, band-30 talents, trainer spells and dungeon picks are the active build's real data; anything else carries an Example tag.

**A, the character's journey** (`home-alive-a-*`)
- One focal moment: a full-bleed band tinted in class colour and Horde red, the crest at 112 px, the name in Marcellus, and the DPS figure on the right.
- A 1 to 60 track with the band markers, then what 30 brings: 4 Fury talents, 3 trainer spells, new dungeon picks from Razorfen Kraul and Scarlet Monastery.
- The upgrades become six tiles with quality glow and green delta bars. Below them, the auto-sim's top 3 abilities by damage share.
- Fixes the empty-state problem: every block is the player's own data. Costs one new system, the auto-sim on import.
- Weakness: a level 60 has no journey left, so at 60 the track needs a raid-readiness version.

**B, the world is moving** (`home-alive-b-*`)
- A compact character strip with three stats, then two countdowns: first raids on 9 Dec, with the three raid icons, and the guild's next event.
- A "Forever right now" feed with HD icons, times ago, a live dot and filters for All, Your spec, Your guild and Logs.
- In the rail: compact upgrades, your spec's tier this week, and the top movers for your spec and band.
- Feels the most "live", but most of the page is other people. That pulls toward a news site, which the positioning rules out.
- Costs the most new work: an events table, weekly ranking and tier snapshots, and a floor for quiet hours.

**C, atmosphere** (`home-alive-c-*`)
- Keeps today's layout exactly. The Fury talent-pane art is a faded backdrop behind the head, with a class-colour glow and drifting embers.
- Richer surfaces: rows with a quality-coloured edge, a glowing icon frame, the top upgrade marked New with a shimmer, and the answer counting up on load.
- Cheapest: art and CSS only. Every animation is off under prefers-reduced-motion.
- The backdrop is 300 px client art. It only holds up blurred, so it would need an HD trace like the crests got.
- Does not fix the root problem: the lower half is still invitations (Run your first sim, Upload a log).

**Recommendation: A as the base, with two pieces from B and the finish from C.**
From B, take the raid countdown (as a slim line in A's hero, beside the track) and the guild's next event in the rail. Keep the feed off home, or reduce it to "Your spec this week": tier, a BiS change and the top mover, all filtered to the player.
From C, take the surface finish (quality edges, hover glow, count-up, the New shimmer, reduced motion respected). Use the spec art only after an HD trace.
The auto-sim on import is the one new system that matters most: it turns the emptiest block on the page into the hero figure.
