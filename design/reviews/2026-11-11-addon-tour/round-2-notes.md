# Round 2: what changed after the first gate, and what was declined

Boards re-rendered by `design/mocks/addon_tour.py` (heights: tour pages 329, 329, 434, 331, 282, 450, 450, 450, 450, 409, 333; dialogs 342, 363, 437, 342, 424, 382; `tour-aside` added).

## Changed (code, Locale, spec)

- Leads (ai-slop 1 to 3, wow-player S9): page 1 keeps its first sentence; page 2 reads "What you pick on the site shows up on your bags, in your talent window and in your panel."; page 3 "Each boss gets a parse and a biggest loss once the companion uploads the fight."; page 5 says "the build you pick", not "your path".
- "Welcome to Forever Sixty" under the heading on page 1 only (ai-slop tell 7, the round-3 follow-up).
- The points split "2 Holy · 2 Retribution" (`Follow.specSplit`) beside the question on the tour, the panel, the dialog's heading and the chip card's foot; "loads its leveling build" and the chip card's footer are gone (ai-slop 4, info-density 3, 6, 12).
- The next talent on each spec's build (`Follow.defaultNext`): under every tile and on every panel Pick row after the playstyle ("· Next: Divine Strength"), and at the right of every own-build row the player is not following, muted; Following stays the only gold among the rows (info-density 1, 2, 7, 9). Feral and Feral Bear are told apart by their playstyle lines and splits; their next talent differs only once one is followed.
- The chip card's spec buttons are the client recipe the tour's tiles and the panel's Pick use, no more gold slabs (ai-slop, brand 6, ux-designer, wow-player S7).
- "From the site" over the site-sent rows and the paste box, not "Custom build" (brand 4).
- Playstyle lines that named talents the shipped data does not carry: Holy "Healer, Holy Shock and Holy Light" (plate at 40), Retribution "Two-handed weapons, Seal of Command and Vengeance", Feral "Cat form, bleeds and Savage Fury" (wow-player M1, M2; every name checked against Data.lua).
- Own rows 6 px apart (ux-designer); the footer's "/fs help opens this again." at the buttons' edge (ux-designer M4).

## Changed (mock only)

- The crest on a site row is the row's class (ux-designer M1, wow-player M3, brand 1); one crest treatment at 20 and 32 (brand 2).
- The panel board draws no rotation block while nothing is followed (ux-designer M2).
- The guild board is wider, keeps its cells, and carries an undecided level 60 (ux-designer M3, guild-officer 3 and 4).
- Next talents on every board come from Data.lua's orders, never typed (wow-player S3).
- The chip pill draws the crest (ux-designer).
- `tour-aside.png`: the dialog opaque over the world with the tour stepped aside (ux-designer).

## Declined, with the reason

- "Tab to switch" and the companion line in the dialog (ai-slop 5): the owner's item 4 asks for the companion line, and the hint is the dialog's one keyboard cue.
- "path" in "Pick a path" and "Your path" (wow-player S5): the owner's own words for these lines; the lead now says "build".
- Spec icons on the own rows (wow-player S1): the data ships none, and the one-crest rule allows no other class image; the next talent now tells the rows apart.
- Raid builds totalling 50 points (wow-player M4): Data.lua is generated nightly and this lane does not touch it; handed to the data lane.
- A hover with the split on other players' empty spec cells (guild-officer 1): the hello carries the spec alone; needs data. A "Pick your path" target on the own guild row: a follow-up, not this landing.
- Guild sign-up counts on page 4 (info-density 14): needs data.
- The dialog's older chrome beside the tour's plate (brand 5): a later pass, recorded in the spec.

## Round 3 (after the second gate)

Second-gate verdicts: ux-designer SHIP WITH FIXES, wow-player SHIP WITH FIXES, ai-slop-reviewer PASS WITH FIXES, info-density-reviewer DENSE, brand-consistency-reviewer DRIFT (only the dialog's older chrome, deferred in spec 6.11). The verdicts are in `<reviewer>-round-2.md`.

Changed: the panel's split sits on its own muted line under the question and a Pick row's "Next:" is on its own line (the one shared must-fix); "coming" on page 4 is muted; page 1's site and in-game lines are nouns ("Best in slot, the sim, guides, tier lists.", "Your bags, your talent window, your panel."); the chip card's buttons centre their labels and each row carries its next talent, set on every show.

Declined: the dialog's hint inset, the tour footer's right edge and the client button's square corners are the existing recipes, not this lane's; page 3's three sublines are the round-3 copy the coordinator ruled on (hiding the score) and stay; "Next talent:" at the dialog's top names the build followed, the rows the builds on offer.
