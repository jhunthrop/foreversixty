# Addon first-run tour (2.0.38), wow-player, round 3 (final captures in round-3/)

**Verdict: SHIP.** There are no must-fix items. The pages fit their content, page 2 shows a real item tooltip, and "coming" is honest and plain. The tour reads like an in-game panel.

## Should-fix, as raised

1. **Talent names on the last page.** "Sniper Shot" and "Lacerating Strikes" should be checked against the client tables. Both come from Data.lua and the sweep-approved `L.specPlaystyle` on main, not from the tour.
2. **"tiers" on page 1.** If it means tier lists, say "tier lists".
3. **The subtitle.** "Welcome to Forever Sixty" under every heading could go on pages 2 to 5.
4. **Gold on page 4.** Gold means both the job and "coming" there. A muted "coming" would keep gold to one meaning.
5. **The no-guild loop.** It leaves a hole at bottom left. Centre the row or end it at "Back on the site".
6. **The Tooltips only hover.** It covers the lower nodes. That is normal for a client tooltip.

## Missing next to reference surfaces

The page 2 example has no "Requires Level" line.

## Data to confirm with the accuracy reviewer

- **Impaling Harpoon** as the Beast Mastery band-20 main_hand. It is item 5200, quality 3, from Captain Greenskin at 0.3. It is in Data.lua and data/builds/1.60.1.70790/items/hunter.json, so it is a data pick and not a tour bug.
- **Lethal Attacks at 25.** Settled.
- **Advanced Combat Logging under System, Network.** Correct.
