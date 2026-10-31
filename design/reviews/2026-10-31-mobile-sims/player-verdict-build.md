# F5 Sims build, player verdict (2026-10-31)

Verdict: HOLD (one visible bug, then SHIP).

Seen: sims-premium (ios, large ios), sims-lapsed-and, sims-rerun-states (ios, and), result-gear (ios, and), result-offline-and. The other captures were not opened.

## What reads right
- Level-60 Fury Warrior set is believable: 17 icons in character-frame order, epic purple and rare blue borders, class name in warrior tan, "+4 DPS from Howler's Furs" with "± 1.2 · 10,000 iterations · engine 1f542def3". The gold corner sits on the back slot (slot 4) and the legend explains it.
- Howler's Furs (272414, Pix Xizzix vendor) is the Fury BiS back in data/builds/1.60.1.70009/bis/warrior-fury.md; Cape of the Black Baron (13340, dungeon) is its listed alternative, 3 to 5 DPS behind in those tables. The swap and the direction are right.
- 1,043 DPS and 851.9 DPS are believable Fury figures. State words, bar text "40 of 64 combinations", Failed lines, Run again and the lapsed "See plans" line all match the site. Kind words (Top Gear, Sim) are right.

## Blocker
1. Android, rerun state "Starting": the button shows a lone "," instead of "Starting…" (sims-rerun-states-and.png, first card). iOS shows "Starting..." correctly. A player taps Run again and sees a blank button with a comma.

## Should fix
- iOS headings render in the plain system font ("Sims", "Cloak check"); Android and the mock use the condensed all-caps heading. Make them match the mock.
- Offline result: all 17 gear icons are red question marks on a result the player has seen online. If the item catalog is cached, show the real icons.
- Your sims rows say "Ran on the server" on a row that is still running or failed. Say "Server run" there, or drop the line.
- Fixture dates read "Dec 14" and "Dec 12", after today's 31 Oct. Harmless in the product, but do not screenshot them for the store listing. "Tuesday" as a title is fine as a player's own name.
- Largest text: "combina-tions" breaks mid-word in the bar caption. Let it wrap by word.

## Re-check (fix round)
Verdict: SHIP.
Seen: sims-rerun-states-and, result-offline-and, sims-premium-ios.
- Android "Starting…" now reads right (disabled, dimmed).
- Offline result shows the real cached icons, quality borders intact; the changed-slot corner is correctly absent on a plain run.
- iOS heading is now the condensed caps, dates read "14 Dec".
- Kept by spec: "Ran on the server" meta on running or failed rows (the state line above carries the outcome), and the Dec fixture dates. Accepted.
- Nit: iOS shows the small centred "Sims" bar title above the large "SIMS". Normal for iOS, but the large title shows twice; hide the bar title until it collapses.
