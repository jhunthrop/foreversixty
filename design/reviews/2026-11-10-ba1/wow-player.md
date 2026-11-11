# wow-player: BA1, rounds 1 and 2

## Round 1: FIX

Blocking: the close X, a tiny muted "×", did not read as a button; "Waist · The Deadmines · Captain Greenskin · 3…" cut the drop chance off. Also: a dangling "·" where the header meta wrapped; "your sim · 2 days ago ·" against Sim again; 10 px Open Sans to confirm at 1x; Eaglehorn Long Bow as a "World drop" for the accuracy reviewer.

Kept: rare `3d94f0` and epic `b866f5` read as their rarity on the dark plate, better than the client's `0070dd` there; blue secondary and teal primary read as clickable; the client's own Need, Greed and Pass are unchanged.

## Changes for round 2

The close X takes a hairline box in the text colour, strong on its hover ground. The truncation and wrap were mock artefacts (the mock set secondary lines at 12 px; the addon's small role is 10 px). The score is a figure in Open Sans semibold.

## Round 2: SHIP

- Close X: boxed, strong text, visible hover ground; reads as a button (`button-states.png`).
- The drop line fits in full with its unit; the header dot is gone.
- The score is a plain Open Sans figure.
- Secondary blue and primary teal are distinct; hover and held states read.
- Rarity text reads clearly as rare and epic on the umber plate.

Not blockers: the guild name was cut in the meta line ("Employ…"); fixed after this round by moving the number onto the title's line (`70e853db9`). Check the 10 px small role once in the live client at 1x. Follow-ups for other lanes: the "·" before Sim again, the "Change" link that reads like a button, a level requirement on upgrade rows, and Eaglehorn's "World drop" source.

VERDICT: SHIP
