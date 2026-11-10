# Icon lane, review round 3 (2026-11-09)

Captures: the after-* files and contact-sheet.png as committed with this file.

## ux-designer: VERDICT: HOLD

Must: the focus ring was lopsided (about 7 px clear on the icon side, 2 to 3 px after the last
letter) and left the chevron of Guides, Logs and Guild outside. Round 4 fix: each door's padding is
8 px a side (taken back by its margin, so nothing moves), giving 4 px clear inside the 2 px ring on
both sides; a door with a panel draws its ring around the door and its chevron together (the
chevron stays its own button, with its own ring when it holds focus). Round 4 captures focus on
Planner, Guides, Logs, Guild and Best in slot.

Shoulds: icons missing at 1920. Round 4: while the doors carry icons they close up 2 px a side, and
the line moves from a 1640 to a 1600 px bar, so the 1920 wide-column bar (1632 px) shows them.
The Sign out row's quiet grey tile is the account menu's existing treatment for its last, quiet
action (unchanged by this lane, same as before-account-menu-1440). The search captures are full
viewport width, not "bar width" (wording corrected here).

## wow-player: VERDICT: SHIP

Shoulds: the width rule (now 1600 px, icons at 1920 too); the Guild door and Your guild share the
banner (owner's mapping); the account tile's hood at small size is the owner's glyph at 20 px.

## ai-slop-reviewer: VERDICT: FAIL

The same focus ring defect on the chevron doors. Fixed as above.

## info-density-reviewer: VERDICT: FAIL

Musts: the ring's right gap and the chevron outside it. Fixed as above, re-shot for all three
chevron doors.
