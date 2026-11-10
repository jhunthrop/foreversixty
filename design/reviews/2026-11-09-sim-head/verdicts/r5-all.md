# Round 5 (captures in r5-after/)

## ai-slop-reviewer

VERDICT: FAIL

1. **BLOCKING**, head, tell 13: the `account-in` cell showed a frozen loading skeleton (dark bars under "Your account", in Characters and in Devices) at 1440 and 390. Capture after the data resolves, or label the cell as a loading state. Fixed in round 6: the capture now opens the settled account (the parity fixture's account read).
2. Non-blocking, outside the head: on bis-spec, "1083.6 DPS" in the BiS banner runs past its edge.
3. Non-blocking, head, tell 22: premium's "Plan" label over "Free".
4. Non-blocking, head, tell 21: sim-landing-in's "Press Run sim. It runs here, in your browser."
5. Non-blocking, head, tell 7: the same scope sentence under sim-gear and sim-specs.
6. Non-blocking, outside the head: the Sim tabs wrap to two rows at 390.

Landed: "± 0.3 DPS" with no "(95%)", "Saved 4 Nov", the setup download beside the answer, and no tell 23 in any head.

## ux-designer

VERDICT: SHIP WITH FIXES

No blockers. The saved head's band now sits under the figure on the buttons' line. Setup items 5 and 6 are fixed. The sim head is consistent. /bis and home are correct. The 404's smaller "Your character ›" is by spec. Polish: setup's download group sits between the title and the character block. At 390, guild-profile wraps "1 of / 1" (fixed in round 6: a Figure's value never wraps), and premium has a stranded "Plan / Free".

## wow-player

VERDICT: SHIP WITH FIXES

1. **BLOCKING**, planner body: "Send this build to Felsnap (replaces their 25 points)" ran past its "In the game" card. Fixed in round 6: the label wraps inside the card (planner.css).
2. Non-blocking: the planner shows "Matches" in green beside "Load the level 60 build" for a level-34 character.
3. Non-blocking: the planner capture shows a linked Night Elf build for Felsnap ("Opened from a link").
4. Non-blocking, fixture: level-35 warrior ranks (Heroic Strike 9, Execute 5, Whirlwind) are too high for 35.
5. Non-blocking: the home countdown reads 32 days to 9 Dec (the capture clock).
6. Non-blocking: the setup body row still says "Updated 1 h ago from Battle.net" (intended: said once, in the body).
7. Fine: "± 0.3 DPS", "Saved 4 Nov", "No parses yet", "Windshaper Skyborne".

## info-density-reviewer

VERDICT: DENSE WITH ADDITIONS

No head or page is THIN. Setup, /sim, /bis and home heads are dense enough. Agreed follow-ups, all outside the head: changelog sources, setup body copy, the sim spread explainer, the Your sims empty state, and a per-spec figure on the /bis rows.
