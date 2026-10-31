# Leftovers UX verdict

Reviewed from a-* and b-* captures at 1440 and 390 against the stage 2 set. Fixture build: missing icons and API-less states not counted.

## Verdict: SHIP

## Blockers

None.

## Closed from stage 2

- /sim strip: now a "Jump to" row inside the Your characters card, with gold ActionLink treatment (Plan, Sim, Logs, Rankings for Warrior, Forget). Reads as links at 1440 and 390.
- Planner hole: closed at 1440; the gear-not-read card is a full-width row under the head.
- Weights: "Edit the raw request" plus the hint "The JSON this run sends to the engine" reads clearly.
- Guide: prose starts at x=96 under the head; TOC at the column's right edge.
- Home: no void under a short Your upgrades at 1440 or 390.
- /setup: DONE tags in each step head at 1440 (green, right-aligned), consistent with the header chips.
- Report: NOT CREDITED tag beside the sentence; h1 wraps at 390; "Show 2 trash fights" is a full-width row with chevron.

## Shoulds

1. /planner 1440 and 390: the gear-not-read card now sits above the trees and pushes them to y~680 (1440) and ~900 (390). Trees are the tool. Collapse the card to one line, or move it under the trees.
2. /planner 1440: about 300 px of empty ground under the Import card and Share rail before the footer (trees reserve and min-h). Trim the reserve to the settled height.
3. /sim 1440: about 170 px empty under Your sims before the footer; the min-h bump (+14 / +62) over-reserves on the landing. Same fix.
4. /sim: Thrallgar appears three times on the landing (top run card, Jump to row, list row). Keep the run card and list row; drop or merge the Jump to row's repeated identity.
5. /sim 390: the Jump to links wrap to two rows with uneven spacing, and CURRENT plus SIM sit on their own indented row away from the name. Align the tag and button to the row's left edge or keep on the name line.
6. /setup 1440 and 390: the in-card character line ("Obnoxious Yell · Fury Warrior / Open in planner / Open in simulator / Forget") is small muted text with no link affordance. Use the same gold ActionLink style as the /sim Jump to row.
7. /report 1440: the NOT CREDITED tag is 9 px caps and sits tight under the meta line; give it the same 8 px gap above as below.
8. /home 1440: about 50 px of empty padding under the Look up a character hint, below the form.
9. "Not in our item data yet" and "Item <id>" copy: no capture shows it (fixture has no unknown worn items). Attach one home or planner capture with that state next pass.
